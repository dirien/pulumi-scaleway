// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package scaleway

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/dirien/pulumi-scaleway/sdk/v2/go/scaleway/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "scaleway:index/accountProject:AccountProject":
		r = &AccountProject{}
	case "scaleway:index/accountSshKey:AccountSshKey":
		r = &AccountSshKey{}
	case "scaleway:index/appleSiliconServer:AppleSiliconServer":
		r = &AppleSiliconServer{}
	case "scaleway:index/baremetalServer:BaremetalServer":
		r = &BaremetalServer{}
	case "scaleway:index/blockSnapshot:BlockSnapshot":
		r = &BlockSnapshot{}
	case "scaleway:index/blockVolume:BlockVolume":
		r = &BlockVolume{}
	case "scaleway:index/cockpit:Cockpit":
		r = &Cockpit{}
	case "scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser":
		r = &CockpitGrafanaUser{}
	case "scaleway:index/cockpitToken:CockpitToken":
		r = &CockpitToken{}
	case "scaleway:index/container:Container":
		r = &Container{}
	case "scaleway:index/containerCron:ContainerCron":
		r = &ContainerCron{}
	case "scaleway:index/containerDomain:ContainerDomain":
		r = &ContainerDomain{}
	case "scaleway:index/containerNamespace:ContainerNamespace":
		r = &ContainerNamespace{}
	case "scaleway:index/containerToken:ContainerToken":
		r = &ContainerToken{}
	case "scaleway:index/containerTrigger:ContainerTrigger":
		r = &ContainerTrigger{}
	case "scaleway:index/documentdbDatabase:DocumentdbDatabase":
		r = &DocumentdbDatabase{}
	case "scaleway:index/documentdbInstance:DocumentdbInstance":
		r = &DocumentdbInstance{}
	case "scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint":
		r = &DocumentdbPrivateNetworkEndpoint{}
	case "scaleway:index/documentdbPrivilege:DocumentdbPrivilege":
		r = &DocumentdbPrivilege{}
	case "scaleway:index/documentdbReadReplica:DocumentdbReadReplica":
		r = &DocumentdbReadReplica{}
	case "scaleway:index/documentdbUser:DocumentdbUser":
		r = &DocumentdbUser{}
	case "scaleway:index/domainRecord:DomainRecord":
		r = &DomainRecord{}
	case "scaleway:index/domainZone:DomainZone":
		r = &DomainZone{}
	case "scaleway:index/flexibleIp:FlexibleIp":
		r = &FlexibleIp{}
	case "scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress":
		r = &FlexibleIpMacAddress{}
	case "scaleway:index/function:Function":
		r = &Function{}
	case "scaleway:index/functionCron:FunctionCron":
		r = &FunctionCron{}
	case "scaleway:index/functionDomain:FunctionDomain":
		r = &FunctionDomain{}
	case "scaleway:index/functionNamespace:FunctionNamespace":
		r = &FunctionNamespace{}
	case "scaleway:index/functionToken:FunctionToken":
		r = &FunctionToken{}
	case "scaleway:index/functionTrigger:FunctionTrigger":
		r = &FunctionTrigger{}
	case "scaleway:index/iamApiKey:IamApiKey":
		r = &IamApiKey{}
	case "scaleway:index/iamApplication:IamApplication":
		r = &IamApplication{}
	case "scaleway:index/iamGroup:IamGroup":
		r = &IamGroup{}
	case "scaleway:index/iamGroupMembership:IamGroupMembership":
		r = &IamGroupMembership{}
	case "scaleway:index/iamPolicy:IamPolicy":
		r = &IamPolicy{}
	case "scaleway:index/iamSshKey:IamSshKey":
		r = &IamSshKey{}
	case "scaleway:index/iamUser:IamUser":
		r = &IamUser{}
	case "scaleway:index/instanceImage:InstanceImage":
		r = &InstanceImage{}
	case "scaleway:index/instanceIp:InstanceIp":
		r = &InstanceIp{}
	case "scaleway:index/instanceIpReverseDns:InstanceIpReverseDns":
		r = &InstanceIpReverseDns{}
	case "scaleway:index/instancePlacementGroup:InstancePlacementGroup":
		r = &InstancePlacementGroup{}
	case "scaleway:index/instancePrivateNic:InstancePrivateNic":
		r = &InstancePrivateNic{}
	case "scaleway:index/instanceSecurityGroup:InstanceSecurityGroup":
		r = &InstanceSecurityGroup{}
	case "scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules":
		r = &InstanceSecurityGroupRules{}
	case "scaleway:index/instanceServer:InstanceServer":
		r = &InstanceServer{}
	case "scaleway:index/instanceSnapshot:InstanceSnapshot":
		r = &InstanceSnapshot{}
	case "scaleway:index/instanceUserData:InstanceUserData":
		r = &InstanceUserData{}
	case "scaleway:index/instanceVolume:InstanceVolume":
		r = &InstanceVolume{}
	case "scaleway:index/iotDevice:IotDevice":
		r = &IotDevice{}
	case "scaleway:index/iotHub:IotHub":
		r = &IotHub{}
	case "scaleway:index/iotNetwork:IotNetwork":
		r = &IotNetwork{}
	case "scaleway:index/iotRoute:IotRoute":
		r = &IotRoute{}
	case "scaleway:index/ipamIp:IpamIp":
		r = &IpamIp{}
	case "scaleway:index/jobDefinition:JobDefinition":
		r = &JobDefinition{}
	case "scaleway:index/k8sCluster:K8sCluster":
		r = &K8sCluster{}
	case "scaleway:index/k8sPool:K8sPool":
		r = &K8sPool{}
	case "scaleway:index/lb:Lb":
		r = &Lb{}
	case "scaleway:index/lbAcl:LbAcl":
		r = &LbAcl{}
	case "scaleway:index/lbBackend:LbBackend":
		r = &LbBackend{}
	case "scaleway:index/lbCertificate:LbCertificate":
		r = &LbCertificate{}
	case "scaleway:index/lbFrontend:LbFrontend":
		r = &LbFrontend{}
	case "scaleway:index/lbIp:LbIp":
		r = &LbIp{}
	case "scaleway:index/lbRoute:LbRoute":
		r = &LbRoute{}
	case "scaleway:index/mnqNatsAccount:MnqNatsAccount":
		r = &MnqNatsAccount{}
	case "scaleway:index/mnqNatsCredentials:MnqNatsCredentials":
		r = &MnqNatsCredentials{}
	case "scaleway:index/mnqSns:MnqSns":
		r = &MnqSns{}
	case "scaleway:index/mnqSnsCredentials:MnqSnsCredentials":
		r = &MnqSnsCredentials{}
	case "scaleway:index/mnqSnsTopic:MnqSnsTopic":
		r = &MnqSnsTopic{}
	case "scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription":
		r = &MnqSnsTopicSubscription{}
	case "scaleway:index/mnqSqs:MnqSqs":
		r = &MnqSqs{}
	case "scaleway:index/mnqSqsCredentials:MnqSqsCredentials":
		r = &MnqSqsCredentials{}
	case "scaleway:index/mnqSqsQueue:MnqSqsQueue":
		r = &MnqSqsQueue{}
	case "scaleway:index/objectBucket:ObjectBucket":
		r = &ObjectBucket{}
	case "scaleway:index/objectBucketAcl:ObjectBucketAcl":
		r = &ObjectBucketAcl{}
	case "scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration":
		r = &ObjectBucketLockConfiguration{}
	case "scaleway:index/objectBucketPolicy:ObjectBucketPolicy":
		r = &ObjectBucketPolicy{}
	case "scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration":
		r = &ObjectBucketWebsiteConfiguration{}
	case "scaleway:index/objectItem:ObjectItem":
		r = &ObjectItem{}
	case "scaleway:index/rdbAcl:RdbAcl":
		r = &RdbAcl{}
	case "scaleway:index/rdbDatabase:RdbDatabase":
		r = &RdbDatabase{}
	case "scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup":
		r = &RdbDatabaseBackup{}
	case "scaleway:index/rdbInstance:RdbInstance":
		r = &RdbInstance{}
	case "scaleway:index/rdbPrivilege:RdbPrivilege":
		r = &RdbPrivilege{}
	case "scaleway:index/rdbReadReplica:RdbReadReplica":
		r = &RdbReadReplica{}
	case "scaleway:index/rdbUser:RdbUser":
		r = &RdbUser{}
	case "scaleway:index/redisCluster:RedisCluster":
		r = &RedisCluster{}
	case "scaleway:index/registryNamespace:RegistryNamespace":
		r = &RegistryNamespace{}
	case "scaleway:index/sdbSqlDatabase:SdbSqlDatabase":
		r = &SdbSqlDatabase{}
	case "scaleway:index/secret:Secret":
		r = &Secret{}
	case "scaleway:index/secretVersion:SecretVersion":
		r = &SecretVersion{}
	case "scaleway:index/temDomain:TemDomain":
		r = &TemDomain{}
	case "scaleway:index/vpc:Vpc":
		r = &Vpc{}
	case "scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork":
		r = &VpcGatewayNetwork{}
	case "scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork":
		r = &VpcPrivateNetwork{}
	case "scaleway:index/vpcPublicGateway:VpcPublicGateway":
		r = &VpcPublicGateway{}
	case "scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp":
		r = &VpcPublicGatewayDhcp{}
	case "scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation":
		r = &VpcPublicGatewayDhcpReservation{}
	case "scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp":
		r = &VpcPublicGatewayIp{}
	case "scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns":
		r = &VpcPublicGatewayIpReverseDns{}
	case "scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule":
		r = &VpcPublicGatewayPatRule{}
	case "scaleway:index/webhosting:Webhosting":
		r = &Webhosting{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:scaleway" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/accountProject",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/accountSshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/appleSiliconServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/baremetalServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/blockSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/blockVolume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/cockpit",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/cockpitGrafanaUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/cockpitToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/container",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerCron",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/containerTrigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/documentdbDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/documentdbInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/documentdbPrivateNetworkEndpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/documentdbPrivilege",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/documentdbReadReplica",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/documentdbUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/domainRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/domainZone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/flexibleIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/flexibleIpMacAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionCron",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/functionTrigger",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamApiKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamApplication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamGroupMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamSshKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iamUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceImage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceIpReverseDns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instancePlacementGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instancePrivateNic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceSecurityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceSecurityGroupRules",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceSnapshot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceUserData",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/instanceVolume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotDevice",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotHub",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/iotRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/ipamIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/jobDefinition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/k8sCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/k8sPool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lbAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lbBackend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lbCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lbFrontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lbIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/lbRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqNatsAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqNatsCredentials",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSnsCredentials",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSnsTopic",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSnsTopicSubscription",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSqs",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSqsCredentials",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/mnqSqsQueue",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketLockConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectBucketWebsiteConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/objectItem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbAcl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbDatabaseBackup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbPrivilege",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbReadReplica",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/rdbUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/redisCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/registryNamespace",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/sdbSqlDatabase",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/secret",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/secretVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/temDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpc",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcGatewayNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPrivateNetwork",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayDhcp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayDhcpReservation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayIp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayIpReverseDns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/vpcPublicGatewayPatRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"scaleway",
		"index/webhosting",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"scaleway",
		&pkg{version},
	)
}
