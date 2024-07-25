
# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `port` | `string` | *Default*: `"80"` |
| `suites` | `models.SuiteCodeEnum` | *Default*: `1` |
| `environment` | `Environment` | The API environment. <br> **Default: `Environment.TESTING`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthCredentials` | [`BasicAuthCredentials`](auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |

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

## BATester Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| aPI | Gets APIController |

