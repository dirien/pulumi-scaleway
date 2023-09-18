// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway object storage buckets.
// For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).
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
//			_, err := scaleway.NewObjectBucket(ctx, "someBucket", &scaleway.ObjectBucketArgs{
//				Tags: pulumi.StringMap{
//					"key": pulumi.String("value"),
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
// ### Creating the bucket in a specific project
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
//			_, err := scaleway.NewObjectBucket(ctx, "someBucket", &scaleway.ObjectBucketArgs{
//				ProjectId: pulumi.String("11111111-1111-1111-1111-111111111111"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Using object lifecycle
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
//			_, err := scaleway.NewObjectBucket(ctx, "main", &scaleway.ObjectBucketArgs{
//				LifecycleRules: scaleway.ObjectBucketLifecycleRuleArray{
//					&scaleway.ObjectBucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Expiration: &scaleway.ObjectBucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(365),
//						},
//						Id:     pulumi.String("id1"),
//						Prefix: pulumi.String("path1/"),
//						Transitions: scaleway.ObjectBucketLifecycleRuleTransitionArray{
//							&scaleway.ObjectBucketLifecycleRuleTransitionArgs{
//								Days:         pulumi.Int(120),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&scaleway.ObjectBucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Expiration: &scaleway.ObjectBucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(50),
//						},
//						Id:     pulumi.String("id2"),
//						Prefix: pulumi.String("path2/"),
//					},
//					&scaleway.ObjectBucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(false),
//						Expiration: &scaleway.ObjectBucketLifecycleRuleExpirationArgs{
//							Days: pulumi.Int(1),
//						},
//						Id:     pulumi.String("id3"),
//						Prefix: pulumi.String("path3/"),
//						Tags: pulumi.StringMap{
//							"tagKey":    pulumi.String("tagValue"),
//							"terraform": pulumi.String("hashicorp"),
//						},
//					},
//					&scaleway.ObjectBucketLifecycleRuleArgs{
//						Enabled: pulumi.Bool(true),
//						Id:      pulumi.String("id4"),
//						Tags: pulumi.StringMap{
//							"tag1": pulumi.String("value1"),
//						},
//						Transitions: scaleway.ObjectBucketLifecycleRuleTransitionArray{
//							&scaleway.ObjectBucketLifecycleRuleTransitionArgs{
//								Days:         pulumi.Int(0),
//								StorageClass: pulumi.String("GLACIER"),
//							},
//						},
//					},
//					&scaleway.ObjectBucketLifecycleRuleArgs{
//						AbortIncompleteMultipartUploadDays: pulumi.Int(30),
//						Enabled:                            pulumi.Bool(true),
//					},
//				},
//				Region: pulumi.String("fr-par"),
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
// Buckets can be imported using the `{region}/{bucketName}` identifier, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket
//
// ```
//
//	If you are importing a bucket from a specific project (that is not your default project), you can use the following syntaxbash
//
// ```sh
//
//	$ pulumi import scaleway:index/objectBucket:ObjectBucket some_bucket fr-par/some-bucket@11111111-1111-1111-1111-111111111111
//
// ```
type ObjectBucket struct {
	pulumi.CustomResourceState

	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules ObjectBucketCorsRuleArrayOutput `pulumi:"corsRules"`
	// The endpoint URL of the bucket
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules ObjectBucketLifecycleRuleArrayOutput `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrOutput `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See ObjectBucketAcl resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// A list of tags (key / value) for the bucket.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning ObjectBucketVersioningOutput `pulumi:"versioning"`
}

// NewObjectBucket registers a new resource with the given unique name, arguments, and options.
func NewObjectBucket(ctx *pulumi.Context,
	name string, args *ObjectBucketArgs, opts ...pulumi.ResourceOption) (*ObjectBucket, error) {
	if args == nil {
		args = &ObjectBucketArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectBucket
	err := ctx.RegisterResource("scaleway:index/objectBucket:ObjectBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectBucket gets an existing ObjectBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectBucketState, opts ...pulumi.ResourceOption) (*ObjectBucket, error) {
	var resource ObjectBucket
	err := ctx.ReadResource("scaleway:index/objectBucket:ObjectBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectBucket resources.
type objectBucketState struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl *string `pulumi:"acl"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []ObjectBucketCorsRule `pulumi:"corsRules"`
	// The endpoint URL of the bucket
	Endpoint *string `pulumi:"endpoint"`
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules []ObjectBucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See ObjectBucketAcl resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region *string `pulumi:"region"`
	// A list of tags (key / value) for the bucket.
	Tags map[string]string `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *ObjectBucketVersioning `pulumi:"versioning"`
}

type ObjectBucketState struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules ObjectBucketCorsRuleArrayInput
	// The endpoint URL of the bucket
	Endpoint pulumi.StringPtrInput
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy pulumi.BoolPtrInput
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules ObjectBucketLifecycleRuleArrayInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrInput
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See ObjectBucketAcl resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId pulumi.StringPtrInput
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region pulumi.StringPtrInput
	// A list of tags (key / value) for the bucket.
	Tags pulumi.StringMapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning ObjectBucketVersioningPtrInput
}

func (ObjectBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketState)(nil)).Elem()
}

