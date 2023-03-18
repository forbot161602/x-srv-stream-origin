package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/caarlos0/env/v7"

	"github.com/forbot161602/pbc-golang-lib/source/core/base/gbcfg"
)

var mConfig *Config

type Config struct {
	BasePath  string
	GitTag    string `env:"GIT_TAG,notEmpty"`
	GitCommit string `env:"GIT_COMMIT,notEmpty"`

	ServiceCode        string `env:"SRV_CODE" envDefault:"S002"`
	ServiceName        string `env:"SRV_NAME" envDefault:"stream-origin"`
	ServicePort        int    `env:"SRV_PORT" envDefault:"80"`
	ServiceProject     string `env:"SRV_PROJECT" envDefault:"open"`
	ServiceVersion     string `env:"SRV_VERSION" envDefault:"v1"`
	ServiceEnvironment string `env:"SRV_ENVIRONMENT,notEmpty"`
	ServiceLogLevel    string `env:"SRV_LOG_LEVEL" envDefault:"INFO"`
	ServiceTesting     bool   `env:"SRV_TESTING" envDefault:"false"`
	ServiceDebugging   bool   `env:"SRV_DEBUGGING" envDefault:"false"`
	ServiceDeveloping  bool
}

func GetConfig() *Config {
	if mConfig == nil {
		panic("Config hasn't been created.")
	}
	return mConfig
}

func SetConfig() {
	if mConfig == nil {
		mConfig = newConfig()
	}
	return
}

func newConfig() *Config {
	config := (&builder{}).
		initialize().
		setConfig().
		setBasePath().
		setServiceDeveloping().
		build()
	return config
}

type builder struct {
	config *Config
}

func (builder *builder) build() *Config {
	return builder.config
}

func (builder *builder) initialize() *builder {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		panic(err)
	}
	builder.config = config
	return builder
}

func (builder *builder) setConfig() *builder {
	gbcfg.SetConfig(builder.config)
	return builder
}

func (builder *builder) setBasePath() *builder {
	_, modulePath, _, _ := runtime.Caller(0)
	builder.config.BasePath = path.Dir(path.Dir(path.Dir(path.Dir(modulePath))))
	return builder
}

func (builder *builder) setServiceDeveloping() *builder {
	switch srvEnv := builder.config.ServiceEnvironment; srvEnv {
	case "local", "dev", "sit":
		builder.config.ServiceDeveloping = true
	case "uat", "stage", "prod":
		builder.config.ServiceDeveloping = false
	default:
		panic(fmt.Sprintf("Config does not support service environment `%s`.", srvEnv))
	}
	return builder
}

// Base definition

func GetBasePath() string {
	return GetConfig().GetBasePath()
}

func (config *Config) GetBasePath() string {
	return config.BasePath
}

func GetGitTag() string {
	return GetConfig().GetGitTag()
}

func (config *Config) GetGitTag() string {
	return config.GitTag
}

func GetGitCommit() string {
	return GetConfig().GetGitCommit()
}

func (config *Config) GetGitCommit() string {
	return config.GitCommit
}

// Core definition

func GetServiceCode() string {
	return GetConfig().GetServiceCode()
}

func (config *Config) GetServiceCode() string {
	return config.ServiceCode
}

func GetServiceName() string {
	return GetConfig().GetServiceName()
}

func (config *Config) GetServiceName() string {
	return config.ServiceName
}

func GetServicePort() int {
	return GetConfig().GetServicePort()
}

func (config *Config) GetServicePort() int {
	return config.ServicePort
}

func GetServiceProject() string {
	return GetConfig().GetServiceProject()
}

func (config *Config) GetServiceProject() string {
	return config.ServiceProject
}

func GetServiceVersion() string {
	return GetConfig().GetServiceVersion()
}

func (config *Config) GetServiceVersion() string {
	return config.ServiceVersion
}

func GetServiceEnvironment() string {
	return GetConfig().GetServiceEnvironment()
}

func (config *Config) GetServiceEnvironment() string {
	return config.ServiceEnvironment
}

func GetServiceLogLevel() string {
	return GetConfig().GetServiceLogLevel()
}

func (config *Config) GetServiceLogLevel() string {
	return config.ServiceLogLevel
}

func GetServiceTesting() bool {
	return GetConfig().GetServiceTesting()
}

func (config *Config) GetServiceTesting() bool {
	return config.ServiceTesting
}

func GetServiceDebugging() bool {
	return GetConfig().GetServiceDebugging()
}

func (config *Config) GetServiceDebugging() bool {
	return config.ServiceDebugging
}

func GetServiceDeveloping() bool {
	return GetConfig().GetServiceDeveloping()
}

func (config *Config) GetServiceDeveloping() bool {
	return config.ServiceDeveloping
}
