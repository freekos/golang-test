package configs

type EnvConfig struct {
	PORT         int    `json:"port"`
	HOST         string `json:"host"`
	DB_PASSWORD  string `json:"password"`
	DB_USER      string `json:"db_user"`
	DB_NAME      string `json:"db_name"`
	DB_HOST      string `json:"db_host"`
	DB_PORT      int    `json:"db_port"`
	DB_SSL_MODE  string `json:"db_ssl_mode"`
	DB_TIME_ZONE string `json:"db_time_zone"`
}

var Config *EnvConfig = &EnvConfig{
	HOST:        "0.0.0.0",
	PORT:        8080,
	DB_USER:     "postgres",
	DB_NAME:     "postgres",
	DB_PASSWORD: "example",
	DB_HOST:     "postgres_container",
	DB_PORT:     5050,
	DB_SSL_MODE: "disable",
}
