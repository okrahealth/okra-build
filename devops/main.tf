######################################
# Terraform config
######################################

provider "google" {
  # credentials file below will change; it is experimental to test it works on
  # a sample google account
  credentials = "${file("okrahealth-0a94406ba78b.json")}"
  project = "${var.project}"
  region  = "${var.region}"
  zone = "${var.zone}"
  version = "~> 2.0"
}

provider "google-beta" {
  project = "${var.project}"
  region  = "${var.region}"
  version = "~> 2.0"
}

provider "template" {
  version = "~> 1.0"
}

locals {
  environment = "${lookup(var.workspace_to_environment_map, terraform.workspace, "dev")}"
  size = "${local.environment == "dev" ? lookup(var.workspace_to_size_map, terraform.workspace, "small") : var.environment_to_size_map[local.environment]}"
}

// A single Google Cloud Engine instance
resource "google_compute_instance" "default" {
  name         = "okra-build"
  machine_type = "${var.machine_instance_type_map[local.size]}"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }

  network_interface {
    network = "default"

    access_config {
      // Include this section to give the VM an external ip address
    }
  }
}

output "ip" {
  value = "${google_compute_instance.default.network_interface.0.access_config.0.nat_ip}"
}

output "machine_instance_type" {
  value = "${var.machine_instance_type_map[local.size]}"
}
