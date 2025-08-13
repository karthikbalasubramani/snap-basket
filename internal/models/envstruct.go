package models

// Environmental Values Configuration Go Server
type ServerConfig struct {
	Port     string
	Protocol string
}

// Environmental Values Configuration Database
type DatabaseConfig struct {
	Uri            string
	DatabaseName   string
	UserCollection string
}
