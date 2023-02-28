package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"gitlab.com/postgres-ai/database-lab/v3/pkg/client/dblabapi"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SnapshotsDataSource{}

func NewSnapshotsDataSource() datasource.DataSource {
	return &SnapshotsDataSource{}
}

// SnapshotsDataSource defines the data source implementation.
type SnapshotsDataSource struct {
	dblabClient *dblabapi.Client
}

type DBLabSnapshot struct {
	Id   types.String `tfsdk:"id"`
	Pool types.String `tfsdk:"pool"`
}

// SnapshotsDataSourceModel describes the data source data model.
type SnapshotsDataSourceModel struct {
	Snapshots []DBLabSnapshot `tfsdk:"snapshots"`
}

func (d *SnapshotsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snapshots"
}

func (d *SnapshotsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Snapshots data source",

		Attributes: map[string]schema.Attribute{
			"snapshots": schema.ListAttribute{
				MarkdownDescription: "Returned snapshots",
				Computed:            true,

				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"id":   types.StringType,
						"pool": types.StringType,
					},
				},
			},
		},
	}
}

func (d *SnapshotsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	data, ok := req.ProviderData.(*DBLabProviderModel)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *DBLabProviderModel, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	client, err := dblabapi.NewClient(dblabapi.Options{
		Host:              data.Endpoint.ValueString(),
		VerificationToken: data.VerificationToken.ValueString(),
	})
	if err != nil {
		panic("x")
	}
	// client, ok := req.ProviderData.(*http.Client)

	// _, err := client.Do(snapshotsRequest)
	// if err != nil {
	// 	panic("snapshots request failed")
	// }

	// if !ok {
	// 	resp.Diagnostics.AddError(
	// 		"Unexpected Data Source Configure Type",
	// 		fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
	// 	)

	// 	return
	// }

	d.dblabClient = client
}

func (d *SnapshotsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	data := SnapshotsDataSourceModel{
		Snapshots: []DBLabSnapshot{},
	}

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	snapshots, err := d.dblabClient.ListSnapshots(ctx)
	if err != nil {
		panic("x")
	}
	for _, snapshot := range snapshots {
		data.Snapshots = append(data.Snapshots, DBLabSnapshot{
			Id:   types.StringValue(snapshot.ID),
			Pool: types.StringValue(snapshot.Pool),
		})
	}
	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := d.client.Do(httpReq)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
