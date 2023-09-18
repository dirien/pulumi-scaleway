// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates and manages Scaleway Domain zone.\
// For more information, see [the documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/configure-dns-zones/).
//
// ## Examples
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
//			_, err := scaleway.NewDomainZone(ctx, "test", &scaleway.DomainZoneArgs{
//				Domain:    pulumi.String("scaleway-terraform.com"),
//				Subdomain: pulumi.String("test"),
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
// Zone can be imported using the `{subdomain}.{domain}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/domainZone:DomainZone test test.scaleway-terraform.com
//
// ```
type DomainZone struct {
	pulumi.CustomResourceState

	// The domain where the DNS zone will be created.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Message
	Message pulumi.StringOutput `pulumi:"message"`
	// NameServer list for zone.
	Ns pulumi.StringArrayOutput `pulumi:"ns"`
	// NameServer default list for zone.
	NsDefaults pulumi.StringArrayOutput `pulumi:"nsDefaults"`
	// NameServer master list for zone.
	NsMasters pulumi.StringArrayOutput `pulumi:"nsMasters"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The domain zone status.
	Status pulumi.StringOutput `pulumi:"status"`
	// The subdomain(zone name) to create in the domain.
	Subdomain pulumi.StringOutput `pulumi:"subdomain"`
	// The date and time of the last update of the DNS zone.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewDomainZone registers a new resource with the given unique name, arguments, and options.
func NewDomainZone(ctx *pulumi.Context,
	name string, args *DomainZoneArgs, opts ...pulumi.ResourceOption) (*DomainZone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Subdomain == nil {
		return nil, errors.New("invalid value for required argument 'Subdomain'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DomainZone
	err := ctx.RegisterResource("scaleway:index/domainZone:DomainZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainZone gets an existing DomainZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainZoneState, opts ...pulumi.ResourceOption) (*DomainZone, error) {
	var resource DomainZone
	err := ctx.ReadResource("scaleway:index/domainZone:DomainZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainZone resources.
type domainZoneState struct {
	// The domain where the DNS zone will be created.
	Domain *string `pulumi:"domain"`
	// Message
	Message *string `pulumi:"message"`
	// NameServer list for zone.
	Ns []string `pulumi:"ns"`
	// NameServer default list for zone.
	NsDefaults []string `pulumi:"nsDefaults"`
	// NameServer master list for zone.
	NsMasters []string `pulumi:"nsMasters"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The domain zone status.
	Status *string `pulumi:"status"`
	// The subdomain(zone name) to create in the domain.
	Subdomain *string `pulumi:"subdomain"`
	// The date and time of the last update of the DNS zone.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type DomainZoneState struct {
	// The domain where the DNS zone will be created.
	Domain pulumi.StringPtrInput
	// Message
	Message pulumi.StringPtrInput
	// NameServer list for zone.
	Ns pulumi.StringArrayInput
	// NameServer default list for zone.
	NsDefaults pulumi.StringArrayInput
	// NameServer master list for zone.
	NsMasters pulumi.StringArrayInput
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput
	// The domain zone status.
	Status pulumi.StringPtrInput
	// The subdomain(zone name) to create in the domain.
	Subdomain pulumi.StringPtrInput
	// The date and time of the last update of the DNS zone.
	UpdatedAt pulumi.StringPtrInput
}

func (DomainZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneState)(nil)).Elem()
}

type domainZoneArgs struct {
	// The domain where the DNS zone will be created.
	Domain string `pulumi:"domain"`
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId *string `pulumi:"projectId"`
	// The subdomain(zone name) to create in the domain.
	Subdomain string `pulumi:"subdomain"`
}

// The set of arguments for constructing a DomainZone resource.
type DomainZoneArgs struct {
	// The domain where the DNS zone will be created.
	Domain pulumi.StringInput
	// `projectId`) The ID of the project the domain is associated with.
	ProjectId pulumi.StringPtrInput
	// The subdomain(zone name) to create in the domain.
	Subdomain pulumi.StringInput
}

func (DomainZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneArgs)(nil)).Elem()
}

type DomainZoneInput interface {
	pulumi.Input

	ToDomainZoneOutput() DomainZoneOutput
	ToDomainZoneOutputWithContext(ctx context.Context) DomainZoneOutput
}

func (*DomainZone) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZone)(nil)).Elem()
}

func (i *DomainZone) ToDomainZoneOutput() DomainZoneOutput {
	return i.ToDomainZoneOutputWithContext(context.Background())
}

func (i *DomainZone) ToDomainZoneOutputWithContext(ctx context.Context) DomainZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneOutput)
}

func (i *DomainZone) ToOutput(ctx context.Context) pulumix.Output[*DomainZone] {
	return pulumix.Output[*DomainZone]{
		OutputState: i.ToDomainZoneOutputWithContext(ctx).OutputState,
	}
}

// DomainZoneArrayInput is an input type that accepts DomainZoneArray and DomainZoneArrayOutput values.
// You can construct a concrete instance of `DomainZoneArrayInput` via:
//
//	DomainZoneArray{ DomainZoneArgs{...} }
type DomainZoneArrayInput interface {
	pulumi.Input

	ToDomainZoneArrayOutput() DomainZoneArrayOutput
	ToDomainZoneArrayOutputWithContext(context.Context) DomainZoneArrayOutput
}

