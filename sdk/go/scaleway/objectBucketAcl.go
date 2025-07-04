// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
//			_, err := scaleway.NewObjectBucket(ctx, "someBucket", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewObjectBucketAcl(ctx, "main", &scaleway.ObjectBucketAclArgs{
//				Bucket: pulumi.Any(scaleway_object_bucket.Main.Id),
//				Acl:    pulumi.String("private"),
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
// For more information, refer to the [PutBucketAcl API call documentation](https://www.scaleway.com/en/docs/object-storage/api-cli/bucket-operations/#putbucketacl).
//
// ### With Grants
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
//			mainObjectBucket, err := scaleway.NewObjectBucket(ctx, "mainObjectBucket", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewObjectBucketAcl(ctx, "mainObjectBucketAcl", &scaleway.ObjectBucketAclArgs{
//				Bucket: mainObjectBucket.ID(),
//				AccessControlPolicy: &scaleway.ObjectBucketAclAccessControlPolicyArgs{
//					Grants: scaleway.ObjectBucketAclAccessControlPolicyGrantArray{
//						&scaleway.ObjectBucketAclAccessControlPolicyGrantArgs{
//							Grantee: &scaleway.ObjectBucketAclAccessControlPolicyGrantGranteeArgs{
//								Id:   pulumi.String("<project-id>:<project-id>"),
//								Type: pulumi.String("CanonicalUser"),
//							},
//							Permission: pulumi.String("FULL_CONTROL"),
//						},
//						&scaleway.ObjectBucketAclAccessControlPolicyGrantArgs{
//							Grantee: &scaleway.ObjectBucketAclAccessControlPolicyGrantGranteeArgs{
//								Id:   pulumi.String("<project-id>"),
//								Type: pulumi.String("CanonicalUser"),
//							},
//							Permission: pulumi.String("WRITE"),
//						},
//					},
//					Owner: &scaleway.ObjectBucketAclAccessControlPolicyOwnerArgs{
//						Id: pulumi.String("<project-id>"),
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
// ## The ACL
//
// Refer to the [official canned ACL documentation](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) for more information on the different roles.
//
// ## The access control policy
//
// The `accessControlPolicy` configuration block supports the following arguments:
//
// * `grant` - (Required) Set of grant configuration blocks documented below.
// * `owner` - (Required) Configuration block of the bucket owner's display name and ID documented below.
//
// ## The grant
//
// The `grant` configuration block supports the following arguments:
//
// * `grantee` - (Required) Configuration block for the project being granted permissions documented below.
// * `permission` - (Required) Logging permissions assigned to the grantee for the bucket.
//
// ## The permission
//
// The following list shows each access policy permissions supported.
//
// `READ`, `WRITE`, `READ_ACP`, `WRITE_ACP`, `FULL_CONTROL`
//
// For more information about ACL permissions in the S3 bucket, see [ACL permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html).
//
// ## The owner
//
// The `owner` configuration block supports the following arguments:
//
// * `id` - (Required) The ID of the project owner.
// * `displayName` - (Optional) The display name of the owner.
//
// ## the grantee
//
// The `grantee` configuration block supports the following arguments:
//
// * `id` - (Required) The canonical user ID of the grantee.
// * `type` - (Required) Type of grantee. Valid values: CanonicalUser.
//
// ## Import
//
// Bucket ACLs can be imported using the `{region}/{bucketName}/{acl}` identifier, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/objectBucketAcl:ObjectBucketAcl some_bucket fr-par/some-bucket/private
// ```
//
// ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.
//
// If you are using a project different from the default one, you have to specify the project ID at the end of the import command.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/objectBucketAcl:ObjectBucketAcl some_bucket fr-par/some-bucket/private@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
// ```
type ObjectBucketAcl struct {
	pulumi.CustomResourceState

	// A configuration block that sets the ACL permissions for an object per grantee documented below.
	AccessControlPolicy ObjectBucketAclAccessControlPolicyOutput `pulumi:"accessControlPolicy"`
	// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
	Acl pulumi.StringPtrOutput `pulumi:"acl"`
	// The bucket's name or regional ID.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The project ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrOutput `pulumi:"expectedBucketOwner"`
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewObjectBucketAcl registers a new resource with the given unique name, arguments, and options.
func NewObjectBucketAcl(ctx *pulumi.Context,
	name string, args *ObjectBucketAclArgs, opts ...pulumi.ResourceOption) (*ObjectBucketAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ObjectBucketAcl
	err := ctx.RegisterResource("scaleway:index/objectBucketAcl:ObjectBucketAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectBucketAcl gets an existing ObjectBucketAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectBucketAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectBucketAclState, opts ...pulumi.ResourceOption) (*ObjectBucketAcl, error) {
	var resource ObjectBucketAcl
	err := ctx.ReadResource("scaleway:index/objectBucketAcl:ObjectBucketAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectBucketAcl resources.
type objectBucketAclState struct {
	// A configuration block that sets the ACL permissions for an object per grantee documented below.
	AccessControlPolicy *ObjectBucketAclAccessControlPolicy `pulumi:"accessControlPolicy"`
	// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
	Acl *string `pulumi:"acl"`
	// The bucket's name or regional ID.
	Bucket *string `pulumi:"bucket"`
	// The project ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
	Region *string `pulumi:"region"`
}

type ObjectBucketAclState struct {
	// A configuration block that sets the ACL permissions for an object per grantee documented below.
	AccessControlPolicy ObjectBucketAclAccessControlPolicyPtrInput
	// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
	Acl pulumi.StringPtrInput
	// The bucket's name or regional ID.
	Bucket pulumi.StringPtrInput
	// The project ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
	Region pulumi.StringPtrInput
}

func (ObjectBucketAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketAclState)(nil)).Elem()
}

type objectBucketAclArgs struct {
	// A configuration block that sets the ACL permissions for an object per grantee documented below.
	AccessControlPolicy *ObjectBucketAclAccessControlPolicy `pulumi:"accessControlPolicy"`
	// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
	Acl *string `pulumi:"acl"`
	// The bucket's name or regional ID.
	Bucket string `pulumi:"bucket"`
	// The project ID of the expected bucket owner.
	ExpectedBucketOwner *string `pulumi:"expectedBucketOwner"`
	// The projectId you want to attach the resource to
	ProjectId *string `pulumi:"projectId"`
	// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a ObjectBucketAcl resource.
type ObjectBucketAclArgs struct {
	// A configuration block that sets the ACL permissions for an object per grantee documented below.
	AccessControlPolicy ObjectBucketAclAccessControlPolicyPtrInput
	// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
	Acl pulumi.StringPtrInput
	// The bucket's name or regional ID.
	Bucket pulumi.StringInput
	// The project ID of the expected bucket owner.
	ExpectedBucketOwner pulumi.StringPtrInput
	// The projectId you want to attach the resource to
	ProjectId pulumi.StringPtrInput
	// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
	Region pulumi.StringPtrInput
}

func (ObjectBucketAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectBucketAclArgs)(nil)).Elem()
}

type ObjectBucketAclInput interface {
	pulumi.Input

	ToObjectBucketAclOutput() ObjectBucketAclOutput
	ToObjectBucketAclOutputWithContext(ctx context.Context) ObjectBucketAclOutput
}

func (*ObjectBucketAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucketAcl)(nil)).Elem()
}

