// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBlockSnapshot(ctx *pulumi.Context, args *LookupBlockSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupBlockSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBlockSnapshotResult
	err := ctx.Invoke("scaleway:index/getBlockSnapshot:getBlockSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBlockSnapshot.
type LookupBlockSnapshotArgs struct {
	Name       *string `pulumi:"name"`
	ProjectId  *string `pulumi:"projectId"`
	SnapshotId *string `pulumi:"snapshotId"`
	VolumeId   *string `pulumi:"volumeId"`
	Zone       *string `pulumi:"zone"`
}

// A collection of values returned by getBlockSnapshot.
type LookupBlockSnapshotResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Name       *string  `pulumi:"name"`
	ProjectId  *string  `pulumi:"projectId"`
	SnapshotId *string  `pulumi:"snapshotId"`
	Tags       []string `pulumi:"tags"`
	VolumeId   *string  `pulumi:"volumeId"`
	Zone       *string  `pulumi:"zone"`
}

func LookupBlockSnapshotOutput(ctx *pulumi.Context, args LookupBlockSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupBlockSnapshotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBlockSnapshotResultOutput, error) {
			args := v.(LookupBlockSnapshotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getBlockSnapshot:getBlockSnapshot", args, LookupBlockSnapshotResultOutput{}, options).(LookupBlockSnapshotResultOutput), nil
		}).(LookupBlockSnapshotResultOutput)
}

// A collection of arguments for invoking getBlockSnapshot.
type LookupBlockSnapshotOutputArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	ProjectId  pulumi.StringPtrInput `pulumi:"projectId"`
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
	VolumeId   pulumi.StringPtrInput `pulumi:"volumeId"`
	Zone       pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupBlockSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlockSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getBlockSnapshot.
type LookupBlockSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupBlockSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBlockSnapshotResult)(nil)).Elem()
}

func (o LookupBlockSnapshotResultOutput) ToLookupBlockSnapshotResultOutput() LookupBlockSnapshotResultOutput {
	return o
}

func (o LookupBlockSnapshotResultOutput) ToLookupBlockSnapshotResultOutputWithContext(ctx context.Context) LookupBlockSnapshotResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBlockSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBlockSnapshotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupBlockSnapshotResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupBlockSnapshotResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o LookupBlockSnapshotResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupBlockSnapshotResultOutput) VolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) *string { return v.VolumeId }).(pulumi.StringPtrOutput)
}

func (o LookupBlockSnapshotResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBlockSnapshotResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBlockSnapshotResultOutput{})
}
