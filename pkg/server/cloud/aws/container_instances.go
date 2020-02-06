package aws

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/api/annotations"
	"github.com/elotl/cloud-instance-provider/pkg/server/cloud"
	"github.com/elotl/cloud-instance-provider/pkg/util"
	"github.com/elotl/cloud-instance-provider/pkg/util/instanceselector"
	"github.com/elotl/cloud-instance-provider/pkg/util/sets"
	"k8s.io/klog"
)

const (
	containerStatusProvisioning   = "PROVISIONING"
	containerStatusPending        = "PENDING"
	containerStatusRunning        = "RUNNING"
	containerStatusStopped        = "STOPPED"
	maxTasksPerRequest            = 100
	eniAttachmentType             = "ElasticNetworkInterface"
	eniAttached                   = "ATTACHED"
	networkInterfaceName          = "networkInterfaceId"
	containerInstanceStartTimeout = 90 * time.Second
)

var (
	waitForRunningPollInterval = 3 * time.Second
)

func MakeFamilyPrefix(controllerID string) string {
	return "milpa-" + controllerID
}

func makeFamilyName(controllerID, podName string) string {
	// We use an underscore for the separator so that we can easily
	// get our pod name from the family name
	return MakeFamilyPrefix(controllerID) + "_" + podName
}

func SplitTaskDef(taskDef, controllerID string) (string, int) {
	parts := strings.Split(taskDef, "/")
	task := parts[len(parts)-1]
	familyPrefixUnderscore := MakeFamilyPrefix(controllerID) + "_"
	if !strings.HasPrefix(task, familyPrefixUnderscore) {
		return "", 0
	}
	task = task[len(familyPrefixUnderscore):]
	parts = strings.Split(task, ":")
	if len(parts) != 2 {
		return "", 0
	}
	name := parts[0]
	revision, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0
	}
	return name, revision
}

func (c *AwsEC2) ensureTaskLongARNEnabled() (bool, error) {
	output, err := c.ecs.ListAccountSettings(&ecs.ListAccountSettingsInput{
		EffectiveSettings: aws.Bool(true),
		Name:              aws.String(ecs.SettingNameTaskLongArnFormat),
	})
	if err != nil {
		return false, err
	}

	// This should never evaluate to true, unless there is a problem with API
	// This if block ensures that the CLI does not panic in that case
	if len(output.Settings) < 1 {
		return false, fmt.Errorf("Received unexpected response from ECS Settings API: %s", output)
	}
	value := aws.StringValue(output.Settings[0].Value)

	if aws.StringValue(output.Settings[0].Value) != "enabled" {
		output, err := c.ecs.PutAccountSetting(&ecs.PutAccountSettingInput{
			Name:  aws.String("taskLongArnFormat"),
			Value: aws.String("enabled"),
		})
		if err != nil {
			return false, util.WrapError(err, "Error enabling long ARN format for ECS, this is necessary for tagging ENS resources. To run fargate pods, you'll need to enable taskLongArnFormat for ECS. To disable fargate pods and this error, remove 'ecsClusterName' from server.yml")
		}
		value = aws.StringValue(output.Setting.Value)
	}

	return value == "enabled", nil
}

