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

// Creates and manages Scaleway DocumentDB database.
//
// ## Example Usage
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
//			instance, err := scaleway.NewDocumentdbInstance(ctx, "instance", &scaleway.DocumentdbInstanceArgs{
//				NodeType:       pulumi.String("docdb-play2-pico"),
//				Engine:         pulumi.String("FerretDB-1"),
//				UserName:       pulumi.String("my_initial_user"),
//				Password:       pulumi.String("thiZ_is_v&ry_s3cret"),
//				VolumeSizeInGb: pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDocumentdbDatabase(ctx, "main", &scaleway.DocumentdbDatabaseArgs{
//				InstanceId: instance.ID(),
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
// DocumentDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/documentdbDatabase:DocumentdbDatabase mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
// ```
type DocumentdbDatabase struct {
	pulumi.CustomResourceState

	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed pulumi.BoolOutput `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the owner of the database.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region pulumi.StringOutput `pulumi:"region"`
	// Size in gigabytes of the database.
	Size pulumi.StringOutput `pulumi:"size"`
}

// NewDocumentdbDatabase registers a new resource with the given unique name, arguments, and options.
func NewDocumentdbDatabase(ctx *pulumi.Context,
	name string, args *DocumentdbDatabaseArgs, opts ...pulumi.ResourceOption) (*DocumentdbDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentdbDatabase
	err := ctx.RegisterResource("scaleway:index/documentdbDatabase:DocumentdbDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentdbDatabase gets an existing DocumentdbDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentdbDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentdbDatabaseState, opts ...pulumi.ResourceOption) (*DocumentdbDatabase, error) {
	var resource DocumentdbDatabase
	err := ctx.ReadResource("scaleway:index/documentdbDatabase:DocumentdbDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentdbDatabase resources.
type documentdbDatabaseState struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId *string `pulumi:"instanceId"`
	// Whether the database is managed or not.
	Managed *bool `pulumi:"managed"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// The name of the owner of the database.
	Owner *string `pulumi:"owner"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
	// Size in gigabytes of the database.
	Size *string `pulumi:"size"`
}

type DocumentdbDatabaseState struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId pulumi.StringPtrInput
	// Whether the database is managed or not.
	Managed pulumi.BoolPtrInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// The name of the owner of the database.
	Owner pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
	// Size in gigabytes of the database.
	Size pulumi.StringPtrInput
}

func (DocumentdbDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbDatabaseState)(nil)).Elem()
}

type documentdbDatabaseArgs struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database (e.g. `my-new-database`).
	Name *string `pulumi:"name"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the resource exists.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DocumentdbDatabase resource.
type DocumentdbDatabaseArgs struct {
	// UUID of the documentdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database.
	InstanceId pulumi.StringInput
	// Name of the database (e.g. `my-new-database`).
	Name pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// `region`) The region in which the resource exists.
	Region pulumi.StringPtrInput
}

func (DocumentdbDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbDatabaseArgs)(nil)).Elem()
}

type DocumentdbDatabaseInput interface {
	pulumi.Input

	ToDocumentdbDatabaseOutput() DocumentdbDatabaseOutput
	ToDocumentdbDatabaseOutputWithContext(ctx context.Context) DocumentdbDatabaseOutput
}

func (*DocumentdbDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbDatabase)(nil)).Elem()
}

func (i *DocumentdbDatabase) ToDocumentdbDatabaseOutput() DocumentdbDatabaseOutput {
	return i.ToDocumentdbDatabaseOutputWithContext(context.Background())
}

func (i *DocumentdbDatabase) ToDocumentdbDatabaseOutputWithContext(ctx context.Context) DocumentdbDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbDatabaseOutput)
}

// DocumentdbDatabaseArrayInput is an input type that accepts DocumentdbDatabaseArray and DocumentdbDatabaseArrayOutput values.
// You can construct a concrete instance of `DocumentdbDatabaseArrayInput` via:
//
//	DocumentdbDatabaseArray{ DocumentdbDatabaseArgs{...} }
type DocumentdbDatabaseArrayInput interface {
	pulumi.Input

	ToDocumentdbDatabaseArrayOutput() DocumentdbDatabaseArrayOutput
	ToDocumentdbDatabaseArrayOutputWithContext(context.Context) DocumentdbDatabaseArrayOutput
}

type DocumentdbDatabaseArray []DocumentdbDatabaseInput

func (DocumentdbDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbDatabase)(nil)).Elem()
}

func (i DocumentdbDatabaseArray) ToDocumentdbDatabaseArrayOutput() DocumentdbDatabaseArrayOutput {
	return i.ToDocumentdbDatabaseArrayOutputWithContext(context.Background())
}

