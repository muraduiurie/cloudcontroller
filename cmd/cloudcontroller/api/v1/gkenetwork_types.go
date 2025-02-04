// Package v1 is the v1 version of the API.
// Package v1 +kubebuilder:object:generate=true
// +groupName=benzaiten.io
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GKENetworkList contains a list of GKENetwork
// +kubebuilder:object:root=true
type GKENetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Items is the list of GKENetworks
	Items []GKENetwork `json:"items"`
}

// GKENetwork is the Schema for the gkenetworks API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,path=gkenetworks,shortName=gke,singular=gkenetwork
type GKENetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	// Spec defines the desired state of GKENetwork
	Spec GKENetworkSpec `json:"spec"`
	// +kubebuilder:validation:Optional
	// Status defines the observed state of GKENetwork
	Status GKENetworkStatus `json:"status"`
}

// GKENetworkSpec defines the desired state of GKENetwork
type GKENetworkSpec struct {
	// +kubebuilder:validation:Required
	// Name is the name of the GKE network
	Name string `json:"name"`
	// +kubebuilder:validation:Required
	// AutoCreateSubnetworks
	AutoCreateSubnetworks bool `json:"autoCreateSubnetworks"`
}

// GKENetworkStatus defines the observed state of GKENetwork
type GKENetworkStatus struct {
	// +kubebuilder:validation:Optional
	// Phase is the current state of the GKE cluster
	Phase string `json:"phase"`
}
