// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Public Gateway.
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
//			main, err := scaleway.NewVpcPublicGateway(ctx, "main", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				Zone: pulumi.String("nl-ams-1"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupVpcPublicGatewayOutput(ctx, scaleway.GetVpcPublicGatewayOutputArgs{
//				Name: main.Name,
//				Zone: pulumi.String("nl-ams-1"),
//			}, nil)
//			_ = scaleway.LookupVpcPublicGatewayOutput(ctx, scaleway.GetVpcPublicGatewayOutputArgs{
//				PublicGatewayId: main.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupVpcPublicGateway(ctx *pulumi.Context, args *LookupVpcPublicGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPublicGatewayResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGateway:getVpcPublicGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGateway.
type LookupVpcPublicGatewayArgs struct {
	// Exact name of the Public Gateway.
	Name *string `pulumi:"name"`
	// The ID of the Project the Public Gateway is associated with.
	ProjectId       *string `pulumi:"projectId"`
	PublicGatewayId *string `pulumi:"publicGatewayId"`
	// `zone`) The Public Gateway's zone.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getVpcPublicGateway.
type LookupVpcPublicGatewayResult struct {
	AllowedIpRanges []string `pulumi:"allowedIpRanges"`
	Bandwidth       int      `pulumi:"bandwidth"`
	BastionEnabled  bool     `pulumi:"bastionEnabled"`
	BastionPort     int      `pulumi:"bastionPort"`
	CreatedAt       string   `pulumi:"createdAt"`
	EnableSmtp      bool     `pulumi:"enableSmtp"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string   `pulumi:"id"`
	IpId               string   `pulumi:"ipId"`
	MoveToIpam         bool     `pulumi:"moveToIpam"`
	Name               *string  `pulumi:"name"`
	OrganizationId     string   `pulumi:"organizationId"`
	ProjectId          *string  `pulumi:"projectId"`
	PublicGatewayId    *string  `pulumi:"publicGatewayId"`
	RefreshSshKeys     string   `pulumi:"refreshSshKeys"`
	Status             string   `pulumi:"status"`
	Tags               []string `pulumi:"tags"`
	Type               string   `pulumi:"type"`
	UpdatedAt          string   `pulumi:"updatedAt"`
	UpstreamDnsServers []string `pulumi:"upstreamDnsServers"`
	Zone               *string  `pulumi:"zone"`
}

func LookupVpcPublicGatewayOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayResultOutput, error) {
			args := v.(LookupVpcPublicGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getVpcPublicGateway:getVpcPublicGateway", args, LookupVpcPublicGatewayResultOutput{}, options).(LookupVpcPublicGatewayResultOutput), nil
		}).(LookupVpcPublicGatewayResultOutput)
}

// A collection of arguments for invoking getVpcPublicGateway.
type LookupVpcPublicGatewayOutputArgs struct {
	// Exact name of the Public Gateway.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the Project the Public Gateway is associated with.
	ProjectId       pulumi.StringPtrInput `pulumi:"projectId"`
	PublicGatewayId pulumi.StringPtrInput `pulumi:"publicGatewayId"`
	// `zone`) The Public Gateway's zone.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupVpcPublicGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGateway.
type LookupVpcPublicGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayResultOutput) ToLookupVpcPublicGatewayResultOutput() LookupVpcPublicGatewayResultOutput {
	return o
}

func (o LookupVpcPublicGatewayResultOutput) ToLookupVpcPublicGatewayResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayResultOutput {
	return o
}

func (o LookupVpcPublicGatewayResultOutput) AllowedIpRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) []string { return v.AllowedIpRanges }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayResultOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) int { return v.Bandwidth }).(pulumi.IntOutput)
}

func (o LookupVpcPublicGatewayResultOutput) BastionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) bool { return v.BastionEnabled }).(pulumi.BoolOutput)
}

func (o LookupVpcPublicGatewayResultOutput) BastionPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) int { return v.BastionPort }).(pulumi.IntOutput)
}

func (o LookupVpcPublicGatewayResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) EnableSmtp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) bool { return v.EnableSmtp }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.IpId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) MoveToIpam() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) bool { return v.MoveToIpam }).(pulumi.BoolOutput)
}

func (o LookupVpcPublicGatewayResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayResultOutput) PublicGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) *string { return v.PublicGatewayId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcPublicGatewayResultOutput) RefreshSshKeys() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.RefreshSshKeys }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayResultOutput) UpstreamDnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) []string { return v.UpstreamDnsServers }).(pulumi.StringArrayOutput)
}

func (o LookupVpcPublicGatewayResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayResultOutput{})
}
