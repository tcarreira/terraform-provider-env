# Terraform Provider env (Terraform Plugin Framework)

https://registry.terraform.io/providers/tcarreira/env

Provider for easy retrieving environment variables

## Use cases

- you already have a bunch of ENVs exported as part of some workflow and it would be cumbersome to use `TF_VAR_`

## Data sources

- env_var
- env_sensitive

## Example

```ruby
terraform {
  required_providers {
    env = {
      source  = "tcarreira/env"
      version = "0.2.0"
    }
  }
}

provider "env" {}


########## resources ##########

data "env_var" "test1" {
  id = "SHELL"
}

data "env_sensitive" "test2" {
  id = "SHELL"
}


########## outputs ##########

output "example" {
  value = data.env_var.test1.value
}
output "example2" {
  value     = data.env_sensitive.test2.value
  sensitive = true # if not set: Error: Output refers to sensitive values
}

# Outputs:
# example = "/usr/bin/zsh"
# example2 = <sensitive>
```
