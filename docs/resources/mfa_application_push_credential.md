---
page_title: "pingone_mfa_application_push_credential Resource - terraform-provider-pingone"
subcategory: "MFA"
description: |-
  Resource to create and manage push credentials for a mobile MFA application configured in PingOne.
---

# pingone_mfa_application_push_credential (Resource)

Resource to create and manage push credentials for a mobile MFA application configured in PingOne.

## Example Usage

```terraform
resource "pingone_environment" "my_environment" {
  # ...
}

resource "pingone_application" "my_application" {
  # ...
}

resource "pingone_mfa_application_push_credential" "example_fcm" {
  environment_id = pingone_environment.my_environment.id
  application_id = pingone_application.my_application.id

  fcm {
    key = var.fcm_key
  }
}

resource "pingone_mfa_application_push_credential" "example_apns" {
  environment_id = pingone_environment.my_environment.id
  application_id = pingone_application.my_application.id

  apns {
    key               = var.apns_key
    team_id           = var.apns_team_id
    token_signing_key = var.apns_token_signing_key
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String) The ID of the application to create the push notification credential for.
- `environment_id` (String) The ID of the environment to create the application push notification credential in.

### Optional

- `apns` (Block List, Max: 1) A block that specifies the credential settings for the Apple Push Notification Service. (see [below for nested schema](#nestedblock--apns))
- `fcm` (Block List, Max: 1) A block that specifies the credential settings for the Firebase Cloud Messaging service. (see [below for nested schema](#nestedblock--fcm))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--apns"></a>
### Nested Schema for `apns`

Required:

- `key` (String, Sensitive) A string that Apple uses as an identifier to identify an authentication key.
- `team_id` (String) A string that Apple uses as an identifier to identify teams.
- `token_signing_key` (String, Sensitive) A string that Apple uses as the authentication token signing key to securely connect to APNS. This is the contents of a p8 file with a private key format.


<a id="nestedblock--fcm"></a>
### Nested Schema for `fcm`

Required:

- `key` (String, Sensitive) A string that represents the server key of the Firebase cloud messaging service.

## Import

Import is supported using the following syntax:

```shell
$ terraform import pingone_mfa_application_push_credential.example <environment_id>/<application_id>/<push_credential_id>
```