func (c *AwsEC2) EnsureContainerInstanceCluster() error {
	if c.ecs == nil {
		return fmt.Errorf("ECS is not configured on this cluster")
	}
	enabled, err := c.ensureTaskLongARNEnabled()
	if err != nil {
		return util.WrapError(err, "Could not ensure ECS is properly configured")
	}
	if !enabled {
		return fmt.Errorf("error setting up ECS cluster, task long ARN formats is not enabled")
	}

	klog.V(2).Infof("Ensuring ECS cluster %s exists", c.ecsClusterName)
	output, err := c.ecs.DescribeClusters(&ecs.DescribeClustersInput{
		Clusters: aws.StringSlice([]string{c.ecsClusterName}),
	})
	if err != nil {
		return util.WrapError(err, "Could not query ECS for Milpa cluster")
	}

	if len(output.Clusters) == 0 {
		klog.V(2).Infof("Creating ECS cluster %s", c.ecsClusterName)
		val := fmt.Sprintf("Milpa Controller %s", c.controllerID)
		tags := []*ecs.Tag{{
			Key:   aws.String("Created by Milpa Controller"),
			Value: aws.String(val),
		}}
		if c.nametag != "" {
			tags = append(tags, &ecs.Tag{
				Key:   aws.String("Cluster Name"),
				Value: aws.String(c.nametag),
			})
		}
		_, err := c.ecs.CreateCluster(
			&ecs.CreateClusterInput{
				ClusterName: aws.String(c.ecsClusterName),
				Tags:        tags,
			})

		if err != nil {
			return util.WrapError(err, "Error creating ECS Cluster")
		}
	} else if len(output.Clusters) > 1 {
		clusterNames := make([]string, 0, len(output.Clusters))
		for i := range output.Clusters {
			clusterNames[i] = aws.StringValue(output.Clusters[i].ClusterName)
		}
		return fmt.Errorf("Multiple ECS clusters matching specified milpa cluster name found: %v", clusterNames)
	} else {
		clusterStatus := aws.StringValue(output.Clusters[0].Status)
		if clusterStatus != "ACTIVE" {
			return fmt.Errorf("ECS cluster %s is not in an active state.  Status is %s. We assume this is on purpose. Milpa cannot continue if the cluster is not active", c.ecsClusterName, clusterStatus)
		}
	}
	return nil
}

// We need to wait for the task to be running so we can get an eni ID
// from it. We use the ENI to get the tasks's IP addresses. We must
// look for an ENI that's in the attached state.  If a task is stopped,
// then it'll have an ENI in the DELETED state
func getEniAndTaskStatus(task *ecs.Task) (eniID, lastStatus string) {
	lastStatus = aws.StringValue(task.LastStatus)
	for i := range task.Attachments {
		if aws.StringValue(task.Attachments[i].Type) != eniAttachmentType ||
			aws.StringValue(task.Attachments[i].Status) != eniAttached {
			continue
		}
		details := task.Attachments[i].Details
		for j := range details {
			if aws.StringValue(details[j].Name) == networkInterfaceName {
				eniID = aws.StringValue(details[j].Value)
			}
		}
	}
	return eniID, lastStatus
}

// ListTasksPages wasn't working correctly so we'll do it ourselves
func (c *AwsEC2) listTaskARNs() ([]*string, error) {
	done := false
	taskARNs := make([]*string, 0, 100)
	for !done {
		resp, err := c.ecs.ListTasks(&ecs.ListTasksInput{
			Cluster:       aws.String(c.ecsClusterName),
			LaunchType:    aws.String(ecs.LaunchTypeFargate),
			DesiredStatus: aws.String(ecs.DesiredStatusRunning),
		})
		if err != nil {
			return nil, err
		}
		taskARNs = append(taskARNs, resp.TaskArns...)
		if resp.NextToken == nil {
			done = true
		}
	}
	return taskARNs, nil
}

func (c *AwsEC2) ListContainerInstances() ([]cloud.ContainerInstance, error) {
	if c.ecs == nil {
		return []cloud.ContainerInstance{}, nil
	}
	taskARNs, err := c.listTaskARNs()
	if err != nil {
		return nil, util.WrapError(err, "Error listing fargate tasks")
	}

	tasks := make([]*ecs.Task, 0, len(taskARNs))
	queryChunks := breakIntoChunks(taskARNs, maxTasksPerRequest)
	for _, arnChunk := range queryChunks {
		resp, err := c.ecs.DescribeTasks(&ecs.DescribeTasksInput{
			Cluster: aws.String(c.ecsClusterName),
			Include: aws.StringSlice([]string{"TAGS"}),
			Tasks:   arnChunk,
		})
		if err != nil {
			return nil, util.WrapError(err, "Error describing fargate tasks")
		}
		for i := range resp.Tasks {
			if hasECSTagValue(resp.Tasks[i].Tags, cloud.ControllerTagKey, c.controllerID) {
				tasks = append(tasks, resp.Tasks[i])
			}
		}

	}

	insts := make([]cloud.ContainerInstance, len(tasks))
	for i := range tasks {
		insts[i] = cloud.ContainerInstance{
			ID: aws.StringValue(tasks[i].TaskArn),
		}
	}
	return insts, nil
}

