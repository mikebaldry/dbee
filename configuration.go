package dbee

// Configuration holds the connection string for the database
type Configuration struct {
	// Driver is the name of the sql driver to use
	Driver string `json:"driver"`
	// ConnectionString is the options that go to the driver to tell it how to connect
	ConnectionString string `json:"connection_string"`
}
