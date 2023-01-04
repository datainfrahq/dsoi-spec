# Distributed Systems Operator Interface
-----------------------------------------------------------------------------------------------

> DSOI is a functional spec for responses to objectives. An objective is a user defined action.
> An action for the spec is converging desired state to current state for any distributed application. 
>The DSOI spec follows standard kubernetes controller practices and adheres to the k8s API specification. 

- The Distributed system operator interface (DSOI) is a specification which focuses on building kubernetes operator's for distributed system specifically big data systems like apache druid, apache pinot, presto. 

- The spec defines standard practices and design pattern's which can help build k8s controllers for multiple node types. An example of nodeType can be a druid broker, a presto controller etc.

- DSOI covers most common deployments patterns in order to deploy and maintain heterogenous kubernetes clusters. 

## Contents

- [Design Principles](PRINCIPLES.md) 

