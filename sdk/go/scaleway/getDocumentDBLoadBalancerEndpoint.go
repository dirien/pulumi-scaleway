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

func GetDocumentDBLoadBalancerEndpoint(ctx *pulumi.Context, args *GetDocumentDBLoadBalancerEndpointArgs, opts ...pulumi.InvokeOption) (*GetDocumentDBLoadBalancerEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDocumentDBLoadBalancerEndpointResult
	err := ctx.Invoke("scaleway:index/getDocumentDBLoadBalancerEndpoint:getDocumentDBLoadBalancerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocumentDBLoadBalancerEndpoint.
type GetDocumentDBLoadBalancerEndpointArgs struct {
	InstanceId   *string `pulumi:"instanceId"`
	InstanceName *string `pulumi:"instanceName"`
	ProjectId    *string `pulumi:"projectId"`
	Region       *string `pulumi:"region"`
}

// A collection of values returned by getDocumentDBLoadBalancerEndpoint.
type GetDocumentDBLoadBalancerEndpointResult struct {
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	InstanceId   string `pulumi:"instanceId"`
	InstanceName string `pulumi:"instanceName"`
	Ip           string `pulumi:"ip"`
	Name         string `pulumi:"name"`
	Port         int    `pulumi:"port"`
	ProjectId    string `pulumi:"projectId"`
	Region       string `pulumi:"region"`
}

func GetDocumentDBLoadBalancerEndpointOutput(ctx *pulumi.Context, args GetDocumentDBLoadBalancerEndpointOutputArgs, opts ...pulumi.InvokeOption) GetDocumentDBLoadBalancerEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDocumentDBLoadBalancerEndpointResult, error) {
			args := v.(GetDocumentDBLoadBalancerEndpointArgs)
			r, err := GetDocumentDBLoadBalancerEndpoint(ctx, &args, opts...)
			var s GetDocumentDBLoadBalancerEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDocumentDBLoadBalancerEndpointResultOutput)
}

// A collection of arguments for invoking getDocumentDBLoadBalancerEndpoint.
type GetDocumentDBLoadBalancerEndpointOutputArgs struct {
	InstanceId   pulumi.StringPtrInput `pulumi:"instanceId"`
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	ProjectId    pulumi.StringPtrInput `pulumi:"projectId"`
	Region       pulumi.StringPtrInput `pulumi:"region"`
}

func (GetDocumentDBLoadBalancerEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentDBLoadBalancerEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getDocumentDBLoadBalancerEndpoint.
type GetDocumentDBLoadBalancerEndpointResultOutput struct{ *pulumi.OutputState }

func (GetDocumentDBLoadBalancerEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentDBLoadBalancerEndpointResult)(nil)).Elem()
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) ToGetDocumentDBLoadBalancerEndpointResultOutput() GetDocumentDBLoadBalancerEndpointResultOutput {
	return o
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) ToGetDocumentDBLoadBalancerEndpointResultOutputWithContext(ctx context.Context) GetDocumentDBLoadBalancerEndpointResultOutput {
	return o
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetDocumentDBLoadBalancerEndpointResult] {
	return pulumix.Output[GetDocumentDBLoadBalancerEndpointResult]{
		OutputState: o.OutputState,
	}
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDocumentDBLoadBalancerEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.Ip }).(pulumi.StringOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetDocumentDBLoadBalancerEndpointResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentDBLoadBalancerEndpointResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDocumentDBLoadBalancerEndpointResultOutput{})
}
