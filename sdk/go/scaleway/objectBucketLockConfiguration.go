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

// The `ObjectBucketLockConfiguration` resource allows you to create and manage an object lock configuration for [Scaleway Object storage](https://www.scaleway.com/en/docs/storage/object/).
//
// Refer to the [dedicated documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/object-lock/) for more information on object lock.
//
// ## Example Usage
//
// ### Configure an Object Lock for a new bucket
//
// > **Note:** `objectLockEnabled` must be set to `true` before configuring the lock.
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
//			mainObjectBucket, err := scaleway.NewObjectBucket(ctx, "mainObjectBucket", &scaleway.ObjectBucketArgs{
//				Acl:               pulumi.String("public-read"),
//				ObjectLockEnabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewObjectBucketLockConfiguration(ctx, "mainObjectBucketLockConfiguration", &scaleway.ObjectBucketLockConfigurationArgs{
//				Bucket: mainObjectBucket.Name,
//				Rule: &scaleway.ObjectBucketLockConfigurationRuleArgs{
//					DefaultRetention: &scaleway.ObjectBucketLockConfigurationRuleDefaultRetentionArgs{
//						Mode: pulumi.String("GOVERNANCE"),
//						Days: pulumi.Int(1),
//					},
//				},
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
// ### Configure an object Lock for an existing bucket
//
// [Contact Scaleway support](https://console.scaleway.com/support/tickets/create) to enable object lock on an existing bucket.
//
// ## Import
//
// Bucket lock configurations can be imported using the `{region}/{bucketName}` identifier, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration some_bucket fr-par/some-bucket
// ```
//
// ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
//
// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
// ```
type ObjectBucketLockConfiguration struct {
	pulumi.CustomResourceState

	// The bucket's name or regional ID.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the object lock rule for the specified object.
	Rule ObjectBucketLockConfigurationRuleOutput `pulumi:"rule"`
}

// NewObjectBucketLockConfiguration registers a new resource with the given unique name, arguments, and options.
func NewObjectBucketLockConfiguration(ctx *pulumi.Context,
	name string, args *ObjectBucketLockConfigurationArgs, opts ...pulumi.ResourceOption) (*ObjectBucketLockConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Rule == nil {
		return nil, errors.New("invalid value for required argument 'Rule'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectBucketLockConfiguration
	err := ctx.RegisterResource("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectBucketLockConfiguration gets an existing ObjectBucketLockConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectBucketLockConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectBucketLockConfigurationState, opts ...pulumi.ResourceOption) (*ObjectBucketLockConfiguration, error) {
	var resource ObjectBucketLockConfiguration
	err := ctx.ReadResource("scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectBucketLockConfiguration resources.
type objectBucketLockConfigurationState struct {
	// The bucket's name or regional ID.
	Bucket *string `pulumi:"bucket"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Specifies the object lock rule for the specified object.
	Rule *ObjectBucketLockConfigurationRule `pulumi:"rule"`
}

type ObjectBucketLockConfigurationState struct {
	// The bucket's name or regional ID.
	Bucket pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Specifies the object lock rule for the specified object.
	Rule ObjectBucketLockConfigurationRulePtrInput
}

func (ObjectBucketLockConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketLockConfigurationState)(nil)).Elem()
}

type objectBucketLockConfigurationArgs struct {
	// The bucket's name or regional ID.
	Bucket string `pulumi:"bucket"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The region you want to attach the resource to
	Region *string `pulumi:"region"`
	// Specifies the object lock rule for the specified object.
	Rule ObjectBucketLockConfigurationRule `pulumi:"rule"`
}

// The set of arguments for constructing a ObjectBucketLockConfiguration resource.
type ObjectBucketLockConfigurationArgs struct {
	// The bucket's name or regional ID.
	Bucket pulumi.StringInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The region you want to attach the resource to
	Region pulumi.StringPtrInput
	// Specifies the object lock rule for the specified object.
	Rule ObjectBucketLockConfigurationRuleInput
}

func (ObjectBucketLockConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketLockConfigurationArgs)(nil)).Elem()
}

type ObjectBucketLockConfigurationInput interface {
	pulumi.Input

	ToObjectBucketLockConfigurationOutput() ObjectBucketLockConfigurationOutput
	ToObjectBucketLockConfigurationOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationOutput
}

func (*ObjectBucketLockConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucketLockConfiguration)(nil)).Elem()
}

func (i *ObjectBucketLockConfiguration) ToObjectBucketLockConfigurationOutput() ObjectBucketLockConfigurationOutput {
	return i.ToObjectBucketLockConfigurationOutputWithContext(context.Background())
}

func (i *ObjectBucketLockConfiguration) ToObjectBucketLockConfigurationOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketLockConfigurationOutput)
}

// ObjectBucketLockConfigurationArrayInput is an input type that accepts ObjectBucketLockConfigurationArray and ObjectBucketLockConfigurationArrayOutput values.
// You can construct a concrete instance of `ObjectBucketLockConfigurationArrayInput` via:
//
//	ObjectBucketLockConfigurationArray{ ObjectBucketLockConfigurationArgs{...} }
type ObjectBucketLockConfigurationArrayInput interface {
	pulumi.Input

	ToObjectBucketLockConfigurationArrayOutput() ObjectBucketLockConfigurationArrayOutput
	ToObjectBucketLockConfigurationArrayOutputWithContext(context.Context) ObjectBucketLockConfigurationArrayOutput
}

