// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a database.
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
//			_, err := scaleway.LookupRdbDatabase(ctx, &scaleway.LookupRdbDatabaseArgs{
//				InstanceId: "11111111-1111-1111-1111-111111111111",
//				Name:       "foobar",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRdbDatabase(ctx *pulumi.Context, args *LookupRdbDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupRdbDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRdbDatabaseResult
	err := ctx.Invoke("scaleway:index/getRdbDatabase:getRdbDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRdbDatabase.
type LookupRdbDatabaseArgs struct {
	// The RDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// The name of the RDB instance.
	Name   string  `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// A collection of values returned by getRdbDatabase.
type LookupRdbDatabaseResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed bool   `pulumi:"managed"`
	Name    string `pulumi:"name"`
	// The name of the owner of the database.
	Owner  string  `pulumi:"owner"`
	Region *string `pulumi:"region"`
	// Size of the database (in bytes).
	Size string `pulumi:"size"`
}

func LookupRdbDatabaseOutput(ctx *pulumi.Context, args LookupRdbDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupRdbDatabaseResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRdbDatabaseResultOutput, error) {
			args := v.(LookupRdbDatabaseArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getRdbDatabase:getRdbDatabase", args, LookupRdbDatabaseResultOutput{}, options).(LookupRdbDatabaseResultOutput), nil
		}).(LookupRdbDatabaseResultOutput)
}

// A collection of arguments for invoking getRdbDatabase.
type LookupRdbDatabaseOutputArgs struct {
	// The RDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The name of the RDB instance.
	Name   pulumi.StringInput    `pulumi:"name"`
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupRdbDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdbDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getRdbDatabase.
type LookupRdbDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupRdbDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdbDatabaseResult)(nil)).Elem()
}

func (o LookupRdbDatabaseResultOutput) ToLookupRdbDatabaseResultOutput() LookupRdbDatabaseResultOutput {
	return o
}

func (o LookupRdbDatabaseResultOutput) ToLookupRdbDatabaseResultOutputWithContext(ctx context.Context) LookupRdbDatabaseResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRdbDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRdbDatabaseResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o LookupRdbDatabaseResultOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) bool { return v.Managed }).(pulumi.BoolOutput)
}

func (o LookupRdbDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the owner of the database.
func (o LookupRdbDatabaseResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupRdbDatabaseResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// Size of the database (in bytes).
func (o LookupRdbDatabaseResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdbDatabaseResult) string { return v.Size }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRdbDatabaseResultOutput{})
}
