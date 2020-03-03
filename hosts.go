package main

import (
	"log"
	"os"

	"github.com/goodhosts/hostsfile"
)

func init() {
	customPath := "./hosts"
	if _, err := os.Stat(customPath); err != nil {
		return
	}

	customHosts, err := hostsfile.NewCustomHosts(customPath)
	if err != nil {
		log.Println(err)
		return
	}

	sysHosts, err := hostsfile.NewHosts()
	if err != nil {
		log.Println(err)
		return
	}

	for _, line := range customHosts.Lines {
		if err = sysHosts.Add(line.IP, line.Hosts...); err != nil {
			log.Printf("failed to add line: %s", line.Raw)
		}
	}

	if err = sysHosts.Flush(); err != nil {
		log.Println(err)
	}
}
