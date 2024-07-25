
# Getting Started with BATester

## Introduction

### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the batester library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace batester => ".\\BATester" // local path to the SDK

require batester v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the root directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `port` | `string` | *Default*: `"80"` |
| `suites` | `models.SuiteCodeEnum` | *Default*: `1` |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.TESTING`** |
| `httpConfiguration` | [`HttpConfiguration`](doc/http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthCredentials` | [`BasicAuthCredentials`](doc/auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |

The API client can be initialized as follows:

```go
client := batester.NewClient(
    batester.CreateConfiguration(
        batester.WithHttpConfiguration(
            batester.CreateHttpConfiguration(
                batester.WithTimeout(0),
            ),
        ),
        batester.WithEnvironment(batester.TESTING),
        batester.WithBasicAuthCredentials(
            batester.NewBasicAuthCredentials(
                "BasicAuthUserName",
                "BasicAuthPassword",
            ),
        ),
        batester.WithPort("80"),
        batester.WithSuites(1),
    ),
)
```

## Environments

The SDK can be configured to use a different environment for making API calls. Available environments are:

### Fields

| Name | Description |
|  --- | --- |
| production | - |
| testing | **Default** |

## Authorization

This API uses the following authentication schemes.

* [`httpBasic (Basic Authentication)`](doc/auth/basic-authentication.md)

## List of APIs

* [API](doc/controllers/api.md)

## Classes Documentation

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)

