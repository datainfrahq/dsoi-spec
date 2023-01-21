# UPGRADE

- ```Group: dsoi.ballastdata.com```
- ```Version: v1beta1```
- ```Kind: {{ app }}Cluster```

- ```{{ app}}Cluster``` Resource defines a spec which based on the [principles](../../../PRINCIPLES.md) defined.

- When upgrading a cluster with multiple node types rolling upgrade can be done in parallel ie all nodes upgraded at once ( which is ideally not a preferred in stateful applications ), else all nodes can be rolled out in the order defined in the CR. If any node crashes during an upgrade, the operator halts the upgrade.

> When dealing with multiple node, in any ideal scenario, cluster installation should be parallel. The first installation can be easily calculated using generation ```CrObject.Generation```

> Upgrades should be incremental across different nodeTypes. Order defined in the CR should be easily configurable. To understand Order, refer to [terminology](../TERMINOLOGY.md).

> A structured event should be emitted for each action on an upgrade, to track the lifecycle of an upgrade. ```	"k8s.io/client-go/tools/record``` can be imported to structure events.

> Large clusters can take sufficient time to upgrade, a CR status should have fields for timestamps on when the upgrade started, finished and the total time taken to upgrade a cluster.

```
type MyAppStatus {
    // The last time this condition was updated.
    LastUpdateTime string `json:"lastUpdateTime,omitempty"`
    // Last time the condition transitioned from one status to another.
    LastTransitionTime string `json:"lastTransitionTime,omitempty"`
    // Time Taken to update a cluster
    TimeTakenToUpgrade string `json:"timeTakenToUpgrade,omitempty"
}
```

- Incase of a crash during upgrades, ideal scenarios are

1. Wait for manual intervention, get alerted.
2. Rollback after a timeout defined in the operator spec.

> Statefulset Object status ```currentRevision``` and ```update revision``` can be used to calculate if statefulset has reconciled to desired state.

- Deployment Object status ```ReadReplicas``` and ```Replicas``` can be used to calculate if deployment has reconciled to desired state.

- When dealing with multiple nodetypes each nodeType should have the flexibility to rollout individually.

## TRADEOFFS ON ROLLBACK
- Was the previous state NOT a crashed state ?
- Operator should not change the desired state ie the CR on revert for crash, its core function is to reconcile current to desired state.
