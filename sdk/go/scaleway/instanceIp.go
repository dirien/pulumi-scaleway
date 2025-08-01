// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway compute Instance IPs. For more information, see the [API documentation](https://www.scaleway.com/en/developers/api/instance/#path-ips-list-all-flexible-ips).
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
//			_, err := scaleway.NewInstanceIp(ctx, "serverIp", nil)
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
// IPs can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/instanceIp:InstanceIp server_ip fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type InstanceIp struct {
	pulumi.CustomResourceState

	// The IP address.
	Address pulumi.StringOutput `pulumi:"address"`
	// The organization ID the IP is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The IP Prefix.
	Prefix pulumi.StringOutput `pulumi:"prefix"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The reverse dns attached to this IP
	Reverse pulumi.StringOutput `pulumi:"reverse"`
	// The server associated with this IP
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The tags associated with the IP.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The type of the IP (`routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
	Type pulumi.StringOutput `pulumi:"type"`
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceIp registers a new resource with the given unique name, arguments, and options.
func NewInstanceIp(ctx *pulumi.Context,
	name string, args *InstanceIpArgs, opts ...pulumi.ResourceOption) (*InstanceIp, error) {
	if args == nil {
		args = &InstanceIpArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceIp
	err := ctx.RegisterResource("scaleway:index/instanceIp:InstanceIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIp gets an existing InstanceIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIpState, opts ...pulumi.ResourceOption) (*InstanceIp, error) {
	var resource InstanceIp
	err := ctx.ReadResource("scaleway:index/instanceIp:InstanceIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIp resources.
type instanceIpState struct {
	// The IP address.
	Address *string `pulumi:"address"`
	// The organization ID the IP is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The IP Prefix.
	Prefix *string `pulumi:"prefix"`
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The reverse dns attached to this IP
	Reverse *string `pulumi:"reverse"`
	// The server associated with this IP
	ServerId *string `pulumi:"serverId"`
	// The tags associated with the IP.
	Tags []string `pulumi:"tags"`
	// The type of the IP (`routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
	Type *string `pulumi:"type"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

type InstanceIpState struct {
	// The IP address.
	Address pulumi.StringPtrInput
	// The organization ID the IP is associated with.
	OrganizationId pulumi.StringPtrInput
	// The IP Prefix.
	Prefix pulumi.StringPtrInput
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringPtrInput
	// The reverse dns attached to this IP
	Reverse pulumi.StringPtrInput
	// The server associated with this IP
	ServerId pulumi.StringPtrInput
	// The tags associated with the IP.
	Tags pulumi.StringArrayInput
	// The type of the IP (`routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
	Type pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (InstanceIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIpState)(nil)).Elem()
}

type instanceIpArgs struct {
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The tags associated with the IP.
	Tags []string `pulumi:"tags"`
	// The type of the IP (`routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
	Type *string `pulumi:"type"`
	// `zone`) The zone in which the IP should be reserved.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceIp resource.
type InstanceIpArgs struct {
	// `projectId`) The ID of the project the IP is associated with.
	ProjectId pulumi.StringPtrInput
	// The tags associated with the IP.
	Tags pulumi.StringArrayInput
	// The type of the IP (`routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
	Type pulumi.StringPtrInput
	// `zone`) The zone in which the IP should be reserved.
	Zone pulumi.StringPtrInput
}

func (InstanceIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIpArgs)(nil)).Elem()
}

type InstanceIpInput interface {
	pulumi.Input

	ToInstanceIpOutput() InstanceIpOutput
	ToInstanceIpOutputWithContext(ctx context.Context) InstanceIpOutput
}

func (*InstanceIp) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIp)(nil)).Elem()
}

func (i *InstanceIp) ToInstanceIpOutput() InstanceIpOutput {
	return i.ToInstanceIpOutputWithContext(context.Background())
}

func (i *InstanceIp) ToInstanceIpOutputWithContext(ctx context.Context) InstanceIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpOutput)
}

