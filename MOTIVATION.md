## MOTIVATION

- The most common analytical data pipeline consits of 
1. Streaming -  kafka, pulsar
2. Stream processing - flink
3. OLAP stores - ingestion and query - druid
4. Query engines - presto, trino
4. BI Tools - superset

- When running any of the above tools/databases/engines we are dealing with 

1. External Dependencies - Zookeeper, Metadatastore etc.
2. Multiple NodeTypes - Broker, Controller 
3. App Specific Configuration - runtime properties, common runtime, log4j.
4. K8s Specific Configuration - pod template spec.
5. Ordered Upgrades/Installation - Installing/Deploying in a specific order.
6. NodeType specific scaling/downscaling - Each nodeType is unique.

- The above is common for the most of distributed systems in the big data domain.

- The goal with the spec, is to define a standard which can implemented by operator's while constructing a CR.

- The user experience while using/implementing any operator will lower down the learning curve of understanding CR for various operator's.

- The Spec clearly defines abstractions between configuration management between an application, its nodeTypes and k8s configurations.
