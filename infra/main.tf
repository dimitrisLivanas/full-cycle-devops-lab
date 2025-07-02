# This block tells Terraform we want to manage resources in Google Cloud.
provider "google" {
  project = "sturdy-layout-459008-n3"
  region  = "europe-west1"
}

# This block defines a resource we want to create.
# In this case, a Google Cloud Storage bucket.
resource "google_storage_bucket" "my_learning_bucket" {
  # The name must be globally unique across all of Google Cloud.
  name          = "full-cycle-devops-lab-bucket-dimitris" # <-- REPLACE THIS with a unique name
  location      = "EU"      # The location for the bucket's data.
  force_destroy = true    # Allows Terraform to delete the bucket even if it has files in it. Good for labs.
}