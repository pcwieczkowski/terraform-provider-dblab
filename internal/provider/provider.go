package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &DBLabProvider{}

// DBLabProvider defines the provider implementation.
type DBLabProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// DBLabProviderModel describes the provider data model.
type DBLabProviderModel struct {
	Endpoint          types.String `tfsdk:"endpoint"`
	VerificationToken types.String `tfsdk:"verification_token"`
}

func (p *DBLabProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "dblab"
	resp.Version = p.version
}

func (p *DBLabProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The full URL of the DBLab API server.",
				Required:            true,
			},
			"verification_token": schema.StringAttribute{
				MarkdownDescription: "Verification token for the DBLab API.",
				Required:            true,
			},
		},
	}
}

func (p *DBLabProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	data := &DBLabProviderModel{}

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.DataSourceData = data
	resp.ResourceData = data
}

func (p *DBLabProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCloneResource,
	}
}

func (p *DBLabProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSnapshotsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &DBLabProvider{
			version: version,
		}
	}
}
