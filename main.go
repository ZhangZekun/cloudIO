package main

import (
    "webTest/cloudIO/service"
)
func main() {
	server := service.NewServer();
	server.Run();                                          
}									