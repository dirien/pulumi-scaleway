// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Public Gateway PAT (Port Address Translation).
// For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#pat-rules-e75d10).
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
//			sg01, err := scaleway.NewInstanceSecurityGroup(ctx, "sg01", &scaleway.InstanceSecurityGroupArgs{
//				InboundDefaultPolicy:  pulumi.String("drop"),
//				OutboundDefaultPolicy: pulumi.String("accept"),
//				InboundRules: scaleway.InstanceSecurityGroupInboundRuleArray{
//					&scaleway.InstanceSecurityGroupInboundRuleArgs{
//						Action:   pulumi.String("accept"),
//						Port:     pulumi.Int(22),
//						Protocol: pulumi.String("TCP"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			srv01, err := scaleway.NewInstanceServer(ctx, "srv01", &scaleway.InstanceServerArgs{
//				Type:            pulumi.String("PLAY2-NANO"),
//				Image:           pulumi.String("ubuntu_jammy"),
//				SecurityGroupId: sg01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			pn01, err := scaleway.NewVpcPrivateNetwork(ctx, "pn01", nil)
//			if err != nil {
//				return err
//			}
//			pnic01, err := scaleway.NewInstancePrivateNic(ctx, "pnic01", &scaleway.InstancePrivateNicArgs{
//				ServerId:         srv01.ID(),
//				PrivateNetworkId: pn01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			dhcp01, err := scaleway.NewVpcPublicGatewayDhcp(ctx, "dhcp01", &scaleway.VpcPublicGatewayDhcpArgs{
//				Subnet: pulumi.String("192.168.0.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			ip01, err := scaleway.NewVpcPublicGatewayIp(ctx, "ip01", nil)
//			if err != nil {
//				return err
//			}
//			pg01, err := scaleway.NewVpcPublicGateway(ctx, "pg01", &scaleway.VpcPublicGatewayArgs{
//				Type: pulumi.String("VPC-GW-S"),
//				IpId: ip01.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			gn01, err := scaleway.NewVpcGatewayNetwork(ctx, "gn01", &scaleway.VpcGatewayNetworkArgs{
//				GatewayId:        pg01.ID(),
//				PrivateNetworkId: pn01.ID(),
//				DhcpId:           dhcp01.ID(),
//				CleanupDhcp:      pulumi.Bool(true),
//				EnableMasquerade: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			rsv01, err := scaleway.NewVpcPublicGatewayDhcpReservation(ctx, "rsv01", &scaleway.VpcPublicGatewayDhcpReservationArgs{
//				GatewayNetworkId: gn01.ID(),
//				MacAddress:       pnic01.MacAddress,
//				IpAddress:        pulumi.String("192.168.0.7"),
//			})
//			if err != nil {
//				return err
//			}
//			// PAT rule for SSH traffic
//			_, err = scaleway.NewVpcPublicGatewayPatRule(ctx, "pat01", &scaleway.VpcPublicGatewayPatRuleArgs{
//				GatewayId:   pg01.ID(),
//				PrivateIp:   rsv01.IpAddress,
//				PrivatePort: pulumi.Int(22),
//				PublicPort:  pulumi.Int(2202),
//				Protocol:    pulumi.String("tcp"),
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
// Public Gateway PAT rule configurations can be imported using `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule main fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcPublicGatewayPatRule struct {
	pulumi.CustomResourceState

	// The date and time of the creation of the PAT rule configuration.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ID of the Public Gateway.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// The Organization ID the PAT rule configuration is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The private IP address to forward data to.
	PrivateIp pulumi.StringOutput `pulumi:"privateIp"`
	// The private port to translate to.
	PrivatePort pulumi.IntOutput `pulumi:"privatePort"`
	// The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// The public port to listen on.
	PublicPort pulumi.IntOutput `pulumi:"publicPort"`
	// The date and time of the last update of the PAT rule configuration.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// `zone`) The zone in which the Public Gateway DHCP configuration should be created.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayPatRule registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayPatRule(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayPatRuleArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayPatRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.PrivateIp == nil {
		return nil, errors.New("invalid value for required argument 'PrivateIp'")
	}
	if args.PrivatePort == nil {
		return nil, errors.New("invalid value for required argument 'PrivatePort'")
	}
	if args.PublicPort == nil {
		return nil, errors.New("invalid value for required argument 'PublicPort'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayPatRule
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayPatRule gets an existing VpcPublicGatewayPatRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayPatRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayPatRuleState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayPatRule, error) {
	var resource VpcPublicGatewayPatRule
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayPatRule resources.
type vpcPublicGatewayPatRuleState struct {
	// The date and time of the creation of the PAT rule configuration.
	CreatedAt *string `pulumi:"createdAt"`
	// The ID of the Public Gateway.
	GatewayId *string `pulumi:"gatewayId"`
	// The Organization ID the PAT rule configuration is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The private IP address to forward data to.
	PrivateIp *string `pulumi:"privateIp"`
	// The private port to translate to.
	PrivatePort *int `pulumi:"privatePort"`
	// The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
	Protocol *string `pulumi:"protocol"`
	// The public port to listen on.
	PublicPort *int `pulumi:"publicPort"`
	// The date and time of the last update of the PAT rule configuration.
	UpdatedAt *string `pulumi:"updatedAt"`
	// `zone`) The zone in which the Public Gateway DHCP configuration should be created.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayPatRuleState struct {
	// The date and time of the creation of the PAT rule configuration.
	CreatedAt pulumi.StringPtrInput
	// The ID of the Public Gateway.
	GatewayId pulumi.StringPtrInput
	// The Organization ID the PAT rule configuration is associated with.
	OrganizationId pulumi.StringPtrInput
	// The private IP address to forward data to.
	PrivateIp pulumi.StringPtrInput
	// The private port to translate to.
	PrivatePort pulumi.IntPtrInput
	// The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
	Protocol pulumi.StringPtrInput
	// The public port to listen on.
	PublicPort pulumi.IntPtrInput
	// The date and time of the last update of the PAT rule configuration.
	UpdatedAt pulumi.StringPtrInput
	// `zone`) The zone in which the Public Gateway DHCP configuration should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayPatRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayPatRuleState)(nil)).Elem()
}

type vpcPublicGatewayPatRuleArgs struct {
	// The ID of the Public Gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The private IP address to forward data to.
	PrivateIp string `pulumi:"privateIp"`
	// The private port to translate to.
	PrivatePort int `pulumi:"privatePort"`
	// The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
	Protocol *string `pulumi:"protocol"`
	// The public port to listen on.
	PublicPort int `pulumi:"publicPort"`
	// `zone`) The zone in which the Public Gateway DHCP configuration should be created.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayPatRule resource.
type VpcPublicGatewayPatRuleArgs struct {
	// The ID of the Public Gateway.
	GatewayId pulumi.StringInput
	// The private IP address to forward data to.
	PrivateIp pulumi.StringInput
	// The private port to translate to.
	PrivatePort pulumi.IntInput
	// The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
	Protocol pulumi.StringPtrInput
	// The public port to listen on.
	PublicPort pulumi.IntInput
	// `zone`) The zone in which the Public Gateway DHCP configuration should be created.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayPatRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayPatRuleArgs)(nil)).Elem()
}

type VpcPublicGatewayPatRuleInput interface {
	pulumi.Input

	ToVpcPublicGatewayPatRuleOutput() VpcPublicGatewayPatRuleOutput
	ToVpcPublicGatewayPatRuleOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleOutput
}

func (*VpcPublicGatewayPatRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayPatRule)(nil)).Elem()
}

func (i *VpcPublicGatewayPatRule) ToVpcPublicGatewayPatRuleOutput() VpcPublicGatewayPatRuleOutput {
	return i.ToVpcPublicGatewayPatRuleOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayPatRule) ToVpcPublicGatewayPatRuleOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayPatRuleOutput)
}

