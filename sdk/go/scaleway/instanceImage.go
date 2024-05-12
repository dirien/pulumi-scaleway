// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Compute Images.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/instance/#images-41389b).
//
// ## Example Usage
//
// ### From a volume
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
//			volume, err := scaleway.NewInstanceVolume(ctx, "volume", &scaleway.InstanceVolumeArgs{
//				Type:     pulumi.String("b_ssd"),
//				SizeInGb: pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			volumeSnapshot, err := scaleway.NewInstanceSnapshot(ctx, "volumeSnapshot", &scaleway.InstanceSnapshotArgs{
//				VolumeId: volume.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewInstanceImage(ctx, "volumeImage", &scaleway.InstanceImageArgs{
//				RootVolumeId: volumeSnapshot.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### From a server
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
//			_, err := scaleway.NewInstanceServer(ctx, "server", &scaleway.InstanceServerArgs{
//				Image: pulumi.String("ubuntu_jammy"),
//				Type:  pulumi.String("DEV1-S"),
//			})
//			if err != nil {
//				return err
//			}
//			serverSnapshot, err := scaleway.NewInstanceSnapshot(ctx, "serverSnapshot", &scaleway.InstanceSnapshotArgs{
//				VolumeId: pulumi.Any(scaleway_instance_server.Main.Root_volume[0].Volume_id),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewInstanceImage(ctx, "serverImage", &scaleway.InstanceImageArgs{
//				RootVolumeId: serverSnapshot.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Images can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/instanceImage:InstanceImage main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type InstanceImage struct {
	pulumi.CustomResourceState

	// List of IDs of the snapshots of the additional volumes to be attached to the image.
	//
	// > **Important:** For now it is only possible to have 1 additional_volume.
	AdditionalVolumeIds pulumi.StringPtrOutput `pulumi:"additionalVolumeIds"`
	// The description of the extra volumes attached to the image.
	AdditionalVolumes InstanceImageAdditionalVolumeArrayOutput `pulumi:"additionalVolumes"`
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture pulumi.StringPtrOutput `pulumi:"architecture"`
	// Date of the volume creation.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// ID of the server the image is based on (in case it is a backup).
	FromServerId pulumi.StringOutput `pulumi:"fromServerId"`
	// Date of volume latest update.
	ModificationDate pulumi.StringOutput `pulumi:"modificationDate"`
	// The name of the image. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the image is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The ID of the project the image is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Set to `true` if the image is public.
	Public pulumi.BoolPtrOutput `pulumi:"public"`
	// The ID of the snapshot of the volume to be used as root in the image.
	RootVolumeId pulumi.StringOutput `pulumi:"rootVolumeId"`
	// State of the volume.
	State pulumi.StringOutput `pulumi:"state"`
	// A list of tags to apply to the image.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The zone in which the image should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceImage registers a new resource with the given unique name, arguments, and options.
