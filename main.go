package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	buffer_size int = 102
	buffer      [102]byte
	inAddr      string = "0.0.0.0:7"
	outAddr     string = "192.168.0.255:9"
)

func init() {
	flag.StringVar(&inAddr, "listen", inAddr, "listen address")
	flag.StringVar(&outAddr, "send", outAddr, "send address")
	flag.Parse()
}

func main() {
	inAddress, err := net.ResolveUDPAddr("IP4", inAddr)
	catch("resolving listening address", err, false)

	outAddress, err := net.ResolveUDPAddr("IP4", outAddr)
	catch("resolving send address:", err, false)

	socketIn, err := net.ListenUDP("udp", inAddress)
	catch("listening:", err, false)

	for {
		read, err := socketIn.Read(buffer[0:buffer_size])
		catch("reading:", err, true)

		if err == nil && read > 0 {
			if checkWOLPacket(buffer) {
				socketOut, err := net.DialUDP("udp4", nil, outAddress)
				catch("opening socket for writing:", err, true)

				if err == nil {
					written, err := socketOut.Write(buffer[0:buffer_size])
					catch("writing to socket:", err, true)

					if err == nil && written > 0 {
						fmt.Println(written, "bytes written")
					}
				}
			}
		}
	}
}

func catch(description string, err error, relax bool) {
	if err == nil {
		return
	}

	fmt.Println("Error", description, err)

	if !relax {
		os.Exit(1)
	}
}

func checkWOLPacket(packet [102]byte) bool {
	header := []byte{255, 255, 255, 255, 255, 255}
	mac := packet[6:12]
	macs := []byte{}

	if bytes.Equal(packet[0:6], header) {
		for i := 1; i <= 16; i++ {
			macs = append(macs, mac...)
		}
		return bytes.Equal(packet[6:102], macs[0:96])
	}

	return false
}
