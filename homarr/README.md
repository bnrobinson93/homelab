# Overview

Homarr is a new tab page that allows for widgets and URL saving.

# How Charts Work

There are a few files:

`Chart.yaml` - contains dependencies, etc

`values.yaml` - specifies default values

`templates` folder - contains the k8s manifest

# Use

## Install

To install, run the below:

```sh
helm repo add oben01 https://oben01.github.io/charts
helm repo update
helm install homarr oben01/homarr # this will go to current namespace
helm install homarr oben01/homarr --namespace homar --create-namespace
```

## Upgrade

To get current values:

```sh
helm show values oben01/homarr > homarr.yaml
```

Now, specify your own values using this:

```sh
helm upgrade -n homarr homarr oben01/homarr --values=values.yaml
```

## View

```sh
helm ls -A # view all helm charts in all namespaces
```

## Uninstall

```sh
helm uninstall -n homarr homarr
```

# Creating your own helm chart

Create a Chart.yaml in the repo, along with values.yaml. In a templates folder, create the k8s manifest. The templates manifests should use {{ .values.** }} syntax. Something like [this](https://github.com/oben01/charts/blob/main/charts/homarr/templates/homarr-dc.yaml):

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name {{ include "homarr.fullname" . }}
  labels:
    {{- include "homarr.labels" . | ninden 4 }}
spec:
  {{- if not .Values.replicaCount }}
  replicas: {{ .Values.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "homarr.selectorLabels" . | ninden 6 }}}}
  template:
    metadata:
      {{- with .Values.podAnnotations }}
      annotations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      labels:
        {{- include "homarr.labels" . | nindent 8 }}
    {{- with .Values.podLabels }}
        {{- toYaml . | nindent 8 }}
        {{- end }}
    spec:
      {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: {{ .Values.service.port }}
              protocol: TCP
          livenessProbe:
{{ toYaml .Values.livenessProbe | indent 12 }}
            initialDelaySeconds: {{ .Values.livenessProbe.initialDelaySeconds }}
            timeoutSeconds: {{ .Values.livenessProbe.timeoutSeconds }}
            periodSeconds: {{ .Values.livenessProbe.periodSeconds }}
            failureThreshold: {{ .Values.livenessProbe.failureThreshold }}
          readinessProbe:
{{ toYaml .Values.readinessProbe | indent 12 }}
            initialDelaySeconds: {{ .Values.readinessProbe.initialDelaySeconds }}
            timeoutSeconds: {{ .Values.readinessProbe.timeoutSeconds }}
            periodSeconds: {{ .Values.readinessProbe.periodSeconds }}
            failureThreshold: {{ .Values.readinessProbe.failureThreshold }}
          env:
           {{- range $key, $value := .Values.env }}
           - name: {{ $key }}
             value: "{{ $value }}"
           {{- end }}
           {{- range $key, $value := .Values.envSecrets }}
           {{- if (and (ne $value nil) (and (ne $value.key nil) (ne $value.name nil))) }}
           - name: {{ $key }}
             valueFrom:
               secretKeyRef:
                 key: {{ $value.key }}
                 name: {{ $value.name }}
           {{- end }}
           {{- end }}
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
          volumeMounts:
          {{- range .Values.persistence }}
          {{- if .enabled }}
           - name: {{ .name }}
             mountPath: {{ .mountPath }}
          {{- end }}
          {{- end }}
      volumes:
        {{- range .Values.persistence }}
        {{- if .enabled }}
        - name: {{ .name }}
          persistentVolumeClaim:
           claimName: {{ .name }}
        {{- end }}
        {{- end }}
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
```
