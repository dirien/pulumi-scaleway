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

// Creates and manages Scaleway Managed Inference deployments.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/inference/).
//
// ## Example Usage
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
//			_, err := scaleway.NewInferenceDeployment(ctx, "deployment", &scaleway.InferenceDeploymentArgs{
//				AcceptEula: pulumi.Bool(true),
//				ModelName:  pulumi.String("meta/llama-3.1-8b-instruct:fp8"),
//				NodeType:   pulumi.String("L4"),
//				PublicEndpoint: &scaleway.InferenceDeploymentPublicEndpointArgs{
//					IsEnabled: pulumi.Bool(true),
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
// Functions can be imported using, `{region}/{id}`, as shown below:
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/inferenceDeployment:InferenceDeployment deployment fr-par/11111111-1111-1111-1111-111111111111
// ```
type InferenceDeployment struct {
	pulumi.CustomResourceState

	// Some models (e.g Meta Llama) require end-user license agreements. Set `true` to accept.
	AcceptEula pulumi.BoolPtrOutput `pulumi:"acceptEula"`
	// The date and time of the creation of the deployment.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The maximum size of the pool.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// The minimum size of the pool.
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// The model id used for the deployment.
	ModelId pulumi.StringOutput `pulumi:"modelId"`
	// The model name to use for the deployment. Model names can be found in Console or using Scaleway's CLI (`scw inference model list`)
	ModelName pulumi.StringOutput `pulumi:"modelName"`
	// The deployment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (`scw inference node-type list`)
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// Configuration of the deployment's private endpoint.
	PrivateEndpoint InferenceDeploymentPrivateEndpointPtrOutput `pulumi:"privateEndpoint"`
	// `projectId`) The ID of the project the deployment is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Configuration of the deployment's public endpoint.
	PublicEndpoint InferenceDeploymentPublicEndpointPtrOutput `pulumi:"publicEndpoint"`
	// `region`) The region in which the deployment is created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The size of the pool.
	Size pulumi.IntOutput `pulumi:"size"`
	// The status of the deployment.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags associated with the deployment.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The date and time of the last update of the deployment.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewInferenceDeployment registers a new resource with the given unique name, arguments, and options.
func NewInferenceDeployment(ctx *pulumi.Context,
	name string, args *InferenceDeploymentArgs, opts ...pulumi.ResourceOption) (*InferenceDeployment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelName == nil {
		return nil, errors.New("invalid value for required argument 'ModelName'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InferenceDeployment
	err := ctx.RegisterResource("scaleway:index/inferenceDeployment:InferenceDeployment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInferenceDeployment gets an existing InferenceDeployment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInferenceDeployment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InferenceDeploymentState, opts ...pulumi.ResourceOption) (*InferenceDeployment, error) {
	var resource InferenceDeployment
	err := ctx.ReadResource("scaleway:index/inferenceDeployment:InferenceDeployment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InferenceDeployment resources.
type inferenceDeploymentState struct {
	// Some models (e.g Meta Llama) require end-user license agreements. Set `true` to accept.
	AcceptEula *bool `pulumi:"acceptEula"`
	// The date and time of the creation of the deployment.
	CreatedAt *string `pulumi:"createdAt"`
	// The maximum size of the pool.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size of the pool.
	MinSize *int `pulumi:"minSize"`
	// The model id used for the deployment.
	ModelId *string `pulumi:"modelId"`
	// The model name to use for the deployment. Model names can be found in Console or using Scaleway's CLI (`scw inference model list`)
	ModelName *string `pulumi:"modelName"`
	// The deployment name.
	Name *string `pulumi:"name"`
	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (`scw inference node-type list`)
	NodeType *string `pulumi:"nodeType"`
	// Configuration of the deployment's private endpoint.
	PrivateEndpoint *InferenceDeploymentPrivateEndpoint `pulumi:"privateEndpoint"`
	// `projectId`) The ID of the project the deployment is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Configuration of the deployment's public endpoint.
	PublicEndpoint *InferenceDeploymentPublicEndpoint `pulumi:"publicEndpoint"`
	// `region`) The region in which the deployment is created.
	Region *string `pulumi:"region"`
	// The size of the pool.
	Size *int `pulumi:"size"`
	// The status of the deployment.
	Status *string `pulumi:"status"`
	// The tags associated with the deployment.
	Tags []string `pulumi:"tags"`
	// The date and time of the last update of the deployment.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type InferenceDeploymentState struct {
	// Some models (e.g Meta Llama) require end-user license agreements. Set `true` to accept.
	AcceptEula pulumi.BoolPtrInput
	// The date and time of the creation of the deployment.
	CreatedAt pulumi.StringPtrInput
	// The maximum size of the pool.
	MaxSize pulumi.IntPtrInput
	// The minimum size of the pool.
	MinSize pulumi.IntPtrInput
	// The model id used for the deployment.
	ModelId pulumi.StringPtrInput
	// The model name to use for the deployment. Model names can be found in Console or using Scaleway's CLI (`scw inference model list`)
	ModelName pulumi.StringPtrInput
	// The deployment name.
	Name pulumi.StringPtrInput
	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (`scw inference node-type list`)
	NodeType pulumi.StringPtrInput
	// Configuration of the deployment's private endpoint.
	PrivateEndpoint InferenceDeploymentPrivateEndpointPtrInput
	// `projectId`) The ID of the project the deployment is associated with.
	ProjectId pulumi.StringPtrInput
	// Configuration of the deployment's public endpoint.
	PublicEndpoint InferenceDeploymentPublicEndpointPtrInput
	// `region`) The region in which the deployment is created.
	Region pulumi.StringPtrInput
	// The size of the pool.
	Size pulumi.IntPtrInput
	// The status of the deployment.
	Status pulumi.StringPtrInput
	// The tags associated with the deployment.
	Tags pulumi.StringArrayInput
	// The date and time of the last update of the deployment.
	UpdatedAt pulumi.StringPtrInput
}

func (InferenceDeploymentState) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceDeploymentState)(nil)).Elem()
}

type inferenceDeploymentArgs struct {
	// Some models (e.g Meta Llama) require end-user license agreements. Set `true` to accept.
	AcceptEula *bool `pulumi:"acceptEula"`
	// The maximum size of the pool.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size of the pool.
	MinSize *int `pulumi:"minSize"`
	// The model name to use for the deployment. Model names can be found in Console or using Scaleway's CLI (`scw inference model list`)
	ModelName string `pulumi:"modelName"`
	// The deployment name.
	Name *string `pulumi:"name"`
	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (`scw inference node-type list`)
	NodeType string `pulumi:"nodeType"`
	// Configuration of the deployment's private endpoint.
	PrivateEndpoint *InferenceDeploymentPrivateEndpoint `pulumi:"privateEndpoint"`
	// `projectId`) The ID of the project the deployment is associated with.
	ProjectId *string `pulumi:"projectId"`
	// Configuration of the deployment's public endpoint.
	PublicEndpoint *InferenceDeploymentPublicEndpoint `pulumi:"publicEndpoint"`
	// `region`) The region in which the deployment is created.
	Region *string `pulumi:"region"`
	// The tags associated with the deployment.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a InferenceDeployment resource.
type InferenceDeploymentArgs struct {
	// Some models (e.g Meta Llama) require end-user license agreements. Set `true` to accept.
	AcceptEula pulumi.BoolPtrInput
	// The maximum size of the pool.
	MaxSize pulumi.IntPtrInput
	// The minimum size of the pool.
	MinSize pulumi.IntPtrInput
	// The model name to use for the deployment. Model names can be found in Console or using Scaleway's CLI (`scw inference model list`)
	ModelName pulumi.StringInput
	// The deployment name.
	Name pulumi.StringPtrInput
	// The node type to use for the deployment. Node types can be found using Scaleway's CLI (`scw inference node-type list`)
	NodeType pulumi.StringInput
	// Configuration of the deployment's private endpoint.
	PrivateEndpoint InferenceDeploymentPrivateEndpointPtrInput
	// `projectId`) The ID of the project the deployment is associated with.
	ProjectId pulumi.StringPtrInput
	// Configuration of the deployment's public endpoint.
	PublicEndpoint InferenceDeploymentPublicEndpointPtrInput
	// `region`) The region in which the deployment is created.
	Region pulumi.StringPtrInput
	// The tags associated with the deployment.
	Tags pulumi.StringArrayInput
}

func (InferenceDeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceDeploymentArgs)(nil)).Elem()
}

type InferenceDeploymentInput interface {
	pulumi.Input

	ToInferenceDeploymentOutput() InferenceDeploymentOutput
	ToInferenceDeploymentOutputWithContext(ctx context.Context) InferenceDeploymentOutput
}

func (*InferenceDeployment) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceDeployment)(nil)).Elem()
}

