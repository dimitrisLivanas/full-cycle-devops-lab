# DevOps Learning Journey - Full Cycle Project

## Progress Report: Step 1 Completion

**Objective:** The initial task was to build a foundational, standalone HTTP server. This serves as the core application that will be used throughout the subsequent DevOps exercises involving containerization, deployment, and monitoring.

**Technology Selection:**
* The **Go (Golang)** programming language was selected for this task.
* **Reasoning:** Go was chosen strategically due to its strong prevalence in the DevOps ecosystem, its ability to compile to a single, dependency-free binary, and its high performance—all of which are highly desirable for creating efficient, portable microservices.

**Key Concepts Learned & Implemented:**

1.  **HTTP Server Fundamentals:** Learned what an HTTP server is at its core: a program that listens on a specific network port for incoming requests and sends back responses.
2.  **Go `net/http` Standard Library:** Successfully built the server using Go's powerful `net/http` package, demonstrating the ability to create a functional web service without requiring any external frameworks or dependencies.
3.  **Request Handling & Routing:** Implemented a basic request handler using `http.HandleFunc` to route incoming requests to specific Go functions.
4.  **Client-Server Interaction:** Tested the running server using two different types of HTTP clients: a graphical web browser (Chrome) and a command-line tool (`curl`).
5.  **Analyzing Client Behavior:** By logging request details, I observed that browsers are "chatty" clients that automatically request additional resources (`/favicon.ico`), unlike simple clients like `curl`.

**Outcome:**
A simple but robust "Hello, World" HTTP server has been successfully built and tested locally. The project is now stored on GitHub and ready for the next stage: **containerization with Docker**.