func (c *AwsEC2) ListContainerInstancesFilterID(taskARNs []string) ([]cloud.ContainerInstance, error) {
	if c.ecs == nil {
		return []cloud.ContainerInstance{}, nil
	}
	outputARNs, err := c.listTaskARNs()
	if err != nil {
		return nil, util.WrapError(err, "Error listing fargate tasks")
	}
	filterSet := sets.NewString(taskARNs...)
	insts := make([]cloud.ContainerInstance, len(outputARNs))
	for i := range outputARNs {
		if filterSet.Has(aws.StringValue(outputARNs[i])) {
			insts = append(insts, cloud.ContainerInstance{
				ID: aws.StringValue(outputARNs[i]),
			})
		}
	}
	return insts, nil
}

func (c *AwsEC2) describeENIs(ids []string) ([]*ec2.NetworkInterface, error) {
	enis := make([]*ec2.NetworkInterface, 0, len(ids))
	err := c.client.DescribeNetworkInterfacesPages(
		&ec2.DescribeNetworkInterfacesInput{
			NetworkInterfaceIds: aws.StringSlice(ids),
		},
		func(page *ec2.DescribeNetworkInterfacesOutput, lastPage bool) bool {
			enis = append(enis, page.NetworkInterfaces...)
			return !lastPage
		})
	if err != nil {
		return nil, util.WrapError(err, "Error listing ENIs attached to tasks")
	}
	return enis, nil
}

func (c *AwsEC2) GetContainerInstancesStatuses(taskARNs []string) (map[string][]api.UnitStatus, error) {
	if c.ecs == nil {
		return map[string][]api.UnitStatus{}, nil
	}
	tasks := make([]*ecs.Task, 0, len(taskARNs))
	arns := aws.StringSlice(taskARNs)
	queryChunks := breakIntoChunks(arns, maxTasksPerRequest)
	for _, arnChunk := range queryChunks {
		resp, err := c.ecs.DescribeTasks(&ecs.DescribeTasksInput{
			Cluster: aws.String(c.ecsClusterName),
			Tasks:   arnChunk,
		})
		if err != nil {
			return nil, util.WrapError(err, "Error describing fargate tasks")
		}
		tasks = append(tasks, resp.Tasks...)
	}
	tasksToStatuses := make(map[string][]api.UnitStatus)
	for _, task := range tasks {
		taskARN := aws.StringValue(task.TaskArn)
		statuses := make([]api.UnitStatus, len(task.Containers))
		for i, container := range task.Containers {
			statuses[i] = containerToUnitStatus(container)
		}
		tasksToStatuses[taskARN] = statuses
	}
	return tasksToStatuses, nil
}

// StartedAt, FinishedAt and Image are all managed by the
// ContainerInstanceController since we don't get that info from ECS.
func containerToUnitStatus(container *ecs.Container) api.UnitStatus {
	var state api.UnitState
	reason := aws.StringValue(container.Reason)

	switch *container.LastStatus {
	case containerStatusProvisioning, containerStatusPending:
		state = api.UnitState{
			Waiting: &api.UnitStateWaiting{
				Reason: reason,
			},
		}

	case containerStatusRunning:
		state = api.UnitState{
			Running: &api.UnitStateRunning{
			// StartedAt is managed by the ContainerInstanceController
			},
		}

	case containerStatusStopped:
		var exitCode int32
		if container.ExitCode != nil {
			exitCode = int32(*container.ExitCode)
		}

		state = api.UnitState{
			Terminated: &api.UnitStateTerminated{
				ExitCode: exitCode,
				// FinishedAt is managed by the ContainerInstanceController
			},
		}
	}
	status := api.UnitStatus{
		Name:         aws.StringValue(container.Name),
		State:        state,
		RestartCount: 0,
		// Image is merged into this struct in the ContainerInstanceController
	}
	return status
}

