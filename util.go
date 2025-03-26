package main

import "os"

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func loadCerts(clientCertPath, clientKeyPath, serverCertPath string) (string, string, string, error) {
	clientCert, err := readFile(clientCertPath)
	if err != nil {
		return "", "", "", err
	}
	clientKey, err := readFile(clientKeyPath)
	if err != nil {
		return "", "", "", err
	}
	serverCert, err := readFile(serverCertPath)
	if err != nil {
		return "", "", "", err
	}
	return clientCert, clientKey, serverCert, nil
}
