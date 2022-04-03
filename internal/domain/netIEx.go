package domain

type NetInterfaceExtended struct {
	Name                 string
	MulticastPkts        uint
	BroadcastPkts        uint
	OutMulticastPkts     uint
	OutBroadcastPkts     uint
	HCInOctets           uint64
	HCInUcastPkts        uint64
	HCInMulticastPkts    uint64
	HCInBroadcastPkts    uint64
	HCOutOctets          uint64
	HCOutUcastPkts       uint64
	HCOutMulticastPkts   uint64
	HCOutBroadcastPkts   uint64
	LinkUpDownTrapEnable int
	HighSpeed            uint
	PromiscuousMode      int
	ConnectorPresent     int
	Alias                string
}

func NewNetInterfaceExtended() NetInterfaceExtended {
	return NetInterfaceExtended{}
}
