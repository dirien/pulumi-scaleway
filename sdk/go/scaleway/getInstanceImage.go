// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an instance image.
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
//			_, err := scaleway.LookupInstanceImage(ctx, &scaleway.LookupInstanceImageArgs{
//				ImageId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInstanceImage(ctx *pulumi.Context, args *LookupInstanceImageArgs, opts ...pulumi.InvokeOption) (*LookupInstanceImageResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupInstanceImageResult
	err := ctx.Invoke("scaleway:index/getInstanceImage:getInstanceImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceImage.
type LookupInstanceImageArgs struct {
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture *string `pulumi:"architecture"`
	// The image id. Only one of `name` and `imageId` should be specified.
	ImageId *string `pulumi:"imageId"`
	// Use the latest image ID.
	Latest *bool `pulumi:"latest"`
	// The image name. Only one of `name` and `imageId` should be specified.
	Name *string `pulumi:"name"`
	// The ID of the project the image is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `zone`) The zone in which the image exists.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getInstanceImage.
type LookupInstanceImageResult struct {
	// IDs of the additional volumes in this image.
	AdditionalVolumeIds []string `pulumi:"additionalVolumeIds"`
	Architecture        *string  `pulumi:"architecture"`
	// Date of the image creation.
	CreationDate string `pulumi:"creationDate"`
	// ID of the default bootscript for this image.
	DefaultBootscriptId string `pulumi:"defaultBootscriptId"`
	// ID of the server the image if based from.
	FromServerId string `pulumi:"fromServerId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	ImageId *string `pulumi:"imageId"`
	Latest  *bool   `pulumi:"latest"`
	// Date of image latest update.
	ModificationDate string  `pulumi:"modificationDate"`
	Name             *string `pulumi:"name"`
	// The ID of the organization the image is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the project the image is associated with.
	ProjectId string `pulumi:"projectId"`
	// Set to `true` if the image is public.
	Public bool `pulumi:"public"`
	// ID of the root volume in this image.
	RootVolumeId string `pulumi:"rootVolumeId"`
	// State of the image. Possible values are: `available`, `creating` or `error`.
	State string `pulumi:"state"`
	Zone  string `pulumi:"zone"`
}

func LookupInstanceImageOutput(ctx *pulumi.Context, args LookupInstanceImageOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceImageResult, error) {
			args := v.(LookupInstanceImageArgs)
			r, err := LookupInstanceImage(ctx, &args, opts...)
			var s LookupInstanceImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceImageResultOutput)
}

// A collection of arguments for invoking getInstanceImage.
type LookupInstanceImageOutputArgs struct {
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture pulumi.StringPtrInput `pulumi:"architecture"`
	// The image id. Only one of `name` and `imageId` should be specified.
	ImageId pulumi.StringPtrInput `pulumi:"imageId"`
	// Use the latest image ID.
	Latest pulumi.BoolPtrInput `pulumi:"latest"`
	// The image name. Only one of `name` and `imageId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the project the image is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `zone`) The zone in which the image exists.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (LookupInstanceImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceImageArgs)(nil)).Elem()
}

// A collection of values returned by getInstanceImage.
type LookupInstanceImageResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceImageResult)(nil)).Elem()
}

func (o LookupInstanceImageResultOutput) ToLookupInstanceImageResultOutput() LookupInstanceImageResultOutput {
	return o
}

func (o LookupInstanceImageResultOutput) ToLookupInstanceImageResultOutputWithContext(ctx context.Context) LookupInstanceImageResultOutput {
	return o
}

// IDs of the additional volumes in this image.
func (o LookupInstanceImageResultOutput) AdditionalVolumeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) []string { return v.AdditionalVolumeIds }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceImageResultOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) *string { return v.Architecture }).(pulumi.StringPtrOutput)
}

// Date of the image creation.
func (o LookupInstanceImageResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// ID of the default bootscript for this image.
func (o LookupInstanceImageResultOutput) DefaultBootscriptId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.DefaultBootscriptId }).(pulumi.StringOutput)
}

// ID of the server the image if based from.
func (o LookupInstanceImageResultOutput) FromServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.FromServerId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInstanceImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceImageResultOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceImageResultOutput) Latest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) *bool { return v.Latest }).(pulumi.BoolPtrOutput)
}

// Date of image latest update.
func (o LookupInstanceImageResultOutput) ModificationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.ModificationDate }).(pulumi.StringOutput)
}

func (o LookupInstanceImageResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the organization the image is associated with.
func (o LookupInstanceImageResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the project the image is associated with.
func (o LookupInstanceImageResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Set to `true` if the image is public.
func (o LookupInstanceImageResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// ID of the root volume in this image.
func (o LookupInstanceImageResultOutput) RootVolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.RootVolumeId }).(pulumi.StringOutput)
}

// State of the image. Possible values are: `available`, `creating` or `error`.
func (o LookupInstanceImageResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupInstanceImageResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceImageResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceImageResultOutput{})
}
