// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Web Hostings.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/webhosting/).
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
//			byName, err := scaleway.GetWebhostingOffer(ctx, &scaleway.GetWebhostingOfferArgs{
//				Name: pulumi.StringRef("lite"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewWebhosting(ctx, "main", &scaleway.WebhostingArgs{
//				OfferId: pulumi.String(byName.OfferId),
//				Email:   pulumi.String("your@email.com"),
//				Domain:  pulumi.String("yourdomain.com"),
//				Tags: pulumi.StringArray{
//					pulumi.String("webhosting"),
//					pulumi.String("provider"),
//					pulumi.String("terraform"),
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
// ## Import
//
// Hostings can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/webhosting:Webhosting hosting01 fr-par/11111111-1111-1111-1111-111111111111
// ```
type Webhosting struct {
	pulumi.CustomResourceState

	// The URL to connect to cPanel Dashboard and to Webmail interface.
	CpanelUrls WebhostingCpanelUrlArrayOutput `pulumi:"cpanelUrls"`
	// Date and time of hosting's creation (RFC 3339 format).
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The DNS status of the hosting.
	DnsStatus pulumi.StringOutput `pulumi:"dnsStatus"`
	// The domain name of the hosting.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The contact email of the client for the hosting.
	Email pulumi.StringOutput `pulumi:"email"`
	// The ID of the selected offer for the hosting.
	OfferId pulumi.StringOutput `pulumi:"offerId"`
	// The name of the active offer.
	OfferName pulumi.StringOutput `pulumi:"offerName"`
	// The IDs of the selected options for the hosting.
	OptionIds pulumi.StringArrayOutput `pulumi:"optionIds"`
	// The active options of the hosting.
	Options WebhostingOptionArrayOutput `pulumi:"options"`
	// The organization ID the hosting is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The hostname of the host platform.
	PlatformHostname pulumi.StringOutput `pulumi:"platformHostname"`
	// The number of the host platform.
	PlatformNumber pulumi.IntOutput `pulumi:"platformNumber"`
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region of the Hosting.
	Region pulumi.StringOutput `pulumi:"region"`
	// The hosting status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags associated with the hosting.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Date and time of hosting's last update (RFC 3339 format).
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The main hosting cPanel username.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewWebhosting registers a new resource with the given unique name, arguments, and options.
func NewWebhosting(ctx *pulumi.Context,
	name string, args *WebhostingArgs, opts ...pulumi.ResourceOption) (*Webhosting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.OfferId == nil {
		return nil, errors.New("invalid value for required argument 'OfferId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Webhosting
	err := ctx.RegisterResource("scaleway:index/webhosting:Webhosting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhosting gets an existing Webhosting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhosting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhostingState, opts ...pulumi.ResourceOption) (*Webhosting, error) {
	var resource Webhosting
	err := ctx.ReadResource("scaleway:index/webhosting:Webhosting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhosting resources.
type webhostingState struct {
	// The URL to connect to cPanel Dashboard and to Webmail interface.
	CpanelUrls []WebhostingCpanelUrl `pulumi:"cpanelUrls"`
	// Date and time of hosting's creation (RFC 3339 format).
	CreatedAt *string `pulumi:"createdAt"`
	// The DNS status of the hosting.
	DnsStatus *string `pulumi:"dnsStatus"`
	// The domain name of the hosting.
	Domain *string `pulumi:"domain"`
	// The contact email of the client for the hosting.
	Email *string `pulumi:"email"`
	// The ID of the selected offer for the hosting.
	OfferId *string `pulumi:"offerId"`
	// The name of the active offer.
	OfferName *string `pulumi:"offerName"`
	// The IDs of the selected options for the hosting.
	OptionIds []string `pulumi:"optionIds"`
	// The active options of the hosting.
	Options []WebhostingOption `pulumi:"options"`
	// The organization ID the hosting is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// The hostname of the host platform.
	PlatformHostname *string `pulumi:"platformHostname"`
	// The number of the host platform.
	PlatformNumber *int `pulumi:"platformNumber"`
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the Hosting.
	Region *string `pulumi:"region"`
	// The hosting status.
	Status *string `pulumi:"status"`
	// The tags associated with the hosting.
	Tags []string `pulumi:"tags"`
	// Date and time of hosting's last update (RFC 3339 format).
	UpdatedAt *string `pulumi:"updatedAt"`
	// The main hosting cPanel username.
	Username *string `pulumi:"username"`
}

type WebhostingState struct {
	// The URL to connect to cPanel Dashboard and to Webmail interface.
	CpanelUrls WebhostingCpanelUrlArrayInput
	// Date and time of hosting's creation (RFC 3339 format).
	CreatedAt pulumi.StringPtrInput
	// The DNS status of the hosting.
	DnsStatus pulumi.StringPtrInput
	// The domain name of the hosting.
	Domain pulumi.StringPtrInput
	// The contact email of the client for the hosting.
	Email pulumi.StringPtrInput
	// The ID of the selected offer for the hosting.
	OfferId pulumi.StringPtrInput
	// The name of the active offer.
	OfferName pulumi.StringPtrInput
	// The IDs of the selected options for the hosting.
	OptionIds pulumi.StringArrayInput
	// The active options of the hosting.
	Options WebhostingOptionArrayInput
	// The organization ID the hosting is associated with.
	OrganizationId pulumi.StringPtrInput
	// The hostname of the host platform.
	PlatformHostname pulumi.StringPtrInput
	// The number of the host platform.
	PlatformNumber pulumi.IntPtrInput
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the Hosting.
	Region pulumi.StringPtrInput
	// The hosting status.
	Status pulumi.StringPtrInput
	// The tags associated with the hosting.
	Tags pulumi.StringArrayInput
	// Date and time of hosting's last update (RFC 3339 format).
	UpdatedAt pulumi.StringPtrInput
	// The main hosting cPanel username.
	Username pulumi.StringPtrInput
}

func (WebhostingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhostingState)(nil)).Elem()
}

type webhostingArgs struct {
	// The domain name of the hosting.
	Domain string `pulumi:"domain"`
	// The contact email of the client for the hosting.
	Email string `pulumi:"email"`
	// The ID of the selected offer for the hosting.
	OfferId string `pulumi:"offerId"`
	// The IDs of the selected options for the hosting.
	OptionIds []string `pulumi:"optionIds"`
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region of the Hosting.
	Region *string `pulumi:"region"`
	// The tags associated with the hosting.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Webhosting resource.
type WebhostingArgs struct {
	// The domain name of the hosting.
	Domain pulumi.StringInput
	// The contact email of the client for the hosting.
	Email pulumi.StringInput
	// The ID of the selected offer for the hosting.
	OfferId pulumi.StringInput
	// The IDs of the selected options for the hosting.
	OptionIds pulumi.StringArrayInput
	// `projectId`) The ID of the project the VPC is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region of the Hosting.
	Region pulumi.StringPtrInput
	// The tags associated with the hosting.
	Tags pulumi.StringArrayInput
}

func (WebhostingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhostingArgs)(nil)).Elem()
}

type WebhostingInput interface {
	pulumi.Input

	ToWebhostingOutput() WebhostingOutput
	ToWebhostingOutputWithContext(ctx context.Context) WebhostingOutput
}

func (*Webhosting) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhosting)(nil)).Elem()
}

