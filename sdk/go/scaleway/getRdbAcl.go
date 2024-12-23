// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about the Database Instance network Access Control List.
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
//			_, err := scaleway.LookupRdbAcl(ctx, &scaleway.LookupRdbAclArgs{
//				InstanceId: "11111111-1111-1111-1111-111111111111",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRdbAcl(ctx *pulumi.Context, args *LookupRdbAclArgs, opts ...pulumi.InvokeOption) (*LookupRdbAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRdbAclResult
	err := ctx.Invoke("scaleway:index/getRdbAcl:getRdbAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRdbAcl.
type LookupRdbAclArgs struct {
	// The RDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// `region`) The region in which the Database Instance should be created.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getRdbAcl.
type LookupRdbAclResult struct {
	// A list of ACLs rules (structure is described below)
	AclRules []GetRdbAclAclRule `pulumi:"aclRules"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	Region     *string `pulumi:"region"`
}

func LookupRdbAclOutput(ctx *pulumi.Context, args LookupRdbAclOutputArgs, opts ...pulumi.InvokeOption) LookupRdbAclResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRdbAclResultOutput, error) {
			args := v.(LookupRdbAclArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getRdbAcl:getRdbAcl", args, LookupRdbAclResultOutput{}, options).(LookupRdbAclResultOutput), nil
		}).(LookupRdbAclResultOutput)
}

// A collection of arguments for invoking getRdbAcl.
type LookupRdbAclOutputArgs struct {
	// The RDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupRdbAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdbAclArgs)(nil)).Elem()
}

// A collection of values returned by getRdbAcl.
type LookupRdbAclResultOutput struct{ *pulumi.OutputState }

func (LookupRdbAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdbAclResult)(nil)).Elem()
}

func (o LookupRdbAclResultOutput) ToLookupRdbAclResultOutput() LookupRdbAclResultOutput {
	return o
}

func (o LookupRdbAclResultOutput) ToLookupRdbAclResultOutputWithContext(ctx context.Context) LookupRdbAclResultOutput {
	return o
}

// A list of ACLs rules (structure is described below)
func (o LookupRdbAclResultOutput) AclRules() GetRdbAclAclRuleArrayOutput {
	return o.ApplyT(func(v LookupRdbAclResult) []GetRdbAclAclRule { return v.AclRules }).(GetRdbAclAclRuleArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRdbAclResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbAclResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRdbAclResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbAclResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o LookupRdbAclResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRdbAclResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRdbAclResultOutput{})
}