func (i *InferenceDeployment) ToInferenceDeploymentOutput() InferenceDeploymentOutput {
	return i.ToInferenceDeploymentOutputWithContext(context.Background())
}

func (i *InferenceDeployment) ToInferenceDeploymentOutputWithContext(ctx context.Context) InferenceDeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceDeploymentOutput)
}

// InferenceDeploymentArrayInput is an input type that accepts InferenceDeploymentArray and InferenceDeploymentArrayOutput values.
// You can construct a concrete instance of `InferenceDeploymentArrayInput` via:
//
//	InferenceDeploymentArray{ InferenceDeploymentArgs{...} }
type InferenceDeploymentArrayInput interface {
	pulumi.Input

	ToInferenceDeploymentArrayOutput() InferenceDeploymentArrayOutput
	ToInferenceDeploymentArrayOutputWithContext(context.Context) InferenceDeploymentArrayOutput
}

type InferenceDeploymentArray []InferenceDeploymentInput

func (InferenceDeploymentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InferenceDeployment)(nil)).Elem()
}

func (i InferenceDeploymentArray) ToInferenceDeploymentArrayOutput() InferenceDeploymentArrayOutput {
	return i.ToInferenceDeploymentArrayOutputWithContext(context.Background())
}

func (i InferenceDeploymentArray) ToInferenceDeploymentArrayOutputWithContext(ctx context.Context) InferenceDeploymentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceDeploymentArrayOutput)
}

