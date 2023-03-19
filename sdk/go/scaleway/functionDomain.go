// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Function Domain bindings.
// For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/).
//
// ## Examples
//
// ### Basic
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
//			mainFunctionNamespace, err := scaleway.NewFunctionNamespace(ctx, "mainFunctionNamespace", nil)
//			if err != nil {
//				return err
//			}
//			mainFunction, err := scaleway.NewFunction(ctx, "mainFunction", &scaleway.FunctionArgs{
//				NamespaceId: mainFunctionNamespace.ID(),
//				Runtime:     pulumi.String("go118"),
//				Privacy:     pulumi.String("private"),
//				Handler:     pulumi.String("Handle"),
//				ZipFile:     pulumi.String("testfixture/gofunction.zip"),
//				Deploy:      pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewFunctionDomain(ctx, "mainFunctionDomain", &scaleway.FunctionDomainArgs{
//				FunctionId: mainFunction.ID(),
//				Hostname:   pulumi.String("example.com"),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				mainFunction,
//			}))
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
// Domain can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/functionDomain:FunctionDomain main fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type FunctionDomain struct {
	pulumi.CustomResourceState

	// The ID of the function you want to create a domain with.
	FunctionId pulumi.StringOutput `pulumi:"functionId"`
	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function `domainName` for it.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// (Defaults to provider `region`) The region in where the domain was created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URL that triggers the function
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewFunctionDomain registers a new resource with the given unique name, arguments, and options.
func NewFunctionDomain(ctx *pulumi.Context,
	name string, args *FunctionDomainArgs, opts ...pulumi.ResourceOption) (*FunctionDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionId == nil {
		return nil, errors.New("invalid value for required argument 'FunctionId'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource FunctionDomain
	err := ctx.RegisterResource("scaleway:index/functionDomain:FunctionDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionDomain gets an existing FunctionDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionDomainState, opts ...pulumi.ResourceOption) (*FunctionDomain, error) {
	var resource FunctionDomain
	err := ctx.ReadResource("scaleway:index/functionDomain:FunctionDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionDomain resources.
type functionDomainState struct {
	// The ID of the function you want to create a domain with.
	FunctionId *string `pulumi:"functionId"`
	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function `domainName` for it.
	Hostname *string `pulumi:"hostname"`
	// (Defaults to provider `region`) The region in where the domain was created.
	Region *string `pulumi:"region"`
	// The URL that triggers the function
	Url *string `pulumi:"url"`
}

type FunctionDomainState struct {
	// The ID of the function you want to create a domain with.
	FunctionId pulumi.StringPtrInput
	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function `domainName` for it.
	Hostname pulumi.StringPtrInput
	// (Defaults to provider `region`) The region in where the domain was created.
	Region pulumi.StringPtrInput
	// The URL that triggers the function
	Url pulumi.StringPtrInput
}

func (FunctionDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDomainState)(nil)).Elem()
}

type functionDomainArgs struct {
	// The ID of the function you want to create a domain with.
	FunctionId string `pulumi:"functionId"`
	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function `domainName` for it.
	Hostname string `pulumi:"hostname"`
	// (Defaults to provider `region`) The region in where the domain was created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FunctionDomain resource.
type FunctionDomainArgs struct {
	// The ID of the function you want to create a domain with.
	FunctionId pulumi.StringInput
	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function `domainName` for it.
	Hostname pulumi.StringInput
	// (Defaults to provider `region`) The region in where the domain was created.
	Region pulumi.StringPtrInput
}

func (FunctionDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDomainArgs)(nil)).Elem()
}

type FunctionDomainInput interface {
	pulumi.Input

	ToFunctionDomainOutput() FunctionDomainOutput
	ToFunctionDomainOutputWithContext(ctx context.Context) FunctionDomainOutput
}

func (*FunctionDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionDomain)(nil)).Elem()
}

func (i *FunctionDomain) ToFunctionDomainOutput() FunctionDomainOutput {
	return i.ToFunctionDomainOutputWithContext(context.Background())
}

func (i *FunctionDomain) ToFunctionDomainOutputWithContext(ctx context.Context) FunctionDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionDomainOutput)
}

// FunctionDomainArrayInput is an input type that accepts FunctionDomainArray and FunctionDomainArrayOutput values.
// You can construct a concrete instance of `FunctionDomainArrayInput` via:
//
//	FunctionDomainArray{ FunctionDomainArgs{...} }
type FunctionDomainArrayInput interface {
	pulumi.Input

	ToFunctionDomainArrayOutput() FunctionDomainArrayOutput
	ToFunctionDomainArrayOutputWithContext(context.Context) FunctionDomainArrayOutput
}

