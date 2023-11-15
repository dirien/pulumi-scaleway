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

// ## Import
//
// Database Instance can be imported using the `{region}/{id}`, e.g. bash
//
// ```sh
//
//	$ pulumi import scaleway:index/rdbInstance:RdbInstance rdb01 fr-par/11111111-1111-1111-1111-111111111111
//
// ```
type RdbInstance struct {
	pulumi.CustomResourceState

	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolOutput `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntOutput `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntOutput `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance.
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrOutput `pulumi:"disableBackup"`
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringOutput `pulumi:"endpointIp"`
	// (Deprecated) The port of the Database Instance.
	EndpointPort pulumi.IntOutput `pulumi:"endpointPort"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	InitSettings pulumi.StringMapOutput `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrOutput `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance.
	LoadBalancers RdbInstanceLoadBalancerArrayOutput `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The organization ID the Database Instance is associated with.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Password for the first user of the database instance.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork RdbInstancePrivateNetworkPtrOutput `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// List of read replicas of the database instance.
	ReadReplicas RdbInstanceReadReplicaArrayOutput `pulumi:"readReplicas"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapOutput `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntOutput `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewRdbInstance registers a new resource with the given unique name, arguments, and options.
func NewRdbInstance(ctx *pulumi.Context,
	name string, args *RdbInstanceArgs, opts ...pulumi.ResourceOption) (*RdbInstance, error) {
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
	var resource RdbInstance
	err := ctx.RegisterResource("scaleway:index/rdbInstance:RdbInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdbInstance gets an existing RdbInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdbInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdbInstanceState, opts ...pulumi.ResourceOption) (*RdbInstance, error) {
	var resource RdbInstance
	err := ctx.ReadResource("scaleway:index/rdbInstance:RdbInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdbInstance resources.
type rdbInstanceState struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Certificate of the database instance.
	Certificate *string `pulumi:"certificate"`
	// Disable automated backup for the database instance.
	DisableBackup *bool `pulumi:"disableBackup"`
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp *string `pulumi:"endpointIp"`
	// (Deprecated) The port of the Database Instance.
	EndpointPort *int `pulumi:"endpointPort"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine *string `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	InitSettings map[string]string `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// List of load balancer endpoints of the database instance.
	LoadBalancers []RdbInstanceLoadBalancer `pulumi:"loadBalancers"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType *string `pulumi:"nodeType"`
	// The organization ID the Database Instance is associated with.
	OrganizationId *string `pulumi:"organizationId"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork *RdbInstancePrivateNetwork `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// List of read replicas of the database instance.
	ReadReplicas []RdbInstanceReadReplica `pulumi:"readReplicas"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings map[string]string `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType *string `pulumi:"volumeType"`
}

type RdbInstanceState struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntPtrInput
	// Certificate of the database instance.
	Certificate pulumi.StringPtrInput
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrInput
	// (Deprecated) The IP of the Database Instance.
	//
	// Deprecated: Please use the private_network or the load_balancer attribute
	EndpointIp pulumi.StringPtrInput
	// (Deprecated) The port of the Database Instance.
	EndpointPort pulumi.IntPtrInput
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringPtrInput
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	InitSettings pulumi.StringMapInput
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrInput
	// List of load balancer endpoints of the database instance.
	LoadBalancers RdbInstanceLoadBalancerArrayInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringPtrInput
	// The organization ID the Database Instance is associated with.
	OrganizationId pulumi.StringPtrInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// List of private networks endpoints of the database instance.
	PrivateNetwork RdbInstancePrivateNetworkPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// List of read replicas of the database instance.
	ReadReplicas RdbInstanceReadReplicaArrayInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrInput
}

func (RdbInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbInstanceState)(nil)).Elem()
}

type rdbInstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion *bool `pulumi:"backupSameRegion"`
	// Backup schedule frequency in hours.
	BackupScheduleFrequency *int `pulumi:"backupScheduleFrequency"`
	// Backup schedule retention in days.
	BackupScheduleRetention *int `pulumi:"backupScheduleRetention"`
	// Disable automated backup for the database instance.
	DisableBackup *bool `pulumi:"disableBackup"`
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine string `pulumi:"engine"`
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	InitSettings map[string]string `pulumi:"initSettings"`
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster *bool `pulumi:"isHaCluster"`
	// The name of the Database Instance.
	Name *string `pulumi:"name"`
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType string `pulumi:"nodeType"`
	// Password for the first user of the database instance.
	Password *string `pulumi:"password"`
	// List of private networks endpoints of the database instance.
	PrivateNetwork *RdbInstancePrivateNetwork `pulumi:"privateNetwork"`
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId *string `pulumi:"projectId"`
	// `region`) The region
	// in which the Database Instance should be created.
	Region *string `pulumi:"region"`
	// Map of engine settings to be set. Using this option will override default config.
	Settings map[string]string `pulumi:"settings"`
	// The tags associated with the Database Instance.
	Tags []string `pulumi:"tags"`
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName *string `pulumi:"userName"`
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb *int `pulumi:"volumeSizeInGb"`
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a RdbInstance resource.
type RdbInstanceArgs struct {
	// Boolean to store logical backups in the same region as the database instance.
	BackupSameRegion pulumi.BoolPtrInput
	// Backup schedule frequency in hours.
	BackupScheduleFrequency pulumi.IntPtrInput
	// Backup schedule retention in days.
	BackupScheduleRetention pulumi.IntPtrInput
	// Disable automated backup for the database instance.
	DisableBackup pulumi.BoolPtrInput
	// Database Instance's engine version (e.g. `PostgreSQL-11`).
	//
	// > **Important:** Updates to `engine` will recreate the Database Instance.
	Engine pulumi.StringInput
	// Map of engine settings to be set at database initialisation.
	//
	// > **Important:** Updates to `initSettings` will recreate the Database Instance.
	InitSettings pulumi.StringMapInput
	// Enable or disable high availability for the database instance.
	//
	// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
	IsHaCluster pulumi.BoolPtrInput
	// The name of the Database Instance.
	Name pulumi.StringPtrInput
	// The type of database instance you want to create (e.g. `db-dev-s`).
	//
	// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
	// interruption. Keep in mind that you cannot downgrade a Database Instance.
	//
	// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
	// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
	NodeType pulumi.StringInput
	// Password for the first user of the database instance.
	Password pulumi.StringPtrInput
	// List of private networks endpoints of the database instance.
	PrivateNetwork RdbInstancePrivateNetworkPtrInput
	// `projectId`) The ID of the project the Database
	// Instance is associated with.
	ProjectId pulumi.StringPtrInput
	// `region`) The region
	// in which the Database Instance should be created.
	Region pulumi.StringPtrInput
	// Map of engine settings to be set. Using this option will override default config.
	Settings pulumi.StringMapInput
	// The tags associated with the Database Instance.
	Tags pulumi.StringArrayInput
	// Identifier for the first user of the database instance.
	//
	// > **Important:** Updates to `userName` will recreate the Database Instance.
	UserName pulumi.StringPtrInput
	// Volume size (in GB) when `volumeType` is set to `bssd`.
	//
	// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
	VolumeSizeInGb pulumi.IntPtrInput
	// Type of volume where data are stored (`bssd` or `lssd`).
	VolumeType pulumi.StringPtrInput
}

func (RdbInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdbInstanceArgs)(nil)).Elem()
}

type RdbInstanceInput interface {
	pulumi.Input

	ToRdbInstanceOutput() RdbInstanceOutput
	ToRdbInstanceOutputWithContext(ctx context.Context) RdbInstanceOutput
}

func (*RdbInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbInstance)(nil)).Elem()
}

func (i *RdbInstance) ToRdbInstanceOutput() RdbInstanceOutput {
	return i.ToRdbInstanceOutputWithContext(context.Background())
}

func (i *RdbInstance) ToRdbInstanceOutputWithContext(ctx context.Context) RdbInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbInstanceOutput)
}

// RdbInstanceArrayInput is an input type that accepts RdbInstanceArray and RdbInstanceArrayOutput values.
// You can construct a concrete instance of `RdbInstanceArrayInput` via:
//
//	RdbInstanceArray{ RdbInstanceArgs{...} }
type RdbInstanceArrayInput interface {
	pulumi.Input

	ToRdbInstanceArrayOutput() RdbInstanceArrayOutput
	ToRdbInstanceArrayOutputWithContext(context.Context) RdbInstanceArrayOutput
}

type RdbInstanceArray []RdbInstanceInput

func (RdbInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbInstance)(nil)).Elem()
}

func (i RdbInstanceArray) ToRdbInstanceArrayOutput() RdbInstanceArrayOutput {
	return i.ToRdbInstanceArrayOutputWithContext(context.Background())
}

func (i RdbInstanceArray) ToRdbInstanceArrayOutputWithContext(ctx context.Context) RdbInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbInstanceArrayOutput)
}

// RdbInstanceMapInput is an input type that accepts RdbInstanceMap and RdbInstanceMapOutput values.
// You can construct a concrete instance of `RdbInstanceMapInput` via:
//
//	RdbInstanceMap{ "key": RdbInstanceArgs{...} }
type RdbInstanceMapInput interface {
	pulumi.Input

	ToRdbInstanceMapOutput() RdbInstanceMapOutput
	ToRdbInstanceMapOutputWithContext(context.Context) RdbInstanceMapOutput
}

type RdbInstanceMap map[string]RdbInstanceInput

func (RdbInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbInstance)(nil)).Elem()
}

func (i RdbInstanceMap) ToRdbInstanceMapOutput() RdbInstanceMapOutput {
	return i.ToRdbInstanceMapOutputWithContext(context.Background())
}

func (i RdbInstanceMap) ToRdbInstanceMapOutputWithContext(ctx context.Context) RdbInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdbInstanceMapOutput)
}

type RdbInstanceOutput struct{ *pulumi.OutputState }

func (RdbInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdbInstance)(nil)).Elem()
}

func (o RdbInstanceOutput) ToRdbInstanceOutput() RdbInstanceOutput {
	return o
}

func (o RdbInstanceOutput) ToRdbInstanceOutputWithContext(ctx context.Context) RdbInstanceOutput {
	return o
}

// Boolean to store logical backups in the same region as the database instance.
func (o RdbInstanceOutput) BackupSameRegion() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.BoolOutput { return v.BackupSameRegion }).(pulumi.BoolOutput)
}

// Backup schedule frequency in hours.
func (o RdbInstanceOutput) BackupScheduleFrequency() pulumi.IntOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.IntOutput { return v.BackupScheduleFrequency }).(pulumi.IntOutput)
}

// Backup schedule retention in days.
func (o RdbInstanceOutput) BackupScheduleRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.IntOutput { return v.BackupScheduleRetention }).(pulumi.IntOutput)
}

// Certificate of the database instance.
func (o RdbInstanceOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.Certificate }).(pulumi.StringOutput)
}

// Disable automated backup for the database instance.
func (o RdbInstanceOutput) DisableBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.BoolPtrOutput { return v.DisableBackup }).(pulumi.BoolPtrOutput)
}

// (Deprecated) The IP of the Database Instance.
//
// Deprecated: Please use the private_network or the load_balancer attribute
func (o RdbInstanceOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.EndpointIp }).(pulumi.StringOutput)
}

// (Deprecated) The port of the Database Instance.
func (o RdbInstanceOutput) EndpointPort() pulumi.IntOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.IntOutput { return v.EndpointPort }).(pulumi.IntOutput)
}

// Database Instance's engine version (e.g. `PostgreSQL-11`).
//
// > **Important:** Updates to `engine` will recreate the Database Instance.
func (o RdbInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Map of engine settings to be set at database initialisation.
//
// > **Important:** Updates to `initSettings` will recreate the Database Instance.
func (o RdbInstanceOutput) InitSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringMapOutput { return v.InitSettings }).(pulumi.StringMapOutput)
}

// Enable or disable high availability for the database instance.
//
// > **Important:** Updates to `isHaCluster` will recreate the Database Instance.
func (o RdbInstanceOutput) IsHaCluster() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.BoolPtrOutput { return v.IsHaCluster }).(pulumi.BoolPtrOutput)
}

// List of load balancer endpoints of the database instance.
func (o RdbInstanceOutput) LoadBalancers() RdbInstanceLoadBalancerArrayOutput {
	return o.ApplyT(func(v *RdbInstance) RdbInstanceLoadBalancerArrayOutput { return v.LoadBalancers }).(RdbInstanceLoadBalancerArrayOutput)
}

// The name of the Database Instance.
func (o RdbInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The type of database instance you want to create (e.g. `db-dev-s`).
//
// > **Important:** Updates to `nodeType` will upgrade the Database Instance to the desired `nodeType` without any
// interruption. Keep in mind that you cannot downgrade a Database Instance.
//
// > **Important:** Once your instance reaches `diskFull` status, if you are using `lssd` storage, you should upgrade the node_type,
// and if you are using `bssd` storage, you should increase the volume size before making any other change to your instance.
func (o RdbInstanceOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// The organization ID the Database Instance is associated with.
func (o RdbInstanceOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Password for the first user of the database instance.
func (o RdbInstanceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// List of private networks endpoints of the database instance.
func (o RdbInstanceOutput) PrivateNetwork() RdbInstancePrivateNetworkPtrOutput {
	return o.ApplyT(func(v *RdbInstance) RdbInstancePrivateNetworkPtrOutput { return v.PrivateNetwork }).(RdbInstancePrivateNetworkPtrOutput)
}

// `projectId`) The ID of the project the Database
// Instance is associated with.
func (o RdbInstanceOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// List of read replicas of the database instance.
func (o RdbInstanceOutput) ReadReplicas() RdbInstanceReadReplicaArrayOutput {
	return o.ApplyT(func(v *RdbInstance) RdbInstanceReadReplicaArrayOutput { return v.ReadReplicas }).(RdbInstanceReadReplicaArrayOutput)
}

// `region`) The region
// in which the Database Instance should be created.
func (o RdbInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Map of engine settings to be set. Using this option will override default config.
func (o RdbInstanceOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringMapOutput { return v.Settings }).(pulumi.StringMapOutput)
}

// The tags associated with the Database Instance.
func (o RdbInstanceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Identifier for the first user of the database instance.
//
// > **Important:** Updates to `userName` will recreate the Database Instance.
func (o RdbInstanceOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

// Volume size (in GB) when `volumeType` is set to `bssd`.
//
// > **Important:** Once your instance reaches `diskFull` status, you should increase the volume size before making any other change to your instance.
func (o RdbInstanceOutput) VolumeSizeInGb() pulumi.IntOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.IntOutput { return v.VolumeSizeInGb }).(pulumi.IntOutput)
}

// Type of volume where data are stored (`bssd` or `lssd`).
func (o RdbInstanceOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RdbInstance) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type RdbInstanceArrayOutput struct{ *pulumi.OutputState }

func (RdbInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdbInstance)(nil)).Elem()
}

func (o RdbInstanceArrayOutput) ToRdbInstanceArrayOutput() RdbInstanceArrayOutput {
	return o
}

func (o RdbInstanceArrayOutput) ToRdbInstanceArrayOutputWithContext(ctx context.Context) RdbInstanceArrayOutput {
	return o
}

func (o RdbInstanceArrayOutput) Index(i pulumi.IntInput) RdbInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdbInstance {
		return vs[0].([]*RdbInstance)[vs[1].(int)]
	}).(RdbInstanceOutput)
}

type RdbInstanceMapOutput struct{ *pulumi.OutputState }

func (RdbInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdbInstance)(nil)).Elem()
}

func (o RdbInstanceMapOutput) ToRdbInstanceMapOutput() RdbInstanceMapOutput {
	return o
}

func (o RdbInstanceMapOutput) ToRdbInstanceMapOutputWithContext(ctx context.Context) RdbInstanceMapOutput {
	return o
}

func (o RdbInstanceMapOutput) MapIndex(k pulumi.StringInput) RdbInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdbInstance {
		return vs[0].(map[string]*RdbInstance)[vs[1].(string)]
	}).(RdbInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdbInstanceInput)(nil)).Elem(), &RdbInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbInstanceArrayInput)(nil)).Elem(), RdbInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdbInstanceMapInput)(nil)).Elem(), RdbInstanceMap{})
	pulumi.RegisterOutputType(RdbInstanceOutput{})
	pulumi.RegisterOutputType(RdbInstanceArrayOutput{})
	pulumi.RegisterOutputType(RdbInstanceMapOutput{})
}