type objectBucketArgs struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl *string `pulumi:"acl"`
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules []ObjectBucketCorsRule `pulumi:"corsRules"`
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules []ObjectBucketLifecycleRule `pulumi:"lifecycleRules"`
	// The name of the bucket.
	Name *string `pulumi:"name"`
	// Enable object lock
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See ObjectBucketAcl resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region *string `pulumi:"region"`
	// A list of tags (key / value) for the bucket.
	Tags map[string]string `pulumi:"tags"`
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning *ObjectBucketVersioning `pulumi:"versioning"`
}

// The set of arguments for constructing a ObjectBucket resource.
type ObjectBucketArgs struct {
	// (Deprecated) The canned ACL you want to apply to the bucket.
	//
	// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
	Acl pulumi.StringPtrInput
	// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
	CorsRules ObjectBucketCorsRuleArrayInput
	// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
	ForceDestroy pulumi.BoolPtrInput
	// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
	LifecycleRules ObjectBucketLifecycleRuleArrayInput
	// The name of the bucket.
	Name pulumi.StringPtrInput
	// Enable object lock
	ObjectLockEnabled pulumi.BoolPtrInput
	// `projectId`) The ID of the project the bucket is associated with.
	//
	// The `acl` attribute is deprecated. See ObjectBucketAcl resource documentation.
	// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
	ProjectId pulumi.StringPtrInput
	// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
	Region pulumi.StringPtrInput
	// A list of tags (key / value) for the bucket.
	Tags pulumi.StringMapInput
	// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
	Versioning ObjectBucketVersioningPtrInput
}

func (ObjectBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketArgs)(nil)).Elem()
}

type ObjectBucketInput interface {
	pulumi.Input

	ToObjectBucketOutput() ObjectBucketOutput
	ToObjectBucketOutputWithContext(ctx context.Context) ObjectBucketOutput
}

func (*ObjectBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucket)(nil)).Elem()
}

func (i *ObjectBucket) ToObjectBucketOutput() ObjectBucketOutput {
	return i.ToObjectBucketOutputWithContext(context.Background())
}

func (i *ObjectBucket) ToObjectBucketOutputWithContext(ctx context.Context) ObjectBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketOutput)
}

func (i *ObjectBucket) ToOutput(ctx context.Context) pulumix.Output[*ObjectBucket] {
	return pulumix.Output[*ObjectBucket]{
		OutputState: i.ToObjectBucketOutputWithContext(ctx).OutputState,
	}
}

// ObjectBucketArrayInput is an input type that accepts ObjectBucketArray and ObjectBucketArrayOutput values.
// You can construct a concrete instance of `ObjectBucketArrayInput` via:
//
//	ObjectBucketArray{ ObjectBucketArgs{...} }
type ObjectBucketArrayInput interface {
	pulumi.Input

	ToObjectBucketArrayOutput() ObjectBucketArrayOutput
	ToObjectBucketArrayOutputWithContext(context.Context) ObjectBucketArrayOutput
}

type ObjectBucketArray []ObjectBucketInput

func (ObjectBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucket)(nil)).Elem()
}

func (i ObjectBucketArray) ToObjectBucketArrayOutput() ObjectBucketArrayOutput {
	return i.ToObjectBucketArrayOutputWithContext(context.Background())
}

func (i ObjectBucketArray) ToObjectBucketArrayOutputWithContext(ctx context.Context) ObjectBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketArrayOutput)
}

func (i ObjectBucketArray) ToOutput(ctx context.Context) pulumix.Output[[]*ObjectBucket] {
	return pulumix.Output[[]*ObjectBucket]{
		OutputState: i.ToObjectBucketArrayOutputWithContext(ctx).OutputState,
	}
}

// ObjectBucketMapInput is an input type that accepts ObjectBucketMap and ObjectBucketMapOutput values.
// You can construct a concrete instance of `ObjectBucketMapInput` via:
//
//	ObjectBucketMap{ "key": ObjectBucketArgs{...} }
type ObjectBucketMapInput interface {
	pulumi.Input

	ToObjectBucketMapOutput() ObjectBucketMapOutput
	ToObjectBucketMapOutputWithContext(context.Context) ObjectBucketMapOutput
}

type ObjectBucketMap map[string]ObjectBucketInput

func (ObjectBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucket)(nil)).Elem()
}

func (i ObjectBucketMap) ToObjectBucketMapOutput() ObjectBucketMapOutput {
	return i.ToObjectBucketMapOutputWithContext(context.Background())
}

func (i ObjectBucketMap) ToObjectBucketMapOutputWithContext(ctx context.Context) ObjectBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketMapOutput)
}

