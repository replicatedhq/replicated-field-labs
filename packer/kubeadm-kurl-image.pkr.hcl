packer {
  required_plugins {
    googlecompute = {
      version = ">= 0.0.1"
      source = "github.com/hashicorp/googlecompute"
    }
  }
}

source "googlecompute" "kubeadm-kurl-cluster" {
  project_id  = var.project
  image_name  = var.image_name
  zone        = var.zone

  source_image_family = var.source_image_family
  machine_type        = var.machine_type
  disk_size           = var.disk_size

  ssh_username = var.ssh_username 
  account_file = var.service_account_key
}

build {
  sources = ["sources.googlecompute.kubeadm-kurl-cluster"]

  provisioner "shell" {
    inline = [
      "curl https://kurl.sh/latest | sudo bash"
    ]
  }

  provisioner "shell" {
    inline = [
      "sudo groupadd --gid 1020 replicant",
      "sudo useradd --uid 1020 --gid 1020 --shell /bin/bash --create-home replicant"
    ]
  }

  provisioner "shell" {
    inline = [
      "sudo mkdir -p /home/replicant/.kube",
      "sudo cp /etc/kubernetes/admin.conf /home/replicant/.kube/config",
      "sudo chmod 0700 /home/replicant/.kube && sudo chmod 0400 /home/replicant/.kube/config",
      "sudo chown -R replicant /home/replicant/.kube",
      "echo export KUBECONFIG=/home/replicant/.kube/config | sudo tee /home/replicant/.profile"
    ]
  }

}
