## Auto Scaling Spec

- ```Group: dsoi.ballastdata.com```
- ```Version: v1beta1```
- ```Kind: {{ app }}scaling```

- Autoscaling of any cluster node can be categorized into

1. Horizontal Scaling
2. Vertically Scaling

- Each node has unique set of scaling requirements when it comes to scaling

1. Compute
2. Memory
3. Storage

- When Scaling vertically, each app depends on certain metrics. In order to get those metrics
  operator can make external API calls. 

- The operator shall consist of a seperate CR controller. Adding auto scaling logic in a single controller which performs installations and upgrade can result in delays and race conditions as fetching metrics and computing a decision is costly for each nodeType with N nodes every N seconds.

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
