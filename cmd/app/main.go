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
	"github.com/simple-monitoring/internal/service/keenetic"
	"github.com/simple-monitoring/pkg/database/influxdb"
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

func writeSystemInfoToDB(writeAPI influxdb2Api.WriteAPIBlocking) {
	metricsFields := map[string]string{"unit": "system-info"}
	metricsValues := map[string]interface{}{"sys-descr": keenetic.SystemInfo.SysDescr, "sys-up-time": keenetic.SystemInfo.SysUpTime.ToString(), "sys-name": keenetic.SystemInfo.SysName}
	point := influxdb.GetPoint(metricsFields, metricsValues)
	influxdb.WriteInfoToDB(point, writeAPI)
}

func writeNetIInfoToDB(writeAPI influxdb2Api.WriteAPIBlocking) {
}

func writeNetIExInfoToDB(writeAPI influxdb2Api.WriteAPIBlocking) {
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

			time.Sleep(5 * time.Second)
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
