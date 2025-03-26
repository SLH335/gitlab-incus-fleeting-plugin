package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"gitlab.com/gitlab-org/fleeting/fleeting/provider"
)

var _ provider.InstanceGroup = &incusDeployment{}

type incusDeployment struct {
}

func (i *incusDeployment) Init(ctx context.Context, logger hclog.Logger, settings provider.Settings) (provider.ProviderInfo, error) {
	return provider.ProviderInfo{}, nil
}

func (i *incusDeployment) Update(ctx context.Context, fn func(instance string, state provider.State)) error {
	return nil
}

func (i *incusDeployment) Increase(ctx context.Context, n int) (int, error) {
	return n, nil
}

func (i *incusDeployment) Decrease(ctx context.Context, instances []string) ([]string, error) {
	return nil, nil
}

func (i *incusDeployment) ConnectInfo(ctx context.Context, instance string) (provider.ConnectInfo, error) {
	return provider.ConnectInfo{}, nil
}

func (i *incusDeployment) Shutdown(ctx context.Context) error {
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide a container name")
		return
	}
	name := strings.ToLower(os.Args[1])

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		return
	}

	c, err := Connect()
	if err != nil {
		fmt.Printf("Error connecting to Incus: %v\n", err)
		return
	}
	fmt.Println("Connected to Incus")

	err = CreateContainer(c, name)
	if err != nil {
		fmt.Printf("Error creating container %s: %v\n", name, err)
		return
	}
	fmt.Printf("Created container %s\n", name)

	err = StartContainer(c, name)
	if err != nil {
		fmt.Printf("Error starting container %s: %v\n", name, err)
		return
	}
	fmt.Printf("Started container %s\n", name)

	//plugin.Serve(&incusDeployment{})
}
