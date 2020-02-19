// Copyright 2014 The Kubernetes Authors.
// Copyright 2018 Elotl inc.

package milpactl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/api"
	"github.com/elotl/cloud-instance-provider/pkg/labels"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

// Most of this was taken from k8s at commit
// 2296108886a29db5cb7be73412eb562cdbf1cb74 and then modified to suit
// our own structures.
const (
	tabwriterMinWidth     = 10
	tabwriterWidth        = 4
	tabwriterPadding      = 3
	tabwriterPadChar      = ' '
	tabwriterFlags        = 0
	max_pod_reason_length = 50
	//loadBalancerWidth = 16
)

func GetPrinter(cmd *cobra.Command) (ResultPrinter, error) {
	outputFormat, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatalf("Error accessing 'output' flag for command %s: %v", cmd.Name(), err)
	}
	switch outputFormat {
	case "json":
		return &JSONPrinter{}, nil
	case "yaml":
		return &YAMLPrinter{}, nil
	case "":
		fallthrough
	default:
		return NewPrettyPrinter(false, false, false, false, false, []string{}), nil
	}
	return nil, fmt.Errorf("Unknown output printer type: %s", outputFormat)
}

type ResultPrinter interface {
	PrintObj(api.MilpaObject, io.Writer) error
}

type JSONPrinter struct{}

func (p *JSONPrinter) PrintObj(obj api.MilpaObject, w io.Writer) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	dst := bytes.Buffer{}
	err = json.Indent(&dst, data, "", "    ")
	dst.WriteByte('\n')
	_, err = w.Write(dst.Bytes())
	return err
}

type YAMLPrinter struct{}

func (p *YAMLPrinter) PrintObj(obj api.MilpaObject, w io.Writer) error {
	output, err := yaml.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, string(output))
	return err
}

type handlerEntry struct {
	columns   []string
	printFunc reflect.Value
}

type PrintOptions struct {
	NoHeaders          bool
	WithNamespace      bool
	WithKind           bool
	Wide               bool
	ShowAll            bool
	ShowLabels         bool
	AbsoluteTimestamps bool
	KindName           string
	ColumnLabels       []string
}

// PrettyPrinter is an implementation of ResourcePrinter which attempts to provide
// more elegant output. It is not threadsafe, but you may call PrintObj repeatedly; headers
// will only be printed if the object type changes. This makes it useful for printing items
// received from watches.
type PrettyPrinter struct {
	handlerMap map[reflect.Type]*handlerEntry
	Options    PrintOptions
	lastType   reflect.Type
}

// NewPrettyPrinter creates a PrettyPrinter.
func NewPrettyPrinter(noHeaders, wide, showAll, showLabels, absoluteTimestamps bool, columnLabels []string) *PrettyPrinter {
	printer := &PrettyPrinter{
		handlerMap: make(map[reflect.Type]*handlerEntry),
		Options: PrintOptions{
			NoHeaders:          noHeaders,
			WithKind:           false,
			KindName:           "",
			Wide:               wide,
			ShowAll:            showAll,
			ShowLabels:         showLabels,
			AbsoluteTimestamps: absoluteTimestamps,
			ColumnLabels:       columnLabels,
		},
	}
	printer.addDefaultHandlers()
	return printer
}

// Handler adds a print handler with a given set of columns to
// PrettyPrinter instance.  See validatePrintHandlerFunc for
// required method signature.
func (h *PrettyPrinter) Handler(columns []string, printFunc interface{}) error {
	printFuncValue := reflect.ValueOf(printFunc)
	if err := h.validatePrintHandlerFunc(printFuncValue); err != nil {
		klog.Errorf("Unable to add print handler: %v", err)
		return err
	}
	objType := printFuncValue.Type().In(0)
	h.handlerMap[objType] = &handlerEntry{
		columns:   columns,
		printFunc: printFuncValue,
	}
	return nil
}

