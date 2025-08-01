// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `AccountProject` data source is used to retrieve information about a Scaleway project.
//
// Refer to the Organizations and Projects [documentation](https://www.scaleway.com/en/docs/organizations-and-projects/) and [API documentation](https://www.scaleway.com/en/developers/api/account/project-api/) for more information.
//
// ## Retrieve a Scaleway Project
//
// The following commands allow you to:
//
// - retrieve a Project by its name
// - retrieve a Project by its ID
// - retrieve the default project of an Organization
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
//			_, err := scaleway.LookupAccountProject(ctx, &scaleway.LookupAccountProjectArgs{
//				Name: pulumi.StringRef("default"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.LookupAccountProject(ctx, &scaleway.LookupAccountProjectArgs{
//				ProjectId: pulumi.StringRef("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAccountProject(ctx *pulumi.Context, args *LookupAccountProjectArgs, opts ...pulumi.InvokeOption) (*LookupAccountProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountProjectResult
	err := ctx.Invoke("scaleway:index/getAccountProject:getAccountProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccountProject.
type LookupAccountProjectArgs struct {
	// The name of the Project.
	// Only one of the `name` and `projectId` should be specified.
	Name *string `pulumi:"name"`
	// The unique identifier of the Organization with which the Project is associated.
	//
	// If no default `organizationId` is set, one must be set explicitly in this datasource
	OrganizationId *string `pulumi:"organizationId"`
	// The unique identifier of the Project.
	// Only one of the `name` and `projectId` should be specified.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getAccountProject.
type LookupAccountProjectResult struct {
	CreatedAt   string `pulumi:"createdAt"`
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	OrganizationId *string `pulumi:"organizationId"`
	ProjectId      string  `pulumi:"projectId"`
	UpdatedAt      string  `pulumi:"updatedAt"`
}

func LookupAccountProjectOutput(ctx *pulumi.Context, args LookupAccountProjectOutputArgs, opts ...pulumi.InvokeOption) LookupAccountProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAccountProjectResultOutput, error) {
			args := v.(LookupAccountProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getAccountProject:getAccountProject", args, LookupAccountProjectResultOutput{}, options).(LookupAccountProjectResultOutput), nil
		}).(LookupAccountProjectResultOutput)
}

// A collection of arguments for invoking getAccountProject.
type LookupAccountProjectOutputArgs struct {
	// The name of the Project.
	// Only one of the `name` and `projectId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The unique identifier of the Organization with which the Project is associated.
	//
	// If no default `organizationId` is set, one must be set explicitly in this datasource
	OrganizationId pulumi.StringPtrInput `pulumi:"organizationId"`
	// The unique identifier of the Project.
	// Only one of the `name` and `projectId` should be specified.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupAccountProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountProjectArgs)(nil)).Elem()
}

// A collection of values returned by getAccountProject.
type LookupAccountProjectResultOutput struct{ *pulumi.OutputState }

func (LookupAccountProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountProjectResult)(nil)).Elem()
}

func (o LookupAccountProjectResultOutput) ToLookupAccountProjectResultOutput() LookupAccountProjectResultOutput {
	return o
}

func (o LookupAccountProjectResultOutput) ToLookupAccountProjectResultOutputWithContext(ctx context.Context) LookupAccountProjectResultOutput {
	return o
}

func (o LookupAccountProjectResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupAccountProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccountProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAccountProjectResultOutput) OrganizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) *string { return v.OrganizationId }).(pulumi.StringPtrOutput)
}

func (o LookupAccountProjectResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupAccountProjectResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountProjectResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountProjectResultOutput{})
}