type FunctionDomainArray []FunctionDomainInput

func (FunctionDomainArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionDomain)(nil)).Elem()
}

func (i FunctionDomainArray) ToFunctionDomainArrayOutput() FunctionDomainArrayOutput {
	return i.ToFunctionDomainArrayOutputWithContext(context.Background())
}

func (i FunctionDomainArray) ToFunctionDomainArrayOutputWithContext(ctx context.Context) FunctionDomainArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionDomainArrayOutput)
}

// FunctionDomainMapInput is an input type that accepts FunctionDomainMap and FunctionDomainMapOutput values.
// You can construct a concrete instance of `FunctionDomainMapInput` via:
//
//	FunctionDomainMap{ "key": FunctionDomainArgs{...} }
type FunctionDomainMapInput interface {
	pulumi.Input

	ToFunctionDomainMapOutput() FunctionDomainMapOutput
	ToFunctionDomainMapOutputWithContext(context.Context) FunctionDomainMapOutput
}

type FunctionDomainMap map[string]FunctionDomainInput

func (FunctionDomainMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionDomain)(nil)).Elem()
}

func (i FunctionDomainMap) ToFunctionDomainMapOutput() FunctionDomainMapOutput {
	return i.ToFunctionDomainMapOutputWithContext(context.Background())
}

func (i FunctionDomainMap) ToFunctionDomainMapOutputWithContext(ctx context.Context) FunctionDomainMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionDomainMapOutput)
}

type FunctionDomainOutput struct{ *pulumi.OutputState }

func (FunctionDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionDomain)(nil)).Elem()
}

func (o FunctionDomainOutput) ToFunctionDomainOutput() FunctionDomainOutput {
	return o
}

func (o FunctionDomainOutput) ToFunctionDomainOutputWithContext(ctx context.Context) FunctionDomainOutput {
	return o
}

// The ID of the function you want to create a domain with.
func (o FunctionDomainOutput) FunctionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionDomain) pulumi.StringOutput { return v.FunctionId }).(pulumi.StringOutput)
}

// The hostname that should resolve to your function id native domain.
// You should use a CNAME domain record that point to your native function `domainName` for it.
func (o FunctionDomainOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionDomain) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// (Defaults to provider `region`) The region in where the domain was created.
func (o FunctionDomainOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionDomain) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The URL that triggers the function
func (o FunctionDomainOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionDomain) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type FunctionDomainArrayOutput struct{ *pulumi.OutputState }

func (FunctionDomainArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionDomain)(nil)).Elem()
}

func (o FunctionDomainArrayOutput) ToFunctionDomainArrayOutput() FunctionDomainArrayOutput {
	return o
}

func (o FunctionDomainArrayOutput) ToFunctionDomainArrayOutputWithContext(ctx context.Context) FunctionDomainArrayOutput {
	return o
}

func (o FunctionDomainArrayOutput) Index(i pulumi.IntInput) FunctionDomainOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionDomain {
		return vs[0].([]*FunctionDomain)[vs[1].(int)]
	}).(FunctionDomainOutput)
}

type FunctionDomainMapOutput struct{ *pulumi.OutputState }

func (FunctionDomainMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionDomain)(nil)).Elem()
}

func (o FunctionDomainMapOutput) ToFunctionDomainMapOutput() FunctionDomainMapOutput {
	return o
}

func (o FunctionDomainMapOutput) ToFunctionDomainMapOutputWithContext(ctx context.Context) FunctionDomainMapOutput {
	return o
}

func (o FunctionDomainMapOutput) MapIndex(k pulumi.StringInput) FunctionDomainOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionDomain {
		return vs[0].(map[string]*FunctionDomain)[vs[1].(string)]
	}).(FunctionDomainOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionDomainInput)(nil)).Elem(), &FunctionDomain{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionDomainArrayInput)(nil)).Elem(), FunctionDomainArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionDomainMapInput)(nil)).Elem(), FunctionDomainMap{})
	pulumi.RegisterOutputType(FunctionDomainOutput{})
	pulumi.RegisterOutputType(FunctionDomainArrayOutput{})
	pulumi.RegisterOutputType(FunctionDomainMapOutput{})
}
