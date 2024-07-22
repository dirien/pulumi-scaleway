// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway VPC Private Networks.
// For more information, see [the API documentation](https://www.scaleway.com/en/developers/api/vpc/#private-networks-ac2df4).
//
// ## Example Usage
//
// ### Basic
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
// ### With subnets
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
//				Ipv4Subnet: &scaleway.VpcPrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("192.168.0.0/24"),
//				},
//				Ipv6Subnets: scaleway.VpcPrivateNetworkIpv6SubnetArray{
//					&scaleway.VpcPrivateNetworkIpv6SubnetArgs{
//						Subnet: pulumi.String("fd46:78ab:30b8:177c::/64"),
//					},
//					&scaleway.VpcPrivateNetworkIpv6SubnetArgs{
//						Subnet: pulumi.String("fd46:78ab:30b8:c7df::/64"),
//					},
//				},
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
// Private Networks can be imported using `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork vpc_demo fr-par/11111111-1111-1111-1111-111111111111
// ```
type VpcPrivateNetwork struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the subnet.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The IPv4 subnet to associate with the Private Network.
	Ipv4Subnet VpcPrivateNetworkIpv4SubnetOutput `pulumi:"ipv4Subnet"`
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets VpcPrivateNetworkIpv6SubnetArrayOutput `pulumi:"ipv6Subnets"`
	// Private Networks are now all necessarily regional.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional pulumi.BoolOutput `pulumi:"isRegional"`
	// The name of the Private Network. If not provided, it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Organization ID the Private Network is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The ID of the Project the private network is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region of the Private Network.
	Region pulumi.StringOutput `pulumi:"region"`
	// The tags associated with the Private Network.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the subnet.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The VPC in which to create the Private Network.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Use `region` instead.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPrivateNetwork registers a new resource with the given unique name, arguments, and options.
func NewVpcPrivateNetwork(ctx *pulumi.Context,
	name string, args *VpcPrivateNetworkArgs, opts ...pulumi.ResourceOption) (*VpcPrivateNetwork, error) {
	if args == nil {
		args = &VpcPrivateNetworkArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
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
	// The date and time of the creation of the subnet.
	CreatedAt *string `pulumi:"createdAt"`
	// The IPv4 subnet to associate with the Private Network.
	Ipv4Subnet *VpcPrivateNetworkIpv4Subnet `pulumi:"ipv4Subnet"`
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets []VpcPrivateNetworkIpv6Subnet `pulumi:"ipv6Subnets"`
	// Private Networks are now all necessarily regional.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional *bool `pulumi:"isRegional"`
	// The name of the Private Network. If not provided, it will be randomly generated.
	Name *string `pulumi:"name"`
	// The Organization ID the Private Network is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the Project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the Private Network.
	Region *string `pulumi:"region"`
	// The tags associated with the Private Network.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the subnet.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The VPC in which to create the Private Network.
	VpcId *string `pulumi:"vpcId"`
	// Use `region` instead.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone *string `pulumi:"zone"`
}

type VpcPrivateNetworkState struct {
	// The date and time of the creation of the subnet.
	CreatedAt pulumi.StringPtrInput
	// The IPv4 subnet to associate with the Private Network.
	Ipv4Subnet VpcPrivateNetworkIpv4SubnetPtrInput
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets VpcPrivateNetworkIpv6SubnetArrayInput
	// Private Networks are now all necessarily regional.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional pulumi.BoolPtrInput
	// The name of the Private Network. If not provided, it will be randomly generated.
	Name pulumi.StringPtrInput
	// The Organization ID the Private Network is associated with.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The ID of the Project the private network is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the Private Network.
	Region pulumi.StringPtrInput
	// The tags associated with the Private Network.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the subnet.
	UpdatedAt pulumi.StringPtrInput
	// The VPC in which to create the Private Network.
	VpcId pulumi.StringPtrInput
	// Use `region` instead.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone pulumi.StringPtrInput
}

func (VpcPrivateNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPrivateNetworkState)(nil)).Elem()
}

type vpcPrivateNetworkArgs struct {
	// The IPv4 subnet to associate with the Private Network.
	Ipv4Subnet *VpcPrivateNetworkIpv4Subnet `pulumi:"ipv4Subnet"`
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets []VpcPrivateNetworkIpv6Subnet `pulumi:"ipv6Subnets"`
	// Private Networks are now all necessarily regional.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional *bool `pulumi:"isRegional"`
	// The name of the Private Network. If not provided, it will be randomly generated.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the Project the private network is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the Private Network.
	Region *string `pulumi:"region"`
	// The tags associated with the Private Network.
	Tags []string `pulumi:"tags"`
	// The VPC in which to create the Private Network.
	VpcId *string `pulumi:"vpcId"`
	// Use `region` instead.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPrivateNetwork resource.
type VpcPrivateNetworkArgs struct {
	// The IPv4 subnet to associate with the Private Network.
	Ipv4Subnet VpcPrivateNetworkIpv4SubnetPtrInput
	// The IPv6 subnets to associate with the private network.
	Ipv6Subnets VpcPrivateNetworkIpv6SubnetArrayInput
	// Private Networks are now all necessarily regional.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version
	IsRegional pulumi.BoolPtrInput
	// The name of the Private Network. If not provided, it will be randomly generated.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the Project the private network is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the Private Network.
	Region pulumi.StringPtrInput
	// The tags associated with the Private Network.
	Tags pulumi.StringArrayInput
	// The VPC in which to create the Private Network.
	VpcId pulumi.StringPtrInput
	// Use `region` instead.
	//
	// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
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

// The date and time of the creation of the subnet.
func (o VpcPrivateNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The IPv4 subnet to associate with the Private Network.
func (o VpcPrivateNetworkOutput) Ipv4Subnet() VpcPrivateNetworkIpv4SubnetOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) VpcPrivateNetworkIpv4SubnetOutput { return v.Ipv4Subnet }).(VpcPrivateNetworkIpv4SubnetOutput)
}

// The IPv6 subnets to associate with the private network.
func (o VpcPrivateNetworkOutput) Ipv6Subnets() VpcPrivateNetworkIpv6SubnetArrayOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) VpcPrivateNetworkIpv6SubnetArrayOutput { return v.Ipv6Subnets }).(VpcPrivateNetworkIpv6SubnetArrayOutput)
}

// Private Networks are now all necessarily regional.
//
// Deprecated: This field is deprecated and will be removed in the next major version
func (o VpcPrivateNetworkOutput) IsRegional() pulumi.BoolOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.BoolOutput { return v.IsRegional }).(pulumi.BoolOutput)
}

// The name of the Private Network. If not provided, it will be randomly generated.
func (o VpcPrivateNetworkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID the Private Network is associated with.
func (o VpcPrivateNetworkOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The ID of the Project the private network is associated with.
func (o VpcPrivateNetworkOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region of the Private Network.
func (o VpcPrivateNetworkOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The tags associated with the Private Network.
func (o VpcPrivateNetworkOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the subnet.
func (o VpcPrivateNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The VPC in which to create the Private Network.
func (o VpcPrivateNetworkOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPrivateNetwork) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Use `region` instead.
//
// Deprecated: This field is deprecated and will be removed in the next major version, please use `region` instead
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
