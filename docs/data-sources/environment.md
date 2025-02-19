---
page_title: "pingone_environment Data Source - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Datasource to read PingOne environment data
---

# pingone_environment (Data Source)

Datasource to read PingOne environment data

## Example Usage

```terraform
data "pingone_environment" "example_by_name" {
  name = "foo"
}

data "pingone_environment" "example_by_id" {
  environment_id = var.environment_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `environment_id` (String) The ID of the environment.
- `name` (String) The name of the environment.

### Read-Only

- `description` (String) A description of the environment.
- `id` (String) The ID of this resource.
- `license_id` (String) An ID of a valid license to apply to the environment.
- `organization_id` (String) The ID of the PingOne organization tenant to which the environment belongs.
- `region` (String) The region the environment is created in.
- `service` (Set of Object) The services enabled in the environment. (see [below for nested schema](#nestedatt--service))
- `solution` (String) The solution context of the environment.  Blank values indicate a custom solution context, without workforce solution additions.  Expected values are `WORKFORCE`, `CUSTOMER` or no value for a custom solution context.
- `type` (String) The type of the environment.  Options are SANDBOX for a development/testing environment and PRODUCTION for environments that require protection from deletion.

<a id="nestedatt--service"></a>
### Nested Schema for `service`

Read-Only:

- `bookmark` (Set of Object) (see [below for nested schema](#nestedobjatt--service--bookmark))
- `console_url` (String)
- `type` (String)

<a id="nestedobjatt--service--bookmark"></a>
### Nested Schema for `service.bookmark`

Read-Only:

- `name` (String)
- `url` (String)
