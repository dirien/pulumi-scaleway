# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SdbSqlDatabaseArgs', 'SdbSqlDatabase']

@pulumi.input_type
class SdbSqlDatabaseArgs:
    def __init__(__self__, *,
                 max_cpu: Optional[pulumi.Input[int]] = None,
                 min_cpu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SdbSqlDatabase resource.
        :param pulumi.Input[int] max_cpu: The maximum number of CPU units for your database. Defaults to 15.
        :param pulumi.Input[int] min_cpu: The minimum number of CPU units for your database. Defaults to 0.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
               
               > **Important:** Updates to `name` will recreate the database.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        if max_cpu is not None:
            pulumi.set(__self__, "max_cpu", max_cpu)
        if min_cpu is not None:
            pulumi.set(__self__, "min_cpu", min_cpu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="maxCpu")
    def max_cpu(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of CPU units for your database. Defaults to 15.
        """
        return pulumi.get(self, "max_cpu")

    @max_cpu.setter
    def max_cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_cpu", value)

    @property
    @pulumi.getter(name="minCpu")
    def min_cpu(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of CPU units for your database. Defaults to 0.
        """
        return pulumi.get(self, "min_cpu")

    @min_cpu.setter
    def min_cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_cpu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-new-database`).

        > **Important:** Updates to `name` will recreate the database.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

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
class _SdbSqlDatabaseState:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 max_cpu: Optional[pulumi.Input[int]] = None,
                 min_cpu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SdbSqlDatabase resources.
        :param pulumi.Input[str] endpoint: Endpoint of the database
        :param pulumi.Input[int] max_cpu: The maximum number of CPU units for your database. Defaults to 15.
        :param pulumi.Input[int] min_cpu: The minimum number of CPU units for your database. Defaults to 0.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
               
               > **Important:** Updates to `name` will recreate the database.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if max_cpu is not None:
            pulumi.set(__self__, "max_cpu", max_cpu)
        if min_cpu is not None:
            pulumi.set(__self__, "min_cpu", min_cpu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint of the database
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="maxCpu")
    def max_cpu(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum number of CPU units for your database. Defaults to 15.
        """
        return pulumi.get(self, "max_cpu")

    @max_cpu.setter
    def max_cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_cpu", value)

    @property
    @pulumi.getter(name="minCpu")
    def min_cpu(self) -> Optional[pulumi.Input[int]]:
        """
        The minimum number of CPU units for your database. Defaults to 0.
        """
        return pulumi.get(self, "min_cpu")

    @min_cpu.setter
    def min_cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_cpu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database (e.g. `my-new-database`).

        > **Important:** Updates to `name` will recreate the database.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

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


class SdbSqlDatabase(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_cpu: Optional[pulumi.Input[int]] = None,
                 min_cpu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Serverless SQL Databases. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/serverless-databases/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        database = scaleway.SdbSqlDatabase("database",
            max_cpu=8,
            min_cpu=0)
        ```

        ## Import

        Serverless SQL Database can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/sdbSqlDatabase:SdbSqlDatabase database fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] max_cpu: The maximum number of CPU units for your database. Defaults to 15.
        :param pulumi.Input[int] min_cpu: The minimum number of CPU units for your database. Defaults to 0.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
               
               > **Important:** Updates to `name` will recreate the database.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SdbSqlDatabaseArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Serverless SQL Databases. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/serverless-databases/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        database = scaleway.SdbSqlDatabase("database",
            max_cpu=8,
            min_cpu=0)
        ```

        ## Import

        Serverless SQL Database can be imported using the `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/sdbSqlDatabase:SdbSqlDatabase database fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param SdbSqlDatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SdbSqlDatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_cpu: Optional[pulumi.Input[int]] = None,
                 min_cpu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SdbSqlDatabaseArgs.__new__(SdbSqlDatabaseArgs)

            __props__.__dict__["max_cpu"] = max_cpu
            __props__.__dict__["min_cpu"] = min_cpu
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["endpoint"] = None
        super(SdbSqlDatabase, __self__).__init__(
            'scaleway:index/sdbSqlDatabase:SdbSqlDatabase',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            max_cpu: Optional[pulumi.Input[int]] = None,
            min_cpu: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'SdbSqlDatabase':
        """
        Get an existing SdbSqlDatabase resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: Endpoint of the database
        :param pulumi.Input[int] max_cpu: The maximum number of CPU units for your database. Defaults to 15.
        :param pulumi.Input[int] min_cpu: The minimum number of CPU units for your database. Defaults to 0.
        :param pulumi.Input[str] name: Name of the database (e.g. `my-new-database`).
               
               > **Important:** Updates to `name` will recreate the database.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: `region`) The region in which the resource exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SdbSqlDatabaseState.__new__(_SdbSqlDatabaseState)

        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["max_cpu"] = max_cpu
        __props__.__dict__["min_cpu"] = min_cpu
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return SdbSqlDatabase(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        Endpoint of the database
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="maxCpu")
    def max_cpu(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum number of CPU units for your database. Defaults to 15.
        """
        return pulumi.get(self, "max_cpu")

    @property
    @pulumi.getter(name="minCpu")
    def min_cpu(self) -> pulumi.Output[Optional[int]]:
        """
        The minimum number of CPU units for your database. Defaults to 0.
        """
        return pulumi.get(self, "min_cpu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the database (e.g. `my-new-database`).

        > **Important:** Updates to `name` will recreate the database.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region in which the resource exists.
        """
        return pulumi.get(self, "region")

