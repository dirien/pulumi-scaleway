# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RdbDatabaseBackupArgs', 'RdbDatabaseBackup']

@pulumi.input_type
class RdbDatabaseBackupArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 expires_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RdbDatabaseBackup resource.
        :param pulumi.Input[str] database_name: Name of the database of this backup.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Backup.
        :param pulumi.Input[str] expires_at: Expiration date (Format ISO 8601).
               
               > **Important:** `expires_at` cannot be removed after being set.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-database`).
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        RdbDatabaseBackupArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            database_name=database_name,
            instance_id=instance_id,
            expires_at=expires_at,
            name=name,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             database_name: pulumi.Input[str],
             instance_id: pulumi.Input[str],
             expires_at: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'databaseName' in kwargs:
            database_name = kwargs['databaseName']
        if 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if 'expiresAt' in kwargs:
            expires_at = kwargs['expiresAt']

        _setter("database_name", database_name)
        _setter("instance_id", instance_id)
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Name of the database of this backup.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Backup.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date (Format ISO 8601).

        > **Important:** `expires_at` cannot be removed after being set.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-database`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _RdbDatabaseBackupState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RdbDatabaseBackup resources.
        :param pulumi.Input[str] created_at: Creation date (Format ISO 8601).
        :param pulumi.Input[str] database_name: Name of the database of this backup.
        :param pulumi.Input[str] expires_at: Expiration date (Format ISO 8601).
               
               > **Important:** `expires_at` cannot be removed after being set.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Backup.
        :param pulumi.Input[str] instance_name: Name of the instance of the backup.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-database`).
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[int] size: Size of the backup (in bytes).
        :param pulumi.Input[str] updated_at: Updated date (Format ISO 8601).
        """
        _RdbDatabaseBackupState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            created_at=created_at,
            database_name=database_name,
            expires_at=expires_at,
            instance_id=instance_id,
            instance_name=instance_name,
            name=name,
            region=region,
            size=size,
            updated_at=updated_at,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             created_at: Optional[pulumi.Input[str]] = None,
             database_name: Optional[pulumi.Input[str]] = None,
             expires_at: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             instance_name: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             size: Optional[pulumi.Input[int]] = None,
             updated_at: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'createdAt' in kwargs:
            created_at = kwargs['createdAt']
        if 'databaseName' in kwargs:
            database_name = kwargs['databaseName']
        if 'expiresAt' in kwargs:
            expires_at = kwargs['expiresAt']
        if 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if 'instanceName' in kwargs:
            instance_name = kwargs['instanceName']
        if 'updatedAt' in kwargs:
            updated_at = kwargs['updatedAt']

        if created_at is not None:
            _setter("created_at", created_at)
        if database_name is not None:
            _setter("database_name", database_name)
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if instance_id is not None:
            _setter("instance_id", instance_id)
        if instance_name is not None:
            _setter("instance_name", instance_name)
        if name is not None:
            _setter("name", name)
        if region is not None:
            _setter("region", region)
        if size is not None:
            _setter("size", size)
        if updated_at is not None:
            _setter("updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Creation date (Format ISO 8601).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database of this backup.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        Expiration date (Format ISO 8601).

        > **Important:** `expires_at` cannot be removed after being set.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Backup.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the instance of the backup.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-database`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Size of the backup (in bytes).
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Updated date (Format ISO 8601).
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class RdbDatabaseBackup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway RDB database backup.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.RdbDatabaseBackup("main",
            instance_id=data["scaleway_rdb_instance"]["main"]["id"],
            database_name=data["scaleway_rdb_database"]["main"]["name"])
        ```

        ### With expiration

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.RdbDatabaseBackup("main",
            instance_id=data["scaleway_rdb_instance"]["main"]["id"],
            database_name=data["scaleway_rdb_database"]["main"]["name"],
            expires_at="2022-06-16T07:48:44Z")
        ```

        ## Import

        RDB Database can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup mybackup fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Name of the database of this backup.
        :param pulumi.Input[str] expires_at: Expiration date (Format ISO 8601).
               
               > **Important:** `expires_at` cannot be removed after being set.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Backup.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-database`).
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RdbDatabaseBackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway RDB database backup.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.RdbDatabaseBackup("main",
            instance_id=data["scaleway_rdb_instance"]["main"]["id"],
            database_name=data["scaleway_rdb_database"]["main"]["name"])
        ```

        ### With expiration

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.RdbDatabaseBackup("main",
            instance_id=data["scaleway_rdb_instance"]["main"]["id"],
            database_name=data["scaleway_rdb_database"]["main"]["name"],
            expires_at="2022-06-16T07:48:44Z")
        ```

        ## Import

        RDB Database can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup mybackup fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param RdbDatabaseBackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RdbDatabaseBackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RdbDatabaseBackupArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RdbDatabaseBackupArgs.__new__(RdbDatabaseBackupArgs)

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["expires_at"] = expires_at
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            __props__.__dict__["created_at"] = None
            __props__.__dict__["instance_name"] = None
            __props__.__dict__["size"] = None
            __props__.__dict__["updated_at"] = None
        super(RdbDatabaseBackup, __self__).__init__(
            'scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'RdbDatabaseBackup':
        """
        Get an existing RdbDatabaseBackup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: Creation date (Format ISO 8601).
        :param pulumi.Input[str] database_name: Name of the database of this backup.
        :param pulumi.Input[str] expires_at: Expiration date (Format ISO 8601).
               
               > **Important:** `expires_at` cannot be removed after being set.
        :param pulumi.Input[str] instance_id: UUID of the rdb instance.
               
               > **Important:** Updates to `instance_id` will recreate the Backup.
        :param pulumi.Input[str] instance_name: Name of the instance of the backup.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-database`).
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        :param pulumi.Input[int] size: Size of the backup (in bytes).
        :param pulumi.Input[str] updated_at: Updated date (Format ISO 8601).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RdbDatabaseBackupState.__new__(_RdbDatabaseBackupState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_name"] = instance_name
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["size"] = size
        __props__.__dict__["updated_at"] = updated_at
        return RdbDatabaseBackup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Creation date (Format ISO 8601).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        Name of the database of this backup.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        Expiration date (Format ISO 8601).

        > **Important:** `expires_at` cannot be removed after being set.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        UUID of the rdb instance.

        > **Important:** Updates to `instance_id` will recreate the Backup.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> pulumi.Output[str]:
        """
        Name of the instance of the backup.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the database (e.g. `my-database`).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Size of the backup (in bytes).
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Updated date (Format ISO 8601).
        """
        return pulumi.get(self, "updated_at")

