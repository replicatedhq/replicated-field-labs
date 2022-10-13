variable "project" {
  type    = string
  default = "kots-field-labs"
}

variable "image_name" {
  type    = string
  default = "k3s-kurl-latest"
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
  default = "n1-standard-1"
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