func NewInstanceImage(ctx *pulumi.Context,
	name string, args *InstanceImageArgs, opts ...pulumi.ResourceOption) (*InstanceImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RootVolumeId == nil {
		return nil, errors.New("invalid value for required argument 'RootVolumeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceImage
	err := ctx.RegisterResource("scaleway:index/instanceImage:InstanceImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceImage gets an existing InstanceImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceImageState, opts ...pulumi.ResourceOption) (*InstanceImage, error) {
	var resource InstanceImage
	err := ctx.ReadResource("scaleway:index/instanceImage:InstanceImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceImage resources.
type instanceImageState struct {
	// List of IDs of the snapshots of the additional volumes to be attached to the image.
	//
	// > **Important:** For now it is only possible to have 1 additional_volume.
	AdditionalVolumeIds *string `pulumi:"additionalVolumeIds"`
	// The description of the extra volumes attached to the image.
	AdditionalVolumes []InstanceImageAdditionalVolume `pulumi:"additionalVolumes"`
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture *string `pulumi:"architecture"`
	// Date of the volume creation.
	CreationDate *string `pulumi:"creationDate"`
	// ID of the server the image is based on (in case it is a backup).
	FromServerId *string `pulumi:"fromServerId"`
	// Date of volume latest update.
	ModificationDate *string `pulumi:"modificationDate"`
	// The name of the image. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the image is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The ID of the project the image is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Set to `true` if the image is public.
	Public *bool `pulumi:"public"`
	// The ID of the snapshot of the volume to be used as root in the image.
	RootVolumeId *string `pulumi:"rootVolumeId"`
	// State of the volume.
	State *string `pulumi:"state"`
	// A list of tags to apply to the image.
	Tags []string `pulumi:"tags"`
	// The zone in which the image should be created.
	Zone *string `pulumi:"zone"`
}

type InstanceImageState struct {
	// List of IDs of the snapshots of the additional volumes to be attached to the image.
	//
	// > **Important:** For now it is only possible to have 1 additional_volume.
	AdditionalVolumeIds pulumi.StringPtrInput
	// The description of the extra volumes attached to the image.
	AdditionalVolumes InstanceImageAdditionalVolumeArrayInput
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture pulumi.StringPtrInput
	// Date of the volume creation.
	CreationDate pulumi.StringPtrInput
	// ID of the server the image is based on (in case it is a backup).
	FromServerId pulumi.StringPtrInput
	// Date of volume latest update.
	ModificationDate pulumi.StringPtrInput
	// The name of the image. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the image is associated with.
	OrganizationId pulumi.StringPtrInput
	// The ID of the project the image is associated with.
	ProjectId pulumi.StringPtrInput
	// Set to `true` if the image is public.
	Public pulumi.BoolPtrInput
	// The ID of the snapshot of the volume to be used as root in the image.
	RootVolumeId pulumi.StringPtrInput
	// State of the volume.
	State pulumi.StringPtrInput
	// A list of tags to apply to the image.
	Tags pulumi.StringArrayInput
	// The zone in which the image should be created.
	Zone pulumi.StringPtrInput
}

func (InstanceImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceImageState)(nil)).Elem()
}

type instanceImageArgs struct {
	// List of IDs of the snapshots of the additional volumes to be attached to the image.
	//
	// > **Important:** For now it is only possible to have 1 additional_volume.
	AdditionalVolumeIds *string `pulumi:"additionalVolumeIds"`
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture *string `pulumi:"architecture"`
	// The name of the image. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The ID of the project the image is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Set to `true` if the image is public.
	Public *bool `pulumi:"public"`
	// The ID of the snapshot of the volume to be used as root in the image.
	RootVolumeId string `pulumi:"rootVolumeId"`
	// A list of tags to apply to the image.
	Tags []string `pulumi:"tags"`
	// The zone in which the image should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceImage resource.
type InstanceImageArgs struct {
	// List of IDs of the snapshots of the additional volumes to be attached to the image.
	//
	// > **Important:** For now it is only possible to have 1 additional_volume.
	AdditionalVolumeIds pulumi.StringPtrInput
	// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
	Architecture pulumi.StringPtrInput
	// The name of the image. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The ID of the project the image is associated with.
	ProjectId pulumi.StringPtrInput
	// Set to `true` if the image is public.
	Public pulumi.BoolPtrInput
	// The ID of the snapshot of the volume to be used as root in the image.
	RootVolumeId pulumi.StringInput
	// A list of tags to apply to the image.
	Tags pulumi.StringArrayInput
	// The zone in which the image should be created.
	Zone pulumi.StringPtrInput
}

func (InstanceImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceImageArgs)(nil)).Elem()
}

type InstanceImageInput interface {
	pulumi.Input

	ToInstanceImageOutput() InstanceImageOutput
	ToInstanceImageOutputWithContext(ctx context.Context) InstanceImageOutput
}

func (*InstanceImage) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceImage)(nil)).Elem()
}

func (i *InstanceImage) ToInstanceImageOutput() InstanceImageOutput {
	return i.ToInstanceImageOutputWithContext(context.Background())
}

func (i *InstanceImage) ToInstanceImageOutputWithContext(ctx context.Context) InstanceImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceImageOutput)
}

// InstanceImageArrayInput is an input type that accepts InstanceImageArray and InstanceImageArrayOutput values.
// You can construct a concrete instance of `InstanceImageArrayInput` via:
//
//	InstanceImageArray{ InstanceImageArgs{...} }
type InstanceImageArrayInput interface {
	pulumi.Input

	ToInstanceImageArrayOutput() InstanceImageArrayOutput
	ToInstanceImageArrayOutputWithContext(context.Context) InstanceImageArrayOutput
}

type InstanceImageArray []InstanceImageInput

func (InstanceImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceImage)(nil)).Elem()
}

func (i InstanceImageArray) ToInstanceImageArrayOutput() InstanceImageArrayOutput {
	return i.ToInstanceImageArrayOutputWithContext(context.Background())
}

func (i InstanceImageArray) ToInstanceImageArrayOutputWithContext(ctx context.Context) InstanceImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceImageArrayOutput)
}

// InstanceImageMapInput is an input type that accepts InstanceImageMap and InstanceImageMapOutput values.
// You can construct a concrete instance of `InstanceImageMapInput` via:
//
//	InstanceImageMap{ "key": InstanceImageArgs{...} }
type InstanceImageMapInput interface {
	pulumi.Input

	ToInstanceImageMapOutput() InstanceImageMapOutput
	ToInstanceImageMapOutputWithContext(context.Context) InstanceImageMapOutput
}

