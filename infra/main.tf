# Configure the Google Cloud provider
provider "google" {
  project = "sturdy-layout-459008-n3" # <-- REPLACE THIS if you haven't already
  region  = "europe-west8"               # Milan, Italy
}

# This resource enables the Kubernetes Engine API for our project.
resource "google_project_service" "kubernetes_api" {
  project = "sturdy-layout-459008-n3" # <-- Make sure this is correct
  service = "container.googleapis.com"

  # This prevents Terraform from trying to disable the API when you run 'destroy'.
  disable_on_destroy = false
}

# Define a GKE Autopilot cluster resource
resource "google_container_cluster" "primary" {
  # This tells Terraform to wait until the API is enabled before trying to create the cluster.
  depends_on = [google_project_service.kubernetes_api]

  name              = "full-cycle-devops-lab-cluster"
  location          = "europe-west8"
  enable_autopilot  = true
  deletion_protection = false
}