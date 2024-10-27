# prometheus-operator

![Version: 0.7.5](https://img.shields.io/badge/Version-0.7.5-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.30.0 || >= v1.30.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://prometheus-community.github.io/helm-charts | kube-prometheus-stack | 65.3.1 |

## Values

<table height="400px" >
	<thead>
		<th>Key</th>
		<th>Type</th>
		<th>Default</th>
		<th>Description</th>
	</thead>
	<tbody>
		<tr>
			<td id="kube-prometheus-stack--coreDns--enabled"><a href="./values.yaml#L22">kube-prometheus-stack.coreDns.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--enabled"><a href="./values.yaml#L6">kube-prometheus-stack.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--fullnameOverride"><a href="./values.yaml#L8">kube-prometheus-stack.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--grafana--enabled"><a href="./values.yaml#L11">kube-prometheus-stack.grafana.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--grafana--sidecar--dashboards--enableNewTablePanelSyntax"><a href="./values.yaml#L15">kube-prometheus-stack.grafana.sidecar.dashboards.enableNewTablePanelSyntax</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--kubeDns--enabled"><a href="./values.yaml#L18">kube-prometheus-stack.kubeDns.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--prometheusOperator--resources--limits--cpu"><a href="./values.yaml#L29">kube-prometheus-stack.prometheusOperator.resources.limits.cpu</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--prometheusOperator--resources--limits--memory"><a href="./values.yaml#L30">kube-prometheus-stack.prometheusOperator.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"512Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--prometheusOperator--resources--requests--cpu"><a href="./values.yaml#L32">kube-prometheus-stack.prometheusOperator.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"30m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--prometheusOperator--resources--requests--memory"><a href="./values.yaml#L33">kube-prometheus-stack.prometheusOperator.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kube-prometheus-stack--prometheusOperator--secretFieldSelector"><a href="./values.yaml#L25">kube-prometheus-stack.prometheusOperator.secretFieldSelector</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"type!=kubernetes.io/dockercfg,type!=kubernetes.io/service-account-token,type!=helm.sh/release.v1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)