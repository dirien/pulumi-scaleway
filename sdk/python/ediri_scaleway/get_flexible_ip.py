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

__all__ = [
    'GetFlexibleIpResult',
    'AwaitableGetFlexibleIpResult',
    'get_flexible_ip',
    'get_flexible_ip_output',
]

@pulumi.output_type
class GetFlexibleIpResult:
    """
    A collection of values returned by getFlexibleIp.
    """
    def __init__(__self__, created_at=None, description=None, flexible_ip_id=None, id=None, ip_address=None, is_ipv6=None, organization_id=None, project_id=None, reverse=None, server_id=None, status=None, tags=None, updated_at=None, zone=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if flexible_ip_id and not isinstance(flexible_ip_id, str):
            raise TypeError("Expected argument 'flexible_ip_id' to be a str")
        pulumi.set(__self__, "flexible_ip_id", flexible_ip_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if is_ipv6 and not isinstance(is_ipv6, bool):
            raise TypeError("Expected argument 'is_ipv6' to be a bool")
        pulumi.set(__self__, "is_ipv6", is_ipv6)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if reverse and not isinstance(reverse, str):
            raise TypeError("Expected argument 'reverse' to be a str")
        pulumi.set(__self__, "reverse", reverse)
        if server_id and not isinstance(server_id, str):
            raise TypeError("Expected argument 'server_id' to be a str")
        pulumi.set(__self__, "server_id", server_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
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
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="flexibleIpId")
    def flexible_ip_id(self) -> Optional[str]:
        return pulumi.get(self, "flexible_ip_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[str]:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> bool:
        return pulumi.get(self, "is_ipv6")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        (Defaults to provider `organization_id`) The ID of the organization the IP is in.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        (Defaults to provider `project_id`) The ID of the project the IP is in.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def reverse(self) -> str:
        """
        The reverse domain associated with this IP.
        """
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> str:
        """
        The associated server ID if any
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

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


class AwaitableGetFlexibleIpResult(GetFlexibleIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlexibleIpResult(
            created_at=self.created_at,
            description=self.description,
            flexible_ip_id=self.flexible_ip_id,
            id=self.id,
            ip_address=self.ip_address,
            is_ipv6=self.is_ipv6,
            organization_id=self.organization_id,
            project_id=self.project_id,
            reverse=self.reverse,
            server_id=self.server_id,
            status=self.status,
            tags=self.tags,
            updated_at=self.updated_at,
            zone=self.zone)


def get_flexible_ip(flexible_ip_id: Optional[str] = None,
                    ip_address: Optional[str] = None,
                    project_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlexibleIpResult:
    """
    Gets information about a Flexible IP.


    :param str ip_address: The IP address.
           Only one of `ip_address` and `ip_id` should be specified.
    :param str project_id: (Defaults to provider `project_id`) The ID of the project the IP is in.
    """
    __args__ = dict()
    __args__['flexibleIpId'] = flexible_ip_id
    __args__['ipAddress'] = ip_address
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getFlexibleIp:getFlexibleIp', __args__, opts=opts, typ=GetFlexibleIpResult).value

    return AwaitableGetFlexibleIpResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        description=pulumi.get(__ret__, 'description'),
        flexible_ip_id=pulumi.get(__ret__, 'flexible_ip_id'),
        id=pulumi.get(__ret__, 'id'),
        ip_address=pulumi.get(__ret__, 'ip_address'),
        is_ipv6=pulumi.get(__ret__, 'is_ipv6'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        reverse=pulumi.get(__ret__, 'reverse'),
        server_id=pulumi.get(__ret__, 'server_id'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        zone=pulumi.get(__ret__, 'zone'))
def get_flexible_ip_output(flexible_ip_id: Optional[pulumi.Input[Optional[str]]] = None,
                           ip_address: Optional[pulumi.Input[Optional[str]]] = None,
                           project_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFlexibleIpResult]:
    """
    Gets information about a Flexible IP.


    :param str ip_address: The IP address.
           Only one of `ip_address` and `ip_id` should be specified.
    :param str project_id: (Defaults to provider `project_id`) The ID of the project the IP is in.
    """
    __args__ = dict()
    __args__['flexibleIpId'] = flexible_ip_id
    __args__['ipAddress'] = ip_address
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getFlexibleIp:getFlexibleIp', __args__, opts=opts, typ=GetFlexibleIpResult)
    return __ret__.apply(lambda __response__: GetFlexibleIpResult(
        created_at=pulumi.get(__response__, 'created_at'),
        description=pulumi.get(__response__, 'description'),
        flexible_ip_id=pulumi.get(__response__, 'flexible_ip_id'),
        id=pulumi.get(__response__, 'id'),
        ip_address=pulumi.get(__response__, 'ip_address'),
        is_ipv6=pulumi.get(__response__, 'is_ipv6'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        reverse=pulumi.get(__response__, 'reverse'),
        server_id=pulumi.get(__response__, 'server_id'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        zone=pulumi.get(__response__, 'zone')))