// validatePrintHandlerFunc validates print handler signature.
// printFunc is the function that will be called to print an object.
// It must be of the following type: func printFunc(object ObjectType,
// w io.Writer, options PrintOptions) error where ObjectType is the
// type of the object that will be printed.
func (h *PrettyPrinter) validatePrintHandlerFunc(printFunc reflect.Value) error {
	if printFunc.Kind() != reflect.Func {
		return fmt.Errorf("invalid print handler. %#v is not a function", printFunc)
	}
	funcType := printFunc.Type()
	if funcType.NumIn() != 3 || funcType.NumOut() != 1 {
		return fmt.Errorf("invalid print handler." +
			"Must accept 3 parameters and return 1 value.")
	}
	if funcType.In(1) != reflect.TypeOf((*io.Writer)(nil)).Elem() ||
		funcType.In(2) != reflect.TypeOf((*PrintOptions)(nil)).Elem() ||
		funcType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return fmt.Errorf("invalid print handler. The expected signature is: "+
			"func handler(obj %v, w io.Writer, options PrintOptions) error", funcType.In(0))
	}
	return nil
}

var podColumns = []string{"NAME", "UNITS", "RUNNING", "STATUS", "RESTARTS", "NODE", "IP", "AGE"}
var nodeColumns = []string{"NAME", "STATUS", "INSTANCE-TYPE", "INSTANCE", "IP", "AGE"}
var serviceColumns = []string{"NAME", "PORT(S)", "SOURCES", "INGRESS ADDRESS", "AGE"}
var eventColumns = []string{"TIMESTAMP", "NAME", "KIND", "STATUS", "SOURCE", "MESSAGE"}
var usageReportColumns = []string{"USAGE", "TYPE", "HOURS"}
var usageColumns = []string{"CATEGORY", "TYPE", "HOURS"}
var metricsColumns = []string{"NAME", "TIMESTAMP", "CPU", "MEMORY", "DISK", "WINDOW"}

func (h *PrettyPrinter) addDefaultHandlers() {
	h.Handler(podColumns, printPod)
	h.Handler(podColumns, printPodList)
	h.Handler(nodeColumns, printNode)
	h.Handler(nodeColumns, printNodeList)
	h.Handler(eventColumns, printEvent)
	h.Handler(eventColumns, printEventList)
	h.Handler(metricsColumns, printMetrics)
	h.Handler(metricsColumns, printMetricsList)
}

func printNode(node *api.Node, w io.Writer, options PrintOptions) error {
	name := node.Name
	kind := options.KindName

	if options.WithKind {
		name = kind + "/" + name
	}
	instanceType := node.Spec.InstanceType
	if node.Spec.Spot {
		instanceType += " (spot)"
	}
	instanceID := node.Status.InstanceID
	ip := api.GetPublicIP(node.Status.Addresses)
	if ip == "" {
		ip = api.GetPrivateIP(node.Status.Addresses)
	}
	phase := string(node.Status.Phase)
	if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s", name, phase, instanceType, instanceID, ip, translateTimestamp(node.CreationTimestamp)); err != nil {
		return err
	}
	// Display caller specify column labels first.
	if _, err := fmt.Fprint(w, AppendLabels(node.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, AppendAllLabels(options.ShowLabels, node.Labels))
	return err
}

func printNodeList(list *api.NodeList, w io.Writer, options PrintOptions) error {
	for _, node := range list.Items {
		if err := printNode(node, w, options); err != nil {
			return err
		}
	}
	return nil
}

