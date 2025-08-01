// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `getAvailabilityZones` data source is used to retrieve information about the available zones based on its Region.
//
// For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
// you can choose the location that better fits your need (country, latency, etc.).
//
// Refer to the Account [documentation](https://www.scaleway.com/en/docs/console/account/reference-content/products-availability/) for more information.
//
// ## Retrieve the Availability Zones of a Region
//
// The following command allow you to retrieve a the AZs of a Region.
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
//			_, err := scaleway.GetAvailabilityZones(ctx, &scaleway.GetAvailabilityZonesArgs{
//				Region: pulumi.StringRef("nl-ams"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("scaleway:index/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// Region is represented as a Geographical area, such as France. Defaults to `fr-par`.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Region *string `pulumi:"region"`
	// The list of availability zones in each Region
	Zones []string `pulumi:"zones"`
}

func GetAvailabilityZonesOutput(ctx *pulumi.Context, args GetAvailabilityZonesOutputArgs, opts ...pulumi.InvokeOption) GetAvailabilityZonesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetAvailabilityZonesResultOutput, error) {
			args := v.(GetAvailabilityZonesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getAvailabilityZones:getAvailabilityZones", args, GetAvailabilityZonesResultOutput{}, options).(GetAvailabilityZonesResultOutput), nil
		}).(GetAvailabilityZonesResultOutput)
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesOutputArgs struct {
	// Region is represented as a Geographical area, such as France. Defaults to `fr-par`.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetAvailabilityZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesArgs)(nil)).Elem()
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResultOutput struct{ *pulumi.OutputState }

func (GetAvailabilityZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAvailabilityZonesResult)(nil)).Elem()
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutput() GetAvailabilityZonesResultOutput {
	return o
}

func (o GetAvailabilityZonesResultOutput) ToGetAvailabilityZonesResultOutputWithContext(ctx context.Context) GetAvailabilityZonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetAvailabilityZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetAvailabilityZonesResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The list of availability zones in each Region
func (o GetAvailabilityZonesResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAvailabilityZonesResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAvailabilityZonesResultOutput{})
}
