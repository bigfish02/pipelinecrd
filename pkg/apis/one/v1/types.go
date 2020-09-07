package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Network describes a Network resource
type Pipeline struct {
	// TypeMeta is the metadata for the resource, like kind and apiversion
	metav1.TypeMeta `json:",inline"`
	// ObjectMeta contains the metadata for the particular object, including
	// things like...
	//  - name
	//  - namespace
	//  - self link
	//  - labels
	//  - ... etc ...
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec is the custom resource spec
	Spec PipelineSpec `json:"spec"`
}

type PipelineSpec struct {
	Type    PipelineType `json:"type"`
	Trigger string       `json:"trigger"`
	Stages  []Stage      `json:"stages"`
}

type Stage struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Cluster string `json:"cluster"`
}

type PipelineType string

var (
	GlobalPipeline   PipelineType = "Global"
	CustomerPipeline PipelineType = "Customer"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Pipeline `json:"items"`
}
