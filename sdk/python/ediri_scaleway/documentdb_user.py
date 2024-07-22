# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DocumentdbUserArgs', 'DocumentdbUser']

@pulumi.input_type
class DocumentdbUserArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DocumentdbUser resource.
        :param pulumi.Input[str] instance_id: UUID of the documentDB instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database Username.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "password", password)
        if is_admin is not None:
            pulumi.set(__self__, "is_admin", is_admin)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the documentDB instance.

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
        Database Username.

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
class _DocumentdbUserState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 is_admin: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DocumentdbUser resources.
        :param pulumi.Input[str] instance_id: UUID of the documentDB instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database Username.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if is_admin is not None:
            pulumi.set(__self__, "is_admin", is_admin)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the documentDB instance.

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
        Database Username.

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


class DocumentdbUser(pulumi.CustomResource):
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
        Creates and manages Scaleway Database DocumentDB Users.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_random as random

        instance = scaleway.DocumentdbInstance("instance",
            node_type="docdb-play2-pico",
            engine="FerretDB-1",
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret",
            volume_size_in_gb=20)
        db_password = random.RandomPassword("dbPassword",
            length=16,
            special=True)
        db_admin = scaleway.DocumentdbUser("dbAdmin",
            instance_id=instance.id,
            password=db_password.result,
            is_admin=True)
        ```

        ## Import

        Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/documentdbUser:DocumentdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the documentDB instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database Username.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DocumentdbUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Database DocumentDB Users.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_random as random

        instance = scaleway.DocumentdbInstance("instance",
            node_type="docdb-play2-pico",
            engine="FerretDB-1",
            user_name="my_initial_user",
            password="thiZ_is_v&ry_s3cret",
            volume_size_in_gb=20)
        db_password = random.RandomPassword("dbPassword",
            length=16,
            special=True)
        db_admin = scaleway.DocumentdbUser("dbAdmin",
            instance_id=instance.id,
            password=db_password.result,
            is_admin=True)
        ```

        ## Import

        Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/documentdbUser:DocumentdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
        ```

        :param str resource_name: The name of the resource.
        :param DocumentdbUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DocumentdbUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
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
            __props__ = DocumentdbUserArgs.__new__(DocumentdbUserArgs)

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
        super(DocumentdbUser, __self__).__init__(
            'scaleway:index/documentdbUser:DocumentdbUser',
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
            region: Optional[pulumi.Input[str]] = None) -> 'DocumentdbUser':
        """
        Get an existing DocumentdbUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: UUID of the documentDB instance.
               
               > **Important:** Updates to `instance_id` will recreate the Database User.
        :param pulumi.Input[bool] is_admin: Grant admin permissions to the Database User.
        :param pulumi.Input[str] name: Database Username.
               
               > **Important:** Updates to `name` will recreate the Database User.
        :param pulumi.Input[str] password: Database User password.
        :param pulumi.Input[str] region: The Scaleway region this resource resides in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DocumentdbUserState.__new__(_DocumentdbUserState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["is_admin"] = is_admin
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["region"] = region
        return DocumentdbUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the documentDB instance.

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
        Database Username.

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

