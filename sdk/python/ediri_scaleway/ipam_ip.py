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
from ._inputs import *

__all__ = ['IpamIpArgs', 'IpamIp']

@pulumi.input_type
class IpamIpArgs:
    def __init__(__self__, *,
                 sources: pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]],
                 address: Optional[pulumi.Input[str]] = None,
                 is_ipv6: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a IpamIp resource.
        :param pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]] sources: The source in which to book the IP.
        :param pulumi.Input[str] address: Request a specific IP in the requested source pool.
        :param pulumi.Input[bool] is_ipv6: Defines whether to request an IPv6 instead of an IPv4.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[str] region: `region`) The region of the IP.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        """
        pulumi.set(__self__, "sources", sources)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if is_ipv6 is not None:
            pulumi.set(__self__, "is_ipv6", is_ipv6)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]]:
        """
        The source in which to book the IP.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Request a specific IP in the requested source pool.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether to request an IPv6 instead of an IPv4.
        """
        return pulumi.get(self, "is_ipv6")

    @is_ipv6.setter
    def is_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ipv6", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the IP.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the IP.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _IpamIpState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 is_ipv6: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpResourceArgs']]]] = None,
                 reverses: Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpReverseArgs']]]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpamIp resources.
        :param pulumi.Input[str] address: Request a specific IP in the requested source pool.
        :param pulumi.Input[str] created_at: Date and time of IP's creation (RFC 3339 format).
        :param pulumi.Input[bool] is_ipv6: Defines whether to request an IPv6 instead of an IPv4.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[str] region: `region`) The region of the IP.
        :param pulumi.Input[Sequence[pulumi.Input['IpamIpResourceArgs']]] resources: The IP resource.
        :param pulumi.Input[Sequence[pulumi.Input['IpamIpReverseArgs']]] reverses: The reverses DNS for this IP.
        :param pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]] sources: The source in which to book the IP.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        :param pulumi.Input[str] updated_at: Date and time of IP's last update (RFC 3339 format).
        :param pulumi.Input[str] zone: The zone of the IP.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if is_ipv6 is not None:
            pulumi.set(__self__, "is_ipv6", is_ipv6)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if reverses is not None:
            pulumi.set(__self__, "reverses", reverses)
        if sources is not None:
            pulumi.set(__self__, "sources", sources)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Request a specific IP in the requested source pool.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of IP's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether to request an IPv6 instead of an IPv4.
        """
        return pulumi.get(self, "is_ipv6")

    @is_ipv6.setter
    def is_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_ipv6", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region of the IP.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpResourceArgs']]]]:
        """
        The IP resource.
        """
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpResourceArgs']]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def reverses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpReverseArgs']]]]:
        """
        The reverses DNS for this IP.
        """
        return pulumi.get(self, "reverses")

    @reverses.setter
    def reverses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpReverseArgs']]]]):
        pulumi.set(self, "reverses", value)

    @property
    @pulumi.getter
    def sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]]]:
        """
        The source in which to book the IP.
        """
        return pulumi.get(self, "sources")

    @sources.setter
    def sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpamIpSourceArgs']]]]):
        pulumi.set(self, "sources", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the IP.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of IP's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the IP.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class IpamIp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 is_ipv6: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpSourceArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Books and manages Scaleway IPAM IPs.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        vpc01 = scaleway.Vpc("vpc01")
        pn01 = scaleway.VpcPrivateNetwork("pn01",
            vpc_id=vpc01.id,
            ipv4_subnet=scaleway.VpcPrivateNetworkIpv4SubnetArgs(
                subnet="172.16.32.0/22",
            ))
        ip01 = scaleway.IpamIp("ip01", sources=[scaleway.IpamIpSourceArgs(
            private_network_id=pn01.id,
        )])
        ```

        ### Request a specific IPv4

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        vpc01 = scaleway.Vpc("vpc01")
        pn01 = scaleway.VpcPrivateNetwork("pn01",
            vpc_id=vpc01.id,
            ipv4_subnet=scaleway.VpcPrivateNetworkIpv4SubnetArgs(
                subnet="172.16.32.0/22",
            ))
        ip01 = scaleway.IpamIp("ip01",
            address="172.16.32.7/22",
            sources=[scaleway.IpamIpSourceArgs(
                private_network_id=pn01.id,
            )])
        ```

        ### Request an IPv6

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        vpc01 = scaleway.Vpc("vpc01")
        pn01 = scaleway.VpcPrivateNetwork("pn01",
            vpc_id=vpc01.id,
            ipv6_subnets=[scaleway.VpcPrivateNetworkIpv6SubnetArgs(
                subnet="fd46:78ab:30b8:177c::/64",
            )])
        ip01 = scaleway.IpamIp("ip01",
            is_ipv6=True,
            sources=[scaleway.IpamIpSourceArgs(
                private_network_id=pn01.id,
            )])
        ```

        ## Import

        IPAM IPs can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/ipamIp:IpamIp ip_demo fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Request a specific IP in the requested source pool.
        :param pulumi.Input[bool] is_ipv6: Defines whether to request an IPv6 instead of an IPv4.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[str] region: `region`) The region of the IP.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpSourceArgs']]]] sources: The source in which to book the IP.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpamIpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Books and manages Scaleway IPAM IPs.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        vpc01 = scaleway.Vpc("vpc01")
        pn01 = scaleway.VpcPrivateNetwork("pn01",
            vpc_id=vpc01.id,
            ipv4_subnet=scaleway.VpcPrivateNetworkIpv4SubnetArgs(
                subnet="172.16.32.0/22",
            ))
        ip01 = scaleway.IpamIp("ip01", sources=[scaleway.IpamIpSourceArgs(
            private_network_id=pn01.id,
        )])
        ```

        ### Request a specific IPv4

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        vpc01 = scaleway.Vpc("vpc01")
        pn01 = scaleway.VpcPrivateNetwork("pn01",
            vpc_id=vpc01.id,
            ipv4_subnet=scaleway.VpcPrivateNetworkIpv4SubnetArgs(
                subnet="172.16.32.0/22",
            ))
        ip01 = scaleway.IpamIp("ip01",
            address="172.16.32.7/22",
            sources=[scaleway.IpamIpSourceArgs(
                private_network_id=pn01.id,
            )])
        ```

        ### Request an IPv6

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        vpc01 = scaleway.Vpc("vpc01")
        pn01 = scaleway.VpcPrivateNetwork("pn01",
            vpc_id=vpc01.id,
            ipv6_subnets=[scaleway.VpcPrivateNetworkIpv6SubnetArgs(
                subnet="fd46:78ab:30b8:177c::/64",
            )])
        ip01 = scaleway.IpamIp("ip01",
            is_ipv6=True,
            sources=[scaleway.IpamIpSourceArgs(
                private_network_id=pn01.id,
            )])
        ```

        ## Import

        IPAM IPs can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/ipamIp:IpamIp ip_demo fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IpamIpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpamIpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 is_ipv6: Optional[pulumi.Input[bool]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpSourceArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpamIpArgs.__new__(IpamIpArgs)

            __props__.__dict__["address"] = address
            __props__.__dict__["is_ipv6"] = is_ipv6
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            if sources is None and not opts.urn:
                raise TypeError("Missing required property 'sources'")
            __props__.__dict__["sources"] = sources
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["resources"] = None
            __props__.__dict__["reverses"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["zone"] = None
        super(IpamIp, __self__).__init__(
            'scaleway:index/ipamIp:IpamIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            is_ipv6: Optional[pulumi.Input[bool]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpResourceArgs']]]]] = None,
            reverses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpReverseArgs']]]]] = None,
            sources: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpSourceArgs']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'IpamIp':
        """
        Get an existing IpamIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Request a specific IP in the requested source pool.
        :param pulumi.Input[str] created_at: Date and time of IP's creation (RFC 3339 format).
        :param pulumi.Input[bool] is_ipv6: Defines whether to request an IPv6 instead of an IPv4.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[str] region: `region`) The region of the IP.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpResourceArgs']]]] resources: The IP resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpReverseArgs']]]] reverses: The reverses DNS for this IP.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['IpamIpSourceArgs']]]] sources: The source in which to book the IP.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        :param pulumi.Input[str] updated_at: Date and time of IP's last update (RFC 3339 format).
        :param pulumi.Input[str] zone: The zone of the IP.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpamIpState.__new__(_IpamIpState)

        __props__.__dict__["address"] = address
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["is_ipv6"] = is_ipv6
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["resources"] = resources
        __props__.__dict__["reverses"] = reverses
        __props__.__dict__["sources"] = sources
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["zone"] = zone
        return IpamIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Request a specific IP in the requested source pool.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date and time of IP's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="isIpv6")
    def is_ipv6(self) -> pulumi.Output[Optional[bool]]:
        """
        Defines whether to request an IPv6 instead of an IPv4.
        """
        return pulumi.get(self, "is_ipv6")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region of the IP.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Sequence['outputs.IpamIpResource']]:
        """
        The IP resource.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def reverses(self) -> pulumi.Output[Sequence['outputs.IpamIpReverse']]:
        """
        The reverses DNS for this IP.
        """
        return pulumi.get(self, "reverses")

    @property
    @pulumi.getter
    def sources(self) -> pulumi.Output[Sequence['outputs.IpamIpSource']]:
        """
        The source in which to book the IP.
        """
        return pulumi.get(self, "sources")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the IP.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date and time of IP's last update (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone of the IP.
        """
        return pulumi.get(self, "zone")

