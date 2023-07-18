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

// Create and manage Scaleway RDB database privilege.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#user-and-permissions).
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
//			mainRdbInstance, err := scaleway.NewRdbInstance(ctx, "mainRdbInstance", &scaleway.RdbInstanceArgs{
//				NodeType:      pulumi.String("DB-DEV-S"),
//				Engine:        pulumi.String("PostgreSQL-11"),
//				IsHaCluster:   pulumi.Bool(true),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			mainRdbDatabase, err := scaleway.NewRdbDatabase(ctx, "mainRdbDatabase", &scaleway.RdbDatabaseArgs{
//				InstanceId: mainRdbInstance.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			mainRdbUser, err := scaleway.NewRdbUser(ctx, "mainRdbUser", &scaleway.RdbUserArgs{
//				InstanceId: mainRdbInstance.ID(),
//				Password:   pulumi.String("thiZ_is_v&ry_s3cret"),
//				IsAdmin:    pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewRdbPrivilege(ctx, "mainRdbPrivilege", &scaleway.RdbPrivilegeArgs{
//				InstanceId:   mainRdbInstance.ID(),
//				UserName:     pulumi.String("my-db-user"),
//				DatabaseName: pulumi.String("my-db-name"),
//				Permission:   pulumi.String("all"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				mainRdbUser,
//				mainRdbDatabase,
//			}))
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
// The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/rdbPrivilege:RdbPrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
//
// ```
type RdbPrivilege struct {
	pulumi.CustomResourceState

	// Name of the database (e.g. `my-db-name`).
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// UUID of the rdb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission pulumi.StringOutput `pulumi:"permission"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Name of the user (e.g. `my-db-user`).
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewRdbPrivilege registers a new resource with the given unique name, arguments, and options.
func NewRdbPrivilege(ctx *pulumi.Context,
	name string, args *RdbPrivilegeArgs, opts ...pulumi.ResourceOption) (*RdbPrivilege, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Permission == nil {
		return nil, errors.New("invalid value for required argument 'Permission'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdbPrivilege
	err := ctx.RegisterResource("scaleway:index/rdbPrivilege:RdbPrivilege", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbPrivilege gets an existing RdbPrivilege resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbPrivilege(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbPrivilegeState, opts ...pulumi.ResourceOption) (*RdbPrivilege, error) {
	var resource RdbPrivilege
	err := ctx.ReadResource("scaleway:index/rdbPrivilege:RdbPrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbPrivilege resources.
type rdbPrivilegeState struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName *string `pulumi:"databaseName"`
	// UUID of the rdb instance.
	InstanceId *string `pulumi:"instanceId"`
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission *string `pulumi:"permission"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Name of the user (e.g. `my-db-user`).
	UserName *string `pulumi:"userName"`
}

type RdbPrivilegeState struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName pulumi.StringPtrInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringPtrInput
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Name of the user (e.g. `my-db-user`).
	UserName pulumi.StringPtrInput
}

func (RdbPrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbPrivilegeState)(nil)).Elem()
}

type rdbPrivilegeArgs struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName string `pulumi:"databaseName"`
	// UUID of the rdb instance.
	InstanceId string `pulumi:"instanceId"`
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission string `pulumi:"permission"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Name of the user (e.g. `my-db-user`).
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a RdbPrivilege resource.
type RdbPrivilegeArgs struct {
	// Name of the database (e.g. `my-db-name`).
	DatabaseName pulumi.StringInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringInput
	// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
	Permission pulumi.StringInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Name of the user (e.g. `my-db-user`).
	UserName pulumi.StringInput
}

func (RdbPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbPrivilegeArgs)(nil)).Elem()
}

type RdbPrivilegeInput interface {
	pulumi.Input

	ToRdbPrivilegeOutput() RdbPrivilegeOutput
	ToRdbPrivilegeOutputWithContext(ctx context.Context) RdbPrivilegeOutput
}

func (*RdbPrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbPrivilege)(nil)).Elem()
}

func (i *RdbPrivilege) ToRdbPrivilegeOutput() RdbPrivilegeOutput {
	return i.ToRdbPrivilegeOutputWithContext(context.Background())
}

func (i *RdbPrivilege) ToRdbPrivilegeOutputWithContext(ctx context.Context) RdbPrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbPrivilegeOutput)
}

// RdbPrivilegeArrayInput is an input type that accepts RdbPrivilegeArray and RdbPrivilegeArrayOutput values.
// You can construct a concrete instance of `RdbPrivilegeArrayInput` via:
//
//	RdbPrivilegeArray{ RdbPrivilegeArgs{...} }
type RdbPrivilegeArrayInput interface {
	pulumi.Input

	ToRdbPrivilegeArrayOutput() RdbPrivilegeArrayOutput
	ToRdbPrivilegeArrayOutputWithContext(context.Context) RdbPrivilegeArrayOutput
}

type RdbPrivilegeArray []RdbPrivilegeInput

func (RdbPrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbPrivilege)(nil)).Elem()
}

