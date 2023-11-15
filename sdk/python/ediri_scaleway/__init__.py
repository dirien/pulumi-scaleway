# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .account_project import *
from .account_ssh_key import *
from .apple_silicon_server import *
from .baremetal_server import *
from .cockpit import *
from .cockpit_grafana_user import *
from .cockpit_token import *
from .container import *
from .container_cron import *
from .container_domain import *
from .container_namespace import *
from .container_token import *
from .container_trigger import *
from .document_db_database import *
from .document_db_instance import *
from .document_db_private_network_endpoint import *
from .document_db_privilege import *
from .document_db_read_replica import *
from .document_dbuser import *
from .domain_record import *
from .domain_zone import *
from .flexible_ip import *
from .flexible_ip_mac_address import *
from .function import *
from .function_cron import *
from .function_domain import *
from .function_namespace import *
from .function_token import *
from .function_trigger import *
from .get_account_project import *
from .get_account_ssh_key import *
from .get_availability_zones import *
from .get_baremetal_offer import *
from .get_baremetal_option import *
from .get_baremetal_os import *
from .get_baremetal_server import *
from .get_billing_consumptions import *
from .get_billing_invoices import *
from .get_cockpit import *
from .get_cockpit_plan import *
from .get_container import *
from .get_container_namespace import *
from .get_document_db_database import *
from .get_document_db_instance import *
from .get_document_db_load_balancer_endpoint import *
from .get_domain_record import *
from .get_domain_zone import *
from .get_flexible_ip import *
from .get_flexible_ips import *
from .get_function import *
from .get_function_namespace import *
from .get_iam_application import *
from .get_iam_group import *
from .get_iam_ssh_key import *
from .get_iam_user import *
from .get_instance_image import *
from .get_instance_ip import *
from .get_instance_private_nic import *
from .get_instance_security_group import *
from .get_instance_server import *
from .get_instance_servers import *
from .get_instance_snapshot import *
from .get_instance_volume import *
from .get_iot_device import *
from .get_iot_hub import *
from .get_ipam_ip import *
from .get_k8s_cluster import *
from .get_k8s_pool import *
from .get_k8s_version import *
from .get_lb import *
from .get_lb_acls import *
from .get_lb_backend import *
from .get_lb_backends import *
from .get_lb_certificate import *
from .get_lb_frontend import *
from .get_lb_frontends import *
from .get_lb_ip import *
from .get_lb_ips import *
from .get_lb_route import *
from .get_lb_routes import *
from .get_lbs import *
from .get_marketplace_image import *
from .get_mnq_sqs import *
from .get_object_bucket import *
from .get_object_bucket_policy import *
from .get_rdb_acl import *
from .get_rdb_database import *
from .get_rdb_database_backup import *
from .get_rdb_instance import *
from .get_rdb_privilege import *
from .get_redis_cluster import *
from .get_registry_image import *
from .get_registry_namespace import *
from .get_secret import *
from .get_secret_version import *
from .get_tem_domain import *
from .get_vpc import *
from .get_vpc_gateway_network import *
from .get_vpc_private_network import *
from .get_vpc_public_gateway import *
from .get_vpc_public_gateway_dhcp import *
from .get_vpc_public_gateway_dhcp_reservation import *
from .get_vpc_public_gateway_ip import *
from .get_vpc_public_gateway_pat_rule import *
from .get_vpcs import *
from .get_web_host_offer import *
from .get_webhosting import *
from .iam_api_key import *
from .iam_application import *
from .iam_group import *
from .iam_group_membership import *
from .iam_ip import *
from .iam_policy import *
from .iam_ssh_key import *
from .iam_user import *
from .instance_image import *
from .instance_ip import *
from .instance_ip_reverse_dns import *
from .instance_placement_group import *
from .instance_private_nic import *
from .instance_security_group import *
from .instance_security_group_rules import *
from .instance_server import *
from .instance_snapshot import *
from .instance_user_data import *
from .instance_volume import *
from .iot_device import *
from .iot_hub import *
from .iot_network import *
from .iot_route import *
from .ipam_ip import *
from .k8s_cluster import *
from .k8s_pool import *
from .lb import *
from .lb_acl import *
from .lb_backend import *
from .lb_certificate import *
from .lb_frontend import *
from .lb_ip import *
from .lb_route import *
from .mnq_credential import *
from .mnq_namespace import *
from .mnq_nats_account import *
from .mnq_nats_credentials import *
from .mnq_queue import *
from .mnq_sqs import *
from .mnq_sqs_credentials import *
from .mnq_sqs_queue import *
from .object_bucket import *
from .object_bucket_acl import *
from .object_bucket_lock_configuration import *
from .object_bucket_policy import *
from .object_bucket_website_configuration import *
from .object_item import *
from .provider import *
from .rdb_acl import *
from .rdb_database import *
from .rdb_database_backup import *
from .rdb_instance import *
from .rdb_privilege import *
from .rdb_read_replica import *
from .rdb_user import *
from .redis_cluster import *
from .registry_namespace import *
from .secret import *
from .secret_version import *
from .tem_domain import *
from .vpc import *
from .vpc_gateway_network import *
from .vpc_private_network import *
from .vpc_public_gateway import *
from .vpc_public_gateway_dhcp import *
from .vpc_public_gateway_dhcp_reservation import *
from .vpc_public_gateway_ip import *
from .vpc_public_gateway_ip_reverse_dns import *
from .vpc_public_gateway_pat_rule import *
from .web_hosting import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import ediri_scaleway.config as __config
    config = __config
