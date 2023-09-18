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

type LbCertificate struct {
	pulumi.CustomResourceState

	// Main domain of the certificate. A new certificate will be created if this field is changed.
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
	CustomCertificate LbCertificateCustomCertificatePtrOutput `pulumi:"customCertificate"`
	// The identifier (SHA-1) of the certificate
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The load-balancer ID this certificate is attached to.
	//
	// > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
	LbId pulumi.StringOutput `pulumi:"lbId"`
	// Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
	Letsencrypt LbCertificateLetsencryptPtrOutput `pulumi:"letsencrypt"`
	// The name of the certificate backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// The not valid after validity bound timestamp
	NotValidAfter pulumi.StringOutput `pulumi:"notValidAfter"`
	// The not valid before validity bound timestamp
	NotValidBefore pulumi.StringOutput `pulumi:"notValidBefore"`
	// Certificate status
	Status pulumi.StringOutput `pulumi:"status"`
	// Array of alternative domain names.  A new certificate will be created if this field is changed.
	//
	// > **Important:** Updates to `letsencrypt` will recreate the load-balancer certificate.
	SubjectAlternativeNames pulumi.StringArrayOutput `pulumi:"subjectAlternativeNames"`
}

// NewLbCertificate registers a new resource with the given unique name, arguments, and options.
func NewLbCertificate(ctx *pulumi.Context,
	name string, args *LbCertificateArgs, opts ...pulumi.ResourceOption) (*LbCertificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LbId == nil {
		return nil, errors.New("invalid value for required argument 'LbId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbCertificate
	err := ctx.RegisterResource("scaleway:index/lbCertificate:LbCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbCertificate gets an existing LbCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbCertificateState, opts ...pulumi.ResourceOption) (*LbCertificate, error) {
	var resource LbCertificate
	err := ctx.ReadResource("scaleway:index/lbCertificate:LbCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbCertificate resources.
type lbCertificateState struct {
	// Main domain of the certificate. A new certificate will be created if this field is changed.
	CommonName *string `pulumi:"commonName"`
	// Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
	CustomCertificate *LbCertificateCustomCertificate `pulumi:"customCertificate"`
	// The identifier (SHA-1) of the certificate
	Fingerprint *string `pulumi:"fingerprint"`
	// The load-balancer ID this certificate is attached to.
	//
	// > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
	LbId *string `pulumi:"lbId"`
	// Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
	Letsencrypt *LbCertificateLetsencrypt `pulumi:"letsencrypt"`
	// The name of the certificate backend.
	Name *string `pulumi:"name"`
	// The not valid after validity bound timestamp
	NotValidAfter *string `pulumi:"notValidAfter"`
	// The not valid before validity bound timestamp
	NotValidBefore *string `pulumi:"notValidBefore"`
	// Certificate status
	Status *string `pulumi:"status"`
	// Array of alternative domain names.  A new certificate will be created if this field is changed.
	//
	// > **Important:** Updates to `letsencrypt` will recreate the load-balancer certificate.
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
}

type LbCertificateState struct {
	// Main domain of the certificate. A new certificate will be created if this field is changed.
	CommonName pulumi.StringPtrInput
	// Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
	CustomCertificate LbCertificateCustomCertificatePtrInput
	// The identifier (SHA-1) of the certificate
	Fingerprint pulumi.StringPtrInput
	// The load-balancer ID this certificate is attached to.
	//
	// > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
	LbId pulumi.StringPtrInput
	// Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
	Letsencrypt LbCertificateLetsencryptPtrInput
	// The name of the certificate backend.
	Name pulumi.StringPtrInput
	// The not valid after validity bound timestamp
	NotValidAfter pulumi.StringPtrInput
	// The not valid before validity bound timestamp
	NotValidBefore pulumi.StringPtrInput
	// Certificate status
	Status pulumi.StringPtrInput
	// Array of alternative domain names.  A new certificate will be created if this field is changed.
	//
	// > **Important:** Updates to `letsencrypt` will recreate the load-balancer certificate.
	SubjectAlternativeNames pulumi.StringArrayInput
}

func (LbCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbCertificateState)(nil)).Elem()
}

type lbCertificateArgs struct {
	// Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
	CustomCertificate *LbCertificateCustomCertificate `pulumi:"customCertificate"`
	// The load-balancer ID this certificate is attached to.
	//
	// > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
	LbId string `pulumi:"lbId"`
	// Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
	Letsencrypt *LbCertificateLetsencrypt `pulumi:"letsencrypt"`
	// The name of the certificate backend.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a LbCertificate resource.
type LbCertificateArgs struct {
	// Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
	CustomCertificate LbCertificateCustomCertificatePtrInput
	// The load-balancer ID this certificate is attached to.
	//
	// > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
	LbId pulumi.StringInput
	// Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
	Letsencrypt LbCertificateLetsencryptPtrInput
	// The name of the certificate backend.
	Name pulumi.StringPtrInput
}

func (LbCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbCertificateArgs)(nil)).Elem()
}

type LbCertificateInput interface {
	pulumi.Input

	ToLbCertificateOutput() LbCertificateOutput
	ToLbCertificateOutputWithContext(ctx context.Context) LbCertificateOutput
}

func (*LbCertificate) ElementType() reflect.Type {
	return reflect.TypeOf((**LbCertificate)(nil)).Elem()
}

func (i *LbCertificate) ToLbCertificateOutput() LbCertificateOutput {
	return i.ToLbCertificateOutputWithContext(context.Background())
}

func (i *LbCertificate) ToLbCertificateOutputWithContext(ctx context.Context) LbCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbCertificateOutput)
}

