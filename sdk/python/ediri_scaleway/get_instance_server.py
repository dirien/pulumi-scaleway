# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetInstanceServerResult',
    'AwaitableGetInstanceServerResult',
    'get_instance_server',
    'get_instance_server_output',
]

@pulumi.output_type
class GetInstanceServerResult:
    """
    A collection of values returned by getInstanceServer.
    """
    def __init__(__self__, additional_volume_ids=None, boot_type=None, bootscript_id=None, cloud_init=None, enable_dynamic_ip=None, enable_ipv6=None, id=None, image=None, ip_id=None, ip_ids=None, ipv6_address=None, ipv6_gateway=None, ipv6_prefix_length=None, name=None, organization_id=None, placement_group_id=None, placement_group_policy_respected=None, private_ip=None, private_networks=None, project_id=None, public_ip=None, public_ips=None, replace_on_type_change=None, root_volumes=None, routed_ip_enabled=None, security_group_id=None, server_id=None, state=None, tags=None, type=None, user_data=None, zone=None):
        if additional_volume_ids and not isinstance(additional_volume_ids, list):
            raise TypeError("Expected argument 'additional_volume_ids' to be a list")
        pulumi.set(__self__, "additional_volume_ids", additional_volume_ids)
        if boot_type and not isinstance(boot_type, str):
            raise TypeError("Expected argument 'boot_type' to be a str")
        pulumi.set(__self__, "boot_type", boot_type)
        if bootscript_id and not isinstance(bootscript_id, str):
            raise TypeError("Expected argument 'bootscript_id' to be a str")
        pulumi.set(__self__, "bootscript_id", bootscript_id)
        if cloud_init and not isinstance(cloud_init, str):
            raise TypeError("Expected argument 'cloud_init' to be a str")
        pulumi.set(__self__, "cloud_init", cloud_init)
        if enable_dynamic_ip and not isinstance(enable_dynamic_ip, bool):
            raise TypeError("Expected argument 'enable_dynamic_ip' to be a bool")
        pulumi.set(__self__, "enable_dynamic_ip", enable_dynamic_ip)
        if enable_ipv6 and not isinstance(enable_ipv6, bool):
            raise TypeError("Expected argument 'enable_ipv6' to be a bool")
        pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image and not isinstance(image, str):
            raise TypeError("Expected argument 'image' to be a str")
        pulumi.set(__self__, "image", image)
        if ip_id and not isinstance(ip_id, str):
            raise TypeError("Expected argument 'ip_id' to be a str")
        pulumi.set(__self__, "ip_id", ip_id)
        if ip_ids and not isinstance(ip_ids, list):
            raise TypeError("Expected argument 'ip_ids' to be a list")
        pulumi.set(__self__, "ip_ids", ip_ids)
        if ipv6_address and not isinstance(ipv6_address, str):
            raise TypeError("Expected argument 'ipv6_address' to be a str")
        pulumi.set(__self__, "ipv6_address", ipv6_address)
        if ipv6_gateway and not isinstance(ipv6_gateway, str):
            raise TypeError("Expected argument 'ipv6_gateway' to be a str")
        pulumi.set(__self__, "ipv6_gateway", ipv6_gateway)
        if ipv6_prefix_length and not isinstance(ipv6_prefix_length, int):
            raise TypeError("Expected argument 'ipv6_prefix_length' to be a int")
        pulumi.set(__self__, "ipv6_prefix_length", ipv6_prefix_length)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if placement_group_id and not isinstance(placement_group_id, str):
            raise TypeError("Expected argument 'placement_group_id' to be a str")
        pulumi.set(__self__, "placement_group_id", placement_group_id)
        if placement_group_policy_respected and not isinstance(placement_group_policy_respected, bool):
            raise TypeError("Expected argument 'placement_group_policy_respected' to be a bool")
        pulumi.set(__self__, "placement_group_policy_respected", placement_group_policy_respected)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if private_networks and not isinstance(private_networks, list):
            raise TypeError("Expected argument 'private_networks' to be a list")
        pulumi.set(__self__, "private_networks", private_networks)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if public_ips and not isinstance(public_ips, list):
            raise TypeError("Expected argument 'public_ips' to be a list")
        pulumi.set(__self__, "public_ips", public_ips)
        if replace_on_type_change and not isinstance(replace_on_type_change, bool):
            raise TypeError("Expected argument 'replace_on_type_change' to be a bool")
        pulumi.set(__self__, "replace_on_type_change", replace_on_type_change)
        if root_volumes and not isinstance(root_volumes, list):
            raise TypeError("Expected argument 'root_volumes' to be a list")
        pulumi.set(__self__, "root_volumes", root_volumes)
        if routed_ip_enabled and not isinstance(routed_ip_enabled, bool):
            raise TypeError("Expected argument 'routed_ip_enabled' to be a bool")
        pulumi.set(__self__, "routed_ip_enabled", routed_ip_enabled)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_data and not isinstance(user_data, dict):
            raise TypeError("Expected argument 'user_data' to be a dict")
        pulumi.set(__self__, "user_data", user_data)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="additionalVolumeIds")
    def additional_volume_ids(self) -> Sequence[str]:
        """
        The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
        attached to the server.
        """
        return pulumi.get(self, "additional_volume_ids")

    @property
    @pulumi.getter(name="bootType")
    def boot_type(self) -> str:
        return pulumi.get(self, "boot_type")

    @property
    @pulumi.getter(name="bootscriptId")
    def bootscript_id(self) -> str:
        return pulumi.get(self, "bootscript_id")

    @property
    @pulumi.getter(name="cloudInit")
    def cloud_init(self) -> str:
        """
        The cloud init script associated with this server.
        """
        return pulumi.get(self, "cloud_init")

    @property
    @pulumi.getter(name="enableDynamicIp")
    def enable_dynamic_ip(self) -> bool:
        """
        True if dynamic IP in enable on the server.
        """
        return pulumi.get(self, "enable_dynamic_ip")

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> bool:
        """
        Determines if IPv6 is enabled for the server.
        """
        return pulumi.get(self, "enable_ipv6")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def image(self) -> str:
        """
        The UUID and the label of the base image used by the server.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> str:
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter(name="ipIds")
    def ip_ids(self) -> Sequence[str]:
        return pulumi.get(self, "ip_ids")

    @property
    @pulumi.getter(name="ipv6Address")
    def ipv6_address(self) -> str:
        """
        The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
        """
        return pulumi.get(self, "ipv6_address")

    @property
    @pulumi.getter(name="ipv6Gateway")
    def ipv6_gateway(self) -> str:
        """
        The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
        """
        return pulumi.get(self, "ipv6_gateway")

    @property
    @pulumi.getter(name="ipv6PrefixLength")
    def ipv6_prefix_length(self) -> int:
        """
        The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
        """
        return pulumi.get(self, "ipv6_prefix_length")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The ID of the organization the server is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="placementGroupId")
    def placement_group_id(self) -> str:
        """
        The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
        """
        return pulumi.get(self, "placement_group_id")

    @property
    @pulumi.getter(name="placementGroupPolicyRespected")
    def placement_group_policy_respected(self) -> bool:
        """
        True when the placement group policy is respected.
        """
        return pulumi.get(self, "placement_group_policy_respected")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        The Scaleway internal IP address of the server.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Sequence['outputs.GetInstanceServerPrivateNetworkResult']:
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        The public IP address of the server.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="publicIps")
    def public_ips(self) -> Sequence['outputs.GetInstanceServerPublicIpResult']:
        """
        The list of public IPs of the server
        """
        return pulumi.get(self, "public_ips")

    @property
    @pulumi.getter(name="replaceOnTypeChange")
    def replace_on_type_change(self) -> bool:
        return pulumi.get(self, "replace_on_type_change")

    @property
    @pulumi.getter(name="rootVolumes")
    def root_volumes(self) -> Sequence['outputs.GetInstanceServerRootVolumeResult']:
        return pulumi.get(self, "root_volumes")

    @property
    @pulumi.getter(name="routedIpEnabled")
    def routed_ip_enabled(self) -> bool:
        """
        True if the server support routed ip only.
        """
        return pulumi.get(self, "routed_ip_enabled")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> str:
        """
        The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[str]:
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the server. Possible values are: `started`, `stopped` or `standby`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        The tags associated with the server.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The commercial type of the server.
        You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userData")
    def user_data(self) -> Mapping[str, str]:
        """
        The user data associated with the server.
        """
        return pulumi.get(self, "user_data")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetInstanceServerResult(GetInstanceServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceServerResult(
            additional_volume_ids=self.additional_volume_ids,
            boot_type=self.boot_type,
            bootscript_id=self.bootscript_id,
            cloud_init=self.cloud_init,
            enable_dynamic_ip=self.enable_dynamic_ip,
            enable_ipv6=self.enable_ipv6,
            id=self.id,
            image=self.image,
            ip_id=self.ip_id,
            ip_ids=self.ip_ids,
            ipv6_address=self.ipv6_address,
            ipv6_gateway=self.ipv6_gateway,
            ipv6_prefix_length=self.ipv6_prefix_length,
            name=self.name,
            organization_id=self.organization_id,
            placement_group_id=self.placement_group_id,
            placement_group_policy_respected=self.placement_group_policy_respected,
            private_ip=self.private_ip,
            private_networks=self.private_networks,
            project_id=self.project_id,
            public_ip=self.public_ip,
            public_ips=self.public_ips,
            replace_on_type_change=self.replace_on_type_change,
            root_volumes=self.root_volumes,
            routed_ip_enabled=self.routed_ip_enabled,
            security_group_id=self.security_group_id,
            server_id=self.server_id,
            state=self.state,
            tags=self.tags,
            type=self.type,
            user_data=self.user_data,
            zone=self.zone)


def get_instance_server(name: Optional[str] = None,
                        project_id: Optional[str] = None,
                        server_id: Optional[str] = None,
                        zone: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceServerResult:
    """
    Gets information about an instance server.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_instance_server(server_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The server name. Only one of `name` and `server_id` should be specified.
    :param str project_id: The ID of the project the instance server is associated with.
    :param str server_id: The server id. Only one of `name` and `server_id` should be specified.
    :param str zone: `zone`) The zone in which the server exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getInstanceServer:getInstanceServer', __args__, opts=opts, typ=GetInstanceServerResult).value

    return AwaitableGetInstanceServerResult(
        additional_volume_ids=pulumi.get(__ret__, 'additional_volume_ids'),
        boot_type=pulumi.get(__ret__, 'boot_type'),
        bootscript_id=pulumi.get(__ret__, 'bootscript_id'),
        cloud_init=pulumi.get(__ret__, 'cloud_init'),
        enable_dynamic_ip=pulumi.get(__ret__, 'enable_dynamic_ip'),
        enable_ipv6=pulumi.get(__ret__, 'enable_ipv6'),
        id=pulumi.get(__ret__, 'id'),
        image=pulumi.get(__ret__, 'image'),
        ip_id=pulumi.get(__ret__, 'ip_id'),
        ip_ids=pulumi.get(__ret__, 'ip_ids'),
        ipv6_address=pulumi.get(__ret__, 'ipv6_address'),
        ipv6_gateway=pulumi.get(__ret__, 'ipv6_gateway'),
        ipv6_prefix_length=pulumi.get(__ret__, 'ipv6_prefix_length'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        placement_group_id=pulumi.get(__ret__, 'placement_group_id'),
        placement_group_policy_respected=pulumi.get(__ret__, 'placement_group_policy_respected'),
        private_ip=pulumi.get(__ret__, 'private_ip'),
        private_networks=pulumi.get(__ret__, 'private_networks'),
        project_id=pulumi.get(__ret__, 'project_id'),
        public_ip=pulumi.get(__ret__, 'public_ip'),
        public_ips=pulumi.get(__ret__, 'public_ips'),
        replace_on_type_change=pulumi.get(__ret__, 'replace_on_type_change'),
        root_volumes=pulumi.get(__ret__, 'root_volumes'),
        routed_ip_enabled=pulumi.get(__ret__, 'routed_ip_enabled'),
        security_group_id=pulumi.get(__ret__, 'security_group_id'),
        server_id=pulumi.get(__ret__, 'server_id'),
        state=pulumi.get(__ret__, 'state'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        user_data=pulumi.get(__ret__, 'user_data'),
        zone=pulumi.get(__ret__, 'zone'))
def get_instance_server_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                               project_id: Optional[pulumi.Input[Optional[str]]] = None,
                               server_id: Optional[pulumi.Input[Optional[str]]] = None,
                               zone: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceServerResult]:
    """
    Gets information about an instance server.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_instance_server(server_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The server name. Only one of `name` and `server_id` should be specified.
    :param str project_id: The ID of the project the instance server is associated with.
    :param str server_id: The server id. Only one of `name` and `server_id` should be specified.
    :param str zone: `zone`) The zone in which the server exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['serverId'] = server_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getInstanceServer:getInstanceServer', __args__, opts=opts, typ=GetInstanceServerResult)
    return __ret__.apply(lambda __response__: GetInstanceServerResult(
        additional_volume_ids=pulumi.get(__response__, 'additional_volume_ids'),
        boot_type=pulumi.get(__response__, 'boot_type'),
        bootscript_id=pulumi.get(__response__, 'bootscript_id'),
        cloud_init=pulumi.get(__response__, 'cloud_init'),
        enable_dynamic_ip=pulumi.get(__response__, 'enable_dynamic_ip'),
        enable_ipv6=pulumi.get(__response__, 'enable_ipv6'),
        id=pulumi.get(__response__, 'id'),
        image=pulumi.get(__response__, 'image'),
        ip_id=pulumi.get(__response__, 'ip_id'),
        ip_ids=pulumi.get(__response__, 'ip_ids'),
        ipv6_address=pulumi.get(__response__, 'ipv6_address'),
        ipv6_gateway=pulumi.get(__response__, 'ipv6_gateway'),
        ipv6_prefix_length=pulumi.get(__response__, 'ipv6_prefix_length'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        placement_group_id=pulumi.get(__response__, 'placement_group_id'),
        placement_group_policy_respected=pulumi.get(__response__, 'placement_group_policy_respected'),
        private_ip=pulumi.get(__response__, 'private_ip'),
        private_networks=pulumi.get(__response__, 'private_networks'),
        project_id=pulumi.get(__response__, 'project_id'),
        public_ip=pulumi.get(__response__, 'public_ip'),
        public_ips=pulumi.get(__response__, 'public_ips'),
        replace_on_type_change=pulumi.get(__response__, 'replace_on_type_change'),
        root_volumes=pulumi.get(__response__, 'root_volumes'),
        routed_ip_enabled=pulumi.get(__response__, 'routed_ip_enabled'),
        security_group_id=pulumi.get(__response__, 'security_group_id'),
        server_id=pulumi.get(__response__, 'server_id'),
        state=pulumi.get(__response__, 'state'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        user_data=pulumi.get(__response__, 'user_data'),
        zone=pulumi.get(__response__, 'zone')))