func printPod(pod *api.Pod, w io.Writer, options PrintOptions) error {
	name := pod.Name
	kind := options.KindName

	if options.WithKind {
		name = kind + "/" + name
	}

	restarts := 0
	totalUnits := len(pod.Spec.Units)
	runningUnits := 0
	status := "Pod " + string(pod.Status.Phase)
	node := pod.Status.BoundNodeName
	if node == "" {
		node = "<unassigned>"
	}
	ip := api.GetPrivateIP(pod.Status.Addresses)
	if ip == "" {
		ip = api.GetPublicIP(pod.Status.Addresses)
	}
	if ip == "" {
		ip = "<none>"
	}

	reason := ""
	for i := len(pod.Status.UnitStatuses) - 1; i >= 0; i-- {
		unit := pod.Status.UnitStatuses[i]
		restarts += int(unit.RestartCount)
		if unit.State.Waiting != nil && unit.State.Waiting.Reason != "" {
			reason = "Unit Waiting: " + unit.State.Waiting.Reason
		} else if unit.State.Terminated != nil {
			reason = "Unit Terminated: " + fmt.Sprintf("ExitCode:%d", unit.State.Terminated.ExitCode)
		} else if unit.State.Running != nil {
			runningUnits++
		}
	}

	if reason != "" {
		if len(reason) > max_pod_reason_length {
			reason = reason[:max_pod_reason_length] + "..."
		}
		status = fmt.Sprintf("%s - %s", status, reason)
	}
	if _, err := fmt.Fprintf(w, "%s\t%d\t%d\t%s\t%d\t%s\t%s\t%s",
		name,
		totalUnits,
		runningUnits,
		status,
		restarts,
		node,
		ip,
		translateTimestamp(pod.CreationTimestamp),
	); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, AppendLabels(pod.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, AppendAllLabels(options.ShowLabels, pod.Labels)); err != nil {
		return err
	}

	return nil
}

func printPodList(podList *api.PodList, w io.Writer, options PrintOptions) error {
	for _, pod := range podList.Items {
		if err := printPod(pod, w, options); err != nil {
			return err
		}
	}
	return nil
}

func printEvent(event *api.Event, w io.Writer, options PrintOptions) error {
	name := event.InvolvedObject.Name
	kind := options.KindName
	if options.WithKind {
		name = kind + "/" + name
	}
	var when string
	if options.AbsoluteTimestamps {
		when = event.CreationTimestamp.String()
	} else {
		when = translateTimestamp(event.CreationTimestamp)
	}
	if _, err := fmt.Fprintf(
		w, "%s\t%s\t%s\t%s\t%s\t%s",
		when,
		name,
		event.InvolvedObject.Kind,
		event.Status,
		event.Source,
		event.Message,
	); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, AppendLabels(event.Labels, options.ColumnLabels)); err != nil {
		return err
	}
	_, err := fmt.Fprint(w, AppendAllLabels(options.ShowLabels, event.Labels))
	return err
}

// Prints the EventList in a human-friendly format.
func printEventList(list *api.EventList, w io.Writer, options PrintOptions) error {
	for i := range list.Items {
		if err := printEvent(list.Items[i], w, options); err != nil {
			return err
		}
	}
	return nil
}

func printMetrics(metrics *api.Metrics, w io.Writer, options PrintOptions) error {
	d := metrics.Window.Round(time.Second)
	ts := metrics.Timestamp.UTC().Format(time.Stamp)
	cpu := metrics.ResourceUsage["cpu"]
	memory := metrics.ResourceUsage["memory"]
	disk := metrics.ResourceUsage["disk"]
	if _, err := fmt.Fprintf(w, "%s\t%s\t%02f\t%02f\t%02f\t%s\n", metrics.Name, ts, cpu, memory, disk, d.String()); err != nil {
		return err
	}
	return nil
}

func printMetricsList(list *api.MetricsList, w io.Writer, options PrintOptions) error {
	sort.Slice(list.Items, func(i, j int) bool {
		return list.Items[i].Name < list.Items[j].Name
	})
	for _, metrics := range list.Items {
		if err := printMetrics(metrics, w, options); err != nil {
			return err
		}
	}
	return nil
}

// GetNewTabWriter returns a tabwriter that translates tabbed columns in input into properly aligned text.
func GetNewTabWriter(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
}

