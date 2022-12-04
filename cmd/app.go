package main

import (
	server "admin-server/internal"
)

func main() {
	server := server.Server{}

	server.Run()
}
