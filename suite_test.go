package zconfig

import (
	"github.com/golang/mock/gomock"
	"github.com/zonewave/pkgs/mock/aferomock"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type Suite struct {
	suite.Suite
	ctrl      *gomock.Controller
	mockFs    *aferomock.MockFs
	mockAfero *aferomock.MockAfero
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (s *Suite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.mockFs = aferomock.NewMockFs(s.ctrl)
	s.mockAfero = aferomock.NewMockAfero(s.ctrl)
}

func (s *Suite) TearDownSuite() {
	s.ctrl.Finish()
}

// The SetupTest method will be run before every test in the suite.
func (s *Suite) SetupTest() {
}

// The TearDownTest method will be run after every test in the suite.
func (s *Suite) TearDownTest() {
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

type AppConfig struct {
	AppName  string         `json:"app_name" yaml:"app_name" toml:"app_name"`
	AppEnv   string         `json:"app_env" yaml:"app_env" toml:"app_env" env:"APP_ENV"`
	Debug    bool           `json:"debug" yaml:"debug" toml:"debug" env:"APP_DEBUG"`
	Database DatabaseConfig `json:"database" yaml:"database" toml:"database"`
}
type DatabaseConfig struct {
	DSN string `json:"dsn" yaml:"dsn" toml:"dsn" env:"DB_DSN"`
}

var jsonExample = `
{
	"app_name": "test-app",
	"app_env": "test",
	"debug": true,
	"database": {
	  "dsn": "root@tcp(localhost:3306)/test"
	}
  }
`
