

package controllers

import (
	"github.com/go-logr/logr"

	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TestChaosReconciler reconciles aTestChaos object
type TestChaosReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=chaos-mesh.org,resources=testchaos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=chaos-mesh.org,resources=testchaos/status,verbs=get;update;patch

func (r *TestChaosReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("reconciler", "testchaos")

	//  the main logic of `TestChaos`, it prints a log `Hello World!` and returns nothing.
	logger.Info("Hello World!")

	return ctrl.Result{}, nil
}

func (r *TestChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		//exports `TestChaos` object, which represents the yaml schema content the user applies.
		For(&chaosmeshv1alpha1.TestChaos{}).
		Complete(r)
}
