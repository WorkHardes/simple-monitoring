package influxdb

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb2Api "github.com/influxdata/influxdb-client-go/v2/api"
	influxdb2Write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

type InfluxDBSettings struct {
	Bucket string
	Org    string
	Token  string
	Url    string
}

func NewInfluxDBSettings() *InfluxDBSettings {
	return &InfluxDBSettings{}
}

func GetInfluxDBSettings() *InfluxDBSettings {
	idbs := NewInfluxDBSettings()
	idbs.Bucket = "keenetic"
	idbs.Org = "keenetic"
	idbs.Token = "admin_token"
	idbs.Url = "http://localhost:8086"
	return idbs
}

func GetClient(url, token string) influxdb2.Client {
	return influxdb2.NewClient(url, token)
}

func GetPoint(metricsFields map[string]string, metricsValues map[string]interface{}) *influxdb2Write.Point {
	pointName := "keenetic-router"
	return influxdb2.NewPoint(pointName, metricsFields, metricsValues, time.Now())
}

func WriteInfoToDB(point *influxdb2Write.Point, writeAPI influxdb2Api.WriteAPIBlocking) {
	writeAPI.WritePoint(context.Background(), point)
}
