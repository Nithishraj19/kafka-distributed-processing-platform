# Kafka Partitions

Partitions are Kafka's unit of parallelism.

Records with the same key can be routed consistently to a partition, helping preserve ordering for that key.

Partition count should be selected based on throughput, consumer parallelism, storage, and operational requirements.
