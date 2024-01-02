package controller

import (
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Reconciler for nginx configmap operator
type Reconciler struct {
	client.Client
	logger logr.Logger
	scheme *runtime.Scheme
}

// NewReconciler generates a new reconciler
func NewReconciler(mgr manager.Manager) *Reconciler {
	return &Reconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}
}
