---
page_title: "pingone_trusted_email_domain Resource - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Resource to create and manage PingOne Trusted Email Domains.
---

# pingone_trusted_email_domain (Resource)

Resource to create and manage PingOne Trusted Email Domains.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_trusted_email_domain" "my_custom_email_domain" {
  environment_id = pingone_environment.my_environment.id

  domain_name = "demo.bxretail.org"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain_name` (String) A string that specifies the domain name to use, which must be provided and must be unique within an environment (for example, `demo.bxretail.org`).
- `environment_id` (String) The ID of the environment to create the certificate in.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_trusted_email_domain.example <environment_id>/<trusted_email_domain_id>
```
