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

// Create and manage Scaleway DocumentDB database privilege.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
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
//			_, err := scaleway.NewDocumentDBPrivilege(ctx, "main", &scaleway.DocumentDBPrivilegeArgs{
//				DatabaseName: pulumi.String("my-db-name"),
//				InstanceId:   pulumi.String("11111111-1111-1111-1111-111111111111"),
//				Permission:   pulumi.String("all"),
//				UserName:     pulumi.String("my-db-user"),
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
// The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/documentDBPrivilege:DocumentDBPrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
//
// ```
type DocumentDBPrivilege struct {
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

// NewDocumentDBPrivilege registers a new resource with the given unique name, arguments, and options.
func NewDocumentDBPrivilege(ctx *pulumi.Context,
	name string, args *DocumentDBPrivilegeArgs, opts ...pulumi.ResourceOption) (*DocumentDBPrivilege, error) {
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
	var resource DocumentDBPrivilege
	err := ctx.RegisterResource("scaleway:index/documentDBPrivilege:DocumentDBPrivilege", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentDBPrivilege gets an existing DocumentDBPrivilege resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentDBPrivilege(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentDBPrivilegeState, opts ...pulumi.ResourceOption) (*DocumentDBPrivilege, error) {
	var resource DocumentDBPrivilege
	err := ctx.ReadResource("scaleway:index/documentDBPrivilege:DocumentDBPrivilege", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentDBPrivilege resources.
type documentDBPrivilegeState struct {
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

type DocumentDBPrivilegeState struct {
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

func (DocumentDBPrivilegeState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentDBPrivilegeState)(nil)).Elem()
}

type documentDBPrivilegeArgs struct {
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

// The set of arguments for constructing a DocumentDBPrivilege resource.
type DocumentDBPrivilegeArgs struct {
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

func (DocumentDBPrivilegeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentDBPrivilegeArgs)(nil)).Elem()
}

type DocumentDBPrivilegeInput interface {
	pulumi.Input

	ToDocumentDBPrivilegeOutput() DocumentDBPrivilegeOutput
	ToDocumentDBPrivilegeOutputWithContext(ctx context.Context) DocumentDBPrivilegeOutput
}

func (*DocumentDBPrivilege) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentDBPrivilege)(nil)).Elem()
}

func (i *DocumentDBPrivilege) ToDocumentDBPrivilegeOutput() DocumentDBPrivilegeOutput {
	return i.ToDocumentDBPrivilegeOutputWithContext(context.Background())
}

func (i *DocumentDBPrivilege) ToDocumentDBPrivilegeOutputWithContext(ctx context.Context) DocumentDBPrivilegeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBPrivilegeOutput)
}

// DocumentDBPrivilegeArrayInput is an input type that accepts DocumentDBPrivilegeArray and DocumentDBPrivilegeArrayOutput values.
// You can construct a concrete instance of `DocumentDBPrivilegeArrayInput` via:
//
//	DocumentDBPrivilegeArray{ DocumentDBPrivilegeArgs{...} }
type DocumentDBPrivilegeArrayInput interface {
	pulumi.Input

	ToDocumentDBPrivilegeArrayOutput() DocumentDBPrivilegeArrayOutput
	ToDocumentDBPrivilegeArrayOutputWithContext(context.Context) DocumentDBPrivilegeArrayOutput
}

type DocumentDBPrivilegeArray []DocumentDBPrivilegeInput

func (DocumentDBPrivilegeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentDBPrivilege)(nil)).Elem()
}

func (i DocumentDBPrivilegeArray) ToDocumentDBPrivilegeArrayOutput() DocumentDBPrivilegeArrayOutput {
	return i.ToDocumentDBPrivilegeArrayOutputWithContext(context.Background())
}

func (i DocumentDBPrivilegeArray) ToDocumentDBPrivilegeArrayOutputWithContext(ctx context.Context) DocumentDBPrivilegeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBPrivilegeArrayOutput)
}

