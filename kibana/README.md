elasticsearchversion - xpack-7.14.0

kibana - docker.elastic.co/kibana/kibana:7.14.0 [elasticstk]

The elasticsearch version and kibana version are recomended to keep the same.

## Elasticsearch yaml


```yaml
apiVersion: kubedb.com/v1alpha2
kind: Elasticsearch
metadata:
  name: es-topology
  namespace: elastic
spec:
  enableSSL: false
  version: xpack-7.14.0
  storageType: Durable
  topology:
    master:
      replicas: 1
      storage:
        storageClassName: "standard"
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
    data:
      replicas: 2
      storage:
        storageClassName: "standard"
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
    ingest:
      replicas: 1
      storage:
        storageClassName: "standard"
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 1Gi
```

## kibana-config.yaml

elasticsearch host name format - name_of_es_CR.namespace.svc:port_where_es_is_exposed

get elasticsearch.username - `kubectl get secret -n elastic es-topology-elastic-cred -o jsonpath='{.data.username}' | base64 -d`

get elasticsearch.password - `kubectl get secret -n elastic es-topology-elastic-cred -o jsonpath='{.data.password}' | base64 -d`

this is a configmap for creating a superuser at kibana with elasticsearch.username and elasticsearch.password

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: kibana-config
  namespace: kibana
data:
  kibana.yml: |
    ---
    server.name: kibana
    server.port: 5601
    elasticsearch.hosts: ["http://es-topology.elastic.svc:9200"]
    elasticsearch.username: elastic
    elasticsearch.password: '8rjzwRXYQfyI-l8q'

```

## kibana-deployment.yaml



```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kibana
  namespace: kibana
  labels:
    component: kibana
spec:
  replicas: 1
  selector:
    matchLabels:
     component: kibana
  template:
    metadata:
      labels:
        component: kibana
    spec:
      containers:
      - name: kibana
        image: docker.elastic.co/kibana/kibana:7.14.0
        ports:
        - containerPort: 5601
          name: http
        volumeMounts:
          - name: config
            mountPath: /usr/share/kibana/config
            readOnly: true
      volumes:
        - name: config
          configMap:
            name: kibana-config
