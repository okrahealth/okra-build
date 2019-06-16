variable "project" {
  default = "okrahealth"
}

variable "region" {
  default = "us-west1"
}

variable "zone" {
  default = "us-west1-a"
}

variable "workspace_to_environment_map" {
  type = "map"
  default = {
    dev     = "dev"
    prod    = "prod"
  }
}

variable "workspace_to_size_map" {
  type = "map"
  default = {
    dev = "small"
    prod = "large"
  }
}
variable "environment_to_size_map" {
  type = "map"
  default = {
    dev     = "small"
    prod    = "large"
  }
}

variable "machine_instance_type_map" {
  description = "A map from environment to the type of Google Cloud VM instance"
  type = "map"
  default = {
    small  = "g1-small"
    large  = "n1-standard-4"
  }
}
