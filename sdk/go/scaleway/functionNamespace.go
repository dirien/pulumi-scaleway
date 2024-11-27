// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `FunctionNamespace` resource allows you to
// for Scaleway [Serverless Functions](https://www.scaleway.com/en/docs/serverless/functions/).
//
// Refer to the Functions namespace [documentation](https://www.scaleway.com/en/docs/serverless/functions/how-to/create-a-functions-namespace/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-namespaces-list-all-your-namespaces) for more information.
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
//			_, err := scaleway.NewFunctionNamespace(ctx, "main", &scaleway.FunctionNamespaceArgs{
//				Description: pulumi.String("Main function namespace"),
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
// Functions namespaces can be imported using `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/functionNamespace:FunctionNamespace main fr-par/11111111-1111-1111-1111-111111111111
// ```
type FunctionNamespace struct {
	pulumi.CustomResourceState

	// The description of the namespace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The environment variables of the namespace.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// The unique name of the Functions namespace.
	//
	// > **Important** Updates to the `name` argument will recreate the namespace.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Organization ID with which the namespace is associated.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// `projectId`) The unique identifier of the project that contains the namespace.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`). The region in which the namespace is created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint pulumi.StringOutput `pulumi:"registryEndpoint"`
	// The registry namespace ID of the namespace.
	RegistryNamespaceId pulumi.StringOutput `pulumi:"registryNamespaceId"`
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables pulumi.StringMapOutput `pulumi:"secretEnvironmentVariables"`
}

// NewFunctionNamespace registers a new resource with the given unique name, arguments, and options.
func NewFunctionNamespace(ctx *pulumi.Context,
	name string, args *FunctionNamespaceArgs, opts ...pulumi.ResourceOption) (*FunctionNamespace, error) {
	if args == nil {
		args = &FunctionNamespaceArgs{}
	}

	if args.SecretEnvironmentVariables != nil {
		args.SecretEnvironmentVariables = pulumi.ToSecret(args.SecretEnvironmentVariables).(pulumi.StringMapInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secretEnvironmentVariables",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FunctionNamespace
	err := ctx.RegisterResource("scaleway:index/functionNamespace:FunctionNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionNamespace gets an existing FunctionNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionNamespaceState, opts ...pulumi.ResourceOption) (*FunctionNamespace, error) {
	var resource FunctionNamespace
	err := ctx.ReadResource("scaleway:index/functionNamespace:FunctionNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionNamespace resources.
type functionNamespaceState struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The unique name of the Functions namespace.
	//
	// > **Important** Updates to the `name` argument will recreate the namespace.
	Name *string `pulumi:"name"`
	// The Organization ID with which the namespace is associated.
	OrganizationId *string `pulumi:"organizationId"`
	// `projectId`) The unique identifier of the project that contains the namespace.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace is created.
	Region *string `pulumi:"region"`
	// The registry endpoint of the namespace.
	RegistryEndpoint *string `pulumi:"registryEndpoint"`
	// The registry namespace ID of the namespace.
	RegistryNamespaceId *string `pulumi:"registryNamespaceId"`
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
}

type FunctionNamespaceState struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// The environment variables of the namespace.
	EnvironmentVariables pulumi.StringMapInput
	// The unique name of the Functions namespace.
	//
	// > **Important** Updates to the `name` argument will recreate the namespace.
	Name pulumi.StringPtrInput
	// The Organization ID with which the namespace is associated.
	OrganizationId pulumi.StringPtrInput
	// `projectId`) The unique identifier of the project that contains the namespace.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace is created.
	Region pulumi.StringPtrInput
	// The registry endpoint of the namespace.
	RegistryEndpoint pulumi.StringPtrInput
	// The registry namespace ID of the namespace.
	RegistryNamespaceId pulumi.StringPtrInput
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables pulumi.StringMapInput
}

func (FunctionNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionNamespaceState)(nil)).Elem()
}

type functionNamespaceArgs struct {
	// The description of the namespace.
	Description *string `pulumi:"description"`
	// The environment variables of the namespace.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// The unique name of the Functions namespace.
	//
	// > **Important** Updates to the `name` argument will recreate the namespace.
	Name *string `pulumi:"name"`
	// `projectId`) The unique identifier of the project that contains the namespace.
	ProjectId *string `pulumi:"projectId"`
	// `region`). The region in which the namespace is created.
	Region *string `pulumi:"region"`
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables map[string]string `pulumi:"secretEnvironmentVariables"`
}

// The set of arguments for constructing a FunctionNamespace resource.
type FunctionNamespaceArgs struct {
	// The description of the namespace.
	Description pulumi.StringPtrInput
	// The environment variables of the namespace.
	EnvironmentVariables pulumi.StringMapInput
	// The unique name of the Functions namespace.
	//
	// > **Important** Updates to the `name` argument will recreate the namespace.
	Name pulumi.StringPtrInput
	// `projectId`) The unique identifier of the project that contains the namespace.
	ProjectId pulumi.StringPtrInput
	// `region`). The region in which the namespace is created.
	Region pulumi.StringPtrInput
	// The secret environment variables of the namespace.
	SecretEnvironmentVariables pulumi.StringMapInput
}

func (FunctionNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionNamespaceArgs)(nil)).Elem()
}

type FunctionNamespaceInput interface {
	pulumi.Input

	ToFunctionNamespaceOutput() FunctionNamespaceOutput
	ToFunctionNamespaceOutputWithContext(ctx context.Context) FunctionNamespaceOutput
}

func (*FunctionNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionNamespace)(nil)).Elem()
}

func (i *FunctionNamespace) ToFunctionNamespaceOutput() FunctionNamespaceOutput {
	return i.ToFunctionNamespaceOutputWithContext(context.Background())
}