// InferenceDeploymentMapInput is an input type that accepts InferenceDeploymentMap and InferenceDeploymentMapOutput values.
// You can construct a concrete instance of `InferenceDeploymentMapInput` via:
//
//	InferenceDeploymentMap{ "key": InferenceDeploymentArgs{...} }
type InferenceDeploymentMapInput interface {
	pulumi.Input

	ToInferenceDeploymentMapOutput() InferenceDeploymentMapOutput
	ToInferenceDeploymentMapOutputWithContext(context.Context) InferenceDeploymentMapOutput
}

type InferenceDeploymentMap map[string]InferenceDeploymentInput

func (InferenceDeploymentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InferenceDeployment)(nil)).Elem()
}

func (i InferenceDeploymentMap) ToInferenceDeploymentMapOutput() InferenceDeploymentMapOutput {
	return i.ToInferenceDeploymentMapOutputWithContext(context.Background())
}

func (i InferenceDeploymentMap) ToInferenceDeploymentMapOutputWithContext(ctx context.Context) InferenceDeploymentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceDeploymentMapOutput)
}

type InferenceDeploymentOutput struct{ *pulumi.OutputState }

func (InferenceDeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceDeployment)(nil)).Elem()
}

func (o InferenceDeploymentOutput) ToInferenceDeploymentOutput() InferenceDeploymentOutput {
	return o
}

func (o InferenceDeploymentOutput) ToInferenceDeploymentOutputWithContext(ctx context.Context) InferenceDeploymentOutput {
	return o
}

// Some models (e.g Meta Llama) require end-user license agreements. Set `true` to accept.
func (o InferenceDeploymentOutput) AcceptEula() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.BoolPtrOutput { return v.AcceptEula }).(pulumi.BoolPtrOutput)
}

