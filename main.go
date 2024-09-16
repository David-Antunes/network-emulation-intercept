package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/David-Antunes/network-emulation-proxy/xdp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("missing socket id")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments")
		return
	}
	conn, err := net.Dial("unix", os.Args[1])

	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(conn)
  enc := gob.NewEncoder(conn)
	for {
		var frame *xdp.Frame

		err := dec.Decode(&frame)
		if err != nil {
			panic(err)
		}

    err := enc.Encode(&frame)
    if err != nil {
      panic(err)
    }
	}

}