// VpcPublicGatewayPatRuleArrayInput is an input type that accepts VpcPublicGatewayPatRuleArray and VpcPublicGatewayPatRuleArrayOutput values.
// You can construct a concrete instance of `VpcPublicGatewayPatRuleArrayInput` via:
//
//	VpcPublicGatewayPatRuleArray{ VpcPublicGatewayPatRuleArgs{...} }
type VpcPublicGatewayPatRuleArrayInput interface {
	pulumi.Input

	ToVpcPublicGatewayPatRuleArrayOutput() VpcPublicGatewayPatRuleArrayOutput
	ToVpcPublicGatewayPatRuleArrayOutputWithContext(context.Context) VpcPublicGatewayPatRuleArrayOutput
}

type VpcPublicGatewayPatRuleArray []VpcPublicGatewayPatRuleInput

func (VpcPublicGatewayPatRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayPatRule)(nil)).Elem()
}

func (i VpcPublicGatewayPatRuleArray) ToVpcPublicGatewayPatRuleArrayOutput() VpcPublicGatewayPatRuleArrayOutput {
	return i.ToVpcPublicGatewayPatRuleArrayOutputWithContext(context.Background())
}

func (i VpcPublicGatewayPatRuleArray) ToVpcPublicGatewayPatRuleArrayOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayPatRuleArrayOutput)
}

