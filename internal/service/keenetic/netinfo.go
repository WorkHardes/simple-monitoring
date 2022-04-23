package keenetic

import (
	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/domain"
)

var oCounter = 0
var netICounter = -1
var countNetI = -1
var ocheck = false
var iCounter = 0
var NetInterfaces = []domain.NetInterface{}

func convert(dataUnit g.SnmpPDU) interface{} {
	if dataUnit.Type == g.OctetString {
		bytes := dataUnit.Value.([]byte)
		return string(bytes)
	} else {
		return dataUnit.Value
	}
}

func GetNetIfacesInfo(dataUnit g.SnmpPDU, ch chan int) {
	if oCounter == 0 {
		countNetI = dataUnit.Value.(int)
		ch <- countNetI
		oCounter++

		for i := 0; i < countNetI; i++ {
			netInterface := domain.NewNetInterface()
			NetInterfaces = append(NetInterfaces, netInterface)
		}
		return
	}

	netICounter++
	if netICounter != countNetI && !ocheck {
		return
	}
	ocheck = true

	if netICounter == countNetI {
		netICounter = 0
		iCounter++
	}

	switch iCounter {
	case 1:
		{
			Name := convert(dataUnit).(string)
			NetInterfaces[netICounter].Name = Name
		}
	case 2:
		{
			Type := dataUnit.Value.(int)
			NetInterfaces[netICounter].Type = Type
		}
	case 3:
		{
			Mtu := convert(dataUnit).(int)
			NetInterfaces[netICounter].Mtu = Mtu
		}
	case 4:
		{
			Speed := convert(dataUnit).(uint)
			NetInterfaces[netICounter].Speed = Speed
		}
	case 5:
		{
			PhysAddress := convert(dataUnit).(string)
			NetInterfaces[netICounter].PhysAddress = PhysAddress
		}
	case 6:
		{
			AdminStatus := convert(dataUnit).(int)
			NetInterfaces[netICounter].AdminStatus = AdminStatus
		}
	case 7:
		{
			OperStatus := convert(dataUnit).(int)
			NetInterfaces[netICounter].OperStatus = OperStatus
		}
	case 8:
		{
			LastChange := convert(dataUnit).(uint32)
			NetInterfaces[netICounter].LastChange = LastChange
		}
	case 9:
		{
			InOctets := convert(dataUnit).(uint)
			NetInterfaces[netICounter].InOctets = InOctets
		}
	case 10:
		{
			InUcastPkts := convert(dataUnit).(uint)
			NetInterfaces[netICounter].InUcastPkts = InUcastPkts
		}
	case 11:
		{
			InDiscard := convert(dataUnit).(uint)
			NetInterfaces[netICounter].InDiscard = InDiscard
		}
	case 12:
		{
			InErrors := convert(dataUnit).(uint)
			NetInterfaces[netICounter].InErrors = InErrors
		}
	case 13:
		{
			OutOctets := convert(dataUnit).(uint)
			NetInterfaces[netICounter].OutOctets = OutOctets
		}
	case 14:
		{
			OutUcastPkts := convert(dataUnit).(uint)
			NetInterfaces[netICounter].OutUcastPkts = OutUcastPkts
		}
	case 15:
		{
			OutDiscards := convert(dataUnit).(uint)
			NetInterfaces[netICounter].OutDiscards = OutDiscards
		}
	case 16:
		{
			OutErrors := convert(dataUnit).(uint)
			NetInterfaces[netICounter].OutErrors = OutErrors
		}
	}

}
