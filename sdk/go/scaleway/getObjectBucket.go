// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `ObjectBucket` data source is used to retrieve information about an Object Storage bucket.
//
// Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/storage/object/how-to/create-a-bucket/) for more information.
//
// ## Retrieve an Object Storage bucket
//
// The following commands allow you to:
//
// - retrieve a bucket by its name
// - retrieve a bucket by its ID
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
//			main, err := scaleway.NewObjectBucket(ctx, "main", &scaleway.ObjectBucketArgs{
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupObjectBucketOutput(ctx, scaleway.GetObjectBucketOutputArgs{
//				Name: main.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
//
// ## Retrieve a bucket from a specific project
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
//			_, err := scaleway.LookupObjectBucket(ctx, &scaleway.LookupObjectBucketArgs{
//				Name:      pulumi.StringRef("bucket.test.com"),
//				ProjectId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupObjectBucket(ctx *pulumi.Context, args *LookupObjectBucketArgs, opts ...pulumi.InvokeOption) (*LookupObjectBucketResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupObjectBucketResult
	err := ctx.Invoke("scaleway:index/getObjectBucket:getObjectBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getObjectBucket.
type LookupObjectBucketArgs struct {
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project with which the bucket is associated.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the bucket exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getObjectBucket.
type LookupObjectBucketResult struct {
	Acl         string                    `pulumi:"acl"`
	ApiEndpoint string                    `pulumi:"apiEndpoint"`
	CorsRules   []GetObjectBucketCorsRule `pulumi:"corsRules"`
	// The endpoint URL of the bucket
	Endpoint     string `pulumi:"endpoint"`
	ForceDestroy bool   `pulumi:"forceDestroy"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                         `pulumi:"id"`
	LifecycleRules    []GetObjectBucketLifecycleRule `pulumi:"lifecycleRules"`
	Name              *string                        `pulumi:"name"`
	ObjectLockEnabled bool                           `pulumi:"objectLockEnabled"`
	ProjectId         *string                        `pulumi:"projectId"`
	Region            *string                        `pulumi:"region"`
	Tags              map[string]string              `pulumi:"tags"`
	Versionings       []GetObjectBucketVersioning    `pulumi:"versionings"`
}

func LookupObjectBucketOutput(ctx *pulumi.Context, args LookupObjectBucketOutputArgs, opts ...pulumi.InvokeOption) LookupObjectBucketResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupObjectBucketResultOutput, error) {
			args := v.(LookupObjectBucketArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupObjectBucketResult
			secret, err := ctx.InvokePackageRaw("scaleway:index/getObjectBucket:getObjectBucket", args, &rv, "", opts...)
			if err != nil {
				return LookupObjectBucketResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupObjectBucketResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupObjectBucketResultOutput), nil
			}
			return output, nil
		}).(LookupObjectBucketResultOutput)
}

// A collection of arguments for invoking getObjectBucket.
type LookupObjectBucketOutputArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `projectId`) The ID of the project with which the bucket is associated.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the bucket exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupObjectBucketOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectBucketArgs)(nil)).Elem()
}

// A collection of values returned by getObjectBucket.
type LookupObjectBucketResultOutput struct{ *pulumi.OutputState }

func (LookupObjectBucketResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupObjectBucketResult)(nil)).Elem()
}

func (o LookupObjectBucketResultOutput) ToLookupObjectBucketResultOutput() LookupObjectBucketResultOutput {
	return o
}

func (o LookupObjectBucketResultOutput) ToLookupObjectBucketResultOutputWithContext(ctx context.Context) LookupObjectBucketResultOutput {
	return o
}

func (o LookupObjectBucketResultOutput) Acl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.Acl }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.ApiEndpoint }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) CorsRules() GetObjectBucketCorsRuleArrayOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) []GetObjectBucketCorsRule { return v.CorsRules }).(GetObjectBucketCorsRuleArrayOutput)
}

// The endpoint URL of the bucket
func (o LookupObjectBucketResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) ForceDestroy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) bool { return v.ForceDestroy }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupObjectBucketResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupObjectBucketResultOutput) LifecycleRules() GetObjectBucketLifecycleRuleArrayOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) []GetObjectBucketLifecycleRule { return v.LifecycleRules }).(GetObjectBucketLifecycleRuleArrayOutput)
}

func (o LookupObjectBucketResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupObjectBucketResultOutput) ObjectLockEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) bool { return v.ObjectLockEnabled }).(pulumi.BoolOutput)
}

func (o LookupObjectBucketResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupObjectBucketResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupObjectBucketResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupObjectBucketResultOutput) Versionings() GetObjectBucketVersioningArrayOutput {
	return o.ApplyT(func(v LookupObjectBucketResult) []GetObjectBucketVersioning { return v.Versionings }).(GetObjectBucketVersioningArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupObjectBucketResultOutput{})
}
