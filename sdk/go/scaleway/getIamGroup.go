// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an existing IAM group.
//
// For more information, refer to the [IAM API documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e)
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
//			_, err := scaleway.LookupIamGroup(ctx, &scaleway.LookupIamGroupArgs{
//				Name: pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupIamGroup(ctx, &scaleway.LookupIamGroupArgs{
//				GroupId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIamGroup(ctx *pulumi.Context, args *LookupIamGroupArgs, opts ...pulumi.InvokeOption) (*LookupIamGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIamGroupResult
	err := ctx.Invoke("scaleway:index/getIamGroup:getIamGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamGroup.
type LookupIamGroupArgs struct {
	// The ID of the IAM group.
	//
	// > **Note** You must specify at least one: `name` and/or `groupId`.
	GroupId *string `pulumi:"groupId"`
	// The name of the IAM group.
	Name *string `pulumi:"name"`
	// `organizationId`) The ID of the
	// organization the group is associated with.
	OrganizationId *string `pulumi:"organizationId"`
}

// A collection of values returned by getIamGroup.
type LookupIamGroupResult struct {
	ApplicationIds     []string `pulumi:"applicationIds"`
	CreatedAt          string   `pulumi:"createdAt"`
	Description        string   `pulumi:"description"`
	ExternalMembership bool     `pulumi:"externalMembership"`
	GroupId            *string  `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	Name           *string  `pulumi:"name"`
	OrganizationId *string  `pulumi:"organizationId"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      string   `pulumi:"updatedAt"`
	UserIds        []string `pulumi:"userIds"`
}

func LookupIamGroupOutput(ctx *pulumi.Context, args LookupIamGroupOutputArgs, opts ...pulumi.InvokeOption) LookupIamGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIamGroupResultOutput, error) {
			args := v.(LookupIamGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getIamGroup:getIamGroup", args, LookupIamGroupResultOutput{}, options).(LookupIamGroupResultOutput), nil
		}).(LookupIamGroupResultOutput)
}

// A collection of arguments for invoking getIamGroup.
type LookupIamGroupOutputArgs struct {
	// The ID of the IAM group.
	//
	// > **Note** You must specify at least one: `name` and/or `groupId`.
	GroupId pulumi.StringPtrInput `pulumi:"groupId"`
	// The name of the IAM group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `organizationId`) The ID of the
	// organization the group is associated with.
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
}

func (LookupIamGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamGroupArgs)(nil)).Elem()
}

// A collection of values returned by getIamGroup.
type LookupIamGroupResultOutput struct{ *pulumi.OutputState }

func (LookupIamGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamGroupResult)(nil)).Elem()
}

func (o LookupIamGroupResultOutput) ToLookupIamGroupResultOutput() LookupIamGroupResultOutput {
	return o
}

func (o LookupIamGroupResultOutput) ToLookupIamGroupResultOutputWithContext(ctx context.Context) LookupIamGroupResultOutput {
	return o
}

func (o LookupIamGroupResultOutput) ApplicationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIamGroupResult) []string { return v.ApplicationIds }).(pulumi.StringArrayOutput)
}

func (o LookupIamGroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamGroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupIamGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupIamGroupResultOutput) ExternalMembership() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIamGroupResult) bool { return v.ExternalMembership }).(pulumi.BoolOutput)
}

func (o LookupIamGroupResultOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIamGroupResult) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIamGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIamGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIamGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupIamGroupResultOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIamGroupResult) *string { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

func (o LookupIamGroupResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIamGroupResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupIamGroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamGroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupIamGroupResultOutput) UserIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIamGroupResult) []string { return v.UserIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIamGroupResultOutput{})
}
