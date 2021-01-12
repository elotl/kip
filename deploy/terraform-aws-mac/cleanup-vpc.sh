#!/bin/bash
#
# Remove leftover cloud resources from an AWS VPC.
#

function usage() {
    {
        echo "Usage $0 <vpc-id> <cluster_name>"
        echo "You can also set the environment variables"
        echo "VPC_ID and CLUSTER_NAME."
    } >&2
    exit 1
}

function check_prg() {
    $1 --version || {
        {
            echo "Can't find $prg."
        } >&2
        exit 2
    }
}

if [[ "$1" != "" ]]; then
    VPC_ID="$1"
fi
if [[ -z "$VPC_ID" ]]; then
    usage
fi
shift

if [[ "$1" != "" ]]; then
    CLUSTER_NAME="$1"
fi
if [[ -z "$CLUSTER_NAME" ]]; then
    usage
fi
shift

if [[ -n "$1" ]]; then
    usage
fi

check_prg aws
check_prg jq

# Delete instances in VPC.
get_pending_and_running_instance_ids_by_vpc_id() {
    aws ec2 describe-instances --filters \
        'Name=instance-state-name,Values=pending,running' \
        "Name=vpc-id,Values=${1}" |
    jq -r '.Reservations | .[] | .Instances | .[] | .InstanceId'
}

instance_ids=$(get_pending_and_running_instance_ids_by_vpc_id "$VPC_ID")
while [[ -n "$instance_ids" ]]
do
    echo "Terminating instances: $instance_ids"
    aws ec2 terminate-instances --instance-ids "$instance_ids"
    aws ec2 wait instance-terminated --instance-ids "$instance_ids"

    # Update the list to ensure nothing else came up in the VPC that would
    # block the destruction of the security group.
    instance_ids=$(get_pending_and_running_instance_ids_by_vpc_id "$VPC_ID")
done

# Delete LBs.
lbs=$(aws elb describe-load-balancers | jq -r ".LoadBalancerDescriptions | .[] | select(.VPCId==\"$VPC_ID\") | .LoadBalancerName")
if [[ -n "$lbs" ]]; then
    echo "Removing LBs:"
    echo "$lbs"
    for lb in $lbs; do
        aws elb delete-load-balancer --load-balancer-name $lb
    done
fi
v2lbs=$(aws elbv2 describe-load-balancers | jq -r ".LoadBalancers | .[] | select(.VpcId==\"$VPC_ID\") | .LoadBalancerArn")
if [[ -n "$v2lbs" ]]; then
    echo "Removing v2 LBs:"
    echo "$v2lbs"
    for lb in $v2lbs; do
        aws elbv2 delete-load-balancer --load-balancer-arn $lb
    done
fi

# Delete security groups in VPC. This doesn't include the default security
# group that we can't delete. It will be deleted when the VPC gets deleted.
get_security_group_ids_by_vpc_id() {
    aws ec2 describe-security-groups \
        --filters 'Name=vpc-id,Values='"$1" \
        --query 'SecurityGroups[?GroupName != `default`].GroupId' \
        --output text
}

security_group_ids=$(get_security_group_ids_by_vpc_id "$VPC_ID")
while [[ -n "$security_group_ids" ]]
do
    for sg_id in $security_group_ids
    do
        echo "Deleting security group: $sg_id"
        aws ec2 delete-security-group --group-id "$sg_id"
    done
    # There may be dependant resources preventing some security groups from
    # being deleted. We retry the loop if that's the case.
    security_group_ids=$(get_security_group_ids_by_vpc_id "$VPC_ID")
done

# Delete volumes created by this cluster.
vols=$(aws ec2 describe-volumes --filters "Name=tag:kubernetes.io/cluster/$CLUSTER_NAME,Values=owned" "Name=status,Values=creating,available" | jq -r ".Volumes | .[] | .VolumeId")
if [[ -n "$vols" ]]; then
    echo "Removing volumes:"
    echo "$vols"
    for vol in $vols; do
        aws ec2 delete-volume --volume-id $vol
    done
fi

# This is needed just in case the last command executed failed
exit 0
