package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QdrouterdSpec defines the desired state of Qdrouterd
type QdrouterdSpec struct {
	DeploymentPlan        DeploymentPlanType `json:"deploymentPlan,omitempty"`
	Listeners             []Listener         `json:"listeners,omitempty"`
	InterRouterListeners  []Listener         `json:"interRouterListeners,omitempty"`
	EdgeListeners         []Listener         `json:"edgeListeners,omitempty"`
	SslProfiles           []SslProfile       `json:"sslProfiles,omitempty"`
	Addresses             []Address          `json:"addresses,omitempty"`
	AutoLinks             []AutoLink         `json:"autoLinks,omitempty"`
	LinkRoutes            []LinkRoute        `json:"linkRoutes,omitempty"`
	Connectors            []Connector        `json:"connectors,omitempty"`
	InterRouterConnectors []Connector        `json:"interRouterConnectors,omitempty"`
	EdgeConnectors        []Connector        `json:"edgeConnectors,omitempty"`
}

type PhaseType string

const (
	QdrouterdPhaseNone     PhaseType = ""
	QdrouterdPhaseCreating           = "Creating"
	QdrouterdPhaseRunning            = "Running"
	QdrouterdPhaseFailed             = "Failed"
)

type ConditionType string

const (
	QdrouterdConditionProvisioning ConditionType = "Provisioning"
	QdrouterdConditionDeployed     ConditionType = "Deployed"
	QdrouterdConditionScalingUp    ConditionType = "ScalingUp"
	QdrouterdConditionScalingDown  ConditionType = "ScalingDown"
	QdrouterdConditionUpgrading    ConditionType = "Upgrading"
)

type QdrouterdCondition struct {
	Type           ConditionType `json:"type"`
	TransitionTime metav1.Time   `json:"transitionTime,omitempty"`
	Reason         string        `json:"reason,omitempty"`
}

// QdrouterdStatus defines the observed state of Qdrouterd
type QdrouterdStatus struct {
	Phase     PhaseType `json:"phase,omitempty"`
	RevNumber string    `json:"revNumber,omitempty"`
	PodNames  []string  `json:"pods"`

	// Conditions keeps most recent qdrouterd conditions
	Conditions []QdrouterdCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Qdrouterd is the Schema for the qdrouterds API
// +k8s:openapi-gen=true
type Qdrouterd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QdrouterdSpec   `json:"spec,omitempty"`
	Status QdrouterdStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QdrouterdList contains a list of Qdrouterd
type QdrouterdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Qdrouterd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Qdrouterd{}, &QdrouterdList{})
}

type RouterRoleType string

const (
	RouterRoleInterior RouterRoleType = "interior"
	RouterRoleEdge                    = "edge"
)

type PlacementType string

const (
	PlacementAny          PlacementType = "Any"
	PlacementEvery                      = "Every"
	PlacementAntiAffinity               = "AntiAffinity"
	PlacementNode                       = "Node"
)

type DeploymentPlanType struct {
	Image     string                      `json:"image,omitempty"`
	Size      int32                       `json:"size,omitempty"`
	Role      RouterRoleType              `json:"role,omitempty"`
	Placement PlacementType               `json:"placement,omitempty"`
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	Issuer    string                      `json:"issuer,omitempty"`
}

type Address struct {
	Prefix       string `json:"prefix,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Waypoint     bool   `json:"waypoint,omitempty"`
	IngressPhase *int32 `json:"ingressPhase,omitempty"`
	EgressPhase  *int32 `json:"egressPhase,omitempty"`
	Priority     *int32 `json:"priority,omitempty"`
}

type Listener struct {
	Name           string `json:"name,omitempty"`
	Host           string `json:"host,omitempty"`
	Port           int32  `json:"port"`
	RouteContainer bool   `json:"routeContainer,omitempty"`
	Http           bool   `json:"http,omitempty"`
	Cost           int32  `json:"cost,omitempty"`
	SslProfile     string `json:"sslProfile,omitempty"`
	Expose         bool   `json:"expose,omitempty"`
}

type SslProfile struct {
	Name               string `json:"name,omitempty"`
	Credentials        string `json:"credentials,omitempty"`
	CaCert             string `json:"caCert,omitempty"`
	RequireClientCerts bool   `json:"requireClientCerts,omitempty"`
	Ciphers            string `json:"ciphers,omitempty"`
	Protocols          string `json:"protocols,omitempty"`
}

type LinkRoute struct {
	Prefix               string `json:"prefix,omitempty"`
	Pattern              string `json:"pattern,omitempty"`
	Direction            string `json:"direction,omitempty"`
	ContainerId          string `json:"containerId,omitempty"`
	Connection           string `json:"connection,omitempty"`
	AddExternalPrefix    string `json:"addExternalPrefix,omitempty"`
	RemoveExternalPrefix string `json:"removeExternalPrefix,omitempty"`
}

type Connector struct {
	Name           string `json:"name,omitempty"`
	Host           string `json:"host"`
	Port           int32  `json:"port"`
	RouteContainer bool   `json:"routeContainer,omitempty"`
	Cost           int32  `json:"cost,omitempty"`
	SslProfile     string `json:"sslProfile,omitempty"`
}

type AutoLink struct {
	Address        string `json:"address"`
	Direction      string `json:"direction"`
	ContainerId    string `json:"containerId,omitempty"`
	Connection     string `json:"connection,omitempty"`
	ExternalPrefix string `json:"externalPrefix,omitempty"`
	Phase          *int32 `json:"phase,omitempty"`
}
