package ip

import (
	"fmt"
	"github.com/ipipdotnet/ipdb-go"
	"log"
	"os"
	"path"
)

var (
	IP       *ipdb.BaseStation
	Language = "CN"
)

func init() {
	var err error
	// 项目根目录
	rootPath, _ := os.Getwd()
	IP, err = ipdb.NewBaseStation(path.Join(rootPath, "framework/ip/ipipfree.ipdb"))
	if err != nil {
		log.Panicf("Load ipdb error: %v", err)
	}
}

// FindIp 获取ID地址
func FindIp(ip string) string {
	if len(ip) == 0 {
		return ""
	}
	baseStationInfo, err := IP.FindInfo(ip, Language)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", baseStationInfo.CountryName, baseStationInfo.RegionName, baseStationInfo.CityName)
}
