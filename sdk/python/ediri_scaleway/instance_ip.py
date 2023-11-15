# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceIpArgs', 'InstanceIp']

@pulumi.input_type
class InstanceIpArgs:
    def __init__(__self__, *,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceIp resource.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        :param pulumi.Input[str] type: The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
               
               > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

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
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the IP.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)

        > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the IP should be reserved.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceIpState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceIp resources.
        :param pulumi.Input[str] address: The IP address.
        :param pulumi.Input[str] organization_id: The organization ID the IP is associated with.
        :param pulumi.Input[str] prefix: The IP Prefix.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[str] reverse: The reverse dns attached to this IP
        :param pulumi.Input[str] server_id: The server associated with this IP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        :param pulumi.Input[str] type: The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
               
               > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the IP is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Prefix.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

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
    def reverse(self) -> Optional[pulumi.Input[str]]:
        """
        The reverse dns attached to this IP
        """
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The server associated with this IP
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

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
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)

        > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the IP should be reserved.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceIp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Compute Instance IPs. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#ips-268151).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        server_ip = scaleway.InstanceIp("serverIp")
        ```

        ## Import

        IPs can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/instanceIp:InstanceIp server_ip fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        :param pulumi.Input[str] type: The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
               
               > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InstanceIpArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Compute Instance IPs. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#ips-268151).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        server_ip = scaleway.InstanceIp("serverIp")
        ```

        ## Import

        IPs can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/instanceIp:InstanceIp server_ip fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceIpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceIpArgs.__new__(InstanceIpArgs)

            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["zone"] = zone
            __props__.__dict__["address"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["prefix"] = None
            __props__.__dict__["reverse"] = None
            __props__.__dict__["server_id"] = None
        super(InstanceIp, __self__).__init__(
            'scaleway:index/instanceIp:InstanceIp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            reverse: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceIp':
        """
        Get an existing InstanceIp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The IP address.
        :param pulumi.Input[str] organization_id: The organization ID the IP is associated with.
        :param pulumi.Input[str] prefix: The IP Prefix.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the IP is associated with.
        :param pulumi.Input[str] reverse: The reverse dns attached to this IP
        :param pulumi.Input[str] server_id: The server associated with this IP
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IP.
        :param pulumi.Input[str] type: The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)
               
               > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceIpState.__new__(_InstanceIpState)

        __props__.__dict__["address"] = address
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["reverse"] = reverse
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["zone"] = zone
        return InstanceIp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The IP address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the IP is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[str]:
        """
        The IP Prefix.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the IP is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[str]:
        """
        The reverse dns attached to this IP
        """
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The server associated with this IP
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the IP.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the IP (`nat`, `routed_ipv4`, `routed_ipv6`), more information in [the documentation](https://www.scaleway.com/en/docs/compute/instances/api-cli/using-routed-ips/)

        > **Important:** An IP can migrate from `nat` to `routed_ipv4` but cannot be converted back
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the IP should be reserved.
        """
        return pulumi.get(self, "zone")