// VpcPublicGatewayPatRuleMapInput is an input type that accepts VpcPublicGatewayPatRuleMap and VpcPublicGatewayPatRuleMapOutput values.
// You can construct a concrete instance of `VpcPublicGatewayPatRuleMapInput` via:
//
//	VpcPublicGatewayPatRuleMap{ "key": VpcPublicGatewayPatRuleArgs{...} }
type VpcPublicGatewayPatRuleMapInput interface {
	pulumi.Input

	ToVpcPublicGatewayPatRuleMapOutput() VpcPublicGatewayPatRuleMapOutput
	ToVpcPublicGatewayPatRuleMapOutputWithContext(context.Context) VpcPublicGatewayPatRuleMapOutput
}

type VpcPublicGatewayPatRuleMap map[string]VpcPublicGatewayPatRuleInput

func (VpcPublicGatewayPatRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayPatRule)(nil)).Elem()
}

func (i VpcPublicGatewayPatRuleMap) ToVpcPublicGatewayPatRuleMapOutput() VpcPublicGatewayPatRuleMapOutput {
	return i.ToVpcPublicGatewayPatRuleMapOutputWithContext(context.Background())
}

func (i VpcPublicGatewayPatRuleMap) ToVpcPublicGatewayPatRuleMapOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayPatRuleMapOutput)
}

type VpcPublicGatewayPatRuleOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayPatRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayPatRule)(nil)).Elem()
}

func (o VpcPublicGatewayPatRuleOutput) ToVpcPublicGatewayPatRuleOutput() VpcPublicGatewayPatRuleOutput {
	return o
}

func (o VpcPublicGatewayPatRuleOutput) ToVpcPublicGatewayPatRuleOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleOutput {
	return o
}

// The date and time of the creation of the PAT rule configuration.
func (o VpcPublicGatewayPatRuleOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ID of the Public Gateway.
func (o VpcPublicGatewayPatRuleOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// The Organization ID the PAT rule configuration is associated with.
func (o VpcPublicGatewayPatRuleOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The private IP address to forward data to.
func (o VpcPublicGatewayPatRuleOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.PrivateIp }).(pulumi.StringOutput)
}

// The private port to translate to.
func (o VpcPublicGatewayPatRuleOutput) PrivatePort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.IntOutput { return v.PrivatePort }).(pulumi.IntOutput)
}

// The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
func (o VpcPublicGatewayPatRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The public port to listen on.
func (o VpcPublicGatewayPatRuleOutput) PublicPort() pulumi.IntOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.IntOutput { return v.PublicPort }).(pulumi.IntOutput)
}

// The date and time of the last update of the PAT rule configuration.
func (o VpcPublicGatewayPatRuleOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// `zone`) The zone in which the Public Gateway DHCP configuration should be created.
func (o VpcPublicGatewayPatRuleOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayPatRule) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPublicGatewayPatRuleArrayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayPatRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayPatRule)(nil)).Elem()
}

func (o VpcPublicGatewayPatRuleArrayOutput) ToVpcPublicGatewayPatRuleArrayOutput() VpcPublicGatewayPatRuleArrayOutput {
	return o
}

func (o VpcPublicGatewayPatRuleArrayOutput) ToVpcPublicGatewayPatRuleArrayOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleArrayOutput {
	return o
}

func (o VpcPublicGatewayPatRuleArrayOutput) Index(i pulumi.IntInput) VpcPublicGatewayPatRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPublicGatewayPatRule {
		return vs[0].([]*VpcPublicGatewayPatRule)[vs[1].(int)]
	}).(VpcPublicGatewayPatRuleOutput)
}

type VpcPublicGatewayPatRuleMapOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayPatRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayPatRule)(nil)).Elem()
}

func (o VpcPublicGatewayPatRuleMapOutput) ToVpcPublicGatewayPatRuleMapOutput() VpcPublicGatewayPatRuleMapOutput {
	return o
}

func (o VpcPublicGatewayPatRuleMapOutput) ToVpcPublicGatewayPatRuleMapOutputWithContext(ctx context.Context) VpcPublicGatewayPatRuleMapOutput {
	return o
}

func (o VpcPublicGatewayPatRuleMapOutput) MapIndex(k pulumi.StringInput) VpcPublicGatewayPatRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPublicGatewayPatRule {
		return vs[0].(map[string]*VpcPublicGatewayPatRule)[vs[1].(string)]
	}).(VpcPublicGatewayPatRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayPatRuleInput)(nil)).Elem(), &VpcPublicGatewayPatRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayPatRuleArrayInput)(nil)).Elem(), VpcPublicGatewayPatRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayPatRuleMapInput)(nil)).Elem(), VpcPublicGatewayPatRuleMap{})
	pulumi.RegisterOutputType(VpcPublicGatewayPatRuleOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayPatRuleArrayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayPatRuleMapOutput{})
}
