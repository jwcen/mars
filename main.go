package main

import "github.com/jwcen/mars/internal/apiserver/integration/startup"

func main() {
	server := startup.InitWebServer()
	err := server.Run(":8081")
	if err != nil {
		panic("端口启动失败")
	}
}
