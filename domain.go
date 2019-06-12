package main

import (
	"fmt"
	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func listAllDomains() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}

	domains, _ := conn.ListAllDomains(1)
	// check for libvirt domain flags ConnectListAllDomainsFlags
	fmt.Printf("Domains: %s\n", domains)

	conn.Close()
}

func domainDef() {
	domcfg := &libvirtxml.Domain{}
	xmldoc, err := domcfg.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Printf("XML: %s\n", xmldoc)

	domcfg = &libvirtxml.Domain{
		Type: "qemu",
		Name: "Eddie1",
		Memory: &libvirtxml.DomainMemory{
			Value: 4,
			Unit:  "GB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: 2,
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Type: "hvm",
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Disks: []libvirtxml.DomainDisk{
				libvirtxml.DomainDisk{
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: "/home/whitebyte/iso/CentOS-7-x86_64-Minimal-1810.iso",
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda",
					},
				},
			},
		},
	}
	xmldoc, err = domcfg.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Printf("XML: %s", xmldoc)

	conn, _ := libvirt.NewConnect("qemu:///system")
	_, err = conn.DomainCreateXML(xmldoc, 0)
	if err != nil {
		panic(err)
	}

}
