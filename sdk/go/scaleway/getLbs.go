// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about multiple Load Balancers.
//
// For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/concepts/#load-balancers) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-load-balancer-list-load-balancers).
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
//			_, err := scaleway.GetLbs(ctx, &scaleway.GetLbsArgs{
//				Name: pulumi.StringRef("foobar"),
//				Zone: pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.GetLbs(ctx, &scaleway.GetLbsArgs{
//				Tags: []string{
//					"a tag",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLbs(ctx *pulumi.Context, args *GetLbsArgs, opts ...pulumi.InvokeOption) (*GetLbsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLbsResult
	err := ctx.Invoke("scaleway:index/getLbs:getLbs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbs.
type GetLbsArgs struct {
	// The Load Balancer name to filter for. Load Balancers with a matching name are listed.
	Name *string `pulumi:"name"`
	// The ID of the Project the Load Balancer is associated with.
	ProjectId *string `pulumi:"projectId"`
	// List of tags to filter for. Load Balancers with these exact tags are listed.
	Tags []string `pulumi:"tags"`
	// `zone`) The zone in which the Load Balancers exist.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getLbs.
type GetLbsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of retrieved Load Balancers
	Lbs []GetLbsLb `pulumi:"lbs"`
	// The name of the Load Balancer.
	Name *string `pulumi:"name"`
	// The ID of the Organization the Load Balancer is associated with.
	OrganizationId string `pulumi:"organizationId"`
	// The ID of the Project the Load Balancer is associated with.
	ProjectId string `pulumi:"projectId"`
	// The tags associated with the Load Balancer.
	Tags []string `pulumi:"tags"`
	// The zone of the Load Balancer.
	Zone string `pulumi:"zone"`
}

func GetLbsOutput(ctx *pulumi.Context, args GetLbsOutputArgs, opts ...pulumi.InvokeOption) GetLbsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLbsResult, error) {
			args := v.(GetLbsArgs)
			r, err := GetLbs(ctx, &args, opts...)
			var s GetLbsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLbsResultOutput)
}

// A collection of arguments for invoking getLbs.
type GetLbsOutputArgs struct {
	// The Load Balancer name to filter for. Load Balancers with a matching name are listed.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the Project the Load Balancer is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// List of tags to filter for. Load Balancers with these exact tags are listed.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// `zone`) The zone in which the Load Balancers exist.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetLbsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbsArgs)(nil)).Elem()
}

// A collection of values returned by getLbs.
type GetLbsResultOutput struct{ *pulumi.OutputState }

func (GetLbsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLbsResult)(nil)).Elem()
}

func (o GetLbsResultOutput) ToGetLbsResultOutput() GetLbsResultOutput {
	return o
}

func (o GetLbsResultOutput) ToGetLbsResultOutputWithContext(ctx context.Context) GetLbsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLbsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of retrieved Load Balancers
func (o GetLbsResultOutput) Lbs() GetLbsLbArrayOutput {
	return o.ApplyT(func(v GetLbsResult) []GetLbsLb { return v.Lbs }).(GetLbsLbArrayOutput)
}

// The name of the Load Balancer.
func (o GetLbsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLbsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID of the Organization the Load Balancer is associated with.
func (o GetLbsResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

// The ID of the Project the Load Balancer is associated with.
func (o GetLbsResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The tags associated with the Load Balancer.
func (o GetLbsResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLbsResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The zone of the Load Balancer.
func (o GetLbsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetLbsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLbsResultOutput{})
}
