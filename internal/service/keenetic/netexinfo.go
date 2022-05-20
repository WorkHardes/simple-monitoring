package keenetic

import (
	g "github.com/gosnmp/gosnmp"
	"github.com/simple-monitoring/internal/domain"
)

var OoCounter = 0
var NetIfacesCounter = -1
var countNetIfaces int
var IiCounter = 1
var NetIExtended = []domain.NetInterfaceExtended{}

func GetNetIExHandler(dataUnit g.SnmpPDU, ch chan int) {
	if OoCounter == 0 {
		countNetIfaces = <-ch
		OoCounter++

		for i := 0; i < countNetIfaces; i++ {
			netInterface := domain.NewNetInterfaceExtended()
			NetIExtended = append(NetIExtended, netInterface)
		}
	}

	NetIfacesCounter++
	if NetIfacesCounter == countNetIfaces {
		NetIfacesCounter = 0
		IiCounter++
	}

	switch IiCounter {
	case 1:
		{
			Name := convert(dataUnit).(string)
			NetIExtended[NetIfacesCounter].Name = Name
		}
	case 2:
		{
			MulticastPkts := dataUnit.Value.(uint)
			NetIExtended[NetIfacesCounter].MulticastPkts = MulticastPkts
		}
	case 3:
		{
			BroadcastPkts := convert(dataUnit).(uint)
			NetIExtended[NetIfacesCounter].BroadcastPkts = BroadcastPkts
		}
	case 4:
		{
			OutMulticastPkts := convert(dataUnit).(uint)
			NetIExtended[NetIfacesCounter].OutMulticastPkts = OutMulticastPkts
		}
	case 5:
		{
			OutBroadcastPkts := convert(dataUnit).(uint)
			NetIExtended[NetIfacesCounter].OutBroadcastPkts = OutBroadcastPkts
		}
	case 6:
		{
			HCInOctets := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCInOctets = HCInOctets
		}
	case 7:
		{
			HCInUcastPkts := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCInUcastPkts = HCInUcastPkts
		}
	case 8:
		{
			HCInMulticastPkts := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCInMulticastPkts = HCInMulticastPkts
		}
	case 9:
		{
			HCInBroadcastPkts := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCInBroadcastPkts = HCInBroadcastPkts
		}
	case 10:
		{
			HCOutOctets := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCOutOctets = HCOutOctets
		}
	case 11:
		{
			HCOutUcastPkts := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCOutUcastPkts = HCOutUcastPkts
		}
	case 12:
		{
			HCOutMulticastPkts := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCOutMulticastPkts = HCOutMulticastPkts
		}
	case 13:
		{
			HCOutBroadcastPkts := convert(dataUnit).(uint64)
			NetIExtended[NetIfacesCounter].HCOutBroadcastPkts = HCOutBroadcastPkts
		}
	case 14:
		{
			LinkUpDownTrapEnable := convert(dataUnit).(int)
			NetIExtended[NetIfacesCounter].LinkUpDownTrapEnable = LinkUpDownTrapEnable
		}
	case 15:
		{
			HighSpeed := convert(dataUnit).(uint)
			NetIExtended[NetIfacesCounter].HighSpeed = HighSpeed

		}
	case 16:
		{
			PromiscuousMode := convert(dataUnit).(int)
			NetIExtended[NetIfacesCounter].PromiscuousMode = PromiscuousMode
		}
	case 17:
		{
			ConnectorPresent := convert(dataUnit).(int)
			NetIExtended[NetIfacesCounter].ConnectorPresent = ConnectorPresent
		}

	case 18:
		{
			Alias := convert(dataUnit).(string)
			NetIExtended[NetIfacesCounter].Alias = Alias
		}
	}

}
