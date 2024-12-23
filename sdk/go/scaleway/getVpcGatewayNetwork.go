// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a GatewayNetwork (a connection between a Public Gateway and a Private Network).
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
//			main, err := scaleway.NewVpcGatewayNetwork(ctx, "main", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        pulumi.Any(scaleway_vpc_public_gateway.Pg01.Id),
//				PrivateNetworkId: pulumi.Any(scaleway_vpc_private_network.Pn01.Id),
//				DhcpId:           pulumi.Any(scaleway_vpc_public_gateway_dhcp.Dhcp01.Id),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupVpcGatewayNetworkOutput(ctx, scaleway.GetVpcGatewayNetworkOutputArgs{
//				GatewayNetworkId: main.ID(),
//			}, nil)
//			_, err = scaleway.LookupVpcGatewayNetwork(ctx, &scaleway.LookupVpcGatewayNetworkArgs{
//				GatewayId:        pulumi.StringRef(scaleway_vpc_public_gateway.Pg01.Id),
//				PrivateNetworkId: pulumi.StringRef(scaleway_vpc_private_network.Pn01.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVpcGatewayNetwork(ctx *pulumi.Context, args *LookupVpcGatewayNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVpcGatewayNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcGatewayNetworkResult
	err := ctx.Invoke("scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcGatewayNetwork.
type LookupVpcGatewayNetworkArgs struct {
	DhcpId           *string `pulumi:"dhcpId"`
	EnableMasquerade *bool   `pulumi:"enableMasquerade"`
	GatewayId        *string `pulumi:"gatewayId"`
	// ID of the GatewayNetwork.
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
}

// A collection of values returned by getVpcGatewayNetwork.
type LookupVpcGatewayNetworkResult struct {
	CleanupDhcp      bool    `pulumi:"cleanupDhcp"`
	CreatedAt        string  `pulumi:"createdAt"`
	DhcpId           *string `pulumi:"dhcpId"`
	EnableDhcp       bool    `pulumi:"enableDhcp"`
	EnableMasquerade *bool   `pulumi:"enableMasquerade"`
	GatewayId        *string `pulumi:"gatewayId"`
	GatewayNetworkId *string `pulumi:"gatewayNetworkId"`
	// The provider-assigned unique ID for this managed resource.
	Id               string                           `pulumi:"id"`
	IpamConfigs      []GetVpcGatewayNetworkIpamConfig `pulumi:"ipamConfigs"`
	MacAddress       string                           `pulumi:"macAddress"`
	PrivateNetworkId *string                          `pulumi:"privateNetworkId"`
	StaticAddress    string                           `pulumi:"staticAddress"`
	Status           string                           `pulumi:"status"`
	UpdatedAt        string                           `pulumi:"updatedAt"`
	Zone             string                           `pulumi:"zone"`
}

func LookupVpcGatewayNetworkOutput(ctx *pulumi.Context, args LookupVpcGatewayNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVpcGatewayNetworkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcGatewayNetworkResultOutput, error) {
			args := v.(LookupVpcGatewayNetworkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork", args, LookupVpcGatewayNetworkResultOutput{}, options).(LookupVpcGatewayNetworkResultOutput), nil
		}).(LookupVpcGatewayNetworkResultOutput)
}

// A collection of arguments for invoking getVpcGatewayNetwork.
type LookupVpcGatewayNetworkOutputArgs struct {
	DhcpId           pulumi.StringPtrInput `pulumi:"dhcpId"`
	EnableMasquerade pulumi.BoolPtrInput   `pulumi:"enableMasquerade"`
	GatewayId        pulumi.StringPtrInput `pulumi:"gatewayId"`
	// ID of the GatewayNetwork.
	GatewayNetworkId pulumi.StringPtrInput `pulumi:"gatewayNetworkId"`
	PrivateNetworkId pulumi.StringPtrInput `pulumi:"privateNetworkId"`
}

func (LookupVpcGatewayNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcGatewayNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getVpcGatewayNetwork.
type LookupVpcGatewayNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVpcGatewayNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcGatewayNetworkResult)(nil)).Elem()
}

func (o LookupVpcGatewayNetworkResultOutput) ToLookupVpcGatewayNetworkResultOutput() LookupVpcGatewayNetworkResultOutput {
	return o
}

func (o LookupVpcGatewayNetworkResultOutput) ToLookupVpcGatewayNetworkResultOutputWithContext(ctx context.Context) LookupVpcGatewayNetworkResultOutput {
	return o
}

func (o LookupVpcGatewayNetworkResultOutput) CleanupDhcp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) bool { return v.CleanupDhcp }).(pulumi.BoolOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) DhcpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) *string { return v.DhcpId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) EnableDhcp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) bool { return v.EnableDhcp }).(pulumi.BoolOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) EnableMasquerade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) *bool { return v.EnableMasquerade }).(pulumi.BoolPtrOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) GatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) *string { return v.GatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) GatewayNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) *string { return v.GatewayNetworkId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcGatewayNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) IpamConfigs() GetVpcGatewayNetworkIpamConfigArrayOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) []GetVpcGatewayNetworkIpamConfig { return v.IpamConfigs }).(GetVpcGatewayNetworkIpamConfigArrayOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.MacAddress }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) *string { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) StaticAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.StaticAddress }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcGatewayNetworkResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcGatewayNetworkResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcGatewayNetworkResultOutput{})
}
