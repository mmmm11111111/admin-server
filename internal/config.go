package server

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type DBCongig struct {
	User     string
	Password string
	Port     string
	Host     string
	DBname   string
}

type ServerConfig struct {
}

type Config struct {
	Redis RedisConfig
	DB    DBCongig
	ServerConfig
}
