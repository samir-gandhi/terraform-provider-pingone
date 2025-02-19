---
page_title: "pingone_webhook Resource - terraform-provider-pingone"
subcategory: "Platform"
description: |-
  Resource to create and manage PingOne Webhooks / Data Subscriptions.
---

# pingone_webhook (Resource)

Resource to create and manage PingOne Webhooks / Data Subscriptions.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_webhook" "my_webhook" {
  environment_id = pingone_environment.my_environment.id

  name    = "My webhook"
  enabled = true

  http_endpoint_url = "https://audit.bxretail.org/"
  http_endpoint_headers = {
    Authorization = "Basic usernamepassword"
  }

  format = "ACTIVITY"

  filter_options {
    included_action_types = ["ACCOUNT.LINKED", "ACCOUNT.UNLINKED"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) The ID of the environment to create the webhook in.
- `filter_options` (Block List, Min: 1, Max: 1) A block that specifies the PingOne platform event filters to be included to trigger this webhook. (see [below for nested schema](#nestedblock--filter_options))
- `format` (String) A string that specifies one of the supported webhook formats. Options are `ACTIVITY`, `SPLUNK`, and `NEWRELIC`.
- `http_endpoint_url` (String) A string that specifies a valid HTTPS URL to which event messages are sent.
- `name` (String) A string that specifies the webhook name.

### Optional

- `enabled` (Boolean) A boolean that specifies whether a created or updated webhook should be active or suspended. A suspended state (`"enabled":false`) accumulates all matched events, but these events are not delivered until the webhook becomes active again (`"enabled":true`). For suspended webhooks, events accumulate for a maximum of two weeks. Events older than two weeks are deleted. Restarted webhooks receive the saved events (up to two weeks from the restart date). Defaults to `false`.
- `http_endpoint_headers` (Map of String) A map that specifies the headers applied to the outbound request (for example, `Authorization` `Basic usernamepassword`. The purpose of these headers is for the HTTPS endpoint to authenticate the PingOne service, ensuring that the information from PingOne is from a trusted source.
- `verify_tls_certificates` (Boolean) A boolean that specifies whether a certificates should be verified. If this property's value is set to `false`, then all certificates are trusted. (Setting this property's value to false introduces a security risk.) Defaults to `true`.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--filter_options"></a>
### Nested Schema for `filter_options`

Required:

- `included_action_types` (Set of String) A non-empty list that specifies the list of action types that should be matched for the webhook.

Optional:

- `included_application_ids` (Set of String) An array that specifies the list of applications (by ID) whose events are monitored by the webhook (maximum of 10 IDs in the array). If a list of applications is not provided, events are monitored for all applications in the environment.
- `included_population_ids` (Set of String) An array that specifies the list of populations (by ID) whose events are monitored by the webhook (maximum of 10 IDs in the array). This property matches events for users in the specified populations, as opposed to events generated in which the user in one of the populations is the actor.
- `included_tags` (Set of String) An array of tags that events must have to be monitored by the webhook. If tags are not specified, all events are monitored. Currently, the available tags are `adminIdentityEvent`. Identifies the event as the action of an administrator on other administrators.

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_webhook.example <environment_id>/<webhook_id>
```