else:
    config = _utilities.lazy_import('ediri_scaleway.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "scaleway",
  "mod": "index/accountProject",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/accountProject:AccountProject": "AccountProject"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/accountSshKey",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/accountSshKey:AccountSshKey": "AccountSshKey"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/appleSiliconServer",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/appleSiliconServer:AppleSiliconServer": "AppleSiliconServer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/baremetalServer",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/baremetalServer:BaremetalServer": "BaremetalServer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/cockpit",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/cockpit:Cockpit": "Cockpit"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/cockpitGrafanaUser",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser": "CockpitGrafanaUser"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/cockpitToken",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/cockpitToken:CockpitToken": "CockpitToken"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/container",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/container:Container": "Container"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerCron",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/containerCron:ContainerCron": "ContainerCron"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerDomain",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/containerDomain:ContainerDomain": "ContainerDomain"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerNamespace",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/containerNamespace:ContainerNamespace": "ContainerNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerToken",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/containerToken:ContainerToken": "ContainerToken"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/containerTrigger",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/containerTrigger:ContainerTrigger": "ContainerTrigger"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/documentDBDatabase",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/documentDBDatabase:DocumentDBDatabase": "DocumentDBDatabase"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/documentDBInstance",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/documentDBInstance:DocumentDBInstance": "DocumentDBInstance"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/documentDBPrivateNetworkEndpoint",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/documentDBPrivateNetworkEndpoint:DocumentDBPrivateNetworkEndpoint": "DocumentDBPrivateNetworkEndpoint"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/documentDBPrivilege",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/documentDBPrivilege:DocumentDBPrivilege": "DocumentDBPrivilege"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/documentDBReadReplica",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/documentDBReadReplica:DocumentDBReadReplica": "DocumentDBReadReplica"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/documentDBUser",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/documentDBUser:DocumentDBUser": "DocumentDBUser"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/domainRecord",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/domainRecord:DomainRecord": "DomainRecord"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/domainZone",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/domainZone:DomainZone": "DomainZone"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/flexibleIp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/flexibleIp:FlexibleIp": "FlexibleIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/flexibleIpMacAddress",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress": "FlexibleIpMacAddress"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/function",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/function:Function": "Function"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionCron",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/functionCron:FunctionCron": "FunctionCron"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionDomain",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/functionDomain:FunctionDomain": "FunctionDomain"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionNamespace",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/functionNamespace:FunctionNamespace": "FunctionNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionToken",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/functionToken:FunctionToken": "FunctionToken"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/functionTrigger",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/functionTrigger:FunctionTrigger": "FunctionTrigger"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamApiKey",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamApiKey:IamApiKey": "IamApiKey"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamApplication",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamApplication:IamApplication": "IamApplication"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamGroup",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamGroup:IamGroup": "IamGroup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamGroupMembership",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamGroupMembership:IamGroupMembership": "IamGroupMembership"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamIp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamIp:IamIp": "IamIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamPolicy",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamPolicy:IamPolicy": "IamPolicy"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamSshKey",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamSshKey:IamSshKey": "IamSshKey"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iamUser",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iamUser:IamUser": "IamUser"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceImage",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceImage:InstanceImage": "InstanceImage"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceIp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceIp:InstanceIp": "InstanceIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceIpReverseDns",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceIpReverseDns:InstanceIpReverseDns": "InstanceIpReverseDns"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instancePlacementGroup",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instancePlacementGroup:InstancePlacementGroup": "InstancePlacementGroup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instancePrivateNic",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instancePrivateNic:InstancePrivateNic": "InstancePrivateNic"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceSecurityGroup",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceSecurityGroup:InstanceSecurityGroup": "InstanceSecurityGroup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceSecurityGroupRules",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules": "InstanceSecurityGroupRules"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceServer",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceServer:InstanceServer": "InstanceServer"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceSnapshot",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceSnapshot:InstanceSnapshot": "InstanceSnapshot"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceUserData",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceUserData:InstanceUserData": "InstanceUserData"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/instanceVolume",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/instanceVolume:InstanceVolume": "InstanceVolume"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotDevice",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iotDevice:IotDevice": "IotDevice"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotHub",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iotHub:IotHub": "IotHub"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotNetwork",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iotNetwork:IotNetwork": "IotNetwork"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/iotRoute",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/iotRoute:IotRoute": "IotRoute"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/ipamIp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/ipamIp:IpamIp": "IpamIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/k8sCluster",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/k8sCluster:K8sCluster": "K8sCluster"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/k8sPool",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/k8sPool:K8sPool": "K8sPool"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lb",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lb:Lb": "Lb"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lbAcl",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lbAcl:LbAcl": "LbAcl"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lbBackend",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lbBackend:LbBackend": "LbBackend"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lbCertificate",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lbCertificate:LbCertificate": "LbCertificate"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lbFrontend",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lbFrontend:LbFrontend": "LbFrontend"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lbIp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lbIp:LbIp": "LbIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/lbRoute",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/lbRoute:LbRoute": "LbRoute"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqCredential",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqCredential:MnqCredential": "MnqCredential"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqNamespace",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqNamespace:MnqNamespace": "MnqNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqNatsAccount",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqNatsAccount:MnqNatsAccount": "MnqNatsAccount"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqNatsCredentials",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqNatsCredentials:MnqNatsCredentials": "MnqNatsCredentials"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqQueue",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqQueue:MnqQueue": "MnqQueue"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqSqs",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqSqs:MnqSqs": "MnqSqs"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqSqsCredentials",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqSqsCredentials:MnqSqsCredentials": "MnqSqsCredentials"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/mnqSqsQueue",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/mnqSqsQueue:MnqSqsQueue": "MnqSqsQueue"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucket",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/objectBucket:ObjectBucket": "ObjectBucket"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketAcl",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/objectBucketAcl:ObjectBucketAcl": "ObjectBucketAcl"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketLockConfiguration",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/objectBucketLockConfiguration:ObjectBucketLockConfiguration": "ObjectBucketLockConfiguration"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketPolicy",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/objectBucketPolicy:ObjectBucketPolicy": "ObjectBucketPolicy"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectBucketWebsiteConfiguration",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration": "ObjectBucketWebsiteConfiguration"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/objectItem",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/objectItem:ObjectItem": "ObjectItem"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbAcl",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbAcl:RdbAcl": "RdbAcl"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbDatabase",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbDatabase:RdbDatabase": "RdbDatabase"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbDatabaseBackup",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup": "RdbDatabaseBackup"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbInstance",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbInstance:RdbInstance": "RdbInstance"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbPrivilege",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbPrivilege:RdbPrivilege": "RdbPrivilege"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbReadReplica",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbReadReplica:RdbReadReplica": "RdbReadReplica"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/rdbUser",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/rdbUser:RdbUser": "RdbUser"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/redisCluster",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/redisCluster:RedisCluster": "RedisCluster"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/registryNamespace",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/registryNamespace:RegistryNamespace": "RegistryNamespace"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/secret",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/secret:Secret": "Secret"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/secretVersion",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/secretVersion:SecretVersion": "SecretVersion"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/temDomain",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/temDomain:TemDomain": "TemDomain"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpc",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpc:Vpc": "Vpc"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcGatewayNetwork",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcGatewayNetwork:VpcGatewayNetwork": "VpcGatewayNetwork"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPrivateNetwork",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPrivateNetwork:VpcPrivateNetwork": "VpcPrivateNetwork"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGateway",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGateway:VpcPublicGateway": "VpcPublicGateway"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayDhcp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayDhcp:VpcPublicGatewayDhcp": "VpcPublicGatewayDhcp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayDhcpReservation",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayDhcpReservation:VpcPublicGatewayDhcpReservation": "VpcPublicGatewayDhcpReservation"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayIp",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp": "VpcPublicGatewayIp"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayIpReverseDns",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns": "VpcPublicGatewayIpReverseDns"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/vpcPublicGatewayPatRule",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule": "VpcPublicGatewayPatRule"
  }
 },
 {
  "pkg": "scaleway",
  "mod": "index/webHosting",
  "fqn": "ediri_scaleway",
  "classes": {
   "scaleway:index/webHosting:WebHosting": "WebHosting"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "scaleway",
  "token": "pulumi:providers:scaleway",
  "fqn": "ediri_scaleway",
  "class": "Provider"
 }
]
"""
)
