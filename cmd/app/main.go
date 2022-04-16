package main

import (
	"strings"

	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/service/keenetic"
	"github.com/simple-monitoring/pkg/logger"
)

var ch = make(chan int, 1)

func walkFn(dataUnit g.SnmpPDU) error {
	if strings.HasPrefix(dataUnit.Name, ".1.3.6.1.2.1.1") {
		keenetic.SysInfoHandler(dataUnit)
	} else if strings.HasPrefix(dataUnit.Name, ".1.3.6.1.2.1.2") {
		keenetic.GetNetIfacesInfo(dataUnit, ch)
	} else if strings.HasPrefix(dataUnit.Name, ".1.3.6.1.2.1.31") {
		keenetic.GetNetIExHandler(dataUnit, ch)
	}
	return nil
}

func main() {
	routerIP := "192.168.1.1"
	routerManufacturer := "keenetic"
	switch strings.ToLower(routerManufacturer) {
	case "keenetic":
		break
	default:
		logger.Error("This device manufacturer is not supported yet")
		return
	}

	g.Default.Target = routerIP
	if err := g.Default.Connect(); err != nil {
		logger.Errorf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	rootOid := ""
	err := g.Default.BulkWalk(rootOid, walkFn)
	if err != nil {
		logger.Errorf("g.Default.BulkWalk() err: %v", err)
	}

}
