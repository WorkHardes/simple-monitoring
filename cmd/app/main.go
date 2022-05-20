package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	g "github.com/gosnmp/gosnmp"
	influxdb2Api "github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/simple-monitoring/internal/domain"
	"github.com/simple-monitoring/internal/service/keenetic"
	"github.com/simple-monitoring/pkg/database/influxdb"
	"github.com/simple-monitoring/pkg/logger"
)

var ch = make(chan int, 1)
var oldInPkts = 0
var oldOutPkts = 0
var checkPkts = false

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

func writeSystemInfoToDB(writeAPI influxdb2Api.WriteAPIBlocking) {
	metricsFields := map[string]string{"unit": "system-info"}
	metricsValues := map[string]interface{}{
		"sys-descr":   keenetic.SystemInfo.SysDescr,
		"sys-up-time": keenetic.SystemInfo.SysUpTime.ToString(),
		"sys-name":    keenetic.SystemInfo.SysName,
	}
	point := influxdb.GetPoint(metricsFields, metricsValues)
	err := influxdb.WriteInfoToDB(point, writeAPI)
	if err != nil {
		logger.Errorf("%w", err)
	}
}

func writeNetIInfoToDB(writeAPI influxdb2Api.WriteAPIBlocking) {
	metricsFields := map[string]string{"unit": "net-ifaces-info"}
	for i := 0; i < len(keenetic.NetInterfaces); i++ {
		metricsValues := map[string]interface{}{
			"netIface-info": keenetic.NetInterfaces[i].ToString(),
		}
		point := influxdb.GetNetIfacePoint(metricsFields, metricsValues)
		err := influxdb.WriteInfoToDB(point, writeAPI)
		if err != nil {
			logger.Errorf("%w", err)
		}

		if i == 1 {
			inPkts := keenetic.NetInterfaces[i].InUcastPkts - uint(oldInPkts)
			if !checkPkts {
				inPkts = 0
			}
			metricsValues = map[string]interface{}{
				"f0/0-in-pkts": inPkts,
			}
			point = influxdb.GetNetIfacePoint(metricsFields, metricsValues)
			err = influxdb.WriteInfoToDB(point, writeAPI)
			if err != nil {
				logger.Errorf("%w", err)
			}
			oldInPkts = int(keenetic.NetInterfaces[i].InUcastPkts)

			outPkts := keenetic.NetInterfaces[i].OutUcastPkts - uint(oldOutPkts)
			if !checkPkts {
				outPkts = 0
				checkPkts = true
			}
			metricsValues = map[string]interface{}{
				"f0/0-out-pkts": outPkts,
			}
			point = influxdb.GetNetIfacePoint(metricsFields, metricsValues)
			err = influxdb.WriteInfoToDB(point, writeAPI)
			if err != nil {
				logger.Errorf("%w", err)
			}
			oldOutPkts = int(keenetic.NetInterfaces[i].OutUcastPkts)
		}
	}
}

func writeNetIExInfoToDB(writeAPI influxdb2Api.WriteAPIBlocking) {
}

func changeVarsToDefault() {
	keenetic.OidCounter = 0

	keenetic.OCounter = 0
	keenetic.NetICounter = -1
	keenetic.CountNetI = 0
	keenetic.OCheck = false
	keenetic.ICounter = 0
	keenetic.NetInterfaces = []domain.NetInterface{}

	keenetic.OoCounter = 0
	keenetic.NetIfacesCounter = -1
	keenetic.IiCounter = 1
	keenetic.NetIExtended = []domain.NetInterfaceExtended{}
}

func main() {
	routerIP := "192.168.1.1"

	logger.Infof("SNMP target is %s", routerIP)

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

	logger.Infof("Connect to %s success", routerIP)

	rootOid := ""

	idbs := influxdb.GetInfluxDBSettings()
	client := influxdb.GetClient(idbs.Url, idbs.Token)
	writeAPI := client.WriteAPIBlocking(idbs.Org, idbs.Bucket)

	go func() {
		for {
			err := g.Default.BulkWalk(rootOid, walkFn)
			if err != nil {
				logger.Errorf("g.Default.BulkWalk() err: %v", err)
			}

			writeSystemInfoToDB(writeAPI)

			writeNetIInfoToDB(writeAPI)

			writeNetIExInfoToDB(writeAPI)

			changeVarsToDefault()

			time.Sleep(2 * time.Second)
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	_, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	logger.Info("Server stopped")

	client.Close()

}
