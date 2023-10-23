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

type MnqQueue struct {
	pulumi.CustomResourceState

	// The number of seconds the queue retains a message.
	MessageMaxAge pulumi.IntPtrOutput `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize pulumi.IntPtrOutput `pulumi:"messageMaxSize"`
	// The name of the queue. Conflicts with name_prefix.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// The ID of the Namespace associated to
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// The NATS attributes of the queue
	Nats MnqQueueNatsPtrOutput `pulumi:"nats"`
	// The SQS attributes of the queue
	Sqs MnqQueueSqsPtrOutput `pulumi:"sqs"`
}

// NewMnqQueue registers a new resource with the given unique name, arguments, and options.
func NewMnqQueue(ctx *pulumi.Context,
	name string, args *MnqQueueArgs, opts ...pulumi.ResourceOption) (*MnqQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MnqQueue
	err := ctx.RegisterResource("scaleway:index/mnqQueue:MnqQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMnqQueue gets an existing MnqQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMnqQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MnqQueueState, opts ...pulumi.ResourceOption) (*MnqQueue, error) {
	var resource MnqQueue
	err := ctx.ReadResource("scaleway:index/mnqQueue:MnqQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MnqQueue resources.
type mnqQueueState struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The name of the queue. Conflicts with name_prefix.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the Namespace associated to
	NamespaceId *string `pulumi:"namespaceId"`
	// The NATS attributes of the queue
	Nats *MnqQueueNats `pulumi:"nats"`
	// The SQS attributes of the queue
	Sqs *MnqQueueSqs `pulumi:"sqs"`
}

type MnqQueueState struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize pulumi.IntPtrInput
	// The name of the queue. Conflicts with name_prefix.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix pulumi.StringPtrInput
	// The ID of the Namespace associated to
	NamespaceId pulumi.StringPtrInput
	// The NATS attributes of the queue
	Nats MnqQueueNatsPtrInput
	// The SQS attributes of the queue
	Sqs MnqQueueSqsPtrInput
}

func (MnqQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqQueueState)(nil)).Elem()
}

type mnqQueueArgs struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge *int `pulumi:"messageMaxAge"`
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize *int `pulumi:"messageMaxSize"`
	// The name of the queue. Conflicts with name_prefix.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix *string `pulumi:"namePrefix"`
	// The ID of the Namespace associated to
	NamespaceId string `pulumi:"namespaceId"`
	// The NATS attributes of the queue
	Nats *MnqQueueNats `pulumi:"nats"`
	// The SQS attributes of the queue
	Sqs *MnqQueueSqs `pulumi:"sqs"`
}

// The set of arguments for constructing a MnqQueue resource.
type MnqQueueArgs struct {
	// The number of seconds the queue retains a message.
	MessageMaxAge pulumi.IntPtrInput
	// The maximum size of a message. Should be in bytes.
	MessageMaxSize pulumi.IntPtrInput
	// The name of the queue. Conflicts with name_prefix.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	NamePrefix pulumi.StringPtrInput
	// The ID of the Namespace associated to
	NamespaceId pulumi.StringInput
	// The NATS attributes of the queue
	Nats MnqQueueNatsPtrInput
	// The SQS attributes of the queue
	Sqs MnqQueueSqsPtrInput
}

func (MnqQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mnqQueueArgs)(nil)).Elem()
}

type MnqQueueInput interface {
	pulumi.Input

	ToMnqQueueOutput() MnqQueueOutput
	ToMnqQueueOutputWithContext(ctx context.Context) MnqQueueOutput
}

func (*MnqQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqQueue)(nil)).Elem()
}

func (i *MnqQueue) ToMnqQueueOutput() MnqQueueOutput {
	return i.ToMnqQueueOutputWithContext(context.Background())
}

func (i *MnqQueue) ToMnqQueueOutputWithContext(ctx context.Context) MnqQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqQueueOutput)
}

func (i *MnqQueue) ToOutput(ctx context.Context) pulumix.Output[*MnqQueue] {
	return pulumix.Output[*MnqQueue]{
		OutputState: i.ToMnqQueueOutputWithContext(ctx).OutputState,
	}
}

// MnqQueueArrayInput is an input type that accepts MnqQueueArray and MnqQueueArrayOutput values.
// You can construct a concrete instance of `MnqQueueArrayInput` via:
//
//	MnqQueueArray{ MnqQueueArgs{...} }
type MnqQueueArrayInput interface {
	pulumi.Input

	ToMnqQueueArrayOutput() MnqQueueArrayOutput
	ToMnqQueueArrayOutputWithContext(context.Context) MnqQueueArrayOutput
}

type MnqQueueArray []MnqQueueInput

func (MnqQueueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqQueue)(nil)).Elem()
}

func (i MnqQueueArray) ToMnqQueueArrayOutput() MnqQueueArrayOutput {
	return i.ToMnqQueueArrayOutputWithContext(context.Background())
}

func (i MnqQueueArray) ToMnqQueueArrayOutputWithContext(ctx context.Context) MnqQueueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqQueueArrayOutput)
}

func (i MnqQueueArray) ToOutput(ctx context.Context) pulumix.Output[[]*MnqQueue] {
	return pulumix.Output[[]*MnqQueue]{
		OutputState: i.ToMnqQueueArrayOutputWithContext(ctx).OutputState,
	}
}

// MnqQueueMapInput is an input type that accepts MnqQueueMap and MnqQueueMapOutput values.
// You can construct a concrete instance of `MnqQueueMapInput` via:
//
//	MnqQueueMap{ "key": MnqQueueArgs{...} }
type MnqQueueMapInput interface {
	pulumi.Input

	ToMnqQueueMapOutput() MnqQueueMapOutput
	ToMnqQueueMapOutputWithContext(context.Context) MnqQueueMapOutput
}

type MnqQueueMap map[string]MnqQueueInput

func (MnqQueueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqQueue)(nil)).Elem()
}

func (i MnqQueueMap) ToMnqQueueMapOutput() MnqQueueMapOutput {
	return i.ToMnqQueueMapOutputWithContext(context.Background())
}

func (i MnqQueueMap) ToMnqQueueMapOutputWithContext(ctx context.Context) MnqQueueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MnqQueueMapOutput)
}

func (i MnqQueueMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MnqQueue] {
	return pulumix.Output[map[string]*MnqQueue]{
		OutputState: i.ToMnqQueueMapOutputWithContext(ctx).OutputState,
	}
}

type MnqQueueOutput struct{ *pulumi.OutputState }

func (MnqQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MnqQueue)(nil)).Elem()
}

func (o MnqQueueOutput) ToMnqQueueOutput() MnqQueueOutput {
	return o
}

func (o MnqQueueOutput) ToMnqQueueOutputWithContext(ctx context.Context) MnqQueueOutput {
	return o
}

func (o MnqQueueOutput) ToOutput(ctx context.Context) pulumix.Output[*MnqQueue] {
	return pulumix.Output[*MnqQueue]{
		OutputState: o.OutputState,
	}
}

// The number of seconds the queue retains a message.
func (o MnqQueueOutput) MessageMaxAge() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.IntPtrOutput { return v.MessageMaxAge }).(pulumi.IntPtrOutput)
}

// The maximum size of a message. Should be in bytes.
func (o MnqQueueOutput) MessageMaxSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.IntPtrOutput { return v.MessageMaxSize }).(pulumi.IntPtrOutput)
}

// The name of the queue. Conflicts with name_prefix.
func (o MnqQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. Conflicts with name.
func (o MnqQueueOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// The ID of the Namespace associated to
func (o MnqQueueOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MnqQueue) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// The NATS attributes of the queue
func (o MnqQueueOutput) Nats() MnqQueueNatsPtrOutput {
	return o.ApplyT(func(v *MnqQueue) MnqQueueNatsPtrOutput { return v.Nats }).(MnqQueueNatsPtrOutput)
}

// The SQS attributes of the queue
func (o MnqQueueOutput) Sqs() MnqQueueSqsPtrOutput {
	return o.ApplyT(func(v *MnqQueue) MnqQueueSqsPtrOutput { return v.Sqs }).(MnqQueueSqsPtrOutput)
}

type MnqQueueArrayOutput struct{ *pulumi.OutputState }

func (MnqQueueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MnqQueue)(nil)).Elem()
}

func (o MnqQueueArrayOutput) ToMnqQueueArrayOutput() MnqQueueArrayOutput {
	return o
}

func (o MnqQueueArrayOutput) ToMnqQueueArrayOutputWithContext(ctx context.Context) MnqQueueArrayOutput {
	return o
}

func (o MnqQueueArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MnqQueue] {
	return pulumix.Output[[]*MnqQueue]{
		OutputState: o.OutputState,
	}
}

func (o MnqQueueArrayOutput) Index(i pulumi.IntInput) MnqQueueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MnqQueue {
		return vs[0].([]*MnqQueue)[vs[1].(int)]
	}).(MnqQueueOutput)
}

type MnqQueueMapOutput struct{ *pulumi.OutputState }

func (MnqQueueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MnqQueue)(nil)).Elem()
}

func (o MnqQueueMapOutput) ToMnqQueueMapOutput() MnqQueueMapOutput {
	return o
}

func (o MnqQueueMapOutput) ToMnqQueueMapOutputWithContext(ctx context.Context) MnqQueueMapOutput {
	return o
}

func (o MnqQueueMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MnqQueue] {
	return pulumix.Output[map[string]*MnqQueue]{
		OutputState: o.OutputState,
	}
}

func (o MnqQueueMapOutput) MapIndex(k pulumi.StringInput) MnqQueueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MnqQueue {
		return vs[0].(map[string]*MnqQueue)[vs[1].(string)]
	}).(MnqQueueOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MnqQueueInput)(nil)).Elem(), &MnqQueue{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqQueueArrayInput)(nil)).Elem(), MnqQueueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MnqQueueMapInput)(nil)).Elem(), MnqQueueMap{})
	pulumi.RegisterOutputType(MnqQueueOutput{})
	pulumi.RegisterOutputType(MnqQueueArrayOutput{})
	pulumi.RegisterOutputType(MnqQueueMapOutput{})
}
