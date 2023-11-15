# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MnqNatsAccountArgs', 'MnqNatsAccount']

@pulumi.input_type
class MnqNatsAccountArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MnqNatsAccount resource.
        :param pulumi.Input[str] name: The unique name of the nats account.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               account is associated with.
        :param pulumi.Input[str] region: `region`). The region
               in which the account should be created.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the nats account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the
        account is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which the account should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MnqNatsAccountState:
    def __init__(__self__, *,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MnqNatsAccount resources.
        :param pulumi.Input[str] endpoint: The endpoint of the NATS service for this account.
        :param pulumi.Input[str] name: The unique name of the nats account.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               account is associated with.
        :param pulumi.Input[str] region: `region`). The region
               in which the account should be created.
        """
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
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
        The endpoint of the NATS service for this account.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the nats account.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the
        account is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which the account should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class MnqNatsAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Messaging and queuing Nats Accounts.
        For further information please check
        our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.MnqNatsAccount("main")
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/mnqNatsAccount:MnqNatsAccount main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The unique name of the nats account.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               account is associated with.
        :param pulumi.Input[str] region: `region`). The region
               in which the account should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MnqNatsAccountArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Messaging and queuing Nats Accounts.
        For further information please check
        our [documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/nats-overview/)

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.MnqNatsAccount("main")
        ```

        ## Import

        Namespaces can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/mnqNatsAccount:MnqNatsAccount main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param MnqNatsAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqNatsAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
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
            __props__ = MnqNatsAccountArgs.__new__(MnqNatsAccountArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["endpoint"] = None
        super(MnqNatsAccount, __self__).__init__(
            'scaleway:index/mnqNatsAccount:MnqNatsAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'MnqNatsAccount':
        """
        Get an existing MnqNatsAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint: The endpoint of the NATS service for this account.
        :param pulumi.Input[str] name: The unique name of the nats account.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the
               account is associated with.
        :param pulumi.Input[str] region: `region`). The region
               in which the account should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqNatsAccountState.__new__(_MnqNatsAccountState)

        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return MnqNatsAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint of the NATS service for this account.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the nats account.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the
        account is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region
        in which the account should be created.
        """
        return pulumi.get(self, "region")

