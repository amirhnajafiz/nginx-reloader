package controller

import (
	v1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

const (
	resourceAnnotationKey   = "snappcloud.io/nginx-operator"
	resourceAnnotationValue = "enable"
)

// annotationCheck is a selector function for our main loop
func annotationCheck(annotations map[string]string) bool {
	if value, ok := annotations[resourceAnnotationKey]; ok {
		return value == resourceAnnotationValue
	}

	return false
}

// ignoreNonNginxConfigmaps is the main filter
func ignoreNonNginxConfigmaps() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(event event.CreateEvent) bool {
			return annotationCheck(event.Object.GetAnnotations())
		},
		UpdateFunc: func(event event.UpdateEvent) bool {
			return annotationCheck(event.ObjectOld.GetAnnotations()) || annotationCheck(event.ObjectNew.GetAnnotations())
		},
		DeleteFunc: func(event event.DeleteEvent) bool {
			return annotationCheck(event.Object.GetAnnotations())
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
