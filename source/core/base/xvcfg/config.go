package xvcfg

import (
	"github.com/starryck/strk-tc-x-lib-go/source/entry/xbprecfg"
)

var mConfig *Config

func GetConfig() *Config {
	if mConfig == nil {
		mConfig = newConfig()
	}
	return mConfig
}

func newConfig() *Config {
	config := (&configBuilder{}).initialize().
		parseEnv().
		setBasePath().
		setServiceID().
		setServiceDeveloping().
		build()
	return config
}

// Base definition

func GetBasePath() string {
	return GetConfig().GetBasePath()
}

func GetGitTag() string {
	return GetConfig().GetGitTag()
}

func GetGitCommit() string {
	return GetConfig().GetGitCommit()
}

// Core definition

func GetServiceID() string {
	return GetConfig().GetServiceID()
}

func GetServiceCode() string {
	return GetConfig().GetServiceCode()
}

func GetServiceName() string {
	return GetConfig().GetServiceName()
}

func GetServicePort() int {
	return GetConfig().GetServicePort()
}

func GetServiceProject() string {
	return GetConfig().GetServiceProject()
}

func GetServiceVersion() string {
	return GetConfig().GetServiceVersion()
}

func GetServiceEnvironment() string {
	return GetConfig().GetServiceEnvironment()
}

func GetServiceLogLevel() string {
	return GetConfig().GetServiceLogLevel()
}

func GetServiceTesting() bool {
	return GetConfig().GetServiceTesting()
}

func GetServiceDebugging() bool {
	return GetConfig().GetServiceDebugging()
}

func GetServiceDeveloping() bool {
	return GetConfig().GetServiceDeveloping()
}

// Service definition

func GetStreamDirectoryPath() string {
	return GetConfig().GetStreamDirectoryPath()
}

type Config struct {
	xbprecfg.Config
	ServiceCode         string `json:"serviceCode" env:"SRV_CODE" envDefault:"S002"`
	ServiceName         string `json:"serviceName" env:"SRV_NAME" envDefault:"srv-stream-origin"`
	StreamDirectoryPath string `json:"streamDirectoryPath" env:"STREAM_DIRECTORY_PATH,notEmpty"`
}

func (config *Config) GetServiceCode() string {
	return config.ServiceCode
}

func (config *Config) GetServiceName() string {
	return config.ServiceName
}

func (config *Config) GetStreamDirectoryPath() string {
	return config.StreamDirectoryPath
}

type configBuilder struct {
	config *Config
}

func (builder *configBuilder) build() *Config {
	return builder.config
}

func (builder *configBuilder) initialize() *configBuilder {
	builder.config = &Config{}
	return builder
}

func (builder *configBuilder) parseEnv() *configBuilder {
	xbprecfg.ParseEnv(builder.config)
	return builder
}

func (builder *configBuilder) setBasePath() *configBuilder {
	builder.config.BasePath = xbprecfg.MakeBasePath(5)
	return builder
}

func (builder *configBuilder) setServiceID() *configBuilder {
	builder.config.ServiceID = xbprecfg.MakeServiceID()
	return builder
}

func (builder *configBuilder) setServiceDeveloping() *configBuilder {
	builder.config.ServiceDeveloping = xbprecfg.MakeServiceDeveloping(builder.config.ServiceEnvironment)
	return builder
}
