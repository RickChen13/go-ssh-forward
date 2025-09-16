package Const

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

var (
	BASE_PATH string
	WD_PATH   string
	Machineid string
)

func init() {
	getRoot()
	getWd()

}

func getRoot() {
	// 获取程序的运行目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return
	}
	BASE_PATH = dir
}

func getWd() {
	// 获取执行命令的目录
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	WD_PATH = dir

	if strings.Contains(os.Args[0], os.TempDir()) {
		BASE_PATH = WD_PATH
	}
}

func GetMachineid() (string, error) {
	if Machineid != "" {
		return Machineid, nil
	}
	id, err := machineid.ID()
	if err != nil {
		return "", err
	}
	Machineid = id
	return id, nil
}
