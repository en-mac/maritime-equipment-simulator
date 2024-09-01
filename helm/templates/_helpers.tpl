{{- define "maritime-equipment-simulator.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "maritime-equipment-simulator.name" -}}
{{- printf "%s" .Release.Name -}}
{{- end -}}
