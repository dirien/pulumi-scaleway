// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway VPC Private Networks.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc/api/#private-networks-ac2df4).
//
// ## Example
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
//			_, err := scaleway.NewVpcPrivateNetwork(ctx, "pnPriv", &scaleway.VpcPrivateNetworkArgs{
//				Tags: pulumi.StringArray{
//					pulumi.String("demo"),
//					pulumi.String("terraform"),
//				},
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
// Private networks can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork vpc_demo fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type VpcPrivateNetwork struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the private network
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The name of the private network. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The organization ID the private network is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The tags associated with the private network.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the private network
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the private network should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPrivateNetwork registers a new resource with the given unique name, arguments, and options.
func NewVpcPrivateNetwork(ctx *pulumi.Context,
	name string, args *VpcPrivateNetworkArgs, opts ...pulumi.ResourceOption) (*VpcPrivateNetwork, error) {
	if args == nil {
		args = &VpcPrivateNetworkArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource VpcPrivateNetwork
	err := ctx.RegisterResource("scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPrivateNetwork gets an existing VpcPrivateNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPrivateNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPrivateNetworkState, opts ...pulumi.ResourceOption) (*VpcPrivateNetwork, error) {
	var resource VpcPrivateNetwork
	err := ctx.ReadResource("scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPrivateNetwork resources.
type vpcPrivateNetworkState struct {
	// The date and time of the creation of the private network
	CreatedAt *string `pulumi:"createdAt"`
	// The name of the private network. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// The organization ID the private network is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the private network.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the private network
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the private network should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPrivateNetworkState struct {
	// The date and time of the creation of the private network
	CreatedAt pulumi.StringPtrInput
	// The name of the private network. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// The organization ID the private network is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId pulumi.StringPtrInput
	// The tags associated with the private network.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the private network
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the private network should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPrivateNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPrivateNetworkState)(nil)).Elem()
}

type vpcPrivateNetworkArgs struct {
	// The name of the private network. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the private network.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the private network should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPrivateNetwork resource.
type VpcPrivateNetworkArgs struct {
	// The name of the private network. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the private network is associated with.
	ProjectId pulumi.StringPtrInput
	// The tags associated with the private network.
	Tags pulumi.StringArrayInput
	// `zone`) The zone in which the private network should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPrivateNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPrivateNetworkArgs)(nil)).Elem()
}

type VpcPrivateNetworkInput interface {
	pulumi.Input

	ToVpcPrivateNetworkOutput() VpcPrivateNetworkOutput
	ToVpcPrivateNetworkOutputWithContext(ctx context.Context) VpcPrivateNetworkOutput
}

func (*VpcPrivateNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPrivateNetwork)(nil)).Elem()
}

func (i *VpcPrivateNetwork) ToVpcPrivateNetworkOutput() VpcPrivateNetworkOutput {
	return i.ToVpcPrivateNetworkOutputWithContext(context.Background())
}

func (i *VpcPrivateNetwork) ToVpcPrivateNetworkOutputWithContext(ctx context.Context) VpcPrivateNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPrivateNetworkOutput)
}

// VpcPrivateNetworkArrayInput is an input type that accepts VpcPrivateNetworkArray and VpcPrivateNetworkArrayOutput values.
// You can construct a concrete instance of `VpcPrivateNetworkArrayInput` via:
//
//	VpcPrivateNetworkArray{ VpcPrivateNetworkArgs{...} }
type VpcPrivateNetworkArrayInput interface {
	pulumi.Input

	ToVpcPrivateNetworkArrayOutput() VpcPrivateNetworkArrayOutput
	ToVpcPrivateNetworkArrayOutputWithContext(context.Context) VpcPrivateNetworkArrayOutput
}

type VpcPrivateNetworkArray []VpcPrivateNetworkInput

func (VpcPrivateNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPrivateNetwork)(nil)).Elem()
}

func (i VpcPrivateNetworkArray) ToVpcPrivateNetworkArrayOutput() VpcPrivateNetworkArrayOutput {
	return i.ToVpcPrivateNetworkArrayOutputWithContext(context.Background())
}

func (i VpcPrivateNetworkArray) ToVpcPrivateNetworkArrayOutputWithContext(ctx context.Context) VpcPrivateNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPrivateNetworkArrayOutput)
}

