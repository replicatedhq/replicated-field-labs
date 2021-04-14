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

  regular_instances = {
    for name, instance in local.provisioner_pairs :
    name => instance
    if length(instance.public_ips) > 0
  }
  jump_boxes = {
    for instance in google_compute_instance.kots-field-labs :
    replace(instance.name, "jump-", "") => instance
    if length(regexall("jump-.*", instance.name)) > 0
  }
  airgap_instances = {
    for name, instance in local.provisioner_pairs :
    name => {
      instance = instance
      jump_box = lookup(local.jump_boxes, instance.name)
    }
    if length(instance.public_ips) == 0
  }
}

resource "google_compute_instance" "airgapped-instance" {
  for_each     = local.airgap_instances
  name         = each.key
  zone         = "us-central1-b"
  machine_type = each.value.instance.machine_type

  provisioner "file" {
    content     = each.value.instance.provision_sh
    destination = "/tmp/provision.sh"
    connection {
      host         = self.network_interface.0.network_ip
      user         = var.user
      bastion_host = each.value.jump_box.network_interface.0.access_config.0.nat_ip
      bastion_user = var.user
    }
  }
  provisioner "remote-exec" {
    inline = [
      "bash /tmp/provision.sh"
    ]
    connection {
      host         = self.network_interface.0.network_ip
      user         = var.user
      bastion_host = each.value.jump_box.network_interface.0.access_config.0.nat_ip
      bastion_user = var.user
    }
  }
  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1804-lts"
      size  = each.value.instance.boot_disk_gb
    }
  }

  network_interface {
    network = "default"
  }
}

// create an instance
resource "google_compute_instance" "kots-field-labs" {
  for_each     = local.regular_instances
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
    access_config {}
  }
}

output "instance_ips" {
  value = {
    for instance in google_compute_instance.kots-field-labs :
    instance.name => instance.network_interface.0.access_config.0.nat_ip
    if length(instance.network_interface.0.access_config) > 0
  }
}
output "airgap_instances" {
  value = [
  for instance in local.airgap_instances: instance.instance.name
  ]
}
