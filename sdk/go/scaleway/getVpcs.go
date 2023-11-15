// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about multiple Virtual Private Clouds.
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
//			_, err := scaleway.GetVpcs(ctx, &scaleway.GetVpcsArgs{
//				Name:   pulumi.StringRef("tf-vpc-datasource"),
//				Region: pulumi.StringRef("nl-ams"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVpcs(ctx *pulumi.Context, args *GetVpcsArgs, opts ...pulumi.InvokeOption) (*GetVpcsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcsResult
	err := ctx.Invoke("scaleway:index/getVpcs:getVpcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcs.
type GetVpcsArgs struct {
	// The VPC name used as filter. VPCs with a name like it are listed.
	Name *string `pulumi:"name"`
	// The ID of the project the VPC is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which vpcs exist.
	Region *string `pulumi:"region"`
	// List of tags used as filter. VPCs with these exact tags are listed.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getVpcs.
type GetVpcsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	// The organization ID the VPC is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the project the VPC is associated with.
	ProjectId string   `pulumi:"projectId"`
	Region    string   `pulumi:"region"`
	Tags      []string `pulumi:"tags"`
	// List of found vpcs
	Vpcs []GetVpcsVpc `pulumi:"vpcs"`
}

func GetVpcsOutput(ctx *pulumi.Context, args GetVpcsOutputArgs, opts ...pulumi.InvokeOption) GetVpcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcsResult, error) {
			args := v.(GetVpcsArgs)
			r, err := GetVpcs(ctx, &args, opts...)
			var s GetVpcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcsResultOutput)
}

// A collection of arguments for invoking getVpcs.
type GetVpcsOutputArgs struct {
	// The VPC name used as filter. VPCs with a name like it are listed.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the VPC is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`). The region in which vpcs exist.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// List of tags used as filter. VPCs with these exact tags are listed.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
}

func (GetVpcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcsArgs)(nil)).Elem()
}

// A collection of values returned by getVpcs.
type GetVpcsResultOutput struct{ *pulumi.OutputState }

func (GetVpcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcsResult)(nil)).Elem()
}

func (o GetVpcsResultOutput) ToGetVpcsResultOutput() GetVpcsResultOutput {
	return o
}

func (o GetVpcsResultOutput) ToGetVpcsResultOutputWithContext(ctx context.Context) GetVpcsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The organization ID the VPC is associated with.
func (o GetVpcsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the project the VPC is associated with.
func (o GetVpcsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetVpcsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetVpcsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// List of found vpcs
func (o GetVpcsResultOutput) Vpcs() GetVpcsVpcArrayOutput {
	return o.ApplyT(func(v GetVpcsResult) []GetVpcsVpc { return v.Vpcs }).(GetVpcsVpcArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcsResultOutput{})
}
