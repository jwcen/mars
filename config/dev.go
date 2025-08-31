package config

var Config = WebookConfig{
	DB: DBConfig{
		DSN: "root:123456@tcp(localhost:33060)/mars?parseTime=true",
	},
	Redis: RedisConfig{
		Addr: "redis://localhost:6379",
	},
}
