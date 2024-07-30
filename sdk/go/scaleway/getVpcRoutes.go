// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about multiple VPC routes.
func GetVpcRoutes(ctx *pulumi.Context, args *GetVpcRoutesArgs, opts ...pulumi.InvokeOption) (*GetVpcRoutesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVpcRoutesResult
	err := ctx.Invoke("scaleway:index/getVpcRoutes:getVpcRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcRoutes.
type GetVpcRoutesArgs struct {
	// Routes with an IPv6 destination will be listed.
	IsIpv6 *bool `pulumi:"isIpv6"`
	// The next hop private network ID to filter for. routes with a similar next hop private network ID are listed.
	NexthopPrivateNetworkId *string `pulumi:"nexthopPrivateNetworkId"`
	// The next hop resource ID to filter for. routes with a similar next hop resource ID are listed.
	NexthopResourceId *string `pulumi:"nexthopResourceId"`
	// The next hop resource type to filter for. routes with a similar next hop resource type are listed.
	NexthopResourceType *string `pulumi:"nexthopResourceType"`
	// `region`). The region in which the routes exist.
	Region *string `pulumi:"region"`
	// List of tags to filter for. routes with these exact tags are listed.
	Tags []string `pulumi:"tags"`
	// The VPC ID to filter for. routes with a similar VPC ID are listed.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getVpcRoutes.
type GetVpcRoutesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                      string  `pulumi:"id"`
	IsIpv6                  *bool   `pulumi:"isIpv6"`
	NexthopPrivateNetworkId *string `pulumi:"nexthopPrivateNetworkId"`
	NexthopResourceId       *string `pulumi:"nexthopResourceId"`
	NexthopResourceType     *string `pulumi:"nexthopResourceType"`
	Region                  string  `pulumi:"region"`
	// List of retrieved routes
	Routes []GetVpcRoutesRoute `pulumi:"routes"`
	Tags   []string            `pulumi:"tags"`
	VpcId  *string             `pulumi:"vpcId"`
}

func GetVpcRoutesOutput(ctx *pulumi.Context, args GetVpcRoutesOutputArgs, opts ...pulumi.InvokeOption) GetVpcRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVpcRoutesResult, error) {
			args := v.(GetVpcRoutesArgs)
			r, err := GetVpcRoutes(ctx, &args, opts...)
			var s GetVpcRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVpcRoutesResultOutput)
}

// A collection of arguments for invoking getVpcRoutes.
type GetVpcRoutesOutputArgs struct {
	// Routes with an IPv6 destination will be listed.
	IsIpv6 pulumi.BoolPtrInput `pulumi:"isIpv6"`
	// The next hop private network ID to filter for. routes with a similar next hop private network ID are listed.
	NexthopPrivateNetworkId pulumi.StringPtrInput `pulumi:"nexthopPrivateNetworkId"`
	// The next hop resource ID to filter for. routes with a similar next hop resource ID are listed.
	NexthopResourceId pulumi.StringPtrInput `pulumi:"nexthopResourceId"`
	// The next hop resource type to filter for. routes with a similar next hop resource type are listed.
	NexthopResourceType pulumi.StringPtrInput `pulumi:"nexthopResourceType"`
	// `region`). The region in which the routes exist.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// List of tags to filter for. routes with these exact tags are listed.
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	// The VPC ID to filter for. routes with a similar VPC ID are listed.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetVpcRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcRoutesArgs)(nil)).Elem()
}

// A collection of values returned by getVpcRoutes.
type GetVpcRoutesResultOutput struct{ *pulumi.OutputState }

func (GetVpcRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVpcRoutesResult)(nil)).Elem()
}

func (o GetVpcRoutesResultOutput) ToGetVpcRoutesResultOutput() GetVpcRoutesResultOutput {
	return o
}

func (o GetVpcRoutesResultOutput) ToGetVpcRoutesResultOutputWithContext(ctx context.Context) GetVpcRoutesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetVpcRoutesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVpcRoutesResultOutput) IsIpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) *bool { return v.IsIpv6 }).(pulumi.BoolPtrOutput)
}

func (o GetVpcRoutesResultOutput) NexthopPrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) *string { return v.NexthopPrivateNetworkId }).(pulumi.StringPtrOutput)
}

func (o GetVpcRoutesResultOutput) NexthopResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) *string { return v.NexthopResourceId }).(pulumi.StringPtrOutput)
}

func (o GetVpcRoutesResultOutput) NexthopResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) *string { return v.NexthopResourceType }).(pulumi.StringPtrOutput)
}

func (o GetVpcRoutesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) string { return v.Region }).(pulumi.StringOutput)
}

// List of retrieved routes
func (o GetVpcRoutesResultOutput) Routes() GetVpcRoutesRouteArrayOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) []GetVpcRoutesRoute { return v.Routes }).(GetVpcRoutesRouteArrayOutput)
}

func (o GetVpcRoutesResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GetVpcRoutesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVpcRoutesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVpcRoutesResultOutput{})
}
