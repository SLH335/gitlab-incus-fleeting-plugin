package main

import (
	"github.com/lxc/incus/client"
	"github.com/lxc/incus/shared/api"
)

func Connect() (incus.InstanceServer, error) {
	c, err := incus.ConnectIncusUnix("", nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func CreateContainer(c incus.InstanceServer, name string) error {
	req := api.InstancesPost{
		Name: name,
		Source: api.InstanceSource{
			Alias:    "debian/bookworm/cloud",
			Mode:     "pull",
			Protocol: "simplestreams",
			Server:   "https://images.linuxcontainers.org",
			Type:     "image",
		},
		Type: "container",
	}

	op, err := c.CreateInstance(req)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func StartContainer(c incus.InstanceServer, name string) error {
	reqState := api.InstanceStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err := c.UpdateInstanceState(name, reqState, "")
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}
