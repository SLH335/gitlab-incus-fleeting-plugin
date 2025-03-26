package main

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"gitlab.com/gitlab-org/fleeting/fleeting/plugin"
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
	plugin.Serve(&incusDeployment{})
}
