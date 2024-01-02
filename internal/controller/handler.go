package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/opdev/subreconciler"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Reconciler for nginx configmap operator
type Reconciler struct {
	client.Client
	configmap *v1.ConfigMap
	logger    logr.Logger
	scheme    *runtime.Scheme
}

// NewReconciler generates a new reconciler
func NewReconciler(mgr manager.Manager) *Reconciler {
	return &Reconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.logger = log.FromContext(ctx)
	r.configmap = &v1.ConfigMap{}

	// get configmap object
	switch err := r.Get(ctx, req.NamespacedName, r.configmap); {
	case apierrors.IsNotFound(err):
		// configmap not found
		r.logger.Info(fmt.Sprintf("Configmap %s in namespace %s not found!", req.Name, req.Namespace))
		return subreconciler.Evaluate(subreconciler.DoNotRequeue())
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default:
		// check delete event
		if r.configmap.ObjectMeta.DeletionTimestamp != nil {
			// delete the deployment related to that configmap
		}

		// check if exists update
		// else create
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
