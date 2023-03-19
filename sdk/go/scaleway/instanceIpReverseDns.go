// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages Scaleway Compute Instance IPs Reverse DNS.
//
// Please check our [guide](https://www.scaleway.com/en/docs/compute/instances/how-to/configure-reverse-dns/) for more details
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
//			serverIp, err := scaleway.NewInstanceIp(ctx, "serverIp", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDomainRecord(ctx, "tfA", &scaleway.DomainRecordArgs{
//				DnsZone:  pulumi.String("scaleway.com"),
//				Type:     pulumi.String("A"),
//				Data:     serverIp.Address,
//				Ttl:      pulumi.Int(3600),
//				Priority: pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewInstanceIpReverseDns(ctx, "reverse", &scaleway.InstanceIpReverseDnsArgs{
//				IpId:    serverIp.ID(),
//				Reverse: pulumi.String("www.scaleway.com"),
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
// IPs reverse DNS can be imported using the `{zone}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/instanceIpReverseDns:InstanceIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
//
// ```
type InstanceIpReverseDns struct {
	pulumi.CustomResourceState

	// The IP ID
	IpId pulumi.StringOutput `pulumi:"ipId"`
	// The reverse DNS for this IP.
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceIpReverseDns registers a new resource with the given unique name, arguments, and options.
func NewInstanceIpReverseDns(ctx *pulumi.Context,
	name string, args *InstanceIpReverseDnsArgs, opts ...pulumi.ResourceOption) (*InstanceIpReverseDns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpId == nil {
		return nil, errors.New("invalid value for required argument 'IpId'")
	}
	if args.Reverse == nil {
		return nil, errors.New("invalid value for required argument 'Reverse'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceIpReverseDns
	err := ctx.RegisterResource("scaleway:index/instanceIpReverseDns:InstanceIpReverseDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIpReverseDns gets an existing InstanceIpReverseDns resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIpReverseDns(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIpReverseDnsState, opts ...pulumi.ResourceOption) (*InstanceIpReverseDns, error) {
	var resource InstanceIpReverseDns
	err := ctx.ReadResource("scaleway:index/instanceIpReverseDns:InstanceIpReverseDns", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIpReverseDns resources.
type instanceIpReverseDnsState struct {
	// The IP ID
	IpId *string `pulumi:"ipId"`
	// The reverse DNS for this IP.
	Reverse *string `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

type InstanceIpReverseDnsState struct {
	// The IP ID
	IpId pulumi.StringPtrInput
	// The reverse DNS for this IP.
	Reverse pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (InstanceIpReverseDnsState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIpReverseDnsState)(nil)).Elem()
}

type instanceIpReverseDnsArgs struct {
	// The IP ID
	IpId string `pulumi:"ipId"`
	// The reverse DNS for this IP.
	Reverse string `pulumi:"reverse"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceIpReverseDns resource.
type InstanceIpReverseDnsArgs struct {
	// The IP ID
	IpId pulumi.StringInput
	// The reverse DNS for this IP.
	Reverse pulumi.StringInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (InstanceIpReverseDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIpReverseDnsArgs)(nil)).Elem()
}

type InstanceIpReverseDnsInput interface {
	pulumi.Input

	ToInstanceIpReverseDnsOutput() InstanceIpReverseDnsOutput
	ToInstanceIpReverseDnsOutputWithContext(ctx context.Context) InstanceIpReverseDnsOutput
}

func (*InstanceIpReverseDns) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIpReverseDns)(nil)).Elem()
}

func (i *InstanceIpReverseDns) ToInstanceIpReverseDnsOutput() InstanceIpReverseDnsOutput {
	return i.ToInstanceIpReverseDnsOutputWithContext(context.Background())
}

func (i *InstanceIpReverseDns) ToInstanceIpReverseDnsOutputWithContext(ctx context.Context) InstanceIpReverseDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpReverseDnsOutput)
}

// InstanceIpReverseDnsArrayInput is an input type that accepts InstanceIpReverseDnsArray and InstanceIpReverseDnsArrayOutput values.
// You can construct a concrete instance of `InstanceIpReverseDnsArrayInput` via:
//
//	InstanceIpReverseDnsArray{ InstanceIpReverseDnsArgs{...} }
type InstanceIpReverseDnsArrayInput interface {
	pulumi.Input

	ToInstanceIpReverseDnsArrayOutput() InstanceIpReverseDnsArrayOutput
	ToInstanceIpReverseDnsArrayOutputWithContext(context.Context) InstanceIpReverseDnsArrayOutput
}

type InstanceIpReverseDnsArray []InstanceIpReverseDnsInput

func (InstanceIpReverseDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIpReverseDns)(nil)).Elem()
}

