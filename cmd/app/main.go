package main

import (
	"fmt"
	"log"
	"strings"

	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/service/keenetic"
)

func walkFn(dataUnit g.SnmpPDU) error {
	if strings.HasPrefix(dataUnit.Name, ".1.3.6.1.2.1.1") {
		keenetic.SysInfoHandler(dataUnit)
	} else if strings.HasPrefix(dataUnit.Name, ".1.3.6.1.2.1.2") {
		keenetic.NetInterfacesHandler(dataUnit)
	} else if strings.HasPrefix(dataUnit.Name, ".1.3.6.1.2.1.31") {
		keenetic.NetInterfacesExtendedHandler(dataUnit)
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
		fmt.Println("This device manufacturer is not supported yet.")
		return
	}

	g.Default.Target = routerIP
	if err := g.Default.Connect(); err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	rootOid := ""
	err := g.Default.BulkWalk(rootOid, walkFn)
	if err != nil {
		log.Fatalf("g.Default.BulkWalk() err: %v", err)
	}

}
