package configs

import "fmt"

type AppConfig struct {
	App      App      `mapstructure:"app"`
	Postgres Postgres `mapstructure:"postgres"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type Postgres struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"db"`
	LogLevel uint8  `mapstructure:"log_level"`
}

func (Postgres) GenerateDSN(user, password, dbName string) string {
	return fmt.Sprintf("postgres://%s:%s@postgres:5432/%s", user, password, dbName)
}
