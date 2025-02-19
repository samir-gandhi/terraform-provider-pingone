---
page_title: "pingone_role Data Source - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Datasource to read PingOne admin role data
---

# pingone_role (Data Source)

Datasource to read PingOne admin role data

## Example Usage

```terraform
data "pingone_role" "organisation_admin" {
  name = "Organization Admin"
}

data "pingone_role" "environment_admin" {
  name = "Environment Admin"
}

data "pingone_role" "identity_data_admin" {
  name = "Identity Data Admin"
}

data "pingone_role" "client_application_developer" {
  name = "Client Application Developer"
}

data "pingone_role" "identity_data_admin_ro" {
  name = "Identity Data Read Only"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the role.

### Read-Only

- `description` (String) A description of the role.
- `id` (String) The ID of this resource.
