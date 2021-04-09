provider "google" {
  project = "smart-proxy-839"
  region  = "us-central1"
}

variable "user" {
  default     = "dex"
  description = "ssh user for provisioning instances"
}

locals {
  // load the json file produced by the go program, with a list of pairs of instance-name + provision-script,
  // one instance per participant per exercise
  provisioner_pairs = jsondecode(file("${path.module}/provisioner_pairs.json"))
}

//// shared postgres instance for labs 2, 3, and 4
//resource "google_sql_database_instance" "shared_postgres" {
//  name             = "shared_postgres"
//  database_version = "POSTGRES_11"
//  region           = "us-central1"
//
//  settings {
//    # Second-generation instance tiers are based on the machine
//    # type. See argument reference below.
//    tier = "db-f1-micro"
//    ip_configuration {
//
//      dynamic "authorized_networks" {
//        for_each = {
//        for instance in google_compute_instance.kots-field-labs:
//        instance.name => instance
//        if length(instance.network_interface) > 0
//        }
//        iterator = instance
//
//        content {
//          name  = instance.value.name
//          value = instance.value.network_interface.0.access_config.0.nat_ip
//        }
//      }
//    }
//  }
//}
//
//
//resource "google_sql_user" "pg_user" {
//  instance = google_sql_database_instance.shared_postgres.name
//  name     = "pg_user"
//  password = "pg_password"
//}

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