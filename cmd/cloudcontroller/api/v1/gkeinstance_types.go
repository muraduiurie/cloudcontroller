// Package v1 is the v1 version of the API.
// Package v1 +kubebuilder:object:generate=true
// +groupName=benzaiten.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKEInstanceList contains a list of GKEInstance
// +kubebuilder:object:root=true
type GKEInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Items is the list of GKEInstances
	Items []GKEInstance `json:"items"`
}

// GKEInstance is the Schema for the gkeinstances API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gkeinstances,shortName=gke,singular=gkeinstance
type GKEInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Spec defines the desired state of GKEInstance
	Spec GKEInstanceSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status defines the observed state of GKEInstance
	Status GKEInstanceStatus `json:"status"`
}

// GKEInstanceSpec defines the desired state of GKEInstance
type GKEInstanceSpec struct {
	// +kubebuilder:validation:Required
	// Name is the name of the GKE instance
	Name string `json:"name"`
}

// GKEInstanceStatus defines the observed state of GKEInstance
type GKEInstanceStatus struct {
	// +kubebuilder:validation:Optional
	// Phase is the current state of the GKE instance
	Phase string `json:"phase"`
}
