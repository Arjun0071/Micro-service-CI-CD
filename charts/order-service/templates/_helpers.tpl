{{- define "order-service.name" -}}
{{ .Chart.Name }}
{{- end }}

{{- define "order-service.fullname" -}}
{{ .Release.Name }}-{{ .Chart.Name }}
{{- end }}