type InstanceImageMap map[string]InstanceImageInput

func (InstanceImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceImage)(nil)).Elem()
}

func (i InstanceImageMap) ToInstanceImageMapOutput() InstanceImageMapOutput {
	return i.ToInstanceImageMapOutputWithContext(context.Background())
}

func (i InstanceImageMap) ToInstanceImageMapOutputWithContext(ctx context.Context) InstanceImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceImageMapOutput)
}

type InstanceImageOutput struct{ *pulumi.OutputState }

func (InstanceImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceImage)(nil)).Elem()
}

func (o InstanceImageOutput) ToInstanceImageOutput() InstanceImageOutput {
	return o
}

func (o InstanceImageOutput) ToInstanceImageOutputWithContext(ctx context.Context) InstanceImageOutput {
	return o
}

// List of IDs of the snapshots of the additional volumes to be attached to the image.
//
// > **Important:** For now it is only possible to have 1 additional_volume.
func (o InstanceImageOutput) AdditionalVolumeIds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringPtrOutput { return v.AdditionalVolumeIds }).(pulumi.StringPtrOutput)
}

// The description of the extra volumes attached to the image.
func (o InstanceImageOutput) AdditionalVolumes() InstanceImageAdditionalVolumeArrayOutput {
	return o.ApplyT(func(v *InstanceImage) InstanceImageAdditionalVolumeArrayOutput { return v.AdditionalVolumes }).(InstanceImageAdditionalVolumeArrayOutput)
}

// The architecture the image is compatible with. Possible values are: `x8664` or `arm`.
func (o InstanceImageOutput) Architecture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringPtrOutput { return v.Architecture }).(pulumi.StringPtrOutput)
}

// Date of the volume creation.
func (o InstanceImageOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// ID of the server the image is based on (in case it is a backup).
func (o InstanceImageOutput) FromServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.FromServerId }).(pulumi.StringOutput)
}

// Date of volume latest update.
func (o InstanceImageOutput) ModificationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.ModificationDate }).(pulumi.StringOutput)
}

// The name of the image. If not provided it will be randomly generated.
func (o InstanceImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the image is associated with.
func (o InstanceImageOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the project the image is associated with.
func (o InstanceImageOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Set to `true` if the image is public.
func (o InstanceImageOutput) Public() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.BoolPtrOutput { return v.Public }).(pulumi.BoolPtrOutput)
}

// The ID of the snapshot of the volume to be used as root in the image.
func (o InstanceImageOutput) RootVolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.RootVolumeId }).(pulumi.StringOutput)
}

// State of the volume.
func (o InstanceImageOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A list of tags to apply to the image.
func (o InstanceImageOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone in which the image should be created.
func (o InstanceImageOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceImage) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceImageArrayOutput struct{ *pulumi.OutputState }

func (InstanceImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceImage)(nil)).Elem()
}

func (o InstanceImageArrayOutput) ToInstanceImageArrayOutput() InstanceImageArrayOutput {
	return o
}

func (o InstanceImageArrayOutput) ToInstanceImageArrayOutputWithContext(ctx context.Context) InstanceImageArrayOutput {
	return o
}

func (o InstanceImageArrayOutput) Index(i pulumi.IntInput) InstanceImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceImage {
		return vs[0].([]*InstanceImage)[vs[1].(int)]
	}).(InstanceImageOutput)
}

type InstanceImageMapOutput struct{ *pulumi.OutputState }

func (InstanceImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceImage)(nil)).Elem()
}

func (o InstanceImageMapOutput) ToInstanceImageMapOutput() InstanceImageMapOutput {
	return o
}

func (o InstanceImageMapOutput) ToInstanceImageMapOutputWithContext(ctx context.Context) InstanceImageMapOutput {
	return o
}

func (o InstanceImageMapOutput) MapIndex(k pulumi.StringInput) InstanceImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceImage {
		return vs[0].(map[string]*InstanceImage)[vs[1].(string)]
	}).(InstanceImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceImageInput)(nil)).Elem(), &InstanceImage{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceImageArrayInput)(nil)).Elem(), InstanceImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceImageMapInput)(nil)).Elem(), InstanceImageMap{})
	pulumi.RegisterOutputType(InstanceImageOutput{})
	pulumi.RegisterOutputType(InstanceImageArrayOutput{})
	pulumi.RegisterOutputType(InstanceImageMapOutput{})
}
