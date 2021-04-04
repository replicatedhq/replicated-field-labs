provider "google" {
  project = "smart-proxy-839"
  region = "us-central1"
}

variable "user" {
  default = "dex"
  description = "ssh user for provisioning instances"
}

locals {
  // load the json file produced by the go program, with a list of pairs of instance-name + provision-script,
  // one instance per participant per exercise
  provisioner_pairs = jsondecode(file("${path.module}/provisioner_pairs.json"))
}

// create an instance
resource "google_compute_instance" "kots-field-labs" {
  for_each = local.provisioner_pairs
  name = each.key
  zone = "us-central1-b"

  provisioner "file" {
    content = each.value
    destination = "/tmp/provision.sh"
    connection {
      host        = self.network_interface[0].access_config[0].nat_ip
      user        = var.user
    }
  }
  provisioner "remote-exec" {
    inline = [
      "bash /tmp/provision.sh"
    ]
    connection {
      host        = self.network_interface[0].access_config[0].nat_ip
      user        = var.user
    }
  }
  // might be able to get away with n1-standard-2, most of these are just running nginx
  machine_type = "n1-standard-4"
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1804-lts"
      size = "200"
    }
  }
  network_interface {
    network = "default"
    access_config {}
  }
}

output "instance_ips" {
  value = {
  for instance in google_compute_instance.kots-field-labs:
  instance.name => instance.network_interface[0].access_config[0].nat_ip
  }
}