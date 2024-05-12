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

// Creates and manages Scaleway VPC Public Gateway Network.
// It allows attaching Private Networks to the VPC Public Gateway and your DHCP config
// For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#step-3-attach-private-networks-to-the-vpc-public-gateway).
//
// ## Example Usage
//
// ### Create a gateway network with IPAM config
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
//			vpc01, err := scaleway.NewVpc(ctx, "vpc01", nil)
//			if err != nil {
//				return err
//			}
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", &scaleway.VpcPrivateNetworkArgs{
//				Ipv4Subnet: &scaleway.VpcPrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.64.0/22"),
//				},
//				VpcId: vpc01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			pg01, err := scaleway.NewVpcPublicGateway(ctx, "pg01", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewVpcGatewayNetwork(ctx, "main", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				EnableMasquerade: pulumi.Bool(true),
//				IpamConfigs: scaleway.VpcGatewayNetworkIpamConfigArray{
//					&scaleway.VpcGatewayNetworkIpamConfigArgs{
//						PushDefaultRoute: pulumi.Bool(true),
//					},
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
// ### Create a gateway network with a booked IPAM IP
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
//			vpc01, err := scaleway.NewVpc(ctx, "vpc01", nil)
//			if err != nil {
//				return err
//			}
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", &scaleway.VpcPrivateNetworkArgs{
//				Ipv4Subnet: &scaleway.VpcPrivateNetworkIpv4SubnetArgs{
//					Subnet: pulumi.String("172.16.64.0/22"),
//				},
//				VpcId: vpc01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			ip01, err := scaleway.NewIpamIp(ctx, "ip01", &scaleway.IpamIpArgs{
//				Address: pulumi.String("172.16.64.7"),
//				Sources: scaleway.IpamIpSourceArray{
//					&scaleway.IpamIpSourceArgs{
//						PrivateNetworkId: pn01.ID(),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			pg01, err := scaleway.NewVpcPublicGateway(ctx, "pg01", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewVpcGatewayNetwork(ctx, "main", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				EnableMasquerade: pulumi.Bool(true),
//				IpamConfigs: scaleway.VpcGatewayNetworkIpamConfigArray{
//					&scaleway.VpcGatewayNetworkIpamConfigArgs{
//						PushDefaultRoute: pulumi.Bool(true),
//						IpamIpId:         ip01.ID(),
//					},
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
// ### Create a gateway network with DHCP
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
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", nil)
//			if err != nil {
//				return err
//			}
//			gw01, err := scaleway.NewVpcPublicGatewayIp(ctx, "gw01", nil)
//			if err != nil {
//				return err
//			}
//			dhcp01, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "dhcp01", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet:           pulumi.String("192.168.1.0/24"),
//				PushDefaultRoute: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			pg01, err := scaleway.NewVpcPublicGateway(ctx, "pg01", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: gw01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewVpcGatewayNetwork(ctx, "main", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				DhcpId:           dhcp01.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
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
// ### Create a gateway network with a static IP address
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
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", nil)
//			if err != nil {
//				return err
//			}
//			pg01, err := scaleway.NewVpcPublicGateway(ctx, "pg01", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewVpcGatewayNetwork(ctx, "main", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				EnableDhcp:       pulumi.Bool(false),
//				EnableMasquerade: pulumi.Bool(true),
//				StaticAddress:    pulumi.String("192.168.1.42/24"),
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
// Gateway network can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcGatewayNetwork struct {
	pulumi.CustomResourceState

	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp pulumi.BoolPtrOutput `pulumi:"cleanupDhcp"`
	// The date and time of the creation of the gateway network.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId pulumi.StringPtrOutput `pulumi:"dhcpId"`
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp pulumi.BoolPtrOutput `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrOutput `pulumi:"enableMasquerade"`
	// The ID of the public gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs VpcGatewayNetworkIpamConfigArrayOutput `pulumi:"ipamConfigs"`
	// The mac address of the creation of the gateway network.
	MacAddress pulumi.StringOutput `pulumi:"macAddress"`
	// The ID of the private network.
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress pulumi.StringOutput `pulumi:"staticAddress"`
	// The status of the Public Gateway's connection to the Private Network.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date and time of the last update of the gateway network.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the gateway network should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcGatewayNetwork registers a new resource with the given unique name, arguments, and options.
func NewVpcGatewayNetwork(ctx *pulumi.Context,
	name string, args *VpcGatewayNetworkArgs, opts ...pulumi.ResourceOption) (*VpcGatewayNetwork, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.PrivateNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateNetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcGatewayNetwork
	err := ctx.RegisterResource("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcGatewayNetwork gets an existing VpcGatewayNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcGatewayNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcGatewayNetworkState, opts ...pulumi.ResourceOption) (*VpcGatewayNetwork, error) {
	var resource VpcGatewayNetwork
	err := ctx.ReadResource("scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcGatewayNetwork resources.
type vpcGatewayNetworkState struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp *bool `pulumi:"cleanupDhcp"`
	// The date and time of the creation of the gateway network.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId *string `pulumi:"dhcpId"`
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// The ID of the public gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs []VpcGatewayNetworkIpamConfig `pulumi:"ipamConfigs"`
	// The mac address of the creation of the gateway network.
	MacAddress *string `pulumi:"macAddress"`
	// The ID of the private network.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress *string `pulumi:"staticAddress"`
	// The status of the Public Gateway's connection to the Private Network.
	Status *string `pulumi:"status"`
	// The date and time of the last update of the gateway network.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the gateway network should be created.
	Zone *string `pulumi:"zone"`
}

type VpcGatewayNetworkState struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp pulumi.BoolPtrInput
	// The date and time of the creation of the gateway network.
	CreatedAt pulumi.StringPtrInput
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId pulumi.StringPtrInput
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp pulumi.BoolPtrInput
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrInput
	// The ID of the public gateway.
	GatewayId pulumi.StringPtrInput
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs VpcGatewayNetworkIpamConfigArrayInput
	// The mac address of the creation of the gateway network.
	MacAddress pulumi.StringPtrInput
	// The ID of the private network.
	PrivateNetworkId pulumi.StringPtrInput
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress pulumi.StringPtrInput
	// The status of the Public Gateway's connection to the Private Network.
	Status pulumi.StringPtrInput
	// The date and time of the last update of the gateway network.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the gateway network should be created.
	Zone pulumi.StringPtrInput
}

func (VpcGatewayNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGatewayNetworkState)(nil)).Elem()
}

type vpcGatewayNetworkArgs struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp *bool `pulumi:"cleanupDhcp"`
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId *string `pulumi:"dhcpId"`
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp *bool `pulumi:"enableDhcp"`
	// Enable masquerade on this network
	EnableMasquerade *bool `pulumi:"enableMasquerade"`
	// The ID of the public gateway.
	GatewayId string `pulumi:"gatewayId"`
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs []VpcGatewayNetworkIpamConfig `pulumi:"ipamConfigs"`
	// The ID of the private network.
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress *string `pulumi:"staticAddress"`
	// `zone`) The zone in which the gateway network should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcGatewayNetwork resource.
type VpcGatewayNetworkArgs struct {
	// Remove DHCP config on this network on destroy. It requires DHCP id.
	CleanupDhcp pulumi.BoolPtrInput
	// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	DhcpId pulumi.StringPtrInput
	// Enable DHCP config on this network. It requires DHCP id.
	EnableDhcp pulumi.BoolPtrInput
	// Enable masquerade on this network
	EnableMasquerade pulumi.BoolPtrInput
	// The ID of the public gateway.
	GatewayId pulumi.StringInput
	// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	IpamConfigs VpcGatewayNetworkIpamConfigArrayInput
	// The ID of the private network.
	PrivateNetworkId pulumi.StringInput
	// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
	StaticAddress pulumi.StringPtrInput
	// `zone`) The zone in which the gateway network should be created.
	Zone pulumi.StringPtrInput
}

func (VpcGatewayNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcGatewayNetworkArgs)(nil)).Elem()
}

type VpcGatewayNetworkInput interface {
	pulumi.Input

	ToVpcGatewayNetworkOutput() VpcGatewayNetworkOutput
	ToVpcGatewayNetworkOutputWithContext(ctx context.Context) VpcGatewayNetworkOutput
}

func (*VpcGatewayNetwork) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGatewayNetwork)(nil)).Elem()
}

func (i *VpcGatewayNetwork) ToVpcGatewayNetworkOutput() VpcGatewayNetworkOutput {
	return i.ToVpcGatewayNetworkOutputWithContext(context.Background())
}

func (i *VpcGatewayNetwork) ToVpcGatewayNetworkOutputWithContext(ctx context.Context) VpcGatewayNetworkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGatewayNetworkOutput)
}

// VpcGatewayNetworkArrayInput is an input type that accepts VpcGatewayNetworkArray and VpcGatewayNetworkArrayOutput values.
// You can construct a concrete instance of `VpcGatewayNetworkArrayInput` via:
//
//	VpcGatewayNetworkArray{ VpcGatewayNetworkArgs{...} }
type VpcGatewayNetworkArrayInput interface {
	pulumi.Input

	ToVpcGatewayNetworkArrayOutput() VpcGatewayNetworkArrayOutput
	ToVpcGatewayNetworkArrayOutputWithContext(context.Context) VpcGatewayNetworkArrayOutput
}

type VpcGatewayNetworkArray []VpcGatewayNetworkInput

func (VpcGatewayNetworkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcGatewayNetwork)(nil)).Elem()
}

func (i VpcGatewayNetworkArray) ToVpcGatewayNetworkArrayOutput() VpcGatewayNetworkArrayOutput {
	return i.ToVpcGatewayNetworkArrayOutputWithContext(context.Background())
}

func (i VpcGatewayNetworkArray) ToVpcGatewayNetworkArrayOutputWithContext(ctx context.Context) VpcGatewayNetworkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGatewayNetworkArrayOutput)
}

// VpcGatewayNetworkMapInput is an input type that accepts VpcGatewayNetworkMap and VpcGatewayNetworkMapOutput values.
// You can construct a concrete instance of `VpcGatewayNetworkMapInput` via:
//
//	VpcGatewayNetworkMap{ "key": VpcGatewayNetworkArgs{...} }
type VpcGatewayNetworkMapInput interface {
	pulumi.Input

	ToVpcGatewayNetworkMapOutput() VpcGatewayNetworkMapOutput
	ToVpcGatewayNetworkMapOutputWithContext(context.Context) VpcGatewayNetworkMapOutput
}

type VpcGatewayNetworkMap map[string]VpcGatewayNetworkInput

func (VpcGatewayNetworkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcGatewayNetwork)(nil)).Elem()
}

func (i VpcGatewayNetworkMap) ToVpcGatewayNetworkMapOutput() VpcGatewayNetworkMapOutput {
	return i.ToVpcGatewayNetworkMapOutputWithContext(context.Background())
}

func (i VpcGatewayNetworkMap) ToVpcGatewayNetworkMapOutputWithContext(ctx context.Context) VpcGatewayNetworkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcGatewayNetworkMapOutput)
}

type VpcGatewayNetworkOutput struct{ *pulumi.OutputState }

func (VpcGatewayNetworkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcGatewayNetwork)(nil)).Elem()
}

func (o VpcGatewayNetworkOutput) ToVpcGatewayNetworkOutput() VpcGatewayNetworkOutput {
	return o
}

func (o VpcGatewayNetworkOutput) ToVpcGatewayNetworkOutputWithContext(ctx context.Context) VpcGatewayNetworkOutput {
	return o
}

// Remove DHCP config on this network on destroy. It requires DHCP id.
func (o VpcGatewayNetworkOutput) CleanupDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.BoolPtrOutput { return v.CleanupDhcp }).(pulumi.BoolPtrOutput)
}

// The date and time of the creation of the gateway network.
func (o VpcGatewayNetworkOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the public gateway DHCP config. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
func (o VpcGatewayNetworkOutput) DhcpId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringPtrOutput { return v.DhcpId }).(pulumi.StringPtrOutput)
}

// Enable DHCP config on this network. It requires DHCP id.
func (o VpcGatewayNetworkOutput) EnableDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.BoolPtrOutput { return v.EnableDhcp }).(pulumi.BoolPtrOutput)
}

// Enable masquerade on this network
func (o VpcGatewayNetworkOutput) EnableMasquerade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.BoolPtrOutput { return v.EnableMasquerade }).(pulumi.BoolPtrOutput)
}

// The ID of the public gateway.
func (o VpcGatewayNetworkOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// Auto-configure the Gateway Network using Scaleway's IPAM (IP address management service). Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
func (o VpcGatewayNetworkOutput) IpamConfigs() VpcGatewayNetworkIpamConfigArrayOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) VpcGatewayNetworkIpamConfigArrayOutput { return v.IpamConfigs }).(VpcGatewayNetworkIpamConfigArrayOutput)
}

// The mac address of the creation of the gateway network.
func (o VpcGatewayNetworkOutput) MacAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.MacAddress }).(pulumi.StringOutput)
}

// The ID of the private network.
func (o VpcGatewayNetworkOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// Enable DHCP config on this network. Only one of `dhcpId`, `staticAddress` and `ipamConfig` should be specified.
func (o VpcGatewayNetworkOutput) StaticAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.StaticAddress }).(pulumi.StringOutput)
}

// The status of the Public Gateway's connection to the Private Network.
func (o VpcGatewayNetworkOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date and time of the last update of the gateway network.
func (o VpcGatewayNetworkOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the gateway network should be created.
func (o VpcGatewayNetworkOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcGatewayNetwork) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcGatewayNetworkArrayOutput struct{ *pulumi.OutputState }

func (VpcGatewayNetworkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcGatewayNetwork)(nil)).Elem()
}

func (o VpcGatewayNetworkArrayOutput) ToVpcGatewayNetworkArrayOutput() VpcGatewayNetworkArrayOutput {
	return o
}

func (o VpcGatewayNetworkArrayOutput) ToVpcGatewayNetworkArrayOutputWithContext(ctx context.Context) VpcGatewayNetworkArrayOutput {
	return o
}

func (o VpcGatewayNetworkArrayOutput) Index(i pulumi.IntInput) VpcGatewayNetworkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcGatewayNetwork {
		return vs[0].([]*VpcGatewayNetwork)[vs[1].(int)]
	}).(VpcGatewayNetworkOutput)
}

type VpcGatewayNetworkMapOutput struct{ *pulumi.OutputState }

func (VpcGatewayNetworkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcGatewayNetwork)(nil)).Elem()
}

func (o VpcGatewayNetworkMapOutput) ToVpcGatewayNetworkMapOutput() VpcGatewayNetworkMapOutput {
	return o
}

func (o VpcGatewayNetworkMapOutput) ToVpcGatewayNetworkMapOutputWithContext(ctx context.Context) VpcGatewayNetworkMapOutput {
	return o
}

func (o VpcGatewayNetworkMapOutput) MapIndex(k pulumi.StringInput) VpcGatewayNetworkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcGatewayNetwork {
		return vs[0].(map[string]*VpcGatewayNetwork)[vs[1].(string)]
	}).(VpcGatewayNetworkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGatewayNetworkInput)(nil)).Elem(), &VpcGatewayNetwork{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGatewayNetworkArrayInput)(nil)).Elem(), VpcGatewayNetworkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcGatewayNetworkMapInput)(nil)).Elem(), VpcGatewayNetworkMap{})
	pulumi.RegisterOutputType(VpcGatewayNetworkOutput{})
	pulumi.RegisterOutputType(VpcGatewayNetworkArrayOutput{})
	pulumi.RegisterOutputType(VpcGatewayNetworkMapOutput{})
}
