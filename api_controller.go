package batester

import (
    "batester/models"
    "context"
)

// APIController represents a controller struct.
type APIController struct {
    baseController
}

// NewAPIController creates a new instance of APIController.
// It takes a baseController as a parameter and returns a pointer to the APIController.
func NewAPIController(baseController baseController) *APIController {
    aPIController := APIController{baseController: baseController}
    return &aPIController
}

// GetBasicAuthTest takes context as parameters and
// returns an models.ApiResponse with string data and
// an error if there was an issue with the request or response.
func (a *APIController) GetBasicAuthTest(ctx context.Context) (
    models.ApiResponse[string],
    error) {
    req := a.prepareRequest(ctx, "GET", "/auth/basic")
    req.Authenticate(NewAuth("httpBasic"))
    str, resp, err := req.CallAsText()
    var result string = str

    if err != nil {
        return models.NewApiResponse(result, resp), err
    }
    return models.NewApiResponse(result, resp), err
}
