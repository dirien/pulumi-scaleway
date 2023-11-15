// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an instance snapshot.
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
//			_, err := scaleway.LookupInstanceSnapshot(ctx, &scaleway.LookupInstanceSnapshotArgs{
//				Name: pulumi.StringRef("my-snapshot-name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupInstanceSnapshot(ctx, &scaleway.LookupInstanceSnapshotArgs{
//				SnapshotId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstanceSnapshot(ctx *pulumi.Context, args *LookupInstanceSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupInstanceSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceSnapshotResult
	err := ctx.Invoke("scaleway:index/getInstanceSnapshot:getInstanceSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceSnapshot.
type LookupInstanceSnapshotArgs struct {
	// The snapshot name.
	// Only one of `name` and `snapshotId` should be specified.
	Name *string `pulumi:"name"`
	// The snapshot id.
	// Only one of `name` and `snapshotId` should be specified.
	SnapshotId *string `pulumi:"snapshotId"`
	// `zone`) The zone in which the snapshot exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceSnapshot.
type LookupInstanceSnapshotResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id             string                      `pulumi:"id"`
	Imports        []GetInstanceSnapshotImport `pulumi:"imports"`
	Name           *string                     `pulumi:"name"`
	OrganizationId string                      `pulumi:"organizationId"`
	ProjectId      string                      `pulumi:"projectId"`
	SizeInGb       int                         `pulumi:"sizeInGb"`
	SnapshotId     *string                     `pulumi:"snapshotId"`
	Tags           []string                    `pulumi:"tags"`
	Type           string                      `pulumi:"type"`
	VolumeId       string                      `pulumi:"volumeId"`
	Zone           *string                     `pulumi:"zone"`
}

func LookupInstanceSnapshotOutput(ctx *pulumi.Context, args LookupInstanceSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceSnapshotResult, error) {
			args := v.(LookupInstanceSnapshotArgs)
			r, err := LookupInstanceSnapshot(ctx, &args, opts...)
			var s LookupInstanceSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceSnapshotResultOutput)
}

// A collection of arguments for invoking getInstanceSnapshot.
type LookupInstanceSnapshotOutputArgs struct {
	// The snapshot name.
	// Only one of `name` and `snapshotId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The snapshot id.
	// Only one of `name` and `snapshotId` should be specified.
	SnapshotId pulumi.StringPtrInput `pulumi:"snapshotId"`
	// `zone`) The zone in which the snapshot exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupInstanceSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceSnapshot.
type LookupInstanceSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceSnapshotResult)(nil)).Elem()
}

func (o LookupInstanceSnapshotResultOutput) ToLookupInstanceSnapshotResultOutput() LookupInstanceSnapshotResultOutput {
	return o
}

func (o LookupInstanceSnapshotResultOutput) ToLookupInstanceSnapshotResultOutputWithContext(ctx context.Context) LookupInstanceSnapshotResultOutput {
	return o
}

func (o LookupInstanceSnapshotResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceSnapshotResultOutput) Imports() GetInstanceSnapshotImportArrayOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) []GetInstanceSnapshotImport { return v.Imports }).(GetInstanceSnapshotImportArrayOutput)
}

func (o LookupInstanceSnapshotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceSnapshotResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupInstanceSnapshotResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupInstanceSnapshotResultOutput) SizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) int { return v.SizeInGb }).(pulumi.IntOutput)
}

func (o LookupInstanceSnapshotResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceSnapshotResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupInstanceSnapshotResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) string { return v.VolumeId }).(pulumi.StringOutput)
}

func (o LookupInstanceSnapshotResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceSnapshotResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceSnapshotResultOutput{})
}
