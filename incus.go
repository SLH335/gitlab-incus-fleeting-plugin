package main

import (
	"os"

	"github.com/lxc/incus/client"
	"github.com/lxc/incus/shared/api"
)

func Connect() (incus.InstanceServer, error) {
	url := os.Getenv("INCUS_URL")

	clientCertPath := os.Getenv("CLIENT_CERT_PATH")
	clientKeyPath := os.Getenv("CLIENT_KEY_PATH")
	serverCertPath := os.Getenv("SERVER_CERT_PATH")

	clientCert, clientKey, serverCert, err := loadCerts(clientCertPath, clientKeyPath, serverCertPath)
	if err != nil {
		return nil, err
	}

	args := incus.ConnectionArgs{
		TLSClientCert: clientCert,
		TLSClientKey:  clientKey,
		TLSServerCert: serverCert,
	}

	c, err := incus.ConnectIncus(url, &args)
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

func DeleteContainer(c incus.InstanceServer, name string) error {
	op, err := c.DeleteInstance(name)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}

func StopContainer(c incus.InstanceServer, name string) error {
	reqState := api.InstanceStatePut{
		Action:  "stop",
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