func (i *ObjectBucketAcl) ToObjectBucketAclOutput() ObjectBucketAclOutput {
	return i.ToObjectBucketAclOutputWithContext(context.Background())
}

func (i *ObjectBucketAcl) ToObjectBucketAclOutputWithContext(ctx context.Context) ObjectBucketAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketAclOutput)
}

// ObjectBucketAclArrayInput is an input type that accepts ObjectBucketAclArray and ObjectBucketAclArrayOutput values.
// You can construct a concrete instance of `ObjectBucketAclArrayInput` via:
//
//	ObjectBucketAclArray{ ObjectBucketAclArgs{...} }
type ObjectBucketAclArrayInput interface {
	pulumi.Input

	ToObjectBucketAclArrayOutput() ObjectBucketAclArrayOutput
	ToObjectBucketAclArrayOutputWithContext(context.Context) ObjectBucketAclArrayOutput
}

type ObjectBucketAclArray []ObjectBucketAclInput

func (ObjectBucketAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucketAcl)(nil)).Elem()
}

func (i ObjectBucketAclArray) ToObjectBucketAclArrayOutput() ObjectBucketAclArrayOutput {
	return i.ToObjectBucketAclArrayOutputWithContext(context.Background())
}

func (i ObjectBucketAclArray) ToObjectBucketAclArrayOutputWithContext(ctx context.Context) ObjectBucketAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketAclArrayOutput)
}

