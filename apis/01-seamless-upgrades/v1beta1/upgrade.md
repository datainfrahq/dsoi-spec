## Upgrades

- The {{app}} resource is defined by Group: dsoi.datainfra.io, Version: v1beta1 and follows the defined principles.

- When upgrading a cluster with multiple node types, rolling upgrade can be done in parallel, but it is not the preferred method for stateful applications.   Nodes can also be upgraded in the order defined in the CR. If any node crashes during an upgrade, the operator halts the upgrade.

- Cluster installation should be parallel in any ideal scenario, and the first installation can be calculated using generation CrObject.Generation. Upgrades should be incremental across different node types, and the order defined in the CR should be easily configurable.

- A structured event should be emitted for each action on an upgrade to track the lifecycle of an upgrade. The k8s.io/client-go/tools/record package can be used to structure events.

- For large clusters, a CR status should have fields for timestamps on when the upgrade started, finished, and the total time taken to upgrade a cluster. In case of a crash during upgrades, ideal scenarios are to wait for manual intervention, get alerted or rollback after a timeout defined in the operator spec.

- The statefulset object status currentRevision and update revision can be used to calculate if the statefulset has reconciled to the desired state. The deployment object status ReadReplicas and Replicas can be used to calculate if the deployment has reconciled to the desired state.

- When dealing with multiple node types, each node type should have the flexibility to rollout individually.

- Tradeoffs on rollback: if the previous state was not a crashed state, the operator should not change the desired state in the CR on revert, as its core function is to reconcile current to the desired state.






