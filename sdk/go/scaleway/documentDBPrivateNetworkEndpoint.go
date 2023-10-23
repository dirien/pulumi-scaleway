// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type DocumentDBPrivateNetworkEndpoint struct {
	pulumi.CustomResourceState

	// The hostname of your endpoint
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Instance on which the endpoint is attached
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The IP of your private network service
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The IP with the given mask within the private subnet
	IpNet pulumi.StringOutput `pulumi:"ipNet"`
	// The name of your private service
	Name pulumi.StringOutput `pulumi:"name"`
	// The port of your private service
	Port pulumi.IntOutput `pulumi:"port"`
	// The private network ID
	PrivateNetworkId pulumi.StringOutput `pulumi:"privateNetworkId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// The zone you want to attach the resource to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDocumentDBPrivateNetworkEndpoint registers a new resource with the given unique name, arguments, and options.
func NewDocumentDBPrivateNetworkEndpoint(ctx *pulumi.Context,
	name string, args *DocumentDBPrivateNetworkEndpointArgs, opts ...pulumi.ResourceOption) (*DocumentDBPrivateNetworkEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PrivateNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateNetworkId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentDBPrivateNetworkEndpoint
	err := ctx.RegisterResource("scaleway:index/documentDBPrivateNetworkEndpoint:DocumentDBPrivateNetworkEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentDBPrivateNetworkEndpoint gets an existing DocumentDBPrivateNetworkEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentDBPrivateNetworkEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentDBPrivateNetworkEndpointState, opts ...pulumi.ResourceOption) (*DocumentDBPrivateNetworkEndpoint, error) {
	var resource DocumentDBPrivateNetworkEndpoint
	err := ctx.ReadResource("scaleway:index/documentDBPrivateNetworkEndpoint:DocumentDBPrivateNetworkEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentDBPrivateNetworkEndpoint resources.
type documentDBPrivateNetworkEndpointState struct {
	// The hostname of your endpoint
	Hostname *string `pulumi:"hostname"`
	// Instance on which the endpoint is attached
	InstanceId *string `pulumi:"instanceId"`
	// The IP of your private network service
	Ip *string `pulumi:"ip"`
	// The IP with the given mask within the private subnet
	IpNet *string `pulumi:"ipNet"`
	// The name of your private service
	Name *string `pulumi:"name"`
	// The port of your private service
	Port *int `pulumi:"port"`
	// The private network ID
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

type DocumentDBPrivateNetworkEndpointState struct {
	// The hostname of your endpoint
	Hostname pulumi.StringPtrInput
	// Instance on which the endpoint is attached
	InstanceId pulumi.StringPtrInput
	// The IP of your private network service
	Ip pulumi.StringPtrInput
	// The IP with the given mask within the private subnet
	IpNet pulumi.StringPtrInput
	// The name of your private service
	Name pulumi.StringPtrInput
	// The port of your private service
	Port pulumi.IntPtrInput
	// The private network ID
	PrivateNetworkId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (DocumentDBPrivateNetworkEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentDBPrivateNetworkEndpointState)(nil)).Elem()
}

type documentDBPrivateNetworkEndpointArgs struct {
	// Instance on which the endpoint is attached
	InstanceId string `pulumi:"instanceId"`
	// The IP with the given mask within the private subnet
	IpNet *string `pulumi:"ipNet"`
	// The port of your private service
	Port *int `pulumi:"port"`
	// The private network ID
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The zone you want to attach the resource to
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a DocumentDBPrivateNetworkEndpoint resource.
type DocumentDBPrivateNetworkEndpointArgs struct {
	// Instance on which the endpoint is attached
	InstanceId pulumi.StringInput
	// The IP with the given mask within the private subnet
	IpNet pulumi.StringPtrInput
	// The port of your private service
	Port pulumi.IntPtrInput
	// The private network ID
	PrivateNetworkId pulumi.StringInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The zone you want to attach the resource to
	Zone pulumi.StringPtrInput
}

func (DocumentDBPrivateNetworkEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentDBPrivateNetworkEndpointArgs)(nil)).Elem()
}

type DocumentDBPrivateNetworkEndpointInput interface {
	pulumi.Input

	ToDocumentDBPrivateNetworkEndpointOutput() DocumentDBPrivateNetworkEndpointOutput
	ToDocumentDBPrivateNetworkEndpointOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointOutput
}

func (*DocumentDBPrivateNetworkEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentDBPrivateNetworkEndpoint)(nil)).Elem()
}

func (i *DocumentDBPrivateNetworkEndpoint) ToDocumentDBPrivateNetworkEndpointOutput() DocumentDBPrivateNetworkEndpointOutput {
	return i.ToDocumentDBPrivateNetworkEndpointOutputWithContext(context.Background())
}

func (i *DocumentDBPrivateNetworkEndpoint) ToDocumentDBPrivateNetworkEndpointOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBPrivateNetworkEndpointOutput)
}

func (i *DocumentDBPrivateNetworkEndpoint) ToOutput(ctx context.Context) pulumix.Output[*DocumentDBPrivateNetworkEndpoint] {
	return pulumix.Output[*DocumentDBPrivateNetworkEndpoint]{
		OutputState: i.ToDocumentDBPrivateNetworkEndpointOutputWithContext(ctx).OutputState,
	}
}

// DocumentDBPrivateNetworkEndpointArrayInput is an input type that accepts DocumentDBPrivateNetworkEndpointArray and DocumentDBPrivateNetworkEndpointArrayOutput values.
// You can construct a concrete instance of `DocumentDBPrivateNetworkEndpointArrayInput` via:
//
//	DocumentDBPrivateNetworkEndpointArray{ DocumentDBPrivateNetworkEndpointArgs{...} }
type DocumentDBPrivateNetworkEndpointArrayInput interface {
	pulumi.Input

	ToDocumentDBPrivateNetworkEndpointArrayOutput() DocumentDBPrivateNetworkEndpointArrayOutput
	ToDocumentDBPrivateNetworkEndpointArrayOutputWithContext(context.Context) DocumentDBPrivateNetworkEndpointArrayOutput
}

type DocumentDBPrivateNetworkEndpointArray []DocumentDBPrivateNetworkEndpointInput

func (DocumentDBPrivateNetworkEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentDBPrivateNetworkEndpoint)(nil)).Elem()
}

func (i DocumentDBPrivateNetworkEndpointArray) ToDocumentDBPrivateNetworkEndpointArrayOutput() DocumentDBPrivateNetworkEndpointArrayOutput {
	return i.ToDocumentDBPrivateNetworkEndpointArrayOutputWithContext(context.Background())
}

func (i DocumentDBPrivateNetworkEndpointArray) ToDocumentDBPrivateNetworkEndpointArrayOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBPrivateNetworkEndpointArrayOutput)
}

func (i DocumentDBPrivateNetworkEndpointArray) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentDBPrivateNetworkEndpoint] {
	return pulumix.Output[[]*DocumentDBPrivateNetworkEndpoint]{
		OutputState: i.ToDocumentDBPrivateNetworkEndpointArrayOutputWithContext(ctx).OutputState,
	}
}

// DocumentDBPrivateNetworkEndpointMapInput is an input type that accepts DocumentDBPrivateNetworkEndpointMap and DocumentDBPrivateNetworkEndpointMapOutput values.
// You can construct a concrete instance of `DocumentDBPrivateNetworkEndpointMapInput` via:
//
//	DocumentDBPrivateNetworkEndpointMap{ "key": DocumentDBPrivateNetworkEndpointArgs{...} }
type DocumentDBPrivateNetworkEndpointMapInput interface {
	pulumi.Input

	ToDocumentDBPrivateNetworkEndpointMapOutput() DocumentDBPrivateNetworkEndpointMapOutput
	ToDocumentDBPrivateNetworkEndpointMapOutputWithContext(context.Context) DocumentDBPrivateNetworkEndpointMapOutput
}

type DocumentDBPrivateNetworkEndpointMap map[string]DocumentDBPrivateNetworkEndpointInput

func (DocumentDBPrivateNetworkEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentDBPrivateNetworkEndpoint)(nil)).Elem()
}

func (i DocumentDBPrivateNetworkEndpointMap) ToDocumentDBPrivateNetworkEndpointMapOutput() DocumentDBPrivateNetworkEndpointMapOutput {
	return i.ToDocumentDBPrivateNetworkEndpointMapOutputWithContext(context.Background())
}

func (i DocumentDBPrivateNetworkEndpointMap) ToDocumentDBPrivateNetworkEndpointMapOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBPrivateNetworkEndpointMapOutput)
}

func (i DocumentDBPrivateNetworkEndpointMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentDBPrivateNetworkEndpoint] {
	return pulumix.Output[map[string]*DocumentDBPrivateNetworkEndpoint]{
		OutputState: i.ToDocumentDBPrivateNetworkEndpointMapOutputWithContext(ctx).OutputState,
	}
}

type DocumentDBPrivateNetworkEndpointOutput struct{ *pulumi.OutputState }

func (DocumentDBPrivateNetworkEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentDBPrivateNetworkEndpoint)(nil)).Elem()
}

func (o DocumentDBPrivateNetworkEndpointOutput) ToDocumentDBPrivateNetworkEndpointOutput() DocumentDBPrivateNetworkEndpointOutput {
	return o
}

func (o DocumentDBPrivateNetworkEndpointOutput) ToDocumentDBPrivateNetworkEndpointOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointOutput {
	return o
}

func (o DocumentDBPrivateNetworkEndpointOutput) ToOutput(ctx context.Context) pulumix.Output[*DocumentDBPrivateNetworkEndpoint] {
	return pulumix.Output[*DocumentDBPrivateNetworkEndpoint]{
		OutputState: o.OutputState,
	}
}

// The hostname of your endpoint
func (o DocumentDBPrivateNetworkEndpointOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Instance on which the endpoint is attached
func (o DocumentDBPrivateNetworkEndpointOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The IP of your private network service
func (o DocumentDBPrivateNetworkEndpointOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The IP with the given mask within the private subnet
func (o DocumentDBPrivateNetworkEndpointOutput) IpNet() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.IpNet }).(pulumi.StringOutput)
}

// The name of your private service
func (o DocumentDBPrivateNetworkEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port of your private service
func (o DocumentDBPrivateNetworkEndpointOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The private network ID
func (o DocumentDBPrivateNetworkEndpointOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o DocumentDBPrivateNetworkEndpointOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The zone you want to attach the resource to
func (o DocumentDBPrivateNetworkEndpointOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivateNetworkEndpoint) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DocumentDBPrivateNetworkEndpointArrayOutput struct{ *pulumi.OutputState }

func (DocumentDBPrivateNetworkEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentDBPrivateNetworkEndpoint)(nil)).Elem()
}

func (o DocumentDBPrivateNetworkEndpointArrayOutput) ToDocumentDBPrivateNetworkEndpointArrayOutput() DocumentDBPrivateNetworkEndpointArrayOutput {
	return o
}

func (o DocumentDBPrivateNetworkEndpointArrayOutput) ToDocumentDBPrivateNetworkEndpointArrayOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointArrayOutput {
	return o
}

func (o DocumentDBPrivateNetworkEndpointArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DocumentDBPrivateNetworkEndpoint] {
	return pulumix.Output[[]*DocumentDBPrivateNetworkEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o DocumentDBPrivateNetworkEndpointArrayOutput) Index(i pulumi.IntInput) DocumentDBPrivateNetworkEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentDBPrivateNetworkEndpoint {
		return vs[0].([]*DocumentDBPrivateNetworkEndpoint)[vs[1].(int)]
	}).(DocumentDBPrivateNetworkEndpointOutput)
}

type DocumentDBPrivateNetworkEndpointMapOutput struct{ *pulumi.OutputState }

func (DocumentDBPrivateNetworkEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentDBPrivateNetworkEndpoint)(nil)).Elem()
}

func (o DocumentDBPrivateNetworkEndpointMapOutput) ToDocumentDBPrivateNetworkEndpointMapOutput() DocumentDBPrivateNetworkEndpointMapOutput {
	return o
}

func (o DocumentDBPrivateNetworkEndpointMapOutput) ToDocumentDBPrivateNetworkEndpointMapOutputWithContext(ctx context.Context) DocumentDBPrivateNetworkEndpointMapOutput {
	return o
}

func (o DocumentDBPrivateNetworkEndpointMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DocumentDBPrivateNetworkEndpoint] {
	return pulumix.Output[map[string]*DocumentDBPrivateNetworkEndpoint]{
		OutputState: o.OutputState,
	}
}

func (o DocumentDBPrivateNetworkEndpointMapOutput) MapIndex(k pulumi.StringInput) DocumentDBPrivateNetworkEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentDBPrivateNetworkEndpoint {
		return vs[0].(map[string]*DocumentDBPrivateNetworkEndpoint)[vs[1].(string)]
	}).(DocumentDBPrivateNetworkEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBPrivateNetworkEndpointInput)(nil)).Elem(), &DocumentDBPrivateNetworkEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBPrivateNetworkEndpointArrayInput)(nil)).Elem(), DocumentDBPrivateNetworkEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBPrivateNetworkEndpointMapInput)(nil)).Elem(), DocumentDBPrivateNetworkEndpointMap{})
	pulumi.RegisterOutputType(DocumentDBPrivateNetworkEndpointOutput{})
	pulumi.RegisterOutputType(DocumentDBPrivateNetworkEndpointArrayOutput{})
	pulumi.RegisterOutputType(DocumentDBPrivateNetworkEndpointMapOutput{})
}