func (i *LbCertificate) ToOutput(ctx context.Context) pulumix.Output[*LbCertificate] {
	return pulumix.Output[*LbCertificate]{
		OutputState: i.ToLbCertificateOutputWithContext(ctx).OutputState,
	}
}

// LbCertificateArrayInput is an input type that accepts LbCertificateArray and LbCertificateArrayOutput values.
// You can construct a concrete instance of `LbCertificateArrayInput` via:
//
//	LbCertificateArray{ LbCertificateArgs{...} }
type LbCertificateArrayInput interface {
	pulumi.Input

	ToLbCertificateArrayOutput() LbCertificateArrayOutput
	ToLbCertificateArrayOutputWithContext(context.Context) LbCertificateArrayOutput
}

type LbCertificateArray []LbCertificateInput

func (LbCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbCertificate)(nil)).Elem()
}

func (i LbCertificateArray) ToLbCertificateArrayOutput() LbCertificateArrayOutput {
	return i.ToLbCertificateArrayOutputWithContext(context.Background())
}

func (i LbCertificateArray) ToLbCertificateArrayOutputWithContext(ctx context.Context) LbCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbCertificateArrayOutput)
}

func (i LbCertificateArray) ToOutput(ctx context.Context) pulumix.Output[[]*LbCertificate] {
	return pulumix.Output[[]*LbCertificate]{
		OutputState: i.ToLbCertificateArrayOutputWithContext(ctx).OutputState,
	}
}

// LbCertificateMapInput is an input type that accepts LbCertificateMap and LbCertificateMapOutput values.
// You can construct a concrete instance of `LbCertificateMapInput` via:
//
//	LbCertificateMap{ "key": LbCertificateArgs{...} }
type LbCertificateMapInput interface {
	pulumi.Input

	ToLbCertificateMapOutput() LbCertificateMapOutput
	ToLbCertificateMapOutputWithContext(context.Context) LbCertificateMapOutput
}

type LbCertificateMap map[string]LbCertificateInput

func (LbCertificateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbCertificate)(nil)).Elem()
}

func (i LbCertificateMap) ToLbCertificateMapOutput() LbCertificateMapOutput {
	return i.ToLbCertificateMapOutputWithContext(context.Background())
}

func (i LbCertificateMap) ToLbCertificateMapOutputWithContext(ctx context.Context) LbCertificateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbCertificateMapOutput)
}

func (i LbCertificateMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LbCertificate] {
	return pulumix.Output[map[string]*LbCertificate]{
		OutputState: i.ToLbCertificateMapOutputWithContext(ctx).OutputState,
	}
}

type LbCertificateOutput struct{ *pulumi.OutputState }

func (LbCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbCertificate)(nil)).Elem()
}

func (o LbCertificateOutput) ToLbCertificateOutput() LbCertificateOutput {
	return o
}

