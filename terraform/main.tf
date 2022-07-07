variable "gcp_project" {
  description = "gcp project to provision instances to"
}
variable "gcp_zone" {
  description = "gcp zone to provision instances to"
}

provider "google" {
  project = var.gcp_project
  region = join("-", [
    split("-", var.gcp_zone)[0],
    split("-", var.gcp_zone)[1],
    ]
  )
}

variable "user" {
  description = "ssh user for provisioning instances"
}
variable "owner" {
  description = "The value for the owner label on the compute instance (should be your Replicated username)"
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
    replace(instance.name, "-jump", "") => instance
    if length(regexall(".*-jump", instance.name)) > 0
  }
  proxies = {
    for name, instance in local.provisioner_pairs :
    instance.prefix => true...
    if instance.use_proxy
  }
  airgap_instances = {
    for name, instance in local.provisioner_pairs :
    name => {
      instance = instance
      jump_box = lookup(local.jump_boxes, instance.name)
    }
    if length(instance.public_ips) == 0
  }
  names = setunion([
    for name, instance_ip in local.instance_ips :
    split("-", split("lab", name)[0])[1]
  ])
  instance_ips = {
    for instance in google_compute_instance.kots-field-labs :
    instance.name => instance.network_interface.0.access_config.0.nat_ip
    if length(instance.network_interface.0.access_config) > 0
  }
  grouped_by_name = {
    for name in local.names :
    name =>
    {
      ips = [for iname, ip in local.instance_ips :
        "${ip}\tlab${split("lab", iname)[1]}\t# ${iname}"
        if length(regexall(name, iname)) > 0
      ]
      labnames = join(", ", [for iname, ip in local.instance_ips :
        regex("lab(\\d+)", iname)[0]
        if length(regexall(name, iname)) > 0
      ])
    }
  }
}

resource "google_compute_instance" "shared_squid_proxy" {
  for_each     = local.proxies
  name         = "${each.key}-kots-field-labs-squid-proxy"
  zone         = var.gcp_zone
  machine_type = "n1-standard-1"

  labels = {
    user       = var.user
    owner      = var.owner
    expires-on = formatdate("YYYY-MM-DD", timeadd(timestamp(), "336h"))
  }

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
      size  = 10
    }
  }

  provisioner "remote-exec" {
    inline = [
      "sudo apt-get update",
      "sudo apt-get install -y apt-transport-https ca-certificates curl gnupg lsb-release",
      "curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg",
      "echo \"deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null",
      "sudo apt-get update",
      "sudo apt-get install -y docker-ce docker-ce-cli containerd.io",
      "sudo docker run --name squid -d -p 3128:3128 --volume /home/foo/squid/logs:/var/log/squid datadog/squid"
    ]
    connection {
      host = self.network_interface.0.access_config.0.nat_ip
      user = var.user
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }

}

resource "google_compute_instance" "airgapped-instance" {
  for_each     = local.airgap_instances
  name         = each.key
  zone         = var.gcp_zone
  machine_type = each.value.instance.machine_type

  labels = {
    user       = var.user
    owner      = var.owner
    expires-on = formatdate("YYYY-MM-DD", timeadd(timestamp(), "336h"))
  }

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
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
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
  zone         = var.gcp_zone
  machine_type = each.value.machine_type

  labels = {
    user       = var.user
    owner      = var.owner
    expires-on = formatdate("YYYY-MM-DD", timeadd(timestamp(), "336h"))
  }

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
      image = "ubuntu-os-cloud/ubuntu-2004-lts"
      size  = each.value.boot_disk_gb
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }
}

output "instance_ips" {
  value = local.instance_ips
}
output "airgap_instances" {
  value = [
    for instance in local.airgap_instances : instance.instance.name
  ]
}

output "proxies" {
  value = [
    for instance in google_compute_instance.shared_squid_proxy : {
      name    = instance.name
      address = instance.network_interface.0.access_config.0.nat_ip
    }
  ]
}

resource "local_file" "etc_hosts" {
  for_each = local.grouped_by_name


  filename = "${path.module}/etchosts/${each.key}"
  content  = <<EOF
# copy the below and add it to your hosts file with
#
#     echo '
#     <PASTE>
#     ' | sudo tee -a /etc/hosts

${join("\n", each.value.ips)}

EOF

}
resource "local_file" "emails" {
  for_each = local.grouped_by_name


  filename = "${path.module}/emails/${each.key}"
  content  = <<EOF
Hi ${title(each.key)} - your labs have been provisioned and are ready!

Labs ${each.value.labnames} have been provisioned.

1. Check your email and accept the invite to create an account (subject will be "Invitation to join team on replicated" from "contact@replicated.com")
2. Navigate to the first lab and start working from the readme: https://github.com/replicatedhq/kots-field-labs/tree/main/labs/lab00-hello-world
3. You will see prompts to "insert your IP address here" -- those IPs for each participant are found below. The password for each SSH login will be "password" (or try "replicated" if that doesn't work) and the UI password for logging into the browser view will be "password" (or try "replicated" if that doesn't work).

IPs for your instances are below — you can use them raw, or drop the snippet into /etc/hosts

${join("\n", each.value.ips)}

If you get stuck, feel free to reach out. If all goes well, we could provision these same labs for folks on your team if you think it would help.

Best,
${title(var.user)}
EOF

}
