package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/tarm/serial"
)

func panicHandle() {
	x := recover()
	if x == "cannot open file." {
		fmt.Println("Cannot open dir")
	}
	if x == "error#1" {

		fmt.Println("Cannot accessfile")
	}

}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func banner() {
	fmt.Println(" ........................................  ")
	fmt.Println(" ....._..___.._..._...____......_...___..  ")
	fmt.Println(" ....|.||_._||.\\.|.|./.___|....|.|./._.\\.  ")
	fmt.Println(" ._..|.|.|.|.|..\\|.||.|.._.._..|.||.|.|.|  ")
	fmt.Println(" |.|_|.|.|.|.|.|\\..||.|_|.||.|_|.||.|_|.|  ")
	fmt.Println(" .\\___/.|___||_|.\\_|.\\____|.\\___/..\\___/.  ")
	fmt.Println(" ........................................  ")
	fmt.Println("develop by JingJo https://github.com/suthiphong")
	fmt.Println("build for :: windowx86, windowx64")
}

func start() {
	var port string
	var baudrate int
	Clear()
	fmt.Println("##########################")
	fmt.Println("##       INITIAL        ##")
	fmt.Println("##########################")
	fmt.Print("port : ")
	fmt.Scanln(&port)
	fmt.Print("baudrate : ")
	fmt.Scanln(&baudrate)
	_port = port
	_baudrate = baudrate
	defer menu()

}
func menu() {
	Clear()
	banner()
	var i int
	fmt.Println("######################################")
	fmt.Println("         Choose program.             ")
	fmt.Println("######################################")
	fmt.Println("1.PrintReport to Serial.")
	fmt.Println("2.Edit setting.")
	fmt.Println("0.Exit")
	fmt.Print("select : ")
	fmt.Scanln(&i)
	if i < 0 || i > 2 {
		menu()
	}
	switch i {
	case 1:
		print()
	case 2:
		start()
	case 0:
		Clear()
		fmt.Println("Goodbye")
	}

}
func print() {
	Clear()
	banner()
	var selected int
	fmt.Println("######################################")
	fmt.Println("    Report in dir ./Report.          ")
	fmt.Println("######################################")
	dirname := "./Report"
	f, err := os.Open(dirname)
	if err != nil {
		panic("cannot open file.")
	}
	files, err := f.Readdir(-1)
	if err != nil {
		panic("error")
	}
	f.Close()

	for i, file := range files {
		fmt.Println("=> ", i+1, file.Name())
	}
	fmt.Println("=>  0 to Back")
	fmt.Print("select : ")
	fmt.Scanln(&selected)
	if selected == 0 {
		menu()
	}
	if selected > len(files) {
		Clear()
		fmt.Println("Cannot access file please try again.")
		fmt.Println("Press any key.")
		fmt.Scanln()
		print()
	}
	if selected <= len(files) {
		filename := "./Report/" + files[selected-1].Name()
		fmt.Print(filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic("cannot open file.")
		}
		_data = string(data)
		printing()
	}

}
func printing() {
	Clear()
	fmt.Println("printing...")
	c := &serial.Config{Name: _port, Baud: _baudrate}
	s, err := serial.OpenPort(c)
	if err != nil {
		Clear()
		fmt.Println("cannot open port")
		fmt.Println("Press any key.")
		fmt.Scanln()
		menu()
	}

	s.Write([]byte(_data))
	s.Close()
	print()
}

var _port string
var _baudrate int
var _data string

func main() {
	Clear()
	start()
}