// InstanceIpArrayInput is an input type that accepts InstanceIpArray and InstanceIpArrayOutput values.
// You can construct a concrete instance of `InstanceIpArrayInput` via:
//
//	InstanceIpArray{ InstanceIpArgs{...} }
type InstanceIpArrayInput interface {
	pulumi.Input

	ToInstanceIpArrayOutput() InstanceIpArrayOutput
	ToInstanceIpArrayOutputWithContext(context.Context) InstanceIpArrayOutput
}

type InstanceIpArray []InstanceIpInput

func (InstanceIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIp)(nil)).Elem()
}

func (i InstanceIpArray) ToInstanceIpArrayOutput() InstanceIpArrayOutput {
	return i.ToInstanceIpArrayOutputWithContext(context.Background())
}

func (i InstanceIpArray) ToInstanceIpArrayOutputWithContext(ctx context.Context) InstanceIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpArrayOutput)
}

// InstanceIpMapInput is an input type that accepts InstanceIpMap and InstanceIpMapOutput values.
// You can construct a concrete instance of `InstanceIpMapInput` via:
//
//	InstanceIpMap{ "key": InstanceIpArgs{...} }
type InstanceIpMapInput interface {
	pulumi.Input

	ToInstanceIpMapOutput() InstanceIpMapOutput
	ToInstanceIpMapOutputWithContext(context.Context) InstanceIpMapOutput
}

type InstanceIpMap map[string]InstanceIpInput

func (InstanceIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIp)(nil)).Elem()
}

func (i InstanceIpMap) ToInstanceIpMapOutput() InstanceIpMapOutput {
	return i.ToInstanceIpMapOutputWithContext(context.Background())
}

func (i InstanceIpMap) ToInstanceIpMapOutputWithContext(ctx context.Context) InstanceIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIpMapOutput)
}

type InstanceIpOutput struct{ *pulumi.OutputState }

func (InstanceIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceIp)(nil)).Elem()
}

func (o InstanceIpOutput) ToInstanceIpOutput() InstanceIpOutput {
	return o
}

func (o InstanceIpOutput) ToInstanceIpOutputWithContext(ctx context.Context) InstanceIpOutput {
	return o
}

// The IP address.
func (o InstanceIpOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The organization ID the IP is associated with.
func (o InstanceIpOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The IP Prefix.
func (o InstanceIpOutput) Prefix() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.Prefix }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the IP is associated with.
func (o InstanceIpOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The reverse dns attached to this IP
func (o InstanceIpOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

// The server associated with this IP
func (o InstanceIpOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The tags associated with the IP.
func (o InstanceIpOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of the IP (`routedIpv4`, `routedIpv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
func (o InstanceIpOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// `zone`) The zone in which the IP should be reserved.
func (o InstanceIpOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceIp) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceIpArrayOutput struct{ *pulumi.OutputState }

func (InstanceIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceIp)(nil)).Elem()
}

func (o InstanceIpArrayOutput) ToInstanceIpArrayOutput() InstanceIpArrayOutput {
	return o
}

func (o InstanceIpArrayOutput) ToInstanceIpArrayOutputWithContext(ctx context.Context) InstanceIpArrayOutput {
	return o
}

func (o InstanceIpArrayOutput) Index(i pulumi.IntInput) InstanceIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceIp {
		return vs[0].([]*InstanceIp)[vs[1].(int)]
	}).(InstanceIpOutput)
}

type InstanceIpMapOutput struct{ *pulumi.OutputState }

func (InstanceIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceIp)(nil)).Elem()
}

func (o InstanceIpMapOutput) ToInstanceIpMapOutput() InstanceIpMapOutput {
	return o
}

func (o InstanceIpMapOutput) ToInstanceIpMapOutputWithContext(ctx context.Context) InstanceIpMapOutput {
	return o
}

func (o InstanceIpMapOutput) MapIndex(k pulumi.StringInput) InstanceIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceIp {
		return vs[0].(map[string]*InstanceIp)[vs[1].(string)]
	}).(InstanceIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpInput)(nil)).Elem(), &InstanceIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpArrayInput)(nil)).Elem(), InstanceIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceIpMapInput)(nil)).Elem(), InstanceIpMap{})
	pulumi.RegisterOutputType(InstanceIpOutput{})
	pulumi.RegisterOutputType(InstanceIpArrayOutput{})
	pulumi.RegisterOutputType(InstanceIpMapOutput{})
}
