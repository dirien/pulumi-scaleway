# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcPublicGatewayPatRuleArgs', 'VpcPublicGatewayPatRule']

@pulumi.input_type
class VpcPublicGatewayPatRuleArgs:
    def __init__(__self__, *,
                 gateway_id: pulumi.Input[str],
                 private_ip: pulumi.Input[str],
                 private_port: pulumi.Input[int],
                 public_port: pulumi.Input[int],
                 protocol: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcPublicGatewayPatRule resource.
        :param pulumi.Input[str] gateway_id: The ID of the Public Gateway.
        :param pulumi.Input[str] private_ip: The private IP address to forward data to.
        :param pulumi.Input[int] private_port: The private port to translate to.
        :param pulumi.Input[int] public_port: The public port to listen on.
        :param pulumi.Input[str] protocol: The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        :param pulumi.Input[str] zone: `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        pulumi.set(__self__, "gateway_id", gateway_id)
        pulumi.set(__self__, "private_ip", private_ip)
        pulumi.set(__self__, "private_port", private_port)
        pulumi.set(__self__, "public_port", public_port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the Public Gateway.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Input[str]:
        """
        The private IP address to forward data to.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_ip", value)

    @property
    @pulumi.getter(name="privatePort")
    def private_port(self) -> pulumi.Input[int]:
        """
        The private port to translate to.
        """
        return pulumi.get(self, "private_port")

    @private_port.setter
    def private_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "private_port", value)

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> pulumi.Input[int]:
        """
        The public port to listen on.
        """
        return pulumi.get(self, "public_port")

    @public_port.setter
    def public_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "public_port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _VpcPublicGatewayPatRuleState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 private_port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 public_port: Optional[pulumi.Input[int]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcPublicGatewayPatRule resources.
        :param pulumi.Input[str] created_at: The date and time of the creation of the PAT rule configuration.
        :param pulumi.Input[str] gateway_id: The ID of the Public Gateway.
        :param pulumi.Input[str] organization_id: The Organization ID the PAT rule configuration is associated with.
        :param pulumi.Input[str] private_ip: The private IP address to forward data to.
        :param pulumi.Input[int] private_port: The private port to translate to.
        :param pulumi.Input[str] protocol: The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        :param pulumi.Input[int] public_port: The public port to listen on.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the PAT rule configuration.
        :param pulumi.Input[str] zone: `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if gateway_id is not None:
            pulumi.set(__self__, "gateway_id", gateway_id)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if private_ip is not None:
            pulumi.set(__self__, "private_ip", private_ip)
        if private_port is not None:
            pulumi.set(__self__, "private_port", private_port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if public_port is not None:
            pulumi.set(__self__, "public_port", public_port)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the PAT rule configuration.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Public Gateway.
        """
        return pulumi.get(self, "gateway_id")

    @gateway_id.setter
    def gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway_id", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Organization ID the PAT rule configuration is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The private IP address to forward data to.
        """
        return pulumi.get(self, "private_ip")

    @private_ip.setter
    def private_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip", value)

    @property
    @pulumi.getter(name="privatePort")
    def private_port(self) -> Optional[pulumi.Input[int]]:
        """
        The private port to translate to.
        """
        return pulumi.get(self, "private_port")

    @private_port.setter
    def private_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "private_port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> Optional[pulumi.Input[int]]:
        """
        The public port to listen on.
        """
        return pulumi.get(self, "public_port")

    @public_port.setter
    def public_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "public_port", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the PAT rule configuration.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class VpcPublicGatewayPatRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 private_port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 public_port: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Public Gateway PAT (Port Address Translation).
        For more information, see [the API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#pat-rules-e75d10).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        sg01 = scaleway.InstanceSecurityGroup("sg01",
            inbound_default_policy="drop",
            outbound_default_policy="accept",
            inbound_rules=[scaleway.InstanceSecurityGroupInboundRuleArgs(
                action="accept",
                port=22,
                protocol="TCP",
            )])
        srv01 = scaleway.InstanceServer("srv01",
            type="PLAY2-NANO",
            image="ubuntu_jammy",
            security_group_id=sg01.id)
        pn01 = scaleway.VpcPrivateNetwork("pn01")
        pnic01 = scaleway.InstancePrivateNic("pnic01",
            server_id=srv01.id,
            private_network_id=pn01.id)
        dhcp01 = scaleway.VpcPublicGatewayDhcp("dhcp01", subnet="192.168.0.0/24")
        ip01 = scaleway.VpcPublicGatewayIp("ip01")
        pg01 = scaleway.VpcPublicGateway("pg01",
            type="VPC-GW-S",
            ip_id=ip01.id)
        gn01 = scaleway.VpcGatewayNetwork("gn01",
            gateway_id=pg01.id,
            private_network_id=pn01.id,
            dhcp_id=dhcp01.id,
            cleanup_dhcp=True,
            enable_masquerade=True)
        rsv01 = scaleway.VpcPublicGatewayDhcpReservation("rsv01",
            gateway_network_id=gn01.id,
            mac_address=pnic01.mac_address,
            ip_address="192.168.0.7")
        # PAT rule for SSH traffic
        pat01 = scaleway.VpcPublicGatewayPatRule("pat01",
            gateway_id=pg01.id,
            private_ip=rsv01.ip_address,
            private_port=22,
            public_port=2202,
            protocol="tcp")
        ```

        ## Import

        Public Gateway PAT rule configurations can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] gateway_id: The ID of the Public Gateway.
        :param pulumi.Input[str] private_ip: The private IP address to forward data to.
        :param pulumi.Input[int] private_port: The private port to translate to.
        :param pulumi.Input[str] protocol: The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        :param pulumi.Input[int] public_port: The public port to listen on.
        :param pulumi.Input[str] zone: `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcPublicGatewayPatRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Public Gateway PAT (Port Address Translation).
        For more information, see [the API documentation](https://www.scaleway.com/en/developers/api/public-gateway/#pat-rules-e75d10).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        sg01 = scaleway.InstanceSecurityGroup("sg01",
            inbound_default_policy="drop",
            outbound_default_policy="accept",
            inbound_rules=[scaleway.InstanceSecurityGroupInboundRuleArgs(
                action="accept",
                port=22,
                protocol="TCP",
            )])
        srv01 = scaleway.InstanceServer("srv01",
            type="PLAY2-NANO",
            image="ubuntu_jammy",
            security_group_id=sg01.id)
        pn01 = scaleway.VpcPrivateNetwork("pn01")
        pnic01 = scaleway.InstancePrivateNic("pnic01",
            server_id=srv01.id,
            private_network_id=pn01.id)
        dhcp01 = scaleway.VpcPublicGatewayDhcp("dhcp01", subnet="192.168.0.0/24")
        ip01 = scaleway.VpcPublicGatewayIp("ip01")
        pg01 = scaleway.VpcPublicGateway("pg01",
            type="VPC-GW-S",
            ip_id=ip01.id)
        gn01 = scaleway.VpcGatewayNetwork("gn01",
            gateway_id=pg01.id,
            private_network_id=pn01.id,
            dhcp_id=dhcp01.id,
            cleanup_dhcp=True,
            enable_masquerade=True)
        rsv01 = scaleway.VpcPublicGatewayDhcpReservation("rsv01",
            gateway_network_id=gn01.id,
            mac_address=pnic01.mac_address,
            ip_address="192.168.0.7")
        # PAT rule for SSH traffic
        pat01 = scaleway.VpcPublicGatewayPatRule("pat01",
            gateway_id=pg01.id,
            private_ip=rsv01.ip_address,
            private_port=22,
            public_port=2202,
            protocol="tcp")
        ```

        ## Import

        Public Gateway PAT rule configurations can be imported using `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param VpcPublicGatewayPatRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcPublicGatewayPatRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway_id: Optional[pulumi.Input[str]] = None,
                 private_ip: Optional[pulumi.Input[str]] = None,
                 private_port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 public_port: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcPublicGatewayPatRuleArgs.__new__(VpcPublicGatewayPatRuleArgs)

            if gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'gateway_id'")
            __props__.__dict__["gateway_id"] = gateway_id
            if private_ip is None and not opts.urn:
                raise TypeError("Missing required property 'private_ip'")
            __props__.__dict__["private_ip"] = private_ip
            if private_port is None and not opts.urn:
                raise TypeError("Missing required property 'private_port'")
            __props__.__dict__["private_port"] = private_port
            __props__.__dict__["protocol"] = protocol
            if public_port is None and not opts.urn:
                raise TypeError("Missing required property 'public_port'")
            __props__.__dict__["public_port"] = public_port
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["updated_at"] = None
        super(VpcPublicGatewayPatRule, __self__).__init__(
            'scaleway:index/vpcPublicGatewayPatRule:VpcPublicGatewayPatRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            gateway_id: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            private_port: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            public_port: Optional[pulumi.Input[int]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'VpcPublicGatewayPatRule':
        """
        Get an existing VpcPublicGatewayPatRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time of the creation of the PAT rule configuration.
        :param pulumi.Input[str] gateway_id: The ID of the Public Gateway.
        :param pulumi.Input[str] organization_id: The Organization ID the PAT rule configuration is associated with.
        :param pulumi.Input[str] private_ip: The private IP address to forward data to.
        :param pulumi.Input[int] private_port: The private port to translate to.
        :param pulumi.Input[str] protocol: The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        :param pulumi.Input[int] public_port: The public port to listen on.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the PAT rule configuration.
        :param pulumi.Input[str] zone: `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcPublicGatewayPatRuleState.__new__(_VpcPublicGatewayPatRuleState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["gateway_id"] = gateway_id
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["private_ip"] = private_ip
        __props__.__dict__["private_port"] = private_port
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["public_port"] = public_port
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["zone"] = zone
        return VpcPublicGatewayPatRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the PAT rule configuration.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the Public Gateway.
        """
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The Organization ID the PAT rule configuration is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> pulumi.Output[str]:
        """
        The private IP address to forward data to.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="privatePort")
    def private_port(self) -> pulumi.Output[int]:
        """
        The private port to translate to.
        """
        return pulumi.get(self, "private_port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        """
        The protocol the rule should apply to. Possible values are `both`, `tcp` and `udp`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> pulumi.Output[int]:
        """
        The public port to listen on.
        """
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the PAT rule configuration.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the Public Gateway DHCP configuration should be created.
        """
        return pulumi.get(self, "zone")