func (h *PrettyPrinter) PrintObj(obj api.MilpaObject, output io.Writer) error {
	// if output is a tabwriter (when it's called by kubectl get), we use it; create a new tabwriter otherwise
	w, found := output.(*tabwriter.Writer)
	if !found {
		w = GetNewTabWriter(output)
		defer w.Flush()
	}
	t := reflect.TypeOf(obj)
	if handler := h.handlerMap[t]; handler != nil {
		if !h.Options.NoHeaders && t != h.lastType {
			headers := append(handler.columns, formatWideHeaders(h.Options.Wide, t)...)
			headers = append(headers, formatLabelHeaders(h.Options.ColumnLabels)...)
			// LABELS is always the last column.
			headers = append(headers, formatShowLabelsHeader(h.Options.ShowLabels, t)...)
			// if h.Options.WithNamespace {
			// 	headers = append(withNamespacePrefixColumns, headers...)
			// }
			h.printHeader(headers, w)
			h.lastType = t
		}
		args := []reflect.Value{reflect.ValueOf(obj), reflect.ValueOf(w), reflect.ValueOf(h.Options)}
		resultValue := handler.printFunc.Call(args)[0]
		if resultValue.IsNil() {
			return nil
		}
		return resultValue.Interface().(error)
	}
	return fmt.Errorf("error: unknown type %#v", obj)
}

// headers for -o wide
func formatWideHeaders(wide bool, t reflect.Type) []string {
	if wide {
		if t.String() == "*api.Pod" || t.String() == "*api.PodList" {
			return []string{"IP", "NODE"}
		}
		if t.String() == "*api.Service" || t.String() == "*api.ServiceList" {
			return []string{"SELECTOR"}
		}
	}
	return nil
}

func formatLabelHeaders(columnLabels []string) []string {
	formHead := make([]string, len(columnLabels))
	for i, l := range columnLabels {
		p := strings.Split(l, "/")
		formHead[i] = strings.ToUpper((p[len(p)-1]))
	}
	return formHead
}

// headers for --show-labels=true
func formatShowLabelsHeader(showLabels bool, t reflect.Type) []string {
	if showLabels {
		if t.String() != "*api.ThirdPartyResource" && t.String() != "*api.ThirdPartyResourceList" {
			return []string{"LABELS"}
		}
	}
	return nil
}

func (h *PrettyPrinter) printHeader(columnNames []string, w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t")); err != nil {
		return err
	}
	return nil
}

func AppendLabels(itemLabels map[string]string, columnLabels []string) string {
	var buffer bytes.Buffer

	for _, cl := range columnLabels {
		buffer.WriteString(fmt.Sprint("\t"))
		if il, ok := itemLabels[cl]; ok {
			buffer.WriteString(il)
		} else {
			buffer.WriteString("<none>")
		}
	}

	return buffer.String()
}

// Append all labels to a single column. We need this even when
// show-labels flag* is false, since this adds newline delimiter to
// the end of each row.
func AppendAllLabels(showLabels bool, itemLabels map[string]string) string {
	var buffer bytes.Buffer

	if showLabels {
		buffer.WriteString("\t")
		buffer.WriteString(labels.FormatLabels(itemLabels))
	}
	buffer.WriteString("\n")

	return buffer.String()
}

func shortHumanDuration(d time.Duration) string {
	// Allow deviation no more than 2 seconds(excluded) to tolerate machine time
	// inconsistence, it can be considered as almost now.
	if seconds := int(d.Seconds()); seconds < -1 {
		return fmt.Sprintf("<invalid>")
	} else if seconds < 0 {
		return fmt.Sprintf("0s")
	} else if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	} else if minutes := int(d.Minutes()); minutes < 60 {
		return fmt.Sprintf("%dm", minutes)
	} else if hours := int(d.Hours()); hours < 24 {
		return fmt.Sprintf("%dh", hours)
	} else if hours < 24*364 {
		return fmt.Sprintf("%dd", hours/24)
	}
	return fmt.Sprintf("%dy", int(d.Hours()/24/365))
}

// translateTimestamp returns the elapsed time since timestamp in
// human-readable approximation.
func translateTimestamp(timestamp api.Time) string {
	if timestamp.IsZero() {
		return "<unknown>"
	}
	return shortHumanDuration(time.Since(timestamp.Time))
}
