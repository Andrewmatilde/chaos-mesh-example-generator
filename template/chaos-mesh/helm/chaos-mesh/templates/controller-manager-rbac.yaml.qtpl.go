// Code generated by qtc from "controller-manager-rbac.yaml.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:1
package templates

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:1
import "strings"

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:2
func StreamControllerManagerRbacYaml(qw422016 *qt422016.Writer, name string) {
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:2
	qw422016.N().S(`
{{- if .Values.rbac.create }}
kind: ServiceAccount
apiVersion: v1
metadata:
  namespace: {{ .Release.Namespace }}
  name: {{ .Values.controllerManager.serviceAccount }}
  labels:
    app.kubernetes.io/name: {{ template "chaos-mesh.name" . }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/component: controller-manager
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+"  "_" }}
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: {{ .Release.Name }}:chaos-controller-manager
  labels:
    app.kubernetes.io/name: {{ template "chaos-mesh.name" . }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/component: controller-manager
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+"  "_" }}
rules:
{{- if .Values.clusterScoped }}
- apiGroups: [""]
  resources:
  - services
  - events
  - namespaces
  verbs: ["*"]
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["create", "get", "list", "watch", "update"]
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["get", "list", "watch", "create", "update", "delete"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["persistentvolumeclaims"]
  verbs: ["get", "list", "watch", "create", "update", "delete"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch","update", "delete"]
- apiGroups: ["apps"]
  resources: ["statefulsets"]
  verbs: ["*"]
{{- end }}
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["persistentvolumes"]
  verbs: ["get", "list", "watch", "patch","update"]
- apiGroups: ["certificates.k8s.io"]
  resources: ["certificatesigningrequests", "certificatesigningrequests/approval"]
  verbs: ["get", "delete", "create", "update"]
- apiGroups: ["certificates.k8s.io"]
  resources:
    - "signers"
  resourceNames:
    - "kubernetes.io/legacy-unknown"
  verbs: ["approve"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "create", "list", "watch", "update", "delete"]
- apiGroups: ["admissionregistration.k8s.io"]
  resources: ["mutatingwebhookconfigurations","validatingwebhookconfigurations"]
  verbs: ["get", "create", "delete", "update", "patch"]
- apiGroups: ["chaos-mesh.org"]
  resources:
    - podchaos
    - networkchaos
    - iochaos
    - timechaos
    - kernelchaos
    - stresschaos
    - `)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:85
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:85
	qw422016.N().S(`chaos
  verbs: ["*"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: {{ .Release.Name }}:chaos-controller-manager
  labels:
    app.kubernetes.io/name: {{ template "chaos-mesh.name" . }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/component: controller-manager
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+"  "_" }}
subjects:
- kind: ServiceAccount
  name: {{ .Values.controllerManager.serviceAccount }}
  namespace: {{ .Release.Namespace }}
roleRef:
  kind: ClusterRole
  name: {{ .Release.Name }}:chaos-controller-manager
  apiGroup: rbac.authorization.k8s.io
{{- if (not .Values.clusterScoped) }}
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  namespace: {{ .Release.Namespace }}
  name: {{ .Release.Name }}:chaos-controller-manager
  labels:
    app.kubernetes.io/name: {{ template "chaos-mesh.name" . }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/component: controller-manager
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+"  "_" }}
rules:
- apiGroups: [""]
  resources:
    - services
    - events
    - namespaces
  verbs: ["*"]
- apiGroups: [""]
  resources: ["endpoints"]
  verbs: ["create", "get", "list", "watch", "update"]
- apiGroups: ["batch"]
  resources: ["jobs"]
  verbs: ["get", "list", "watch", "create", "update", "delete"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["persistentvolumeclaims"]
  verbs: ["get", "list", "watch", "create", "update", "delete"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get", "list", "watch","update", "delete"]
- apiGroups: ["apps"]
  resources: ["statefulsets"]
  verbs: ["*"]
- apiGroups: ["chaos-mesh.org"]
  resources:
  - podchaos
  - networkchaos
  - iochaos
  - timechaos
  - kernelchaos
  - stresschaos
  - `)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:152
	qw422016.E().V(strings.ToLower(name))
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:152
	qw422016.N().S(`chaos
  verbs: ["*"]
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  namespace: {{ .Release.Namespace }}
  name: {{ .Release.Name }}:chaos-controller-manager
  labels:
    app.kubernetes.io/name: {{ template "chaos-mesh.name" . }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/component: controller-manager
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+"  "_" }}
subjects:
- kind: ServiceAccount
  name: {{ .Values.controllerManager.serviceAccount }}
roleRef:
  kind: Role
  name: {{ .Release.Name }}:chaos-controller-manager
  apiGroup: rbac.authorization.k8s.io
{{- end }}
{{- end }}
`)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
}

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
func WriteControllerManagerRbacYaml(qq422016 qtio422016.Writer, name string) {
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	qw422016 := qt422016.AcquireWriter(qq422016)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	StreamControllerManagerRbacYaml(qw422016, name)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	qt422016.ReleaseWriter(qw422016)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
}

//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
func ControllerManagerRbacYaml(name string) string {
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	qb422016 := qt422016.AcquireByteBuffer()
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	WriteControllerManagerRbacYaml(qb422016, name)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	qs422016 := string(qb422016.B)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	qt422016.ReleaseByteBuffer(qb422016)
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
	return qs422016
//line template/chaos-mesh/helm/chaos-mesh/templates/controller-manager-rbac.yaml.qtpl:175
}
