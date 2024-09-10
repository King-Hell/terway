package webhook

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/AliyunContainerService/terway/pkg/apis/network.alibabacloud.com/v1beta1"
	"github.com/AliyunContainerService/terway/types/controlplane"
)

func Test_setNodeAffinityByZones(t *testing.T) {
	type args struct {
		pod       *corev1.Pod
		zones     []string
		prevZones []string
	}
	tests := []struct {
		name string
		args args
		want *corev1.Pod
	}{
		{
			name: "ds pod should ignore",
			args: args{
				pod: &corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						OwnerReferences: []metav1.OwnerReference{
							{
								Kind: "DaemonSet",
							},
						},
					},
				},
				zones: []string{"foo", "bar"},
			},
			want: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "DaemonSet",
						},
					},
				},
			},
		}, {
			name: "pod with no affinity",
			args: args{
				pod:   &corev1.Pod{},
				zones: []string{"foo", "bar"},
			},
			want: &corev1.Pod{
				Spec: corev1.PodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: []corev1.NodeSelectorTerm{
						{
							MatchExpressions: []corev1.NodeSelectorRequirement{
								{
									Key:      corev1.LabelTopologyZone,
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"foo", "bar"},
								},
							},
						},
					},
				}}}},
			},
		}, {
			name: "pod with exist affinity",
			args: args{
				pod: &corev1.Pod{Spec: corev1.PodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: []corev1.NodeSelectorTerm{
						{
							MatchExpressions: []corev1.NodeSelectorRequirement{
								{
									Key:      corev1.LabelTopologyZone,
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"exist"},
								},
							},
							MatchFields: []corev1.NodeSelectorRequirement{
								{
									Key:      "metadata.name",
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"foo"},
								},
							},
						},
						{
							MatchFields: []corev1.NodeSelectorRequirement{
								{
									Key:      "metadata.name",
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"foo"},
								},
							},
						},
					},
				}}}}},
				zones: []string{"foo", "bar"},
			},
			want: &corev1.Pod{Spec: corev1.PodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
				NodeSelectorTerms: []corev1.NodeSelectorTerm{
					{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"exist"},
							},
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo", "bar"},
							},
						},
						MatchFields: []corev1.NodeSelectorRequirement{
							{
								Key:      "metadata.name",
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo"},
							},
						},
					},
					{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo", "bar"},
							},
						},
						MatchFields: []corev1.NodeSelectorRequirement{
							{
								Key:      "metadata.name",
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo"},
							},
						},
					},
				},
			}}}}},
		}, {
			name: "pod with multi zone set",
			args: args{
				pod: &corev1.Pod{Spec: corev1.PodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
					NodeSelectorTerms: []corev1.NodeSelectorTerm{
						{
							MatchExpressions: []corev1.NodeSelectorRequirement{
								{
									Key:      corev1.LabelTopologyZone,
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"exist"},
								},
							},
							MatchFields: []corev1.NodeSelectorRequirement{
								{
									Key:      "metadata.name",
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"foo"},
								},
							},
						},
						{
							MatchFields: []corev1.NodeSelectorRequirement{
								{
									Key:      "metadata.name",
									Operator: corev1.NodeSelectorOpIn,
									Values:   []string{"foo"},
								},
							},
						},
					},
				}}}}},
				zones:     []string{"foo", "bar"},
				prevZones: []string{"foo", "bar"},
			},
			want: &corev1.Pod{Spec: corev1.PodSpec{Affinity: &corev1.Affinity{NodeAffinity: &corev1.NodeAffinity{RequiredDuringSchedulingIgnoredDuringExecution: &corev1.NodeSelector{
				NodeSelectorTerms: []corev1.NodeSelectorTerm{
					{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"exist"},
							},
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo", "bar"},
							},
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo", "bar"},
							},
						},
						MatchFields: []corev1.NodeSelectorRequirement{
							{
								Key:      "metadata.name",
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo"},
							},
						},
					},
					{
						MatchExpressions: []corev1.NodeSelectorRequirement{
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo", "bar"},
							},
							{
								Key:      corev1.LabelTopologyZone,
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo", "bar"},
							},
						},
						MatchFields: []corev1.NodeSelectorRequirement{
							{
								Key:      "metadata.name",
								Operator: corev1.NodeSelectorOpIn,
								Values:   []string{"foo"},
							},
						},
					},
				},
			}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setNodeAffinityByZones(tt.args.pod, tt.args.zones, tt.args.prevZones)
			if !reflect.DeepEqual(tt.args.pod, tt.want) {
				t.Errorf("setNodeAffinityByZones() = %v, want %v", tt.args.pod, tt.want)
			}
		})
	}
}

