---
page_title: "Provider: env"
description: |-
  The env provider is used to easy retrieve environment variables
---

# Env Provider

The env provider allows the use of environment variables within Terraform
configurations. This is a *logical provider*, which means that it works
entirely within Terraform's logic, and doesn't interact with any other
services.

Using environment variables is not reccommended within a Terraform configuration,
the reccommended way is using [variables](https://developer.hashicorp.com/terraform/language/values/variables#declaring-an-input-variable)
that can be passed with an environment variable in the form of
[TF_VAR_name](https://developer.hashicorp.com/terraform/cli/config/environment-variables#tf_var_name).

But sometimes, an environment variable is already being consistently used
(eg: for configuring a provider) and it would be useful to use the same Env
in order to avoid repetition and possible divergence.
If this is the case, please use this provider.

For example:

```terraform
data "env_var" "example" {
  id       = "EXAMPLE_ENV"
  required = true # (optional) plan will error if not found
}

output "out" {
  value = data.env_var.example.value
}
```