package config

import (
	"github.com/forbot161602/x-lib-go/source/entry/precfg"
)

var mConfig *Config

func GetConfig() *Config {
	if mConfig == nil {
		mConfig = newConfig()
	}
	return mConfig
}

func newConfig() *Config {
	config := (&builder{}).initialize().
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
	precfg.Config
	ServiceCode         string `json:"serviceCode" env:"SRV_CODE" envDefault:"S002"`
	ServiceName         string `json:"serviceName" env:"SRV_NAME" envDefault:"stream-origin"`
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

type builder struct {
	config *Config
}

func (builder *builder) build() *Config {
	return builder.config
}

func (builder *builder) initialize() *builder {
	builder.config = &Config{}
	return builder
}

func (builder *builder) parseEnv() *builder {
	precfg.ParseEnv(builder.config)
	return builder
}

func (builder *builder) setBasePath() *builder {
	builder.config.BasePath = precfg.MakeBasePath(5)
	return builder
}

func (builder *builder) setServiceID() *builder {
	builder.config.ServiceID = precfg.MakeServiceID()
	return builder
}

func (builder *builder) setServiceDeveloping() *builder {
	builder.config.ServiceDeveloping = precfg.MakeServiceDeveloping(builder.config.ServiceEnvironment)
	return builder
}
