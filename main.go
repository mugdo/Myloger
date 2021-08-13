package main

import (
	"log"
	"net"

	"github.com/mugdo/Mylogr/customlog"
)

func getUserIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	//Get user ip
	IP := getUserIP()
	//create a loger fiele user given name
	object := customlog.CreateFile("test.txt")

	//method call
	object.Delete(1, IP)
	object.Update("John", 4)
	object.Warning(3, "bittu", IP)
	object.Endpoint(3, "log/mugdo/github.gom")

}