// VpcPrivateNetworkMapInput is an input type that accepts VpcPrivateNetworkMap and VpcPrivateNetworkMapOutput values.
// You can construct a concrete instance of `VpcPrivateNetworkMapInput` via:
//
//	VpcPrivateNetworkMap{ "key": VpcPrivateNetworkArgs{...} }
type VpcPrivateNetworkMapInput interface {
	pulumi.Input

	ToVpcPrivateNetworkMapOutput() VpcPrivateNetworkMapOutput
	ToVpcPrivateNetworkMapOutputWithContext(context.Context) VpcPrivateNetworkMapOutput
}

type VpcPrivateNetworkMap map[string]VpcPrivateNetworkInput

func (VpcPrivateNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPrivateNetwork)(nil)).Elem()
}

func (i VpcPrivateNetworkMap) ToVpcPrivateNetworkMapOutput() VpcPrivateNetworkMapOutput {
	return i.ToVpcPrivateNetworkMapOutputWithContext(context.Background())
}

func (i VpcPrivateNetworkMap) ToVpcPrivateNetworkMapOutputWithContext(ctx context.Context) VpcPrivateNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPrivateNetworkMapOutput)
}

type VpcPrivateNetworkOutput struct{ *pulumi.OutputState }

func (VpcPrivateNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPrivateNetwork)(nil)).Elem()
}

func (o VpcPrivateNetworkOutput) ToVpcPrivateNetworkOutput() VpcPrivateNetworkOutput {
	return o
}

func (o VpcPrivateNetworkOutput) ToVpcPrivateNetworkOutputWithContext(ctx context.Context) VpcPrivateNetworkOutput {
	return o
}

// The date and time of the creation of the private network
func (o VpcPrivateNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The name of the private network. If not provided it will be randomly generated.
func (o VpcPrivateNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The organization ID the private network is associated with.
func (o VpcPrivateNetworkOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the private network is associated with.
func (o VpcPrivateNetworkOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The tags associated with the private network.
func (o VpcPrivateNetworkOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the private network
func (o VpcPrivateNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the private network should be created.
func (o VpcPrivateNetworkOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPrivateNetworkArrayOutput struct{ *pulumi.OutputState }

func (VpcPrivateNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPrivateNetwork)(nil)).Elem()
}

func (o VpcPrivateNetworkArrayOutput) ToVpcPrivateNetworkArrayOutput() VpcPrivateNetworkArrayOutput {
	return o
}

func (o VpcPrivateNetworkArrayOutput) ToVpcPrivateNetworkArrayOutputWithContext(ctx context.Context) VpcPrivateNetworkArrayOutput {
	return o
}

func (o VpcPrivateNetworkArrayOutput) Index(i pulumi.IntInput) VpcPrivateNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPrivateNetwork {
		return vs[0].([]*VpcPrivateNetwork)[vs[1].(int)]
	}).(VpcPrivateNetworkOutput)
}

type VpcPrivateNetworkMapOutput struct{ *pulumi.OutputState }

func (VpcPrivateNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPrivateNetwork)(nil)).Elem()
}

func (o VpcPrivateNetworkMapOutput) ToVpcPrivateNetworkMapOutput() VpcPrivateNetworkMapOutput {
	return o
}

func (o VpcPrivateNetworkMapOutput) ToVpcPrivateNetworkMapOutputWithContext(ctx context.Context) VpcPrivateNetworkMapOutput {
	return o
}

func (o VpcPrivateNetworkMapOutput) MapIndex(k pulumi.StringInput) VpcPrivateNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPrivateNetwork {
		return vs[0].(map[string]*VpcPrivateNetwork)[vs[1].(string)]
	}).(VpcPrivateNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPrivateNetworkInput)(nil)).Elem(), &VpcPrivateNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPrivateNetworkArrayInput)(nil)).Elem(), VpcPrivateNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPrivateNetworkMapInput)(nil)).Elem(), VpcPrivateNetworkMap{})
	pulumi.RegisterOutputType(VpcPrivateNetworkOutput{})
	pulumi.RegisterOutputType(VpcPrivateNetworkArrayOutput{})
	pulumi.RegisterOutputType(VpcPrivateNetworkMapOutput{})
}
