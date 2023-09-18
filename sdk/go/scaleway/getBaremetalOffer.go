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

// Gets information about a baremetal offer. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
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
//			_, err := scaleway.GetBaremetalOffer(ctx, &scaleway.GetBaremetalOfferArgs{
//				OfferId: pulumi.StringRef("25dcf38b-c90c-4b18-97a2-6956e9d1e113"),
//				Zone:    pulumi.StringRef("fr-par-2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetBaremetalOffer(ctx *pulumi.Context, args *GetBaremetalOfferArgs, opts ...pulumi.InvokeOption) (*GetBaremetalOfferResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBaremetalOfferResult
	err := ctx.Invoke("scaleway:index/getBaremetalOffer:getBaremetalOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBaremetalOffer.
type GetBaremetalOfferArgs struct {
	IncludeDisabled *bool `pulumi:"includeDisabled"`
	// The offer name. Only one of `name` and `offerId` should be specified.
	Name *string `pulumi:"name"`
	// The offer id. Only one of `name` and `offerId` should be specified.
	OfferId *string `pulumi:"offerId"`
	// Period of subscription the desired offer. Should be `hourly` or `monthly`.
	SubscriptionPeriod *string `pulumi:"subscriptionPeriod"`
	// `zone`) The zone in which the offer should be created.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getBaremetalOffer.
type GetBaremetalOfferResult struct {
	// Available Bandwidth with the offer.
	Bandwidth int `pulumi:"bandwidth"`
	// Commercial range of the offer.
	CommercialRange string `pulumi:"commercialRange"`
	// A list of cpu specifications. (Structure is documented below.)
	Cpus []GetBaremetalOfferCpus `pulumi:"cpus"`
	// A list of disk specifications. (Structure is documented below.)
	Disks []GetBaremetalOfferDisk `pulumi:"disks"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	IncludeDisabled *bool  `pulumi:"includeDisabled"`
	// A list of memory specifications. (Structure is documented below.)
	Memories []GetBaremetalOfferMemory `pulumi:"memories"`
	// Name of the CPU.
	Name    *string `pulumi:"name"`
	OfferId *string `pulumi:"offerId"`
	// Stock status for this offer. Possible values are: `empty`, `low` or `available`.
	Stock              string  `pulumi:"stock"`
	SubscriptionPeriod *string `pulumi:"subscriptionPeriod"`
	Zone               string  `pulumi:"zone"`
}

func GetBaremetalOfferOutput(ctx *pulumi.Context, args GetBaremetalOfferOutputArgs, opts ...pulumi.InvokeOption) GetBaremetalOfferResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBaremetalOfferResult, error) {
			args := v.(GetBaremetalOfferArgs)
			r, err := GetBaremetalOffer(ctx, &args, opts...)
			var s GetBaremetalOfferResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBaremetalOfferResultOutput)
}

// A collection of arguments for invoking getBaremetalOffer.
type GetBaremetalOfferOutputArgs struct {
	IncludeDisabled pulumi.BoolPtrInput `pulumi:"includeDisabled"`
	// The offer name. Only one of `name` and `offerId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The offer id. Only one of `name` and `offerId` should be specified.
	OfferId pulumi.StringPtrInput `pulumi:"offerId"`
	// Period of subscription the desired offer. Should be `hourly` or `monthly`.
	SubscriptionPeriod pulumi.StringPtrInput `pulumi:"subscriptionPeriod"`
	// `zone`) The zone in which the offer should be created.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetBaremetalOfferOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaremetalOfferArgs)(nil)).Elem()
}

// A collection of values returned by getBaremetalOffer.
type GetBaremetalOfferResultOutput struct{ *pulumi.OutputState }

func (GetBaremetalOfferResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBaremetalOfferResult)(nil)).Elem()
}

func (o GetBaremetalOfferResultOutput) ToGetBaremetalOfferResultOutput() GetBaremetalOfferResultOutput {
	return o
}

func (o GetBaremetalOfferResultOutput) ToGetBaremetalOfferResultOutputWithContext(ctx context.Context) GetBaremetalOfferResultOutput {
	return o
}

func (o GetBaremetalOfferResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetBaremetalOfferResult] {
	return pulumix.Output[GetBaremetalOfferResult]{
		OutputState: o.OutputState,
	}
}

// Available Bandwidth with the offer.
func (o GetBaremetalOfferResultOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) int { return v.Bandwidth }).(pulumi.IntOutput)
}

// Commercial range of the offer.
func (o GetBaremetalOfferResultOutput) CommercialRange() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.CommercialRange }).(pulumi.StringOutput)
}

// A list of cpu specifications. (Structure is documented below.)
func (o GetBaremetalOfferResultOutput) Cpus() GetBaremetalOfferCpusArrayOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) []GetBaremetalOfferCpus { return v.Cpus }).(GetBaremetalOfferCpusArrayOutput)
}

// A list of disk specifications. (Structure is documented below.)
func (o GetBaremetalOfferResultOutput) Disks() GetBaremetalOfferDiskArrayOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) []GetBaremetalOfferDisk { return v.Disks }).(GetBaremetalOfferDiskArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBaremetalOfferResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBaremetalOfferResultOutput) IncludeDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *bool { return v.IncludeDisabled }).(pulumi.BoolPtrOutput)
}

// A list of memory specifications. (Structure is documented below.)
func (o GetBaremetalOfferResultOutput) Memories() GetBaremetalOfferMemoryArrayOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) []GetBaremetalOfferMemory { return v.Memories }).(GetBaremetalOfferMemoryArrayOutput)
}

// Name of the CPU.
func (o GetBaremetalOfferResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOfferResultOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

// Stock status for this offer. Possible values are: `empty`, `low` or `available`.
func (o GetBaremetalOfferResultOutput) Stock() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.Stock }).(pulumi.StringOutput)
}

func (o GetBaremetalOfferResultOutput) SubscriptionPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) *string { return v.SubscriptionPeriod }).(pulumi.StringPtrOutput)
}

func (o GetBaremetalOfferResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetBaremetalOfferResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBaremetalOfferResultOutput{})
}
