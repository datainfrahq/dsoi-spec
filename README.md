# Distributed Systems Operator Interface
-----------------------------------------------------------------------------------------------
![CI](https://github.com/servicemeshinterface/smi-spec/workflows/CI/badge.svg)

- DSOI Spec is a specification which focuses on defining a standard spec for building CR’s for kubernetes operator’s for distributed systems specifically big data systems such druid, presto, pinot etc.

- The spec defines standard practices and design pattern's which can help build k8s controllers for multiple node types. An example of nodeType can be a druid broker, a presto controller etc.

- DSOI covers most common deployments patterns in order to deploy and maintain heterogenous distributed system clusters.

- The DSOI spec follows standard kubernetes controller practices and adheres to the k8s API specification.

- The DSOI spec is not static and expects to evolve. Not bound to any specific application.

- DSOI spec consists of kubernetes native CRD's and spec's. The goal is not to re-introduce any new concepts but to encapsulate existing spec in a single structured spec.

## Contents

- [Motivation](MOTIVATION.md)
- [Design Principles](PRINCIPLES.md) 
- [Presentation](https://docs.google.com/presentation/d/15BXVpppKGI5S3nR87C6aS6-WlkuqC9MCITNU6ff-xmY/edit?usp=sharing)
 
### Note
- Apache®, Apache Druid, Druid® are either registered trademarks or trademarks of the Apache Software Foundation in the United States and/or other countries.
- Apache®, Apache Pinot, Pinot® are either registered trademarks or trademarks of the Apache Software Foundation in the United States and/or other countries.
- Presto, Presto DB are registered trademarks of The Linux Foundation & Presto Foundation.
