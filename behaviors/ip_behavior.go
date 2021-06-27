package behaviors

import (
	"log"
	"net"
)

type IpBehavior struct {
}

func (i *IpBehavior) DoCommand(command string) string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func (i *IpBehavior) Description() string {
	return "Shows IP"
}

func (i *IpBehavior) Command() string{
	return "ip"
}
