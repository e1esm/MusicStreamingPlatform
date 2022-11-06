package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func GetConfig() *Config {
	return &Config{DB: &DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "E9602922",
		Dbname:   "musicstreamingplatform",
	}}
}