func (i RdbPrivilegeArray) ToRdbPrivilegeArrayOutput() RdbPrivilegeArrayOutput {
	return i.ToRdbPrivilegeArrayOutputWithContext(context.Background())
}

func (i RdbPrivilegeArray) ToRdbPrivilegeArrayOutputWithContext(ctx context.Context) RdbPrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbPrivilegeArrayOutput)
}

// RdbPrivilegeMapInput is an input type that accepts RdbPrivilegeMap and RdbPrivilegeMapOutput values.
// You can construct a concrete instance of `RdbPrivilegeMapInput` via:
//
//	RdbPrivilegeMap{ "key": RdbPrivilegeArgs{...} }
type RdbPrivilegeMapInput interface {
	pulumi.Input

	ToRdbPrivilegeMapOutput() RdbPrivilegeMapOutput
	ToRdbPrivilegeMapOutputWithContext(context.Context) RdbPrivilegeMapOutput
}

type RdbPrivilegeMap map[string]RdbPrivilegeInput

func (RdbPrivilegeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbPrivilege)(nil)).Elem()
}

func (i RdbPrivilegeMap) ToRdbPrivilegeMapOutput() RdbPrivilegeMapOutput {
	return i.ToRdbPrivilegeMapOutputWithContext(context.Background())
}

func (i RdbPrivilegeMap) ToRdbPrivilegeMapOutputWithContext(ctx context.Context) RdbPrivilegeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbPrivilegeMapOutput)
}

type RdbPrivilegeOutput struct{ *pulumi.OutputState }

func (RdbPrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbPrivilege)(nil)).Elem()
}

func (o RdbPrivilegeOutput) ToRdbPrivilegeOutput() RdbPrivilegeOutput {
	return o
}

func (o RdbPrivilegeOutput) ToRdbPrivilegeOutputWithContext(ctx context.Context) RdbPrivilegeOutput {
	return o
}

// Name of the database (e.g. `my-db-name`).
func (o RdbPrivilegeOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbPrivilege) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// UUID of the rdb instance.
func (o RdbPrivilegeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbPrivilege) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
func (o RdbPrivilegeOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbPrivilege) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o RdbPrivilegeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbPrivilege) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Name of the user (e.g. `my-db-user`).
func (o RdbPrivilegeOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbPrivilege) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type RdbPrivilegeArrayOutput struct{ *pulumi.OutputState }

func (RdbPrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbPrivilege)(nil)).Elem()
}

func (o RdbPrivilegeArrayOutput) ToRdbPrivilegeArrayOutput() RdbPrivilegeArrayOutput {
	return o
}

func (o RdbPrivilegeArrayOutput) ToRdbPrivilegeArrayOutputWithContext(ctx context.Context) RdbPrivilegeArrayOutput {
	return o
}

func (o RdbPrivilegeArrayOutput) Index(i pulumi.IntInput) RdbPrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdbPrivilege {
		return vs[0].([]*RdbPrivilege)[vs[1].(int)]
	}).(RdbPrivilegeOutput)
}

type RdbPrivilegeMapOutput struct{ *pulumi.OutputState }

func (RdbPrivilegeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbPrivilege)(nil)).Elem()
}

func (o RdbPrivilegeMapOutput) ToRdbPrivilegeMapOutput() RdbPrivilegeMapOutput {
	return o
}

func (o RdbPrivilegeMapOutput) ToRdbPrivilegeMapOutputWithContext(ctx context.Context) RdbPrivilegeMapOutput {
	return o
}

func (o RdbPrivilegeMapOutput) MapIndex(k pulumi.StringInput) RdbPrivilegeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdbPrivilege {
		return vs[0].(map[string]*RdbPrivilege)[vs[1].(string)]
	}).(RdbPrivilegeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdbPrivilegeInput)(nil)).Elem(), &RdbPrivilege{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbPrivilegeArrayInput)(nil)).Elem(), RdbPrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbPrivilegeMapInput)(nil)).Elem(), RdbPrivilegeMap{})
	pulumi.RegisterOutputType(RdbPrivilegeOutput{})
	pulumi.RegisterOutputType(RdbPrivilegeArrayOutput{})
	pulumi.RegisterOutputType(RdbPrivilegeMapOutput{})
}
