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

// ## Import
//
// IoT devices can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/iotDevice:IotDevice device01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type IotDevice struct {
	pulumi.CustomResourceState

	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure pulumi.BoolPtrOutput `pulumi:"allowInsecure"`
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections pulumi.BoolPtrOutput `pulumi:"allowMultipleConnections"`
	// The certificate bundle of the device.
	Certificate IotDeviceCertificateOutput `pulumi:"certificate"`
	// The date and time the device was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the IoT device (e.g. `living room`).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the hub on which this device will be created.
	HubId pulumi.StringOutput `pulumi:"hubId"`
	// The current connection status of the device.
	IsConnected pulumi.BoolOutput `pulumi:"isConnected"`
	// The last MQTT activity of the device.
	LastActivityAt pulumi.StringOutput `pulumi:"lastActivityAt"`
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters IotDeviceMessageFiltersPtrOutput `pulumi:"messageFilters"`
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// The current status of the device.
	Status pulumi.StringOutput `pulumi:"status"`
	// The date and time the device resource was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewIotDevice registers a new resource with the given unique name, arguments, and options.
func NewIotDevice(ctx *pulumi.Context,
	name string, args *IotDeviceArgs, opts ...pulumi.ResourceOption) (*IotDevice, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HubId == nil {
		return nil, errors.New("invalid value for required argument 'HubId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IotDevice
	err := ctx.RegisterResource("scaleway:index/iotDevice:IotDevice", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotDevice gets an existing IotDevice resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotDeviceState, opts ...pulumi.ResourceOption) (*IotDevice, error) {
	var resource IotDevice
	err := ctx.ReadResource("scaleway:index/iotDevice:IotDevice", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotDevice resources.
type iotDeviceState struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure *bool `pulumi:"allowInsecure"`
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections *bool `pulumi:"allowMultipleConnections"`
	// The certificate bundle of the device.
	Certificate *IotDeviceCertificate `pulumi:"certificate"`
	// The date and time the device was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the IoT device (e.g. `living room`).
	Description *string `pulumi:"description"`
	// The ID of the hub on which this device will be created.
	HubId *string `pulumi:"hubId"`
	// The current connection status of the device.
	IsConnected *bool `pulumi:"isConnected"`
	// The last MQTT activity of the device.
	LastActivityAt *string `pulumi:"lastActivityAt"`
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters *IotDeviceMessageFilters `pulumi:"messageFilters"`
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name *string `pulumi:"name"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// The current status of the device.
	Status *string `pulumi:"status"`
	// The date and time the device resource was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type IotDeviceState struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure pulumi.BoolPtrInput
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections pulumi.BoolPtrInput
	// The certificate bundle of the device.
	Certificate IotDeviceCertificatePtrInput
	// The date and time the device was created.
	CreatedAt pulumi.StringPtrInput
	// The description of the IoT device (e.g. `living room`).
	Description pulumi.StringPtrInput
	// The ID of the hub on which this device will be created.
	HubId pulumi.StringPtrInput
	// The current connection status of the device.
	IsConnected pulumi.BoolPtrInput
	// The last MQTT activity of the device.
	LastActivityAt pulumi.StringPtrInput
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters IotDeviceMessageFiltersPtrInput
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// The current status of the device.
	Status pulumi.StringPtrInput
	// The date and time the device resource was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (IotDeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDeviceState)(nil)).Elem()
}

type iotDeviceArgs struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure *bool `pulumi:"allowInsecure"`
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections *bool `pulumi:"allowMultipleConnections"`
	// The certificate bundle of the device.
	Certificate *IotDeviceCertificate `pulumi:"certificate"`
	// The description of the IoT device (e.g. `living room`).
	Description *string `pulumi:"description"`
	// The ID of the hub on which this device will be created.
	HubId string `pulumi:"hubId"`
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters *IotDeviceMessageFilters `pulumi:"messageFilters"`
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name *string `pulumi:"name"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a IotDevice resource.
type IotDeviceArgs struct {
	// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
	//
	// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
	AllowInsecure pulumi.BoolPtrInput
	// Allow more than one simultaneous connection using the same device credentials.
	//
	// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
	AllowMultipleConnections pulumi.BoolPtrInput
	// The certificate bundle of the device.
	Certificate IotDeviceCertificatePtrInput
	// The description of the IoT device (e.g. `living room`).
	Description pulumi.StringPtrInput
	// The ID of the hub on which this device will be created.
	HubId pulumi.StringInput
	// Rules that define which messages are authorized or denied based on their topic.
	MessageFilters IotDeviceMessageFiltersPtrInput
	// The name of the IoT device you want to create (e.g. `my-device`).
	//
	// > **Important:** Updates to `name` will destroy and recreate a new resource.
	Name pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
}

func (IotDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDeviceArgs)(nil)).Elem()
}

type IotDeviceInput interface {
	pulumi.Input

	ToIotDeviceOutput() IotDeviceOutput
	ToIotDeviceOutputWithContext(ctx context.Context) IotDeviceOutput
}

func (*IotDevice) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDevice)(nil)).Elem()
}

