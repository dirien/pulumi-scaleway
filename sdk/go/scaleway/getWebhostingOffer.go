// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a webhosting offer.
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
//			_, err := scaleway.GetWebhostingOffer(ctx, &scaleway.GetWebhostingOfferArgs{
//				Name: pulumi.StringRef("performance"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.GetWebhostingOffer(ctx, &scaleway.GetWebhostingOfferArgs{
//				OfferId: pulumi.StringRef("de2426b4-a9e9-11ec-b909-0242ac120002"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetWebhostingOffer(ctx *pulumi.Context, args *GetWebhostingOfferArgs, opts ...pulumi.InvokeOption) (*GetWebhostingOfferResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetWebhostingOfferResult
	err := ctx.Invoke("scaleway:index/getWebhostingOffer:getWebhostingOffer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebhostingOffer.
type GetWebhostingOfferArgs struct {
	// The offer name. Only one of `name` and `offerId` should be specified.
	Name *string `pulumi:"name"`
	// The offer id. Only one of `name` and `offerId` should be specified.
	OfferId *string `pulumi:"offerId"`
	// `region`) The region in which offer exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getWebhostingOffer.
type GetWebhostingOfferResult struct {
	// The unique identifier used for billing.
	BillingOperationPath string `pulumi:"billingOperationPath"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    *string `pulumi:"name"`
	OfferId *string `pulumi:"offerId"`
	// The offer price.
	Price string `pulumi:"price"`
	// The offer product.
	Products []GetWebhostingOfferProduct `pulumi:"products"`
	Region   string                      `pulumi:"region"`
}

func GetWebhostingOfferOutput(ctx *pulumi.Context, args GetWebhostingOfferOutputArgs, opts ...pulumi.InvokeOption) GetWebhostingOfferResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWebhostingOfferResult, error) {
			args := v.(GetWebhostingOfferArgs)
			r, err := GetWebhostingOffer(ctx, &args, opts...)
			var s GetWebhostingOfferResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWebhostingOfferResultOutput)
}

// A collection of arguments for invoking getWebhostingOffer.
type GetWebhostingOfferOutputArgs struct {
	// The offer name. Only one of `name` and `offerId` should be specified.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The offer id. Only one of `name` and `offerId` should be specified.
	OfferId pulumi.StringPtrInput `pulumi:"offerId"`
	// `region`) The region in which offer exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetWebhostingOfferOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhostingOfferArgs)(nil)).Elem()
}

// A collection of values returned by getWebhostingOffer.
type GetWebhostingOfferResultOutput struct{ *pulumi.OutputState }

func (GetWebhostingOfferResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWebhostingOfferResult)(nil)).Elem()
}

func (o GetWebhostingOfferResultOutput) ToGetWebhostingOfferResultOutput() GetWebhostingOfferResultOutput {
	return o
}

func (o GetWebhostingOfferResultOutput) ToGetWebhostingOfferResultOutputWithContext(ctx context.Context) GetWebhostingOfferResultOutput {
	return o
}

// The unique identifier used for billing.
func (o GetWebhostingOfferResultOutput) BillingOperationPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) string { return v.BillingOperationPath }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetWebhostingOfferResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetWebhostingOfferResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetWebhostingOfferResultOutput) OfferId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) *string { return v.OfferId }).(pulumi.StringPtrOutput)
}

// The offer price.
func (o GetWebhostingOfferResultOutput) Price() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) string { return v.Price }).(pulumi.StringOutput)
}

// The offer product.
func (o GetWebhostingOfferResultOutput) Products() GetWebhostingOfferProductArrayOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) []GetWebhostingOfferProduct { return v.Products }).(GetWebhostingOfferProductArrayOutput)
}

func (o GetWebhostingOfferResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetWebhostingOfferResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWebhostingOfferResultOutput{})
}
