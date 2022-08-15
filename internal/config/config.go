package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
	defaultHTTPPort               = "8000"

	defaultNatsPort      = "4224"
	defaultNatsClusterID = "nats_cluster"
	defaultNatsClientID  = "nats_client"

	defaultPostgresPort = "5432"
	defaultPostgresHost = "localhost"

	defaultLimiterRPS   = 10
	defaultLimiterBurst = 2
	defaultLimiterTTL   = 10 * time.Minute
)

type (
	Config struct {
		HTTP     HTTPConfig
		DB       DatabaseConfig
		NATS     NATSConfig
		CacheTTL time.Duration `mapstructure:"ttl"`
	}

	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		Timeout            TimeoutConfig `mapstructure:"timeouts"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	NATSConfig struct {
		Port      string `mapstructure:"port"`
		ClusterID string `mapstructure:"clusterId"`
		ClientID  string `mapstructure:"clientId"`
	}

	DatabaseConfig struct {
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		User     string `mapstructure:"user"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		Password string `mapstructure:"password"`
		SSLMode  string `mapstructure:"sslMode"`
	}

	TimeoutConfig struct {
		Read  time.Duration `mapstructure:"read"`
		Write time.Duration `mapstructure:"write"`
	}
)

func Init(cfgDir string) (*Config, error) {
	setupDefaultValues()

	if err := parseConfigFile(cfgDir); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshall(&cfg); err != nil {
		return nil, err
	}

	parseEnvFile(&cfg)
	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshall(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("database", &cfg.DB); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("nats", &cfg.NATS); err != nil {
		return err
	}
	return viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL)
}

func parseEnvFile(cfg *Config) {
	cfg.DB.Postgres.Password = os.Getenv("DB_PASSWORD")
}

func setupDefaultValues() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderMegabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)

	viper.SetDefault("nats.port", defaultNatsPort)
	viper.SetDefault("nats.clusterId", defaultNatsClusterID)
	viper.SetDefault("nats.clientId", defaultNatsClientID)

	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)

	viper.SetDefault("database.postgres.host", defaultPostgresHost)
	viper.SetDefault("database.postgres.port", defaultPostgresPort)
}
