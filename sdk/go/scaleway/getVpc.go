// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a Scaleway Virtual Private Cloud.
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
//			_, err := scaleway.LookupVpc(ctx, &scaleway.LookupVpcArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupVpc(ctx, &scaleway.LookupVpcArgs{
//				VpcId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupVpc(ctx, &scaleway.LookupVpcArgs{
//				IsDefault: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVpc(ctx *pulumi.Context, args *LookupVpcArgs, opts ...pulumi.InvokeOption) (*LookupVpcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpcResult
	err := ctx.Invoke("scaleway:index/getVpc:getVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpc.
type LookupVpcArgs struct {
	// Whether the targeted VPC is the default VPC.
	IsDefault *bool `pulumi:"isDefault"`
	// Name of the VPC. A maximum of one of `name` and `vpcId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the Organization the VPC is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the Project the VPC is associated with.
	ProjectId *string `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// ID of the VPC. A maximum of one of `name` and `vpcId` should be specified.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVpc.
type LookupVpcResult struct {
	CreatedAt     string `pulumi:"createdAt"`
	EnableRouting bool   `pulumi:"enableRouting"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	IsDefault      *bool    `pulumi:"isDefault"`
	Name           *string  `pulumi:"name"`
	OrganizationId string   `pulumi:"organizationId"`
	ProjectId      *string  `pulumi:"projectId"`
	Region         *string  `pulumi:"region"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      string   `pulumi:"updatedAt"`
	VpcId          *string  `pulumi:"vpcId"`
}

func LookupVpcOutput(ctx *pulumi.Context, args LookupVpcOutputArgs, opts ...pulumi.InvokeOption) LookupVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpcResult, error) {
			args := v.(LookupVpcArgs)
			r, err := LookupVpc(ctx, &args, opts...)
			var s LookupVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpcResultOutput)
}

// A collection of arguments for invoking getVpc.
type LookupVpcOutputArgs struct {
	// Whether the targeted VPC is the default VPC.
	IsDefault pulumi.BoolPtrInput `pulumi:"isDefault"`
	// Name of the VPC. A maximum of one of `name` and `vpcId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the Organization the VPC is associated with.
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
	// `projectId`) The ID of the Project the VPC is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	Region    pulumi.StringPtrInput `pulumi:"region"`
	// ID of the VPC. A maximum of one of `name` and `vpcId` should be specified.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcArgs)(nil)).Elem()
}

// A collection of values returned by getVpc.
type LookupVpcResultOutput struct{ *pulumi.OutputState }

func (LookupVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpcResult)(nil)).Elem()
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutput() LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) ToLookupVpcResultOutputWithContext(ctx context.Context) LookupVpcResultOutput {
	return o
}

func (o LookupVpcResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) EnableRouting() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpcResult) bool { return v.EnableRouting }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

func (o LookupVpcResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVpcResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupVpcResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupVpcResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpcResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupVpcResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpcResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupVpcResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpcResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpcResultOutput{})
}