type DomainZoneArray []DomainZoneInput

func (DomainZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZone)(nil)).Elem()
}

func (i DomainZoneArray) ToDomainZoneArrayOutput() DomainZoneArrayOutput {
	return i.ToDomainZoneArrayOutputWithContext(context.Background())
}

func (i DomainZoneArray) ToDomainZoneArrayOutputWithContext(ctx context.Context) DomainZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneArrayOutput)
}

func (i DomainZoneArray) ToOutput(ctx context.Context) pulumix.Output[[]*DomainZone] {
	return pulumix.Output[[]*DomainZone]{
		OutputState: i.ToDomainZoneArrayOutputWithContext(ctx).OutputState,
	}
}

// DomainZoneMapInput is an input type that accepts DomainZoneMap and DomainZoneMapOutput values.
// You can construct a concrete instance of `DomainZoneMapInput` via:
//
//	DomainZoneMap{ "key": DomainZoneArgs{...} }
type DomainZoneMapInput interface {
	pulumi.Input

	ToDomainZoneMapOutput() DomainZoneMapOutput
	ToDomainZoneMapOutputWithContext(context.Context) DomainZoneMapOutput
}

type DomainZoneMap map[string]DomainZoneInput

func (DomainZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZone)(nil)).Elem()
}

func (i DomainZoneMap) ToDomainZoneMapOutput() DomainZoneMapOutput {
	return i.ToDomainZoneMapOutputWithContext(context.Background())
}

func (i DomainZoneMap) ToDomainZoneMapOutputWithContext(ctx context.Context) DomainZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneMapOutput)
}

func (i DomainZoneMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DomainZone] {
	return pulumix.Output[map[string]*DomainZone]{
		OutputState: i.ToDomainZoneMapOutputWithContext(ctx).OutputState,
	}
}

type DomainZoneOutput struct{ *pulumi.OutputState }

func (DomainZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZone)(nil)).Elem()
}

func (o DomainZoneOutput) ToDomainZoneOutput() DomainZoneOutput {
	return o
}

func (o DomainZoneOutput) ToDomainZoneOutputWithContext(ctx context.Context) DomainZoneOutput {
	return o
}

func (o DomainZoneOutput) ToOutput(ctx context.Context) pulumix.Output[*DomainZone] {
	return pulumix.Output[*DomainZone]{
		OutputState: o.OutputState,
	}
}

// The domain where the DNS zone will be created.
func (o DomainZoneOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Message
func (o DomainZoneOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// NameServer list for zone.
func (o DomainZoneOutput) Ns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringArrayOutput { return v.Ns }).(pulumi.StringArrayOutput)
}

// NameServer default list for zone.
func (o DomainZoneOutput) NsDefaults() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringArrayOutput { return v.NsDefaults }).(pulumi.StringArrayOutput)
}

// NameServer master list for zone.
func (o DomainZoneOutput) NsMasters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringArrayOutput { return v.NsMasters }).(pulumi.StringArrayOutput)
}

// `projectId`) The ID of the project the domain is associated with.
func (o DomainZoneOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The domain zone status.
func (o DomainZoneOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The subdomain(zone name) to create in the domain.
func (o DomainZoneOutput) Subdomain() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.Subdomain }).(pulumi.StringOutput)
}

// The date and time of the last update of the DNS zone.
func (o DomainZoneOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type DomainZoneArrayOutput struct{ *pulumi.OutputState }

func (DomainZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZone)(nil)).Elem()
}

func (o DomainZoneArrayOutput) ToDomainZoneArrayOutput() DomainZoneArrayOutput {
	return o
}

func (o DomainZoneArrayOutput) ToDomainZoneArrayOutputWithContext(ctx context.Context) DomainZoneArrayOutput {
	return o
}

func (o DomainZoneArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DomainZone] {
	return pulumix.Output[[]*DomainZone]{
		OutputState: o.OutputState,
	}
}

func (o DomainZoneArrayOutput) Index(i pulumi.IntInput) DomainZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainZone {
		return vs[0].([]*DomainZone)[vs[1].(int)]
	}).(DomainZoneOutput)
}

type DomainZoneMapOutput struct{ *pulumi.OutputState }

func (DomainZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZone)(nil)).Elem()
}

func (o DomainZoneMapOutput) ToDomainZoneMapOutput() DomainZoneMapOutput {
	return o
}

func (o DomainZoneMapOutput) ToDomainZoneMapOutputWithContext(ctx context.Context) DomainZoneMapOutput {
	return o
}

func (o DomainZoneMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DomainZone] {
	return pulumix.Output[map[string]*DomainZone]{
		OutputState: o.OutputState,
	}
}

func (o DomainZoneMapOutput) MapIndex(k pulumi.StringInput) DomainZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainZone {
		return vs[0].(map[string]*DomainZone)[vs[1].(string)]
	}).(DomainZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneInput)(nil)).Elem(), &DomainZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneArrayInput)(nil)).Elem(), DomainZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneMapInput)(nil)).Elem(), DomainZoneMap{})
	pulumi.RegisterOutputType(DomainZoneOutput{})
	pulumi.RegisterOutputType(DomainZoneArrayOutput{})
	pulumi.RegisterOutputType(DomainZoneMapOutput{})
}
