target "docker-metadata-action" {}

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
