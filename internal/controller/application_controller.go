package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ApplicationReconciler reconciles a Application object
type ApplicationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=application.app.example.com,resources=applications,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=application.app.example.com,resources=applications/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=application.app.example.com,resources=applications/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ApplicationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := r.Log.WithValues("application", req.NamespacedName)
	l.Info("Reconciling Application")

	var app applicationv1alpha1.Application
	if err := r.Get(ctx, req.NamespacedName, &app); err != nil {
		if client.IgnoreNotFound(err) != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	finalizer := "your-finalizer-name"

	if !controllerutil.ContainsFinalizer(&app, finalizer) {
		controllerutil.AddFinalizer(&app, finalizer)
		if err := r.Update(ctx, &app); err != nil {
			return reconcile.Result{}, err
		}
	}

	if !app.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, &app)
	}

	return r.reconcileCreate(ctx, &app)
}

func (r *ApplicationReconciler) reconcileCreate(ctx context.Context, app *minettodevv1alpha1.Application) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	l.Info("Creating deployment")
	err := r.createOrUpdateDeployment(ctx, app)
	if err != nil {
		return ctrl.Result{}, err
	}
	l.Info("Creating service")
	err = r.createService(ctx, app)
	if err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *ApplicationReconciler) createOrUpdateDeployment(ctx context.Context, app *minettodevv1alpha1.Application) error {
	// Add implementation for creating or updating the deployment
}

func (r *ApplicationReconciler) createService(ctx context.Context, app *minettodevv1alpha1.Application) error {
	// Add implementation for creating the service
}

func (r *ApplicationReconciler) reconcileDelete(ctx context.Context, app *minettodevv1alpha1.Application) (ctrl.Result, error) {
	// Add implementation for reconciling deletion
}

// SetupWithManager sets up the controller with the Manager.
func (r *ApplicationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&applicationv1alpha1.Application{}).
		Complete(r)
}