func (i *Webhosting) ToWebhostingOutput() WebhostingOutput {
	return i.ToWebhostingOutputWithContext(context.Background())
}

func (i *Webhosting) ToWebhostingOutputWithContext(ctx context.Context) WebhostingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhostingOutput)
}

// WebhostingArrayInput is an input type that accepts WebhostingArray and WebhostingArrayOutput values.
// You can construct a concrete instance of `WebhostingArrayInput` via:
//
//	WebhostingArray{ WebhostingArgs{...} }
type WebhostingArrayInput interface {
	pulumi.Input

	ToWebhostingArrayOutput() WebhostingArrayOutput
	ToWebhostingArrayOutputWithContext(context.Context) WebhostingArrayOutput
}

type WebhostingArray []WebhostingInput

func (WebhostingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhosting)(nil)).Elem()
}

func (i WebhostingArray) ToWebhostingArrayOutput() WebhostingArrayOutput {
	return i.ToWebhostingArrayOutputWithContext(context.Background())
}

func (i WebhostingArray) ToWebhostingArrayOutputWithContext(ctx context.Context) WebhostingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhostingArrayOutput)
}

// WebhostingMapInput is an input type that accepts WebhostingMap and WebhostingMapOutput values.
// You can construct a concrete instance of `WebhostingMapInput` via:
//
//	WebhostingMap{ "key": WebhostingArgs{...} }
type WebhostingMapInput interface {
	pulumi.Input

	ToWebhostingMapOutput() WebhostingMapOutput
	ToWebhostingMapOutputWithContext(context.Context) WebhostingMapOutput
}

