//go:build !linux

package device

import (
	"github.com/abdullahjan-emumba/wireguard/conn"
	"github.com/abdullahjan-emumba/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
