package ip

import (
	"fmt"
	"github.com/ipipdotnet/ipdb-go"
	"log"
)

var (
	IP       *ipdb.BaseStation
	Language = "CN"
)

func init() {
	var err error
	// 项目根目录
	IP, err = ipdb.NewBaseStation("./ipipfree.ipdb")
	if err != nil {
		log.Panicf("Load ipdb error: %v", err)
	}
}

// FindInfo 获取IP信息
func FindInfo(addr string, language string) (baseStationInfo *ipdb.BaseStationInfo, err error) {
	baseStationInfo, err = IP.FindInfo(addr, language)
	return
}

// FindIp 获取IP地址
func FindIp(addr string, language string) string {
	if len(addr) == 0 {
		return ""
	}
	baseStationInfo, err := IP.FindInfo(addr, language)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", baseStationInfo.CountryName, baseStationInfo.RegionName, baseStationInfo.CityName)
}
