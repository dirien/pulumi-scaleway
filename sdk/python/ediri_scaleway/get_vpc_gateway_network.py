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
    'GetVpcGatewayNetworkResult',
    'AwaitableGetVpcGatewayNetworkResult',
    'get_vpc_gateway_network',
    'get_vpc_gateway_network_output',
]

@pulumi.output_type
class GetVpcGatewayNetworkResult:
    """
    A collection of values returned by getVpcGatewayNetwork.
    """
    def __init__(__self__, cleanup_dhcp=None, created_at=None, dhcp_id=None, enable_dhcp=None, enable_masquerade=None, gateway_id=None, gateway_network_id=None, id=None, mac_address=None, private_network_id=None, static_address=None, updated_at=None, zone=None):
        if cleanup_dhcp and not isinstance(cleanup_dhcp, bool):
            raise TypeError("Expected argument 'cleanup_dhcp' to be a bool")
        pulumi.set(__self__, "cleanup_dhcp", cleanup_dhcp)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if dhcp_id and not isinstance(dhcp_id, str):
            raise TypeError("Expected argument 'dhcp_id' to be a str")
        pulumi.set(__self__, "dhcp_id", dhcp_id)
        if enable_dhcp and not isinstance(enable_dhcp, bool):
            raise TypeError("Expected argument 'enable_dhcp' to be a bool")
        pulumi.set(__self__, "enable_dhcp", enable_dhcp)
        if enable_masquerade and not isinstance(enable_masquerade, bool):
            raise TypeError("Expected argument 'enable_masquerade' to be a bool")
        pulumi.set(__self__, "enable_masquerade", enable_masquerade)
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        pulumi.set(__self__, "gateway_id", gateway_id)
        if gateway_network_id and not isinstance(gateway_network_id, str):
            raise TypeError("Expected argument 'gateway_network_id' to be a str")
        pulumi.set(__self__, "gateway_network_id", gateway_network_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mac_address and not isinstance(mac_address, str):
            raise TypeError("Expected argument 'mac_address' to be a str")
        pulumi.set(__self__, "mac_address", mac_address)
        if private_network_id and not isinstance(private_network_id, str):
            raise TypeError("Expected argument 'private_network_id' to be a str")
        pulumi.set(__self__, "private_network_id", private_network_id)
        if static_address and not isinstance(static_address, str):
            raise TypeError("Expected argument 'static_address' to be a str")
        pulumi.set(__self__, "static_address", static_address)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="cleanupDhcp")
    def cleanup_dhcp(self) -> bool:
        return pulumi.get(self, "cleanup_dhcp")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dhcpId")
    def dhcp_id(self) -> Optional[str]:
        return pulumi.get(self, "dhcp_id")

    @property
    @pulumi.getter(name="enableDhcp")
    def enable_dhcp(self) -> bool:
        return pulumi.get(self, "enable_dhcp")

    @property
    @pulumi.getter(name="enableMasquerade")
    def enable_masquerade(self) -> Optional[bool]:
        return pulumi.get(self, "enable_masquerade")

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> Optional[str]:
        return pulumi.get(self, "gateway_id")

    @property
    @pulumi.getter(name="gatewayNetworkId")
    def gateway_network_id(self) -> Optional[str]:
        return pulumi.get(self, "gateway_network_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="macAddress")
    def mac_address(self) -> str:
        return pulumi.get(self, "mac_address")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[str]:
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter(name="staticAddress")
    def static_address(self) -> str:
        return pulumi.get(self, "static_address")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetVpcGatewayNetworkResult(GetVpcGatewayNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcGatewayNetworkResult(
            cleanup_dhcp=self.cleanup_dhcp,
            created_at=self.created_at,
            dhcp_id=self.dhcp_id,
            enable_dhcp=self.enable_dhcp,
            enable_masquerade=self.enable_masquerade,
            gateway_id=self.gateway_id,
            gateway_network_id=self.gateway_network_id,
            id=self.id,
            mac_address=self.mac_address,
            private_network_id=self.private_network_id,
            static_address=self.static_address,
            updated_at=self.updated_at,
            zone=self.zone)


def get_vpc_gateway_network(dhcp_id: Optional[str] = None,
                            enable_masquerade: Optional[bool] = None,
                            gateway_id: Optional[str] = None,
                            gateway_network_id: Optional[str] = None,
                            private_network_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcGatewayNetworkResult:
    """
    Gets information about a gateway network.

    ## Example Usage

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.VpcGatewayNetwork("main",
        gateway_id=scaleway_vpc_public_gateway["pg01"]["id"],
        private_network_id=scaleway_vpc_private_network["pn01"]["id"],
        dhcp_id=scaleway_vpc_public_gateway_dhcp["dhcp01"]["id"],
        cleanup_dhcp=True,
        enable_masquerade=True)
    by_id = scaleway.get_vpc_gateway_network_output(gateway_network_id=main.id)
    by_gateway_and_pn = scaleway.get_vpc_gateway_network(gateway_id=scaleway_vpc_public_gateway["pg01"]["id"],
        private_network_id=scaleway_vpc_private_network["pn01"]["id"])
    ```


    :param str dhcp_id: ID of the public gateway DHCP config
    :param bool enable_masquerade: If masquerade is enabled on requested network
    :param str gateway_id: ID of the public gateway the gateway network is linked to
    :param str gateway_network_id: ID of the gateway network.
           
           > Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
    :param str private_network_id: ID of the private network the gateway network is linked to
    """
    __args__ = dict()
    __args__['dhcpId'] = dhcp_id
    __args__['enableMasquerade'] = enable_masquerade
    __args__['gatewayId'] = gateway_id
    __args__['gatewayNetworkId'] = gateway_network_id
    __args__['privateNetworkId'] = private_network_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getVpcGatewayNetwork:getVpcGatewayNetwork', __args__, opts=opts, typ=GetVpcGatewayNetworkResult).value

    return AwaitableGetVpcGatewayNetworkResult(
        cleanup_dhcp=pulumi.get(__ret__, 'cleanup_dhcp'),
        created_at=pulumi.get(__ret__, 'created_at'),
        dhcp_id=pulumi.get(__ret__, 'dhcp_id'),
        enable_dhcp=pulumi.get(__ret__, 'enable_dhcp'),
        enable_masquerade=pulumi.get(__ret__, 'enable_masquerade'),
        gateway_id=pulumi.get(__ret__, 'gateway_id'),
        gateway_network_id=pulumi.get(__ret__, 'gateway_network_id'),
        id=pulumi.get(__ret__, 'id'),
        mac_address=pulumi.get(__ret__, 'mac_address'),
        private_network_id=pulumi.get(__ret__, 'private_network_id'),
        static_address=pulumi.get(__ret__, 'static_address'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_vpc_gateway_network)
def get_vpc_gateway_network_output(dhcp_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   enable_masquerade: Optional[pulumi.Input[Optional[bool]]] = None,
                                   gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   gateway_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   private_network_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcGatewayNetworkResult]:
    """
    Gets information about a gateway network.

    ## Example Usage

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.VpcGatewayNetwork("main",
        gateway_id=scaleway_vpc_public_gateway["pg01"]["id"],
        private_network_id=scaleway_vpc_private_network["pn01"]["id"],
        dhcp_id=scaleway_vpc_public_gateway_dhcp["dhcp01"]["id"],
        cleanup_dhcp=True,
        enable_masquerade=True)
    by_id = scaleway.get_vpc_gateway_network_output(gateway_network_id=main.id)
    by_gateway_and_pn = scaleway.get_vpc_gateway_network(gateway_id=scaleway_vpc_public_gateway["pg01"]["id"],
        private_network_id=scaleway_vpc_private_network["pn01"]["id"])
    ```


    :param str dhcp_id: ID of the public gateway DHCP config
    :param bool enable_masquerade: If masquerade is enabled on requested network
    :param str gateway_id: ID of the public gateway the gateway network is linked to
    :param str gateway_network_id: ID of the gateway network.
           
           > Only one of `gateway_network_id` or filters should be specified. You can use all the filters you want.
    :param str private_network_id: ID of the private network the gateway network is linked to
    """
    ...
