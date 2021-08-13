package customlog

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

//Interface create
type Clog interface {
	Update(time string, fileName string, name string, id int)
	Delete(time string, fileName string, id int, IP net.IP)
	Endpoint(time string, fileName string, id int, path string)
	Warning(time string, fileName string, id int, name string, IP net.IP)
}

// structe for custom loger "a vale can be more than one type"
type LogService struct {
	File string
}

//Method for loger

func (Ls *LogService) Update(Ntime string, name string, id int) {
	// conver int to string , i can also use id interface type: it's take all type of value
	ID := strconv.Itoa(id)

	data := "\n Type : (Update)" + " Date & time : [" + Ntime + "] userId : {" + ID + "}, Name : [" + name + "] : `upadate data.'"
	//get file name using getfile function
	Getfilename := Ls.getfile()
	//wrrite date by user given it coverted into string
	Getfilename.WriteString(data)

}
func (Ls *LogService) Delete(Ntime string, id int, IP net.IP) {
	ID := strconv.Itoa(id)
	data := "\n Type : (Delete)" + "Time & Date : [" + Ntime + "] userID : {" + ID + "} " + "IP : {" + IP.String() + "}: `Delete data By Id`."
	Getfilename := Ls.getfile()
	Getfilename.WriteString(data)

}
func (Ls *LogService) Endpoint(Ntime string, id int, path string) {
	ID := strconv.Itoa(id)
	data := "\n Type : (Endpoint) " + "Time & Date : [" + Ntime + "] userID : {" + ID + "} " + "Endpoint : {" + path + "}: ``."
	Getfilename := Ls.getfile()
	Getfilename.WriteString(data)
}
func (Ls *LogService) Warning(Ntime string, id int, name string, IP net.IP) {
	ID := strconv.Itoa(id)
	data := "\n Type : (Warning) " + "Time & Date : [" + Ntime + "] userID : {" + ID + "} " + "IP : {" + IP.String() + "}: `Something worng`."
	Getfilename := Ls.getfile()
	Getfilename.WriteString(data)

}

// crete a new file 0666 is the file parmission for Read and write
func CreateFile(file string) LogService {
	var err error
	//set the name on struct
	Logfile := LogService{File: file}
	Ls, err := os.OpenFile(file, os.O_CREATE, 0666)

	if err != nil {
		log.Fatal("File not working", Ls)

	}
	return Logfile

}

//get method for append message
func (Ls *LogService) getfile() *os.File {
	f, err := os.OpenFile(Ls.File, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	return f
}
