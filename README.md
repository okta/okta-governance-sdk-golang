[<img src="https://cdn.brandfolder.io/2VK5Y09C/at/bb3mqsj5ssrgxtc5fbvtx/Logo-H_Developer-blue.svg" align="right" width="384px"/>](https://developer.okta.com/)
[![Build Status](https://img.shields.io/travis/okta/okta-sdk-golang.svg?logo=travis)](https://travis-ci.org/okta/okta-sdk-golang)
[![License](https://img.shields.io/github/license/okta/okta-sdk-golang.svg)](https://opensource.org/licenses/Apache-2.0)
[![Support](https://img.shields.io/badge/support-Developer%20Forum-blue.svg)][devforum]
[![API Reference](https://img.shields.io/badge/docs-reference-lightgrey.svg)][sdkapiref]

> [!WARNING]
> **Beta API Notice:** Some APIs exposed through this SDK, such as Grants, are currently in Beta.
> These Beta APIs are subject to change and may break or behave unexpectedly without prior notice.
> We recommend using them with caution in production environments.


# Okta Governance Golang SDK

* [Release status](#release-status)
* [Need help?](#need-help)
* [Getting started](#getting-started)
* [Usage guide](#usage-guide)
* [Configuration reference](#configuration-reference)
* [Upgrading Guide](#upgrading-to-20x)
* [Building the SDK](#building-the-sdk)
* [Contributing](#contributing)

This repository contains the Okta Internal Governance SDK for Golang. This SDK can be
used in your server-side code to interact with the Okta Internal Governance APIs and manage Okta campaigns, entitlements, reviews, grants etc.

* Manage campaigns with the [Campaigns API](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Campaigns/)
* Manage Entitlements with the [Entitlements API](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Entitlements/), [Entitlements-bundles](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Entitlement-Bundles/)
* Reassign and list Reviews with the [Reviews API](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Reviews/)
* Manage [Grants](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Grants/), [Principal Access](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Principal-Access/), [Principal Entitlements](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Principal-Entitlements/) and [much more](https://developer.okta.com/docs/api/iga/)!

## Release status

This library uses semantic versioning and follows Okta's [library version
policy](https://developer.okta.com/code/library-versions/).

| Version | Status                                                    |
|---------|-----------------------------------------------------------|
| 1.x     | :heavy_check_mark: Release                             |

## Need help?

If you run into problems using the SDK, you can

* Ask questions on the [Okta Developer Forums][devforum]
* Post [issues][github-issues] here on GitHub (for code errors)

### Install current release

To install the Okta Golang SDK in your project:
- Create a module file by running `go mod init`
    - You can skip this step if you already use `go mod`
- Run `go get github.com/okta/okta-governance-sdk-golang/v1@latest`. This will add
  the SDK to your `go.mod` file.
- Import the package in your project with `import
   "github.com/okta/okta-governance-sdk-golang/v1/governance"`

### You'll also need

* An Okta account, called an _organization_ (sign up for a free [developer
  organization](https://developer.okta.com/signup) if you need one)
* An [API
  token](https://developer.okta.com/docs/api/getting_started/getting_a_token)

### Initialize a client

Construct a client instance by passing it your Okta domain name and API token:

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v5/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)
}
```

Hard-coding the Okta domain and API token works for quick tests, but for real
projects you should use a more secure way of storing these values (such as
environment variables). This library supports a few different configuration
sources, covered in the [configuration reference](#configuration-reference)
section.

## Usage guide

These examples will help you understand how to use this library. You can also
browse the full [API reference documentation][sdkapiref].

Once you initialize a `client`, you can call methods to make requests to the
Okta API. Most methods are grouped by the API endpoint they belong to. For
example, methods that call the [Users
API](https://developer.okta.com/docs/api/resources/users) are organized under
`client.User`.

## Caching

In the default configuration the client utilizes a memory cache that has a time
to live on its cached values. See [Configuration Setter
Object](#configuration-setter-object)  `WithCache(cache bool)`,
`WithCacheTtl(i int32)`, and `WithCacheTti(i int32)`.  This helps to
keep HTTP requests to the Okta API at a minimum. In the case where the client
needs to be certain it is accessing recent data; for instance, list items,
delete an item, then list items again; be sure to make use of the refresh next
facility to clear the request cache. See [Refreshing Cache for Specific
Call](#refreshing-cache-for-specific-call). To completely disable the request
memory cache configure the client with `WithCache(false)`.

NOTE: Regardless of cache manager, Access Tokens from OAuth requests are always
cached.

## Connection Retry / Rate Limiting

By default this SDK retries requests that are returned with a 429 exception. To
disable this functionality set `OKTA_CLIENT_REQUEST_TIMEOUT` and
`OKTA_CLIENT_RATELIMIT_MAXRETRIES` to `0`.

Setting only one of the values to zero disables that check. Meaning, by
default, four retry attempts will be made. If you set
`OKTA_CLIENT_REQUEST_TIMEOUT` to 45 seconds and
`OKTA_CLIENT_RATELIMIT_MAXRETRIES` to `0`. This SDK will continue to retry
indefinitely for 45 seconds. If both values are non zero, this SDK attempts to
retry until either of the conditions are met (not both).

We use the Date header from the server to calculate the delta, as it's more
reliable than system time. But always add 1 second to account for some clock
skew in our service:

```go
backoff_seconds = header['X-Rate-Limit-Reset'] - header['Date'] + 1s
```

If the `backoff_seconds` calculation exceeds the request timeout, the initial
429 response will be allowed through without additional attempts.

When creating your client, you can pass in these settings like you would with
any other configuration.

```go
import (
	"fmt"
	"context"
	"github.com/okta/okta-sdk-golang/v5/okta"
)

func main() {
  config, err := okta.NewConfiguration(
    okta.WithOrgUrl("https://{yourOktaDomain}"),
    okta.WithToken("{apiToken}"),
    okta.WithRequestTimeout(45),
    okta.WithRateLimitMaxRetries(3),
  )
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := okta.NewAPIClient(config)
}
```

## Configuration reference

This library reuses the same configuration as in the okta-sdk-golang. Please refer [here]().

## Building the SDK

> [!WARNING]
> This sdk is still WIP, some endpoints might still break since they are in Beta.

In most cases, you won't need to build the SDK from source. If you want to
build it yourself, you'll need these prerequisites:

- Clone the repo
- Run `make build` from the root of the project

## ⚠️ Running Tests

Some of the integration tests currently rely on hardcoded IDs tied to a specific Okta tenant (e.g. campaign and entitlement IDs).
If you plan to run these tests locally, you must update the test data to match your Okta environment.


## Contributing

We're happy to accept contributions and PRs! Please see the [contribution
guide](/okta/okta-sdk-golang/blob/master/CONTRIBUTING.md) to understand how to
structure a contribution.

[devforum]: https://devforum.okta.com/
[sdkapiref]: https://godoc.org/github.com/okta/okta-sdk-golang/v4/okta
[lang-landing]: https://developer.okta.com/code/go/
[github-issues]: /okta/okta-sdk-golang/issues
[github-releases]: /okta/okta-sdk-golang/releases

