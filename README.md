<h2 align="center">
  <picture>
    <img alt="DataInfra Logo" src="https://raw.githubusercontent.com/datainfrahq/.github/main/images/logo.svg" width="500" height="100">
  </picture>
  <br>
  Distributed Systems Operator Interface
  </br>
</h2>


<div align="center">

![CI](https://github.com/datainfrahq/dsoi-spec/workflows/CI/badge.svg) [![Slack](https://img.shields.io/badge/slack-brightgreen.svg?logo=slack&label=Community&style=flat&color=%2373DC8C&)](https://launchpass.com/datainfra-workspace) 

</div>

The Distributed System Operator Interface (DSOI) specification is a set of best practices for building Kubernetes operators for distributed systems. The spec defines standard practices that can help define custom resources. It consists of Kubernetes-native CRDs and specs and is not bound to any specific application.


Operators built using the DSOI Spec
- [Parseable Kubernetes Operator](https://github.com/parseablehq/operator)
- [Control Plane For Apache Pinot](https://github.com/datainfrahq/pinot-operator)

### :earth_americas: Background

The most common analytical data pipeline consists of a number of different components, including streaming tools like Kafka and Pulsar, stream processing tools like Flink, OLAP stores for ingestion and querying like Druid, query engines like Presto and Trino, and BI tools like Superset. When running any of these tools, databases, or engines, there are a number of common challenges and requirements that must be addressed, including external dependencies like Zookeeper and metadata stores, multiple node types like brokers and controllers, app-specific configuration like runtime properties and log4j, and Kubernetes-specific configuration like pod template specs.

## :dart: Goals

The goal of the DSOI spec is to define a standard that can be implemented by operators when constructing a CR. By doing so, it is expected to lower the learning curve of understanding CRs for various operators and provide a consistent user experience across different operators.

## :bulb: Purpose

The purpose of the DSOI spec is to provide a structured specification for building Kubernetes operators that can be used with distributed systems. By adhering to this specification, developers can create consistent and reliable operators that are easy to maintain and evolve over time.

## :mag_right: Scope

The DSOI spec is designed to be flexible and extensible. It is not meant to introduce new concepts but rather to encapsulate existing best practices into a single, structured spec. It is expected to evolve over time based on feedback and community contributions.

## :heavy_plus_sign: Benefits

Adhering to the DSOI spec provides a number of benefits for developers and operators of distributed systems, including:
- **Standardization**: The spec provides a standardized approach for managing configuration between applications, node types, and Kubernetes configurations.
- **Consistency**: By adhering to the spec, developers can create consistent and reliable operators that are easy to maintain and evolve over time.
- **Interoperability**: The spec promotes interoperability across different operators and can be adapted to a wide range of use cases.

## 🚩 Principles

The Distributed System Operator Interface (DSOI) is designed to adhere to a set of guiding principles that promote best practices in building Kubernetes operators for distributed systems. These principles include:

- **Separation of Concerns**: Each section of the system should address a distinct concern, such as application, Kubernetes infrastructure, cloud components, storage, network, compute, configuration management, and reconciliation of configuration.

- **Building Abstractions**: Encapsulate each concern in a purposeful suppression, such as control plane and data plane, handling business logic with a separate component (controller), heterogeneous configurations, and designing interfaces.

- **Define Contracts**: Reach agreement between the various components responsible for reconciling desired, original, and current state. The DSOI spec defines a contract between the operator and its resources.

- **One Controller per CRD**: Each CRD should have its own controller for installation, scaling, and application specifications. The controller and application should interact with each other.

- **Eventual Consistency**: Controllers should be eventually consistent, built on observed state rather than state machines.

By following these principles, DSOI ensures a **consistent**, **scalable**, and **maintainable** architecture for building Kubernetes operators for distributed systems.
 
## :notebook_with_decorative_cover: NOTE

- Apache®, Apache Druid, Druid® are either registered trademarks or trademarks of the Apache Software Foundation in the United States and/or other countries.
- Apache®, Apache Pinot, Pinot® are either registered trademarks or trademarks of the Apache Software Foundation in the United States and/or other countries.
- Presto, Presto DB are registered trademarks of The Linux Foundation & Presto Foundation.
