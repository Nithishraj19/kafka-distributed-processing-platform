# AWS MSK Reference Architecture

A production AWS implementation can use Amazon MSK as the managed Kafka layer.

Reference components:

- VPC
- Private subnets
- Amazon MSK
- IAM and security controls
- Go producer services
- Go consumer worker services
- Kubernetes workloads
- CloudWatch/Prometheus/Grafana monitoring
- Redis/PostgreSQL/ClickHouse downstream systems

This document is an architecture reference. It does not claim that an AWS MSK cluster has been deployed by this repository.