func (i InstanceIpReverseDnsArray) ToInstanceIpReverseDnsArrayOutput() InstanceIpReverseDnsArrayOutput {
	return i.ToInstanceIpReverseDnsArrayOutputWithContext(context.Background())
}

func (i InstanceIpReverseDnsArray) ToInstanceIpReverseDnsArrayOutputWithContext(ctx context.Context) InstanceIpReverseDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpReverseDnsArrayOutput)
}

// InstanceIpReverseDnsMapInput is an input type that accepts InstanceIpReverseDnsMap and InstanceIpReverseDnsMapOutput values.
// You can construct a concrete instance of `InstanceIpReverseDnsMapInput` via:
//
//	InstanceIpReverseDnsMap{ "key": InstanceIpReverseDnsArgs{...} }
type InstanceIpReverseDnsMapInput interface {
	pulumi.Input

	ToInstanceIpReverseDnsMapOutput() InstanceIpReverseDnsMapOutput
	ToInstanceIpReverseDnsMapOutputWithContext(context.Context) InstanceIpReverseDnsMapOutput
}

type InstanceIpReverseDnsMap map[string]InstanceIpReverseDnsInput

func (InstanceIpReverseDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIpReverseDns)(nil)).Elem()
}

func (i InstanceIpReverseDnsMap) ToInstanceIpReverseDnsMapOutput() InstanceIpReverseDnsMapOutput {
	return i.ToInstanceIpReverseDnsMapOutputWithContext(context.Background())
}

func (i InstanceIpReverseDnsMap) ToInstanceIpReverseDnsMapOutputWithContext(ctx context.Context) InstanceIpReverseDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpReverseDnsMapOutput)
}

type InstanceIpReverseDnsOutput struct{ *pulumi.OutputState }

func (InstanceIpReverseDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIpReverseDns)(nil)).Elem()
}

func (o InstanceIpReverseDnsOutput) ToInstanceIpReverseDnsOutput() InstanceIpReverseDnsOutput {
	return o
}

func (o InstanceIpReverseDnsOutput) ToInstanceIpReverseDnsOutputWithContext(ctx context.Context) InstanceIpReverseDnsOutput {
	return o
}

// The IP ID
func (o InstanceIpReverseDnsOutput) IpId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIpReverseDns) pulumi.StringOutput { return v.IpId }).(pulumi.StringOutput)
}

// The reverse DNS for this IP.
func (o InstanceIpReverseDnsOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIpReverseDns) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// `zone`) The zone in which the IP should be reserved.
func (o InstanceIpReverseDnsOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIpReverseDns) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceIpReverseDnsArrayOutput struct{ *pulumi.OutputState }

func (InstanceIpReverseDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIpReverseDns)(nil)).Elem()
}

func (o InstanceIpReverseDnsArrayOutput) ToInstanceIpReverseDnsArrayOutput() InstanceIpReverseDnsArrayOutput {
	return o
}

func (o InstanceIpReverseDnsArrayOutput) ToInstanceIpReverseDnsArrayOutputWithContext(ctx context.Context) InstanceIpReverseDnsArrayOutput {
	return o
}

func (o InstanceIpReverseDnsArrayOutput) Index(i pulumi.IntInput) InstanceIpReverseDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceIpReverseDns {
		return vs[0].([]*InstanceIpReverseDns)[vs[1].(int)]
	}).(InstanceIpReverseDnsOutput)
}

type InstanceIpReverseDnsMapOutput struct{ *pulumi.OutputState }

func (InstanceIpReverseDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIpReverseDns)(nil)).Elem()
}

func (o InstanceIpReverseDnsMapOutput) ToInstanceIpReverseDnsMapOutput() InstanceIpReverseDnsMapOutput {
	return o
}

func (o InstanceIpReverseDnsMapOutput) ToInstanceIpReverseDnsMapOutputWithContext(ctx context.Context) InstanceIpReverseDnsMapOutput {
	return o
}

func (o InstanceIpReverseDnsMapOutput) MapIndex(k pulumi.StringInput) InstanceIpReverseDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceIpReverseDns {
		return vs[0].(map[string]*InstanceIpReverseDns)[vs[1].(string)]
	}).(InstanceIpReverseDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpReverseDnsInput)(nil)).Elem(), &InstanceIpReverseDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpReverseDnsArrayInput)(nil)).Elem(), InstanceIpReverseDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpReverseDnsMapInput)(nil)).Elem(), InstanceIpReverseDnsMap{})
	pulumi.RegisterOutputType(InstanceIpReverseDnsOutput{})
	pulumi.RegisterOutputType(InstanceIpReverseDnsArrayOutput{})
	pulumi.RegisterOutputType(InstanceIpReverseDnsMapOutput{})
}
