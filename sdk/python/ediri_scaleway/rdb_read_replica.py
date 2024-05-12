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

__all__ = ['RdbReadReplicaArgs', 'RdbReadReplica']

@pulumi.input_type
class RdbReadReplicaArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 direct_access: Optional[pulumi.Input['RdbReadReplicaDirectAccessArgs']] = None,
                 private_network: Optional[pulumi.Input['RdbReadReplicaPrivateNetworkArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 same_zone: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a RdbReadReplica resource.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        :param pulumi.Input['RdbReadReplicaDirectAccessArgs'] direct_access: Creates a direct access endpoint to rdb replica.
        :param pulumi.Input['RdbReadReplicaPrivateNetworkArgs'] private_network: Create an endpoint in a private network.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database read replica should be created.
        :param pulumi.Input[bool] same_zone: Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if direct_access is not None:
            pulumi.set(__self__, "direct_access", direct_access)
        if private_network is not None:
            pulumi.set(__self__, "private_network", private_network)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if same_zone is not None:
            pulumi.set(__self__, "same_zone", same_zone)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the rdb instance.

        > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="directAccess")
    def direct_access(self) -> Optional[pulumi.Input['RdbReadReplicaDirectAccessArgs']]:
        """
        Creates a direct access endpoint to rdb replica.
        """
        return pulumi.get(self, "direct_access")

    @direct_access.setter
    def direct_access(self, value: Optional[pulumi.Input['RdbReadReplicaDirectAccessArgs']]):
        pulumi.set(self, "direct_access", value)

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> Optional[pulumi.Input['RdbReadReplicaPrivateNetworkArgs']]:
        """
        Create an endpoint in a private network.
        """
        return pulumi.get(self, "private_network")

    @private_network.setter
    def private_network(self, value: Optional[pulumi.Input['RdbReadReplicaPrivateNetworkArgs']]):
        pulumi.set(self, "private_network", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the Database read replica should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sameZone")
    def same_zone(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        return pulumi.get(self, "same_zone")

    @same_zone.setter
    def same_zone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "same_zone", value)


@pulumi.input_type
class _RdbReadReplicaState:
    def __init__(__self__, *,
                 direct_access: Optional[pulumi.Input['RdbReadReplicaDirectAccessArgs']] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 private_network: Optional[pulumi.Input['RdbReadReplicaPrivateNetworkArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 same_zone: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering RdbReadReplica resources.
        :param pulumi.Input['RdbReadReplicaDirectAccessArgs'] direct_access: Creates a direct access endpoint to rdb replica.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        :param pulumi.Input['RdbReadReplicaPrivateNetworkArgs'] private_network: Create an endpoint in a private network.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database read replica should be created.
        :param pulumi.Input[bool] same_zone: Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        if direct_access is not None:
            pulumi.set(__self__, "direct_access", direct_access)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if private_network is not None:
            pulumi.set(__self__, "private_network", private_network)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if same_zone is not None:
            pulumi.set(__self__, "same_zone", same_zone)

    @property
    @pulumi.getter(name="directAccess")
    def direct_access(self) -> Optional[pulumi.Input['RdbReadReplicaDirectAccessArgs']]:
        """
        Creates a direct access endpoint to rdb replica.
        """
        return pulumi.get(self, "direct_access")

    @direct_access.setter
    def direct_access(self, value: Optional[pulumi.Input['RdbReadReplicaDirectAccessArgs']]):
        pulumi.set(self, "direct_access", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the rdb instance.

        > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> Optional[pulumi.Input['RdbReadReplicaPrivateNetworkArgs']]:
        """
        Create an endpoint in a private network.
        """
        return pulumi.get(self, "private_network")

    @private_network.setter
    def private_network(self, value: Optional[pulumi.Input['RdbReadReplicaPrivateNetworkArgs']]):
        pulumi.set(self, "private_network", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the Database read replica should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="sameZone")
    def same_zone(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        return pulumi.get(self, "same_zone")

    @same_zone.setter
    def same_zone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "same_zone", value)


class RdbReadReplica(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 direct_access: Optional[pulumi.Input[pulumi.InputType['RdbReadReplicaDirectAccessArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 private_network: Optional[pulumi.Input[pulumi.InputType['RdbReadReplicaPrivateNetworkArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 same_zone: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Database read replicas.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        instance = scaleway.RdbInstance("instance",
            node_type="db-dev-s",
            engine="PostgreSQL-14",
            is_ha_cluster=False,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret",
            tags=[
                "terraform-test",
                "scaleway_rdb_read_replica",
                "minimal",
            ])
        replica = scaleway.RdbReadReplica("replica",
            instance_id=instance.id,
            direct_access=scaleway.RdbReadReplicaDirectAccessArgs())
        ```

        ### Private network

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        instance = scaleway.RdbInstance("instance",
            node_type="db-dev-s",
            engine="PostgreSQL-14",
            is_ha_cluster=False,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret")
        pn = scaleway.VpcPrivateNetwork("pn")
        replica = scaleway.RdbReadReplica("replica",
            instance_id=instance.id,
            private_network=scaleway.RdbReadReplicaPrivateNetworkArgs(
                private_network_id=pn.id,
                service_ip="192.168.1.254/24",
            ))
        ```

        ## Import

        Database Read replica can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/rdbReadReplica:RdbReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RdbReadReplicaDirectAccessArgs']] direct_access: Creates a direct access endpoint to rdb replica.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        :param pulumi.Input[pulumi.InputType['RdbReadReplicaPrivateNetworkArgs']] private_network: Create an endpoint in a private network.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database read replica should be created.
        :param pulumi.Input[bool] same_zone: Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RdbReadReplicaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Database read replicas.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        instance = scaleway.RdbInstance("instance",
            node_type="db-dev-s",
            engine="PostgreSQL-14",
            is_ha_cluster=False,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret",
            tags=[
                "terraform-test",
                "scaleway_rdb_read_replica",
                "minimal",
            ])
        replica = scaleway.RdbReadReplica("replica",
            instance_id=instance.id,
            direct_access=scaleway.RdbReadReplicaDirectAccessArgs())
        ```

        ### Private network

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        instance = scaleway.RdbInstance("instance",
            node_type="db-dev-s",
            engine="PostgreSQL-14",
            is_ha_cluster=False,
            disable_backup=True,
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret")
        pn = scaleway.VpcPrivateNetwork("pn")
        replica = scaleway.RdbReadReplica("replica",
            instance_id=instance.id,
            private_network=scaleway.RdbReadReplicaPrivateNetworkArgs(
                private_network_id=pn.id,
                service_ip="192.168.1.254/24",
            ))
        ```

        ## Import

        Database Read replica can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/rdbReadReplica:RdbReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param RdbReadReplicaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RdbReadReplicaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 direct_access: Optional[pulumi.Input[pulumi.InputType['RdbReadReplicaDirectAccessArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 private_network: Optional[pulumi.Input[pulumi.InputType['RdbReadReplicaPrivateNetworkArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 same_zone: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RdbReadReplicaArgs.__new__(RdbReadReplicaArgs)

            __props__.__dict__["direct_access"] = direct_access
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["private_network"] = private_network
            __props__.__dict__["region"] = region
            __props__.__dict__["same_zone"] = same_zone
        super(RdbReadReplica, __self__).__init__(
            'scaleway:index/rdbReadReplica:RdbReadReplica',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            direct_access: Optional[pulumi.Input[pulumi.InputType['RdbReadReplicaDirectAccessArgs']]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            private_network: Optional[pulumi.Input[pulumi.InputType['RdbReadReplicaPrivateNetworkArgs']]] = None,
            region: Optional[pulumi.Input[str]] = None,
            same_zone: Optional[pulumi.Input[bool]] = None) -> 'RdbReadReplica':
        """
        Get an existing RdbReadReplica resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RdbReadReplicaDirectAccessArgs']] direct_access: Creates a direct access endpoint to rdb replica.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        :param pulumi.Input[pulumi.InputType['RdbReadReplicaPrivateNetworkArgs']] private_network: Create an endpoint in a private network.
        :param pulumi.Input[str] region: `region`) The region
               in which the Database read replica should be created.
        :param pulumi.Input[bool] same_zone: Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RdbReadReplicaState.__new__(_RdbReadReplicaState)

        __props__.__dict__["direct_access"] = direct_access
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["private_network"] = private_network
        __props__.__dict__["region"] = region
        __props__.__dict__["same_zone"] = same_zone
        return RdbReadReplica(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="directAccess")
    def direct_access(self) -> pulumi.Output[Optional['outputs.RdbReadReplicaDirectAccess']]:
        """
        Creates a direct access endpoint to rdb replica.
        """
        return pulumi.get(self, "direct_access")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the rdb instance.

        > **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="privateNetwork")
    def private_network(self) -> pulumi.Output[Optional['outputs.RdbReadReplicaPrivateNetwork']]:
        """
        Create an endpoint in a private network.
        """
        return pulumi.get(self, "private_network")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region
        in which the Database read replica should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="sameZone")
    def same_zone(self) -> pulumi.Output[Optional[bool]]:
        """
        Defines whether to create the replica in the same availability zone as the main instance nodes or not.
        """
        return pulumi.get(self, "same_zone")

