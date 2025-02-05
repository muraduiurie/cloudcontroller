package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// ---------------------------------------------------
// GCPKubernetesCluster
// ---------------------------------------------------
func (in *GCPKubernetesCluster) DeepCopyInto(out *GCPKubernetesCluster) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = GCPKubernetesClusterSpec{
		Name:             in.Spec.Name,
		InitialNodeCount: in.Spec.InitialNodeCount,
	}
}

func (in *GCPKubernetesCluster) DeepCopyObject() runtime.Object {
	out := GCPKubernetesCluster{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *GCPKubernetesClusterList) DeepCopyObject() runtime.Object {
	out := GCPKubernetesClusterList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]GCPKubernetesCluster, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

// ---------------------------------------------------
// GCPInstance
// ---------------------------------------------------
func (in *GCPInstance) DeepCopyInto(out *GCPInstance) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = GCPInstanceSpec{
		Name: in.Spec.Name,
	}
}

func (in *GCPInstance) DeepCopyObject() runtime.Object {
	out := GCPInstance{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *GCPInstanceList) DeepCopyObject() runtime.Object {
	out := GCPInstanceList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]GCPInstance, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

// ---------------------------------------------------
// GCPNetwork
// ---------------------------------------------------
func (in *GCPNetwork) DeepCopyInto(out *GCPNetwork) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = GCPNetworkSpec{
		Name:                  in.Spec.Name,
		AutoCreateSubnetworks: in.Spec.AutoCreateSubnetworks,
	}
}

func (in *GCPNetwork) DeepCopyObject() runtime.Object {
	out := GCPNetwork{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *GCPNetworkList) DeepCopyObject() runtime.Object {
	out := GCPNetworkList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]GCPNetwork, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
