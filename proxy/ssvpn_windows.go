package proxy

import "net"

func vpnDial(address string) (net.Conn, error) {
	panic("not on Windows")
}

func sendTrafficStats(stat *trafficSurvey) error {
	return nil
}
