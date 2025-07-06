# DevOps Learning Journey - Full Cycle Project

This repository documents my hands-on learning journey through a series of DevOps tasks, from building a simple web server to deploying and monitoring it in a cloud environment.

---
## Step 1: Building a Go HTTP Server

**Objective:** Create a foundational, standalone HTTP server.
* **Technology:** Go (Golang)
* **Key Concepts:** `net/http` standard library, request handlers, and testing with `curl`.

---
## Step 2: Containerization with Docker

**Objective:** Package the Go application into a portable Docker container.
* **Technology:** Docker
* **Key Concepts:** `Dockerfile` basics (`FROM`, `RUN`, `CMD`), `.dockerignore`.

---
## Step 3: Orchestration with Docker Compose

**Objective:** Simplify the management of the application service for local development.
* **Technology:** Docker Compose
* **Key Concepts:** `docker-compose.yml`, defining services, `docker compose up/down`.

---
## Step 4: Optimizing with Multi-Stage Docker Builds

**Objective:** Drastically reduce the final image size for better security and efficiency.
* **Technology:** Docker
* **Key Concepts:** Multi-stage builds (`FROM ... AS builder`), copying artifacts between stages.

---
## Step 5: Intro to IaC with Terraform

**Objective:** Learn the fundamentals of Infrastructure as Code (IaC) by managing a simple cloud resource.
* **Technology:** Terraform
* **Key Concepts:** IaC principles, declarative syntax (HCL), the `init`/`plan`/`apply`/`destroy` workflow.

---
## Step 6: Provisioning a Kubernetes Cluster

**Objective:** Provision a production-grade Kubernetes cluster in the cloud using Terraform.
* **Technologies:** Terraform, Google Kubernetes Engine (GKE)
* **Key Concepts:** GKE Autopilot, configuring `kubectl`, managing cloud APIs as code.

---
## Step 7: Deploying an App to Kubernetes

**Objective:** Deploy a containerized application to the GKE cluster.
* **Technology:** Kubernetes
* **Key Concepts:** Kubernetes manifests (YAML), `Deployment` and `Service` objects, `type: LoadBalancer`.

---
## Step 8: Infrastructure Cleanup

**Objective:** Learn to properly tear down cloud infrastructure to manage costs.
* **Technology:** Terraform
* **Key Concepts:** `terraform destroy`, `deletion_protection` as a safety feature.

---
## Step 9: Deploying a Local Monitoring Stack

**Objective:** Run a local observability backend to receive telemetry data.
* **Technologies:** Docker Compose, Jaeger
* **Key Concepts:** The role of an observability backend, using Docker Compose for local infrastructure, debugging container startup issues (`ps`, `logs`), and resolving Docker port permissions.

---
## Step 10: Instrumenting an Application with OpenTelemetry

**Objective:** Modify the Go application to generate and export traces to the monitoring backend.
* **Technology:** Go, OpenTelemetry
* **Key Concepts:** Application **instrumentation**, using OpenTelemetry SDKs, configuring an **exporter** (OTLP), and understanding Docker container networking (`localhost` vs. service name).

---
### **Next Steps:**

Deploy the fully instrumented application to Kubernetes and integrate it with Google Cloud's native monitoring services.