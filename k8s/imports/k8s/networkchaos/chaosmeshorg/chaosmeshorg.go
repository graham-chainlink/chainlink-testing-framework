// chaos-meshorg
package chaosmeshorg

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/smartcontractkit/chainlink-testing-framework/k8s/imports/k8s/networkchaos/chaosmeshorg/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/smartcontractkit/chainlink-testing-framework/k8s/imports/k8s/networkchaos/chaosmeshorg/internal"
)

// NetworkChaos is the Schema for the networkchaos API.
type NetworkChaos interface {
	cdk8s.ApiObject
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion() *string
	// The chart in which this object is defined.
	Chart() cdk8s.Chart
	// The object kind.
	Kind() *string
	// Metadata associated with this API object.
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Create a dependency between this ApiObject and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	AddDependency(dependencies ...constructs.IConstruct)
	// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
	//
	// Example:
	//     kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
	//
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	// Renders the object to Kubernetes JSON.
	ToJson() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for NetworkChaos
type jsiiProxy_NetworkChaos struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_NetworkChaos) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkChaos) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkChaos) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkChaos) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkChaos) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkChaos) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkChaos) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

// Defines a "NetworkChaos" API object.
func NewNetworkChaos(scope constructs.Construct, id *string, props *NetworkChaosProps) NetworkChaos {
	_init_.Initialize()

	j := jsiiProxy_NetworkChaos{}

	_jsii_.Create(
		"chaos-meshorg.NetworkChaos",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "NetworkChaos" API object.
func NewNetworkChaos_Override(n NetworkChaos, scope constructs.Construct, id *string, props *NetworkChaosProps) {
	_init_.Initialize()

	_jsii_.Create(
		"chaos-meshorg.NetworkChaos",
		[]interface{}{scope, id, props},
		n,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NetworkChaos_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"chaos-meshorg.NetworkChaos",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "NetworkChaos".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func NetworkChaos_Manifest(props *NetworkChaosProps) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"chaos-meshorg.NetworkChaos",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func NetworkChaos_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"chaos-meshorg.NetworkChaos",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func NetworkChaos_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"chaos-meshorg.NetworkChaos",
		"GVK",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkChaos) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addDependency",
		args,
	)
}

func (n *jsiiProxy_NetworkChaos) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		n,
		"addJsonPatch",
		args,
	)
}

