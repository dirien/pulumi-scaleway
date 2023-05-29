# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

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
    def __init__(__self__, created_at=None, id=None, name=None, organization_id=None, private_network_id=None, project_id=None, subnets=None, tags=None, updated_at=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
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
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def subnets(self) -> Sequence[str]:
        """
        The subnets CIDR associated with private network.
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

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
            name=self.name,
            organization_id=self.organization_id,
            private_network_id=self.private_network_id,
            project_id=self.project_id,
            subnets=self.subnets,
            tags=self.tags,
            updated_at=self.updated_at,
            zone=self.zone)


def get_vpc_private_network(name: Optional[str] = None,
                            private_network_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcPrivateNetworkResult:
    """
    Gets information about a private network.

    ## Example Usage

    N/A, the usage will be meaningful in the next releases of VPC.


    :param str name: Name of the private network. One of `name` and `private_network_id` should be specified.
    :param str private_network_id: ID of the private network. One of `name` and `private_network_id` should be specified.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['privateNetworkId'] = private_network_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcPrivateNetwork:getVpcPrivateNetwork', __args__, opts=opts, typ=GetVpcPrivateNetworkResult).value

    return AwaitableGetVpcPrivateNetworkResult(
        created_at=__ret__.created_at,
        id=__ret__.id,
        name=__ret__.name,
        organization_id=__ret__.organization_id,
        private_network_id=__ret__.private_network_id,
        project_id=__ret__.project_id,
        subnets=__ret__.subnets,
        tags=__ret__.tags,
        updated_at=__ret__.updated_at,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_vpc_private_network)
def get_vpc_private_network_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                   private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcPrivateNetworkResult]:
    """
    Gets information about a private network.

    ## Example Usage

    N/A, the usage will be meaningful in the next releases of VPC.


    :param str name: Name of the private network. One of `name` and `private_network_id` should be specified.
    :param str private_network_id: ID of the private network. One of `name` and `private_network_id` should be specified.
    """
    ...
