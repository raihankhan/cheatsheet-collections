`Kubernetes` is a portable, extensible, open-source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes services, support, and tools are widely available.


K8s as an abbreviation results from counting the eight letters between the "K" and the "s".


A Kubernetes cluster consists of a set of worker machines, called `nodes`, that run containerized applications. Every cluster has at least one worker node.

### Node Components
Node components run on every node, maintaining running pods and providing the Kubernetes runtime environment.

#### pods
Pods are the smallest deployable units of computing that you can create and manage in Kubernetes.

A Pod (as in a pod of whales or pea pod) is a group of one or more containers, with shared storage and network resources, and a specification for how to run the containers. A Pod's contents are always co-located and co-scheduled, and run in a shared context. A Pod models an application-specific "logical host": it contains one or more application containers which are relatively tightly coupled. 
In terms of `Docker` concepts, `a Pod` is similar to a group of Docker containers with shared namespaces and shared filesystem volumes.

Kubernetes is designed to accommodate configurations that meet all of the following criteria:

No more than `110 pods` per node

No more than `5000 nodes`

No more than `150000 total pods`

No more than `300000 total containers`

You can scale your cluster by adding or removing nodes. The way you do this depends on how your cluster is deployed.



#### kubelet
An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod.

The kubelet takes a set of PodSpecs that are provided through various mechanisms and ensures that the containers described in those PodSpecs are running and healthy. The kubelet doesn't manage containers which were not created by Kubernetes.

#### kube-proxy
kube-proxy is a network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.

kube-proxy maintains network rules on nodes. These network rules allow network communication to your Pods from network sessions inside or outside of your cluster.

kube-proxy uses the operating system packet filtering layer if there is one and it's available. Otherwise, kube-proxy forwards the traffic itself.

#### Container runtime
The container runtime is the software that is responsible for running containers.

Kubernetes supports several container runtimes: Docker, containerd, CRI-O, and any implementation of the Kubernetes CRI (Container Runtime Interface).

#### Addons
Addons use Kubernetes resources (DaemonSet, Deployment, etc) to implement cluster features. Because these are providing cluster-level features, namespaced resources for addons belong within the kube-system namespace.

Selected addons are described below; for an extended list of available addons, please see Addons.

#### DNS
While the other addons are not strictly required, all Kubernetes clusters should have cluster DNS, as many examples rely on it.

Cluster DNS is a DNS server, in addition to the other DNS server(s) in your environment, which serves DNS records for Kubernetes services.

Containers started by Kubernetes automatically include this DNS server in their DNS searches.

#### Web UI (Dashboard)
Dashboard is a general purpose, web-based UI for Kubernetes clusters. It allows users to manage and troubleshoot applications running in the cluster, as well as the cluster itself.

#### Container Resource Monitoring
Container Resource Monitoring records generic time-series metrics about containers in a central database, and provides a UI for browsing that data.

#### Cluster-level Logging
A cluster-level logging mechanism is responsible for saving container logs to a central log store with search/browsing interface.

 #### Deploying a `kubernetes` clustser using kind

Create A cluster : `kind create cluster --name cluster-name`

Delete A cluster : `kind delete cluster --name cluster-name`

To interact with a specific cluster:  `kubectl cluster-info --context kind-cluster-name`

To View Active clusters : `kind get clusters`

list out all of the nodes : `kubectl get nodes`

view Kubectl version in json :  `kubectl version -o json`

Describe a node in details : `kubectl describe nodes node-name`

! [Udemy CourSe](https://www.youtube.com/watch?v=2CAU4xWdKVM&list=PLMPZQTftRCS8Pp4wiiUruly5ODScvAwcQ&index=3)