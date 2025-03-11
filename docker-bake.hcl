target "docker-metadata-action" {}

variable "firecracker" { default = "" }
target "firecracker" {
  inherits   = ["docker-metadata-action"]
  context    = "./library/firecracker"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  args = {
    VERSION = "${firecracker}"
  }
}

variable "go-busybox" { default = "" }
target "go-busybox" {
  inherits   = ["docker-metadata-action"]
  context    = "./library/go-busybox"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
    "linux/ppc64le",
    "linux/riscv64",
  ]
  args = {
    VERSION = "${go-busybox}"
  }
}

variable "kine" { default = "" }
target "kine" {
  inherits   = ["docker-metadata-action"]
  context    = "./library/kine"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  args = {
    VERSION = "${kine}"
  }
}

variable "mc" { default = "" }
target "mc" {
  inherits   = ["docker-metadata-action"]
  context    = "./library/mc"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
    "linux/ppc64le",
  ]
  args = {
    VERSION = "${mc}"
  }
}

variable "zig" { default = "" }
target "zig" {
  inherits   = ["docker-metadata-action"]
  context    = "./library/zig"
  dockerfile = "Dockerfile"
  platforms = [
    "linux/amd64",
    "linux/arm64",
    "linux/ppc64le",
    "linux/riscv64",
  ]
  args = {
    VERSION = "${zig}"
  }
}
