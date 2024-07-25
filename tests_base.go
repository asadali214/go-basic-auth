package batester

var apiController APIController

// init is an initialization function that sets up the controllers.
// It creates a configuration from the environment with a specified HTTP configuration and initializes the client.
// Then, it assigns the different controllers from the client to the corresponding variables for further use.
func init() {
    
    config := CreateConfigurationFromEnvironment(
        WithEnvironment("testing"),
        WithHttpConfiguration(
            CreateHttpConfiguration(
                WithTimeout(100),
            ),
        ),
    )
    
    client := NewClient(config)
    
    apiController = *client.APIController()
}