type WebhostingMap map[string]WebhostingInput

func (WebhostingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhosting)(nil)).Elem()
}

func (i WebhostingMap) ToWebhostingMapOutput() WebhostingMapOutput {
	return i.ToWebhostingMapOutputWithContext(context.Background())
}

func (i WebhostingMap) ToWebhostingMapOutputWithContext(ctx context.Context) WebhostingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhostingMapOutput)
}

type WebhostingOutput struct{ *pulumi.OutputState }

func (WebhostingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhosting)(nil)).Elem()
}

func (o WebhostingOutput) ToWebhostingOutput() WebhostingOutput {
	return o
}

func (o WebhostingOutput) ToWebhostingOutputWithContext(ctx context.Context) WebhostingOutput {
	return o
}

// The URL to connect to cPanel Dashboard and to Webmail interface.
func (o WebhostingOutput) CpanelUrls() WebhostingCpanelUrlArrayOutput {
	return o.ApplyT(func(v *Webhosting) WebhostingCpanelUrlArrayOutput { return v.CpanelUrls }).(WebhostingCpanelUrlArrayOutput)
}

// Date and time of hosting's creation (RFC 3339 format).
func (o WebhostingOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The DNS status of the hosting.
func (o WebhostingOutput) DnsStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.DnsStatus }).(pulumi.StringOutput)
}

// The domain name of the hosting.
func (o WebhostingOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The contact email of the client for the hosting.
func (o WebhostingOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The ID of the selected offer for the hosting.
func (o WebhostingOutput) OfferId() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.OfferId }).(pulumi.StringOutput)
}

// The name of the active offer.
func (o WebhostingOutput) OfferName() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.OfferName }).(pulumi.StringOutput)
}

// The IDs of the selected options for the hosting.
func (o WebhostingOutput) OptionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringArrayOutput { return v.OptionIds }).(pulumi.StringArrayOutput)
}

// The active options of the hosting.
func (o WebhostingOutput) Options() WebhostingOptionArrayOutput {
	return o.ApplyT(func(v *Webhosting) WebhostingOptionArrayOutput { return v.Options }).(WebhostingOptionArrayOutput)
}

// The organization ID the hosting is associated with.
func (o WebhostingOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The hostname of the host platform.
func (o WebhostingOutput) PlatformHostname() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.PlatformHostname }).(pulumi.StringOutput)
}

// The number of the host platform.
func (o WebhostingOutput) PlatformNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.IntOutput { return v.PlatformNumber }).(pulumi.IntOutput)
}

// `projectId`) The ID of the project the VPC is associated with.
func (o WebhostingOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region of the Hosting.
func (o WebhostingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The hosting status.
func (o WebhostingOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the hosting.
func (o WebhostingOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Date and time of hosting's last update (RFC 3339 format).
func (o WebhostingOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The main hosting cPanel username.
func (o WebhostingOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Webhosting) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type WebhostingArrayOutput struct{ *pulumi.OutputState }

func (WebhostingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Webhosting)(nil)).Elem()
}

func (o WebhostingArrayOutput) ToWebhostingArrayOutput() WebhostingArrayOutput {
	return o
}

func (o WebhostingArrayOutput) ToWebhostingArrayOutputWithContext(ctx context.Context) WebhostingArrayOutput {
	return o
}

func (o WebhostingArrayOutput) Index(i pulumi.IntInput) WebhostingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Webhosting {
		return vs[0].([]*Webhosting)[vs[1].(int)]
	}).(WebhostingOutput)
}

type WebhostingMapOutput struct{ *pulumi.OutputState }

func (WebhostingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Webhosting)(nil)).Elem()
}

func (o WebhostingMapOutput) ToWebhostingMapOutput() WebhostingMapOutput {
	return o
}

func (o WebhostingMapOutput) ToWebhostingMapOutputWithContext(ctx context.Context) WebhostingMapOutput {
	return o
}

func (o WebhostingMapOutput) MapIndex(k pulumi.StringInput) WebhostingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Webhosting {
		return vs[0].(map[string]*Webhosting)[vs[1].(string)]
	}).(WebhostingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WebhostingInput)(nil)).Elem(), &Webhosting{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhostingArrayInput)(nil)).Elem(), WebhostingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebhostingMapInput)(nil)).Elem(), WebhostingMap{})
	pulumi.RegisterOutputType(WebhostingOutput{})
	pulumi.RegisterOutputType(WebhostingArrayOutput{})
	pulumi.RegisterOutputType(WebhostingMapOutput{})
}