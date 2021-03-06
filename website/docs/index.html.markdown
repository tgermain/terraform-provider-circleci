---
layout: "circleci"
page_title: "Provider: CircleCI"
sidebar_current: "docs-circleci-index"
description: |-
  The CircleCI provider is used to to interact with the CircleCI API.
---

# CircleCI Provider

The CircleCI provider is used to interact with CircleCI API.

## Authentication

This provider requires a CircleCI API token in order to manage
resources.

## Example Usage

```hcl
provider "circleci" {
  api_token    = "${file("circleci_token")}"
  vcs_type     = "github"
  organization = "my_org"
}

# Create a context
resource "circleci_context" "build" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `api_token` - (Required) A CircleCI API token. This can also be set via the `CIRCLECI_TOKEN` environment variable.
* `vcs_type` - (Optional) The version control system, either `"github"` or `"bitbucket"`. Defaults to `"github"`. This can also be set via the `CIRCLECI_VCS_TYPE` environment variable.
* `organization` - (Optional) The organization where resources will be created. If unset, an organization must be provided with each resource. This can also be set via the `CIRCLECI_ORGANIZATION` environment variable.
* `url` - (Optional) The URL for the the CircleCI API (v1). Defaults to `"https://circleci.com/api/v1.1/"`. This value should generally only be set for testing. This can also be set via the `CIRCLECI_URL` environment variable.
* `graphql_url` - (Optional) The URL for the the CircleCI GraphQL API (unstable). Defaults to `"https://circleci.com/graphql-unstable"`. This value should generally only be set for testing. This can also be set via the `CIRCLECI_GRAPHQL_URL` environment variable.
