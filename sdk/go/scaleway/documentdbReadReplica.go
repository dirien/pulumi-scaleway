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

// Creates and manages Scaleway DocumentDB Database read replicas.
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
//			_, err := scaleway.NewDocumentdbReadReplica(ctx, "replica", &scaleway.DocumentdbReadReplicaArgs{
//				DirectAccess: nil,
//				InstanceId:   pulumi.String("11111111-1111-1111-1111-111111111111"),
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
//			pn, err := scaleway.NewVpcPrivateNetwork(ctx, "pn", nil)
//			if err != nil {
//				return err
//			}
//			instance, err := scaleway.NewDocumentdbInstance(ctx, "instance", &scaleway.DocumentdbInstanceArgs{
//				NodeType:       pulumi.String("docdb-play2-pico"),
//				Engine:         pulumi.String("FerretDB-1"),
//				UserName:       pulumi.String("my_initial_user"),
//				Password:       pulumi.String("thiZ_is_v&ry_s3cret"),
//				VolumeSizeInGb: pulumi.Int(20),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = scaleway.NewDocumentdbReadReplica(ctx, "replica", &scaleway.DocumentdbReadReplicaArgs{
//				InstanceId: instance.ID(),
//				PrivateNetwork: &scaleway.DocumentdbReadReplicaPrivateNetworkArgs{
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
// Database Read replica can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/documentdbReadReplica:DocumentdbReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
// ```
type DocumentdbReadReplica struct {
	pulumi.CustomResourceState

	// Creates a direct access endpoint to documentdb replica.
	DirectAccess DocumentdbReadReplicaDirectAccessPtrOutput `pulumi:"directAccess"`
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork DocumentdbReadReplicaPrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewDocumentdbReadReplica registers a new resource with the given unique name, arguments, and options.
func NewDocumentdbReadReplica(ctx *pulumi.Context,
	name string, args *DocumentdbReadReplicaArgs, opts ...pulumi.ResourceOption) (*DocumentdbReadReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentdbReadReplica
	err := ctx.RegisterResource("scaleway:index/documentdbReadReplica:DocumentdbReadReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentdbReadReplica gets an existing DocumentdbReadReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentdbReadReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentdbReadReplicaState, opts ...pulumi.ResourceOption) (*DocumentdbReadReplica, error) {
	var resource DocumentdbReadReplica
	err := ctx.ReadResource("scaleway:index/documentdbReadReplica:DocumentdbReadReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentdbReadReplica resources.
type documentdbReadReplicaState struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess *DocumentdbReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId *string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *DocumentdbReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

type DocumentdbReadReplicaState struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess DocumentdbReadReplicaDirectAccessPtrInput
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringPtrInput
	// Create an endpoint in a private network.
	PrivateNetwork DocumentdbReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (DocumentdbReadReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbReadReplicaState)(nil)).Elem()
}

type documentdbReadReplicaArgs struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess *DocumentdbReadReplicaDirectAccess `pulumi:"directAccess"`
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId string `pulumi:"instanceId"`
	// Create an endpoint in a private network.
	PrivateNetwork *DocumentdbReadReplicaPrivateNetwork `pulumi:"privateNetwork"`
	// `region`) The region
	// in which the Database read replica should be created.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a DocumentdbReadReplica resource.
type DocumentdbReadReplicaArgs struct {
	// Creates a direct access endpoint to documentdb replica.
	DirectAccess DocumentdbReadReplicaDirectAccessPtrInput
	// UUID of the documentdb instance.
	//
	// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
	InstanceId pulumi.StringInput
	// Create an endpoint in a private network.
	PrivateNetwork DocumentdbReadReplicaPrivateNetworkPtrInput
	// `region`) The region
	// in which the Database read replica should be created.
	Region pulumi.StringPtrInput
}

func (DocumentdbReadReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbReadReplicaArgs)(nil)).Elem()
}

type DocumentdbReadReplicaInput interface {
	pulumi.Input

	ToDocumentdbReadReplicaOutput() DocumentdbReadReplicaOutput
	ToDocumentdbReadReplicaOutputWithContext(ctx context.Context) DocumentdbReadReplicaOutput
}

func (*DocumentdbReadReplica) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbReadReplica)(nil)).Elem()
}

func (i *DocumentdbReadReplica) ToDocumentdbReadReplicaOutput() DocumentdbReadReplicaOutput {
	return i.ToDocumentdbReadReplicaOutputWithContext(context.Background())
}

func (i *DocumentdbReadReplica) ToDocumentdbReadReplicaOutputWithContext(ctx context.Context) DocumentdbReadReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbReadReplicaOutput)
}

// DocumentdbReadReplicaArrayInput is an input type that accepts DocumentdbReadReplicaArray and DocumentdbReadReplicaArrayOutput values.
// You can construct a concrete instance of `DocumentdbReadReplicaArrayInput` via:
//
//	DocumentdbReadReplicaArray{ DocumentdbReadReplicaArgs{...} }
type DocumentdbReadReplicaArrayInput interface {
	pulumi.Input

	ToDocumentdbReadReplicaArrayOutput() DocumentdbReadReplicaArrayOutput
	ToDocumentdbReadReplicaArrayOutputWithContext(context.Context) DocumentdbReadReplicaArrayOutput
}

type DocumentdbReadReplicaArray []DocumentdbReadReplicaInput

func (DocumentdbReadReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbReadReplica)(nil)).Elem()
}

