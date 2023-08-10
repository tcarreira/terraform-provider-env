data "env_var" "example" {
  id       = "EXAMPLE_ENV"
  required = true # (optional) plan will error if not found
}

output "out" {
  value = data.env_var.example.value
}
