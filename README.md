# DevOps Learning Journey - Full Cycle Project

This repository documents my hands-on learning journey through a series of DevOps tasks, from building a simple web server to deploying and monitoring it in a cloud environment.

---
## Step 1: Building a Go HTTP Server

**Objective:** Create a foundational, standalone HTTP server to serve as the core application for this project.

* **Technology:** Go (Golang)
* **Key Concepts:**
    * **HTTP Server Fundamentals:** Understood the role of an HTTP server that listens for requests and sends responses.
    * **Go `net/http`:** Used Go's standard library to build a web server without external dependencies.
    * **Client-Server Interaction:** Tested the server using both a web browser and `curl`, observing the different types of requests they generate (e.g., for `/favicon.ico`).

---
## Step 2: Containerization with Docker

**Objective:** Package the Go application and its dependencies into a portable Docker container, solving the "it works on my machine" problem.

* **Technology:** Docker
* **Key Concepts:**
    * **Containerization:** Learned the core concept of creating a lightweight, executable package containing everything needed to run the application.
    * **Dockerfile:** Authored a `Dockerfile` using instructions like `FROM`, `WORKDIR`, `COPY`, `RUN`, and `CMD` to create a repeatable and isolated build process.
    * **Go Modules:** Understood the necessity of `go.mod` for modern Go projects, which was required for the Docker build to succeed.
    * **Build Optimization:** Added a `.dockerignore` file to keep the Docker build context small and efficient by excluding unnecessary files.

---
## Step 3: Orchestration with Docker Compose

**Objective:** Simplify the management and definition of the application service for local development.

* **Technology:** Docker Compose
* **Key Concepts:**
    * **Declarative Configuration:** Used a `docker-compose.yml` file to declaratively define the application's services, builds, and port mappings.
    * **Service Management:** Replaced long `docker run` commands with simple `docker compose up` and `docker compose down` commands to easily manage the application's lifecycle.
    * **Modern CLI:** Adapted to the new `docker compose` (with a space) syntax, which is integrated directly into the Docker CLI.

---
### **Next Steps:**

The next stage is to improve the Docker build process with a multi-stage build before moving on to infrastructure and cloud deployment.