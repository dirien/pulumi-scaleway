# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FlexibleIpMacAddressArgs', 'FlexibleIpMacAddress']

@pulumi.input_type
class FlexibleIpMacAddressArgs:
    def __init__(__self__, *,
                 flexible_ip_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 flexible_ip_ids_to_duplicates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlexibleIpMacAddress resource.
        :param pulumi.Input[str] flexible_ip_id: The ID of the flexible IP for which to generate a virtual MAC.
        :param pulumi.Input[str] type: The type of the virtual MAC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] flexible_ip_ids_to_duplicates: The IDs of the flexible IPs on which to duplicate the virtual MAC.
               > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        :param pulumi.Input[str] zone: The zone of the Virtual Mac Address.
        """
        pulumi.set(__self__, "flexible_ip_id", flexible_ip_id)
        pulumi.set(__self__, "type", type)
        if flexible_ip_ids_to_duplicates is not None:
            pulumi.set(__self__, "flexible_ip_ids_to_duplicates", flexible_ip_ids_to_duplicates)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="flexibleIpId")
    def flexible_ip_id(self) -> pulumi.Input[str]:
        """
        The ID of the flexible IP for which to generate a virtual MAC.
        """
        return pulumi.get(self, "flexible_ip_id")

    @flexible_ip_id.setter
    def flexible_ip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "flexible_ip_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the virtual MAC.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="flexibleIpIdsToDuplicates")
    def flexible_ip_ids_to_duplicates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of the flexible IPs on which to duplicate the virtual MAC.
        > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        """
        return pulumi.get(self, "flexible_ip_ids_to_duplicates")

    @flexible_ip_ids_to_duplicates.setter
    def flexible_ip_ids_to_duplicates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "flexible_ip_ids_to_duplicates", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the Virtual Mac Address.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _FlexibleIpMacAddressState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 flexible_ip_id: Optional[pulumi.Input[str]] = None,
                 flexible_ip_ids_to_duplicates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FlexibleIpMacAddress resources.
        :param pulumi.Input[str] address: The Virtual MAC address.
        :param pulumi.Input[str] created_at: The date at which the Virtual Mac Address was created (RFC 3339 format).
        :param pulumi.Input[str] flexible_ip_id: The ID of the flexible IP for which to generate a virtual MAC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] flexible_ip_ids_to_duplicates: The IDs of the flexible IPs on which to duplicate the virtual MAC.
               > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        :param pulumi.Input[str] status: The Virtual MAC status.
        :param pulumi.Input[str] type: The type of the virtual MAC.
        :param pulumi.Input[str] updated_at: The date at which the Virtual Mac Address was last updated (RFC 3339 format).
        :param pulumi.Input[str] zone: The zone of the Virtual Mac Address.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if flexible_ip_id is not None:
            pulumi.set(__self__, "flexible_ip_id", flexible_ip_id)
        if flexible_ip_ids_to_duplicates is not None:
            pulumi.set(__self__, "flexible_ip_ids_to_duplicates", flexible_ip_ids_to_duplicates)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        The Virtual MAC address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date at which the Virtual Mac Address was created (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="flexibleIpId")
    def flexible_ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the flexible IP for which to generate a virtual MAC.
        """
        return pulumi.get(self, "flexible_ip_id")

    @flexible_ip_id.setter
    def flexible_ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flexible_ip_id", value)

    @property
    @pulumi.getter(name="flexibleIpIdsToDuplicates")
    def flexible_ip_ids_to_duplicates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The IDs of the flexible IPs on which to duplicate the virtual MAC.
        > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        """
        return pulumi.get(self, "flexible_ip_ids_to_duplicates")

    @flexible_ip_ids_to_duplicates.setter
    def flexible_ip_ids_to_duplicates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "flexible_ip_ids_to_duplicates", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The Virtual MAC status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the virtual MAC.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date at which the Virtual Mac Address was last updated (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the Virtual Mac Address.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class FlexibleIpMacAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flexible_ip_id: Optional[pulumi.Input[str]] = None,
                 flexible_ip_ids_to_duplicates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Flexible IP Mac Addresses.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/flexible-ip/api).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_flexible_ip = scaleway.FlexibleIp("mainFlexibleIp")
        main_flexible_ip_mac_address = scaleway.FlexibleIpMacAddress("mainFlexibleIpMacAddress",
            flexible_ip_id=main_flexible_ip.id,
            type="kvm")
        ```

        ### Duplicate on many other flexible IPs

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        my_offer = scaleway.get_baremetal_offer(name="EM-B112X-SSD")
        base = scaleway.BaremetalServer("base",
            offer=my_offer.offer_id,
            install_config_afterward=True)
        ip01 = scaleway.FlexibleIp("ip01", server_id=base.id)
        ip02 = scaleway.FlexibleIp("ip02", server_id=base.id)
        ip03 = scaleway.FlexibleIp("ip03", server_id=base.id)
        main = scaleway.FlexibleIpMacAddress("main",
            flexible_ip_id=ip01.id,
            type="kvm",
            flexible_ip_ids_to_duplicates=[
                ip02.id,
                ip03.id,
            ])
        ```

        ## Import

        Flexible IP Mac Addresses can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flexible_ip_id: The ID of the flexible IP for which to generate a virtual MAC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] flexible_ip_ids_to_duplicates: The IDs of the flexible IPs on which to duplicate the virtual MAC.
               > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        :param pulumi.Input[str] type: The type of the virtual MAC.
        :param pulumi.Input[str] zone: The zone of the Virtual Mac Address.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlexibleIpMacAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Flexible IP Mac Addresses.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/flexible-ip/api).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_flexible_ip = scaleway.FlexibleIp("mainFlexibleIp")
        main_flexible_ip_mac_address = scaleway.FlexibleIpMacAddress("mainFlexibleIpMacAddress",
            flexible_ip_id=main_flexible_ip.id,
            type="kvm")
        ```

        ### Duplicate on many other flexible IPs

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        my_offer = scaleway.get_baremetal_offer(name="EM-B112X-SSD")
        base = scaleway.BaremetalServer("base",
            offer=my_offer.offer_id,
            install_config_afterward=True)
        ip01 = scaleway.FlexibleIp("ip01", server_id=base.id)
        ip02 = scaleway.FlexibleIp("ip02", server_id=base.id)
        ip03 = scaleway.FlexibleIp("ip03", server_id=base.id)
        main = scaleway.FlexibleIpMacAddress("main",
            flexible_ip_id=ip01.id,
            type="kvm",
            flexible_ip_ids_to_duplicates=[
                ip02.id,
                ip03.id,
            ])
        ```

        ## Import

        Flexible IP Mac Addresses can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param FlexibleIpMacAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlexibleIpMacAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flexible_ip_id: Optional[pulumi.Input[str]] = None,
                 flexible_ip_ids_to_duplicates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlexibleIpMacAddressArgs.__new__(FlexibleIpMacAddressArgs)

            if flexible_ip_id is None and not opts.urn:
                raise TypeError("Missing required property 'flexible_ip_id'")
            __props__.__dict__["flexible_ip_id"] = flexible_ip_id
            __props__.__dict__["flexible_ip_ids_to_duplicates"] = flexible_ip_ids_to_duplicates
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["zone"] = zone
            __props__.__dict__["address"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(FlexibleIpMacAddress, __self__).__init__(
            'scaleway:index/flexibleIpMacAddress:FlexibleIpMacAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            flexible_ip_id: Optional[pulumi.Input[str]] = None,
            flexible_ip_ids_to_duplicates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'FlexibleIpMacAddress':
        """
        Get an existing FlexibleIpMacAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: The Virtual MAC address.
        :param pulumi.Input[str] created_at: The date at which the Virtual Mac Address was created (RFC 3339 format).
        :param pulumi.Input[str] flexible_ip_id: The ID of the flexible IP for which to generate a virtual MAC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] flexible_ip_ids_to_duplicates: The IDs of the flexible IPs on which to duplicate the virtual MAC.
               > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        :param pulumi.Input[str] status: The Virtual MAC status.
        :param pulumi.Input[str] type: The type of the virtual MAC.
        :param pulumi.Input[str] updated_at: The date at which the Virtual Mac Address was last updated (RFC 3339 format).
        :param pulumi.Input[str] zone: The zone of the Virtual Mac Address.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlexibleIpMacAddressState.__new__(_FlexibleIpMacAddressState)

        __props__.__dict__["address"] = address
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["flexible_ip_id"] = flexible_ip_id
        __props__.__dict__["flexible_ip_ids_to_duplicates"] = flexible_ip_ids_to_duplicates
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["zone"] = zone
        return FlexibleIpMacAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        The Virtual MAC address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date at which the Virtual Mac Address was created (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="flexibleIpId")
    def flexible_ip_id(self) -> pulumi.Output[str]:
        """
        The ID of the flexible IP for which to generate a virtual MAC.
        """
        return pulumi.get(self, "flexible_ip_id")

    @property
    @pulumi.getter(name="flexibleIpIdsToDuplicates")
    def flexible_ip_ids_to_duplicates(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IDs of the flexible IPs on which to duplicate the virtual MAC.
        > **Important:** The flexible IPs need to be attached to the same server for the operation to work.
        """
        return pulumi.get(self, "flexible_ip_ids_to_duplicates")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The Virtual MAC status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the virtual MAC.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date at which the Virtual Mac Address was last updated (RFC 3339 format).
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone of the Virtual Mac Address.
        """
        return pulumi.get(self, "zone")

