// /*
// Copyright 2019 Alexander Eldeib.
// */

package controllers

import (
	"context"

	"github.com/alexeldeib/cerberus/pkg/configmap"
	"github.com/alexeldeib/cerberus/pkg/imds"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// ConfigMapController reconciles a managed identity
type ConfigMapController struct {
	client.Client
}

// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch

func (r *ConfigMapController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	var ns v1.Namespace
	if err := r.Get(ctx, req.NamespacedName, &ns); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	data, err := imds.New()
	if err != nil {
		return ctrl.Result{}, err
	}

	cm, err := configmap.New(data, "default")
	if err != nil {
		return ctrl.Result{}, err
	}

	_, err = controllerutil.CreateOrUpdate(ctx, r, cm, func() error {
		return nil
	})
	return ctrl.Result{}, err
}

func (r *ConfigMapController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Namespace{}).
		Owns(&corev1.ConfigMap{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 1}).
		Complete(r)
}
