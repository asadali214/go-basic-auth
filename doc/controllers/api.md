# API

```go
aPIController := client.APIController()
```

## Class Name

`APIController`


# Get Basic Auth Test

```go
GetBasicAuthTest(
    ctx context.Context) (
    models.ApiResponse[string],
    error)
```

## Response Type

`string`

## Example Usage

```go
ctx := context.Background()

apiResponse, err := aPIController.GetBasicAuthTest(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

