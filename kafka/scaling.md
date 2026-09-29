# Consumer Scaling

Consumer scaling is driven by partition count.

Example:

```text
12 Kafka partitions
        |
        +-- Go Worker 1
        +-- Go Worker 2
        +-- Go Worker 3
        +-- Go Worker 4
```

Adding consumers can increase parallelism until all available partitions are assigned.

Operational scaling should also consider processing latency, downstream capacity, consumer lag, CPU, memory, network throughput, and rebalancing.
