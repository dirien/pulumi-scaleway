// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway RDB database.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := scaleway.NewRdbDatabase(ctx, "main", &scaleway.RdbDatabaseArgs{
//				InstanceId: pulumi.Any(scaleway_rdb_instance.Main.Id),
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
// RDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/rdbDatabase:RdbDatabase rdb01_mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
//
// ```
type RdbDatabase struct {
	pulumi.CustomResourceState

	// UUID of the rdb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed pulumi.BoolOutput `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the owner of the database.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Size of the database (in bytes).
	Size pulumi.StringOutput `pulumi:"size"`
}

// NewRdbDatabase registers a new resource with the given unique name, arguments, and options.
func NewRdbDatabase(ctx *pulumi.Context,
	name string, args *RdbDatabaseArgs, opts ...pulumi.ResourceOption) (*RdbDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RdbDatabase
	err := ctx.RegisterResource("scaleway:index/rdbDatabase:RdbDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbDatabase gets an existing RdbDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbDatabaseState, opts ...pulumi.ResourceOption) (*RdbDatabase, error) {
	var resource RdbDatabase
	err := ctx.ReadResource("scaleway:index/rdbDatabase:RdbDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbDatabase resources.
type rdbDatabaseState struct {
	// UUID of the rdb instance.
	InstanceId *string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed *bool `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// The name of the owner of the database.
	Owner *string `pulumi:"owner"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Size of the database (in bytes).
	Size *string `pulumi:"size"`
}

type RdbDatabaseState struct {
	// UUID of the rdb instance.
	InstanceId pulumi.StringPtrInput
	// Whether the database is managed or not.
	Managed pulumi.BoolPtrInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// The name of the owner of the database.
	Owner pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Size of the database (in bytes).
	Size pulumi.StringPtrInput
}

func (RdbDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbDatabaseState)(nil)).Elem()
}

type rdbDatabaseArgs struct {
	// UUID of the rdb instance.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RdbDatabase resource.
type RdbDatabaseArgs struct {
	// UUID of the rdb instance.
	InstanceId pulumi.StringInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (RdbDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbDatabaseArgs)(nil)).Elem()
}

type RdbDatabaseInput interface {
	pulumi.Input

	ToRdbDatabaseOutput() RdbDatabaseOutput
	ToRdbDatabaseOutputWithContext(ctx context.Context) RdbDatabaseOutput
}

func (*RdbDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbDatabase)(nil)).Elem()
}

func (i *RdbDatabase) ToRdbDatabaseOutput() RdbDatabaseOutput {
	return i.ToRdbDatabaseOutputWithContext(context.Background())
}

func (i *RdbDatabase) ToRdbDatabaseOutputWithContext(ctx context.Context) RdbDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbDatabaseOutput)
}

// RdbDatabaseArrayInput is an input type that accepts RdbDatabaseArray and RdbDatabaseArrayOutput values.
// You can construct a concrete instance of `RdbDatabaseArrayInput` via:
//
//	RdbDatabaseArray{ RdbDatabaseArgs{...} }
type RdbDatabaseArrayInput interface {
	pulumi.Input

	ToRdbDatabaseArrayOutput() RdbDatabaseArrayOutput
	ToRdbDatabaseArrayOutputWithContext(context.Context) RdbDatabaseArrayOutput
}

type RdbDatabaseArray []RdbDatabaseInput

func (RdbDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbDatabase)(nil)).Elem()
}

func (i RdbDatabaseArray) ToRdbDatabaseArrayOutput() RdbDatabaseArrayOutput {
	return i.ToRdbDatabaseArrayOutputWithContext(context.Background())
}

func (i RdbDatabaseArray) ToRdbDatabaseArrayOutputWithContext(ctx context.Context) RdbDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbDatabaseArrayOutput)
}

// RdbDatabaseMapInput is an input type that accepts RdbDatabaseMap and RdbDatabaseMapOutput values.
// You can construct a concrete instance of `RdbDatabaseMapInput` via:
//
//	RdbDatabaseMap{ "key": RdbDatabaseArgs{...} }
type RdbDatabaseMapInput interface {
	pulumi.Input

	ToRdbDatabaseMapOutput() RdbDatabaseMapOutput
	ToRdbDatabaseMapOutputWithContext(context.Context) RdbDatabaseMapOutput
}

type RdbDatabaseMap map[string]RdbDatabaseInput

func (RdbDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbDatabase)(nil)).Elem()
}

func (i RdbDatabaseMap) ToRdbDatabaseMapOutput() RdbDatabaseMapOutput {
	return i.ToRdbDatabaseMapOutputWithContext(context.Background())
}

func (i RdbDatabaseMap) ToRdbDatabaseMapOutputWithContext(ctx context.Context) RdbDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbDatabaseMapOutput)
}

type RdbDatabaseOutput struct{ *pulumi.OutputState }

func (RdbDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbDatabase)(nil)).Elem()
}

func (o RdbDatabaseOutput) ToRdbDatabaseOutput() RdbDatabaseOutput {
	return o
}

func (o RdbDatabaseOutput) ToRdbDatabaseOutputWithContext(ctx context.Context) RdbDatabaseOutput {
	return o
}

// UUID of the rdb instance.
func (o RdbDatabaseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabase) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o RdbDatabaseOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdbDatabase) pulumi.BoolOutput { return v.Managed }).(pulumi.BoolOutput)
}

// Name of the database (e.g. `my-new-database`).
func (o RdbDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the owner of the database.
func (o RdbDatabaseOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabase) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o RdbDatabaseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabase) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Size of the database (in bytes).
func (o RdbDatabaseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabase) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

type RdbDatabaseArrayOutput struct{ *pulumi.OutputState }

func (RdbDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbDatabase)(nil)).Elem()
}

func (o RdbDatabaseArrayOutput) ToRdbDatabaseArrayOutput() RdbDatabaseArrayOutput {
	return o
}

func (o RdbDatabaseArrayOutput) ToRdbDatabaseArrayOutputWithContext(ctx context.Context) RdbDatabaseArrayOutput {
	return o
}

func (o RdbDatabaseArrayOutput) Index(i pulumi.IntInput) RdbDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdbDatabase {
		return vs[0].([]*RdbDatabase)[vs[1].(int)]
	}).(RdbDatabaseOutput)
}

type RdbDatabaseMapOutput struct{ *pulumi.OutputState }

func (RdbDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbDatabase)(nil)).Elem()
}

func (o RdbDatabaseMapOutput) ToRdbDatabaseMapOutput() RdbDatabaseMapOutput {
	return o
}

func (o RdbDatabaseMapOutput) ToRdbDatabaseMapOutputWithContext(ctx context.Context) RdbDatabaseMapOutput {
	return o
}

func (o RdbDatabaseMapOutput) MapIndex(k pulumi.StringInput) RdbDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdbDatabase {
		return vs[0].(map[string]*RdbDatabase)[vs[1].(string)]
	}).(RdbDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdbDatabaseInput)(nil)).Elem(), &RdbDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbDatabaseArrayInput)(nil)).Elem(), RdbDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbDatabaseMapInput)(nil)).Elem(), RdbDatabaseMap{})
	pulumi.RegisterOutputType(RdbDatabaseOutput{})
	pulumi.RegisterOutputType(RdbDatabaseArrayOutput{})
	pulumi.RegisterOutputType(RdbDatabaseMapOutput{})
}