func Test_setResourceRequest1(t *testing.T) {
	type args struct {
		pod         *corev1.Pod
		podNetworks []controlplane.PodNetworks
		enableTrunk bool
	}
	tests := []struct {
		name string
		args args
		want *corev1.Pod
	}{
		{
			name: "ignore count=0",
			args: args{
				pod:         &corev1.Pod{},
				podNetworks: nil,
				enableTrunk: false,
			},
			want: &corev1.Pod{},
		},
		{
			name: "request one",
			args: args{
				pod: &corev1.Pod{Spec: corev1.PodSpec{Containers: []corev1.Container{
					{
						Name: "a",
						Resources: corev1.ResourceRequirements{
							Limits: map[corev1.ResourceName]resource.Quantity{
								"bar": resource.MustParse(strconv.Itoa(100)),
							},
						},
					}, {
						Name: "b",
					},
				}}},
				podNetworks: []controlplane.PodNetworks{
					{},
				},
				enableTrunk: true,
			},
			want: &corev1.Pod{Spec: corev1.PodSpec{Containers: []corev1.Container{
				{
					Name: "a",
					Resources: corev1.ResourceRequirements{
						Limits: map[corev1.ResourceName]resource.Quantity{
							"bar":               resource.MustParse(strconv.Itoa(100)),
							"aliyun/member-eni": resource.MustParse(strconv.Itoa(1)),
						},
						Requests: map[corev1.ResourceName]resource.Quantity{
							"aliyun/member-eni": resource.MustParse(strconv.Itoa(1)),
						},
					},
				}, {
					Name: "b",
				},
			}}},
		},
		{
			name: "eni only",
			args: args{
				pod: &corev1.Pod{Spec: corev1.PodSpec{Containers: []corev1.Container{
					{
						Name: "a",
						Resources: corev1.ResourceRequirements{
							Limits: map[corev1.ResourceName]resource.Quantity{
								"bar": resource.MustParse(strconv.Itoa(100)),
							},
						},
					}, {
						Name: "b",
					},
				}}},
				podNetworks: []controlplane.PodNetworks{
					{},
				},
				enableTrunk: false,
			},
			want: &corev1.Pod{Spec: corev1.PodSpec{Containers: []corev1.Container{
				{
					Name: "a",
					Resources: corev1.ResourceRequirements{
						Limits: map[corev1.ResourceName]resource.Quantity{
							"bar":        resource.MustParse(strconv.Itoa(100)),
							"aliyun/eni": resource.MustParse(strconv.Itoa(1)),
						},
						Requests: map[corev1.ResourceName]resource.Quantity{
							"aliyun/eni": resource.MustParse(strconv.Itoa(1)),
						},
					},
				}, {
					Name: "b",
				},
			}}},
		},
		{
			name: "allow override to aliyun/eni",
			args: args{
				pod: &corev1.Pod{Spec: corev1.PodSpec{Containers: []corev1.Container{
					{
						Name: "a",
						Resources: corev1.ResourceRequirements{
							Limits: map[corev1.ResourceName]resource.Quantity{
								"bar": resource.MustParse(strconv.Itoa(100)),
							},
						},
					}, {
						Name: "b",
					},
				}}},
				podNetworks: []controlplane.PodNetworks{
					{
						ENIOptions: v1beta1.ENIOptions{
							ENIAttachType: "ENI",
						},
					},
				},
				enableTrunk: true,
			},
			want: &corev1.Pod{Spec: corev1.PodSpec{Containers: []corev1.Container{
				{
					Name: "a",
					Resources: corev1.ResourceRequirements{
						Limits: map[corev1.ResourceName]resource.Quantity{
							"bar":        resource.MustParse(strconv.Itoa(100)),
							"aliyun/eni": resource.MustParse(strconv.Itoa(1)),
						},
						Requests: map[corev1.ResourceName]resource.Quantity{
							"aliyun/eni": resource.MustParse(strconv.Itoa(1)),
						},
					},
				}, {
					Name: "b",
				},
			}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setResourceRequest(tt.args.pod, tt.args.podNetworks, tt.args.enableTrunk)
			assert.Equal(t, tt.want, tt.args.pod)
		})
	}
}

func TestPodMatchSelectorReturnsTrueWhenLabelsMatch(t *testing.T) {
	labelSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{"key": "value"},
	}
	labelsSet := labels.Set{"key": "value"}

	result, err := PodMatchSelector(labelSelector, labelsSet)
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestPodMatchSelectorReturnsFalseWhenLabelsDoNotMatch(t *testing.T) {
	labelSelector := &metav1.LabelSelector{
		MatchLabels: map[string]string{"key": "value"},
	}
	labelsSet := labels.Set{"key": "different"}

	result, err := PodMatchSelector(labelSelector, labelsSet)
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestPodMatchSelectorReturnsErrorForInvalidLabelSelector(t *testing.T) {
	labelSelector := &metav1.LabelSelector{
		MatchExpressions: []metav1.LabelSelectorRequirement{
			{
				Key:      "key",
				Operator: "InvalidOperator",
				Values:   []string{"value"},
			},
		},
	}
	labelsSet := labels.Set{"key": "value"}

	result, err := PodMatchSelector(labelSelector, labelsSet)
	assert.Error(t, err)
	assert.False(t, result)
}

func TestPreviousZoneReturnsEmptyWhenPodIsNotFixedName(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind: "ReplicaSet",
				},
			},
		},
	}

	zone, err := getPreviousZone(context.Background(), fakeClient, pod)
	assert.NoError(t, err)
	assert.Equal(t, "", zone)
}

