package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure EnvProvider satisfies various provider interfaces.
var _ provider.Provider = &EnvProvider{}

// EnvProvider defines the provider implementation.
type EnvProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// EnvProviderModel describes the provider data model.
type EnvProviderModel struct{}

func (p *EnvProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "env"
	resp.Version = p.version
}

func (p *EnvProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *EnvProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data EnvProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
}

func (p *EnvProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *EnvProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewVarDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &EnvProvider{
			version: version,
		}
	}
}
