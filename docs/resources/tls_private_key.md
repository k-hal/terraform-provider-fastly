---
layout: "fastly"
page_title: "Fastly: tls_private_key"
sidebar_current: "docs-fastly-resource-tls_private_key"
description: |-
Uploads a Custom TLS Private Key
---

# fastly_tls_private_key

Uploads a Custom TLS Private Key to Fastly. This can be combined with a `fastly_tls_custom_certificate` resource to provide a TLS Certificate able to be applied to a Fastly Service.

The Private Key resource requires a key in PEM format, and a name to identify it.

## Example Usage

Basic usage:

```hcl
resource "tls_private_key" "demo" {
  algorithm = "RSA"
}

resource "fastly_tls_private_key" "demo" {
  key_pem = tls_private_key.demo.private_key_pem
  name    = "tf-demo"
}
```

## Import

A Private Key can be imported using its ID, e.g.

```
$ terraform import fastly_tls_private_key.demo xxxxxxxxxxx
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **key_pem** (String, Sensitive) Private key in PEM format.
- **name** (String) Customisable name of the private key.

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **created_at** (String) Time-stamp (GMT) when the private key was created.
- **key_length** (Number) The key length used to generate the private key.
- **key_type** (String) The algorithm used to generate the private key. Must be RSA.
- **public_key_sha1** (String) Useful for safely identifying the key.
- **replace** (Boolean) Whether Fastly recommends replacing this private key.