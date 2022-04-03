package keenetic

import (
	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/domain"
)

var ooCounter = 0
var netIfacesCounter = -1
var countNetIfaces int
var iiCounter = 1

var netIExtended = []domain.NetInterfaceExtended{}

func GetNetIExHandler(dataUnit g.SnmpPDU, ch chan int) {
	if ooCounter == 0 {
		countNetIfaces = <-ch
		ooCounter++

		for i := 0; i < countNetIfaces; i++ {
			netInterface := domain.NewNetInterfaceExtended()
			netIExtended = append(netIExtended, netInterface)
		}
	}

	netIfacesCounter++
	if netIfacesCounter == countNetIfaces {
		netIfacesCounter = 0
		iiCounter++
	}

	switch iiCounter {
	case 1:
		{
			Name := convert(dataUnit).(string)
			netIExtended[netIfacesCounter].Name = Name
		}
	case 2:
		{
			MulticastPkts := dataUnit.Value.(uint)
			netIExtended[netIfacesCounter].MulticastPkts = MulticastPkts
		}
	case 3:
		{
			BroadcastPkts := convert(dataUnit).(uint)
			netIExtended[netIfacesCounter].BroadcastPkts = BroadcastPkts
		}
	case 4:
		{
			OutMulticastPkts := convert(dataUnit).(uint)
			netIExtended[netIfacesCounter].OutMulticastPkts = OutMulticastPkts
		}
	case 5:
		{
			OutBroadcastPkts := convert(dataUnit).(uint)
			netIExtended[netIfacesCounter].OutBroadcastPkts = OutBroadcastPkts
		}
	case 6:
		{
			HCInOctets := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCInOctets = HCInOctets
		}
	case 7:
		{
			HCInUcastPkts := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCInUcastPkts = HCInUcastPkts
		}
	case 8:
		{
			HCInMulticastPkts := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCInMulticastPkts = HCInMulticastPkts
		}
	case 9:
		{
			HCInBroadcastPkts := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCInBroadcastPkts = HCInBroadcastPkts
		}
	case 10:
		{
			HCOutOctets := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCOutOctets = HCOutOctets
		}
	case 11:
		{
			HCOutUcastPkts := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCOutUcastPkts = HCOutUcastPkts
		}
	case 12:
		{
			HCOutMulticastPkts := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCOutMulticastPkts = HCOutMulticastPkts
		}
	case 13:
		{
			HCOutBroadcastPkts := convert(dataUnit).(uint64)
			netIExtended[netIfacesCounter].HCOutBroadcastPkts = HCOutBroadcastPkts
		}
	case 14:
		{
			LinkUpDownTrapEnable := convert(dataUnit).(int)
			netIExtended[netIfacesCounter].LinkUpDownTrapEnable = LinkUpDownTrapEnable
		}
	case 15:
		{
			HighSpeed := convert(dataUnit).(uint)
			netIExtended[netIfacesCounter].HighSpeed = HighSpeed

		}
	case 16:
		{
			PromiscuousMode := convert(dataUnit).(int)
			netIExtended[netIfacesCounter].PromiscuousMode = PromiscuousMode
		}
	case 17:
		{
			ConnectorPresent := convert(dataUnit).(int)
			netIExtended[netIfacesCounter].ConnectorPresent = ConnectorPresent
		}

	case 18:
		{
			Alias := convert(dataUnit).(string)
			netIExtended[netIfacesCounter].Alias = Alias
		}
	}

}
