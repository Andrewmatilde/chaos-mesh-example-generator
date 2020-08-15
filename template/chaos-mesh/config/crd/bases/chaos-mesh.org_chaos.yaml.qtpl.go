// Code generated by qtc from "chaos-mesh.org_chaos.yaml.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:1
package bases

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:1
import "github.com/iancoleman/strcase"

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:2
import "strings"

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:3
func StreamCRDYaml(qw422016 *qt422016.Writer, name string) {
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:3
	qw422016.N().S(`
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.2.5
  creationTimestamp: null
  name: `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:11
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:11
	qw422016.N().S(`chaos.chaos-mesh.org
spec:
  group: chaos-mesh.org
  names:
    kind: `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:15
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:15
	qw422016.N().S(`Chaos
    listKind: `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:16
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:16
	qw422016.N().S(`ChaosList
    plural: `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:17
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:17
	qw422016.N().S(`chaos
    singular: `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:18
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:18
	qw422016.N().S(`chaos
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:22
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:22
	qw422016.N().S(`Chaos is the Schema for the `)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:22
	qw422016.E().V(strcase.ToCamel(name))
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:22
	qw422016.N().S(`chaos API
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
      type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
}

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
func WriteCRDYaml(qq422016 qtio422016.Writer, name string) {
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	StreamCRDYaml(qw422016, name)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
}

//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
func CRDYaml(name string) string {
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	WriteCRDYaml(qb422016, name)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	qs422016 := string(qb422016.B)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
	return qs422016
//line template/chaos-mesh/config/crd/bases/chaos-mesh.org_chaos.yaml.qtpl:48
}