// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Load Balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/dirien/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.LookupLb(ctx, &scaleway.LookupLbArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupLb(ctx, &scaleway.LookupLbArgs{
//				LbId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLb(ctx *pulumi.Context, args *LookupLbArgs, opts ...pulumi.InvokeOption) (*LookupLbResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupLbResult
	err := ctx.Invoke("scaleway:index/getLb:getLb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLb.
type LookupLbArgs struct {
	LbId *string `pulumi:"lbId"`
	// The load balancer name.
	Name      *string `pulumi:"name"`
	ReleaseIp *bool   `pulumi:"releaseIp"`
	// (Defaults to provider `zone`) The zone in which the LB exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLb.
type LookupLbResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The load-balancer public IP Address.
	IpAddress       string                `pulumi:"ipAddress"`
	IpId            string                `pulumi:"ipId"`
	LbId            *string               `pulumi:"lbId"`
	Name            *string               `pulumi:"name"`
	OrganizationId  string                `pulumi:"organizationId"`
	PrivateNetworks []GetLbPrivateNetwork `pulumi:"privateNetworks"`
	// (Defaults to provider `projectId`) The ID of the project the LB is associated with.
	ProjectId             string `pulumi:"projectId"`
	Region                string `pulumi:"region"`
	ReleaseIp             *bool  `pulumi:"releaseIp"`
	SslCompatibilityLevel string `pulumi:"sslCompatibilityLevel"`
	// The tags associated with the load-balancer.
	Tags []string `pulumi:"tags"`
	// The type of the load-balancer.
	Type string `pulumi:"type"`
	// (Defaults to provider `zone`) The zone in which the LB exists.
	Zone *string `pulumi:"zone"`
}

func LookupLbOutput(ctx *pulumi.Context, args LookupLbOutputArgs, opts ...pulumi.InvokeOption) LookupLbResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbResult, error) {
			args := v.(LookupLbArgs)
			r, err := LookupLb(ctx, &args, opts...)
			var s LookupLbResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLbResultOutput)
}

// A collection of arguments for invoking getLb.
type LookupLbOutputArgs struct {
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The load balancer name.
	Name      pulumi.StringPtrInput `pulumi:"name"`
	ReleaseIp pulumi.BoolPtrInput   `pulumi:"releaseIp"`
	// (Defaults to provider `zone`) The zone in which the LB exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupLbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbArgs)(nil)).Elem()
}

// A collection of values returned by getLb.
type LookupLbResultOutput struct{ *pulumi.OutputState }

func (LookupLbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbResult)(nil)).Elem()
}

func (o LookupLbResultOutput) ToLookupLbResultOutput() LookupLbResultOutput {
	return o
}

func (o LookupLbResultOutput) ToLookupLbResultOutputWithContext(ctx context.Context) LookupLbResultOutput {
	return o
}

func (o LookupLbResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.Id }).(pulumi.StringOutput)
}

// The load-balancer public IP Address.
func (o LookupLbResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupLbResultOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.IpId }).(pulumi.StringOutput)
}

func (o LookupLbResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupLbResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLbResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupLbResultOutput) PrivateNetworks() GetLbPrivateNetworkArrayOutput {
	return o.ApplyT(func(v LookupLbResult) []GetLbPrivateNetwork { return v.PrivateNetworks }).(GetLbPrivateNetworkArrayOutput)
}

// (Defaults to provider `projectId`) The ID of the project the LB is associated with.
func (o LookupLbResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupLbResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupLbResultOutput) ReleaseIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLbResult) *bool { return v.ReleaseIp }).(pulumi.BoolPtrOutput)
}

func (o LookupLbResultOutput) SslCompatibilityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.SslCompatibilityLevel }).(pulumi.StringOutput)
}

// The tags associated with the load-balancer.
func (o LookupLbResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLbResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the load-balancer.
func (o LookupLbResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbResult) string { return v.Type }).(pulumi.StringOutput)
}

// (Defaults to provider `zone`) The zone in which the LB exists.
func (o LookupLbResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbResultOutput{})
}
