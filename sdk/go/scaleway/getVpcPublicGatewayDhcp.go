// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a public gateway DHCP.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "main", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet: pulumi.String("192.168.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupVpcPublicGatewayDhcpOutput(ctx, scaleway.GetVpcPublicGatewayDhcpOutputArgs{
//				DhcpId: main.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupVpcPublicGatewayDhcp(ctx *pulumi.Context, args *LookupVpcPublicGatewayDhcpArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayDhcpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPublicGatewayDhcpResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGatewayDhcp:getVpcPublicGatewayDhcp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGatewayDhcp.
type LookupVpcPublicGatewayDhcpArgs struct {
	DhcpId string `pulumi:"dhcpId"`
}

// A collection of values returned by getVpcPublicGatewayDhcp.
type LookupVpcPublicGatewayDhcpResult struct {
	Address             string   `pulumi:"address"`
	CreatedAt           string   `pulumi:"createdAt"`
	DhcpId              string   `pulumi:"dhcpId"`
	DnsLocalName        string   `pulumi:"dnsLocalName"`
	DnsSearches         []string `pulumi:"dnsSearches"`
	DnsServersOverrides []string `pulumi:"dnsServersOverrides"`
	EnableDynamic       bool     `pulumi:"enableDynamic"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	OrganizationId   string `pulumi:"organizationId"`
	PoolHigh         string `pulumi:"poolHigh"`
	PoolLow          string `pulumi:"poolLow"`
	ProjectId        string `pulumi:"projectId"`
	PushDefaultRoute bool   `pulumi:"pushDefaultRoute"`
	PushDnsServer    bool   `pulumi:"pushDnsServer"`
	RebindTimer      int    `pulumi:"rebindTimer"`
	RenewTimer       int    `pulumi:"renewTimer"`
	Subnet           string `pulumi:"subnet"`
	UpdatedAt        string `pulumi:"updatedAt"`
	ValidLifetime    int    `pulumi:"validLifetime"`
	Zone             string `pulumi:"zone"`
}

func LookupVpcPublicGatewayDhcpOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayDhcpOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayDhcpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayDhcpResult, error) {
			args := v.(LookupVpcPublicGatewayDhcpArgs)
			r, err := LookupVpcPublicGatewayDhcp(ctx, &args, opts...)
			var s LookupVpcPublicGatewayDhcpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcPublicGatewayDhcpResultOutput)
}

// A collection of arguments for invoking getVpcPublicGatewayDhcp.
type LookupVpcPublicGatewayDhcpOutputArgs struct {
	DhcpId pulumi.StringInput `pulumi:"dhcpId"`
}

func (LookupVpcPublicGatewayDhcpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayDhcpArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGatewayDhcp.
type LookupVpcPublicGatewayDhcpResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayDhcpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayDhcpResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayDhcpResultOutput) ToLookupVpcPublicGatewayDhcpResultOutput() LookupVpcPublicGatewayDhcpResultOutput {
	return o
}

func (o LookupVpcPublicGatewayDhcpResultOutput) ToLookupVpcPublicGatewayDhcpResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayDhcpResultOutput {
	return o
}

func (o LookupVpcPublicGatewayDhcpResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) DhcpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.DhcpId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) DnsLocalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.DnsLocalName }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) DnsSearches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) []string { return v.DnsSearches }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) DnsServersOverrides() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) []string { return v.DnsServersOverrides }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) EnableDynamic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) bool { return v.EnableDynamic }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayDhcpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) PoolHigh() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.PoolHigh }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) PoolLow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.PoolLow }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) PushDefaultRoute() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) bool { return v.PushDefaultRoute }).(pulumi.BoolOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) PushDnsServer() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) bool { return v.PushDnsServer }).(pulumi.BoolOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) RebindTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) int { return v.RebindTimer }).(pulumi.IntOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) RenewTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) int { return v.RenewTimer }).(pulumi.IntOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.Subnet }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) ValidLifetime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) int { return v.ValidLifetime }).(pulumi.IntOutput)
}

func (o LookupVpcPublicGatewayDhcpResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayDhcpResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayDhcpResultOutput{})
}
