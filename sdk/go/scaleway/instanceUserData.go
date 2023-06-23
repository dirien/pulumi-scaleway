// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Compute Instance User Data values.
//
// User data is a key value store API you can use to provide data from and to your server without authentication. It is the mechanism by which a user can pass information contained in a local file to an Instance at launch time.
//
// The typical use case is to pass something like a shell script or a configuration file as user data.
//
// For more information about [userData](https://developers.scaleway.com/en/products/instance/api/#patch-9ef3ec)  check our documentation guide [here](https://www.scaleway.com/en/docs/compute/instances/how-to/use-boot-modes/#how-to-use-cloud-init).
//
// About cloud-init documentation please check this [link](https://cloudinit.readthedocs.io/en/latest/).
//
// ## Examples
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			userData := map[string]interface{}{
//				"cloud-init": "#cloud-config\napt-update: true\napt-upgrade: true\n",
//				"foo":        "bar",
//			}
//			if param := cfg.GetBool("userData"); param != nil {
//				userData = param
//			}
//			mainInstanceServer, err := scaleway.NewInstanceServer(ctx, "mainInstanceServer", &scaleway.InstanceServerArgs{
//				Image: pulumi.String("ubuntu_focal"),
//				Type:  pulumi.String("DEV1-S"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewInstanceUserData(ctx, "mainInstanceUserData", &scaleway.InstanceUserDataArgs{
//				ServerId: mainInstanceServer.ID(),
//				Key:      pulumi.String("foo"),
//				Value:    pulumi.String("bar"),
//			})
//			if err != nil {
//				return err
//			}
//			var data []*scaleway.InstanceUserData
//			for index := 0; index < userData; index++ {
//				key0 := index
//				val0 := index
//				__res, err := scaleway.NewInstanceUserData(ctx, fmt.Sprintf("data-%v", key0), &scaleway.InstanceUserDataArgs{
//					ServerId: mainInstanceServer.ID(),
//					Key:      pulumi.Any(key0),
//					Value:    pulumi.Any(val0),
//				})
//				if err != nil {
//					return err
//				}
//				data = append(data, __res)
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// User data can be imported using the `{zone}/{key}/{server_id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/instanceUserData:InstanceUserData main fr-par-1/cloud-init/11111111-1111-1111-1111-111111111111
//
// ```
type InstanceUserData struct {
	pulumi.CustomResourceState

	// Key of the user data.
	Key pulumi.StringOutput `pulumi:"key"`
	// The ID of the server associated with.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// Value associated with your key
	Value pulumi.StringOutput `pulumi:"value"`
	// `zone`) The zone in which the server should be created.
	//
	// > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
	// You can define values using:
	// - string
	// - UTF-8 encoded file content using file
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceUserData registers a new resource with the given unique name, arguments, and options.
func NewInstanceUserData(ctx *pulumi.Context,
	name string, args *InstanceUserDataArgs, opts ...pulumi.ResourceOption) (*InstanceUserData, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource InstanceUserData
	err := ctx.RegisterResource("scaleway:index/instanceUserData:InstanceUserData", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceUserData gets an existing InstanceUserData resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceUserData(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceUserDataState, opts ...pulumi.ResourceOption) (*InstanceUserData, error) {
	var resource InstanceUserData
	err := ctx.ReadResource("scaleway:index/instanceUserData:InstanceUserData", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceUserData resources.
type instanceUserDataState struct {
	// Key of the user data.
	Key *string `pulumi:"key"`
	// The ID of the server associated with.
	ServerId *string `pulumi:"serverId"`
	// Value associated with your key
	Value *string `pulumi:"value"`
	// `zone`) The zone in which the server should be created.
	//
	// > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
	// You can define values using:
	// - string
	// - UTF-8 encoded file content using file
	Zone *string `pulumi:"zone"`
}

type InstanceUserDataState struct {
	// Key of the user data.
	Key pulumi.StringPtrInput
	// The ID of the server associated with.
	ServerId pulumi.StringPtrInput
	// Value associated with your key
	Value pulumi.StringPtrInput
	// `zone`) The zone in which the server should be created.
	//
	// > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
	// You can define values using:
	// - string
	// - UTF-8 encoded file content using file
	Zone pulumi.StringPtrInput
}

func (InstanceUserDataState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceUserDataState)(nil)).Elem()
}

type instanceUserDataArgs struct {
	// Key of the user data.
	Key string `pulumi:"key"`
	// The ID of the server associated with.
	ServerId string `pulumi:"serverId"`
	// Value associated with your key
	Value string `pulumi:"value"`
	// `zone`) The zone in which the server should be created.
	//
	// > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
	// You can define values using:
	// - string
	// - UTF-8 encoded file content using file
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceUserData resource.
type InstanceUserDataArgs struct {
	// Key of the user data.
	Key pulumi.StringInput
	// The ID of the server associated with.
	ServerId pulumi.StringInput
	// Value associated with your key
	Value pulumi.StringInput
	// `zone`) The zone in which the server should be created.
	//
	// > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
	// You can define values using:
	// - string
	// - UTF-8 encoded file content using file
	Zone pulumi.StringPtrInput
}

func (InstanceUserDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceUserDataArgs)(nil)).Elem()
}

type InstanceUserDataInput interface {
	pulumi.Input

	ToInstanceUserDataOutput() InstanceUserDataOutput
	ToInstanceUserDataOutputWithContext(ctx context.Context) InstanceUserDataOutput
}

func (*InstanceUserData) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceUserData)(nil)).Elem()
}

