package main

type Server struct {
	ID            int    `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	SSHPassword   string `gorm:"column:ssh_password"`
	SSHPort       int    `gorm:"column:ssh_port"`
	SSHUsername   string `gorm:"column:ssh_username"`
	AdminDomain   string `gorm:"column:admin_domain"`
	AdminPassword string `gorm:"column:admin_password"`
	IPv4          string `gorm:"column:ipv4"`
	IPv6          string `gorm:"column:ipv6"`
}

func (Server) TableName() string {
	return "core.servers"
}
