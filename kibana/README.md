You can use this gist to use the kubebuilder client from any of your go code.
This avoids needing to use multiple clients like we do today.   kubebuilder_client.go

ElasticSearch resources
elasticsearch-tutorial

Sharding-in-elasticsearch


Go-elasticsearch

Elasticsearch-go-developers-guide


http-settings

webhookValidation

Node-affinity

Pod-Affinity

how-finalizers-work

Understand-helm-upgrade-flags-reset-values-reuse-value

data_tiers

shard filtering

opensearch vs elastic stk vs open source
Codebase

init() function
cobra cli
Informers

Pkg / controller and distribution 

kubedb.com/elasticsearch

Clustering
Combined cluster
Topology cluster
Concepts
Helm chart elasticsearch : elastic/helm-charts
kubedb/ elasticsearch codebase : kubedb/elasticsearch

Elasticsearch spec
.
|--version : where docker images are specified
	      Name format: {Security Plugin Name}-{Application Version}-{Modification Tag}
       Samples: searchguard-7.9.3, xpack-7.9.1-v1, opendistro-1.12.0, etc.

|--kernelSettings: (optional) ****

|--disableSecurity: disable security plugins like safeguard, xpack etc

|--internalUsers: 

|--rolesmapping:
 
|--topology:

|--terminationPolicy:

- Wipeout : if CR gets deleted, deletes every CR resources


- Halt: if CR gets deleted, keeps PVC and secrets, so that if we deploy CR again, it recreates ES cluster from existing auth_secrets and PVCs.

-DoNotTerminate: if CR gets deleted, prevents deletion
	






Helm Charts - elasticsearch

Templates-

PDB - You can specify only one of maxUnavailable and minAvailable in a single PodDisruptionBudget.
pdb-example
specifying-a-pod disruption budget


PodSecurityPolicy - is deprecated as of Kubernetes v1.21, and will be removed in v1.25.
	
A Pod Security Policy is a cluster-level resource that controls security sensitive aspects of the pod specification. The PodSecurityPolicy objects define a set of conditions that a pod must run with in order to be accepted into the system, as well as defaults for the related fields. 
policy-reference

	
	yq merge
	kubedb elasticsearch-init
	kubedb elastic search custom config






















Kibana

setup kibana with elasticsearch
kibana configure settings
kibana config yml
kibana deployment
How-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes

Kibana gist
https://gist.github.com/chancez/048a8e5ec6f5049ee3f2356aac6fa1d4

kibana-plugin
Kibana-searchguard-plugin-installation

Opendistro kibana version compatibility
opendistro version vs es version


Kibana docker images:

bitnami/kibana
kibana by elastic team
amazon/opendistro-for-elasticsearch-kibana
opensearch kibana
Searchguard-kibana

Kibana distro docs:

Searchguard
Opensearch
Opendistro
Elastic Stk


………………………………………………….
kibana/current/xpack-security-authorization.html
kibana-kubernetes


Sidecar

giantswarm/kibana-sidecar
InMoment/kibana-sidecar
k8s-service-mesh-istio.html































