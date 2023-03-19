// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Function Triggers. For the moment, the feature is limited to CRON Schedule (time-based).
//
// For more information consult
// the [documentation](https://www.scaleway.com/en/docs/compute/functions/api-cli/fun-uploading-with-serverless-framework/#configuring-events)
// .
//
// For more details about the limitation
// check [functions-limitations](https://www.scaleway.com/en/docs/compute/functions/reference-content/functions-limitations/).
//
// You can check also
// our [functions cron api documentation](https://developers.scaleway.com/en/products/functions/api/#crons-942bf4).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/dirien/pulumi-scaleway/sdk/go/scaleway"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mainFunctionNamespace, err := scaleway.NewFunctionNamespace(ctx, "mainFunctionNamespace", nil)
//			if err != nil {
//				return err
//			}
//			mainFunction, err := scaleway.NewFunction(ctx, "mainFunction", &scaleway.FunctionArgs{
//				NamespaceId: mainFunctionNamespace.ID(),
//				Runtime:     pulumi.String("node14"),
//				Privacy:     pulumi.String("private"),
//				Handler:     pulumi.String("handler.handle"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"test": "scw",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = scaleway.NewFunctionCron(ctx, "mainFunctionCron", &scaleway.FunctionCronArgs{
//				FunctionId: mainFunction.ID(),
//				Schedule:   pulumi.String("0 0 * * *"),
//				Args:       pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"my_var": "terraform",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = scaleway.NewFunctionCron(ctx, "func", &scaleway.FunctionCronArgs{
//				FunctionId: mainFunction.ID(),
//				Schedule:   pulumi.String("0 1 * * *"),
//				Args:       pulumi.String(json1),
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
// Container Cron can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/functionCron:FunctionCron main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type FunctionCron struct {
	pulumi.CustomResourceState

	// The key-value mapping to define arguments that will be passed to your function’s event object
	// during
	Args pulumi.StringOutput `pulumi:"args"`
	// The function ID to link with your cron.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region pulumi.StringOutput `pulumi:"region"`
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// The cron status.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewFunctionCron registers a new resource with the given unique name, arguments, and options.
func NewFunctionCron(ctx *pulumi.Context,
	name string, args *FunctionCronArgs, opts ...pulumi.ResourceOption) (*FunctionCron, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Args == nil {
		return nil, errors.New("invalid value for required argument 'Args'")
	}
	if args.FunctionId == nil {
		return nil, errors.New("invalid value for required argument 'FunctionId'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FunctionCron
	err := ctx.RegisterResource("scaleway:index/functionCron:FunctionCron", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionCron gets an existing FunctionCron resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionCron(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionCronState, opts ...pulumi.ResourceOption) (*FunctionCron, error) {
	var resource FunctionCron
	err := ctx.ReadResource("scaleway:index/functionCron:FunctionCron", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionCron resources.
type functionCronState struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	// during
	Args *string `pulumi:"args"`
	// The function ID to link with your cron.
	FunctionId *string `pulumi:"functionId"`
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region *string `pulumi:"region"`
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule *string `pulumi:"schedule"`
	// The cron status.
	Status *string `pulumi:"status"`
}

type FunctionCronState struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	// during
	Args pulumi.StringPtrInput
	// The function ID to link with your cron.
	FunctionId pulumi.StringPtrInput
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region pulumi.StringPtrInput
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule pulumi.StringPtrInput
	// The cron status.
	Status pulumi.StringPtrInput
}

func (FunctionCronState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionCronState)(nil)).Elem()
}

type functionCronArgs struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	// during
	Args string `pulumi:"args"`
	// The function ID to link with your cron.
	FunctionId string `pulumi:"functionId"`
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region *string `pulumi:"region"`
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule string `pulumi:"schedule"`
}

// The set of arguments for constructing a FunctionCron resource.
type FunctionCronArgs struct {
	// The key-value mapping to define arguments that will be passed to your function’s event object
	// during
	Args pulumi.StringInput
	// The function ID to link with your cron.
	FunctionId pulumi.StringInput
	// (Defaults to provider `region`) The region
	// in where the job was created.
	Region pulumi.StringPtrInput
	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	Schedule pulumi.StringInput
}

func (FunctionCronArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionCronArgs)(nil)).Elem()
}

type FunctionCronInput interface {
	pulumi.Input

	ToFunctionCronOutput() FunctionCronOutput
	ToFunctionCronOutputWithContext(ctx context.Context) FunctionCronOutput
}

func (*FunctionCron) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionCron)(nil)).Elem()
}

