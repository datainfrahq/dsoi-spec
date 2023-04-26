## 📦 🚀  Cluster Installation

### Defining CRD

To define your CRD, follow these steps:

-  Define the Group, Version, and Kind for your CRD.
   - Group: `dsoi.datainfra.io`
   - Version: `v1beta1`
   - Kind: `{{ app }}`

-  Example:
```
apiVersion: datainfra.io/v1beta1
kind: Pinot
metadata:
  name: pinot
```

## 📄 Specification

#### External

- An app can rely on an external database, or any third party dependency.

```
kind: MyApp
metadata:
  name: myapp
spec:
  external:
    zookeeper:
     spec:
        ....
        ....
        ....
    metadataStore:
      spec:
       ....
       ....
       ....
```

- Example:
```
apiVersion: datainfra.io/v1beta1
kind: Pinot
metadata:
  name: pinot-basic
spec:
  external:
    zookeeper:
     spec:
       zkAddress: zk-pinot-zookeeper-headless.pinot:2181
```


#### Deployment Order

- Distributed applications often have multiple nodes, and the order in which these nodes are rolled out can be important. 
- This can be achieved using the deploymentOrder array.

```
kind: MyApp
metadata:
  name: bigapp
spec:
  deploymentOrder:
  - node1
  - node2
  - node3

```
- Example

```
apiVersion: datainfra.io/v1beta1
kind: Pinot
metadata:
  name: pinot-basic
spec:
  deploymentOrder:
    - controller
    - broker
    - server
    - minion
```
#### K8s Config Group

- K8s specific configurations as defined by the K8s API. The goal is not reinvent any k8s spec. 

- Example
```
k8sConfig:

  - name: controller
    serviceAccountName: "default"
    port:
    - name: controller 
      containerPort: 9000
      protocol: TCP
    service:
      type: LoadBalancer
      ports:
      - protocol: TCP
        port: 9000
        targetPort: 9000
    livenessProbe:
      initialDelaySeconds: 60
      periodSeconds: 10
      httpGet:
        path: "/health"
        port: 9000
    readinessProbe:
      initialDelaySeconds: 60
      periodSeconds: 10
      httpGet:
        path: "/health"
        port: 9000
    env:
    - name: LOG4J_CONSOLE_LEVEL
      value: info
    image: apachepinot/pinot:latest
    storageConfig:
    - name: pinotcontroller
      mountPath: "/var/pinot/controller/data"
      spec:
        accessModes:
        - ReadWriteOnce
        storageClassName: ${STORAGE_CLASS_NAME}
        resources:
          requests:
            storage: 1Gi
```

#### App NodeConfig Group

- Application specific configurations. The goal is for app developers to add specific configs without learning the k8s specific configs. 

```
kind: MyAppCluster
metadata:
  name: bigapp
spec:
  appConfigGroup:
  - name: nodename
    configs-one:
    ....
    configs-two:
    ....
    configs-three:
    ....

```

- Example:

```
pinotNodeConfig:
  - name: controller
    java_opts: "-XX:ActiveProcessorCount=2 -Xms256M -Xmx1G -XX:+UseG1GC -XX:MaxGCPauseMillis=200
                -Xlog:gc*:file=/opt/pinot/gc-pinot-controller.log -Dlog4j2.configurationFile=/opt/pinot/conf/log4j2.xml
                -Dplugins.dir=/opt/pinot/plugins"
    data: |-
        controller.port=9000
        controller.data.dir=/var/pinot/controller/data 
        pinot.set.instance.id.to.hostname=true
        controller.task.scheduler.enabled=true
```

#### Map K8s Config Group and App Config group to Each NodeType

- After defining the configs, map which config is needed for which ```nodeType```.
- Example

```
apiVersion: datainfra.io/v1beta1
kind: Pinot
metadata:
  name: pinot-basic
spec:

  nodes:

    - name: pinot-controller
      kind: Statefulset
      replicas: 1
      nodeType: controller
      k8sConfig: controller
      pinotNodeConfig: controller
    
    - name: pinot-broker
      kind: Statefulset
      replicas: 1
      nodeType: broker
      k8sConfig: broker
      pinotNodeConfig: broker

    - name: pinot-server
      kind: Statefulset
      replicas: 1
      nodeType: server
      k8sConfig: server
      pinotNodeConfig: server
    
    - name: pinot-minion
      kind: Statefulset
      replicas: 1
      nodeType: minion
      k8sConfig: minion
      pinotNodeConfig: minion
```
