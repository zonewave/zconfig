package test

// AppConfig is the config struct
type AppConfig struct {
	AppName  string         `json:"app_name" yaml:"app_name" toml:"app_name"`
	AppEnv   string         `json:"app_env" yaml:"app_env" toml:"app_env" env:"APP_ENV"`
	Debug    bool           `json:"debug" yaml:"debug" toml:"debug" env:"APP_DEBUG"`
	Database DatabaseConfig `json:"database" yaml:"database" toml:"database"`
}

// DatabaseConfig is the database config struct
type DatabaseConfig struct {
	DSN string `json:"dsn" yaml:"dsn" toml:"dsn" env:"DB_DSN"`
}

// JSONExample is the json example
var JSONExample = `
{
	"app_name": "test-app",
	"app_env": "test",
	"debug": true,
	"database": {
	  "dsn": "root@tcp(localhost:3306)/test"
	}
  }
`
