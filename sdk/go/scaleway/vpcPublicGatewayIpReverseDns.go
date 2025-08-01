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

// Manages Scaleway Public Gateway public (flexible) IPs' reverse DNS.
// For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).
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
//			mainVpcPublicGatewayIp, err := scaleway.NewVpcPublicGatewayIp(ctx, "mainVpcPublicGatewayIp", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "tfA", &scaleway.DomainRecordArgs{
//				DnsZone:  pulumi.String("example.com"),
//				Type:     pulumi.String("A"),
//				Data:     mainVpcPublicGatewayIp.Address,
//				Ttl:      pulumi.Int(3600),
//				Priority: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewVpcPublicGatewayIpReverseDns(ctx, "mainVpcPublicGatewayIpReverseDns", &scaleway.VpcPublicGatewayIpReverseDnsArgs{
//				GatewayIpId: mainVpcPublicGatewayIp.ID(),
//				Reverse:     pulumi.String("tf.example.com"),
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
// Public Gateway IP reverse DNS can be imported using `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type VpcPublicGatewayIpReverseDns struct {
	pulumi.CustomResourceState

	// The Public Gateway IP ID
	GatewayIpId pulumi.StringOutput `pulumi:"gatewayIpId"`
	// The reverse domain name for this IP address
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVpcPublicGatewayIpReverseDns registers a new resource with the given unique name, arguments, and options.
func NewVpcPublicGatewayIpReverseDns(ctx *pulumi.Context,
	name string, args *VpcPublicGatewayIpReverseDnsArgs, opts ...pulumi.ResourceOption) (*VpcPublicGatewayIpReverseDns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayIpId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayIpId'")
	}
	if args.Reverse == nil {
		return nil, errors.New("invalid value for required argument 'Reverse'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcPublicGatewayIpReverseDns
	err := ctx.RegisterResource("scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcPublicGatewayIpReverseDns gets an existing VpcPublicGatewayIpReverseDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPublicGatewayIpReverseDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcPublicGatewayIpReverseDnsState, opts ...pulumi.ResourceOption) (*VpcPublicGatewayIpReverseDns, error) {
	var resource VpcPublicGatewayIpReverseDns
	err := ctx.ReadResource("scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcPublicGatewayIpReverseDns resources.
type vpcPublicGatewayIpReverseDnsState struct {
	// The Public Gateway IP ID
	GatewayIpId *string `pulumi:"gatewayIpId"`
	// The reverse domain name for this IP address
	Reverse *string `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

type VpcPublicGatewayIpReverseDnsState struct {
	// The Public Gateway IP ID
	GatewayIpId pulumi.StringPtrInput
	// The reverse domain name for this IP address
	Reverse pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayIpReverseDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayIpReverseDnsState)(nil)).Elem()
}

type vpcPublicGatewayIpReverseDnsArgs struct {
	// The Public Gateway IP ID
	GatewayIpId string `pulumi:"gatewayIpId"`
	// The reverse domain name for this IP address
	Reverse string `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a VpcPublicGatewayIpReverseDns resource.
type VpcPublicGatewayIpReverseDnsArgs struct {
	// The Public Gateway IP ID
	GatewayIpId pulumi.StringInput
	// The reverse domain name for this IP address
	Reverse pulumi.StringInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (VpcPublicGatewayIpReverseDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcPublicGatewayIpReverseDnsArgs)(nil)).Elem()
}

type VpcPublicGatewayIpReverseDnsInput interface {
	pulumi.Input

	ToVpcPublicGatewayIpReverseDnsOutput() VpcPublicGatewayIpReverseDnsOutput
	ToVpcPublicGatewayIpReverseDnsOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsOutput
}

func (*VpcPublicGatewayIpReverseDns) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayIpReverseDns)(nil)).Elem()
}

func (i *VpcPublicGatewayIpReverseDns) ToVpcPublicGatewayIpReverseDnsOutput() VpcPublicGatewayIpReverseDnsOutput {
	return i.ToVpcPublicGatewayIpReverseDnsOutputWithContext(context.Background())
}

func (i *VpcPublicGatewayIpReverseDns) ToVpcPublicGatewayIpReverseDnsOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayIpReverseDnsOutput)
}

// VpcPublicGatewayIpReverseDnsArrayInput is an input type that accepts VpcPublicGatewayIpReverseDnsArray and VpcPublicGatewayIpReverseDnsArrayOutput values.
// You can construct a concrete instance of `VpcPublicGatewayIpReverseDnsArrayInput` via:
//
//	VpcPublicGatewayIpReverseDnsArray{ VpcPublicGatewayIpReverseDnsArgs{...} }
type VpcPublicGatewayIpReverseDnsArrayInput interface {
	pulumi.Input

	ToVpcPublicGatewayIpReverseDnsArrayOutput() VpcPublicGatewayIpReverseDnsArrayOutput
	ToVpcPublicGatewayIpReverseDnsArrayOutputWithContext(context.Context) VpcPublicGatewayIpReverseDnsArrayOutput
}

type VpcPublicGatewayIpReverseDnsArray []VpcPublicGatewayIpReverseDnsInput

func (VpcPublicGatewayIpReverseDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayIpReverseDns)(nil)).Elem()
}