func SecurityContextToUserGroup(sc *api.SecurityContext) *string {
	var user *string
	if sc != nil && sc.RunAsUser != nil {
		if sc.RunAsGroup != nil {
			s := fmt.Sprintf("%d:%d", *sc.RunAsUser, *sc.RunAsGroup)
			user = &s
		} else {
			s := fmt.Sprintf("%d", *sc.RunAsUser)
			user = &s
		}
	}
	return user
}

func unitToContainerDef(podname string, unit *api.Unit, emptyDirVols sets.String) *ecs.ContainerDefinition {
	containerDef := &ecs.ContainerDefinition{
		Name:       aws.String(unit.Name),
		Image:      aws.String(unit.Image),
		EntryPoint: aws.StringSlice(unit.Command),
		Command:    aws.StringSlice(unit.Args),
		User:       SecurityContextToUserGroup(unit.SecurityContext),
		// Note: Hostname parameter is not supported in VPC networking mode
	}
	if unit.WorkingDir != "" {
		containerDef.WorkingDirectory = aws.String(unit.WorkingDir)
	}
	for _, env := range unit.Env {
		containerDef.Environment = append(
			containerDef.Environment,
			&ecs.KeyValuePair{
				Name:  aws.String(env.Name),
				Value: aws.String(env.Value),
			})
	}
	for _, vol := range unit.VolumeMounts {
		if emptyDirVols.Has(vol.Name) {
			containerDef.MountPoints = append(
				containerDef.MountPoints,
				&ecs.MountPoint{
					SourceVolume:  aws.String(vol.Name),
					ContainerPath: aws.String(vol.MountPath),
				})
		}
	}
	return containerDef
}

