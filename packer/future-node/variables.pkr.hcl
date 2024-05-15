variable "project" {
  type    = string
  default = "kots-field-labs"
}

variable "image_name" {
  type    = string
  default = "future-node"
}

variable "zone" {
  type    = string
  default = "europe-west1-b"
}

variable "source_image_family" {
  type    = string
  default = "ubuntu-2204-lts"
}

variable "machine_type" {
  type    = string
  default = "n1-standard-2"
}

variable "disk_size" {
  type    = string
  default = "50"
}

variable "ssh_username" {
  type    = string
  default = "ubuntu"
}

variable "service_account_key" {
  type    = string
}