// DocumentDBPrivilegeMapInput is an input type that accepts DocumentDBPrivilegeMap and DocumentDBPrivilegeMapOutput values.
// You can construct a concrete instance of `DocumentDBPrivilegeMapInput` via:
//
//	DocumentDBPrivilegeMap{ "key": DocumentDBPrivilegeArgs{...} }
type DocumentDBPrivilegeMapInput interface {
	pulumi.Input

	ToDocumentDBPrivilegeMapOutput() DocumentDBPrivilegeMapOutput
	ToDocumentDBPrivilegeMapOutputWithContext(context.Context) DocumentDBPrivilegeMapOutput
}

type DocumentDBPrivilegeMap map[string]DocumentDBPrivilegeInput

func (DocumentDBPrivilegeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentDBPrivilege)(nil)).Elem()
}

func (i DocumentDBPrivilegeMap) ToDocumentDBPrivilegeMapOutput() DocumentDBPrivilegeMapOutput {
	return i.ToDocumentDBPrivilegeMapOutputWithContext(context.Background())
}

func (i DocumentDBPrivilegeMap) ToDocumentDBPrivilegeMapOutputWithContext(ctx context.Context) DocumentDBPrivilegeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDBPrivilegeMapOutput)
}

type DocumentDBPrivilegeOutput struct{ *pulumi.OutputState }

func (DocumentDBPrivilegeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentDBPrivilege)(nil)).Elem()
}

func (o DocumentDBPrivilegeOutput) ToDocumentDBPrivilegeOutput() DocumentDBPrivilegeOutput {
	return o
}

func (o DocumentDBPrivilegeOutput) ToDocumentDBPrivilegeOutputWithContext(ctx context.Context) DocumentDBPrivilegeOutput {
	return o
}

// Name of the database (e.g. `my-db-name`).
func (o DocumentDBPrivilegeOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivilege) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// UUID of the rdb instance.
func (o DocumentDBPrivilegeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivilege) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
func (o DocumentDBPrivilegeOutput) Permission() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivilege) pulumi.StringOutput { return v.Permission }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o DocumentDBPrivilegeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivilege) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Name of the user (e.g. `my-db-user`).
func (o DocumentDBPrivilegeOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentDBPrivilege) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type DocumentDBPrivilegeArrayOutput struct{ *pulumi.OutputState }

func (DocumentDBPrivilegeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentDBPrivilege)(nil)).Elem()
}

func (o DocumentDBPrivilegeArrayOutput) ToDocumentDBPrivilegeArrayOutput() DocumentDBPrivilegeArrayOutput {
	return o
}

func (o DocumentDBPrivilegeArrayOutput) ToDocumentDBPrivilegeArrayOutputWithContext(ctx context.Context) DocumentDBPrivilegeArrayOutput {
	return o
}

func (o DocumentDBPrivilegeArrayOutput) Index(i pulumi.IntInput) DocumentDBPrivilegeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentDBPrivilege {
		return vs[0].([]*DocumentDBPrivilege)[vs[1].(int)]
	}).(DocumentDBPrivilegeOutput)
}

type DocumentDBPrivilegeMapOutput struct{ *pulumi.OutputState }

func (DocumentDBPrivilegeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentDBPrivilege)(nil)).Elem()
}

func (o DocumentDBPrivilegeMapOutput) ToDocumentDBPrivilegeMapOutput() DocumentDBPrivilegeMapOutput {
	return o
}

func (o DocumentDBPrivilegeMapOutput) ToDocumentDBPrivilegeMapOutputWithContext(ctx context.Context) DocumentDBPrivilegeMapOutput {
	return o
}

func (o DocumentDBPrivilegeMapOutput) MapIndex(k pulumi.StringInput) DocumentDBPrivilegeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentDBPrivilege {
		return vs[0].(map[string]*DocumentDBPrivilege)[vs[1].(string)]
	}).(DocumentDBPrivilegeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBPrivilegeInput)(nil)).Elem(), &DocumentDBPrivilege{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBPrivilegeArrayInput)(nil)).Elem(), DocumentDBPrivilegeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentDBPrivilegeMapInput)(nil)).Elem(), DocumentDBPrivilegeMap{})
	pulumi.RegisterOutputType(DocumentDBPrivilegeOutput{})
	pulumi.RegisterOutputType(DocumentDBPrivilegeArrayOutput{})
	pulumi.RegisterOutputType(DocumentDBPrivilegeMapOutput{})
}
