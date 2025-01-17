package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Device implements api.Device
type Device struct {
	ApiClient ApiClient
}

func (d *Device) client(systemID string) ApiClient {
	return d.ApiClient.GetSubObject("devices").GetSubObject(fmt.Sprintf("%v", systemID))
}

// Get fetches a device with a given system_id
func (d *Device) Get(systemID string) (device *entity.Device, err error) {
	device = new(entity.Device)
	err = d.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, device)
	})
	return
}

// Update updates a given Device
func (d *Device) Update(systemID string, deviceParams *entity.DeviceUpdateParams) (device *entity.Device, err error) {
	qsp, err := query.Values(deviceParams)
	if err != nil {
		return
	}
	device = new(entity.Device)
	err = d.client(systemID).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, device)
	})
	return
}

// Delete deletes a given Device
func (d *Device) Delete(systemID string) error {
	return d.client(systemID).Delete()
}
