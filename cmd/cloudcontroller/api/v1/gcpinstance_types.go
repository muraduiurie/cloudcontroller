// Package v1 is the v1 version of the API.
// Package v1 +kubebuilder:object:generate=true
// +groupName=benzaiten.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GCPInstanceList contains a list of GCPInstance
// +kubebuilder:object:root=true
type GCPInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Items is the list of GCPInstances
	Items []GCPInstance `json:"items"`
}

// GCPInstance is the Schema for the gcpinstances API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gcpinstances,shortName=gi,singular=gcpinstance
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=".status.phase"
type GCPInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Spec defines the desired state of GCPInstance
	Spec GCPInstanceSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status defines the observed state of GCPInstance
	Status GCPInstanceStatus `json:"status"`
}

// GCPInstanceSpec defines the desired state of GCPInstance
type GCPInstanceSpec struct {
	// +kubebuilder:validation:Required
	// Name is the name of the GCP instance
	Name string `json:"name"`
}

// GCPInstanceStatus defines the observed state of GCPInstance
type GCPInstanceStatus struct {
	// +kubebuilder:validation:Optional
	// Phase is the current state of the GCP instance
	Phase string `json:"phase"`
}
