package test

type AppConfig struct {
	AppName  string         `json:"app_name" yaml:"app_name" toml:"app_name"`
	AppEnv   string         `json:"app_env" yaml:"app_env" toml:"app_env" env:"APP_ENV"`
	Debug    bool           `json:"debug" yaml:"debug" toml:"debug" env:"APP_DEBUG"`
	Database DatabaseConfig `json:"database" yaml:"database" toml:"database"`
}
type DatabaseConfig struct {
	DSN string `json:"dsn" yaml:"dsn" toml:"dsn" env:"DB_DSN"`
}

var JsonExample = `
{
	"app_name": "test-app",
	"app_env": "test",
	"debug": true,
	"database": {
	  "dsn": "root@tcp(localhost:3306)/test"
	}
  }
`