func (o LbCertificateOutput) ToLbCertificateOutputWithContext(ctx context.Context) LbCertificateOutput {
	return o
}

func (o LbCertificateOutput) ToOutput(ctx context.Context) pulumix.Output[*LbCertificate] {
	return pulumix.Output[*LbCertificate]{
		OutputState: o.OutputState,
	}
}

// Main domain of the certificate. A new certificate will be created if this field is changed.
func (o LbCertificateOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.CommonName }).(pulumi.StringOutput)
}

// Configuration block for custom certificate chain. Only one of `letsencrypt` and `customCertificate` should be specified.
func (o LbCertificateOutput) CustomCertificate() LbCertificateCustomCertificatePtrOutput {
	return o.ApplyT(func(v *LbCertificate) LbCertificateCustomCertificatePtrOutput { return v.CustomCertificate }).(LbCertificateCustomCertificatePtrOutput)
}

// The identifier (SHA-1) of the certificate
func (o LbCertificateOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The load-balancer ID this certificate is attached to.
//
// > **Important:** Updates to `lbId` will recreate the load-balancer certificate.
func (o LbCertificateOutput) LbId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.LbId }).(pulumi.StringOutput)
}

// Configuration block for Let's Encrypt configuration. Only one of `letsencrypt` and `customCertificate` should be specified.
func (o LbCertificateOutput) Letsencrypt() LbCertificateLetsencryptPtrOutput {
	return o.ApplyT(func(v *LbCertificate) LbCertificateLetsencryptPtrOutput { return v.Letsencrypt }).(LbCertificateLetsencryptPtrOutput)
}

// The name of the certificate backend.
func (o LbCertificateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The not valid after validity bound timestamp
func (o LbCertificateOutput) NotValidAfter() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.NotValidAfter }).(pulumi.StringOutput)
}

// The not valid before validity bound timestamp
func (o LbCertificateOutput) NotValidBefore() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.NotValidBefore }).(pulumi.StringOutput)
}

// Certificate status
func (o LbCertificateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Array of alternative domain names.  A new certificate will be created if this field is changed.
//
// > **Important:** Updates to `letsencrypt` will recreate the load-balancer certificate.
func (o LbCertificateOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LbCertificate) pulumi.StringArrayOutput { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

type LbCertificateArrayOutput struct{ *pulumi.OutputState }

func (LbCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbCertificate)(nil)).Elem()
}

func (o LbCertificateArrayOutput) ToLbCertificateArrayOutput() LbCertificateArrayOutput {
	return o
}

func (o LbCertificateArrayOutput) ToLbCertificateArrayOutputWithContext(ctx context.Context) LbCertificateArrayOutput {
	return o
}

func (o LbCertificateArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LbCertificate] {
	return pulumix.Output[[]*LbCertificate]{
		OutputState: o.OutputState,
	}
}

func (o LbCertificateArrayOutput) Index(i pulumi.IntInput) LbCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbCertificate {
		return vs[0].([]*LbCertificate)[vs[1].(int)]
	}).(LbCertificateOutput)
}

type LbCertificateMapOutput struct{ *pulumi.OutputState }

func (LbCertificateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbCertificate)(nil)).Elem()
}

func (o LbCertificateMapOutput) ToLbCertificateMapOutput() LbCertificateMapOutput {
	return o
}

func (o LbCertificateMapOutput) ToLbCertificateMapOutputWithContext(ctx context.Context) LbCertificateMapOutput {
	return o
}

func (o LbCertificateMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LbCertificate] {
	return pulumix.Output[map[string]*LbCertificate]{
		OutputState: o.OutputState,
	}
}

func (o LbCertificateMapOutput) MapIndex(k pulumi.StringInput) LbCertificateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbCertificate {
		return vs[0].(map[string]*LbCertificate)[vs[1].(string)]
	}).(LbCertificateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbCertificateInput)(nil)).Elem(), &LbCertificate{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbCertificateArrayInput)(nil)).Elem(), LbCertificateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbCertificateMapInput)(nil)).Elem(), LbCertificateMap{})
	pulumi.RegisterOutputType(LbCertificateOutput{})
	pulumi.RegisterOutputType(LbCertificateArrayOutput{})
	pulumi.RegisterOutputType(LbCertificateMapOutput{})
}
