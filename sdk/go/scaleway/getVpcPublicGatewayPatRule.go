// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a public gateway PAT rule. For further information please check the
// API [documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#get-8faeea)
func LookupVpcPublicGatewayPatRule(ctx *pulumi.Context, args *LookupVpcPublicGatewayPatRuleArgs, opts ...pulumi.InvokeOption) (*LookupVpcPublicGatewayPatRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcPublicGatewayPatRuleResult
	err := ctx.Invoke("scaleway:index/getVpcPublicGatewayPatRule:getVpcPublicGatewayPatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcPublicGatewayPatRule.
type LookupVpcPublicGatewayPatRuleArgs struct {
	// The ID of the PAT rule to retrieve
	PatRuleId string `pulumi:"patRuleId"`
	// `zone`) The zone in which
	// the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getVpcPublicGatewayPatRule.
type LookupVpcPublicGatewayPatRuleResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The ID of the public gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	OrganizationId string `pulumi:"organizationId"`
	PatRuleId      string `pulumi:"patRuleId"`
	// The Private IP to forward data to (IP address).
	PrivateIp string `pulumi:"privateIp"`
	// The Private port to translate to.
	PrivatePort int `pulumi:"privatePort"`
	// The Protocol the rule should apply to. Possible values are both, tcp and udp.
	Protocol string `pulumi:"protocol"`
	// The Public port to listen on.
	PublicPort int     `pulumi:"publicPort"`
	UpdatedAt  string  `pulumi:"updatedAt"`
	Zone       *string `pulumi:"zone"`
}

func LookupVpcPublicGatewayPatRuleOutput(ctx *pulumi.Context, args LookupVpcPublicGatewayPatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVpcPublicGatewayPatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcPublicGatewayPatRuleResult, error) {
			args := v.(LookupVpcPublicGatewayPatRuleArgs)
			r, err := LookupVpcPublicGatewayPatRule(ctx, &args, opts...)
			var s LookupVpcPublicGatewayPatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcPublicGatewayPatRuleResultOutput)
}

// A collection of arguments for invoking getVpcPublicGatewayPatRule.
type LookupVpcPublicGatewayPatRuleOutputArgs struct {
	// The ID of the PAT rule to retrieve
	PatRuleId pulumi.StringInput `pulumi:"patRuleId"`
	// `zone`) The zone in which
	// the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupVpcPublicGatewayPatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayPatRuleArgs)(nil)).Elem()
}

// A collection of values returned by getVpcPublicGatewayPatRule.
type LookupVpcPublicGatewayPatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVpcPublicGatewayPatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcPublicGatewayPatRuleResult)(nil)).Elem()
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) ToLookupVpcPublicGatewayPatRuleResultOutput() LookupVpcPublicGatewayPatRuleResultOutput {
	return o
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) ToLookupVpcPublicGatewayPatRuleResultOutputWithContext(ctx context.Context) LookupVpcPublicGatewayPatRuleResultOutput {
	return o
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the public gateway.
func (o LookupVpcPublicGatewayPatRuleResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcPublicGatewayPatRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) PatRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.PatRuleId }).(pulumi.StringOutput)
}

// The Private IP to forward data to (IP address).
func (o LookupVpcPublicGatewayPatRuleResultOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// The Private port to translate to.
func (o LookupVpcPublicGatewayPatRuleResultOutput) PrivatePort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) int { return v.PrivatePort }).(pulumi.IntOutput)
}

// The Protocol the rule should apply to. Possible values are both, tcp and udp.
func (o LookupVpcPublicGatewayPatRuleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// The Public port to listen on.
func (o LookupVpcPublicGatewayPatRuleResultOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) int { return v.PublicPort }).(pulumi.IntOutput)
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcPublicGatewayPatRuleResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcPublicGatewayPatRuleResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcPublicGatewayPatRuleResultOutput{})
}