func (i ObjectBucketMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ObjectBucket] {
	return pulumix.Output[map[string]*ObjectBucket]{
		OutputState: i.ToObjectBucketMapOutputWithContext(ctx).OutputState,
	}
}

type ObjectBucketOutput struct{ *pulumi.OutputState }

func (ObjectBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucket)(nil)).Elem()
}

func (o ObjectBucketOutput) ToObjectBucketOutput() ObjectBucketOutput {
	return o
}

func (o ObjectBucketOutput) ToObjectBucketOutputWithContext(ctx context.Context) ObjectBucketOutput {
	return o
}

func (o ObjectBucketOutput) ToOutput(ctx context.Context) pulumix.Output[*ObjectBucket] {
	return pulumix.Output[*ObjectBucket]{
		OutputState: o.OutputState,
	}
}

// (Deprecated) The canned ACL you want to apply to the bucket.
//
// Deprecated: ACL attribute is deprecated. Please use the resource scaleway_object_bucket_acl instead.
func (o ObjectBucketOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
func (o ObjectBucketOutput) CorsRules() ObjectBucketCorsRuleArrayOutput {
	return o.ApplyT(func(v *ObjectBucket) ObjectBucketCorsRuleArrayOutput { return v.CorsRules }).(ObjectBucketCorsRuleArrayOutput)
}

// The endpoint URL of the bucket
func (o ObjectBucketOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Enable deletion of objects in bucket before destroying, locked objects or under legal hold are also deleted and **not** recoverable
func (o ObjectBucketOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// Lifecycle configuration is a set of rules that define actions that Scaleway Object Storage applies to a group of objects
func (o ObjectBucketOutput) LifecycleRules() ObjectBucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v *ObjectBucket) ObjectBucketLifecycleRuleArrayOutput { return v.LifecycleRules }).(ObjectBucketLifecycleRuleArrayOutput)
}

// The name of the bucket.
func (o ObjectBucketOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Enable object lock
func (o ObjectBucketOutput) ObjectLockEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.BoolPtrOutput { return v.ObjectLockEnabled }).(pulumi.BoolPtrOutput)
}

// `projectId`) The ID of the project the bucket is associated with.
//
// The `acl` attribute is deprecated. See ObjectBucketAcl resource documentation.
// Please check the [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation for supported values.
func (o ObjectBucketOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The [region](https://developers.scaleway.com/en/quickstart/#region-definition) in which the bucket should be created.
func (o ObjectBucketOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// A list of tags (key / value) for the bucket.
func (o ObjectBucketOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ObjectBucket) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
func (o ObjectBucketOutput) Versioning() ObjectBucketVersioningOutput {
	return o.ApplyT(func(v *ObjectBucket) ObjectBucketVersioningOutput { return v.Versioning }).(ObjectBucketVersioningOutput)
}

type ObjectBucketArrayOutput struct{ *pulumi.OutputState }

func (ObjectBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucket)(nil)).Elem()
}

func (o ObjectBucketArrayOutput) ToObjectBucketArrayOutput() ObjectBucketArrayOutput {
	return o
}

func (o ObjectBucketArrayOutput) ToObjectBucketArrayOutputWithContext(ctx context.Context) ObjectBucketArrayOutput {
	return o
}

func (o ObjectBucketArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ObjectBucket] {
	return pulumix.Output[[]*ObjectBucket]{
		OutputState: o.OutputState,
	}
}

func (o ObjectBucketArrayOutput) Index(i pulumi.IntInput) ObjectBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectBucket {
		return vs[0].([]*ObjectBucket)[vs[1].(int)]
	}).(ObjectBucketOutput)
}

type ObjectBucketMapOutput struct{ *pulumi.OutputState }

func (ObjectBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucket)(nil)).Elem()
}

func (o ObjectBucketMapOutput) ToObjectBucketMapOutput() ObjectBucketMapOutput {
	return o
}

func (o ObjectBucketMapOutput) ToObjectBucketMapOutputWithContext(ctx context.Context) ObjectBucketMapOutput {
	return o
}

func (o ObjectBucketMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ObjectBucket] {
	return pulumix.Output[map[string]*ObjectBucket]{
		OutputState: o.OutputState,
	}
}

func (o ObjectBucketMapOutput) MapIndex(k pulumi.StringInput) ObjectBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectBucket {
		return vs[0].(map[string]*ObjectBucket)[vs[1].(string)]
	}).(ObjectBucketOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketInput)(nil)).Elem(), &ObjectBucket{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketArrayInput)(nil)).Elem(), ObjectBucketArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketMapInput)(nil)).Elem(), ObjectBucketMap{})
	pulumi.RegisterOutputType(ObjectBucketOutput{})
	pulumi.RegisterOutputType(ObjectBucketArrayOutput{})
	pulumi.RegisterOutputType(ObjectBucketMapOutput{})
}
