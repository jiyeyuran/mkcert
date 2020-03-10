package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goodhosts/hostsfile"
)

func installHosts() (err error) {
	customPath := "./hosts"
	if _, err = os.Stat(customPath); err != nil {
		return nil
	}

	customHosts, err := hostsfile.NewCustomHosts(customPath)
	if err != nil {
		return
	}

	osHostsFilePath := os.ExpandEnv(filepath.FromSlash(hostsfile.HostsFilePath))
	if _, err = os.Stat(osHostsFilePath); err != nil {
		file, errInner := os.Create(osHostsFilePath)
		if errInner != nil {
			err = fmt.Errorf("请以管理员权限运行该程序: %s", errInner)
			return
		}
		file.Close()
	}

	sysHosts, err := hostsfile.NewHosts()
	if err != nil {
		log.Println(err)
		return
	}

	if !sysHosts.IsWritable() {
		err = fmt.Errorf("请以管理员权限运行该程序")
		return
	}

	for _, line := range customHosts.Lines {
		if err = sysHosts.Add(line.IP, line.Hosts...); err != nil {
			err = fmt.Errorf("行写入失败: %s", line.Raw)
			return
		}
	}

	if err = sysHosts.Flush(); err != nil {
		err = fmt.Errorf("hosts写入失败：%s", err)
	}

	return
}
