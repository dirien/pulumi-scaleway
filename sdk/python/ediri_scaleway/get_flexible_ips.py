# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetFlexibleIpsResult',
    'AwaitableGetFlexibleIpsResult',
    'get_flexible_ips',
    'get_flexible_ips_output',
]

@pulumi.output_type
class GetFlexibleIpsResult:
    """
    A collection of values returned by getFlexibleIps.
    """
    def __init__(__self__, id=None, ips=None, organization_id=None, project_id=None, server_ids=None, tags=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if server_ids and not isinstance(server_ids, list):
            raise TypeError("Expected argument 'server_ids' to be a list")
        pulumi.set(__self__, "server_ids", server_ids)
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
    @pulumi.getter
    def ips(self) -> Sequence['outputs.GetFlexibleIpsIpResult']:
        """
        List of found flexible IPS
        """
        return pulumi.get(self, "ips")

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
    @pulumi.getter(name="serverIds")
    def server_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "server_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence[str]]:
        """
        The list of tags which are attached to the flexible IP.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        (Defaults to provider `zone`) The zone in which the MAC address exist.
        """
        return pulumi.get(self, "zone")


class AwaitableGetFlexibleIpsResult(GetFlexibleIpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlexibleIpsResult(
            id=self.id,
            ips=self.ips,
            organization_id=self.organization_id,
            project_id=self.project_id,
            server_ids=self.server_ids,
            tags=self.tags,
            zone=self.zone)


def get_flexible_ips(project_id: Optional[str] = None,
                     server_ids: Optional[Sequence[str]] = None,
                     tags: Optional[Sequence[str]] = None,
                     zone: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlexibleIpsResult:
    """
    Gets information about multiple Flexible IPs.

    ## Example Usage

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    fips_by_tags = scaleway.get_flexible_ips(tags=["a tag"])
    my_offer = scaleway.get_baremetal_offer(name="EM-B112X-SSD")
    base = scaleway.BaremetalServer("base",
        offer=my_offer.offer_id,
        install_config_afterward=True)
    first = scaleway.FlexibleIp("first",
        server_id=base.id,
        tags=[
            "foo",
            "first",
        ])
    second = scaleway.FlexibleIp("second",
        server_id=base.id,
        tags=[
            "foo",
            "second",
        ])
    fips_by_server_id = scaleway.get_flexible_ips_output(server_ids=[base.id])
    ```


    :param str project_id: (Defaults to provider `project_id`) The ID of the project the IP is in.
    :param Sequence[str] server_ids: List of server IDs used as filter. IPs with these exact server IDs are listed.
    :param Sequence[str] tags: List of tags used as filter. IPs with these exact tags are listed.
    :param str zone: `zone`) The zone in which IPs exist.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['serverIds'] = server_ids
    __args__['tags'] = tags
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getFlexibleIps:getFlexibleIps', __args__, opts=opts, typ=GetFlexibleIpsResult).value

    return AwaitableGetFlexibleIpsResult(
        id=pulumi.get(__ret__, 'id'),
        ips=pulumi.get(__ret__, 'ips'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        server_ids=pulumi.get(__ret__, 'server_ids'),
        tags=pulumi.get(__ret__, 'tags'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_flexible_ips)
def get_flexible_ips_output(project_id: Optional[pulumi.Input[Optional[str]]] = None,
                            server_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            tags: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            zone: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlexibleIpsResult]:
    """
    Gets information about multiple Flexible IPs.

    ## Example Usage

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    fips_by_tags = scaleway.get_flexible_ips(tags=["a tag"])
    my_offer = scaleway.get_baremetal_offer(name="EM-B112X-SSD")
    base = scaleway.BaremetalServer("base",
        offer=my_offer.offer_id,
        install_config_afterward=True)
    first = scaleway.FlexibleIp("first",
        server_id=base.id,
        tags=[
            "foo",
            "first",
        ])
    second = scaleway.FlexibleIp("second",
        server_id=base.id,
        tags=[
            "foo",
            "second",
        ])
    fips_by_server_id = scaleway.get_flexible_ips_output(server_ids=[base.id])
    ```


    :param str project_id: (Defaults to provider `project_id`) The ID of the project the IP is in.
    :param Sequence[str] server_ids: List of server IDs used as filter. IPs with these exact server IDs are listed.
    :param Sequence[str] tags: List of tags used as filter. IPs with these exact tags are listed.
    :param str zone: `zone`) The zone in which IPs exist.
    """
    ...