func (i *IotDevice) ToIotDeviceOutput() IotDeviceOutput {
	return i.ToIotDeviceOutputWithContext(context.Background())
}

func (i *IotDevice) ToIotDeviceOutputWithContext(ctx context.Context) IotDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDeviceOutput)
}

func (i *IotDevice) ToOutput(ctx context.Context) pulumix.Output[*IotDevice] {
	return pulumix.Output[*IotDevice]{
		OutputState: i.ToIotDeviceOutputWithContext(ctx).OutputState,
	}
}

// IotDeviceArrayInput is an input type that accepts IotDeviceArray and IotDeviceArrayOutput values.
// You can construct a concrete instance of `IotDeviceArrayInput` via:
//
//	IotDeviceArray{ IotDeviceArgs{...} }
type IotDeviceArrayInput interface {
	pulumi.Input

	ToIotDeviceArrayOutput() IotDeviceArrayOutput
	ToIotDeviceArrayOutputWithContext(context.Context) IotDeviceArrayOutput
}

type IotDeviceArray []IotDeviceInput

func (IotDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotDevice)(nil)).Elem()
}

func (i IotDeviceArray) ToIotDeviceArrayOutput() IotDeviceArrayOutput {
	return i.ToIotDeviceArrayOutputWithContext(context.Background())
}

func (i IotDeviceArray) ToIotDeviceArrayOutputWithContext(ctx context.Context) IotDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDeviceArrayOutput)
}

func (i IotDeviceArray) ToOutput(ctx context.Context) pulumix.Output[[]*IotDevice] {
	return pulumix.Output[[]*IotDevice]{
		OutputState: i.ToIotDeviceArrayOutputWithContext(ctx).OutputState,
	}
}

// IotDeviceMapInput is an input type that accepts IotDeviceMap and IotDeviceMapOutput values.
// You can construct a concrete instance of `IotDeviceMapInput` via:
//
//	IotDeviceMap{ "key": IotDeviceArgs{...} }
type IotDeviceMapInput interface {
	pulumi.Input

	ToIotDeviceMapOutput() IotDeviceMapOutput
	ToIotDeviceMapOutputWithContext(context.Context) IotDeviceMapOutput
}

type IotDeviceMap map[string]IotDeviceInput

func (IotDeviceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotDevice)(nil)).Elem()
}

func (i IotDeviceMap) ToIotDeviceMapOutput() IotDeviceMapOutput {
	return i.ToIotDeviceMapOutputWithContext(context.Background())
}

func (i IotDeviceMap) ToIotDeviceMapOutputWithContext(ctx context.Context) IotDeviceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDeviceMapOutput)
}

func (i IotDeviceMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IotDevice] {
	return pulumix.Output[map[string]*IotDevice]{
		OutputState: i.ToIotDeviceMapOutputWithContext(ctx).OutputState,
	}
}

type IotDeviceOutput struct{ *pulumi.OutputState }

func (IotDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDevice)(nil)).Elem()
}

func (o IotDeviceOutput) ToIotDeviceOutput() IotDeviceOutput {
	return o
}

