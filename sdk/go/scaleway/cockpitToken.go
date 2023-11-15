// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Cockpit Tokens.
//
// For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#tokens).
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
//			mainCockpit, err := scaleway.LookupCockpit(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewCockpitToken(ctx, "mainCockpitToken", &scaleway.CockpitTokenArgs{
//				ProjectId: *pulumi.String(mainCockpit.ProjectId),
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
//			mainCockpit, err := scaleway.LookupCockpit(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewCockpitToken(ctx, "mainCockpitToken", &scaleway.CockpitTokenArgs{
//				ProjectId: *pulumi.String(mainCockpit.ProjectId),
//				Scopes: &scaleway.CockpitTokenScopesArgs{
//					QueryMetrics: pulumi.Bool(true),
//					WriteMetrics: pulumi.Bool(false),
//					QueryLogs:    pulumi.Bool(true),
//					WriteLogs:    pulumi.Bool(false),
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
// Cockpits can be imported using the token ID, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/cockpitToken:CockpitToken main 11111111-1111-1111-1111-111111111111
//
// ```
type CockpitToken struct {
	pulumi.CustomResourceState

	// The name of the token.
	Name pulumi.StringOutput `pulumi:"name"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Allowed scopes.
	Scopes CockpitTokenScopesOutput `pulumi:"scopes"`
	// The secret key of the token.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
}

// NewCockpitToken registers a new resource with the given unique name, arguments, and options.
func NewCockpitToken(ctx *pulumi.Context,
	name string, args *CockpitTokenArgs, opts ...pulumi.ResourceOption) (*CockpitToken, error) {
	if args == nil {
		args = &CockpitTokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretKey",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CockpitToken
	err := ctx.RegisterResource("scaleway:index/cockpitToken:CockpitToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCockpitToken gets an existing CockpitToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCockpitToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CockpitTokenState, opts ...pulumi.ResourceOption) (*CockpitToken, error) {
	var resource CockpitToken
	err := ctx.ReadResource("scaleway:index/cockpitToken:CockpitToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CockpitToken resources.
type cockpitTokenState struct {
	// The name of the token.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Allowed scopes.
	Scopes *CockpitTokenScopes `pulumi:"scopes"`
	// The secret key of the token.
	SecretKey *string `pulumi:"secretKey"`
}

type CockpitTokenState struct {
	// The name of the token.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// Allowed scopes.
	Scopes CockpitTokenScopesPtrInput
	// The secret key of the token.
	SecretKey pulumi.StringPtrInput
}

func (CockpitTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitTokenState)(nil)).Elem()
}

type cockpitTokenArgs struct {
	// The name of the token.
	Name *string `pulumi:"name"`
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Allowed scopes.
	Scopes *CockpitTokenScopes `pulumi:"scopes"`
}

// The set of arguments for constructing a CockpitToken resource.
type CockpitTokenArgs struct {
	// The name of the token.
	Name pulumi.StringPtrInput
	// `projectId`) The ID of the project the cockpit is associated with.
	ProjectId pulumi.StringPtrInput
	// Allowed scopes.
	Scopes CockpitTokenScopesPtrInput
}

func (CockpitTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cockpitTokenArgs)(nil)).Elem()
}

type CockpitTokenInput interface {
	pulumi.Input

	ToCockpitTokenOutput() CockpitTokenOutput
	ToCockpitTokenOutputWithContext(ctx context.Context) CockpitTokenOutput
}

func (*CockpitToken) ElementType() reflect.Type {
	return reflect.TypeOf((**CockpitToken)(nil)).Elem()
}

func (i *CockpitToken) ToCockpitTokenOutput() CockpitTokenOutput {
	return i.ToCockpitTokenOutputWithContext(context.Background())
}

func (i *CockpitToken) ToCockpitTokenOutputWithContext(ctx context.Context) CockpitTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitTokenOutput)
}

// CockpitTokenArrayInput is an input type that accepts CockpitTokenArray and CockpitTokenArrayOutput values.
// You can construct a concrete instance of `CockpitTokenArrayInput` via:
//
//	CockpitTokenArray{ CockpitTokenArgs{...} }
type CockpitTokenArrayInput interface {
	pulumi.Input

	ToCockpitTokenArrayOutput() CockpitTokenArrayOutput
	ToCockpitTokenArrayOutputWithContext(context.Context) CockpitTokenArrayOutput
}

type CockpitTokenArray []CockpitTokenInput

func (CockpitTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CockpitToken)(nil)).Elem()
}

func (i CockpitTokenArray) ToCockpitTokenArrayOutput() CockpitTokenArrayOutput {
	return i.ToCockpitTokenArrayOutputWithContext(context.Background())
}

func (i CockpitTokenArray) ToCockpitTokenArrayOutputWithContext(ctx context.Context) CockpitTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitTokenArrayOutput)
}

// CockpitTokenMapInput is an input type that accepts CockpitTokenMap and CockpitTokenMapOutput values.
// You can construct a concrete instance of `CockpitTokenMapInput` via:
//
//	CockpitTokenMap{ "key": CockpitTokenArgs{...} }
type CockpitTokenMapInput interface {
	pulumi.Input

	ToCockpitTokenMapOutput() CockpitTokenMapOutput
	ToCockpitTokenMapOutputWithContext(context.Context) CockpitTokenMapOutput
}

type CockpitTokenMap map[string]CockpitTokenInput

func (CockpitTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CockpitToken)(nil)).Elem()
}

func (i CockpitTokenMap) ToCockpitTokenMapOutput() CockpitTokenMapOutput {
	return i.ToCockpitTokenMapOutputWithContext(context.Background())
}

func (i CockpitTokenMap) ToCockpitTokenMapOutputWithContext(ctx context.Context) CockpitTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CockpitTokenMapOutput)
}

type CockpitTokenOutput struct{ *pulumi.OutputState }

func (CockpitTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CockpitToken)(nil)).Elem()
}

func (o CockpitTokenOutput) ToCockpitTokenOutput() CockpitTokenOutput {
	return o
}

func (o CockpitTokenOutput) ToCockpitTokenOutputWithContext(ctx context.Context) CockpitTokenOutput {
	return o
}

// The name of the token.
func (o CockpitTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// `projectId`) The ID of the project the cockpit is associated with.
func (o CockpitTokenOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Allowed scopes.
func (o CockpitTokenOutput) Scopes() CockpitTokenScopesOutput {
	return o.ApplyT(func(v *CockpitToken) CockpitTokenScopesOutput { return v.Scopes }).(CockpitTokenScopesOutput)
}

// The secret key of the token.
func (o CockpitTokenOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v *CockpitToken) pulumi.StringOutput { return v.SecretKey }).(pulumi.StringOutput)
}

type CockpitTokenArrayOutput struct{ *pulumi.OutputState }

func (CockpitTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CockpitToken)(nil)).Elem()
}

func (o CockpitTokenArrayOutput) ToCockpitTokenArrayOutput() CockpitTokenArrayOutput {
	return o
}

func (o CockpitTokenArrayOutput) ToCockpitTokenArrayOutputWithContext(ctx context.Context) CockpitTokenArrayOutput {
	return o
}

func (o CockpitTokenArrayOutput) Index(i pulumi.IntInput) CockpitTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CockpitToken {
		return vs[0].([]*CockpitToken)[vs[1].(int)]
	}).(CockpitTokenOutput)
}

type CockpitTokenMapOutput struct{ *pulumi.OutputState }

func (CockpitTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CockpitToken)(nil)).Elem()
}

func (o CockpitTokenMapOutput) ToCockpitTokenMapOutput() CockpitTokenMapOutput {
	return o
}

func (o CockpitTokenMapOutput) ToCockpitTokenMapOutputWithContext(ctx context.Context) CockpitTokenMapOutput {
	return o
}

func (o CockpitTokenMapOutput) MapIndex(k pulumi.StringInput) CockpitTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CockpitToken {
		return vs[0].(map[string]*CockpitToken)[vs[1].(string)]
	}).(CockpitTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitTokenInput)(nil)).Elem(), &CockpitToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitTokenArrayInput)(nil)).Elem(), CockpitTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CockpitTokenMapInput)(nil)).Elem(), CockpitTokenMap{})
	pulumi.RegisterOutputType(CockpitTokenOutput{})
	pulumi.RegisterOutputType(CockpitTokenArrayOutput{})
	pulumi.RegisterOutputType(CockpitTokenMapOutput{})
}
