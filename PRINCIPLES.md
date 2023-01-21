## Principles

- Seperation of Concerns

> A seperation is a distinct section. Each section address a concern.
> 1. An application, K8s Infra and cloud components
> 2. Storage, Network and Compute
> 3. Configuration Management and Reconcilation of configuration

- Building Abstractions

> An Abstraction is a purposeful suppression, incapsulating each concern.
> 1. Control Plane and Data Plane
> 2. Handling each buisness logic with a seperate component ( controller )
> 3. Hetrogenous configurations

- Design Interface

> An interface are methods for communication between incapsulated abstractions.
> 1. K8s Client and Application
> 2. External components ie metadata, deepstroage, zookeeper etc
> 3. Metrics collection internal or external required for building logic

- Define Contracts

> A contract is an agreement, incase of DSOI spec, it is between the various components responsible for reconciling desired, original and current state.
> 1. One controller per CRD for installation, scaling and application specifications.
> 2. Interaction between controller and application.
> 3. Controllers are eventually consistent, build on observed state. NOT state machines.