func (o IotDeviceOutput) ToIotDeviceOutputWithContext(ctx context.Context) IotDeviceOutput {
	return o
}

func (o IotDeviceOutput) ToOutput(ctx context.Context) pulumix.Output[*IotDevice] {
	return pulumix.Output[*IotDevice]{
		OutputState: o.OutputState,
	}
}

// Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
//
// > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
func (o IotDeviceOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.BoolPtrOutput { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

// Allow more than one simultaneous connection using the same device credentials.
//
// > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
func (o IotDeviceOutput) AllowMultipleConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.BoolPtrOutput { return v.AllowMultipleConnections }).(pulumi.BoolPtrOutput)
}

// The certificate bundle of the device.
func (o IotDeviceOutput) Certificate() IotDeviceCertificateOutput {
	return o.ApplyT(func(v *IotDevice) IotDeviceCertificateOutput { return v.Certificate }).(IotDeviceCertificateOutput)
}

// The date and time the device was created.
func (o IotDeviceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the IoT device (e.g. `living room`).
func (o IotDeviceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the hub on which this device will be created.
func (o IotDeviceOutput) HubId() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.HubId }).(pulumi.StringOutput)
}

// The current connection status of the device.
func (o IotDeviceOutput) IsConnected() pulumi.BoolOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.BoolOutput { return v.IsConnected }).(pulumi.BoolOutput)
}

// The last MQTT activity of the device.
func (o IotDeviceOutput) LastActivityAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.LastActivityAt }).(pulumi.StringOutput)
}

// Rules that define which messages are authorized or denied based on their topic.
func (o IotDeviceOutput) MessageFilters() IotDeviceMessageFiltersPtrOutput {
	return o.ApplyT(func(v *IotDevice) IotDeviceMessageFiltersPtrOutput { return v.MessageFilters }).(IotDeviceMessageFiltersPtrOutput)
}

// The name of the IoT device you want to create (e.g. `my-device`).
//
// > **Important:** Updates to `name` will destroy and recreate a new resource.
func (o IotDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o IotDeviceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current status of the device.
func (o IotDeviceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The date and time the device resource was updated.
func (o IotDeviceOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *IotDevice) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type IotDeviceArrayOutput struct{ *pulumi.OutputState }

func (IotDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IotDevice)(nil)).Elem()
}

func (o IotDeviceArrayOutput) ToIotDeviceArrayOutput() IotDeviceArrayOutput {
	return o
}

func (o IotDeviceArrayOutput) ToIotDeviceArrayOutputWithContext(ctx context.Context) IotDeviceArrayOutput {
	return o
}

func (o IotDeviceArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IotDevice] {
	return pulumix.Output[[]*IotDevice]{
		OutputState: o.OutputState,
	}
}

func (o IotDeviceArrayOutput) Index(i pulumi.IntInput) IotDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IotDevice {
		return vs[0].([]*IotDevice)[vs[1].(int)]
	}).(IotDeviceOutput)
}

type IotDeviceMapOutput struct{ *pulumi.OutputState }

func (IotDeviceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IotDevice)(nil)).Elem()
}

func (o IotDeviceMapOutput) ToIotDeviceMapOutput() IotDeviceMapOutput {
	return o
}

func (o IotDeviceMapOutput) ToIotDeviceMapOutputWithContext(ctx context.Context) IotDeviceMapOutput {
	return o
}

func (o IotDeviceMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IotDevice] {
	return pulumix.Output[map[string]*IotDevice]{
		OutputState: o.OutputState,
	}
}

func (o IotDeviceMapOutput) MapIndex(k pulumi.StringInput) IotDeviceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IotDevice {
		return vs[0].(map[string]*IotDevice)[vs[1].(string)]
	}).(IotDeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IotDeviceInput)(nil)).Elem(), &IotDevice{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotDeviceArrayInput)(nil)).Elem(), IotDeviceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotDeviceMapInput)(nil)).Elem(), IotDeviceMap{})
	pulumi.RegisterOutputType(IotDeviceOutput{})
	pulumi.RegisterOutputType(IotDeviceArrayOutput{})
	pulumi.RegisterOutputType(IotDeviceMapOutput{})
}
