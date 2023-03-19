# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceIpReverseDnsArgs', 'InstanceIpReverseDns']

@pulumi.input_type
class InstanceIpReverseDnsArgs:
    def __init__(__self__, *,
                 ip_id: pulumi.Input[str],
                 reverse: pulumi.Input[str],
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceIpReverseDns resource.
        :param pulumi.Input[str] ip_id: The IP ID
        :param pulumi.Input[str] reverse: The reverse DNS for this IP.
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        pulumi.set(__self__, "ip_id", ip_id)
        pulumi.set(__self__, "reverse", reverse)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> pulumi.Input[str]:
        """
        The IP ID
        """
        return pulumi.get(self, "ip_id")

    @ip_id.setter
    def ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_id", value)

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Input[str]:
        """
        The reverse DNS for this IP.
        """
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: pulumi.Input[str]):
        pulumi.set(self, "reverse", value)

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
class _InstanceIpReverseDnsState:
    def __init__(__self__, *,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceIpReverseDns resources.
        :param pulumi.Input[str] ip_id: The IP ID
        :param pulumi.Input[str] reverse: The reverse DNS for this IP.
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        if ip_id is not None:
            pulumi.set(__self__, "ip_id", ip_id)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The IP ID
        """
        return pulumi.get(self, "ip_id")

    @ip_id.setter
    def ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_id", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input[str]]:
        """
        The reverse DNS for this IP.
        """
        return pulumi.get(self, "reverse")

    @reverse.setter
    def reverse(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reverse", value)

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


class InstanceIpReverseDns(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Scaleway Compute Instance IPs Reverse DNS.

        Please check our [guide](https://www.scaleway.com/en/docs/compute/instances/how-to/configure-reverse-dns/) for more details

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        server_ip = scaleway.InstanceIp("serverIp")
        tf_a = scaleway.DomainRecord("tfA",
            dns_zone="scaleway.com",
            type="A",
            data=server_ip.address,
            ttl=3600,
            priority=1)
        reverse = scaleway.InstanceIpReverseDns("reverse",
            ip_id=server_ip.id,
            reverse="www.scaleway.com")
        ```

        ## Import

        IPs reverse DNS can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/instanceIpReverseDns:InstanceIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip_id: The IP ID
        :param pulumi.Input[str] reverse: The reverse DNS for this IP.
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceIpReverseDnsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Scaleway Compute Instance IPs Reverse DNS.

        Please check our [guide](https://www.scaleway.com/en/docs/compute/instances/how-to/configure-reverse-dns/) for more details

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        server_ip = scaleway.InstanceIp("serverIp")
        tf_a = scaleway.DomainRecord("tfA",
            dns_zone="scaleway.com",
            type="A",
            data=server_ip.address,
            ttl=3600,
            priority=1)
        reverse = scaleway.InstanceIpReverseDns("reverse",
            ip_id=server_ip.id,
            reverse="www.scaleway.com")
        ```

        ## Import

        IPs reverse DNS can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/instanceIpReverseDns:InstanceIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceIpReverseDnsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIpReverseDnsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceIpReverseDnsArgs.__new__(InstanceIpReverseDnsArgs)

            if ip_id is None and not opts.urn:
                raise TypeError("Missing required property 'ip_id'")
            __props__.__dict__["ip_id"] = ip_id
            if reverse is None and not opts.urn:
                raise TypeError("Missing required property 'reverse'")
            __props__.__dict__["reverse"] = reverse
            __props__.__dict__["zone"] = zone
        super(InstanceIpReverseDns, __self__).__init__(
            'scaleway:index/instanceIpReverseDns:InstanceIpReverseDns',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ip_id: Optional[pulumi.Input[str]] = None,
            reverse: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceIpReverseDns':
        """
        Get an existing InstanceIpReverseDns resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip_id: The IP ID
        :param pulumi.Input[str] reverse: The reverse DNS for this IP.
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceIpReverseDnsState.__new__(_InstanceIpReverseDnsState)

        __props__.__dict__["ip_id"] = ip_id
        __props__.__dict__["reverse"] = reverse
        __props__.__dict__["zone"] = zone
        return InstanceIpReverseDns(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> pulumi.Output[str]:
        """
        The IP ID
        """
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[str]:
        """
        The reverse DNS for this IP.
        """
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the IP should be reserved.
        """
        return pulumi.get(self, "zone")

