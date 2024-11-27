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

__all__ = ['VpcPublicGatewayIpReverseDnsArgs', 'VpcPublicGatewayIpReverseDns']

@pulumi.input_type
class VpcPublicGatewayIpReverseDnsArgs:
    def __init__(__self__, *,
                 gateway_ip_id: pulumi.Input[str],
                 reverse: pulumi.Input[str],
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcPublicGatewayIpReverseDns resource.
        :param pulumi.Input[str] gateway_ip_id: The Public Gateway IP ID
        :param pulumi.Input[str] reverse: The reverse domain name for this IP address
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        pulumi.set(__self__, "gateway_ip_id", gateway_ip_id)
        pulumi.set(__self__, "reverse", reverse)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="gatewayIpId")
    def gateway_ip_id(self) -> pulumi.Input[str]:
        """
        The Public Gateway IP ID
        """
        return pulumi.get(self, "gateway_ip_id")

    @gateway_ip_id.setter
    def gateway_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_ip_id", value)

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Input[str]:
        """
        The reverse domain name for this IP address
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
class _VpcPublicGatewayIpReverseDnsState:
    def __init__(__self__, *,
                 gateway_ip_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcPublicGatewayIpReverseDns resources.
        :param pulumi.Input[str] gateway_ip_id: The Public Gateway IP ID
        :param pulumi.Input[str] reverse: The reverse domain name for this IP address
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        if gateway_ip_id is not None:
            pulumi.set(__self__, "gateway_ip_id", gateway_ip_id)
        if reverse is not None:
            pulumi.set(__self__, "reverse", reverse)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="gatewayIpId")
    def gateway_ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Public Gateway IP ID
        """
        return pulumi.get(self, "gateway_ip_id")

    @gateway_ip_id.setter
    def gateway_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_ip_id", value)

    @property
    @pulumi.getter
    def reverse(self) -> Optional[pulumi.Input[str]]:
        """
        The reverse domain name for this IP address
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


class VpcPublicGatewayIpReverseDns(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_ip_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Scaleway Public Gateway public (flexible) IPs' reverse DNS.
        For more information, see [the API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_vpc_public_gateway_ip = scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp")
        tf_a = scaleway.DomainRecord("tfA",
            dns_zone="example.com",
            type="A",
            data=main_vpc_public_gateway_ip.address,
            ttl=3600,
            priority=1)
        main_vpc_public_gateway_ip_reverse_dns = scaleway.VpcPublicGatewayIpReverseDns("mainVpcPublicGatewayIpReverseDns",
            gateway_ip_id=main_vpc_public_gateway_ip.id,
            reverse="tf.example.com")
        ```

        ## Import

        Public Gateway IP reverse DNS can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_ip_id: The Public Gateway IP ID
        :param pulumi.Input[str] reverse: The reverse domain name for this IP address
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcPublicGatewayIpReverseDnsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Scaleway Public Gateway public (flexible) IPs' reverse DNS.
        For more information, see [the API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#path-ips-list-ips).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_vpc_public_gateway_ip = scaleway.VpcPublicGatewayIp("mainVpcPublicGatewayIp")
        tf_a = scaleway.DomainRecord("tfA",
            dns_zone="example.com",
            type="A",
            data=main_vpc_public_gateway_ip.address,
            ttl=3600,
            priority=1)
        main_vpc_public_gateway_ip_reverse_dns = scaleway.VpcPublicGatewayIpReverseDns("mainVpcPublicGatewayIpReverseDns",
            gateway_ip_id=main_vpc_public_gateway_ip.id,
            reverse="tf.example.com")
        ```

        ## Import

        Public Gateway IP reverse DNS can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns reverse fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param VpcPublicGatewayIpReverseDnsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcPublicGatewayIpReverseDnsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_ip_id: Optional[pulumi.Input[str]] = None,
                 reverse: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcPublicGatewayIpReverseDnsArgs.__new__(VpcPublicGatewayIpReverseDnsArgs)

            if gateway_ip_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_ip_id'")
            __props__.__dict__["gateway_ip_id"] = gateway_ip_id
            if reverse is None and not opts.urn:
                raise TypeError("Missing required property 'reverse'")
            __props__.__dict__["reverse"] = reverse
            __props__.__dict__["zone"] = zone
        super(VpcPublicGatewayIpReverseDns, __self__).__init__(
            'scaleway:index/vpcPublicGatewayIpReverseDns:VpcPublicGatewayIpReverseDns',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            gateway_ip_id: Optional[pulumi.Input[str]] = None,
            reverse: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'VpcPublicGatewayIpReverseDns':
        """
        Get an existing VpcPublicGatewayIpReverseDns resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_ip_id: The Public Gateway IP ID
        :param pulumi.Input[str] reverse: The reverse domain name for this IP address
        :param pulumi.Input[str] zone: `zone`) The zone in which the IP should be reserved.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcPublicGatewayIpReverseDnsState.__new__(_VpcPublicGatewayIpReverseDnsState)

        __props__.__dict__["gateway_ip_id"] = gateway_ip_id
        __props__.__dict__["reverse"] = reverse
        __props__.__dict__["zone"] = zone
        return VpcPublicGatewayIpReverseDns(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="gatewayIpId")
    def gateway_ip_id(self) -> pulumi.Output[str]:
        """
        The Public Gateway IP ID
        """
        return pulumi.get(self, "gateway_ip_id")

    @property
    @pulumi.getter
    def reverse(self) -> pulumi.Output[str]:
        """
        The reverse domain name for this IP address
        """
        return pulumi.get(self, "reverse")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the IP should be reserved.
        """
        return pulumi.get(self, "zone")

