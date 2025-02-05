package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

const (
	GroupName    = "benzaiten.io"
	GroupVersion = "v1"
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemaGroupVersion,
		&GCPKubernetesCluster{},
		&GCPKubernetesClusterList{},
		&GCPInstance{},
		&GCPInstanceList{},
		&GCPNetwork{},
		&GCPNetworkList{},
	)

	metav1.AddToGroupVersion(scheme, SchemaGroupVersion)
	return nil
}

var SchemaGroupVersion = schema.GroupVersion{
	Group:   GroupName,
	Version: GroupVersion,
}