func (i *InstanceUserData) ToInstanceUserDataOutput() InstanceUserDataOutput {
	return i.ToInstanceUserDataOutputWithContext(context.Background())
}

func (i *InstanceUserData) ToInstanceUserDataOutputWithContext(ctx context.Context) InstanceUserDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceUserDataOutput)
}

// InstanceUserDataArrayInput is an input type that accepts InstanceUserDataArray and InstanceUserDataArrayOutput values.
// You can construct a concrete instance of `InstanceUserDataArrayInput` via:
//
//	InstanceUserDataArray{ InstanceUserDataArgs{...} }
type InstanceUserDataArrayInput interface {
	pulumi.Input

	ToInstanceUserDataArrayOutput() InstanceUserDataArrayOutput
	ToInstanceUserDataArrayOutputWithContext(context.Context) InstanceUserDataArrayOutput
}

type InstanceUserDataArray []InstanceUserDataInput

func (InstanceUserDataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceUserData)(nil)).Elem()
}

func (i InstanceUserDataArray) ToInstanceUserDataArrayOutput() InstanceUserDataArrayOutput {
	return i.ToInstanceUserDataArrayOutputWithContext(context.Background())
}

func (i InstanceUserDataArray) ToInstanceUserDataArrayOutputWithContext(ctx context.Context) InstanceUserDataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceUserDataArrayOutput)
}

// InstanceUserDataMapInput is an input type that accepts InstanceUserDataMap and InstanceUserDataMapOutput values.
// You can construct a concrete instance of `InstanceUserDataMapInput` via:
//
//	InstanceUserDataMap{ "key": InstanceUserDataArgs{...} }
type InstanceUserDataMapInput interface {
	pulumi.Input

	ToInstanceUserDataMapOutput() InstanceUserDataMapOutput
	ToInstanceUserDataMapOutputWithContext(context.Context) InstanceUserDataMapOutput
}

type InstanceUserDataMap map[string]InstanceUserDataInput

func (InstanceUserDataMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceUserData)(nil)).Elem()
}

func (i InstanceUserDataMap) ToInstanceUserDataMapOutput() InstanceUserDataMapOutput {
	return i.ToInstanceUserDataMapOutputWithContext(context.Background())
}

func (i InstanceUserDataMap) ToInstanceUserDataMapOutputWithContext(ctx context.Context) InstanceUserDataMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceUserDataMapOutput)
}

type InstanceUserDataOutput struct{ *pulumi.OutputState }

func (InstanceUserDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceUserData)(nil)).Elem()
}

func (o InstanceUserDataOutput) ToInstanceUserDataOutput() InstanceUserDataOutput {
	return o
}

func (o InstanceUserDataOutput) ToInstanceUserDataOutputWithContext(ctx context.Context) InstanceUserDataOutput {
	return o
}

// Key of the user data.
func (o InstanceUserDataOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceUserData) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The ID of the server associated with.
func (o InstanceUserDataOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceUserData) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// Value associated with your key
func (o InstanceUserDataOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceUserData) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

// `zone`) The zone in which the server should be created.
//
// > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
// You can define values using:
// - string
// - UTF-8 encoded file content using file
func (o InstanceUserDataOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceUserData) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceUserDataArrayOutput struct{ *pulumi.OutputState }

func (InstanceUserDataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceUserData)(nil)).Elem()
}

func (o InstanceUserDataArrayOutput) ToInstanceUserDataArrayOutput() InstanceUserDataArrayOutput {
	return o
}

func (o InstanceUserDataArrayOutput) ToInstanceUserDataArrayOutputWithContext(ctx context.Context) InstanceUserDataArrayOutput {
	return o
}

func (o InstanceUserDataArrayOutput) Index(i pulumi.IntInput) InstanceUserDataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceUserData {
		return vs[0].([]*InstanceUserData)[vs[1].(int)]
	}).(InstanceUserDataOutput)
}

type InstanceUserDataMapOutput struct{ *pulumi.OutputState }

func (InstanceUserDataMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceUserData)(nil)).Elem()
}

func (o InstanceUserDataMapOutput) ToInstanceUserDataMapOutput() InstanceUserDataMapOutput {
	return o
}

func (o InstanceUserDataMapOutput) ToInstanceUserDataMapOutputWithContext(ctx context.Context) InstanceUserDataMapOutput {
	return o
}

func (o InstanceUserDataMapOutput) MapIndex(k pulumi.StringInput) InstanceUserDataOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceUserData {
		return vs[0].(map[string]*InstanceUserData)[vs[1].(string)]
	}).(InstanceUserDataOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceUserDataInput)(nil)).Elem(), &InstanceUserData{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceUserDataArrayInput)(nil)).Elem(), InstanceUserDataArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceUserDataMapInput)(nil)).Elem(), InstanceUserDataMap{})
	pulumi.RegisterOutputType(InstanceUserDataOutput{})
	pulumi.RegisterOutputType(InstanceUserDataArrayOutput{})
	pulumi.RegisterOutputType(InstanceUserDataMapOutput{})
}
