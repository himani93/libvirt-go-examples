package main

import (
	"fmt"
	libvirt "github.com/libvirt/libvirt-go"
)

func connect() {
	fmt.Println("connect")
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}
	fmt.Println("close")
	conn.Close()
}
