package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &SensitiveDataSource{}
	_ datasource.DataSourceWithConfigure = &SensitiveDataSource{}
)

func NewSensitiveDataSource() datasource.DataSource {
	return &SensitiveDataSource{}
}

// SensitiveDataSource defines the data source implementation.
type SensitiveDataSource struct{}

// SensitiveDataSourceModel describes the data source data model.
type SensitiveDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Required types.Bool   `tfsdk:"required"`
	Value    types.String `tfsdk:"value"`
}

func (d *SensitiveDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sensitive"
}

func (d *SensitiveDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Environment Variable data source marked as sensitive",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "env name",
				Required:            true,
			},
			"required": schema.BoolAttribute{
				MarkdownDescription: "require ENV to be set",
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "env value",
				Computed:            true,
				Sensitive:           true,
			},
		},
	}
}

func (d *SensitiveDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *SensitiveDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SensitiveDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	v, found := os.LookupEnv(data.Id.ValueString())
	if !found && data.Required.ValueBool() {
		errMsg := "env " + data.Id.String() + " not found"
		resp.Diagnostics.AddError(errMsg, errMsg+". Export that environment variable or set required=false on this data resource")
		return
	}
	data.Value = types.StringValue(v)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
