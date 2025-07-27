package config

var Config = WebookConfig{
	DB: DBConfig{
		DSN: "root:123456@tcp(localhost:33060)/mars",
	},
}