// ObjectBucketAclMapInput is an input type that accepts ObjectBucketAclMap and ObjectBucketAclMapOutput values.
// You can construct a concrete instance of `ObjectBucketAclMapInput` via:
//
//	ObjectBucketAclMap{ "key": ObjectBucketAclArgs{...} }
type ObjectBucketAclMapInput interface {
	pulumi.Input

	ToObjectBucketAclMapOutput() ObjectBucketAclMapOutput
	ToObjectBucketAclMapOutputWithContext(context.Context) ObjectBucketAclMapOutput
}

type ObjectBucketAclMap map[string]ObjectBucketAclInput

func (ObjectBucketAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucketAcl)(nil)).Elem()
}

func (i ObjectBucketAclMap) ToObjectBucketAclMapOutput() ObjectBucketAclMapOutput {
	return i.ToObjectBucketAclMapOutputWithContext(context.Background())
}

func (i ObjectBucketAclMap) ToObjectBucketAclMapOutputWithContext(ctx context.Context) ObjectBucketAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectBucketAclMapOutput)
}

type ObjectBucketAclOutput struct{ *pulumi.OutputState }

func (ObjectBucketAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectBucketAcl)(nil)).Elem()
}

func (o ObjectBucketAclOutput) ToObjectBucketAclOutput() ObjectBucketAclOutput {
	return o
}

func (o ObjectBucketAclOutput) ToObjectBucketAclOutputWithContext(ctx context.Context) ObjectBucketAclOutput {
	return o
}

// A configuration block that sets the ACL permissions for an object per grantee documented below.
func (o ObjectBucketAclOutput) AccessControlPolicy() ObjectBucketAclAccessControlPolicyOutput {
	return o.ApplyT(func(v *ObjectBucketAcl) ObjectBucketAclAccessControlPolicyOutput { return v.AccessControlPolicy }).(ObjectBucketAclAccessControlPolicyOutput)
}

// The canned ACL you want to apply to the bucket. Refer to the [AWS Canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl_overview.html#canned-acl) documentation page to find a list of all the supported canned ACLs.
func (o ObjectBucketAclOutput) Acl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectBucketAcl) pulumi.StringPtrOutput { return v.Acl }).(pulumi.StringPtrOutput)
}

// The bucket's name or regional ID.
func (o ObjectBucketAclOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketAcl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The project ID of the expected bucket owner.
func (o ObjectBucketAclOutput) ExpectedBucketOwner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectBucketAcl) pulumi.StringPtrOutput { return v.ExpectedBucketOwner }).(pulumi.StringPtrOutput)
}

// The projectId you want to attach the resource to
func (o ObjectBucketAclOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketAcl) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The [region](https://www.scaleway.com/en/developers/api/#regions-and-zones) in which the bucket should be created.
func (o ObjectBucketAclOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectBucketAcl) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ObjectBucketAclArrayOutput struct{ *pulumi.OutputState }

func (ObjectBucketAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectBucketAcl)(nil)).Elem()
}

func (o ObjectBucketAclArrayOutput) ToObjectBucketAclArrayOutput() ObjectBucketAclArrayOutput {
	return o
}

func (o ObjectBucketAclArrayOutput) ToObjectBucketAclArrayOutputWithContext(ctx context.Context) ObjectBucketAclArrayOutput {
	return o
}

func (o ObjectBucketAclArrayOutput) Index(i pulumi.IntInput) ObjectBucketAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ObjectBucketAcl {
		return vs[0].([]*ObjectBucketAcl)[vs[1].(int)]
	}).(ObjectBucketAclOutput)
}

type ObjectBucketAclMapOutput struct{ *pulumi.OutputState }

func (ObjectBucketAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectBucketAcl)(nil)).Elem()
}

func (o ObjectBucketAclMapOutput) ToObjectBucketAclMapOutput() ObjectBucketAclMapOutput {
	return o
}

func (o ObjectBucketAclMapOutput) ToObjectBucketAclMapOutputWithContext(ctx context.Context) ObjectBucketAclMapOutput {
	return o
}

func (o ObjectBucketAclMapOutput) MapIndex(k pulumi.StringInput) ObjectBucketAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ObjectBucketAcl {
		return vs[0].(map[string]*ObjectBucketAcl)[vs[1].(string)]
	}).(ObjectBucketAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketAclInput)(nil)).Elem(), &ObjectBucketAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketAclArrayInput)(nil)).Elem(), ObjectBucketAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectBucketAclMapInput)(nil)).Elem(), ObjectBucketAclMap{})
	pulumi.RegisterOutputType(ObjectBucketAclOutput{})
	pulumi.RegisterOutputType(ObjectBucketAclArrayOutput{})
	pulumi.RegisterOutputType(ObjectBucketAclMapOutput{})
}
