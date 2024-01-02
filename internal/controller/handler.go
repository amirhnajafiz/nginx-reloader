package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/opdev/subreconciler"
	"k8s.io/api/apps/v1beta1"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
		// configmap not found, create it
		r.logger.Info(fmt.Sprintf("Configmap %s in namespace %s not found!", req.Name, req.Namespace))
		return r.CreateDeployment(ctx)
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	default:
		// check delete event
		if r.configmap.ObjectMeta.DeletionTimestamp != nil {
			return r.DeleteDeployment(ctx)
		}
	}

	return r.UpdateDeployment(ctx)
}

func (r *Reconciler) UpdateDeployment(ctx context.Context) (ctrl.Result, error) {
	deployment := &v1beta1.Deployment{}
	key := client.ObjectKey{
		Name:      fmt.Sprintf("%s-nginx-deployment", r.configmap.Name),
		Namespace: r.configmap.Namespace,
	}

	// get deployment object
	switch err := r.Get(ctx, key, deployment); {
	case apierrors.IsNotFound(err):
		// deployment not found
		r.logger.Info(fmt.Sprintf("Deployment %s in namespace %s not found!", key.Name, key.Namespace))
		return subreconciler.Evaluate(subreconciler.DoNotRequeue())
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	// rollout restart
	if deployment.Spec.Template.ObjectMeta.Annotations == nil {
		deployment.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	deployment.Spec.Template.ObjectMeta.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

	// update deployment
	if err := r.Update(ctx, deployment); err != nil {
		r.logger.Error(err, "failed to update deployment")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}

func (r *Reconciler) CreateDeployment(ctx context.Context) (ctrl.Result, error) {
	name := fmt.Sprintf("%s-nginx-deployment", r.configmap.Name)
	replicas := int32(1)
	deployment := &v1beta1.Deployment{
		ObjectMeta: ctrl.ObjectMeta{
			Name:        name,
			Namespace:   r.configmap.Namespace,
			Labels:      r.configmap.Labels,
			Annotations: r.configmap.Annotations,
		},
		Spec: v1beta1.DeploymentSpec{
			Selector: &v12.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Replicas: &replicas,
			Template: core.PodTemplateSpec{
				ObjectMeta: ctrl.ObjectMeta{
					Labels: map[string]string{
						"app":  name,
						"type": "nginx",
					},
					Annotations: make(map[string]string),
				},
				Spec: core.PodSpec{
					Containers: []core.Container{
						{
							Name:  name,
							Image: "nginx:1.7.9",
							Ports: []core.ContainerPort{
								{
									HostPort:      80,
									ContainerPort: 80,
								},
							},
							VolumeMounts: []core.VolumeMount{
								{
									Name:      "data",
									MountPath: "/usr/share/nginx/html",
								},
							},
							Resources: core.ResourceRequirements{
								Limits: core.ResourceList{
									core.ResourceCPU:    resource.Quantity{Format: "250m"},
									core.ResourceMemory: resource.Quantity{Format: "64Mi"},
								},
								Requests: core.ResourceList{
									core.ResourceCPU:    resource.Quantity{Format: "250m"},
									core.ResourceMemory: resource.Quantity{Format: "64Mi"},
								},
							},
						},
					},
					Volumes: []core.Volume{
						{
							Name: "data",
							VolumeSource: core.VolumeSource{
								ConfigMap: &core.ConfigMapVolumeSource{
									LocalObjectReference: core.LocalObjectReference{
										Name: r.configmap.Name,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if err := r.Create(ctx, deployment); err != nil {
		r.logger.Error(err, "failed to create deployment")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}

func (r *Reconciler) DeleteDeployment(ctx context.Context) (ctrl.Result, error) {
	deployment := &v1beta1.Deployment{}
	key := client.ObjectKey{
		Name:      fmt.Sprintf("%s-nginx-deployment", r.configmap.Name),
		Namespace: r.configmap.Namespace,
	}

	// get deployment object
	switch err := r.Get(ctx, key, deployment); {
	case apierrors.IsNotFound(err):
		// deployment not found
		r.logger.Info(fmt.Sprintf("Deployment %s in namespace %s not found!", key.Name, key.Namespace))
		return subreconciler.Evaluate(subreconciler.DoNotRequeue())
	case err != nil:
		// error in fetch
		r.logger.Error(err, "failed to fetch object")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	// delete deployment
	if err := r.Delete(ctx, deployment); err != nil {
		r.logger.Error(err, "failed to delete deployment")
		return subreconciler.Evaluate(subreconciler.Requeue())
	}

	return subreconciler.Evaluate(subreconciler.DoNotRequeue())
}
