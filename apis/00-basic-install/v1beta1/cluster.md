# Cluster Installation 

- ```Group: dsoi.ballastdata.com```
- ```Version: v1beta1```
- ```Kind: {{ app }}Cluster```

- ```{{ app}}Cluster``` Resource defines a spec which based on the [principles](../../../PRINCIPLES.md) defined.
- Refer to [terminology](../TERMINOLOGY.md).

## Specification

#### External

- An app can rely on an external database, or any third party app.

```
kind: MyAppCluster
metadata:
  name: bigapp
spec:
  external:
    zookeeper:
        ....
        ....
        ....
    metadataStore:
       ....
       ....
       ....
```

#### Order

- Distributed applications will have multiple nodes. An order defines in which order to rollout the nodes. User can change the order of N number of nodes. ```OrderForCreateUpdate``` is an array. 

```
kind: MyAppCluster
metadata:
  name: bigapp
spec:
  OrderForCreateUpdate:
  - node1
  - node2
  - node3
  ...
  ...

```

#### Node K8s Config Group

- K8s specific configurations as defined by the K8s API. The goal is not reinvent any k8s spec. 

```
kind: MyAppCluster
metadata:
  name: bigapp
spec:
  nodeK8sConfigGroup:
    high-mem:
        storageConfigs:
        - name: storage1
            mountPath: "/druid/data/segment"
            volumeClaimTemplates:
            - metadata:
                name: data-volume
            spec:
                accessModes:
                - ReadWriteOnce
                resources:
                requests:
                    storage: 2Gi
                storageClassName: gp2
        serviceAccountName: "high-mem"
        resources:
          limits:
            cpu: "2"
            memory: 3Gi
          requests:
            cpu: "1"
            memory: 1Gi
```

#### App Config Group

- Application specific configurations. The goal is for app developers to add specific configs without learning the k8s specific configs. 
```
kind: MyAppCluster
metadata:
  name: bigapp
spec:
  appConfigGroup:
    myconfigs:
      scope: common # common to all the nodes
      ....
    node1configs:
      nodeType: broker
      scope: node # default to node scope
      ....
      ....
      ....
    node2configs:
      nodeType: coordinator
      scope: node 
      ....
      ....
      ....
```

#### Map config group to nodes

- After defining the configs, map which config is needed for which ```nodeType```.

```
kind: MyAppCluster
metadata:
  name: bigapp
spec:
  nodes:
    {{ nodeType }}
    - name: az1-{{ nodeType }}
      replicas: 2 # user wants to 2 replicas ( instances ) of brokeraz1 of type broker.
      nodeK8sConfigGroup: "node1ConfigGroup"
      appConfigGroup: "app1ConfigGroup"
    - name: az2-{{ nodeType }}
      replicas: 1
      nodeConfigGroup: "node1ConfigGroup"
      appConfigGroup: "app2ConfigGroup"
```
