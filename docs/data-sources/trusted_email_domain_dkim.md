---
page_title: "pingone_trusted_email_domain_dkim Data Source - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Datasource to retrieve Trusted Email Domain DKIM status.
---

# pingone_trusted_email_domain_dkim (Data Source)

Datasource to retrieve Trusted Email Domain DKIM status.

## Example Usage

```terraform
data "pingone_trusted_email_domain_dkim" "email_domain_dkim" {
  environment_id = pingone_environment.my_environment.id

  trusted_email_domain_id = pingone_trusted_email_domain.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment.
- `trusted_email_domain_id` (String) A string that specifies the auto-generated ID of the email domain.

### Read-Only

- `id` (String) The ID of this resource.
- `region` (Set of Object) The regions collection specifies the properties for the 4 AWS SES regions that are used for sending email for the environment. The regions are determined by the geography where this environment was provisioned (North America, Canada, Europe & Asia-Pacific). (see [below for nested schema](#nestedatt--region))
- `type` (String) A string that specifies the type of DNS record.

<a id="nestedatt--region"></a>
### Nested Schema for `region`

Read-Only:

- `name` (String)
- `status` (String)
- `token` (Set of Object) (see [below for nested schema](#nestedobjatt--region--token))

<a id="nestedobjatt--region--token"></a>
### Nested Schema for `region.token`

Read-Only:

- `key` (String)
- `value` (String)
