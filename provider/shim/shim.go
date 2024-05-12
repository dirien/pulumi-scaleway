package shim

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	prov "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ prov.Provider = &ScalewayProvider{}

type ScalewayProvider struct{}

func (a ScalewayProvider) Metadata(ctx context.Context, request prov.MetadataRequest, response *prov.MetadataResponse) {
}

func (a ScalewayProvider) Schema(ctx context.Context, request prov.SchemaRequest, response *prov.SchemaResponse) {

}

func (a ScalewayProvider) Configure(ctx context.Context, request prov.ConfigureRequest, response *prov.ConfigureResponse) {

}

func (a ScalewayProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (a ScalewayProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func Framework() func() prov.Provider {
	return func() prov.Provider {
		return &ScalewayProvider{}
	}
}
