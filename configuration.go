package batester

import (
    "batester/models"
    "log"
    "os"
    "strconv"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
    environment          Environment
    port                 string
    suites               models.SuiteCodeEnum
    httpConfiguration    HttpConfiguration
    basicAuthCredentials BasicAuthCredentials
}

// newConfiguration creates a new Configuration with the provided options.
func newConfiguration(options ...ConfigurationOptions) Configuration {
    config := Configuration{}
    
    for _, option := range options {
        option(&config)
    }
    return config
}

// cloneWithOptions provides configuration with the provided options.
func (config Configuration) cloneWithOptions(options ...ConfigurationOptions) Configuration {
    for _, option := range options {
        option(&config)
    }
    return config
}

// WithEnvironment is an option that sets the Environment in the Configuration.
func WithEnvironment(environment Environment) ConfigurationOptions {
    return func(c *Configuration) {
        c.environment = environment
    }
}

// WithPort is an option that sets the Port in the Configuration.
func WithPort(port string) ConfigurationOptions {
    return func(c *Configuration) {
        c.port = port
    }
}

// WithSuites is an option that sets the Suites in the Configuration.
func WithSuites(suites models.SuiteCodeEnum) ConfigurationOptions {
    return func(c *Configuration) {
        c.suites = suites
    }
}

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
    return func(c *Configuration) {
        c.httpConfiguration = httpConfiguration
    }
}

// WithBasicAuthCredentials is an option that sets the BasicAuthCredentials in the Configuration.
func WithBasicAuthCredentials(basicAuthCredentials BasicAuthCredentials) ConfigurationOptions {
    return func(c *Configuration) {
        c.basicAuthCredentials = basicAuthCredentials
    }
}

// Environment returns the Environment from the Configuration.
func (c Configuration) Environment() Environment {
    return c.environment
}

// Port returns the Port from the Configuration.
func (c Configuration) Port() string {
    return c.port
}

// Suites returns the Suites from the Configuration.
func (c Configuration) Suites() models.SuiteCodeEnum {
    return c.suites
}

// HttpConfiguration returns the HttpConfiguration from the Configuration.
func (c Configuration) HttpConfiguration() HttpConfiguration {
    return c.httpConfiguration
}

// BasicAuthCredentials returns the BasicAuthCredentials from the Configuration.
func (c Configuration) BasicAuthCredentials() BasicAuthCredentials {
    return c.basicAuthCredentials
}

// CreateConfigurationFromEnvironment creates a new Configuration with default settings.
// It also configures various Configuration options.
func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
    config := DefaultConfiguration()
    
    environment := os.Getenv("BATESTER_ENVIRONMENT")
    if environment != "" {
        config.environment = Environment(environment)
    }
    port := os.Getenv("BATESTER_PORT")
    if port != "" {
        config.port = port
    }
    suites := os.Getenv("BATESTER_SUITES")
    if suites != "" {
        temp, err := strconv.Atoi(suites)
        if err != nil {
            log.Fatalln(err)
        }
        config.suites = models.SuiteCodeEnum(temp)
    }
    basicAuthUserName := os.Getenv("BATESTER_BASIC_AUTH_USER_NAME")
    if basicAuthUserName != "" {
        config.basicAuthCredentials.username = basicAuthUserName
    }
    basicAuthPassword := os.Getenv("BATESTER_BASIC_AUTH_PASSWORD")
    if basicAuthPassword != "" {
        config.basicAuthCredentials.password = basicAuthPassword
    }
    for _, option := range options {
        option(&config)
    }
    return config
}

// Server represents available servers.
type Server string

const (
    ENUMDEFAULT Server = "default"
    AUTHSERVER  Server = "auth server"
)

// Environment represents available environments.
type Environment string

const (
    PRODUCTION Environment = "production"
    TESTING    Environment = "testing"
)

// CreateRetryConfiguration creates a new RetryConfiguration with the provided options.
func CreateRetryConfiguration(options ...RetryConfigurationOptions) RetryConfiguration {
    config := DefaultRetryConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

// CreateHttpConfiguration creates a new HttpConfiguration with the provided options.
func CreateHttpConfiguration(options ...HttpConfigurationOptions) HttpConfiguration {
    config := DefaultHttpConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

// CreateConfiguration creates a new Configuration with the provided options.
func CreateConfiguration(options ...ConfigurationOptions) Configuration {
    config := DefaultConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}
