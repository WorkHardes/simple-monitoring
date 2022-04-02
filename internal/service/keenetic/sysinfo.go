package keenetic

import (
	"fmt"

	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/domain"
)

var oidCounter = 0

func getSysDescr(dataUnit g.SnmpPDU) string {
	bytes := dataUnit.Value.([]byte)
	return string(bytes)
}

func getSysUpTime(dataUnit g.SnmpPDU) *domain.WorkTime {
	timeTicks := int(dataUnit.Value.(uint32))
	workSeconds := timeTicks / 100

	wt := domain.NewWorkTime()
	wt.Days = workSeconds / 86400
	wt.Hours = workSeconds / 3600
	wt.Minutes = workSeconds % 3600 / 60
	wt.Seconds = workSeconds % 216000 % 60

	return wt
}

func getSysName(dataUnit g.SnmpPDU) string {
	bytes := dataUnit.Value.([]byte)
	return string(bytes)
}

func SysInfoHandler(dataUnit g.SnmpPDU) {
	switch oidCounter {
	case 0:
		{
			sysDescr := getSysDescr(dataUnit)
			fmt.Println("sysDescr:", sysDescr)
		}
	case 1:
		{
			break
		}
	case 2:
		{
			sysUpTime := getSysUpTime(dataUnit)
			fmt.Println("sysUpTime:", sysUpTime)
		}
	case 3:
		{
			// sysContact
			break
		}
	case 4:
		{
			sysContact := getSysName(dataUnit)
			fmt.Println("sysName:", sysContact)
		}
	case 5:
		// sysLocation
		break
	case 6:
		// sysServices
		break
	}
	oidCounter++
}
