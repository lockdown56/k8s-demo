package patch

import (
	"encoding/json"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func mockCurr() corev1.Service {
	return corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Protocol:   "tcp",
					Port:       80,
					TargetPort: intstr.IntOrString{},
				},
			},
			Selector: nil,
			Type:     "ClusterIp",
		},
	}

}

// demo
func TestTwoWayMergePatch4Service(t *testing.T) {
	curr := mockCurr()
	modified := mockCurr()
	modified.Spec.Ports[0].Port = 8080

	currJson, _ := json.Marshal(curr)
	modifiedJson, _ := json.Marshal(modified)

	patch, err := TwoWayMergePatch4Service(currJson, modifiedJson)
	if err != nil {
		t.Fatalf("patch failed, %v", err)
	}

	t.Logf("patch: %s", patch)
}

func TestTwoWayMergePatchWithKind(t *testing.T) {
	curr := mockCurr()
	modified := mockCurr()
	modified.Spec.Ports[0].Port = 8080

	currJson, _ := json.Marshal(curr)
	modifiedJson, _ := json.Marshal(modified)

	groupVersionKind := schema.GroupVersionKind{
		Group:   "",
		Version: "v1",
		Kind:    "Service",
	}

	patch, err := TwoWayMergePatchWithKind(currJson, modifiedJson, groupVersionKind)
	if err != nil {
		t.Fatalf("patch failed, %v", err)
	}

	t.Logf("patch: %s", patch)

}
