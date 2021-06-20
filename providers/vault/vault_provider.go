package vault

import (
	"errors"
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
	"os"
	"strings"
)

type VaultProvider struct {
	terraformutils.Provider
	token   string
	address string
}

func (p *VaultProvider) Init(args []string) error {

	if address := os.Getenv("VAULT_ADDR"); address != "" {
		p.address = os.Getenv("VAULT_ADDR")
	}

	if token := os.Getenv("VAULT_TOKEN"); token != "" {
		p.token = os.Getenv("VAULT_TOKEN")
	}

	if len(args) > 0 && args[0] != "" {
		p.address = args[0]
	}

	if len(args) > 1 && args[1] != "" {
		p.token = args[1]
	}

	return nil
}

func (p *VaultProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"token":   cty.StringVal(p.token),
		"address": cty.StringVal(p.address),
	})
}

func (p *VaultProvider) GetName() string {
	return "vault"
}

func (p *VaultProvider) InitService(serviceName string, verbose bool) error {
	if service, ok := p.GetSupportedService()[serviceName]; ok {
		p.Service = service
		p.Service.SetName(serviceName)
		p.Service.SetVerbose(verbose)
		p.Service.SetProviderName(p.GetName())
		p.Service.SetArgs(map[string]interface{}{
			"token":   p.token,
			"address": p.address,
		})
		if err := service.(*ServiceGenerator).setVaultClient(); err != nil {
			return err
		}
		return nil
	}
	return errors.New(p.GetName() + ": " + serviceName + " not supported service")
}

func getSupportedEngineServices() []string {
	var services []string
	mapping := map[string][]string{
		"secret_backend": {"ad", "aws", "azure", "consul", "gcp", "nomad", "pki", "rabbitmq", "terraform_cloud"},
		"secret_backend_role": {"ad", "aws", "azure", "consul", "database", "pki", "rabbitmq", "ssh"},
		"auth_backend": {"gcp", "github", "jwt", "ldap", "okta"},
		"auth_backend_role": {"alicloud", "approle", "aws", "azure", "cert", "gcp", "jwt", "kubernetes", "token"},
		"auth_backend_user": {"ldap", "okta"},
		"auth_backend_group": {"ldap", "okta"},
	}
	for resource, mountTypes := range mapping {
		for _, mountType := range mountTypes {
			services = append(services, fmt.Sprintf("%s_%s", mountType, resource))
		}
	}
	return services
}

func (p *VaultProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	generators := make(map[string]terraformutils.ServiceGenerator)
	for _, service := range getSupportedEngineServices() {
		split := strings.SplitN(service, "_", 2)
		generators[service] = &ServiceGenerator{mountType: split[0], resource: split[1]}
	}
	generators["policy"] = &ServiceGenerator{resource: "policy"}
	return generators
}

func (VaultProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (VaultProvider) GetProviderData(_ ...string) map[string]interface{} {
	return map[string]interface{}{}
}
