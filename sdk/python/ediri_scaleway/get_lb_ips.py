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
    'GetLbIpsResult',
    'AwaitableGetLbIpsResult',
    'get_lb_ips',
    'get_lb_ips_output',
]

@pulumi.output_type
class GetLbIpsResult:
    """
    A collection of values returned by getLbIps.
    """
    def __init__(__self__, id=None, ip_cidr_range=None, ip_type=None, ips=None, organization_id=None, project_id=None, tags=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_cidr_range and not isinstance(ip_cidr_range, str):
            raise TypeError("Expected argument 'ip_cidr_range' to be a str")
        pulumi.set(__self__, "ip_cidr_range", ip_cidr_range)
        if ip_type and not isinstance(ip_type, str):
            raise TypeError("Expected argument 'ip_type' to be a str")
        pulumi.set(__self__, "ip_type", ip_type)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipCidrRange")
    def ip_cidr_range(self) -> Optional[str]:
        return pulumi.get(self, "ip_cidr_range")

    @property
    @pulumi.getter(name="ipType")
    def ip_type(self) -> Optional[str]:
        return pulumi.get(self, "ip_type")

    @property
    @pulumi.getter
    def ips(self) -> Sequence['outputs.GetLbIpsIpResult']:
        """
        List of retrieved IPs
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The ID of the Organization the Load Balancer is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of the Project the Load Balancer is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The zone of the Load Balancer.
        """
        return pulumi.get(self, "zone")


class AwaitableGetLbIpsResult(GetLbIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbIpsResult(
            id=self.id,
            ip_cidr_range=self.ip_cidr_range,
            ip_type=self.ip_type,
            ips=self.ips,
            organization_id=self.organization_id,
            project_id=self.project_id,
            tags=self.tags,
            zone=self.zone)


def get_lb_ips(ip_cidr_range: Optional[str] = None,
               ip_type: Optional[str] = None,
               project_id: Optional[str] = None,
               tags: Optional[Sequence[str]] = None,
               zone: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbIpsResult:
    """
    Gets information about multiple Load Balancer IP addresses.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_lb_ips(ip_cidr_range="0.0.0.0/0",
        zone="fr-par-2")
    ips_by_tags_and_type = scaleway.get_lb_ips(ip_type="ipv4",
        tags=["a tag"])
    ```


    :param str ip_cidr_range: The IP CIDR range to filter for. IPs within a matching CIDR block are listed.
    :param str ip_type: The IP type used as a filter.
    :param str project_id: The ID of the Project the Load Balancer is associated with.
    :param Sequence[str] tags: List of tags used as filter. IPs with these exact tags are listed.
    :param str zone: `zone`) The zone in which the IPs exist.
    """
    __args__ = dict()
    __args__['ipCidrRange'] = ip_cidr_range
    __args__['ipType'] = ip_type
    __args__['projectId'] = project_id
    __args__['tags'] = tags
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getLbIps:getLbIps', __args__, opts=opts, typ=GetLbIpsResult).value

    return AwaitableGetLbIpsResult(
        id=pulumi.get(__ret__, 'id'),
        ip_cidr_range=pulumi.get(__ret__, 'ip_cidr_range'),
        ip_type=pulumi.get(__ret__, 'ip_type'),
        ips=pulumi.get(__ret__, 'ips'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        tags=pulumi.get(__ret__, 'tags'),
        zone=pulumi.get(__ret__, 'zone'))
def get_lb_ips_output(ip_cidr_range: Optional[pulumi.Input[Optional[str]]] = None,
                      ip_type: Optional[pulumi.Input[Optional[str]]] = None,
                      project_id: Optional[pulumi.Input[Optional[str]]] = None,
                      tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      zone: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetLbIpsResult]:
    """
    Gets information about multiple Load Balancer IP addresses.

    For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-flex-ips/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-ip-addresses-list-ip-addresses).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_lb_ips(ip_cidr_range="0.0.0.0/0",
        zone="fr-par-2")
    ips_by_tags_and_type = scaleway.get_lb_ips(ip_type="ipv4",
        tags=["a tag"])
    ```


    :param str ip_cidr_range: The IP CIDR range to filter for. IPs within a matching CIDR block are listed.
    :param str ip_type: The IP type used as a filter.
    :param str project_id: The ID of the Project the Load Balancer is associated with.
    :param Sequence[str] tags: List of tags used as filter. IPs with these exact tags are listed.
    :param str zone: `zone`) The zone in which the IPs exist.
    """
    __args__ = dict()
    __args__['ipCidrRange'] = ip_cidr_range
    __args__['ipType'] = ip_type
    __args__['projectId'] = project_id
    __args__['tags'] = tags
    __args__['zone'] = zone
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getLbIps:getLbIps', __args__, opts=opts, typ=GetLbIpsResult)
    return __ret__.apply(lambda __response__: GetLbIpsResult(
        id=pulumi.get(__response__, 'id'),
        ip_cidr_range=pulumi.get(__response__, 'ip_cidr_range'),
        ip_type=pulumi.get(__response__, 'ip_type'),
        ips=pulumi.get(__response__, 'ips'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        tags=pulumi.get(__response__, 'tags'),
        zone=pulumi.get(__response__, 'zone')))
