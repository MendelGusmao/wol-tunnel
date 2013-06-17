package main

import "testing"

var (
	packet = [102]byte{
		255, 255, 255, 255, 255, 255,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
		19, 9, 87, 30, 7, 86,
	}
)

func TestCheckWOLPacket(t *testing.T) {
	_, ok := checkWOLPacket(packet)
	if !ok {
		t.Fatal()
	}
}

func TestCheckWOLPacket2(t *testing.T) {
	packet[30] = 89
	_, ok := checkWOLPacket(packet)
	if ok {
		t.Fatal()
	}
}
