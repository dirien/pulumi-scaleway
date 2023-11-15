// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about Scaleway Secrets.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/).
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
//			_, err := scaleway.NewSecret(ctx, "main", &scaleway.SecretArgs{
//				Description: pulumi.String("barr"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupSecret(ctx, &scaleway.LookupSecretArgs{
//				SecretId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupSecret(ctx, &scaleway.LookupSecretArgs{
//				Name: pulumi.StringRef("your_secret_name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSecret(ctx *pulumi.Context, args *LookupSecretArgs, opts ...pulumi.InvokeOption) (*LookupSecretResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecretResult
	err := ctx.Invoke("scaleway:index/getSecret:getSecret", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecret.
type LookupSecretArgs struct {
	// The secret name.
	// Only one of `name` and `secretId` should be specified.
	Name *string `pulumi:"name"`
	// The organization ID the Project is associated with.
	// If no default organizationId is set, one must be set explicitly in this datasource
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The ID of the
	// project the secret is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the secret exists.
	Region *string `pulumi:"region"`
	// The secret id.
	// Only one of `name` and `secretId` should be specified.
	SecretId *string `pulumi:"secretId"`
}

// A collection of values returned by getSecret.
type LookupSecretResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	Name           *string  `pulumi:"name"`
	OrganizationId string   `pulumi:"organizationId"`
	ProjectId      *string  `pulumi:"projectId"`
	Region         *string  `pulumi:"region"`
	SecretId       *string  `pulumi:"secretId"`
	Status         string   `pulumi:"status"`
	Tags           []string `pulumi:"tags"`
	UpdatedAt      string   `pulumi:"updatedAt"`
	VersionCount   int      `pulumi:"versionCount"`
}

func LookupSecretOutput(ctx *pulumi.Context, args LookupSecretOutputArgs, opts ...pulumi.InvokeOption) LookupSecretResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecretResult, error) {
			args := v.(LookupSecretArgs)
			r, err := LookupSecret(ctx, &args, opts...)
			var s LookupSecretResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecretResultOutput)
}

// A collection of arguments for invoking getSecret.
type LookupSecretOutputArgs struct {
	// The secret name.
	// Only one of `name` and `secretId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The organization ID the Project is associated with.
	// If no default organizationId is set, one must be set explicitly in this datasource
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
	// `projectId`) The ID of the
	// project the secret is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the secret exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The secret id.
	// Only one of `name` and `secretId` should be specified.
	SecretId pulumi.StringPtrInput `pulumi:"secretId"`
}

func (LookupSecretOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretArgs)(nil)).Elem()
}

// A collection of values returned by getSecret.
type LookupSecretResultOutput struct{ *pulumi.OutputState }

func (LookupSecretResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecretResult)(nil)).Elem()
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutput() LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) ToLookupSecretResultOutputWithContext(ctx context.Context) LookupSecretResultOutput {
	return o
}

func (o LookupSecretResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecretResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) SecretId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecretResult) *string { return v.SecretId }).(pulumi.StringPtrOutput)
}

func (o LookupSecretResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSecretResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupSecretResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecretResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupSecretResultOutput) VersionCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecretResult) int { return v.VersionCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecretResultOutput{})
}
