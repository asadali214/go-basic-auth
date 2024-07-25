
# HttpConfiguration

The following parameters are configurable for the HttpConfiguration:

## Properties

| Name | Type | Description |
|  --- | --- | --- |
| `timeout` | `float64` | Timeout in milliseconds.<br>*Default*: `0` |
| `transport` | `http.RoundTripper` | Establishes network connection and caches them for reuse.<br>*Default*: `http.DefaultTransport` |
| `retryConfiguration` | [`batester.RetryConfiguration`](retry-configuration.md) | Configurations to retry requests.<br>*Default*: `DefaultRetryConfiguration()` |

The httpConfiguration can be initialized as follows:

```go
httpConfiguration := CreateHttpConfiguration(
    batester.WithTimeout(0),
    batester.WithTransport(http.DefaultTransport),
    batester.WithRetryConfiguration(DefaultRetryConfiguration()),
)
```

