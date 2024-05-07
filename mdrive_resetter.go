package main

import (
	"go.bug.st/serial"
)
import (
	"fmt"
	"log"
)

func main() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	openPort()
}

func openPort() {
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		//		Parity:   serial.NoParity,
		//		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open("/dev/tty.usbserial-1120", mode)
	if err != nil {
		log.Fatal(err)
	}

	//port.Write([]byte("\rFD\r \rCK=0…\r \n*FD\n \n*FDÌ\n"))
	//
	// factory default string: ^MFD^M ^d1000 ^MCK=0…^M ^d1000 ^J*FD^J^d1000^J*FDÌ^J

	//
	// ^MFD^M
	//
	//^d1000
	//	time.Sleep(1 * time.Second)

	//	port.Write([]byte("\rFD\r ^d1000 \rCK=0…\r ^d1000 \n*FD\n^d1000\n*FDÌ\n"))

	//fmt.Printf("Clearing memory from address 1\n")
	//port.Write([]byte("\rCP 1\r^pFD\r "))
	//time.Sleep(2 * time.Second)

	//fmt.Printf("Attempting Factory reset\n")
	//port.Write([]byte("\rFD\r "))
	//time.Sleep(1 * time.Second)
	//
	//fmt.Printf("Stopping all movement\n")
	//port.Write([]byte("CK=0\r "))
	//time.Sleep(5 * time.Second)

	//^d1000
	//^J*FD^J
	//fmt.Printf("Attempting Factory reset with *\n")
	//port.Write([]byte("\r*FD\r "))
	//time.Sleep(1 * time.Second)
	//
	//^J*FDÌ^J
	//fmt.Printf("Attempting Factory reset with Ì\n")
	//port.Write([]byte("\r*FDÌ\r "))

	//time.Sleep(1 * time.Second)
	//	port.Write([]byte("S\n"))
	//	// S\n
	//	time.Sleep(1 * time.Second)
	//	// S\n
	//	port.Write([]byte("FD\n"))

	//	time.Sleep(1 * time.Second)
	// S\n
	//	port.Write([]byte("S\r ^d2000\r "))
	port.Write([]byte("PR AL^M "))
	//	port.Write([]byte("S\r "))

	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Printf("Sent bytes\n")

	buff := make([]byte, 10000)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			//break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		fmt.Printf("%v", string(buff[:n]))
	}
}
