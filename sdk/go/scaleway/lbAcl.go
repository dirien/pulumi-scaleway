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

// Creates and manages Scaleway Load-Balancer ACLs. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls).
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
//			_, err := scaleway.NewLbAcl(ctx, "acl01", &scaleway.LbAclArgs{
//				FrontendId:  pulumi.Any(scaleway_lb_frontend.Frt01.Id),
//				Description: pulumi.String("Exclude well-known IPs"),
//				Index:       pulumi.Int(0),
//				Action: &scaleway.LbAclActionArgs{
//					Type: pulumi.String("allow"),
//				},
//				Match: &scaleway.LbAclMatchArgs{
//					IpSubnets: pulumi.StringArray{
//						pulumi.String("192.168.0.1"),
//						pulumi.String("192.168.0.2"),
//						pulumi.String("192.168.10.0/24"),
//					},
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
// Load-Balancer ACL can be imported using the `{zone}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/lbAcl:LbAcl acl01 fr-par-1/11111111-1111-1111-1111-111111111111
// ```
type LbAcl struct {
	pulumi.CustomResourceState

	// Action to undertake when an ACL filter matches.
	Action LbAclActionOutput `pulumi:"action"`
	// IsDate and time of ACL's creation (RFC 3339 format)
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The ACL description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId pulumi.StringOutput `pulumi:"frontendId"`
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index pulumi.IntOutput `pulumi:"index"`
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match LbAclMatchPtrOutput `pulumi:"match"`
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringOutput `pulumi:"name"`
	// IsDate and time of ACL's update (RFC 3339 format)
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewLbAcl registers a new resource with the given unique name, arguments, and options.
func NewLbAcl(ctx *pulumi.Context,
	name string, args *LbAclArgs, opts ...pulumi.ResourceOption) (*LbAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.FrontendId == nil {
		return nil, errors.New("invalid value for required argument 'FrontendId'")
	}
	if args.Index == nil {
		return nil, errors.New("invalid value for required argument 'Index'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbAcl
	err := ctx.RegisterResource("scaleway:index/lbAcl:LbAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbAcl gets an existing LbAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbAclState, opts ...pulumi.ResourceOption) (*LbAcl, error) {
	var resource LbAcl
	err := ctx.ReadResource("scaleway:index/lbAcl:LbAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbAcl resources.
type lbAclState struct {
	// Action to undertake when an ACL filter matches.
	Action *LbAclAction `pulumi:"action"`
	// IsDate and time of ACL's creation (RFC 3339 format)
	CreatedAt *string `pulumi:"createdAt"`
	// The ACL description.
	Description *string `pulumi:"description"`
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId *string `pulumi:"frontendId"`
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index *int `pulumi:"index"`
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match *LbAclMatch `pulumi:"match"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
	// IsDate and time of ACL's update (RFC 3339 format)
	UpdatedAt *string `pulumi:"updatedAt"`
}

type LbAclState struct {
	// Action to undertake when an ACL filter matches.
	Action LbAclActionPtrInput
	// IsDate and time of ACL's creation (RFC 3339 format)
	CreatedAt pulumi.StringPtrInput
	// The ACL description.
	Description pulumi.StringPtrInput
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId pulumi.StringPtrInput
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index pulumi.IntPtrInput
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match LbAclMatchPtrInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
	// IsDate and time of ACL's update (RFC 3339 format)
	UpdatedAt pulumi.StringPtrInput
}

func (LbAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbAclState)(nil)).Elem()
}

type lbAclArgs struct {
	// Action to undertake when an ACL filter matches.
	Action LbAclAction `pulumi:"action"`
	// The ACL description.
	Description *string `pulumi:"description"`
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId string `pulumi:"frontendId"`
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index int `pulumi:"index"`
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match *LbAclMatch `pulumi:"match"`
	// The ACL name. If not provided it will be randomly generated.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a LbAcl resource.
type LbAclArgs struct {
	// Action to undertake when an ACL filter matches.
	Action LbAclActionInput
	// The ACL description.
	Description pulumi.StringPtrInput
	// The load-balancer Frontend ID to attach the ACL to.
	FrontendId pulumi.StringInput
	// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
	Index pulumi.IntInput
	// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
	Match LbAclMatchPtrInput
	// The ACL name. If not provided it will be randomly generated.
	Name pulumi.StringPtrInput
}

func (LbAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbAclArgs)(nil)).Elem()
}

type LbAclInput interface {
	pulumi.Input

	ToLbAclOutput() LbAclOutput
	ToLbAclOutputWithContext(ctx context.Context) LbAclOutput
}

func (*LbAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**LbAcl)(nil)).Elem()
}

func (i *LbAcl) ToLbAclOutput() LbAclOutput {
	return i.ToLbAclOutputWithContext(context.Background())
}

func (i *LbAcl) ToLbAclOutputWithContext(ctx context.Context) LbAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbAclOutput)
}

// LbAclArrayInput is an input type that accepts LbAclArray and LbAclArrayOutput values.
// You can construct a concrete instance of `LbAclArrayInput` via:
//
//	LbAclArray{ LbAclArgs{...} }
type LbAclArrayInput interface {
	pulumi.Input

	ToLbAclArrayOutput() LbAclArrayOutput
	ToLbAclArrayOutputWithContext(context.Context) LbAclArrayOutput
}

type LbAclArray []LbAclInput

func (LbAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbAcl)(nil)).Elem()
}

func (i LbAclArray) ToLbAclArrayOutput() LbAclArrayOutput {
	return i.ToLbAclArrayOutputWithContext(context.Background())
}

func (i LbAclArray) ToLbAclArrayOutputWithContext(ctx context.Context) LbAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbAclArrayOutput)
}

// LbAclMapInput is an input type that accepts LbAclMap and LbAclMapOutput values.
// You can construct a concrete instance of `LbAclMapInput` via:
//
//	LbAclMap{ "key": LbAclArgs{...} }
type LbAclMapInput interface {
	pulumi.Input

	ToLbAclMapOutput() LbAclMapOutput
	ToLbAclMapOutputWithContext(context.Context) LbAclMapOutput
}

type LbAclMap map[string]LbAclInput

func (LbAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbAcl)(nil)).Elem()
}

func (i LbAclMap) ToLbAclMapOutput() LbAclMapOutput {
	return i.ToLbAclMapOutputWithContext(context.Background())
}

func (i LbAclMap) ToLbAclMapOutputWithContext(ctx context.Context) LbAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbAclMapOutput)
}

type LbAclOutput struct{ *pulumi.OutputState }

func (LbAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbAcl)(nil)).Elem()
}

func (o LbAclOutput) ToLbAclOutput() LbAclOutput {
	return o
}

func (o LbAclOutput) ToLbAclOutputWithContext(ctx context.Context) LbAclOutput {
	return o
}

// Action to undertake when an ACL filter matches.
func (o LbAclOutput) Action() LbAclActionOutput {
	return o.ApplyT(func(v *LbAcl) LbAclActionOutput { return v.Action }).(LbAclActionOutput)
}

// IsDate and time of ACL's creation (RFC 3339 format)
func (o LbAclOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAcl) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The ACL description.
func (o LbAclOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbAcl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The load-balancer Frontend ID to attach the ACL to.
func (o LbAclOutput) FrontendId() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAcl) pulumi.StringOutput { return v.FrontendId }).(pulumi.StringOutput)
}

// The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
func (o LbAclOutput) Index() pulumi.IntOutput {
	return o.ApplyT(func(v *LbAcl) pulumi.IntOutput { return v.Index }).(pulumi.IntOutput)
}

// The ACL match rule. At least `ipSubnet` or `httpFilter` and `httpFilterValue` are required.
func (o LbAclOutput) Match() LbAclMatchPtrOutput {
	return o.ApplyT(func(v *LbAcl) LbAclMatchPtrOutput { return v.Match }).(LbAclMatchPtrOutput)
}

// The ACL name. If not provided it will be randomly generated.
func (o LbAclOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAcl) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IsDate and time of ACL's update (RFC 3339 format)
func (o LbAclOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LbAcl) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type LbAclArrayOutput struct{ *pulumi.OutputState }

func (LbAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbAcl)(nil)).Elem()
}

func (o LbAclArrayOutput) ToLbAclArrayOutput() LbAclArrayOutput {
	return o
}

func (o LbAclArrayOutput) ToLbAclArrayOutputWithContext(ctx context.Context) LbAclArrayOutput {
	return o
}

func (o LbAclArrayOutput) Index(i pulumi.IntInput) LbAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbAcl {
		return vs[0].([]*LbAcl)[vs[1].(int)]
	}).(LbAclOutput)
}

type LbAclMapOutput struct{ *pulumi.OutputState }

func (LbAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbAcl)(nil)).Elem()
}

func (o LbAclMapOutput) ToLbAclMapOutput() LbAclMapOutput {
	return o
}

func (o LbAclMapOutput) ToLbAclMapOutputWithContext(ctx context.Context) LbAclMapOutput {
	return o
}

func (o LbAclMapOutput) MapIndex(k pulumi.StringInput) LbAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbAcl {
		return vs[0].(map[string]*LbAcl)[vs[1].(string)]
	}).(LbAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbAclInput)(nil)).Elem(), &LbAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbAclArrayInput)(nil)).Elem(), LbAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbAclMapInput)(nil)).Elem(), LbAclMap{})
	pulumi.RegisterOutputType(LbAclOutput{})
	pulumi.RegisterOutputType(LbAclArrayOutput{})
	pulumi.RegisterOutputType(LbAclMapOutput{})
}