```


<!-----
NEW: Check the "Suppress top comment" option to remove this info from the output.

Conversion time: 0.678 seconds.


Using this Markdown file:

1. Paste this output into your source file.
2. See the notes and action items below regarding this conversion run.
3. Check the rendered output (headings, lists, code blocks, tables) for proper
   formatting and use a linkchecker before you publish this page.

Conversion notes:

* Docs to Markdown version 1.0β31
* Thu Nov 18 2021 03:53:04 GMT-0800 (PST)
* Source doc: Elastic Search
----->


You can use this gist to use the kubebuilder client from any of your go code.

This avoids needing to use multiple clients like we do today.   [kubebuilder_client.go](https://gist.github.com/tamalsaha/6215f161788d0293c066d1afd88eb0d4)


## ElasticSearch resources

[elasticsearch-tutorial](https://logz.io/blog/elasticsearch-tutorial/)

[Sharding-in-elasticsearch](https://codingexplained.com/coding/elasticsearch/understanding-sharding-in-elasticsearch) \


[Go-elasticsearch](https://github.com/elastic/go-elasticsearch)

[Elasticsearch-go-developers-guide](https://developer.okta.com/blog/2021/04/23/elasticsearch-go-developers-guide) \


[http-settings](https://www.elastic.co/guide/en/elasticsearch/reference/7.15/modules-network.html#http-settings)

[webhookValidation](https://medium.com/swlh/kubernetes-validating-webhook-implementation-60f3352b66a)

[Node-affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity)

[Pod-Affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity)

[how-finalizers-work](https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers/#how-finalizers-work)

[Understand-helm-upgrade-flags-reset-values-reuse-value](https://medium.com/@kcatstack/understand-helm-upgrade-flags-reset-values-reuse-values-6e58ac8f127e)

[data_tiers](https://www.elastic.co/guide/en/elasticsearch/reference/current/data-tiers.html)

[shard filtering](https://www.alibabacloud.com/blog/allocate-indexes-to-hot-and-warm-nodes-in-elasticsearch-through-shard-filtering_597456)

[opensearch vs elastic stk vs open source](https://aws.plainenglish.io/the-difference-between-elasticsearch-open-distro-and-opensearch-d43c9a2c31b1)


## Codebase

[init() function](https://tutorialedge.net/golang/the-go-init-function/)

[cobra cli](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177)

[Informers](https://dev.to/davidsbond/go-creating-dynamic-kubernetes-informers-1npi)

Pkg / controller and distribution 


## [kubedb.com/elasticsearch](https://kubedb.com/docs/v2021.09.30/guides/elasticsearch/)



1. Clustering
    1. Combined cluster
    2. Topology cluster
2. Concepts
3. Helm chart elasticsearch : [elastic/helm-charts](https://github.com/elastic/helm-charts)
4. kubedb/ elasticsearch codebase : [kubedb/elasticsearch](https://github.com/kubedb/elasticsearch)

Elasticsearch spec

.

|--version : where docker images are specified

	      Name format: {Security Plugin Name}-{Application Version}-{Modification Tag}


           Samples: searchguard-7.9.3, xpack-7.9.1-v1, opendistro-1.12.0, etc.

|--kernelSettings: (optional) ****

|--disableSecurity: disable security plugins like safeguard, xpack etc

**|--internalUsers: **

**|--rolesmapping:**

** **

**|--topology:**

**|--terminationPolicy:**


    -** Wipeout :** if CR gets deleted, deletes every CR resources \



    - **Halt: **if CR gets deleted, keeps PVC and secrets, so that if we deploy CR again, it recreates ES cluster from existing auth_secrets and PVCs.


    -**DoNotTerminate: **if CR gets deleted, prevents deletion

	

Helm Charts - elasticsearch

Templates-

**PDB** - You can specify only one of maxUnavailable and minAvailable in a single PodDisruptionBudget.


    [pdb-example](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/#pdb-example)


    [specifying-a-pod disruption budget](https://kubernetes.io/docs/tasks/run-application/configure-pdb/#specifying-a-poddisruptionbudget)

**PodSecurityPolicy -** is deprecated as of Kubernetes v1.21, and will be removed in v1.25.

	


    A Pod Security Policy is a cluster-level resource that controls security sensitive aspects of the pod specification. The PodSecurityPolicy objects define a set of conditions that a pod must run with in order to be accepted into the system, as well as defaults for the related fields. 


    [policy-reference](https://kubernetes.io/docs/concepts/policy/pod-security-policy/#policy-reference)

	

	[yq merge](https://mikefarah.gitbook.io/yq/v/v3.x/commands/merge)

	[kubedb elasticsearch-init](https://github.com/kubedb/elasticsearch-docker/tree/7.14.2-searchguard-v2021.11.10/elasticsearch-init)

	[kubedb elastic search custom config](https://kubedb.com/docs/v2021.09.30/guides/elasticsearch/configuration/overview/)

Kibana

[setup kibana with elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/7.15/security-minimal-setup.html)

[kibana configure settings](https://www.elastic.co/guide/en/kibana/current/settings.html)

[kibana config yml](https://github.com/elastic/kibana/blob/main/config/kibana.yml)

[kibana deployment](https://github.com/kubernetes/kubernetes/blob/master/cluster/addons/fluentd-elasticsearch/kibana-deployment.yaml)

[How-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-elasticsearch-fluentd-and-kibana-efk-logging-stack-on-kubernetes)

Kibana gist

[https://gist.github.com/chancez/048a8e5ec6f5049ee3f2356aac6fa1d4](https://gist.github.com/chancez/048a8e5ec6f5049ee3f2356aac6fa1d4)

kibana-plugin

[Kibana-searchguard-plugin-installation](https://docs.search-guard.com/latest/kibana-plugin-installation)

Opendistro kibana version compatibility

[opendistro version vs es version](https://opendistro.github.io/for-elasticsearch-docs/version-history/)

Kibana docker images:

[bitnami/kibana](https://hub.docker.com/r/bitnami/kibana)

[kibana by elastic team](https://hub.docker.com/_/kibana)

[amazon/opendistro-for-elasticsearch-kibana](https://hub.docker.com/r/amazon/opendistro-for-elasticsearch-kibana)

[opensearch kibana](https://hub.docker.com/r/opensearchproject/opensearch-dashboards)

[Searchguard-kibana](https://hub.docker.com/r/floragunncom/sg-kibana/tags)

Kibana distro docs:

[Searchguard](https://docs.search-guard.com/latest/kibana-plugin-installation)

[Opensearch](https://opensearch.org/docs/latest/dashboards/install/index/)

[Opendistro](https://opendistro.github.io/for-elasticsearch-docs/docs/install/docker/#sample-docker-compose-file)

[Elastic Stk](https://www.elastic.co/guide/en/kibana/current/index.html)

………………………………………………….

[kibana/current/xpack-security-authorization.html](https://www.elastic.co/guide/en/kibana/current/xpack-security-authorization.html)

[kibana-kubernetes](https://imti.co/kibana-kubernetes/)

Sidecar

[giantswarm/kibana-sidecar](https://github.com/giantswarm/kibana-sidecar)

[InMoment/kibana-sidecar](https://github.com/InMoment/kibana-sidecar)

[k8s-service-mesh-istio.html](https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-service-mesh-istio.html)
