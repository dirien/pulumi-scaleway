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
    'GetVpcPrivateNetworkResult',
    'AwaitableGetVpcPrivateNetworkResult',
    'get_vpc_private_network',
    'get_vpc_private_network_output',
]

@pulumi.output_type
class GetVpcPrivateNetworkResult:
    """
    A collection of values returned by getVpcPrivateNetwork.
    """
    def __init__(__self__, created_at=None, id=None, ipv4_subnets=None, ipv6_subnets=None, is_regional=None, name=None, organization_id=None, private_network_id=None, project_id=None, region=None, tags=None, updated_at=None, vpc_id=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv4_subnets and not isinstance(ipv4_subnets, list):
            raise TypeError("Expected argument 'ipv4_subnets' to be a list")
        pulumi.set(__self__, "ipv4_subnets", ipv4_subnets)
        if ipv6_subnets and not isinstance(ipv6_subnets, list):
            raise TypeError("Expected argument 'ipv6_subnets' to be a list")
        pulumi.set(__self__, "ipv6_subnets", ipv6_subnets)
        if is_regional and not isinstance(is_regional, bool):
            raise TypeError("Expected argument 'is_regional' to be a bool")
        pulumi.set(__self__, "is_regional", is_regional)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if private_network_id and not isinstance(private_network_id, str):
            raise TypeError("Expected argument 'private_network_id' to be a str")
        pulumi.set(__self__, "private_network_id", private_network_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipv4Subnets")
    def ipv4_subnets(self) -> Sequence['outputs.GetVpcPrivateNetworkIpv4SubnetResult']:
        """
        The IPv4 subnet associated with the Private Network.
        """
        return pulumi.get(self, "ipv4_subnets")

    @property
    @pulumi.getter(name="ipv6Subnets")
    def ipv6_subnets(self) -> Sequence['outputs.GetVpcPrivateNetworkIpv6SubnetResult']:
        """
        The IPv6 subnets associated with the Private Network.
        """
        return pulumi.get(self, "ipv6_subnets")

    @property
    @pulumi.getter(name="isRegional")
    def is_regional(self) -> bool:
        return pulumi.get(self, "is_regional")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[str]:
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetVpcPrivateNetworkResult(GetVpcPrivateNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcPrivateNetworkResult(
            created_at=self.created_at,
            id=self.id,
            ipv4_subnets=self.ipv4_subnets,
            ipv6_subnets=self.ipv6_subnets,
            is_regional=self.is_regional,
            name=self.name,
            organization_id=self.organization_id,
            private_network_id=self.private_network_id,
            project_id=self.project_id,
            region=self.region,
            tags=self.tags,
            updated_at=self.updated_at,
            vpc_id=self.vpc_id,
            zone=self.zone)


def get_vpc_private_network(name: Optional[str] = None,
                            private_network_id: Optional[str] = None,
                            project_id: Optional[str] = None,
                            region: Optional[str] = None,
                            vpc_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPrivateNetworkResult:
    """
    Gets information about a Private Network.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_name = scaleway.get_vpc_private_network(name="foobar")
    my_name_and_vpc_id = scaleway.get_vpc_private_network(name="foobar",
        vpc_id="11111111-1111-1111-1111-111111111111")
    my_id = scaleway.get_vpc_private_network(private_network_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: Name of the Private Network. Cannot be used with `private_network_id`.
    :param str private_network_id: ID of the Private Network. Cannot be used with `name` or `vpc_id`.
    :param str project_id: The ID of the Project the Private Network is associated with.
    :param str vpc_id: ID of the VPC the Private Network is in. Cannot be used with `private_network_id`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['privateNetworkId'] = private_network_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork', __args__, opts=opts, typ=GetVpcPrivateNetworkResult).value

    return AwaitableGetVpcPrivateNetworkResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        ipv4_subnets=pulumi.get(__ret__, 'ipv4_subnets'),
        ipv6_subnets=pulumi.get(__ret__, 'ipv6_subnets'),
        is_regional=pulumi.get(__ret__, 'is_regional'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        private_network_id=pulumi.get(__ret__, 'private_network_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        zone=pulumi.get(__ret__, 'zone'))
def get_vpc_private_network_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                   private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   region: Optional[pulumi.Input[Optional[str]]] = None,
                                   vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpcPrivateNetworkResult]:
    """
    Gets information about a Private Network.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_name = scaleway.get_vpc_private_network(name="foobar")
    my_name_and_vpc_id = scaleway.get_vpc_private_network(name="foobar",
        vpc_id="11111111-1111-1111-1111-111111111111")
    my_id = scaleway.get_vpc_private_network(private_network_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: Name of the Private Network. Cannot be used with `private_network_id`.
    :param str private_network_id: ID of the Private Network. Cannot be used with `name` or `vpc_id`.
    :param str project_id: The ID of the Project the Private Network is associated with.
    :param str vpc_id: ID of the VPC the Private Network is in. Cannot be used with `private_network_id`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['privateNetworkId'] = private_network_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork', __args__, opts=opts, typ=GetVpcPrivateNetworkResult)
    return __ret__.apply(lambda __response__: GetVpcPrivateNetworkResult(
        created_at=pulumi.get(__response__, 'created_at'),
        id=pulumi.get(__response__, 'id'),
        ipv4_subnets=pulumi.get(__response__, 'ipv4_subnets'),
        ipv6_subnets=pulumi.get(__response__, 'ipv6_subnets'),
        is_regional=pulumi.get(__response__, 'is_regional'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        private_network_id=pulumi.get(__response__, 'private_network_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        tags=pulumi.get(__response__, 'tags'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        vpc_id=pulumi.get(__response__, 'vpc_id'),
        zone=pulumi.get(__response__, 'zone')))
