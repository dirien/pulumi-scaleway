// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Database instance authorized IPs.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#acl-rules-allowed-ips).
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
//			_, err := scaleway.NewRdbAcl(ctx, "main", &scaleway.RdbAclArgs{
//				InstanceId: pulumi.Any(scaleway_rdb_instance.Main.Id),
//				AclRules: scaleway.RdbAclAclRuleArray{
//					&scaleway.RdbAclAclRuleArgs{
//						Ip:          pulumi.String("1.2.3.4/32"),
//						Description: pulumi.String("foo"),
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
// Database Instance can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/rdbAcl:RdbAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type RdbAcl struct {
	pulumi.CustomResourceState

	// A list of ACLs (structure is described below)
	AclRules RdbAclAclRuleArrayOutput `pulumi:"aclRules"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database ACL.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRdbAcl registers a new resource with the given unique name, arguments, and options.
func NewRdbAcl(ctx *pulumi.Context,
	name string, args *RdbAclArgs, opts ...pulumi.ResourceOption) (*RdbAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AclRules == nil {
		return nil, errors.New("invalid value for required argument 'AclRules'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RdbAcl
	err := ctx.RegisterResource("scaleway:index/rdbAcl:RdbAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbAcl gets an existing RdbAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbAclState, opts ...pulumi.ResourceOption) (*RdbAcl, error) {
	var resource RdbAcl
	err := ctx.ReadResource("scaleway:index/rdbAcl:RdbAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbAcl resources.
type rdbAclState struct {
	// A list of ACLs (structure is described below)
	AclRules []RdbAclAclRule `pulumi:"aclRules"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database ACL.
	InstanceId *string `pulumi:"instanceId"`
	// `region`) The region in which the Database Instance should be created.
	Region *string `pulumi:"region"`
}

type RdbAclState struct {
	// A list of ACLs (structure is described below)
	AclRules RdbAclAclRuleArrayInput
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database ACL.
	InstanceId pulumi.StringPtrInput
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringPtrInput
}

func (RdbAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbAclState)(nil)).Elem()
}

type rdbAclArgs struct {
	// A list of ACLs (structure is described below)
	AclRules []RdbAclAclRule `pulumi:"aclRules"`
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database ACL.
	InstanceId string `pulumi:"instanceId"`
	// `region`) The region in which the Database Instance should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RdbAcl resource.
type RdbAclArgs struct {
	// A list of ACLs (structure is described below)
	AclRules RdbAclAclRuleArrayInput
	// UUID of the rdb instance.
	//
	// > **Important:** Updates to `instanceId` will recreate the Database ACL.
	InstanceId pulumi.StringInput
	// `region`) The region in which the Database Instance should be created.
	Region pulumi.StringPtrInput
}

func (RdbAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbAclArgs)(nil)).Elem()
}

type RdbAclInput interface {
	pulumi.Input

	ToRdbAclOutput() RdbAclOutput
	ToRdbAclOutputWithContext(ctx context.Context) RdbAclOutput
}

func (*RdbAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbAcl)(nil)).Elem()
}

func (i *RdbAcl) ToRdbAclOutput() RdbAclOutput {
	return i.ToRdbAclOutputWithContext(context.Background())
}

func (i *RdbAcl) ToRdbAclOutputWithContext(ctx context.Context) RdbAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbAclOutput)
}

// RdbAclArrayInput is an input type that accepts RdbAclArray and RdbAclArrayOutput values.
// You can construct a concrete instance of `RdbAclArrayInput` via:
//
//	RdbAclArray{ RdbAclArgs{...} }
type RdbAclArrayInput interface {
	pulumi.Input

	ToRdbAclArrayOutput() RdbAclArrayOutput
	ToRdbAclArrayOutputWithContext(context.Context) RdbAclArrayOutput
}

type RdbAclArray []RdbAclInput

func (RdbAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbAcl)(nil)).Elem()
}

func (i RdbAclArray) ToRdbAclArrayOutput() RdbAclArrayOutput {
	return i.ToRdbAclArrayOutputWithContext(context.Background())
}

func (i RdbAclArray) ToRdbAclArrayOutputWithContext(ctx context.Context) RdbAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbAclArrayOutput)
}

// RdbAclMapInput is an input type that accepts RdbAclMap and RdbAclMapOutput values.
// You can construct a concrete instance of `RdbAclMapInput` via:
//
//	RdbAclMap{ "key": RdbAclArgs{...} }
type RdbAclMapInput interface {
	pulumi.Input

	ToRdbAclMapOutput() RdbAclMapOutput
	ToRdbAclMapOutputWithContext(context.Context) RdbAclMapOutput
}

type RdbAclMap map[string]RdbAclInput

func (RdbAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbAcl)(nil)).Elem()
}

func (i RdbAclMap) ToRdbAclMapOutput() RdbAclMapOutput {
	return i.ToRdbAclMapOutputWithContext(context.Background())
}

func (i RdbAclMap) ToRdbAclMapOutputWithContext(ctx context.Context) RdbAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbAclMapOutput)
}

type RdbAclOutput struct{ *pulumi.OutputState }

func (RdbAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbAcl)(nil)).Elem()
}

func (o RdbAclOutput) ToRdbAclOutput() RdbAclOutput {
	return o
}

func (o RdbAclOutput) ToRdbAclOutputWithContext(ctx context.Context) RdbAclOutput {
	return o
}

// A list of ACLs (structure is described below)
func (o RdbAclOutput) AclRules() RdbAclAclRuleArrayOutput {
	return o.ApplyT(func(v *RdbAcl) RdbAclAclRuleArrayOutput { return v.AclRules }).(RdbAclAclRuleArrayOutput)
}

// UUID of the rdb instance.
//
// > **Important:** Updates to `instanceId` will recreate the Database ACL.
func (o RdbAclOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbAcl) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// `region`) The region in which the Database Instance should be created.
func (o RdbAclOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbAcl) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RdbAclArrayOutput struct{ *pulumi.OutputState }

func (RdbAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbAcl)(nil)).Elem()
}

func (o RdbAclArrayOutput) ToRdbAclArrayOutput() RdbAclArrayOutput {
	return o
}

func (o RdbAclArrayOutput) ToRdbAclArrayOutputWithContext(ctx context.Context) RdbAclArrayOutput {
	return o
}

func (o RdbAclArrayOutput) Index(i pulumi.IntInput) RdbAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdbAcl {
		return vs[0].([]*RdbAcl)[vs[1].(int)]
	}).(RdbAclOutput)
}

type RdbAclMapOutput struct{ *pulumi.OutputState }

func (RdbAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbAcl)(nil)).Elem()
}

func (o RdbAclMapOutput) ToRdbAclMapOutput() RdbAclMapOutput {
	return o
}

func (o RdbAclMapOutput) ToRdbAclMapOutputWithContext(ctx context.Context) RdbAclMapOutput {
	return o
}

func (o RdbAclMapOutput) MapIndex(k pulumi.StringInput) RdbAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdbAcl {
		return vs[0].(map[string]*RdbAcl)[vs[1].(string)]
	}).(RdbAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdbAclInput)(nil)).Elem(), &RdbAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbAclArrayInput)(nil)).Elem(), RdbAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbAclMapInput)(nil)).Elem(), RdbAclMap{})
	pulumi.RegisterOutputType(RdbAclOutput{})
	pulumi.RegisterOutputType(RdbAclArrayOutput{})
	pulumi.RegisterOutputType(RdbAclMapOutput{})
}
