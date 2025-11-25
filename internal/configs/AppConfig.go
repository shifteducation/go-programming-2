package configs

import "strconv"

type AppConfig struct {
	App App `mapstructure:"app"`
	DB  DB  `mapstructure:"db"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type DB struct {
	DSN      string `mapstructure:"dsn"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"db_name"`
	LogLevel uint8  `mapstructure:"log_level"`
}

func (db DB) GetDBConnectionString() string {
	return db.User + db.Password + "@" + db.Host + ":" + strconv.Itoa(db.Port) + "/" + db.DbName
}
