package keenetic

import (
	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/domain"
)

var oidCounter = 0
var SystemInfo = domain.SystemInfo{}

func getSysUpTime(dataUnit g.SnmpPDU) domain.WorkTime {
	timeTicks := int(dataUnit.Value.(uint32))
	workSeconds := timeTicks / 100

	wt := domain.NewWorkTime()
	wt.Days = workSeconds / 86400
	wt.Hours = workSeconds / 3600
	wt.Minutes = workSeconds % 3600 / 60
	wt.Seconds = workSeconds % 216000 % 60

	return wt
}

func SysInfoHandler(dataUnit g.SnmpPDU) {
	switch oidCounter {
	case 0:
		{
			sysDescr := convert(dataUnit).(string)
			SystemInfo.SysDescr = sysDescr
		}
	case 1:
		{
			break
		}
	case 2:
		{
			sysUpTime := getSysUpTime(dataUnit)
			SystemInfo.SysUpTime = sysUpTime
		}
	case 3:
		{
			// sysContact
			break
		}
	case 4:
		{
			sysName := convert(dataUnit).(string)
			SystemInfo.SysName = sysName
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
