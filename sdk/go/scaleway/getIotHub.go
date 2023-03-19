// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an IOT Hub.
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
//			_, err := scaleway.LookupIotHub(ctx, &scaleway.LookupIotHubArgs{
//				HubId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIotHub(ctx *pulumi.Context, args *LookupIotHubArgs, opts ...pulumi.InvokeOption) (*LookupIotHubResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupIotHubResult
	err := ctx.Invoke("scaleway:index/getIotHub:getIotHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIotHub.
type LookupIotHubArgs struct {
	// The Hub ID.
	// Only one of the `name` and `hubId` should be specified.
	HubId *string `pulumi:"hubId"`
	// The name of the Hub.
	// Only one of the `name` and `hubId` should be specified.
	Name *string `pulumi:"name"`
	// `region`) The region in which the hub exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getIotHub.
type LookupIotHubResult struct {
	ConnectedDeviceCount   int     `pulumi:"connectedDeviceCount"`
	CreatedAt              string  `pulumi:"createdAt"`
	DeviceAutoProvisioning bool    `pulumi:"deviceAutoProvisioning"`
	DeviceCount            int     `pulumi:"deviceCount"`
	DisableEvents          bool    `pulumi:"disableEvents"`
	Enabled                bool    `pulumi:"enabled"`
	Endpoint               string  `pulumi:"endpoint"`
	EventsTopicPrefix      string  `pulumi:"eventsTopicPrefix"`
	HubCa                  string  `pulumi:"hubCa"`
	HubCaChallenge         string  `pulumi:"hubCaChallenge"`
	HubId                  *string `pulumi:"hubId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string  `pulumi:"id"`
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	ProductPlan    string  `pulumi:"productPlan"`
	ProjectId      string  `pulumi:"projectId"`
	Region         *string `pulumi:"region"`
	Status         string  `pulumi:"status"`
	UpdatedAt      string  `pulumi:"updatedAt"`
}

func LookupIotHubOutput(ctx *pulumi.Context, args LookupIotHubOutputArgs, opts ...pulumi.InvokeOption) LookupIotHubResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIotHubResult, error) {
			args := v.(LookupIotHubArgs)
			r, err := LookupIotHub(ctx, &args, opts...)
			var s LookupIotHubResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIotHubResultOutput)
}

// A collection of arguments for invoking getIotHub.
type LookupIotHubOutputArgs struct {
	// The Hub ID.
	// Only one of the `name` and `hubId` should be specified.
	HubId pulumi.StringPtrInput `pulumi:"hubId"`
	// The name of the Hub.
	// Only one of the `name` and `hubId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// `region`) The region in which the hub exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupIotHubOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubArgs)(nil)).Elem()
}

// A collection of values returned by getIotHub.
type LookupIotHubResultOutput struct{ *pulumi.OutputState }

func (LookupIotHubResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIotHubResult)(nil)).Elem()
}

func (o LookupIotHubResultOutput) ToLookupIotHubResultOutput() LookupIotHubResultOutput {
	return o
}

func (o LookupIotHubResultOutput) ToLookupIotHubResultOutputWithContext(ctx context.Context) LookupIotHubResultOutput {
	return o
}

func (o LookupIotHubResultOutput) ConnectedDeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIotHubResult) int { return v.ConnectedDeviceCount }).(pulumi.IntOutput)
}

func (o LookupIotHubResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) DeviceAutoProvisioning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIotHubResult) bool { return v.DeviceAutoProvisioning }).(pulumi.BoolOutput)
}

func (o LookupIotHubResultOutput) DeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIotHubResult) int { return v.DeviceCount }).(pulumi.IntOutput)
}

func (o LookupIotHubResultOutput) DisableEvents() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIotHubResult) bool { return v.DisableEvents }).(pulumi.BoolOutput)
}

func (o LookupIotHubResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIotHubResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupIotHubResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) EventsTopicPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.EventsTopicPrefix }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) HubCa() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.HubCa }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) HubCaChallenge() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.HubCaChallenge }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) HubId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResult) *string { return v.HubId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIotHubResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubResultOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) ProductPlan() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.ProductPlan }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIotHubResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupIotHubResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupIotHubResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIotHubResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIotHubResultOutput{})
}
