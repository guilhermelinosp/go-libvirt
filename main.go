package main

import (
	"fmt"
	"log"

	"github.com/libvirt/libvirt-go"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("failed to connect to libvirt: %v", err)
	}
	defer conn.Close()

	domains, err := conn.ListAllDomains(libvirt.ConnectListAllDomainsFlags(0))
	if err != nil {
		log.Fatalf("failed to list domains: %v", err)
	}

	fmt.Println("Active Domains:")
	for _, domain := range domains {
		name, err := domain.GetName()
		if err != nil {
			log.Printf("failed to get domain name: %v", err)
			continue
		}

		id, _ := domain.GetID()
		state, _, err := domain.GetState()
		if err != nil {
			log.Printf("failed to get domain state: %v", err)
			continue
		}

		fmt.Printf("ID: %d, Name: %s, State: %s\n", id, name, domainStateToString(state))

		domain.Free()
	}
}

func domainStateToString(state libvirt.DomainState) string {
	switch state {
	case libvirt.DOMAIN_RUNNING:
		return "Running"
	case libvirt.DOMAIN_BLOCKED:
		return "Blocked"
	case libvirt.DOMAIN_PAUSED:
		return "Paused"
	case libvirt.DOMAIN_SHUTDOWN:
		return "Shutdown"
	case libvirt.DOMAIN_SHUTOFF:
		return "Shut off"
	case libvirt.DOMAIN_CRASHED:
		return "Crashed"
	default:
		return "Unknown"
	}
}
