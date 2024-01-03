package controller

import (
	v1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

func ignoreNonNginxConfigmaps() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(event event.CreateEvent) bool {
			if value, ok := event.Object.GetAnnotations()["snappcloud.io/nginx-operator"]; ok {
				return value == "enable"
			}

			return false
		},
		UpdateFunc: func(event event.UpdateEvent) bool {
			if value, ok := event.ObjectOld.GetAnnotations()["snappcloud.io/nginx-operator"]; ok {
				return value == "enable"
			}

			return false
		},
		DeleteFunc: func(event event.DeleteEvent) bool {
			if value, ok := event.Object.GetAnnotations()["snappcloud.io/nginx-operator"]; ok {
				return value == "enable"
			}

			return false
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1.ConfigMap{}).
		WithEventFilter(ignoreNonNginxConfigmaps()).
		Complete(r)
}
