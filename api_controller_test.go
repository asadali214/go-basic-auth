package batester

import (
    "context"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

// TestAPIControllerBasicAuthTest tests the behavior of the APIController
func TestAPIControllerBasicAuthTest(t *testing.T) {
    ctx := context.Background()
    apiResponse, err := apiController.GetBasicAuthTest(ctx)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expected := `You've passed the test!`
    testHelper.RawBodyMatcher(t, expected, apiResponse.Response.Body)
}
