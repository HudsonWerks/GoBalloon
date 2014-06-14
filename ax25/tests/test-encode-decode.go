package main

import (
	"bytes"
	"fmt"
	"github.com/chrissnell/GoBalloon/ax25"
	"log"
	//	"os"
	//	"os/signal"
	// "github.com/chrissnell/go-base91"
)

func main() {

	psource := ax25.APRSAddress{
		Callsign: "NW5W",
		SSID:     7,
	}

	pdest := ax25.APRSAddress{
		Callsign: "APZ001",
		SSID:     0,
	}

	path1 := ax25.APRSAddress{
		Callsign: "WIDE1",
		SSID:     1,
	}

	path2 := ax25.APRSAddress{
		Callsign: "WIDE2",
		SSID:     1,
	}

	a := ax25.APRSData{
		Source: psource,
		Dest:   pdest,
		Path:   []ax25.APRSAddress{path1, path2},
		Body:   "!4715.68N/12228.20W-GoBalloon Test http://nw5w.com",
	}

	packet, err := ax25.EncodeAX25Command(a)
	if err != nil {
		log.Fatalf("Unable to create packet: %v", err)
	} else {
		fmt.Printf("--> %v\n", string(packet))
	}

	fmt.Println("Byte#\tHexVal\tChar\tChar>>1\tBinary")
	fmt.Println("-----\t------\t----\t-------\t------")
	for k, v := range packet {
		rs := v >> 1
		fmt.Printf("%4d \t%#x \t%v \t%v\t%08b\n", k, v, string(v), string(rs), v)
	}

	//var buf bytes.Buffer
	buf := bytes.NewReader(packet)
	d := ax25.NewDecoder(buf)

	msg, err := d.Next()

	fmt.Printf("%+v\n", msg)

}
