// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about an DocumentDB load balancer endpoint.
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
//			_, err := scaleway.GetDocumentdbLoadBalancerEndpoint(ctx, &scaleway.GetDocumentdbLoadBalancerEndpointArgs{
//				InstanceId: pulumi.StringRef("11111111-1111-1111-1111-111111111111"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetDocumentdbLoadBalancerEndpoint(ctx *pulumi.Context, args *GetDocumentdbLoadBalancerEndpointArgs, opts ...pulumi.InvokeOption) (*GetDocumentdbLoadBalancerEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDocumentdbLoadBalancerEndpointResult
	err := ctx.Invoke("scaleway:index/getDocumentdbLoadBalancerEndpoint:getDocumentdbLoadBalancerEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointArgs struct {
	// The DocumentDB Instance on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
	InstanceId *string `pulumi:"instanceId"`
	// The DocumentDB Instance Name on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
	InstanceName *string `pulumi:"instanceName"`
	// The ID of the project the DocumentDB endpoint is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region in which the DocumentDB endpoint exists.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointResult struct {
	// The hostname of your endpoint.
	Hostname string `pulumi:"hostname"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	InstanceId   string `pulumi:"instanceId"`
	InstanceName string `pulumi:"instanceName"`
	// The IP of your load balancer service.
	Ip string `pulumi:"ip"`
	// The name of your load balancer service.
	Name string `pulumi:"name"`
	// The port of your load balancer service.
	Port      int    `pulumi:"port"`
	ProjectId string `pulumi:"projectId"`
	Region    string `pulumi:"region"`
}

func GetDocumentdbLoadBalancerEndpointOutput(ctx *pulumi.Context, args GetDocumentdbLoadBalancerEndpointOutputArgs, opts ...pulumi.InvokeOption) GetDocumentdbLoadBalancerEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDocumentdbLoadBalancerEndpointResult, error) {
			args := v.(GetDocumentdbLoadBalancerEndpointArgs)
			r, err := GetDocumentdbLoadBalancerEndpoint(ctx, &args, opts...)
			var s GetDocumentdbLoadBalancerEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDocumentdbLoadBalancerEndpointResultOutput)
}

// A collection of arguments for invoking getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointOutputArgs struct {
	// The DocumentDB Instance on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The DocumentDB Instance Name on which the endpoint is attached. Only one of `instanceName` and `instanceId` should be specified.
	InstanceName pulumi.StringPtrInput `pulumi:"instanceName"`
	// The ID of the project the DocumentDB endpoint is associated with.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
	// `region`) The region in which the DocumentDB endpoint exists.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetDocumentdbLoadBalancerEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentdbLoadBalancerEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getDocumentdbLoadBalancerEndpoint.
type GetDocumentdbLoadBalancerEndpointResultOutput struct{ *pulumi.OutputState }

func (GetDocumentdbLoadBalancerEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDocumentdbLoadBalancerEndpointResult)(nil)).Elem()
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ToGetDocumentdbLoadBalancerEndpointResultOutput() GetDocumentdbLoadBalancerEndpointResultOutput {
	return o
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ToGetDocumentdbLoadBalancerEndpointResultOutputWithContext(ctx context.Context) GetDocumentdbLoadBalancerEndpointResultOutput {
	return o
}

// The hostname of your endpoint.
func (o GetDocumentdbLoadBalancerEndpointResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDocumentdbLoadBalancerEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

// The IP of your load balancer service.
func (o GetDocumentdbLoadBalancerEndpointResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Ip }).(pulumi.StringOutput)
}

// The name of your load balancer service.
func (o GetDocumentdbLoadBalancerEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// The port of your load balancer service.
func (o GetDocumentdbLoadBalancerEndpointResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) int { return v.Port }).(pulumi.IntOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetDocumentdbLoadBalancerEndpointResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDocumentdbLoadBalancerEndpointResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDocumentdbLoadBalancerEndpointResultOutput{})
}
