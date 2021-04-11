provider "google" {
  project = "smart-proxy-839"
  region  = "us-central1"
}

variable "user" {
  description = "ssh user for provisioning instances"
}
variable "provisioner_pairs_json" {
  description = "path to json file containing instance details"
}

locals {
  // load the json file produced by the go program, with a list of pairs of instance-name + provision-script,
  // one instance per participant per exercise
  provisioner_pairs = jsondecode(file("${path.module}/../${var.provisioner_pairs_json}"))
}

// create an instance
resource "google_compute_instance" "kots-field-labs" {
  for_each     = local.provisioner_pairs
  name         = each.key
  zone         = "us-central1-b"
  machine_type = each.value.machine_type

  provisioner "file" {
    content     = each.value.provision_sh
    destination = "/tmp/provision.sh"
    connection {
      host = self.network_interface.0.access_config.0.nat_ip
      user = var.user
    }
  }
  provisioner "remote-exec" {
    inline = [
      "bash /tmp/provision.sh"
    ]
    connection {
      host = self.network_interface.0.access_config.0.nat_ip
      user = var.user
    }
  }
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1804-lts"
      size  = each.value.boot_disk_gb
    }
  }

  network_interface {
    network = "default"
    // skip creating public ip if the map is empty. Wee bit of a hack
    dynamic access_config {
      for_each = each.value.public_ips
      iterator = ignored

      content {}
    }

  }
}

output "instance_ips" {
  value = {
    for instance in google_compute_instance.kots-field-labs:
    instance.name => instance.network_interface.0.access_config.0.nat_ip
    if length(instance.network_interface.0.access_config) > 0
  }
}