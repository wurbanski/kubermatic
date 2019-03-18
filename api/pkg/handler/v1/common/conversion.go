package common

import (
	apiv1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"
	kubermaticapiv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"

	corev1 "k8s.io/api/core/v1"
)

func ConvertInternalSSHKeysToExternal(internalKeys []*kubermaticapiv1.UserSSHKey) []*apiv1.SSHKey {
	apiKeys := make([]*apiv1.SSHKey, len(internalKeys))
	for index, key := range internalKeys {
		apiKey := &apiv1.SSHKey{
			ObjectMeta: apiv1.ObjectMeta{
				ID:                key.Name,
				Name:              key.Spec.Name,
				CreationTimestamp: apiv1.NewTime(key.CreationTimestamp.Time),
			},
			Spec: apiv1.SSHKeySpec{
				Fingerprint: key.Spec.Fingerprint,
				PublicKey:   key.Spec.PublicKey,
			},
		}
		apiKeys[index] = apiKey
	}
	return apiKeys
}

// ConvertInternalEventToExternal converts Kubernetes Events to Kubermatic ones (used in the API).
func ConvertInternalEventToExternal(event corev1.Event) apiv1.Event {
	result := apiv1.Event{
		ObjectMeta: apiv1.ObjectMeta{
			ID:                string(event.ObjectMeta.UID),
			Name:              event.ObjectMeta.Name,
			CreationTimestamp: apiv1.NewTime(event.ObjectMeta.CreationTimestamp.Time),
		},
		Message:            event.Message,
		Type:               event.Type,
		InvolvedObjectName: event.InvolvedObject.Name,
		LastTimestamp:      apiv1.NewTime(event.LastTimestamp.Time),
		Count:              event.Count,
	}
	return result
}