func (i DocumentdbDatabaseArray) ToDocumentdbDatabaseArrayOutputWithContext(ctx context.Context) DocumentdbDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbDatabaseArrayOutput)
}

// DocumentdbDatabaseMapInput is an input type that accepts DocumentdbDatabaseMap and DocumentdbDatabaseMapOutput values.
// You can construct a concrete instance of `DocumentdbDatabaseMapInput` via:
//
//	DocumentdbDatabaseMap{ "key": DocumentdbDatabaseArgs{...} }
type DocumentdbDatabaseMapInput interface {
	pulumi.Input

	ToDocumentdbDatabaseMapOutput() DocumentdbDatabaseMapOutput
	ToDocumentdbDatabaseMapOutputWithContext(context.Context) DocumentdbDatabaseMapOutput
}

type DocumentdbDatabaseMap map[string]DocumentdbDatabaseInput

func (DocumentdbDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbDatabase)(nil)).Elem()
}

func (i DocumentdbDatabaseMap) ToDocumentdbDatabaseMapOutput() DocumentdbDatabaseMapOutput {
	return i.ToDocumentdbDatabaseMapOutputWithContext(context.Background())
}

func (i DocumentdbDatabaseMap) ToDocumentdbDatabaseMapOutputWithContext(ctx context.Context) DocumentdbDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbDatabaseMapOutput)
}

type DocumentdbDatabaseOutput struct{ *pulumi.OutputState }

func (DocumentdbDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbDatabase)(nil)).Elem()
}

func (o DocumentdbDatabaseOutput) ToDocumentdbDatabaseOutput() DocumentdbDatabaseOutput {
	return o
}

func (o DocumentdbDatabaseOutput) ToDocumentdbDatabaseOutputWithContext(ctx context.Context) DocumentdbDatabaseOutput {
	return o
}

// UUID of the documentdb instance.
//
// > **Important:** Updates to `instanceId` will recreate the Database.
func (o DocumentdbDatabaseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Whether the database is managed or not.
func (o DocumentdbDatabaseOutput) Managed() pulumi.BoolOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.BoolOutput { return v.Managed }).(pulumi.BoolOutput)
}

// Name of the database (e.g. `my-new-database`).
func (o DocumentdbDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the owner of the database.
func (o DocumentdbDatabaseOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// The projectId you want to attach the resource to
func (o DocumentdbDatabaseOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region in which the resource exists.
func (o DocumentdbDatabaseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Size in gigabytes of the database.
func (o DocumentdbDatabaseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbDatabase) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

type DocumentdbDatabaseArrayOutput struct{ *pulumi.OutputState }

func (DocumentdbDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbDatabase)(nil)).Elem()
}

func (o DocumentdbDatabaseArrayOutput) ToDocumentdbDatabaseArrayOutput() DocumentdbDatabaseArrayOutput {
	return o
}

func (o DocumentdbDatabaseArrayOutput) ToDocumentdbDatabaseArrayOutputWithContext(ctx context.Context) DocumentdbDatabaseArrayOutput {
	return o
}

func (o DocumentdbDatabaseArrayOutput) Index(i pulumi.IntInput) DocumentdbDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentdbDatabase {
		return vs[0].([]*DocumentdbDatabase)[vs[1].(int)]
	}).(DocumentdbDatabaseOutput)
}

type DocumentdbDatabaseMapOutput struct{ *pulumi.OutputState }

func (DocumentdbDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbDatabase)(nil)).Elem()
}

func (o DocumentdbDatabaseMapOutput) ToDocumentdbDatabaseMapOutput() DocumentdbDatabaseMapOutput {
	return o
}

func (o DocumentdbDatabaseMapOutput) ToDocumentdbDatabaseMapOutputWithContext(ctx context.Context) DocumentdbDatabaseMapOutput {
	return o
}

func (o DocumentdbDatabaseMapOutput) MapIndex(k pulumi.StringInput) DocumentdbDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentdbDatabase {
		return vs[0].(map[string]*DocumentdbDatabase)[vs[1].(string)]
	}).(DocumentdbDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbDatabaseInput)(nil)).Elem(), &DocumentdbDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbDatabaseArrayInput)(nil)).Elem(), DocumentdbDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbDatabaseMapInput)(nil)).Elem(), DocumentdbDatabaseMap{})
	pulumi.RegisterOutputType(DocumentdbDatabaseOutput{})
	pulumi.RegisterOutputType(DocumentdbDatabaseArrayOutput{})
	pulumi.RegisterOutputType(DocumentdbDatabaseMapOutput{})
}
