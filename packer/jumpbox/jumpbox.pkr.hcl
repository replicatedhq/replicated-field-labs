packer {
  required_plugins {
    googlecompute = {
      version = ">= 0.0.1"
      source = "github.com/hashicorp/googlecompute"
    }
  }
}

source "googlecompute" "airgap-jumpbox" {
  project_id  = var.project
  image_name  = var.image_name
  zone        = var.zone

  source_image_family = var.source_image_family
  machine_type        = var.machine_type
  disk_size           = var.disk_size

  metadata = {
    "user-data" = file("user-data.yaml")
  }
  ssh_username = var.ssh_username 
  account_file = var.service_account_key
}

build {
  sources = ["sources.googlecompute.airgap-jumpbox"]

  provisioner "shell" {
    inline = [
      "cloud-init status --wait",
    ]
  }
}
