package domain

type NetInterface struct {
	Name         string
	Type         int
	Mtu          int
	Speed        uint
	PhysAddress  string
	AdminStatus  int
	OperStatus   int
	LastChange   uint32
	InOctets     uint
	InUcastPkts  uint
	InDiscard    uint
	InErrors     uint
	OutOctets    uint
	OutUcastPkts uint
	OutDiscards  uint
	OutErrors    uint
}

func NewNetInterface() NetInterface {
	return NetInterface{}
}
