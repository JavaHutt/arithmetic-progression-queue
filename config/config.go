package config

import (
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// Arithmetic Processor minimal number of workers
	concurrencyLimitMinValue = 1

	// HTTP Server port restrictions
	httpServerPortMinValue = 1
	httpServerPortMaxValue = 65535

	// HTTP Server shutdown timeout limits in seconds
	httpShutdownTimeoutMinValue = 0
	httpShutdownTimeoutMaxValue = 100

	// The main service shutdown timeout limits in seconds
	serviceShutdownTimeoutMinValue = 0
	serviceShutdownTimeoutMaxValue = 100
)

type Config interface {
	ConcurrencyLimit() int
	HTTPServerPort() string
	HTTPShutdownTimeout() time.Duration
	ServiceShutdownTimeout() time.Duration
}

// New constructs new config
func New() (Config, error) {
	cfg, err := configureViper()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func configureViper() (*configData, error) {
	err := setDefaults()
	if err != nil {
		return nil, fmt.Errorf("set defaults error: %w", err)
	}

	if err = bindFlags(); err != nil {
		return nil, fmt.Errorf("binding flags error: %w", err)
	}

	err = loadConfigToViper(".")
	if err != nil {
		fmt.Printf("cannot load the config file: %v\n", err)
	}

	viper.AutomaticEnv()

	var cfg configData
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal viper error: %w", err)
	}

	err = cfg.validate()
	if err != nil {
		return nil, fmt.Errorf("validate error: %w", err)
	}

	return &cfg, nil
}

func setDefaults() error {
	viper.SetDefault("CONCURRENCY_LIMIT", 4)
	viper.SetDefault("HTTP_SERVER_PORT", 8001)
	viper.SetDefault("HTTP_SHUTDOWN_TIMEOUT", 3)
	viper.SetDefault("SERVICE_SHUTDOWN_TIMEOUT", 10)

	return nil
}

func bindFlags() error {
	pflag.Bool("tracing", false, "If errors in the API should be traced")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}
	viper.RegisterAlias("TRACING_ENABLED", "tracing")
	return nil
}

func loadConfigToViper(path string) error {
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}

type configData struct {
	ConcurrencyLimit_       int `mapstructure:"CONCURRENCY_LIMIT"`
	HTTPServerPort_         int `mapstructure:"HTTP_SERVER_PORT"`
	HTTPShutdownTimeout_    int `mapstructure:"HTTP_SHUTDOWN_TIMEOUT"`
	ServiceShutdownTimeout_ int `mapstructure:"SERVICE_SHUTDOWN_TIMEOUT"`
}

func (cfg *configData) validate() (err error) {
	if cfg.ConcurrencyLimit_ < concurrencyLimitMinValue {
		err = fmt.Errorf("concurrency limit (%d) is too low: %w", cfg.HTTPShutdownTimeout_, err)
	}
	if cfg.HTTPServerPort_ < httpServerPortMinValue || cfg.HTTPServerPort_ > httpServerPortMaxValue {
		err = fmt.Errorf("http server port (%d) is out of range: %w", cfg.HTTPShutdownTimeout_, err)
	}
	if cfg.HTTPShutdownTimeout_ < httpShutdownTimeoutMinValue || cfg.HTTPShutdownTimeout_ > httpShutdownTimeoutMaxValue {
		err = fmt.Errorf("http shutdown timeout (%d) is out of range: %w", cfg.HTTPServerPort_, err)
	}
	if cfg.ServiceShutdownTimeout_ < serviceShutdownTimeoutMinValue || cfg.ServiceShutdownTimeout_ > serviceShutdownTimeoutMaxValue {
		err = fmt.Errorf("service shutdown timeout (%d) is out of range: %w", cfg.ServiceShutdownTimeout_, err)
	}
	return
}

func (cfg *configData) ConcurrencyLimit() int {
	return cfg.ConcurrencyLimit_
}

func (cfg *configData) HTTPServerPort() string {
	return fmt.Sprint(cfg.HTTPServerPort_)
}

func (cfg *configData) HTTPShutdownTimeout() time.Duration {
	return time.Duration(cfg.HTTPShutdownTimeout_) * time.Second
}

func (cfg *configData) ServiceShutdownTimeout() time.Duration {
	return time.Duration(cfg.ServiceShutdownTimeout_) * time.Second
}
