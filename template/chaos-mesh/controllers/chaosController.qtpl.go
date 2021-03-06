// Code generated by qtc from "chaosController.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/chaos-mesh/controllers/chaosController.qtpl:1
package controllers

//line template/chaos-mesh/controllers/chaosController.qtpl:1
import "github.com/iancoleman/strcase"

//line template/chaos-mesh/controllers/chaosController.qtpl:2
import "strings"

//line template/chaos-mesh/controllers/chaosController.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/chaos-mesh/controllers/chaosController.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/chaos-mesh/controllers/chaosController.qtpl:3
func StreamChaosContoller(qw422016 *qt422016.Writer, name string) {
//line template/chaos-mesh/controllers/chaosController.qtpl:3
	qw422016.N().S(`

package controllers

import (
	"github.com/go-logr/logr"

	chaosmeshv1alpha1 "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// `)
//line template/chaos-mesh/controllers/chaosController.qtpl:16
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:16
	qw422016.N().S(`ChaosReconciler reconciles a`)
//line template/chaos-mesh/controllers/chaosController.qtpl:16
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:16
	qw422016.N().S(`Chaos object
type `)
//line template/chaos-mesh/controllers/chaosController.qtpl:17
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:17
	qw422016.N().S(`ChaosReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=chaos-mesh.org,resources=`)
//line template/chaos-mesh/controllers/chaosController.qtpl:22
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:22
	qw422016.N().S(`chaos,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=chaos-mesh.org,resources=`)
//line template/chaos-mesh/controllers/chaosController.qtpl:23
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:23
	qw422016.N().S(`chaos/status,verbs=get;update;patch

func (r *`)
//line template/chaos-mesh/controllers/chaosController.qtpl:25
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:25
	qw422016.N().S(`ChaosReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	logger := r.Log.WithValues("reconciler", "`)
//line template/chaos-mesh/controllers/chaosController.qtpl:26
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:26
	qw422016.N().S(`chaos")

	//  the main logic of `)
//line template/chaos-mesh/controllers/chaosController.qtpl:26
	qw422016.N().S("`")
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S(`Chaos`)
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S("`")
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S(`, it prints a log `)
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S("`")
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S(`Hello World!`)
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S("`")
//line template/chaos-mesh/controllers/chaosController.qtpl:28
	qw422016.N().S(` and returns nothing.
	logger.Info("Hello World!")

	return ctrl.Result{}, nil
}

func (r *`)
//line template/chaos-mesh/controllers/chaosController.qtpl:34
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:34
	qw422016.N().S(`ChaosReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		//exports `)
//line template/chaos-mesh/controllers/chaosController.qtpl:34
	qw422016.N().S("`")
//line template/chaos-mesh/controllers/chaosController.qtpl:36
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:36
	qw422016.N().S(`Chaos`)
//line template/chaos-mesh/controllers/chaosController.qtpl:36
	qw422016.N().S("`")
//line template/chaos-mesh/controllers/chaosController.qtpl:36
	qw422016.N().S(` object, which represents the yaml schema content the user applies.
		For(&chaosmeshv1alpha1.`)
//line template/chaos-mesh/controllers/chaosController.qtpl:37
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/controllers/chaosController.qtpl:37
	qw422016.N().S(`Chaos{}).
		Complete(r)
}
`)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
}

//line template/chaos-mesh/controllers/chaosController.qtpl:40
func WriteChaosContoller(qq422016 qtio422016.Writer, name string) {
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	StreamChaosContoller(qw422016, name)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
}

//line template/chaos-mesh/controllers/chaosController.qtpl:40
func ChaosContoller(name string) string {
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	WriteChaosContoller(qb422016, name)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	qs422016 := string(qb422016.B)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
//line template/chaos-mesh/controllers/chaosController.qtpl:40
	return qs422016
//line template/chaos-mesh/controllers/chaosController.qtpl:40
}
