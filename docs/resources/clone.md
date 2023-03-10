---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dblab_clone Resource - terraform-provider-dblab"
subcategory: ""
description: |-
  Database clone
---

# dblab_clone (Resource)

Database clone



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database_password` (String, Sensitive) Custom credentials for the database clone.
- `database_username` (String) Custom credentials for the database clone.
- `protected` (Boolean) When enabled no one can delete this clone via API & UI and automated deletion is also disabled. Note that destroy operation will unset it first.

### Optional

- `id` (String) Identifier of the clone, can be provided by user or generated by the server.
- `snapshot_id` (String) Snapshot identifier to use for the clone. If unset the latest snapshot will be used.

### Read-Only

- `database_name` (String) The name of the database in the clone.
- `port` (String) The port to acces clone's PostgreSQL instance.