func (i VpcPublicGatewayIpReverseDnsArray) ToVpcPublicGatewayIpReverseDnsArrayOutput() VpcPublicGatewayIpReverseDnsArrayOutput {
	return i.ToVpcPublicGatewayIpReverseDnsArrayOutputWithContext(context.Background())
}

func (i VpcPublicGatewayIpReverseDnsArray) ToVpcPublicGatewayIpReverseDnsArrayOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayIpReverseDnsArrayOutput)
}

// VpcPublicGatewayIpReverseDnsMapInput is an input type that accepts VpcPublicGatewayIpReverseDnsMap and VpcPublicGatewayIpReverseDnsMapOutput values.
// You can construct a concrete instance of `VpcPublicGatewayIpReverseDnsMapInput` via:
//
//	VpcPublicGatewayIpReverseDnsMap{ "key": VpcPublicGatewayIpReverseDnsArgs{...} }
type VpcPublicGatewayIpReverseDnsMapInput interface {
	pulumi.Input

	ToVpcPublicGatewayIpReverseDnsMapOutput() VpcPublicGatewayIpReverseDnsMapOutput
	ToVpcPublicGatewayIpReverseDnsMapOutputWithContext(context.Context) VpcPublicGatewayIpReverseDnsMapOutput
}

type VpcPublicGatewayIpReverseDnsMap map[string]VpcPublicGatewayIpReverseDnsInput

func (VpcPublicGatewayIpReverseDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayIpReverseDns)(nil)).Elem()
}

func (i VpcPublicGatewayIpReverseDnsMap) ToVpcPublicGatewayIpReverseDnsMapOutput() VpcPublicGatewayIpReverseDnsMapOutput {
	return i.ToVpcPublicGatewayIpReverseDnsMapOutputWithContext(context.Background())
}

func (i VpcPublicGatewayIpReverseDnsMap) ToVpcPublicGatewayIpReverseDnsMapOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPublicGatewayIpReverseDnsMapOutput)
}

type VpcPublicGatewayIpReverseDnsOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayIpReverseDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPublicGatewayIpReverseDns)(nil)).Elem()
}

func (o VpcPublicGatewayIpReverseDnsOutput) ToVpcPublicGatewayIpReverseDnsOutput() VpcPublicGatewayIpReverseDnsOutput {
	return o
}

func (o VpcPublicGatewayIpReverseDnsOutput) ToVpcPublicGatewayIpReverseDnsOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsOutput {
	return o
}

// The Public Gateway IP ID
func (o VpcPublicGatewayIpReverseDnsOutput) GatewayIpId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIpReverseDns) pulumi.StringOutput { return v.GatewayIpId }).(pulumi.StringOutput)
}

// The reverse domain name for this IP address
func (o VpcPublicGatewayIpReverseDnsOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIpReverseDns) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// `zone`) The zone in which the IP should be reserved.
func (o VpcPublicGatewayIpReverseDnsOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcPublicGatewayIpReverseDns) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VpcPublicGatewayIpReverseDnsArrayOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayIpReverseDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcPublicGatewayIpReverseDns)(nil)).Elem()
}

func (o VpcPublicGatewayIpReverseDnsArrayOutput) ToVpcPublicGatewayIpReverseDnsArrayOutput() VpcPublicGatewayIpReverseDnsArrayOutput {
	return o
}

func (o VpcPublicGatewayIpReverseDnsArrayOutput) ToVpcPublicGatewayIpReverseDnsArrayOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsArrayOutput {
	return o
}

func (o VpcPublicGatewayIpReverseDnsArrayOutput) Index(i pulumi.IntInput) VpcPublicGatewayIpReverseDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcPublicGatewayIpReverseDns {
		return vs[0].([]*VpcPublicGatewayIpReverseDns)[vs[1].(int)]
	}).(VpcPublicGatewayIpReverseDnsOutput)
}

type VpcPublicGatewayIpReverseDnsMapOutput struct{ *pulumi.OutputState }

func (VpcPublicGatewayIpReverseDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcPublicGatewayIpReverseDns)(nil)).Elem()
}

func (o VpcPublicGatewayIpReverseDnsMapOutput) ToVpcPublicGatewayIpReverseDnsMapOutput() VpcPublicGatewayIpReverseDnsMapOutput {
	return o
}

func (o VpcPublicGatewayIpReverseDnsMapOutput) ToVpcPublicGatewayIpReverseDnsMapOutputWithContext(ctx context.Context) VpcPublicGatewayIpReverseDnsMapOutput {
	return o
}

func (o VpcPublicGatewayIpReverseDnsMapOutput) MapIndex(k pulumi.StringInput) VpcPublicGatewayIpReverseDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcPublicGatewayIpReverseDns {
		return vs[0].(map[string]*VpcPublicGatewayIpReverseDns)[vs[1].(string)]
	}).(VpcPublicGatewayIpReverseDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayIpReverseDnsInput)(nil)).Elem(), &VpcPublicGatewayIpReverseDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayIpReverseDnsArrayInput)(nil)).Elem(), VpcPublicGatewayIpReverseDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcPublicGatewayIpReverseDnsMapInput)(nil)).Elem(), VpcPublicGatewayIpReverseDnsMap{})
	pulumi.RegisterOutputType(VpcPublicGatewayIpReverseDnsOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayIpReverseDnsArrayOutput{})
	pulumi.RegisterOutputType(VpcPublicGatewayIpReverseDnsMapOutput{})
}
