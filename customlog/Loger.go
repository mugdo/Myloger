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
	Update(fileName string, name string, id int)
	Delete(fileName string, id int, IP net.IP)
	Endpoint(fileName string, id int, path string)
	Warning(fileName string, id int, name string, IP net.IP)
}

// structe for custom loger "a vale can be more than one type"
type LogService struct {
	File string
}

//Method for loger

func (Ls *LogService) Update(name string, id int) {
	// conver int to string , i can also use id interface type: it's take all type of value
	ID := strconv.Itoa(id)
	data := "\n Type : (Update)" + "userId : {" + ID + "}, Name : [" + name + "] : `upadate data.'"
	//get file name using getfile function
	Getfilename := Ls.getfile()
	//wrrite date by user given it coverted into string
	Getfilename.WriteString(data)

}
func (Ls *LogService) Delete(id int, IP net.IP) {
	ID := strconv.Itoa(id)
	data := "\n Type : (Delete)" + " userID : {" + ID + "} " + "IP : {" + IP.String() + "}: `Delete data By Id`."
	Getfilename := Ls.getfile()
	Getfilename.WriteString(data)

}
func (Ls *LogService) Endpoint(id int, path string) {
	ID := strconv.Itoa(id)
	data := "\n Type : (Endpoint)" + " userID : {" + ID + "} " + "Endpoint : {" + path + "}: ``."
	Getfilename := Ls.getfile()
	Getfilename.WriteString(data)
}
func (Ls *LogService) Warning(id int, name string, IP net.IP) {
	ID := strconv.Itoa(id)
	data := "\n Type : (Warning)" + " userID : {" + ID + "} " + "IP : {" + IP.String() + "}: `Something worng`."
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
