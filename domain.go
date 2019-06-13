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
		Name: "Eddie4",
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
			Emulator: "/usr/bin/kvm-spice",
			Disks: []libvirtxml.DomainDisk{
				libvirtxml.DomainDisk{
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: "/home/whitebyte/iso/eddie4.iso",
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "sda",
					},
				},
			},
			Serials: []libvirtxml.DomainSerial{
				libvirtxml.DomainSerial{
					Protocol: &libvirtxml.DomainChardevProtocol{
						Type: "serial",
					},
					Target: &libvirtxml.DomainSerialTarget{
						Port: new(uint),
					},
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				libvirtxml.DomainConsole{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "serial",
						Port: new(uint),
					},
				},
				libvirtxml.DomainConsole{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "virtio",
					},
				},
			},
			Graphics: []libvirtxml.DomainGraphic{
				libvirtxml.DomainGraphic{
					Spice: &libvirtxml.DomainGraphicSpice{
						AutoPort: "yes",
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
