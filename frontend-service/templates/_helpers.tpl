{{- define "frontend.name" -}}
{{ .Chart.Name }}
{{- end }}

{{- define "frontend.fullname" -}}
{{ .Release.Name }}-{{ .Chart.Name }}
{{- end }}