func (c *AwsEC2) getFargateTags(pod *api.Pod) []*ecs.Tag {
	tags := []*ecs.Tag{
		&ecs.Tag{
			Key:   aws.String("Name"),
			Value: aws.String(pod.Name),
		},
		&ecs.Tag{
			Key:   aws.String(cloud.ControllerTagKey),
			Value: aws.String(c.controllerID),
		},
		&ecs.Tag{
			Key:   aws.String(cloud.NametagTagKey),
			Value: aws.String(c.nametag),
		},
	}
	filteredLabels, _ := filterLabelsForTags("fargate task", pod.Labels)
	for k, v := range filteredLabels {
		tags = append(tags, &ecs.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}
	return tags
}

func (c *AwsEC2) StartContainerInstance(pod *api.Pod) (string, error) {
	if c.ecs == nil {
		return "", fmt.Errorf("cannot start container instance: container instances client is not configured")
	}

	familyName := makeFamilyName(c.controllerID, pod.Name)
	taskDef := &ecs.RegisterTaskDefinitionInput{
		Family:                  aws.String(familyName),
		RequiresCompatibilities: []*string{aws.String(ecs.CompatibilityFargate)},
		NetworkMode:             aws.String(ecs.NetworkModeAwsvpc),
		ContainerDefinitions:    []*ecs.ContainerDefinition{},
	}
	emptyDirVols := sets.NewString()
	for _, vol := range pod.Spec.Volumes {
		if vol.EmptyDir != nil {
			emptyDirVols.Insert(vol.Name)
		}
	}
	for _, unit := range pod.Spec.Units {
		containerDef := unitToContainerDef(pod.Name, &unit, emptyDirVols)
		taskDef.ContainerDefinitions = append(taskDef.ContainerDefinitions, containerDef)
	}
	basicTags := c.getFargateTags(pod)
	taskDef.Tags = basicTags
	cpu, memory, err := instanceselector.ResourcesToContainerInstance(&pod.Spec.Resources)
	if err != nil {
		return "", util.WrapError(err, "could not find a suitable container instance size for pod resource spec")
	}
	taskDef.Cpu = aws.String(strconv.FormatInt(cpu, 10))
	taskDef.Memory = aws.String(strconv.FormatInt(memory, 10))
	if executionRole, ok := pod.Annotations[annotations.PodTaskExecutionRole]; ok {
		taskDef.ExecutionRoleArn = aws.String(executionRole)
	}
	if taskRole, ok := pod.Annotations[annotations.PodTaskRole]; ok {
		taskDef.TaskRoleArn = aws.String(taskRole)
	}

	taskDefOutput, err := c.ecs.RegisterTaskDefinition(taskDef)
	if err != nil {
		return "", util.WrapError(err, "Error registering task definition")
	}

	// Start the task.
	var taskDefARN *string
	if taskDefOutput.TaskDefinition != nil {
		taskDefARN = taskDefOutput.TaskDefinition.TaskDefinitionArn
	}
	assignPublicIPAddress := ecs.AssignPublicIpEnabled
	if pod.Spec.Resources.PrivateIPOnly {
		assignPublicIPAddress = ecs.AssignPublicIpDisabled
	}
	sgs := c.bootSecurityGroupIDs
	runTaskInput := &ecs.RunTaskInput{
		Cluster:    aws.String(c.ecsClusterName),
		Count:      aws.Int64(1),
		LaunchType: aws.String(ecs.LaunchTypeFargate),
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				AssignPublicIp: aws.String(assignPublicIPAddress),
				SecurityGroups: aws.StringSlice(sgs),
				Subnets:        []*string{&c.subnetID},
			},
		},
		PlatformVersion: aws.String("LATEST"),
		StartedBy:       aws.String(pod.UID),
		TaskDefinition:  taskDefARN,
		Tags:            basicTags,
	}

	runTaskOutput, err := c.ecs.RunTask(runTaskInput)
	if err != nil || len(runTaskOutput.Tasks) == 0 {
		if len(runTaskOutput.Failures) != 0 {
			err = fmt.Errorf("reason: %s", *runTaskOutput.Failures[0].Reason)
		}
		err = fmt.Errorf("failed to run task: %v", err)
		return "", err
	}
	if len(runTaskOutput.Tasks) < 1 {
		return "", fmt.Errorf("The cloud returned no Task info from RunTask")
	}
	return aws.StringValue(runTaskOutput.Tasks[0].TaskArn), nil
}

func (c *AwsEC2) StopContainerInstance(containerInstanceID string) error {
	if c.ecs == nil {
		return fmt.Errorf("cannot stop containerInstanceID %s: container instances client is not configured", containerInstanceID)
	}

	klog.V(2).Infof("Stopping container instance %s", containerInstanceID)
	stopTaskInput := &ecs.StopTaskInput{
		Cluster: aws.String(c.ecsClusterName),
		Reason:  aws.String("Stopped by Milpa"),
		Task:    aws.String(containerInstanceID),
	}

	// Stop the task and then Deregister Task Def
	stopTaskOutput, err := c.ecs.StopTask(stopTaskInput)
	if err != nil {
		return util.WrapError(err, "Failed to stop fargate task")
	}
	// if we got a task back, get the TaskDefinitionArn
	if stopTaskOutput.Task != nil {
		err := c.DeregisterTaskDefinition(aws.StringValue(stopTaskOutput.Task.TaskDefinitionArn))
		if err != nil {
			klog.Warningf("Error deleting task definition: %s. The task definition will be cleaned up later", err.Error())
		}
	} else {
		klog.Warningf("Task definition could not be found for %s, defering deletion of task definition", containerInstanceID)
	}
	return nil
}

func (c *AwsEC2) DeregisterTaskDefinition(taskARN string) error {
	klog.V(2).Infof("Deregistering task definition %s", taskARN)
	_, err := c.ecs.DeregisterTaskDefinition(
		&ecs.DeregisterTaskDefinitionInput{
			TaskDefinition: aws.String(taskARN),
		})
	return err
}

