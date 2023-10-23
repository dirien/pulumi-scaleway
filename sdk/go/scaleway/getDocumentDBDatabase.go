// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about DocumentDB database. More on our official [site](https://www.scaleway.com/en/developers/api/document_db/)
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
//			_, err := scaleway.LookupDocumentDBDatabase(ctx, &scaleway.LookupDocumentDBDatabaseArgs{
//				InstanceId: "11111111-1111-1111-1111-111111111111",
//				Name:       pulumi.StringRef("foobar"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDocumentDBDatabase(ctx *pulumi.Context, args *LookupDocumentDBDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDocumentDBDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDocumentDBDatabaseResult
	err := ctx.Invoke("scaleway:index/getDocumentDBDatabase:getDocumentDBDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocumentDBDatabase.
type LookupDocumentDBDatabaseArgs struct {
	// The DocumentDB instance ID.
	InstanceId string `pulumi:"instanceId"`
	// The name of the DocumentDB instance.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDocumentDBDatabase.
type LookupDocumentDBDatabaseResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string `pulumi:"id"`
	InstanceId string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed bool    `pulumi:"managed"`
	Name    *string `pulumi:"name"`
	// The name of the owner of the database.
	Owner     string  `pulumi:"owner"`
	ProjectId string  `pulumi:"projectId"`
	Region    *string `pulumi:"region"`
	// Size of the database (in bytes).
	Size string `pulumi:"size"`
}

func LookupDocumentDBDatabaseOutput(ctx *pulumi.Context, args LookupDocumentDBDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentDBDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentDBDatabaseResult, error) {
			args := v.(LookupDocumentDBDatabaseArgs)
			r, err := LookupDocumentDBDatabase(ctx, &args, opts...)
			var s LookupDocumentDBDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentDBDatabaseResultOutput)
}

// A collection of arguments for invoking getDocumentDBDatabase.
type LookupDocumentDBDatabaseOutputArgs struct {
	// The DocumentDB instance ID.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The name of the DocumentDB instance.
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupDocumentDBDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentDBDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getDocumentDBDatabase.
type LookupDocumentDBDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentDBDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentDBDatabaseResult)(nil)).Elem()
}

func (o LookupDocumentDBDatabaseResultOutput) ToLookupDocumentDBDatabaseResultOutput() LookupDocumentDBDatabaseResultOutput {
	return o
}

func (o LookupDocumentDBDatabaseResultOutput) ToLookupDocumentDBDatabaseResultOutputWithContext(ctx context.Context) LookupDocumentDBDatabaseResultOutput {
	return o
}

func (o LookupDocumentDBDatabaseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDocumentDBDatabaseResult] {
	return pulumix.Output[LookupDocumentDBDatabaseResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDocumentDBDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDocumentDBDatabaseResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o LookupDocumentDBDatabaseResultOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) bool { return v.Managed }).(pulumi.BoolOutput)
}

func (o LookupDocumentDBDatabaseResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The name of the owner of the database.
func (o LookupDocumentDBDatabaseResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupDocumentDBDatabaseResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupDocumentDBDatabaseResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// Size of the database (in bytes).
func (o LookupDocumentDBDatabaseResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentDBDatabaseResult) string { return v.Size }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentDBDatabaseResultOutput{})
}
