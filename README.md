# DevOps Learning Journey - Full Cycle Project

This repository documents my hands-on learning journey through a series of DevOps tasks, from building a simple web server to deploying and monitoring it in a cloud environment.

---
## Step 1: Building a Go HTTP Server

**Objective:** Create a foundational, standalone HTTP server to serve as the core application for this project.

* **Technology:** Go (Golang)
* **Key Concepts:**
    * **HTTP Server Fundamentals:** Understood the role of an HTTP server that listens for requests and sends responses.
    * **Go `net/http`:** Used Go's standard library to build a web server without external dependencies.
    * **Client-Server Interaction:** Tested the server using both a web browser and `curl`.

---
## Step 2: Containerization with Docker

**Objective:** Package the Go application and its dependencies into a portable Docker container.

* **Technology:** Docker
* **Key Concepts:**
    * **Containerization:** Learned the core concept of creating a lightweight, executable package containing everything needed to run the application, solving the "it works on my machine" problem.
    * **Dockerfile:** Authored an initial `Dockerfile` using instructions like `FROM`, `WORKDIR`, `COPY`, `RUN`, and `CMD`.

---
## Step 3: Orchestration with Docker Compose

**Objective:** Simplify the management and definition of the application service for local development.

* **Technology:** Docker Compose
* **Key Concepts:**
    * **Declarative Configuration:** Used a `docker-compose.yml` file to define the application's services, builds, and port mappings.
    * **Service Management:** Replaced long `docker run` commands with simple `docker compose up` and `docker compose down` commands.

---
## Step 4: Optimizing with Multi-Stage Docker Builds

**Objective:** Drastically reduce the final image size for better security and efficiency.

* **Technology:** Docker
* **Key Concepts:**
    * **Multi-Stage Builds:** Used multiple `FROM` statements to separate the build environment (with the Go compiler) from the final runtime environment.
    * **Artifact Copying:** Copied only the compiled binary from the build stage to the final stage using `COPY --from=builder`.
    * **Layer Caching:** Optimized the Dockerfile order (`COPY go.mod` before `COPY . .`) to leverage Docker's layer cache for faster rebuilds.

---
## Step 5: Intro to IaC with Terraform

**Objective:** Learn the fundamentals of Infrastructure as Code (IaC) by managing a simple cloud resource.

* **Technology:** Terraform
* **Key Concepts:**
    * **IaC Principles:** Managed infrastructure with code for automation, repeatability, and version control.
    * **Declarative Syntax:** Used HCL to define the desired state of a Google Cloud Storage bucket.
    * **Core Workflow:** Learned the `terraform init`, `plan`, `apply`, and `destroy` workflow.
    * **State Management:** Understood that Terraform creates a state file to keep track of managed resources and that it should be ignored in version control for local projects.

---
## Step 6: Provisioning a Kubernetes Cluster

**Objective:** Provision a production-grade Kubernetes cluster in the cloud using Terraform.

* **Technologies:** Terraform, Google Kubernetes Engine (GKE)
* **Key Concepts:**
    * **GKE Autopilot:** Provisioned a serverless Kubernetes cluster where Google manages the underlying nodes.
    * **`kubectl`:** Configured the standard Kubernetes CLI to connect to the new GKE cluster.
    * **Cloud API Management:** Learned to enable cloud service APIs as part of the Terraform code using the `google_project_service` resource.

---
### **Next Steps:**

Deploy the containerized application to the new Kubernetes cluster.