func (c *AwsEC2) WaitForContainerInstanceRunning(pod *api.Pod) (*api.Pod, error) {
	if c.ecs == nil {
		return nil, fmt.Errorf("Could not wait for container instance running: ECS client is not configured")
	}
	klog.V(2).Infof("Waiting for task %s to be running", pod.Status.BoundInstanceID)
	lastStatus := ""
	observedPending := false
	eniID := ""
	start := time.Now()
	for time.Since(start) < containerInstanceStartTimeout {
		describeTasksInput := &ecs.DescribeTasksInput{
			Cluster: aws.String(c.ecsClusterName),
			Tasks:   []*string{aws.String(pod.Status.BoundInstanceID)},
		}
		taskDesc, err := c.ecs.DescribeTasks(describeTasksInput)
		if err != nil {
			return nil, util.WrapError(err, "Error waiting for container instance to be running")
		}
		if len(taskDesc.Tasks) > 0 {
			eniID, lastStatus = getEniAndTaskStatus(taskDesc.Tasks[0])
		}
		if eniID != "" && lastStatus == containerStatusRunning {
			break
		} else if lastStatus == containerStatusPending {
			observedPending = true
		} else if lastStatus == containerStatusStopped && observedPending {
			// We protect against an eventual consistency error here:
			// Failure case: An old stopped task could be in STOPPED
			// state when we first start this call. Only acknowledge a
			// STOPPED task if we've first seen the task in PENDING. Worst
			// case scenerio, we timeout waiting for the task to start
			return nil, fmt.Errorf("Task has been stopped")
		}
		time.Sleep(waitForRunningPollInterval)
	}
	if lastStatus == containerStatusStopped {
		return nil, fmt.Errorf("task %s is in STOPPED state", pod.Name)
	} else if eniID == "" {
		return nil, fmt.Errorf("timed out waiting for ENI for pod %s", pod.Name)
	}
	if !pod.Spec.Resources.PrivateIPOnly {
		addys, err := c.getENIAddresses(eniID)
		if err != nil {
			klog.Errorf("Error getting addresses from cloud for pod %s: %s", pod.Name, err.Error())
		} else {
			pod.Status.Addresses = append(pod.Status.Addresses, addys...)
		}
	}
	return pod, nil
}

func (c *AwsEC2) getENIAddresses(eniID string) ([]api.NetworkAddress, error) {
	enis, err := c.describeENIs([]string{eniID})
	var addresses []api.NetworkAddress
	if err != nil {
		return nil, err
	} else if len(enis) > 0 {
		addresses = api.NewNetworkAddresses(
			aws.StringValue(enis[0].PrivateIpAddress),
			aws.StringValue(enis[0].PrivateDnsName),
		)
		if enis[0].Association != nil &&
			enis[0].Association.PublicIp != nil {
			addresses = api.SetPublicAddresses(
				aws.StringValue(enis[0].Association.PublicIp),
				aws.StringValue(enis[0].Association.PublicDnsName),
				addresses)
		}
	}
	return addresses, nil
}

func hasECSTagValue(tags []*ecs.Tag, key, value string) bool {
	for i := range tags {
		if aws.StringValue(tags[i].Key) == key &&
			aws.StringValue(tags[i].Value) == value {
			return true
		}
	}
	return false
}

func (c *AwsEC2) ListTaskDefinitions() ([]string, error) {
	if c.ecs == nil {
		return []string{}, nil
	}
	familyPrefix := MakeFamilyPrefix(c.controllerID)
	params := &ecs.ListTaskDefinitionsInput{}
	taskDefArns := make([]string, 0, 50)
	err := c.ecs.ListTaskDefinitionsPages(params,
		func(page *ecs.ListTaskDefinitionsOutput, lastPage bool) bool {
			for i := range page.TaskDefinitionArns {
				td := aws.StringValue(page.TaskDefinitionArns[i])
				parts := strings.Split(td, "/")
				if strings.HasPrefix(parts[len(parts)-1], familyPrefix) {
					taskDefArns = append(taskDefArns, td)
				}
			}
			return !lastPage
		})
	if err != nil {
		return nil, err
	}
	return taskDefArns, nil
}
