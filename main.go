package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/mugdo/Mylogr/customlog"
)

const dot = "."
const hypen = "-"
const and = "&"

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

	// Get time now
	Ntime := getTime()

	//create a loger fiele user given name
	object := customlog.CreateFile("test.txt")

	//method call
	object.Delete(Ntime, 1, IP)
	object.Update(Ntime, "John", 4)
	object.Warning(Ntime, 3, "bittu", IP)
	object.Endpoint(Ntime, 3, "log/mugdo/github.gom")

}
func getTime() string {
	today := time.Now()
	date := strconv.Itoa(today.Year()) + hypen + strconv.Itoa(int(today.Month())) + hypen + strconv.Itoa(today.Day())
	date += " " + and + " " + strconv.Itoa(today.Hour()) + dot + strconv.Itoa(today.Minute()) + dot + strconv.Itoa(today.Second())

	return date
}
