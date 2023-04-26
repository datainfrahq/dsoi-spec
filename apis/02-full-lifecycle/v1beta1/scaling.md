# Auto Scaling Spec

Group: dsoi.datainfra.io

Version: v1beta1

Kind: {{ app }}scaling

Autoscaling of any cluster node can be categorized into:

- Horizontal Scaling
- Vertically Scaling

Each node has a unique set of scaling requirements when it comes to scaling:

- Compute
- Memory
- Storage

When scaling vertically, each app depends on certain metrics. In order to get those metrics, the operator can make external API calls.

The operator shall consist of a separate CR controller. Adding auto scaling logic in a single controller which performs installations and upgrades can result in delays and race conditions as fetching metrics and computing a decision is costly for each nodeType with N nodes every N seconds.


```
spec:
    external:
        spec:
            url: http://getmymmetrics.com/scalenow
    autoscaling:
      - name: {{ name of node }}
        spec:
            min:
            max:
            threshold:
            coolDown:
      - name: {{ name of node }}
        spec:
            min:
            max:
            threshold:
            coolDown:
        .....
```
