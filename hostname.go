package main

import (
	"fmt"
	libvirt "github.com/libvirt/libvirt-go"
)

func getHostname() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}

	host_name, _ := conn.GetHostname()

	fmt.Printf("Host Name: %s\n", host_name)

	conn.Close()
}
