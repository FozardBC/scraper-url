package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local" env-required:"false"`
	//StoragePath string `yaml:"storage_path" env-required:"false"`
	//HTTPServer  HTTPServer `yaml:"http_server" env-required:"true"`
	Address string `yaml:"addres" env-default:"localhost:8080" env-required:"true"`
	Url     string `yaml:"url" env-required:"true"`
	Depth   int    `yaml:"depth" env-default:"1"`
}

// type HTTPServer struct {
// 	Address     string        `yaml:"address" env-default:"localhost:8080"`
// 	Timeout     time.Duration `yaml:"timeout" env-default:"30s"`
// 	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"120s"`
// 	User        string        `yaml:"user" env-required:"true"`
// 	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SEREVER_PASSWORD"`
// }

func MustLoad() *Config {
	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get projDirectory to set confing file: %s", err.Error())
	}

	configPath := filepath.Join(projectDir, "..", "..", "config", "cfg.yaml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err = cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("can't read config:%s", err.Error())
	}

	return &cfg

}
