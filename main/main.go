package main

import "auth/server"

func main() {
	s := server.New()
	s.Run()
}