func TestPreviousZoneReturnsEmptyWhenPodENINotFound(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind: "StatefulSet",
				},
			},
		},
	}

	zone, err := getPreviousZone(context.Background(), fakeClient, pod)
	assert.NoError(t, err)
	assert.Equal(t, "", zone)
}

func TestPreviousZoneReturnsEmptyWhenPodENIHasDeletionTimestamp(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
	}
	podENI := &v1beta1.PodENI{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:         "default",
			Name:              "test-pod",
			DeletionTimestamp: &metav1.Time{Time: time.Now()},
		},
	}

	_ = fakeClient.Create(context.Background(), podENI)

	zone, err := getPreviousZone(context.Background(), fakeClient, pod)
	assert.NoError(t, err)
	assert.Equal(t, "", zone)
}

func TestPreviousZoneReturnsEmptyWhenPodENIHasNoAllocations(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
	}
	podENI := &v1beta1.PodENI{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
		Spec: v1beta1.PodENISpec{
			Allocations: []v1beta1.Allocation{},
		},
	}

	_ = fakeClient.Create(context.Background(), podENI)

	zone, err := getPreviousZone(context.Background(), fakeClient, pod)
	assert.NoError(t, err)
	assert.Equal(t, "", zone)
}

func TestPreviousZoneReturnsZoneWhenPodENIHasAllocations(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
	}
	podENI := &v1beta1.PodENI{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
		Spec: v1beta1.PodENISpec{
			Zone: "aa-1a",
			Allocations: []v1beta1.Allocation{
				{},
			},
		},
	}

	_ = fakeClient.Create(context.Background(), podENI)

	zone, err := getPreviousZone(context.Background(), fakeClient, pod)
	assert.NoError(t, err)
	assert.Equal(t, "aa-1a", zone)
}

func TestMatchOnePodNetworkingReturnsNilWhenNoPodNetworkings(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
	}

	podNetworking, err := matchOnePodNetworking(context.Background(), "default", fakeClient, pod)
	assert.NoError(t, err)
	assert.Nil(t, podNetworking)
}

func TestMatchOnePodNetworkingReturnsNilWhenNoMatchingPodNetworking(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
		},
		&v1beta1.PodNetworking{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-podnetworking",
			},
			Spec: v1beta1.PodNetworkingSpec{
				Selector: v1beta1.Selector{
					PodSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"key": "different"},
					},
				},
			},
			Status: v1beta1.PodNetworkingStatus{
				Status: v1beta1.NetworkingStatusReady,
			},
		},
	).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
			Labels:    map[string]string{"key": "value"},
		},
	}

	result, err := matchOnePodNetworking(context.Background(), "default", fakeClient, pod)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestMatchOnePodNetworkingReturnsPodNetworkingWhenMatchingPodSelector(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
		},
		&v1beta1.PodNetworking{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-podnetworking",
			},
			Spec: v1beta1.PodNetworkingSpec{
				Selector: v1beta1.Selector{
					PodSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"key": "value"},
					},
				},
			},
			Status: v1beta1.PodNetworkingStatus{
				Status: v1beta1.NetworkingStatusReady,
			},
		},
	).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
			Labels:    map[string]string{"key": "value"},
		},
	}

	result, err := matchOnePodNetworking(context.Background(), "default", fakeClient, pod)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-podnetworking", result.Name)
}

func TestMatchOnePodNetworkingReturnsPodNetworkingWhenMatchingNamespaceSelector(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:   "default",
				Labels: map[string]string{"key": "value"},
			},
		},
		&v1beta1.PodNetworking{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-podnetworking",
			},
			Spec: v1beta1.PodNetworkingSpec{
				Selector: v1beta1.Selector{
					NamespaceSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"key": "value"},
					},
				},
			},
			Status: v1beta1.PodNetworkingStatus{
				Status: v1beta1.NetworkingStatusReady,
			},
		},
	).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
		},
	}

	result, err := matchOnePodNetworking(context.Background(), "default", fakeClient, pod)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-podnetworking", result.Name)
}

func TestMatchOnePodNetworkingReturnsNilWhenPodNetworkingNotReady(t *testing.T) {
	scheme := runtime.NewScheme()
	_ = corev1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	fakeClient := fake.NewClientBuilder().WithScheme(scheme).WithObjects(
		&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: "default",
			},
		},
		&v1beta1.PodNetworking{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test-podnetworking",
			},
			Spec: v1beta1.PodNetworkingSpec{
				Selector: v1beta1.Selector{
					PodSelector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"key": "value"},
					},
				},
			},
			Status: v1beta1.PodNetworkingStatus{
				Status: v1beta1.NetworkingStatusFail,
			},
		},
	).Build()

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name:      "test-pod",
			Labels:    map[string]string{"key": "value"},
		},
	}

	result, err := matchOnePodNetworking(context.Background(), "default", fakeClient, pod)
	assert.NoError(t, err)
	assert.Nil(t, result)
}
