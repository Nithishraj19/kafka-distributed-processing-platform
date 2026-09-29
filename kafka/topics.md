# Kafka Topics

| Topic | Purpose |
|---|---|
| events | General platform events |
| jobs | Work items for processing |
| notifications | Notification events |
| events.retry | Retryable records |
| events.dlq | Records requiring investigation |

Topic design should consider throughput, ordering, retention, partition count, consumer behavior, and downstream capacity.
