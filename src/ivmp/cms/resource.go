package resource

type Resource struct {
	ID   string
	Name string
}

type Platform struct {
	IP   string
	Port ushort
}

type Device struct {
	Resource
	ParentID []string
	NodeID   string
	IP       string
	Port     ushort
	UserName string
	Password ushort
	Address  string
}

type Camera struct {
	Device
	PTZType      int
	Manufacturer string
	Channel      int
}

type NVR struct {
	Device
}

type BusinessGroup struct {
	Resource
}

type VirtualOrganization struct {
	Resource
	ParentID string
}

type Node struct {
	Resource
	Type int
	IP   string
	Port ushort
}

type User struct {
	Resource
	Type     int
	Purview  int
	Password string
}

type Resourcer interface {
	Control()
}
