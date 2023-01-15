# UPGRADE

- ```Group: dsoi.ballastdata.com```
- ```Version: v1beta1```
- ```Kind: {{ app }}Cluster```

- ```{{ app}}Cluster``` Resource defines a spec which based on the [principles](../../../PRINCIPLES.md) defined.

- When dealing with multiple node, in any ideal scenario, cluster installation should be parallel. The first installation can be easily calculated using ```spec.generation```.

- Upgrades should be incremental across different nodeTypes. Order defined in the CR should be easily configurable.

- Emit event for each action on an upgrade.

- To understand Order, refer to [terminology](../TERMINOLOGY.md).

- Incase of a crash during upgrades, ideal scenarios are

1. Wait for manual intervention, get alerted.
2. Rollback after a timeout defined in the operator spec.

## TRADEOFFS ON ROLLBACK

- Was the previous state NOT a crashed state ?
- Operator should not change the desired state ie the CR on revert for crash, its core function is to reconcile current to desired state.
