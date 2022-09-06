package patch

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/mergepatch"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"
)

// TwoWayMergePatch4Service 指定资源对象的方式, 这里以 Service 对象为例
func TwoWayMergePatch4Service(curr []byte, modified []byte) ([]byte, error) {
	// 条件, 理论上可以省略
	preconditions := []mergepatch.PreconditionFunc{
		mergepatch.RequireKeyUnchanged("apiVersion"),
		mergepatch.RequireKeyUnchanged("kind"),
		mergepatch.RequireMetadataKeyUnchanged("name"),
		mergepatch.RequireKeyUnchanged("managedFields"),
	}

	patch, err := strategicpatch.CreateTwoWayMergePatch(curr, modified, corev1.Service{}, preconditions...)
	if err != nil {
		return nil, err
	}

	return patch, nil

}

// TwoWayMergePatchWithKind 针对非特定类型
func TwoWayMergePatchWithKind(curr []byte, modified []byte, kind schema.GroupVersionKind) ([]byte, error) {
	versionedObject, err := scheme.Scheme.New(kind)
	if err != nil {
		return nil, err
	}

	// 条件, 理论上可以省略
	preconditions := []mergepatch.PreconditionFunc{
		mergepatch.RequireKeyUnchanged("apiVersion"),
		mergepatch.RequireKeyUnchanged("kind"),
		mergepatch.RequireMetadataKeyUnchanged("name"),
		mergepatch.RequireKeyUnchanged("managedFields"),
	}

	patch, err := strategicpatch.CreateTwoWayMergePatch(curr, modified, versionedObject, preconditions...)
	if err != nil {
		return nil, err
	}

	return patch, nil
}