func (i *FunctionCron) ToFunctionCronOutput() FunctionCronOutput {
	return i.ToFunctionCronOutputWithContext(context.Background())
}

func (i *FunctionCron) ToFunctionCronOutputWithContext(ctx context.Context) FunctionCronOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionCronOutput)
}

// FunctionCronArrayInput is an input type that accepts FunctionCronArray and FunctionCronArrayOutput values.
// You can construct a concrete instance of `FunctionCronArrayInput` via:
//
//	FunctionCronArray{ FunctionCronArgs{...} }
type FunctionCronArrayInput interface {
	pulumi.Input

	ToFunctionCronArrayOutput() FunctionCronArrayOutput
	ToFunctionCronArrayOutputWithContext(context.Context) FunctionCronArrayOutput
}

type FunctionCronArray []FunctionCronInput

func (FunctionCronArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionCron)(nil)).Elem()
}

func (i FunctionCronArray) ToFunctionCronArrayOutput() FunctionCronArrayOutput {
	return i.ToFunctionCronArrayOutputWithContext(context.Background())
}

func (i FunctionCronArray) ToFunctionCronArrayOutputWithContext(ctx context.Context) FunctionCronArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionCronArrayOutput)
}

// FunctionCronMapInput is an input type that accepts FunctionCronMap and FunctionCronMapOutput values.
// You can construct a concrete instance of `FunctionCronMapInput` via:
//
//	FunctionCronMap{ "key": FunctionCronArgs{...} }
type FunctionCronMapInput interface {
	pulumi.Input

	ToFunctionCronMapOutput() FunctionCronMapOutput
	ToFunctionCronMapOutputWithContext(context.Context) FunctionCronMapOutput
}

type FunctionCronMap map[string]FunctionCronInput

func (FunctionCronMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionCron)(nil)).Elem()
}

func (i FunctionCronMap) ToFunctionCronMapOutput() FunctionCronMapOutput {
	return i.ToFunctionCronMapOutputWithContext(context.Background())
}

func (i FunctionCronMap) ToFunctionCronMapOutputWithContext(ctx context.Context) FunctionCronMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionCronMapOutput)
}

type FunctionCronOutput struct{ *pulumi.OutputState }

func (FunctionCronOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionCron)(nil)).Elem()
}

func (o FunctionCronOutput) ToFunctionCronOutput() FunctionCronOutput {
	return o
}

func (o FunctionCronOutput) ToFunctionCronOutputWithContext(ctx context.Context) FunctionCronOutput {
	return o
}

// The key-value mapping to define arguments that will be passed to your function’s event object
// during
func (o FunctionCronOutput) Args() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Args }).(pulumi.StringOutput)
}

// The function ID to link with your cron.
func (o FunctionCronOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

// (Defaults to provider `region`) The region
// in where the job was created.
func (o FunctionCronOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
// executed.
func (o FunctionCronOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// The cron status.
func (o FunctionCronOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionCron) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type FunctionCronArrayOutput struct{ *pulumi.OutputState }

func (FunctionCronArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionCron)(nil)).Elem()
}

func (o FunctionCronArrayOutput) ToFunctionCronArrayOutput() FunctionCronArrayOutput {
	return o
}

func (o FunctionCronArrayOutput) ToFunctionCronArrayOutputWithContext(ctx context.Context) FunctionCronArrayOutput {
	return o
}

func (o FunctionCronArrayOutput) Index(i pulumi.IntInput) FunctionCronOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionCron {
		return vs[0].([]*FunctionCron)[vs[1].(int)]
	}).(FunctionCronOutput)
}

type FunctionCronMapOutput struct{ *pulumi.OutputState }

func (FunctionCronMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionCron)(nil)).Elem()
}

func (o FunctionCronMapOutput) ToFunctionCronMapOutput() FunctionCronMapOutput {
	return o
}

func (o FunctionCronMapOutput) ToFunctionCronMapOutputWithContext(ctx context.Context) FunctionCronMapOutput {
	return o
}

func (o FunctionCronMapOutput) MapIndex(k pulumi.StringInput) FunctionCronOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionCron {
		return vs[0].(map[string]*FunctionCron)[vs[1].(string)]
	}).(FunctionCronOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionCronInput)(nil)).Elem(), &FunctionCron{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionCronArrayInput)(nil)).Elem(), FunctionCronArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionCronMapInput)(nil)).Elem(), FunctionCronMap{})
	pulumi.RegisterOutputType(FunctionCronOutput{})
	pulumi.RegisterOutputType(FunctionCronArrayOutput{})
	pulumi.RegisterOutputType(FunctionCronMapOutput{})
}
