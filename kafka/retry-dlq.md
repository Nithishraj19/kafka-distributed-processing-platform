# Retry and Dead-Letter Processing

```text
Main Topic
    |
    v
Go Consumer
    |
    +---- success ----> processing complete
    |
    +---- transient --> retry topic
    |
    +---- invalid ----> dead-letter topic
```

Use bounded retry attempts for transient failures. Records that cannot be processed safely can be routed to a dead-letter topic for investigation or replay.
