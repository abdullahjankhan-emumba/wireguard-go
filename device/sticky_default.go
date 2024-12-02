//go:build !linux

package device

import (
	"github.com/abdullahjankhan-emumba/wireguard/conn"
	"github.com/abdullahjankhan-emumba/wireguard/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
