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

// Creates and manages Scaleway Database Instances.
// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
//
// ## Example Usage
//
// ### Example Basic
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
//			_, err := scaleway.NewDocumentdbInstance(ctx, "main", &scaleway.DocumentdbInstanceArgs{
//				Engine:   pulumi.String("FerretDB-1"),
//				NodeType: pulumi.String("docdb-play2-pico"),
//				Password: pulumi.String("thiZ_is_v&ry_s3cret"),
//				Tags: pulumi.StringArray{
//					pulumi.String("terraform-test"),
//					pulumi.String("scaleway_documentdb_instance"),
//					pulumi.String("minimal"),
//				},
//				UserName:       pulumi.String("my_initial_user"),
//				VolumeSizeInGb: pulumi.Int(20),
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
// Database Instance can be imported using the `{region}/{id}`, e.g.
//
// bash
//
// ```sh
// $ pulumi import scaleway:index/documentdbInstance:DocumentdbInstance db fr-par/11111111-1111-1111-1111-111111111111
// ```
type DocumentdbInstance struct {
	pulumi.CustomResourceState

	// Database Instance's engine version (e.g. `FerretDB-1`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Enable or disable high availability for the database instance.
	IsHaCluster pulumi.BoolPtrOutput `pulumi:"isHaCluster"`
	// The name of the Database Instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of database instance you want to create (e.g. `docdb-play2-pico`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// Password for the first user of the database instance.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	TelemetryEnabled pulumi.BoolPtrOutput `pulumi:"telemetryEnabled"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	VolumeSizeInGb pulumi.IntOutput `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewDocumentdbInstance registers a new resource with the given unique name, arguments, and options.
func NewDocumentdbInstance(ctx *pulumi.Context,
	name string, args *DocumentdbInstanceArgs, opts ...pulumi.ResourceOption) (*DocumentdbInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DocumentdbInstance
	err := ctx.RegisterResource("scaleway:index/documentdbInstance:DocumentdbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentdbInstance gets an existing DocumentdbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentdbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentdbInstanceState, opts ...pulumi.ResourceOption) (*DocumentdbInstance, error) {
	var resource DocumentdbInstance
	err := ctx.ReadResource("scaleway:index/documentdbInstance:DocumentdbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentdbInstance resources.
type documentdbInstanceState struct {
	// Database Instance's engine version (e.g. `FerretDB-1`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine *string `pulumi:"engine"`
	// Enable or disable high availability for the database instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `docdb-play2-pico`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	NodeType *string `pulumi:"nodeType"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	TelemetryEnabled *bool `pulumi:"telemetryEnabled"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType *string `pulumi:"volumeType"`
}

type DocumentdbInstanceState struct {
	// Database Instance's engine version (e.g. `FerretDB-1`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringPtrInput
	// Enable or disable high availability for the database instance.
	IsHaCluster pulumi.BoolPtrInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `docdb-play2-pico`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	NodeType pulumi.StringPtrInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	TelemetryEnabled pulumi.BoolPtrInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrInput
}

func (DocumentdbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbInstanceState)(nil)).Elem()
}

type documentdbInstanceArgs struct {
	// Database Instance's engine version (e.g. `FerretDB-1`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine string `pulumi:"engine"`
	// Enable or disable high availability for the database instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `docdb-play2-pico`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	NodeType string `pulumi:"nodeType"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	TelemetryEnabled *bool `pulumi:"telemetryEnabled"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a DocumentdbInstance resource.
type DocumentdbInstanceArgs struct {
	// Database Instance's engine version (e.g. `FerretDB-1`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringInput
	// Enable or disable high availability for the database instance.
	IsHaCluster pulumi.BoolPtrInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `docdb-play2-pico`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	NodeType pulumi.StringInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	TelemetryEnabled pulumi.BoolPtrInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrInput
}

func (DocumentdbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentdbInstanceArgs)(nil)).Elem()
}

type DocumentdbInstanceInput interface {
	pulumi.Input

	ToDocumentdbInstanceOutput() DocumentdbInstanceOutput
	ToDocumentdbInstanceOutputWithContext(ctx context.Context) DocumentdbInstanceOutput
}

func (*DocumentdbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbInstance)(nil)).Elem()
}

func (i *DocumentdbInstance) ToDocumentdbInstanceOutput() DocumentdbInstanceOutput {
	return i.ToDocumentdbInstanceOutputWithContext(context.Background())
}

func (i *DocumentdbInstance) ToDocumentdbInstanceOutputWithContext(ctx context.Context) DocumentdbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbInstanceOutput)
}

// DocumentdbInstanceArrayInput is an input type that accepts DocumentdbInstanceArray and DocumentdbInstanceArrayOutput values.
// You can construct a concrete instance of `DocumentdbInstanceArrayInput` via:
//
//	DocumentdbInstanceArray{ DocumentdbInstanceArgs{...} }
type DocumentdbInstanceArrayInput interface {
	pulumi.Input

	ToDocumentdbInstanceArrayOutput() DocumentdbInstanceArrayOutput
	ToDocumentdbInstanceArrayOutputWithContext(context.Context) DocumentdbInstanceArrayOutput
}

type DocumentdbInstanceArray []DocumentdbInstanceInput

func (DocumentdbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbInstance)(nil)).Elem()
}

func (i DocumentdbInstanceArray) ToDocumentdbInstanceArrayOutput() DocumentdbInstanceArrayOutput {
	return i.ToDocumentdbInstanceArrayOutputWithContext(context.Background())
}

