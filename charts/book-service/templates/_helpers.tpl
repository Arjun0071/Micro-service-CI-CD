{{- define "book-service.name" -}}
{{ .Chart.Name }}
{{- end }}

{{- define "book-service.fullname" -}}
{{ include "book-service.name" . }}-{{ .Release.Name }}
{{- end }}

