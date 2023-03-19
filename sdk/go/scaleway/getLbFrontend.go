// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about Scaleway Load-Balancer Frontends.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/lb/zoned_api/#frontends-a6a28d).
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
//			ip01, err := scaleway.NewLbIp(ctx, "ip01", nil)
//			if err != nil {
//				return err
//			}
//			lb01, err := scaleway.NewLb(ctx, "lb01", &scaleway.LbArgs{
//				IpId: ip01.ID(),
//				Type: pulumi.String("lb-s"),
//			})
//			if err != nil {
//				return err
//			}
//			bkd01, err := scaleway.NewLbBackend(ctx, "bkd01", &scaleway.LbBackendArgs{
//				LbId:            lb01.ID(),
//				ForwardProtocol: pulumi.String("tcp"),
//				ForwardPort:     pulumi.Int(80),
//				ProxyProtocol:   pulumi.String("none"),
//			})
//			if err != nil {
//				return err
//			}
//			frt01, err := scaleway.NewLbFrontend(ctx, "frt01", &scaleway.LbFrontendArgs{
//				LbId:        lb01.ID(),
//				BackendId:   bkd01.ID(),
//				InboundPort: pulumi.Int(80),
//			})
//			if err != nil {
//				return err
//			}
//			_ = scaleway.LookupLbFrontendOutput(ctx, scaleway.GetLbFrontendOutputArgs{
//				FrontendId: frt01.ID(),
//			}, nil)
//			_ = scaleway.LookupLbFrontendOutput(ctx, scaleway.GetLbFrontendOutputArgs{
//				Name: frt01.Name,
//				LbId: lb01.ID(),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func LookupLbFrontend(ctx *pulumi.Context, args *LookupLbFrontendArgs, opts ...pulumi.InvokeOption) (*LookupLbFrontendResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupLbFrontendResult
	err := ctx.Invoke("scaleway:index/getLbFrontend:getLbFrontend", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbFrontend.
type LookupLbFrontendArgs struct {
	// The frontend id.
	// - Only one of `name` and `frontendId` should be specified.
	FrontendId *string `pulumi:"frontendId"`
	// The load-balancer ID this frontend is attached to.
	LbId *string `pulumi:"lbId"`
	// The name of the frontend.
	// - When using the `name` you should specify the `lb-id`
	Name *string `pulumi:"name"`
}

// A collection of values returned by getLbFrontend.
type LookupLbFrontendResult struct {
	Acls           []GetLbFrontendAcl `pulumi:"acls"`
	BackendId      string             `pulumi:"backendId"`
	CertificateId  string             `pulumi:"certificateId"`
	CertificateIds []string           `pulumi:"certificateIds"`
	EnableHttp3    bool               `pulumi:"enableHttp3"`
	FrontendId     *string            `pulumi:"frontendId"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	InboundPort   int     `pulumi:"inboundPort"`
	LbId          *string `pulumi:"lbId"`
	Name          *string `pulumi:"name"`
	TimeoutClient string  `pulumi:"timeoutClient"`
}

func LookupLbFrontendOutput(ctx *pulumi.Context, args LookupLbFrontendOutputArgs, opts ...pulumi.InvokeOption) LookupLbFrontendResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLbFrontendResult, error) {
			args := v.(LookupLbFrontendArgs)
			r, err := LookupLbFrontend(ctx, &args, opts...)
			var s LookupLbFrontendResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLbFrontendResultOutput)
}

// A collection of arguments for invoking getLbFrontend.
type LookupLbFrontendOutputArgs struct {
	// The frontend id.
	// - Only one of `name` and `frontendId` should be specified.
	FrontendId pulumi.StringPtrInput `pulumi:"frontendId"`
	// The load-balancer ID this frontend is attached to.
	LbId pulumi.StringPtrInput `pulumi:"lbId"`
	// The name of the frontend.
	// - When using the `name` you should specify the `lb-id`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupLbFrontendOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbFrontendArgs)(nil)).Elem()
}

// A collection of values returned by getLbFrontend.
type LookupLbFrontendResultOutput struct{ *pulumi.OutputState }

func (LookupLbFrontendResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbFrontendResult)(nil)).Elem()
}

func (o LookupLbFrontendResultOutput) ToLookupLbFrontendResultOutput() LookupLbFrontendResultOutput {
	return o
}

func (o LookupLbFrontendResultOutput) ToLookupLbFrontendResultOutputWithContext(ctx context.Context) LookupLbFrontendResultOutput {
	return o
}

func (o LookupLbFrontendResultOutput) Acls() GetLbFrontendAclArrayOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) []GetLbFrontendAcl { return v.Acls }).(GetLbFrontendAclArrayOutput)
}

func (o LookupLbFrontendResultOutput) BackendId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) string { return v.BackendId }).(pulumi.StringOutput)
}

func (o LookupLbFrontendResultOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) string { return v.CertificateId }).(pulumi.StringOutput)
}

func (o LookupLbFrontendResultOutput) CertificateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) []string { return v.CertificateIds }).(pulumi.StringArrayOutput)
}

func (o LookupLbFrontendResultOutput) EnableHttp3() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) bool { return v.EnableHttp3 }).(pulumi.BoolOutput)
}

func (o LookupLbFrontendResultOutput) FrontendId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) *string { return v.FrontendId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbFrontendResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbFrontendResultOutput) InboundPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) int { return v.InboundPort }).(pulumi.IntOutput)
}

func (o LookupLbFrontendResultOutput) LbId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) *string { return v.LbId }).(pulumi.StringPtrOutput)
}

func (o LookupLbFrontendResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupLbFrontendResultOutput) TimeoutClient() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbFrontendResult) string { return v.TimeoutClient }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbFrontendResultOutput{})
}