func (i DocumentdbInstanceArray) ToDocumentdbInstanceArrayOutputWithContext(ctx context.Context) DocumentdbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbInstanceArrayOutput)
}

// DocumentdbInstanceMapInput is an input type that accepts DocumentdbInstanceMap and DocumentdbInstanceMapOutput values.
// You can construct a concrete instance of `DocumentdbInstanceMapInput` via:
//
//	DocumentdbInstanceMap{ "key": DocumentdbInstanceArgs{...} }
type DocumentdbInstanceMapInput interface {
	pulumi.Input

	ToDocumentdbInstanceMapOutput() DocumentdbInstanceMapOutput
	ToDocumentdbInstanceMapOutputWithContext(context.Context) DocumentdbInstanceMapOutput
}

type DocumentdbInstanceMap map[string]DocumentdbInstanceInput

func (DocumentdbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbInstance)(nil)).Elem()
}

func (i DocumentdbInstanceMap) ToDocumentdbInstanceMapOutput() DocumentdbInstanceMapOutput {
	return i.ToDocumentdbInstanceMapOutputWithContext(context.Background())
}

func (i DocumentdbInstanceMap) ToDocumentdbInstanceMapOutputWithContext(ctx context.Context) DocumentdbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentdbInstanceMapOutput)
}

type DocumentdbInstanceOutput struct{ *pulumi.OutputState }

func (DocumentdbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentdbInstance)(nil)).Elem()
}

func (o DocumentdbInstanceOutput) ToDocumentdbInstanceOutput() DocumentdbInstanceOutput {
	return o
}

func (o DocumentdbInstanceOutput) ToDocumentdbInstanceOutputWithContext(ctx context.Context) DocumentdbInstanceOutput {
	return o
}

// Database Instance's engine version (e.g. `FerretDB-1`).
//
// > **Important:** Updates to `engine` will recreate the Database Instance.
func (o DocumentdbInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Enable or disable high availability for the database instance.
func (o DocumentdbInstanceOutput) IsHaCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.BoolPtrOutput { return v.IsHaCluster }).(pulumi.BoolPtrOutput)
}

// The name of the Database Instance.
func (o DocumentdbInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of database instance you want to create (e.g. `docdb-play2-pico`).
//
// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
// interruption. Keep in mind that you cannot downgrade a Database Instance.
func (o DocumentdbInstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// Password for the first user of the database instance.
func (o DocumentdbInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// `projectId`) The ID of the project the Database
// Instance is associated with.
func (o DocumentdbInstanceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// `region`) The region
// in which the Database Instance should be created.
func (o DocumentdbInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The tags associated with the Database Instance.
func (o DocumentdbInstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Enable telemetry to collects basic anonymous usage data and sends them to FerretDB telemetry service. More about the telemetry [here](https://docs.ferretdb.io/telemetry/#configure-telemetry).
//
// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
func (o DocumentdbInstanceOutput) TelemetryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.BoolPtrOutput { return v.TelemetryEnabled }).(pulumi.BoolPtrOutput)
}

// Identifier for the first user of the database instance.
//
// > **Important:** Updates to `userName` will recreate the Database Instance.
func (o DocumentdbInstanceOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// Volume size (in GB) when `volumeType` is set to `bssd`.
func (o DocumentdbInstanceOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.IntOutput { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// Type of volume where data are stored (`bssd` or `lssd`).
func (o DocumentdbInstanceOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentdbInstance) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type DocumentdbInstanceArrayOutput struct{ *pulumi.OutputState }

func (DocumentdbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentdbInstance)(nil)).Elem()
}

func (o DocumentdbInstanceArrayOutput) ToDocumentdbInstanceArrayOutput() DocumentdbInstanceArrayOutput {
	return o
}

func (o DocumentdbInstanceArrayOutput) ToDocumentdbInstanceArrayOutputWithContext(ctx context.Context) DocumentdbInstanceArrayOutput {
	return o
}

func (o DocumentdbInstanceArrayOutput) Index(i pulumi.IntInput) DocumentdbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentdbInstance {
		return vs[0].([]*DocumentdbInstance)[vs[1].(int)]
	}).(DocumentdbInstanceOutput)
}

type DocumentdbInstanceMapOutput struct{ *pulumi.OutputState }

func (DocumentdbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentdbInstance)(nil)).Elem()
}

func (o DocumentdbInstanceMapOutput) ToDocumentdbInstanceMapOutput() DocumentdbInstanceMapOutput {
	return o
}

func (o DocumentdbInstanceMapOutput) ToDocumentdbInstanceMapOutputWithContext(ctx context.Context) DocumentdbInstanceMapOutput {
	return o
}

func (o DocumentdbInstanceMapOutput) MapIndex(k pulumi.StringInput) DocumentdbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentdbInstance {
		return vs[0].(map[string]*DocumentdbInstance)[vs[1].(string)]
	}).(DocumentdbInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbInstanceInput)(nil)).Elem(), &DocumentdbInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbInstanceArrayInput)(nil)).Elem(), DocumentdbInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentdbInstanceMapInput)(nil)).Elem(), DocumentdbInstanceMap{})
	pulumi.RegisterOutputType(DocumentdbInstanceOutput{})
	pulumi.RegisterOutputType(DocumentdbInstanceArrayOutput{})
	pulumi.RegisterOutputType(DocumentdbInstanceMapOutput{})
}