func (n *jsiiProxy_NetworkChaos) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkChaos) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// NetworkChaos is the Schema for the networkchaos API.
type NetworkChaosProps struct {
	// Spec defines the behavior of a pod chaos experiment.
	Spec     *NetworkChaosSpec        `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

// Spec defines the behavior of a pod chaos experiment.
type NetworkChaosSpec struct {
	// Action defines the specific network chaos action.
	//
	// Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay.
	Action NetworkChaosSpecAction `field:"required" json:"action" yaml:"action"`
	// Mode defines the mode to run chaos action.
	//
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent.
	Mode NetworkChaosSpecMode `field:"required" json:"mode" yaml:"mode"`
	// Selector is used to select pods that are used to inject chaos action.
	Selector *NetworkChaosSpecSelector `field:"required" json:"selector" yaml:"selector"`
	// Bandwidth represents the detail about bandwidth control action.
	Bandwidth *NetworkChaosSpecBandwidth `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Corrupt represents the detail about corrupt action.
	Corrupt *NetworkChaosSpecCorrupt `field:"optional" json:"corrupt" yaml:"corrupt"`
	// Delay represents the detail about delay action.
	Delay *NetworkChaosSpecDelay `field:"optional" json:"delay" yaml:"delay"`
	// Direction represents the direction, this applies on netem and network partition action.
	Direction NetworkChaosSpecDirection `field:"optional" json:"direction" yaml:"direction"`
	// DuplicateSpec represents the detail about loss action.
	Duplicate *NetworkChaosSpecDuplicate `field:"optional" json:"duplicate" yaml:"duplicate"`
	// Duration represents the duration of the chaos action.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// ExternalTargets represents network targets outside k8s.
	ExternalTargets *[]*string `field:"optional" json:"externalTargets" yaml:"externalTargets"`
	// Loss represents the detail about loss action.
	Loss *NetworkChaosSpecLoss `field:"optional" json:"loss" yaml:"loss"`
	// Target represents network target, this applies on netem and network partition action.
	Target *NetworkChaosSpecTarget `field:"optional" json:"target" yaml:"target"`
	// Value is required when the mode is set to `FixedPodMode` / `FixedPercentPodMod` / `RandomMaxPercentPodMod`.
	//
	// If `FixedPodMode`, provide an integer of pods to do chaos action. If `FixedPercentPodMod`, provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF `RandomMaxPercentPodMod`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Action defines the specific network chaos action.
//
// Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay.
type NetworkChaosSpecAction string

const (
	// netem.
	NetworkChaosSpecAction_NETEM NetworkChaosSpecAction = "NETEM"
	// delay.
	NetworkChaosSpecAction_DELAY NetworkChaosSpecAction = "DELAY"
	// loss.
	NetworkChaosSpecAction_LOSS NetworkChaosSpecAction = "LOSS"
	// duplicate.
	NetworkChaosSpecAction_DUPLICATE NetworkChaosSpecAction = "DUPLICATE"
	// corrupt.
	NetworkChaosSpecAction_CORRUPT NetworkChaosSpecAction = "CORRUPT"
	// partition.
	NetworkChaosSpecAction_PARTITION NetworkChaosSpecAction = "PARTITION"
	// bandwidth.
	NetworkChaosSpecAction_BANDWIDTH NetworkChaosSpecAction = "BANDWIDTH"
)

// Bandwidth represents the detail about bandwidth control action.
type NetworkChaosSpecBandwidth struct {
	// Buffer is the maximum amount of bytes that tokens can be available for instantaneously.
	Buffer *float64 `field:"required" json:"buffer" yaml:"buffer"`
	// Limit is the number of bytes that can be queued waiting for tokens to become available.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Rate is the speed knob.
	//
	// Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.
	Rate *string `field:"required" json:"rate" yaml:"rate"`
	// Minburst specifies the size of the peakrate bucket.
	//
	// For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.
	Minburst *float64 `field:"optional" json:"minburst" yaml:"minburst"`
	// Peakrate is the maximum depletion rate of the bucket.
	//
	// The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.
	Peakrate *float64 `field:"optional" json:"peakrate" yaml:"peakrate"`
}

// Corrupt represents the detail about corrupt action.
type NetworkChaosSpecCorrupt struct {
	Corrupt     *string `field:"required" json:"corrupt" yaml:"corrupt"`
	Correlation *string `field:"optional" json:"correlation" yaml:"correlation"`
}

// Delay represents the detail about delay action.
type NetworkChaosSpecDelay struct {
	Latency     *string `field:"required" json:"latency" yaml:"latency"`
	Correlation *string `field:"optional" json:"correlation" yaml:"correlation"`
	Jitter      *string `field:"optional" json:"jitter" yaml:"jitter"`
	// ReorderSpec defines details of packet reorder.
	Reorder *NetworkChaosSpecDelayReorder `field:"optional" json:"reorder" yaml:"reorder"`
}

// ReorderSpec defines details of packet reorder.
type NetworkChaosSpecDelayReorder struct {
	Gap         *float64 `field:"required" json:"gap" yaml:"gap"`
	Reorder     *string  `field:"required" json:"reorder" yaml:"reorder"`
	Correlation *string  `field:"optional" json:"correlation" yaml:"correlation"`
}

// Direction represents the direction, this applies on netem and network partition action.
type NetworkChaosSpecDirection string

const (
	// to.
	NetworkChaosSpecDirection_TO NetworkChaosSpecDirection = "TO"
	// from.
	NetworkChaosSpecDirection_FROM NetworkChaosSpecDirection = "FROM"
	// both.
	NetworkChaosSpecDirection_BOTH   NetworkChaosSpecDirection = "BOTH"
	NetworkChaosSpecDirection_VALUE_ NetworkChaosSpecDirection = "VALUE_"
)

// DuplicateSpec represents the detail about loss action.
type NetworkChaosSpecDuplicate struct {
	Duplicate   *string `field:"required" json:"duplicate" yaml:"duplicate"`
	Correlation *string `field:"optional" json:"correlation" yaml:"correlation"`
}

// Loss represents the detail about loss action.
type NetworkChaosSpecLoss struct {
	Loss        *string `field:"required" json:"loss" yaml:"loss"`
	Correlation *string `field:"optional" json:"correlation" yaml:"correlation"`
}

// Mode defines the mode to run chaos action.
//
// Supported mode: one / all / fixed / fixed-percent / random-max-percent.
type NetworkChaosSpecMode string

const (
	// one.
	NetworkChaosSpecMode_ONE NetworkChaosSpecMode = "ONE"
	// all.
	NetworkChaosSpecMode_ALL NetworkChaosSpecMode = "ALL"
	// fixed.
	NetworkChaosSpecMode_FIXED NetworkChaosSpecMode = "FIXED"
	// fixed-percent.
	NetworkChaosSpecMode_FIXED_PERCENT NetworkChaosSpecMode = "FIXED_PERCENT"
	// random-max-percent.
	NetworkChaosSpecMode_RANDOM_MAX_PERCENT NetworkChaosSpecMode = "RANDOM_MAX_PERCENT"
)

// Selector is used to select pods that are used to inject chaos action.
type NetworkChaosSpecSelector struct {
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on annotations.
	AnnotationSelectors *map[string]*string `field:"optional" json:"annotationSelectors" yaml:"annotationSelectors"`
	// a slice of label selector expressions that can be used to select objects.
	//
	// A list of selectors based on set-based label expressions.
	ExpressionSelectors *[]*NetworkChaosSpecSelectorExpressionSelectors `field:"optional" json:"expressionSelectors" yaml:"expressionSelectors"`
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on fields.
	FieldSelectors *map[string]*string `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on labels.
	LabelSelectors *map[string]*string `field:"optional" json:"labelSelectors" yaml:"labelSelectors"`
	// Namespaces is a set of namespace to which objects belong.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Nodes is a set of node name and objects must belong to these nodes.
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// Map of string keys and values that can be used to select nodes.
	//
	// Selector which must match a node's labels, and objects must belong to these selected nodes.
	NodeSelectors *map[string]*string `field:"optional" json:"nodeSelectors" yaml:"nodeSelectors"`
	// PodPhaseSelectors is a set of condition of a pod at the current time.
	//
	// supported value: Pending / Running / Succeeded / Failed / Unknown.
	PodPhaseSelectors *[]*string `field:"optional" json:"podPhaseSelectors" yaml:"podPhaseSelectors"`
	// Pods is a map of string keys and a set values that used to select pods.
	//
	// The key defines the namespace which pods belong, and the each values is a set of pod names.
	Pods *map[string]*[]*string `field:"optional" json:"pods" yaml:"pods"`
}

// A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type NetworkChaosSpecSelectorExpressionSelectors struct {
	// key is the label key that the selector applies to.
	Key *string `field:"required" json:"key" yaml:"key"`
	// operator represents a key's relationship to a set of values.
	//
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// values is an array of string values.
	//
	// If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

// Target represents network target, this applies on netem and network partition action.
type NetworkChaosSpecTarget struct {
	// Mode defines the mode to run chaos action.
	//
	// Supported mode: one / all / fixed / fixed-percent / random-max-percent.
	Mode NetworkChaosSpecTargetMode `field:"required" json:"mode" yaml:"mode"`
	// Selector is used to select pods that are used to inject chaos action.
	Selector *NetworkChaosSpecTargetSelector `field:"required" json:"selector" yaml:"selector"`
	// Value is required when the mode is set to `FixedPodMode` / `FixedPercentPodMod` / `RandomMaxPercentPodMod`.
	//
	// If `FixedPodMode`, provide an integer of pods to do chaos action. If `FixedPercentPodMod`, provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF `RandomMaxPercentPodMod`,  provide a number from 0-100 to specify the max percent of pods to do chaos action
	Value *string `field:"optional" json:"value" yaml:"value"`
}

// Mode defines the mode to run chaos action.
//
// Supported mode: one / all / fixed / fixed-percent / random-max-percent.
type NetworkChaosSpecTargetMode string

const (
	// one.
	NetworkChaosSpecTargetMode_ONE NetworkChaosSpecTargetMode = "ONE"
	// all.
	NetworkChaosSpecTargetMode_ALL NetworkChaosSpecTargetMode = "ALL"
	// fixed.
	NetworkChaosSpecTargetMode_FIXED NetworkChaosSpecTargetMode = "FIXED"
	// fixed-percent.
	NetworkChaosSpecTargetMode_FIXED_PERCENT NetworkChaosSpecTargetMode = "FIXED_PERCENT"
	// random-max-percent.
	NetworkChaosSpecTargetMode_RANDOM_MAX_PERCENT NetworkChaosSpecTargetMode = "RANDOM_MAX_PERCENT"
)

// Selector is used to select pods that are used to inject chaos action.
type NetworkChaosSpecTargetSelector struct {
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on annotations.
	AnnotationSelectors *map[string]*string `field:"optional" json:"annotationSelectors" yaml:"annotationSelectors"`
	// a slice of label selector expressions that can be used to select objects.
	//
	// A list of selectors based on set-based label expressions.
	ExpressionSelectors *[]*NetworkChaosSpecTargetSelectorExpressionSelectors `field:"optional" json:"expressionSelectors" yaml:"expressionSelectors"`
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on fields.
	FieldSelectors *map[string]*string `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// Map of string keys and values that can be used to select objects.
	//
	// A selector based on labels.
	LabelSelectors *map[string]*string `field:"optional" json:"labelSelectors" yaml:"labelSelectors"`
	// Namespaces is a set of namespace to which objects belong.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Nodes is a set of node name and objects must belong to these nodes.
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// Map of string keys and values that can be used to select nodes.
	//
	// Selector which must match a node's labels, and objects must belong to these selected nodes.
	NodeSelectors *map[string]*string `field:"optional" json:"nodeSelectors" yaml:"nodeSelectors"`
	// PodPhaseSelectors is a set of condition of a pod at the current time.
	//
	// supported value: Pending / Running / Succeeded / Failed / Unknown.
	PodPhaseSelectors *[]*string `field:"optional" json:"podPhaseSelectors" yaml:"podPhaseSelectors"`
	// Pods is a map of string keys and a set values that used to select pods.
	//
	// The key defines the namespace which pods belong, and the each values is a set of pod names.
	Pods *map[string]*[]*string `field:"optional" json:"pods" yaml:"pods"`
}

// A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
type NetworkChaosSpecTargetSelectorExpressionSelectors struct {
	// key is the label key that the selector applies to.
	Key *string `field:"required" json:"key" yaml:"key"`
	// operator represents a key's relationship to a set of values.
	//
	// Valid operators are In, NotIn, Exists and DoesNotExist.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// values is an array of string values.
	//
	// If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}
