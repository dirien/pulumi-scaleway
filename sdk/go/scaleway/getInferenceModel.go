// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `InferenceModel` data source allows you to retrieve information about an inference model available in the Scaleway Inference API, either by providing the model's `name` or its `modelId`.
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
//			_, err := scaleway.LookupInferenceModel(ctx, &scaleway.LookupInferenceModelArgs{
//				Name: pulumi.StringRef("meta/llama-3.1-8b-instruct:fp8"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupInferenceModel(ctx *pulumi.Context, args *LookupInferenceModelArgs, opts ...pulumi.InvokeOption) (*LookupInferenceModelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInferenceModelResult
	err := ctx.Invoke("scaleway:index/getInferenceModel:getInferenceModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInferenceModel.
type LookupInferenceModelArgs struct {
	// The ID of the model to retrieve. Must be a valid UUID with locality (i.e., Scaleway's zoned UUID format).
	ModelId *string `pulumi:"modelId"`
	// The fully qualified name of the model to look up (e.g., "meta/llama-3.1-8b-instruct:fp8"). The provider will search for a model with an exact name match in the selected region and project.
	Name *string `pulumi:"name"`
	Url  *string `pulumi:"url"`
}

// A collection of values returned by getInferenceModel.
type LookupInferenceModelResult struct {
	CreatedAt string `pulumi:"createdAt"`
	// A textual description of the model (if available).
	Description string `pulumi:"description"`
	// Whether the model requires end-user license agreement acceptance before use.
	HasEula bool `pulumi:"hasEula"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	ModelId *string `pulumi:"modelId"`
	Name    *string `pulumi:"name"`
	// List of supported node types and their quantization options. Each entry contains:
	NodesSupports []GetInferenceModelNodesSupport `pulumi:"nodesSupports"`
	// Size, in bits, of the model parameters.
	ParameterSizeBits int    `pulumi:"parameterSizeBits"`
	ProjectId         string `pulumi:"projectId"`
	Region            string `pulumi:"region"`
	Secret            string `pulumi:"secret"`
	// Total size, in bytes, of the model archive.
	SizeBytes int `pulumi:"sizeBytes"`
	// The current status of the model (e.g., ready, error, etc.).
	Status string `pulumi:"status"`
	// Tags associated with the model.
	Tags      []string `pulumi:"tags"`
	UpdatedAt string   `pulumi:"updatedAt"`
	Url       *string  `pulumi:"url"`
}

func LookupInferenceModelOutput(ctx *pulumi.Context, args LookupInferenceModelOutputArgs, opts ...pulumi.InvokeOption) LookupInferenceModelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInferenceModelResultOutput, error) {
			args := v.(LookupInferenceModelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("scaleway:index/getInferenceModel:getInferenceModel", args, LookupInferenceModelResultOutput{}, options).(LookupInferenceModelResultOutput), nil
		}).(LookupInferenceModelResultOutput)
}

// A collection of arguments for invoking getInferenceModel.
type LookupInferenceModelOutputArgs struct {
	// The ID of the model to retrieve. Must be a valid UUID with locality (i.e., Scaleway's zoned UUID format).
	ModelId pulumi.StringPtrInput `pulumi:"modelId"`
	// The fully qualified name of the model to look up (e.g., "meta/llama-3.1-8b-instruct:fp8"). The provider will search for a model with an exact name match in the selected region and project.
	Name pulumi.StringPtrInput `pulumi:"name"`
	Url  pulumi.StringPtrInput `pulumi:"url"`
}

func (LookupInferenceModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceModelArgs)(nil)).Elem()
}

// A collection of values returned by getInferenceModel.
type LookupInferenceModelResultOutput struct{ *pulumi.OutputState }

func (LookupInferenceModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInferenceModelResult)(nil)).Elem()
}

func (o LookupInferenceModelResultOutput) ToLookupInferenceModelResultOutput() LookupInferenceModelResultOutput {
	return o
}

func (o LookupInferenceModelResultOutput) ToLookupInferenceModelResultOutputWithContext(ctx context.Context) LookupInferenceModelResultOutput {
	return o
}

func (o LookupInferenceModelResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A textual description of the model (if available).
func (o LookupInferenceModelResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.Description }).(pulumi.StringOutput)
}

// Whether the model requires end-user license agreement acceptance before use.
func (o LookupInferenceModelResultOutput) HasEula() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) bool { return v.HasEula }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInferenceModelResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInferenceModelResultOutput) ModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) *string { return v.ModelId }).(pulumi.StringPtrOutput)
}

func (o LookupInferenceModelResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// List of supported node types and their quantization options. Each entry contains:
func (o LookupInferenceModelResultOutput) NodesSupports() GetInferenceModelNodesSupportArrayOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) []GetInferenceModelNodesSupport { return v.NodesSupports }).(GetInferenceModelNodesSupportArrayOutput)
}

// Size, in bits, of the model parameters.
func (o LookupInferenceModelResultOutput) ParameterSizeBits() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) int { return v.ParameterSizeBits }).(pulumi.IntOutput)
}

func (o LookupInferenceModelResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o LookupInferenceModelResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupInferenceModelResultOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.Secret }).(pulumi.StringOutput)
}

// Total size, in bytes, of the model archive.
func (o LookupInferenceModelResultOutput) SizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) int { return v.SizeBytes }).(pulumi.IntOutput)
}

// The current status of the model (e.g., ready, error, etc.).
func (o LookupInferenceModelResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.Status }).(pulumi.StringOutput)
}

// Tags associated with the model.
func (o LookupInferenceModelResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupInferenceModelResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o LookupInferenceModelResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInferenceModelResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInferenceModelResultOutput{})
}
