# Kafka + Go Interview Guide

## Why partitions?

Partitions provide Kafka's parallelism and scalability model.

## What is a consumer group?

A consumer group allows multiple Go consumer instances to cooperatively consume a topic.

## What happens when consumers increase?

Kafka reassigns partitions among group members. Consumers beyond the available partition count cannot actively process a partition simultaneously.

## What is consumer lag?

The gap between records available in Kafka and the consumer group's committed/processed position.

## How do you handle failures?

Use bounded retries for transient failures and a dead-letter topic for records requiring investigation.

## Why Go for Kafka services?

Go provides lightweight compiled services, straightforward concurrency, and a strong fit for cloud-native worker and API workloads.

## What affects Kafka capacity?

Traffic rate, record size, partitions, replication, retention, broker resources, network throughput, and consumer/downstream capacity.

## What should you monitor?

Consumer lag, throughput, errors, processing latency, Go runtime health, broker CPU/memory/disk/network, and JVM health where applicable.

## Production Discussion

Be prepared to explain partition strategy, key selection, ordering, consumer scaling, rebalancing, retention, replication, retry/DLQ design, monitoring, capacity planning, and AWS MSK networking/security.
