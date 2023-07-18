// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway RDB database backup.
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
//			_, err := scaleway.NewRdbDatabaseBackup(ctx, "main", &scaleway.RdbDatabaseBackupArgs{
//				InstanceId:   pulumi.Any(data.Scaleway_rdb_instance.Main.Id),
//				DatabaseName: pulumi.Any(data.Scaleway_rdb_database.Main.Name),
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
// ### With expiration
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
//			_, err := scaleway.NewRdbDatabaseBackup(ctx, "main", &scaleway.RdbDatabaseBackupArgs{
//				InstanceId:   pulumi.Any(data.Scaleway_rdb_instance.Main.Id),
//				DatabaseName: pulumi.Any(data.Scaleway_rdb_database.Main.Name),
//				ExpiresAt:    pulumi.String("2022-06-16T07:48:44Z"),
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
// RDB Database can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup mybackup fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type RdbDatabaseBackup struct {
	pulumi.CustomResourceState

	// Creation date (Format ISO 8601).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the database of this backup.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt pulumi.StringPtrOutput `pulumi:"expiresAt"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of the instance of the backup.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// Name of the database (e.g. `my-database`).
	Name pulumi.StringOutput `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Size of the backup (in bytes).
	Size pulumi.IntOutput `pulumi:"size"`
	// Updated date (Format ISO 8601).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewRdbDatabaseBackup registers a new resource with the given unique name, arguments, and options.
func NewRdbDatabaseBackup(ctx *pulumi.Context,
	name string, args *RdbDatabaseBackupArgs, opts ...pulumi.ResourceOption) (*RdbDatabaseBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdbDatabaseBackup
	err := ctx.RegisterResource("scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbDatabaseBackup gets an existing RdbDatabaseBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbDatabaseBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbDatabaseBackupState, opts ...pulumi.ResourceOption) (*RdbDatabaseBackup, error) {
	var resource RdbDatabaseBackup
	err := ctx.ReadResource("scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbDatabaseBackup resources.
type rdbDatabaseBackupState struct {
	// Creation date (Format ISO 8601).
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the database of this backup.
	DatabaseName *string `pulumi:"databaseName"`
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt *string `pulumi:"expiresAt"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the instance of the backup.
	InstanceName *string `pulumi:"instanceName"`
	// Name of the database (e.g. `my-database`).
	Name *string `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Size of the backup (in bytes).
	Size *int `pulumi:"size"`
	// Updated date (Format ISO 8601).
	UpdatedAt *string `pulumi:"updatedAt"`
}

type RdbDatabaseBackupState struct {
	// Creation date (Format ISO 8601).
	CreatedAt pulumi.StringPtrInput
	// Name of the database of this backup.
	DatabaseName pulumi.StringPtrInput
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt pulumi.StringPtrInput
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId pulumi.StringPtrInput
	// Name of the instance of the backup.
	InstanceName pulumi.StringPtrInput
	// Name of the database (e.g. `my-database`).
	Name pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Size of the backup (in bytes).
	Size pulumi.IntPtrInput
	// Updated date (Format ISO 8601).
	UpdatedAt pulumi.StringPtrInput
}

func (RdbDatabaseBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbDatabaseBackupState)(nil)).Elem()
}

type rdbDatabaseBackupArgs struct {
	// Name of the database of this backup.
	DatabaseName string `pulumi:"databaseName"`
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt *string `pulumi:"expiresAt"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database (e.g. `my-database`).
	Name *string `pulumi:"name"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RdbDatabaseBackup resource.
type RdbDatabaseBackupArgs struct {
	// Name of the database of this backup.
	DatabaseName pulumi.StringInput
	// Expiration date (Format ISO 8601).
	//
	// > **Important:** `expiresAt` cannot be removed after being set.
	ExpiresAt pulumi.StringPtrInput
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Backup.
	InstanceId pulumi.StringInput
	// Name of the database (e.g. `my-database`).
	Name pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (RdbDatabaseBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbDatabaseBackupArgs)(nil)).Elem()
}

type RdbDatabaseBackupInput interface {
	pulumi.Input

	ToRdbDatabaseBackupOutput() RdbDatabaseBackupOutput
	ToRdbDatabaseBackupOutputWithContext(ctx context.Context) RdbDatabaseBackupOutput
}

func (*RdbDatabaseBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbDatabaseBackup)(nil)).Elem()
}

func (i *RdbDatabaseBackup) ToRdbDatabaseBackupOutput() RdbDatabaseBackupOutput {
	return i.ToRdbDatabaseBackupOutputWithContext(context.Background())
}

func (i *RdbDatabaseBackup) ToRdbDatabaseBackupOutputWithContext(ctx context.Context) RdbDatabaseBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbDatabaseBackupOutput)
}

// RdbDatabaseBackupArrayInput is an input type that accepts RdbDatabaseBackupArray and RdbDatabaseBackupArrayOutput values.
// You can construct a concrete instance of `RdbDatabaseBackupArrayInput` via:
//
//	RdbDatabaseBackupArray{ RdbDatabaseBackupArgs{...} }
type RdbDatabaseBackupArrayInput interface {
	pulumi.Input

	ToRdbDatabaseBackupArrayOutput() RdbDatabaseBackupArrayOutput
	ToRdbDatabaseBackupArrayOutputWithContext(context.Context) RdbDatabaseBackupArrayOutput
}

type RdbDatabaseBackupArray []RdbDatabaseBackupInput

func (RdbDatabaseBackupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbDatabaseBackup)(nil)).Elem()
}

func (i RdbDatabaseBackupArray) ToRdbDatabaseBackupArrayOutput() RdbDatabaseBackupArrayOutput {
	return i.ToRdbDatabaseBackupArrayOutputWithContext(context.Background())
}

func (i RdbDatabaseBackupArray) ToRdbDatabaseBackupArrayOutputWithContext(ctx context.Context) RdbDatabaseBackupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbDatabaseBackupArrayOutput)
}

// RdbDatabaseBackupMapInput is an input type that accepts RdbDatabaseBackupMap and RdbDatabaseBackupMapOutput values.
// You can construct a concrete instance of `RdbDatabaseBackupMapInput` via:
//
//	RdbDatabaseBackupMap{ "key": RdbDatabaseBackupArgs{...} }
type RdbDatabaseBackupMapInput interface {
	pulumi.Input

	ToRdbDatabaseBackupMapOutput() RdbDatabaseBackupMapOutput
	ToRdbDatabaseBackupMapOutputWithContext(context.Context) RdbDatabaseBackupMapOutput
}

type RdbDatabaseBackupMap map[string]RdbDatabaseBackupInput

func (RdbDatabaseBackupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbDatabaseBackup)(nil)).Elem()
}

func (i RdbDatabaseBackupMap) ToRdbDatabaseBackupMapOutput() RdbDatabaseBackupMapOutput {
	return i.ToRdbDatabaseBackupMapOutputWithContext(context.Background())
}

func (i RdbDatabaseBackupMap) ToRdbDatabaseBackupMapOutputWithContext(ctx context.Context) RdbDatabaseBackupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbDatabaseBackupMapOutput)
}

type RdbDatabaseBackupOutput struct{ *pulumi.OutputState }

func (RdbDatabaseBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbDatabaseBackup)(nil)).Elem()
}

func (o RdbDatabaseBackupOutput) ToRdbDatabaseBackupOutput() RdbDatabaseBackupOutput {
	return o
}

func (o RdbDatabaseBackupOutput) ToRdbDatabaseBackupOutputWithContext(ctx context.Context) RdbDatabaseBackupOutput {
	return o
}

// Creation date (Format ISO 8601).
func (o RdbDatabaseBackupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the database of this backup.
func (o RdbDatabaseBackupOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// Expiration date (Format ISO 8601).
//
// > **Important:** `expiresAt` cannot be removed after being set.
func (o RdbDatabaseBackupOutput) ExpiresAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringPtrOutput { return v.ExpiresAt }).(pulumi.StringPtrOutput)
}

// UUID of the rdb instance.
//
// > **Important:** Updates to `instanceId` will recreate the Backup.
func (o RdbDatabaseBackupOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Name of the instance of the backup.
func (o RdbDatabaseBackupOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// Name of the database (e.g. `my-database`).
func (o RdbDatabaseBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o RdbDatabaseBackupOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Size of the backup (in bytes).
func (o RdbDatabaseBackupOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Updated date (Format ISO 8601).
func (o RdbDatabaseBackupOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbDatabaseBackup) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type RdbDatabaseBackupArrayOutput struct{ *pulumi.OutputState }

func (RdbDatabaseBackupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbDatabaseBackup)(nil)).Elem()
}

func (o RdbDatabaseBackupArrayOutput) ToRdbDatabaseBackupArrayOutput() RdbDatabaseBackupArrayOutput {
	return o
}

func (o RdbDatabaseBackupArrayOutput) ToRdbDatabaseBackupArrayOutputWithContext(ctx context.Context) RdbDatabaseBackupArrayOutput {
	return o
}

func (o RdbDatabaseBackupArrayOutput) Index(i pulumi.IntInput) RdbDatabaseBackupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdbDatabaseBackup {
		return vs[0].([]*RdbDatabaseBackup)[vs[1].(int)]
	}).(RdbDatabaseBackupOutput)
}

type RdbDatabaseBackupMapOutput struct{ *pulumi.OutputState }

func (RdbDatabaseBackupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbDatabaseBackup)(nil)).Elem()
}

func (o RdbDatabaseBackupMapOutput) ToRdbDatabaseBackupMapOutput() RdbDatabaseBackupMapOutput {
	return o
}

func (o RdbDatabaseBackupMapOutput) ToRdbDatabaseBackupMapOutputWithContext(ctx context.Context) RdbDatabaseBackupMapOutput {
	return o
}

func (o RdbDatabaseBackupMapOutput) MapIndex(k pulumi.StringInput) RdbDatabaseBackupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdbDatabaseBackup {
		return vs[0].(map[string]*RdbDatabaseBackup)[vs[1].(string)]
	}).(RdbDatabaseBackupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdbDatabaseBackupInput)(nil)).Elem(), &RdbDatabaseBackup{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbDatabaseBackupArrayInput)(nil)).Elem(), RdbDatabaseBackupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbDatabaseBackupMapInput)(nil)).Elem(), RdbDatabaseBackupMap{})
	pulumi.RegisterOutputType(RdbDatabaseBackupOutput{})
	pulumi.RegisterOutputType(RdbDatabaseBackupArrayOutput{})
	pulumi.RegisterOutputType(RdbDatabaseBackupMapOutput{})
}
