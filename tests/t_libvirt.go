package main

import (
	"github.com/rgbkrk/libvirt-go"
	"log"
)

func main() {
	var domainname string
	domainname = "test33"

	conn, err := libvirt.NewVirConnection("qemu:///system")
	if err != nil {
		log.Println("conn err", err)
	}
	domain, err := conn.LookupDomainByName(domainname)
	if err != nil {
		log.Println("init domain err", err)
	}

	// //update cpu
	// err = domain.SetVcpusFlags(2, 2)
	// log.Println("update vcp err", err)
	err = domain.AttachDevice(`<graphics type='vnc' port='-1' autoport='yes' listen='0.0.0.0' passwd='redhat'><listen type='address' address='0.0.0.0'/></graphics>`)
	log.Println("eeeee--->", err)
}

// func InitDomain(conn *libvirt.VirConnection, domainname string) (*libvirt.VirDomain, error) {
// 	d, err := conn.LookupDomainByName(domainname)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return d, nil
// }

//add disk to the config file ,2 means affect config file,after attach ,must restart domain
// func AttachDevice(domain *libvirt.VirDomain, imgpath string, devname string) error {
// 	xml := "<disk type='file' device='disk'><driver name='qemu' type='qcow2'/> <source file='" + imgpath + "'/><target dev='" + devname + "' bus='virtio'/></disk>"
// 	err = domain.AttachDeviceFlags(xml, 2)
// 	return err
// }

// //update domain cpu count of config file
// func UpdateDomainVcpu(domain *libvirt.VirDomain, vcpus uint) error {
// 	err = domain.SetVcpusFlags(vcpus, 4)
// 	return err
// }
