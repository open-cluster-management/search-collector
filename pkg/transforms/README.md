# Data Model Documentation

Here we document the properties collected for the kubernetes resources and the relationships between those resources. This only covers the edges computed on the collector (intra-cluster).  Inter-cluster edges are computed in the aggregator.

## Resource Properties

- Properties that start with underscore `_` are only for internal use and won't be available for users to search.
- Common properties that we collect for any resource:
    - `kind (string), name (string), namespace (string), created (string), apigroup (string), apiversion (string), label ([]string)`
    - **Deprecated:** `selfLink`. It can be built from the properties above. We don't expect users to search for this.
- Each transform file had a BuildNode() function where we define which properties we want to extract an index for the resource.
- Our goal is to match the properties displayed from `oc get <resource> -o wide`, but we don't have a generic way to do this yet.

## Resource Relationships (Edges)

Each transform has a BuildEdges() function where we find other resources related to each resource.

### Common
Edges for any kubernetes resource.

- **(\*)-[OWNED_BY]->(\*)**
    - Extract owner references from the object's metadata.
- **(\*)-[DEFINED_BY]->(Deployable)**
    - Logic explained on [Deployable (AppDeployable) section](#deployable-appdeployable).
- **(\*)-[OWNED_BY]->(Release)**
    - Logic explained on [Helm Release section](#helm-release-apphelmcr).
- **(\*)-[DEPLOYED_BY]->(Subscription)**
    - Logic explained on [Subscription section](#subscription).

### Application

- **(Application)-[CONTAINS]->(Subscriptions)**
  - Use the annotation `apps.open-cluster-management.io/subscriptions` to link subscriptions associated to the application.
  - When this edge gets built, we capture the application info in the subscription's metadata as `_hostingApplication`. This is used in 2 ways:
    1. Whenever a new edge is built from any node(N) with the subscription(S) as the destination, an edge is added by default from that node(N) to the Application(A). In addition, an edge is added from that node(N) to the subscription's channels.
    2. Whenever a new edge is built from subscription(S) to any node(N), where the node is a deployable, placementrule or channel, an edge is built from the node(N) to the Application(A). So, the info in `_hostingApplication` is used in building edges from nodes to Application.
- **(Application)-[CONTAINS]->(Deployable)**
  - Use the annotation `apps.open-cluster-management.io/deployables` to link deployables associated to the application.


### Channel
- **(Channel)-[USES]->(ConfigMap)** OR **(Channel)-[USES]->(Secret)**
  - Extract from `Spec.ConfigMapRef.Name` or `Spec.SecretRef.Name`

- **(Channel)-[DEPLOYS]->(Deployable)**
  - If channel type is a helm repo, extract from spec.


### Deployable (AppDeployable)
- **(Deployable)-[PROMOTED_TO]-(Channel)**
  - Extract from `Spec.Channels`

- **(Deployable)-[REFERS_TO]-(PlacementRule)**
  - Extract from `Spec.Placement.PlacementRef.Name`

- **(\*)-[DEFINED_BY]->(Deployable)**
    > The logic to build this edge is on the common transform.
    - Use the annotation `apps.open-cluster-management.io/hosting-deployable` on any resource to link to the deployable that created the resource.
    - This is built as part of commonEdges(). The annotation `apps.open-cluster-management.io/hosting-deployable` is saved on each node as `_hostingDeployable`
    - If a node doesn't have the `apps.open-cluster-management.io/hosting-deployable` annotation, we will check recursively if its owner Node has the annotation and create the edge. For example, `(Pod)-[OwnedBy]->(ReplicaSet)` and `(ReplicaSet)-[OwnedBy]->(Deployment)` and the Deployment has `apps.open-cluster-management.io/hosting-deployable` or `apps.open-cluster-management.io/hosting-subscription` annotation, the pod and the replicaset will also have an edge to the deployable or subscription


### Helm Release (appHelmCR)
- **(HelmRelease)-[ATTACHED_TO]->(ConfigMap)**
  - Extract from `Repo.ConfigMapRef.Name`
- **(HelmRelease)-[ATTACHED_TO]->(Release)**
  - Extract from `ObjectMeta.Name`
- **(HelmRelease)-[ATTACHED_TO]->(Secret)**
  - Extract from `Repo.SecretRef.Name`


### HelmRelease (HelmReleaseResource)
- **(\*)-[OWNED_BY]->(Release)**
  - Reads the helm release manifest file to find resources, then link each resource to the HelmRelease resource.


### Pod
- **(Pod)-[ATTACHED_TO]->(ConfigMap)**
- **(Pod)-[ATTACHED_TO]->(Secret)**
- **(Pod)-[ATTACHED_TO]->(PersistentVolume)**
- **(Pod)-[ATTACHED_TO]->(PersistentVolumeClaim)**
- **(Pod)-[RUNS_ON]->(Node)**


### PersistentVolumeClaim
- **(PersistentVolumeClaim)-[BOUND_TO]->(PersistentVolume)**


### Service
- **(Service)-[USED_BY]->(Pod)**


### Subscription
- **(Subscription)-[TO]->(Channel)**
  - Extract from `Spec.Channel`

- **(Subscription)-[REFERS_TO]->(PlacementRule)**
  - Extract from `Spec.Placement.PlacementRef.Name`

- **(Subscription)-[SUBSCRIBES_TO]->(Deployable)**
  - Use the annotation `apps.open-cluster-management.io/deployables` on the subscription.

- **(\*)-[DEPLOYED_BY]->(Subscription)**
  - Use the annotation `apps.open-cluster-management.io/hosting-subscription` on any resource to link to the subscription that created the resource.
  - This is built as part of commonEdges(). The annotation "hosting-subscription" is saved on each node as "_hostingSubscription"
