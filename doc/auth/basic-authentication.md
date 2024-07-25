
# Basic Authentication



Documentation for accessing and setting credentials for httpBasic.

## Auth Credentials

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| Username | `string` | The username to use with basic authentication | `WithUsername` | `Username` |
| Password | `string` | The password to use with basic authentication | `WithPassword` | `Password` |



**Note:** Required auth credentials can be set using `WithBasicAuthCredentials()` by providing a credentials instance with `NewBasicAuthCredentials()` in the configuration initialization and accessed using the `BasicAuthCredentials()` method in the configuration instance.

## Usage Example

### Client Initialization

You must provide credentials in the client as shown in the following code snippet.

```go
client := batester.NewClient(
    batester.CreateConfiguration(
        batester.WithBasicAuthCredentials(
            batester.NewBasicAuthCredentials(
                "BasicAuthUserName",
                "BasicAuthPassword",
            ),
        ),
    ),
)
```