type ObjectBucketLockConfigurationArray []ObjectBucketLockConfigurationInput

func (ObjectBucketLockConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucketLockConfiguration)(nil)).Elem()
}

func (i ObjectBucketLockConfigurationArray) ToObjectBucketLockConfigurationArrayOutput() ObjectBucketLockConfigurationArrayOutput {
	return i.ToObjectBucketLockConfigurationArrayOutputWithContext(context.Background())
}

func (i ObjectBucketLockConfigurationArray) ToObjectBucketLockConfigurationArrayOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketLockConfigurationArrayOutput)
}

// ObjectBucketLockConfigurationMapInput is an input type that accepts ObjectBucketLockConfigurationMap and ObjectBucketLockConfigurationMapOutput values.
// You can construct a concrete instance of `ObjectBucketLockConfigurationMapInput` via:
//
//	ObjectBucketLockConfigurationMap{ "key": ObjectBucketLockConfigurationArgs{...} }
type ObjectBucketLockConfigurationMapInput interface {
	pulumi.Input

	ToObjectBucketLockConfigurationMapOutput() ObjectBucketLockConfigurationMapOutput
	ToObjectBucketLockConfigurationMapOutputWithContext(context.Context) ObjectBucketLockConfigurationMapOutput
}

type ObjectBucketLockConfigurationMap map[string]ObjectBucketLockConfigurationInput

func (ObjectBucketLockConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucketLockConfiguration)(nil)).Elem()
}

func (i ObjectBucketLockConfigurationMap) ToObjectBucketLockConfigurationMapOutput() ObjectBucketLockConfigurationMapOutput {
	return i.ToObjectBucketLockConfigurationMapOutputWithContext(context.Background())
}

func (i ObjectBucketLockConfigurationMap) ToObjectBucketLockConfigurationMapOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketLockConfigurationMapOutput)
}

type ObjectBucketLockConfigurationOutput struct{ *pulumi.OutputState }

func (ObjectBucketLockConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucketLockConfiguration)(nil)).Elem()
}

func (o ObjectBucketLockConfigurationOutput) ToObjectBucketLockConfigurationOutput() ObjectBucketLockConfigurationOutput {
	return o
}

func (o ObjectBucketLockConfigurationOutput) ToObjectBucketLockConfigurationOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationOutput {
	return o
}

// The bucket's name or regional ID.
func (o ObjectBucketLockConfigurationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketLockConfiguration) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The projectId you want to attach the resource to
func (o ObjectBucketLockConfigurationOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketLockConfiguration) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The region you want to attach the resource to
func (o ObjectBucketLockConfigurationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketLockConfiguration) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the object lock rule for the specified object.
func (o ObjectBucketLockConfigurationOutput) Rule() ObjectBucketLockConfigurationRuleOutput {
	return o.ApplyT(func(v *ObjectBucketLockConfiguration) ObjectBucketLockConfigurationRuleOutput { return v.Rule }).(ObjectBucketLockConfigurationRuleOutput)
}

type ObjectBucketLockConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ObjectBucketLockConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucketLockConfiguration)(nil)).Elem()
}

func (o ObjectBucketLockConfigurationArrayOutput) ToObjectBucketLockConfigurationArrayOutput() ObjectBucketLockConfigurationArrayOutput {
	return o
}

func (o ObjectBucketLockConfigurationArrayOutput) ToObjectBucketLockConfigurationArrayOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationArrayOutput {
	return o
}

func (o ObjectBucketLockConfigurationArrayOutput) Index(i pulumi.IntInput) ObjectBucketLockConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectBucketLockConfiguration {
		return vs[0].([]*ObjectBucketLockConfiguration)[vs[1].(int)]
	}).(ObjectBucketLockConfigurationOutput)
}

type ObjectBucketLockConfigurationMapOutput struct{ *pulumi.OutputState }

func (ObjectBucketLockConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucketLockConfiguration)(nil)).Elem()
}

func (o ObjectBucketLockConfigurationMapOutput) ToObjectBucketLockConfigurationMapOutput() ObjectBucketLockConfigurationMapOutput {
	return o
}

func (o ObjectBucketLockConfigurationMapOutput) ToObjectBucketLockConfigurationMapOutputWithContext(ctx context.Context) ObjectBucketLockConfigurationMapOutput {
	return o
}

func (o ObjectBucketLockConfigurationMapOutput) MapIndex(k pulumi.StringInput) ObjectBucketLockConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectBucketLockConfiguration {
		return vs[0].(map[string]*ObjectBucketLockConfiguration)[vs[1].(string)]
	}).(ObjectBucketLockConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketLockConfigurationInput)(nil)).Elem(), &ObjectBucketLockConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketLockConfigurationArrayInput)(nil)).Elem(), ObjectBucketLockConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketLockConfigurationMapInput)(nil)).Elem(), ObjectBucketLockConfigurationMap{})
	pulumi.RegisterOutputType(ObjectBucketLockConfigurationOutput{})
	pulumi.RegisterOutputType(ObjectBucketLockConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ObjectBucketLockConfigurationMapOutput{})
}