func (i *FunctionNamespace) ToFunctionNamespaceOutputWithContext(ctx context.Context) FunctionNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionNamespaceOutput)
}

// FunctionNamespaceArrayInput is an input type that accepts FunctionNamespaceArray and FunctionNamespaceArrayOutput values.
// You can construct a concrete instance of `FunctionNamespaceArrayInput` via:
//
//	FunctionNamespaceArray{ FunctionNamespaceArgs{...} }
type FunctionNamespaceArrayInput interface {
	pulumi.Input

	ToFunctionNamespaceArrayOutput() FunctionNamespaceArrayOutput
	ToFunctionNamespaceArrayOutputWithContext(context.Context) FunctionNamespaceArrayOutput
}

type FunctionNamespaceArray []FunctionNamespaceInput

func (FunctionNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionNamespace)(nil)).Elem()
}

func (i FunctionNamespaceArray) ToFunctionNamespaceArrayOutput() FunctionNamespaceArrayOutput {
	return i.ToFunctionNamespaceArrayOutputWithContext(context.Background())
}

func (i FunctionNamespaceArray) ToFunctionNamespaceArrayOutputWithContext(ctx context.Context) FunctionNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionNamespaceArrayOutput)
}

// FunctionNamespaceMapInput is an input type that accepts FunctionNamespaceMap and FunctionNamespaceMapOutput values.
// You can construct a concrete instance of `FunctionNamespaceMapInput` via:
//
//	FunctionNamespaceMap{ "key": FunctionNamespaceArgs{...} }
type FunctionNamespaceMapInput interface {
	pulumi.Input

	ToFunctionNamespaceMapOutput() FunctionNamespaceMapOutput
	ToFunctionNamespaceMapOutputWithContext(context.Context) FunctionNamespaceMapOutput
}

type FunctionNamespaceMap map[string]FunctionNamespaceInput

func (FunctionNamespaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionNamespace)(nil)).Elem()
}

func (i FunctionNamespaceMap) ToFunctionNamespaceMapOutput() FunctionNamespaceMapOutput {
	return i.ToFunctionNamespaceMapOutputWithContext(context.Background())
}

func (i FunctionNamespaceMap) ToFunctionNamespaceMapOutputWithContext(ctx context.Context) FunctionNamespaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionNamespaceMapOutput)
}

type FunctionNamespaceOutput struct{ *pulumi.OutputState }

func (FunctionNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionNamespace)(nil)).Elem()
}

func (o FunctionNamespaceOutput) ToFunctionNamespaceOutput() FunctionNamespaceOutput {
	return o
}

func (o FunctionNamespaceOutput) ToFunctionNamespaceOutputWithContext(ctx context.Context) FunctionNamespaceOutput {
	return o
}

// The description of the namespace.
func (o FunctionNamespaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The environment variables of the namespace.
func (o FunctionNamespaceOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringMapOutput { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// The unique name of the Functions namespace.
//
// > **Important** Updates to the `name` argument will recreate the namespace.
func (o FunctionNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Organization ID with which the namespace is associated.
func (o FunctionNamespaceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// `projectId`) The unique identifier of the project that contains the namespace.
func (o FunctionNamespaceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`). The region in which the namespace is created.
func (o FunctionNamespaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The registry endpoint of the namespace.
func (o FunctionNamespaceOutput) RegistryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringOutput { return v.RegistryEndpoint }).(pulumi.StringOutput)
}

// The registry namespace ID of the namespace.
func (o FunctionNamespaceOutput) RegistryNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringOutput { return v.RegistryNamespaceId }).(pulumi.StringOutput)
}

// The secret environment variables of the namespace.
func (o FunctionNamespaceOutput) SecretEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FunctionNamespace) pulumi.StringMapOutput { return v.SecretEnvironmentVariables }).(pulumi.StringMapOutput)
}

type FunctionNamespaceArrayOutput struct{ *pulumi.OutputState }

func (FunctionNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FunctionNamespace)(nil)).Elem()
}

func (o FunctionNamespaceArrayOutput) ToFunctionNamespaceArrayOutput() FunctionNamespaceArrayOutput {
	return o
}

func (o FunctionNamespaceArrayOutput) ToFunctionNamespaceArrayOutputWithContext(ctx context.Context) FunctionNamespaceArrayOutput {
	return o
}

func (o FunctionNamespaceArrayOutput) Index(i pulumi.IntInput) FunctionNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FunctionNamespace {
		return vs[0].([]*FunctionNamespace)[vs[1].(int)]
	}).(FunctionNamespaceOutput)
}

type FunctionNamespaceMapOutput struct{ *pulumi.OutputState }

func (FunctionNamespaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FunctionNamespace)(nil)).Elem()
}

func (o FunctionNamespaceMapOutput) ToFunctionNamespaceMapOutput() FunctionNamespaceMapOutput {
	return o
}

func (o FunctionNamespaceMapOutput) ToFunctionNamespaceMapOutputWithContext(ctx context.Context) FunctionNamespaceMapOutput {
	return o
}

func (o FunctionNamespaceMapOutput) MapIndex(k pulumi.StringInput) FunctionNamespaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FunctionNamespace {
		return vs[0].(map[string]*FunctionNamespace)[vs[1].(string)]
	}).(FunctionNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionNamespaceInput)(nil)).Elem(), &FunctionNamespace{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionNamespaceArrayInput)(nil)).Elem(), FunctionNamespaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionNamespaceMapInput)(nil)).Elem(), FunctionNamespaceMap{})
	pulumi.RegisterOutputType(FunctionNamespaceOutput{})
	pulumi.RegisterOutputType(FunctionNamespaceArrayOutput{})
	pulumi.RegisterOutputType(FunctionNamespaceMapOutput{})
}
