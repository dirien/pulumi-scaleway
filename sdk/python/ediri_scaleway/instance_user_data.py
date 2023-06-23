# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['InstanceUserDataArgs', 'InstanceUserData']

@pulumi.input_type
class InstanceUserDataArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 server_id: pulumi.Input[str],
                 value: pulumi.Input[str],
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceUserData resource.
        :param pulumi.Input[str] key: Key of the user data.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[str] value: Value associated with your key
        :param pulumi.Input[str] zone: `zone`) The zone in which the server should be created.
               
               > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
               You can define values using:
               - string
               - UTF-8 encoded file content using file
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "server_id", server_id)
        pulumi.set(__self__, "value", value)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Key of the user data.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Input[str]:
        """
        The ID of the server associated with.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Value associated with your key
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the server should be created.

        > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        You can define values using:
        - string
        - UTF-8 encoded file content using file
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceUserDataState:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceUserData resources.
        :param pulumi.Input[str] key: Key of the user data.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[str] value: Value associated with your key
        :param pulumi.Input[str] zone: `zone`) The zone in which the server should be created.
               
               > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
               You can define values using:
               - string
               - UTF-8 encoded file content using file
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key of the user data.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the server associated with.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Value associated with your key
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the server should be created.

        > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        You can define values using:
        - string
        - UTF-8 encoded file content using file
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceUserData(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Compute Instance User Data values.

        User data is a key value store API you can use to provide data from and to your server without authentication. It is the mechanism by which a user can pass information contained in a local file to an Instance at launch time.

        The typical use case is to pass something like a shell script or a configuration file as user data.

        For more information about [user_data](https://developers.scaleway.com/en/products/instance/api/#patch-9ef3ec)  check our documentation guide [here](https://www.scaleway.com/en/docs/compute/instances/how-to/use-boot-modes/#how-to-use-cloud-init).

        About cloud-init documentation please check this [link](https://cloudinit.readthedocs.io/en/latest/).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        config = pulumi.Config()
        user_data = config.get_object("userData")
        if user_data is None:
            user_data = {
                "cloud-init": \"\"\"#cloud-config
        apt-update: true
        apt-upgrade: true
        \"\"\",
                "foo": "bar",
            }
        main_instance_server = scaleway.InstanceServer("mainInstanceServer",
            image="ubuntu_focal",
            type="DEV1-S")
        # User data with a single value
        main_instance_user_data = scaleway.InstanceUserData("mainInstanceUserData",
            server_id=main_instance_server.id,
            key="foo",
            value="bar")
        # User Data with many keys.
        data = []
        for range in [{"value": i} for i in range(0, user_data)]:
            data.append(scaleway.InstanceUserData(f"data-{range['value']}",
                server_id=main_instance_server.id,
                key=range["key"],
                value=range["value"]))
        ```

        ## Import

        User data can be imported using the `{zone}/{key}/{server_id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/instanceUserData:InstanceUserData main fr-par-1/cloud-init/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: Key of the user data.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[str] value: Value associated with your key
        :param pulumi.Input[str] zone: `zone`) The zone in which the server should be created.
               
               > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
               You can define values using:
               - string
               - UTF-8 encoded file content using file
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceUserDataArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Compute Instance User Data values.

        User data is a key value store API you can use to provide data from and to your server without authentication. It is the mechanism by which a user can pass information contained in a local file to an Instance at launch time.

        The typical use case is to pass something like a shell script or a configuration file as user data.

        For more information about [user_data](https://developers.scaleway.com/en/products/instance/api/#patch-9ef3ec)  check our documentation guide [here](https://www.scaleway.com/en/docs/compute/instances/how-to/use-boot-modes/#how-to-use-cloud-init).

        About cloud-init documentation please check this [link](https://cloudinit.readthedocs.io/en/latest/).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        config = pulumi.Config()
        user_data = config.get_object("userData")
        if user_data is None:
            user_data = {
                "cloud-init": \"\"\"#cloud-config
        apt-update: true
        apt-upgrade: true
        \"\"\",
                "foo": "bar",
            }
        main_instance_server = scaleway.InstanceServer("mainInstanceServer",
            image="ubuntu_focal",
            type="DEV1-S")
        # User data with a single value
        main_instance_user_data = scaleway.InstanceUserData("mainInstanceUserData",
            server_id=main_instance_server.id,
            key="foo",
            value="bar")
        # User Data with many keys.
        data = []
        for range in [{"value": i} for i in range(0, user_data)]:
            data.append(scaleway.InstanceUserData(f"data-{range['value']}",
                server_id=main_instance_server.id,
                key=range["key"],
                value=range["value"]))
        ```

        ## Import

        User data can be imported using the `{zone}/{key}/{server_id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/instanceUserData:InstanceUserData main fr-par-1/cloud-init/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceUserDataArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceUserDataArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceUserDataArgs.__new__(InstanceUserDataArgs)

            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            if server_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_id'")
            __props__.__dict__["server_id"] = server_id
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
            __props__.__dict__["zone"] = zone
        super(InstanceUserData, __self__).__init__(
            'scaleway:index/instanceUserData:InstanceUserData',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceUserData':
        """
        Get an existing InstanceUserData resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: Key of the user data.
        :param pulumi.Input[str] server_id: The ID of the server associated with.
        :param pulumi.Input[str] value: Value associated with your key
        :param pulumi.Input[str] zone: `zone`) The zone in which the server should be created.
               
               > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
               You can define values using:
               - string
               - UTF-8 encoded file content using file
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceUserDataState.__new__(_InstanceUserDataState)

        __props__.__dict__["key"] = key
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["value"] = value
        __props__.__dict__["zone"] = zone
        return InstanceUserData(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Key of the user data.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The ID of the server associated with.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[str]:
        """
        Value associated with your key
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the server should be created.

        > **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        You can define values using:
        - string
        - UTF-8 encoded file content using file
        """
        return pulumi.get(self, "zone")