// The date and time of the creation of the deployment.
func (o InferenceDeploymentOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The maximum size of the pool.
func (o InferenceDeploymentOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// The minimum size of the pool.
func (o InferenceDeploymentOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.IntOutput { return v.MinSize }).(pulumi.IntOutput)
}

// The model id used for the deployment.
func (o InferenceDeploymentOutput) ModelId() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.ModelId }).(pulumi.StringOutput)
}

// The model name to use for the deployment. Model names can be found in Console or using Scaleway's CLI (`scw inference model list`)
func (o InferenceDeploymentOutput) ModelName() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.ModelName }).(pulumi.StringOutput)
}

// The deployment name.
func (o InferenceDeploymentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The node type to use for the deployment. Node types can be found using Scaleway's CLI (`scw inference node-type list`)
func (o InferenceDeploymentOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// Configuration of the deployment's private endpoint.
func (o InferenceDeploymentOutput) PrivateEndpoint() InferenceDeploymentPrivateEndpointPtrOutput {
	return o.ApplyT(func(v *InferenceDeployment) InferenceDeploymentPrivateEndpointPtrOutput { return v.PrivateEndpoint }).(InferenceDeploymentPrivateEndpointPtrOutput)
}

// `projectId`) The ID of the project the deployment is associated with.
func (o InferenceDeploymentOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Configuration of the deployment's public endpoint.
func (o InferenceDeploymentOutput) PublicEndpoint() InferenceDeploymentPublicEndpointPtrOutput {
	return o.ApplyT(func(v *InferenceDeployment) InferenceDeploymentPublicEndpointPtrOutput { return v.PublicEndpoint }).(InferenceDeploymentPublicEndpointPtrOutput)
}

// `region`) The region in which the deployment is created.
func (o InferenceDeploymentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The size of the pool.
func (o InferenceDeploymentOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The status of the deployment.
func (o InferenceDeploymentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags associated with the deployment.
func (o InferenceDeploymentOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The date and time of the last update of the deployment.
func (o InferenceDeploymentOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceDeployment) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type InferenceDeploymentArrayOutput struct{ *pulumi.OutputState }

func (InferenceDeploymentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InferenceDeployment)(nil)).Elem()
}

func (o InferenceDeploymentArrayOutput) ToInferenceDeploymentArrayOutput() InferenceDeploymentArrayOutput {
	return o
}

func (o InferenceDeploymentArrayOutput) ToInferenceDeploymentArrayOutputWithContext(ctx context.Context) InferenceDeploymentArrayOutput {
	return o
}

func (o InferenceDeploymentArrayOutput) Index(i pulumi.IntInput) InferenceDeploymentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InferenceDeployment {
		return vs[0].([]*InferenceDeployment)[vs[1].(int)]
	}).(InferenceDeploymentOutput)
}

type InferenceDeploymentMapOutput struct{ *pulumi.OutputState }

func (InferenceDeploymentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InferenceDeployment)(nil)).Elem()
}

func (o InferenceDeploymentMapOutput) ToInferenceDeploymentMapOutput() InferenceDeploymentMapOutput {
	return o
}

func (o InferenceDeploymentMapOutput) ToInferenceDeploymentMapOutputWithContext(ctx context.Context) InferenceDeploymentMapOutput {
	return o
}

func (o InferenceDeploymentMapOutput) MapIndex(k pulumi.StringInput) InferenceDeploymentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InferenceDeployment {
		return vs[0].(map[string]*InferenceDeployment)[vs[1].(string)]
	}).(InferenceDeploymentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InferenceDeploymentInput)(nil)).Elem(), &InferenceDeployment{})
	pulumi.RegisterInputType(reflect.TypeOf((*InferenceDeploymentArrayInput)(nil)).Elem(), InferenceDeploymentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InferenceDeploymentMapInput)(nil)).Elem(), InferenceDeploymentMap{})
	pulumi.RegisterOutputType(InferenceDeploymentOutput{})
	pulumi.RegisterOutputType(InferenceDeploymentArrayOutput{})
	pulumi.RegisterOutputType(InferenceDeploymentMapOutput{})
}
