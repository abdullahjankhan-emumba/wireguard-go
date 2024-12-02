//go:build !linux

package device

import (
	"github.com/abdullahjankhan-emumba/wireguard-go/conn"
	"github.com/abdullahjankhan-emumba/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
