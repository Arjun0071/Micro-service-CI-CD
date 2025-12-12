{{- define "user-service.name" -}}
{{ .Chart.Name }}
{{- end }}

{{- define "user-service.fullname" -}}
{{ .Release.Name }}-{{ .Chart.Name }}
{{- end }}

