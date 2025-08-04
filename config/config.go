package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host     string `mapstructure:"host"`
		HTTPPort int    `mapstructure:"http_port"`
	}

	JWT struct {
		SecretKey          string        `mapstructure:"secret_key"`
		AccessTokenExpiry  time.Duration `mapstructure:"access_token_expiry"`
		RefreshTokenExpiry time.Duration `mapstructure:"refresh_token_expiry"`
		Issuer             string        `mapstructure:"issuer"`
	} `mapstructure:"jwt"`

	GRPCClients struct {
		AuthServiceAddr    string `mapstructure:"auth_service_addr"`
		ControlPlaneAddr   string `mapstructure:"control_plane_addr"`
		KubeManagerAddr    string `mapstructure:"kube_manager_addr"`
		BillingServiceAddr string `mapstructure:"billing_service_addr"`
		NodeServiceAddr    string `mapstructure:"node_service_addr"`
	} `mapstructure:"grpc_clients"`

	Metrics struct {
		PrometheusEnabled bool `mapstructure:"prometheus_enabled"`
		PrometheusPort    int  `mapstructure:"prometheus_port"`
	} `mapstructure:"metrics"`

	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"logging"`
}

var AppConfig Config

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("erreur de chargement du fichier de configuration: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("erreur d'analyse du fichier de configuration: %w", err)
	}

	AppConfig = config

	log.Println("[CONFIG] Fichier config.yaml chargé avec succès.")
	return &config, nil
}
