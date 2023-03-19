// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and manages Scaleway Database read replicas.
// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
//			instance, err := scaleway.NewRdbInstance(ctx, "instance", &scaleway.RdbInstanceArgs{
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform-test"),
//					pulumi.String("scaleway_rdb_read_replica"),
//					pulumi.String("minimal"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewRdbReadReplica(ctx, "replica", &scaleway.RdbReadReplicaArgs{
//				InstanceId:   instance.ID(),
//				DirectAccess: nil,
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
// ### Private network
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
//			instance, err := scaleway.NewRdbInstance(ctx, "instance", &scaleway.RdbInstanceArgs{
//				NodeType:      pulumi.String("db-dev-s"),
//				Engine:        pulumi.String("PostgreSQL-14"),
//				IsHaCluster:   pulumi.Bool(false),
//				DisableBackup: pulumi.Bool(true),
//				UserName:      pulumi.String("my_initial_user"),
//				Password:      pulumi.String("thiZ_is_v&ry_s3cret"),
//			})
//			if err != nil {
//				return err
//			}
//			pn, err := scaleway.NewVpcPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewRdbReadReplica(ctx, "replica", &scaleway.RdbReadReplicaArgs{
//				InstanceId: instance.ID(),
//				PrivateNetwork: &scaleway.RdbReadReplicaPrivateNetworkArgs{
//					PrivateNetworkId: pn.ID(),
//					ServiceIp:        pulumi.String("192.168.1.254/24"),
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
// Database Read replica can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/rdbReadReplica:RdbReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type RdbReadReplica struct {
	pulumi.CustomResourceState

	// Creates a direct access endpoint to rdb replica.
	DirectAccess RdbReadReplicaDirectAccessPtrOutput `pulumi:"directAccess"`
	// UUID of the rdb instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork RdbReadReplicaPrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `region`) The region in which the Database read replica should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewRdbReadReplica registers a new resource with the given unique name, arguments, and options.
func NewRdbReadReplica(ctx *pulumi.Context,
	name string, args *RdbReadReplicaArgs, opts ...pulumi.ResourceOption) (*RdbReadReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RdbReadReplica
	err := ctx.RegisterResource("scaleway:index/rdbReadReplica:RdbReadReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbReadReplica gets an existing RdbReadReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbReadReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbReadReplicaState, opts ...pulumi.ResourceOption) (*RdbReadReplica, error) {
	var resource RdbReadReplica
	err := ctx.ReadResource("scaleway:index/rdbReadReplica:RdbReadReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbReadReplica resources.
type rdbReadReplicaState struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess *RdbReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the rdb instance.
	InstanceId *string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *RdbReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

type RdbReadReplicaState struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess RdbReadReplicaDirectAccessPtrInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringPtrInput
	// Create an endpoint in a private network.
	PrivateNetwork RdbReadReplicaPrivateNetworkPtrInput
	// `region`) The region in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (RdbReadReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbReadReplicaState)(nil)).Elem()
}

type rdbReadReplicaArgs struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess *RdbReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the rdb instance.
	InstanceId string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *RdbReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a RdbReadReplica resource.
type RdbReadReplicaArgs struct {
	// Creates a direct access endpoint to rdb replica.
	DirectAccess RdbReadReplicaDirectAccessPtrInput
	// UUID of the rdb instance.
	InstanceId pulumi.StringInput
	// Create an endpoint in a private network.
	PrivateNetwork RdbReadReplicaPrivateNetworkPtrInput
	// `region`) The region in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (RdbReadReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbReadReplicaArgs)(nil)).Elem()
}

type RdbReadReplicaInput interface {
	pulumi.Input

	ToRdbReadReplicaOutput() RdbReadReplicaOutput
	ToRdbReadReplicaOutputWithContext(ctx context.Context) RdbReadReplicaOutput
}

func (*RdbReadReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbReadReplica)(nil)).Elem()
}

func (i *RdbReadReplica) ToRdbReadReplicaOutput() RdbReadReplicaOutput {
	return i.ToRdbReadReplicaOutputWithContext(context.Background())
}

func (i *RdbReadReplica) ToRdbReadReplicaOutputWithContext(ctx context.Context) RdbReadReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbReadReplicaOutput)
}

// RdbReadReplicaArrayInput is an input type that accepts RdbReadReplicaArray and RdbReadReplicaArrayOutput values.
// You can construct a concrete instance of `RdbReadReplicaArrayInput` via:
//
//	RdbReadReplicaArray{ RdbReadReplicaArgs{...} }
type RdbReadReplicaArrayInput interface {
	pulumi.Input

	ToRdbReadReplicaArrayOutput() RdbReadReplicaArrayOutput
	ToRdbReadReplicaArrayOutputWithContext(context.Context) RdbReadReplicaArrayOutput
}

