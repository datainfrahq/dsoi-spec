# Distributed Systems Operator Interface
-----------------------------------------------------------------------------------------------

> DSOI is a functional spec for responses to objectives. An objective is a user defined action.
> An action for the spec is converging desired state to current state for any distributed application. 
>The DSOI spec follows standard kubernetes controller practices and adheres to the k8s API specification. 

- The Distributed system operator interface (DSOI) is a specification which focuses on building kubernetes operator's for distributed system specifically big data systems like apache druid, apache pinot, presto. 

- The spec defines standard practices and design pattern's which can help build k8s controllers for multiple node types. An example of nodeType can be a druid broker, a presto controller etc.

- DSOI covers most common deployments patterns in order to deploy and maintain heterogenous kubernetes clusters.

## Principles

- Seperation of Concerns

> A clean consise seperation between 
> 1. An application, K8s Infra and cloud components.
> 2. Storage, Network and Compute
> 3. Configuration Management and Reconcilation of configuration

- Building Abstractions

> An Abstraction is a purposeful suppression, incapsulating each concern.
> 1. Control Plane and Data Plane
> 2. One CRD per controller
> 3. Hetrogenous configurations

- Design Interface

> An interface are methods for communication between incapsulated abstractions.
> 1. K8s Client and Application
> 2. External components ie metadata, deepstroage, zookeeper etc.
> 3. Metrics collection

- Define Contracts
