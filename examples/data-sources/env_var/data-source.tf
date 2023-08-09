data "env_var" "example" {
  id = "EXAMPLE_ENV"
}
output "out" {
  value = data.env_var.example
}

# if 'required', an error occurs when env is not found
data "env_var" "example" {
  id       = "EXAMPLE_ENV"
  required = true
}