type RdbReadReplicaArray []RdbReadReplicaInput

func (RdbReadReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbReadReplica)(nil)).Elem()
}

func (i RdbReadReplicaArray) ToRdbReadReplicaArrayOutput() RdbReadReplicaArrayOutput {
	return i.ToRdbReadReplicaArrayOutputWithContext(context.Background())
}

func (i RdbReadReplicaArray) ToRdbReadReplicaArrayOutputWithContext(ctx context.Context) RdbReadReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbReadReplicaArrayOutput)
}

// RdbReadReplicaMapInput is an input type that accepts RdbReadReplicaMap and RdbReadReplicaMapOutput values.
// You can construct a concrete instance of `RdbReadReplicaMapInput` via:
//
//	RdbReadReplicaMap{ "key": RdbReadReplicaArgs{...} }
type RdbReadReplicaMapInput interface {
	pulumi.Input

	ToRdbReadReplicaMapOutput() RdbReadReplicaMapOutput
	ToRdbReadReplicaMapOutputWithContext(context.Context) RdbReadReplicaMapOutput
}

type RdbReadReplicaMap map[string]RdbReadReplicaInput

func (RdbReadReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbReadReplica)(nil)).Elem()
}

func (i RdbReadReplicaMap) ToRdbReadReplicaMapOutput() RdbReadReplicaMapOutput {
	return i.ToRdbReadReplicaMapOutputWithContext(context.Background())
}

func (i RdbReadReplicaMap) ToRdbReadReplicaMapOutputWithContext(ctx context.Context) RdbReadReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbReadReplicaMapOutput)
}

type RdbReadReplicaOutput struct{ *pulumi.OutputState }

func (RdbReadReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbReadReplica)(nil)).Elem()
}

func (o RdbReadReplicaOutput) ToRdbReadReplicaOutput() RdbReadReplicaOutput {
	return o
}

func (o RdbReadReplicaOutput) ToRdbReadReplicaOutputWithContext(ctx context.Context) RdbReadReplicaOutput {
	return o
}

// Creates a direct access endpoint to rdb replica.
func (o RdbReadReplicaOutput) DirectAccess() RdbReadReplicaDirectAccessPtrOutput {
	return o.ApplyT(func(v *RdbReadReplica) RdbReadReplicaDirectAccessPtrOutput { return v.DirectAccess }).(RdbReadReplicaDirectAccessPtrOutput)
}

// UUID of the rdb instance.
func (o RdbReadReplicaOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbReadReplica) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Create an endpoint in a private network.
func (o RdbReadReplicaOutput) PrivateNetwork() RdbReadReplicaPrivateNetworkPtrOutput {
	return o.ApplyT(func(v *RdbReadReplica) RdbReadReplicaPrivateNetworkPtrOutput { return v.PrivateNetwork }).(RdbReadReplicaPrivateNetworkPtrOutput)
}

// `region`) The region in which the Database read replica should be created.
func (o RdbReadReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbReadReplica) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type RdbReadReplicaArrayOutput struct{ *pulumi.OutputState }

func (RdbReadReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbReadReplica)(nil)).Elem()
}

func (o RdbReadReplicaArrayOutput) ToRdbReadReplicaArrayOutput() RdbReadReplicaArrayOutput {
	return o
}

func (o RdbReadReplicaArrayOutput) ToRdbReadReplicaArrayOutputWithContext(ctx context.Context) RdbReadReplicaArrayOutput {
	return o
}

func (o RdbReadReplicaArrayOutput) Index(i pulumi.IntInput) RdbReadReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdbReadReplica {
		return vs[0].([]*RdbReadReplica)[vs[1].(int)]
	}).(RdbReadReplicaOutput)
}

type RdbReadReplicaMapOutput struct{ *pulumi.OutputState }

func (RdbReadReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbReadReplica)(nil)).Elem()
}

func (o RdbReadReplicaMapOutput) ToRdbReadReplicaMapOutput() RdbReadReplicaMapOutput {
	return o
}

func (o RdbReadReplicaMapOutput) ToRdbReadReplicaMapOutputWithContext(ctx context.Context) RdbReadReplicaMapOutput {
	return o
}

func (o RdbReadReplicaMapOutput) MapIndex(k pulumi.StringInput) RdbReadReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdbReadReplica {
		return vs[0].(map[string]*RdbReadReplica)[vs[1].(string)]
	}).(RdbReadReplicaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdbReadReplicaInput)(nil)).Elem(), &RdbReadReplica{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbReadReplicaArrayInput)(nil)).Elem(), RdbReadReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbReadReplicaMapInput)(nil)).Elem(), RdbReadReplicaMap{})
	pulumi.RegisterOutputType(RdbReadReplicaOutput{})
	pulumi.RegisterOutputType(RdbReadReplicaArrayOutput{})
	pulumi.RegisterOutputType(RdbReadReplicaMapOutput{})
}
