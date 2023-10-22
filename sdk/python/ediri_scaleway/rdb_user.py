# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RdbUserArgs', 'RdbUser']

@pulumi.input_type
class RdbUserArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RdbUser resource.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database User name.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        RdbUserArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            instance_id=instance_id,
            password=password,
            is_admin=is_admin,
            name=name,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             instance_id: pulumi.Input[str],
             password: pulumi.Input[str],
             is_admin: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if 'isAdmin' in kwargs:
            is_admin = kwargs['isAdmin']

        _setter("instance_id", instance_id)
        _setter("password", password)
        if is_admin is not None:
            _setter("is_admin", is_admin)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database User.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        Database User password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> Optional[pulumi.Input[bool]]:
        """
        Grant admin permissions to the Database User.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Database User name.

        > **Important:** Updates to `name` will recreate the Database User.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The Scaleway region this resource resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _RdbUserState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RdbUser resources.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database User name.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        _RdbUserState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            instance_id=instance_id,
            is_admin=is_admin,
            name=name,
            password=password,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             instance_id: Optional[pulumi.Input[str]] = None,
             is_admin: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             password: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if 'isAdmin' in kwargs:
            is_admin = kwargs['isAdmin']

        if instance_id is not None:
            _setter("instance_id", instance_id)
        if is_admin is not None:
            _setter("is_admin", is_admin)
        if name is not None:
            _setter("name", name)
        if password is not None:
            _setter("password", password)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database User.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> Optional[pulumi.Input[bool]]:
        """
        Grant admin permissions to the Database User.
        """
        return pulumi.get(self, "is_admin")

    @is_admin.setter
    def is_admin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_admin", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Database User name.

        > **Important:** Updates to `name` will recreate the Database User.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        Database User password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The Scaleway region this resource resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class RdbUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Database Users.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_random as random

        db_password = random.RandomPassword("dbPassword",
            length=16,
            special=True)
        db_admin = scaleway.RdbUser("dbAdmin",
            instance_id=scaleway_rdb_instance["main"]["id"],
            password=db_password.result,
            is_admin=True)
        ```

        ## Import

        Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/rdbUser:RdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database User name.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RdbUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Database Users.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_random as random

        db_password = random.RandomPassword("dbPassword",
            length=16,
            special=True)
        db_admin = scaleway.RdbUser("dbAdmin",
            instance_id=scaleway_rdb_instance["main"]["id"],
            password=db_password.result,
            is_admin=True)
        ```

        ## Import

        Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/rdbUser:RdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
        ```

        :param str resource_name: The name of the resource.
        :param RdbUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RdbUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RdbUserArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RdbUserArgs.__new__(RdbUserArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["is_admin"] = is_admin
            __props__.__dict__["name"] = name
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["region"] = region
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(RdbUser, __self__).__init__(
            'scaleway:index/rdbUser:RdbUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            is_admin: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'RdbUser':
        """
        Get an existing RdbUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database User name.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RdbUserState.__new__(_RdbUserState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_admin"] = is_admin
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["region"] = region
        return RdbUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Database User.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="isAdmin")
    def is_admin(self) -> pulumi.Output[Optional[bool]]:
        """
        Grant admin permissions to the Database User.
        """
        return pulumi.get(self, "is_admin")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Database User name.

        > **Important:** Updates to `name` will recreate the Database User.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Database User password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The Scaleway region this resource resides in.
        """
        return pulumi.get(self, "region")

