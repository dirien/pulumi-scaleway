# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IotNetworkArgs', 'IotNetwork']

@pulumi.input_type
class IotNetworkArgs:
    def __init__(__self__, *,
                 hub_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 topic_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IotNetwork resource.
        :param pulumi.Input[str] hub_id: The hub ID to which the Network will be attached to.
        :param pulumi.Input[str] type: The network type to create (e.g. `sigfox`).
        :param pulumi.Input[str] name: The name of the IoT Network you want to create (e.g. `my-net`).
        :param pulumi.Input[str] topic_prefix: The prefix that will be prepended to all topics for this Network.
        """
        pulumi.set(__self__, "hub_id", hub_id)
        pulumi.set(__self__, "type", type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if topic_prefix is not None:
            pulumi.set(__self__, "topic_prefix", topic_prefix)

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> pulumi.Input[str]:
        """
        The hub ID to which the Network will be attached to.
        """
        return pulumi.get(self, "hub_id")

    @hub_id.setter
    def hub_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "hub_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The network type to create (e.g. `sigfox`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IoT Network you want to create (e.g. `my-net`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="topicPrefix")
    def topic_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix that will be prepended to all topics for this Network.
        """
        return pulumi.get(self, "topic_prefix")

    @topic_prefix.setter
    def topic_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_prefix", value)


@pulumi.input_type
class _IotNetworkState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 hub_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 topic_prefix: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IotNetwork resources.
        :param pulumi.Input[str] created_at: The date and time the Network was created.
        :param pulumi.Input[str] endpoint: The endpoint to use when interacting with the network.
        :param pulumi.Input[str] hub_id: The hub ID to which the Network will be attached to.
        :param pulumi.Input[str] name: The name of the IoT Network you want to create (e.g. `my-net`).
        :param pulumi.Input[str] secret: The endpoint key to keep secret.
        :param pulumi.Input[str] topic_prefix: The prefix that will be prepended to all topics for this Network.
        :param pulumi.Input[str] type: The network type to create (e.g. `sigfox`).
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if hub_id is not None:
            pulumi.set(__self__, "hub_id", hub_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if topic_prefix is not None:
            pulumi.set(__self__, "topic_prefix", topic_prefix)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the Network was created.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint to use when interacting with the network.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> Optional[pulumi.Input[str]]:
        """
        The hub ID to which the Network will be attached to.
        """
        return pulumi.get(self, "hub_id")

    @hub_id.setter
    def hub_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hub_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IoT Network you want to create (e.g. `my-net`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint key to keep secret.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter(name="topicPrefix")
    def topic_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The prefix that will be prepended to all topics for this Network.
        """
        return pulumi.get(self, "topic_prefix")

    @topic_prefix.setter
    def topic_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_prefix", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The network type to create (e.g. `sigfox`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class IotNetwork(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hub_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 topic_prefix: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_iot_hub = scaleway.IotHub("mainIotHub", product_plan="plan_shared")
        main_iot_network = scaleway.IotNetwork("mainIotNetwork",
            hub_id=main_iot_hub.id,
            type="sigfox")
        ```

        ## Import

        IoT Networks can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iotNetwork:IotNetwork net01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hub_id: The hub ID to which the Network will be attached to.
        :param pulumi.Input[str] name: The name of the IoT Network you want to create (e.g. `my-net`).
        :param pulumi.Input[str] topic_prefix: The prefix that will be prepended to all topics for this Network.
        :param pulumi.Input[str] type: The network type to create (e.g. `sigfox`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IotNetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_iot_hub = scaleway.IotHub("mainIotHub", product_plan="plan_shared")
        main_iot_network = scaleway.IotNetwork("mainIotNetwork",
            hub_id=main_iot_hub.id,
            type="sigfox")
        ```

        ## Import

        IoT Networks can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iotNetwork:IotNetwork net01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IotNetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IotNetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hub_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 topic_prefix: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IotNetworkArgs.__new__(IotNetworkArgs)

            if hub_id is None and not opts.urn:
                raise TypeError("Missing required property 'hub_id'")
            __props__.__dict__["hub_id"] = hub_id
            __props__.__dict__["name"] = name
            __props__.__dict__["topic_prefix"] = topic_prefix
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["created_at"] = None
            __props__.__dict__["endpoint"] = None
            __props__.__dict__["secret"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(IotNetwork, __self__).__init__(
            'scaleway:index/iotNetwork:IotNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            hub_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            topic_prefix: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'IotNetwork':
        """
        Get an existing IotNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time the Network was created.
        :param pulumi.Input[str] endpoint: The endpoint to use when interacting with the network.
        :param pulumi.Input[str] hub_id: The hub ID to which the Network will be attached to.
        :param pulumi.Input[str] name: The name of the IoT Network you want to create (e.g. `my-net`).
        :param pulumi.Input[str] secret: The endpoint key to keep secret.
        :param pulumi.Input[str] topic_prefix: The prefix that will be prepended to all topics for this Network.
        :param pulumi.Input[str] type: The network type to create (e.g. `sigfox`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IotNetworkState.__new__(_IotNetworkState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["hub_id"] = hub_id
        __props__.__dict__["name"] = name
        __props__.__dict__["secret"] = secret
        __props__.__dict__["topic_prefix"] = topic_prefix
        __props__.__dict__["type"] = type
        return IotNetwork(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time the Network was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint to use when interacting with the network.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="hubId")
    def hub_id(self) -> pulumi.Output[str]:
        """
        The hub ID to which the Network will be attached to.
        """
        return pulumi.get(self, "hub_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the IoT Network you want to create (e.g. `my-net`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        The endpoint key to keep secret.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter(name="topicPrefix")
    def topic_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        The prefix that will be prepended to all topics for this Network.
        """
        return pulumi.get(self, "topic_prefix")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The network type to create (e.g. `sigfox`).
        """
        return pulumi.get(self, "type")