func (i DocumentdbReadReplicaArray) ToDocumentdbReadReplicaArrayOutput() DocumentdbReadReplicaArrayOutput {
	return i.ToDocumentdbReadReplicaArrayOutputWithContext(context.Background())
}

func (i DocumentdbReadReplicaArray) ToDocumentdbReadReplicaArrayOutputWithContext(ctx context.Context) DocumentdbReadReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbReadReplicaArrayOutput)
}

// DocumentdbReadReplicaMapInput is an input type that accepts DocumentdbReadReplicaMap and DocumentdbReadReplicaMapOutput values.
// You can construct a concrete instance of `DocumentdbReadReplicaMapInput` via:
//
//	DocumentdbReadReplicaMap{ "key": DocumentdbReadReplicaArgs{...} }
type DocumentdbReadReplicaMapInput interface {
	pulumi.Input

	ToDocumentdbReadReplicaMapOutput() DocumentdbReadReplicaMapOutput
	ToDocumentdbReadReplicaMapOutputWithContext(context.Context) DocumentdbReadReplicaMapOutput
}

type DocumentdbReadReplicaMap map[string]DocumentdbReadReplicaInput

func (DocumentdbReadReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbReadReplica)(nil)).Elem()
}

func (i DocumentdbReadReplicaMap) ToDocumentdbReadReplicaMapOutput() DocumentdbReadReplicaMapOutput {
	return i.ToDocumentdbReadReplicaMapOutputWithContext(context.Background())
}

func (i DocumentdbReadReplicaMap) ToDocumentdbReadReplicaMapOutputWithContext(ctx context.Context) DocumentdbReadReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbReadReplicaMapOutput)
}

type DocumentdbReadReplicaOutput struct{ *pulumi.OutputState }

func (DocumentdbReadReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbReadReplica)(nil)).Elem()
}

func (o DocumentdbReadReplicaOutput) ToDocumentdbReadReplicaOutput() DocumentdbReadReplicaOutput {
	return o
}

func (o DocumentdbReadReplicaOutput) ToDocumentdbReadReplicaOutputWithContext(ctx context.Context) DocumentdbReadReplicaOutput {
	return o
}

// Creates a direct access endpoint to documentdb replica.
func (o DocumentdbReadReplicaOutput) DirectAccess() DocumentdbReadReplicaDirectAccessPtrOutput {
	return o.ApplyT(func(v *DocumentdbReadReplica) DocumentdbReadReplicaDirectAccessPtrOutput { return v.DirectAccess }).(DocumentdbReadReplicaDirectAccessPtrOutput)
}

// UUID of the documentdb instance.
//
// > **Important:** The replica musts contains at least one of `directAccess` or `privateNetwork`. It can contain both.
func (o DocumentdbReadReplicaOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbReadReplica) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Create an endpoint in a private network.
func (o DocumentdbReadReplicaOutput) PrivateNetwork() DocumentdbReadReplicaPrivateNetworkPtrOutput {
	return o.ApplyT(func(v *DocumentdbReadReplica) DocumentdbReadReplicaPrivateNetworkPtrOutput { return v.PrivateNetwork }).(DocumentdbReadReplicaPrivateNetworkPtrOutput)
}

// `region`) The region
// in which the Database read replica should be created.
func (o DocumentdbReadReplicaOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbReadReplica) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type DocumentdbReadReplicaArrayOutput struct{ *pulumi.OutputState }

func (DocumentdbReadReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbReadReplica)(nil)).Elem()
}

func (o DocumentdbReadReplicaArrayOutput) ToDocumentdbReadReplicaArrayOutput() DocumentdbReadReplicaArrayOutput {
	return o
}

func (o DocumentdbReadReplicaArrayOutput) ToDocumentdbReadReplicaArrayOutputWithContext(ctx context.Context) DocumentdbReadReplicaArrayOutput {
	return o
}

func (o DocumentdbReadReplicaArrayOutput) Index(i pulumi.IntInput) DocumentdbReadReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentdbReadReplica {
		return vs[0].([]*DocumentdbReadReplica)[vs[1].(int)]
	}).(DocumentdbReadReplicaOutput)
}

type DocumentdbReadReplicaMapOutput struct{ *pulumi.OutputState }

func (DocumentdbReadReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbReadReplica)(nil)).Elem()
}

func (o DocumentdbReadReplicaMapOutput) ToDocumentdbReadReplicaMapOutput() DocumentdbReadReplicaMapOutput {
	return o
}

func (o DocumentdbReadReplicaMapOutput) ToDocumentdbReadReplicaMapOutputWithContext(ctx context.Context) DocumentdbReadReplicaMapOutput {
	return o
}

func (o DocumentdbReadReplicaMapOutput) MapIndex(k pulumi.StringInput) DocumentdbReadReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentdbReadReplica {
		return vs[0].(map[string]*DocumentdbReadReplica)[vs[1].(string)]
	}).(DocumentdbReadReplicaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbReadReplicaInput)(nil)).Elem(), &DocumentdbReadReplica{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbReadReplicaArrayInput)(nil)).Elem(), DocumentdbReadReplicaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbReadReplicaMapInput)(nil)).Elem(), DocumentdbReadReplicaMap{})
	pulumi.RegisterOutputType(DocumentdbReadReplicaOutput{})
	pulumi.RegisterOutputType(DocumentdbReadReplicaArrayOutput{})
	pulumi.RegisterOutputType(DocumentdbReadReplicaMapOutput{})
}
