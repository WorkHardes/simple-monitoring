package keenetic

import (
	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/domain"
)

var OCounter = 0
var NetICounter = -1
var CountNetI = -1
var OCheck = false
var ICounter = 0
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
	if OCounter == 0 {
		CountNetI = dataUnit.Value.(int)
		ch <- CountNetI
		OCounter++

		for i := 0; i < CountNetI; i++ {
			netInterface := domain.NewNetInterface()
			NetInterfaces = append(NetInterfaces, netInterface)
		}
		return
	}

	NetICounter++
	if NetICounter != CountNetI && !OCheck {
		return
	}
	OCheck = true

	if NetICounter == CountNetI {
		NetICounter = 0
		ICounter++
	}

	switch ICounter {
	case 1:
		{
			Name := convert(dataUnit).(string)
			NetInterfaces[NetICounter].Name = Name
		}
	case 2:
		{
			Type := dataUnit.Value.(int)
			NetInterfaces[NetICounter].Type = Type
		}
	case 3:
		{
			Mtu := convert(dataUnit).(int)
			NetInterfaces[NetICounter].Mtu = Mtu
		}
	case 4:
		{
			Speed := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].Speed = Speed
		}
	case 5:
		{
			PhysAddress := convert(dataUnit).(string)
			NetInterfaces[NetICounter].PhysAddress = PhysAddress
		}
	case 6:
		{
			AdminStatus := convert(dataUnit).(int)
			NetInterfaces[NetICounter].AdminStatus = AdminStatus
		}
	case 7:
		{
			OperStatus := convert(dataUnit).(int)
			NetInterfaces[NetICounter].OperStatus = OperStatus
		}
	case 8:
		{
			LastChange := convert(dataUnit).(uint32)
			NetInterfaces[NetICounter].LastChange = LastChange
		}
	case 9:
		{
			InOctets := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].InOctets = InOctets
		}
	case 10:
		{
			InUcastPkts := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].InUcastPkts = InUcastPkts
		}
	case 11:
		{
			InDiscard := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].InDiscard = InDiscard
		}
	case 12:
		{
			InErrors := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].InErrors = InErrors
		}
	case 13:
		{
			OutOctets := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].OutOctets = OutOctets
		}
	case 14:
		{
			OutUcastPkts := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].OutUcastPkts = OutUcastPkts
		}
	case 15:
		{
			OutDiscards := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].OutDiscards = OutDiscards
		}
	case 16:
		{
			OutErrors := convert(dataUnit).(uint)
			NetInterfaces[NetICounter].OutErrors = OutErrors
		}
	}

}
