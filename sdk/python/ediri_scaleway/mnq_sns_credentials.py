# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['MnqSnsCredentialsArgs', 'MnqSnsCredentials']

@pulumi.input_type
class MnqSnsCredentialsArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input['MnqSnsCredentialsPermissionsArgs']] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MnqSnsCredentials resource.
        :param pulumi.Input[str] name: The unique name of the SNS credentials.
        :param pulumi.Input['MnqSnsCredentialsPermissionsArgs'] permissions: . List of permissions associated with these credentials. Only one of the following permissions may be set:
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region in which SNS is enabled.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the SNS credentials.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input['MnqSnsCredentialsPermissionsArgs']]:
        """
        . List of permissions associated with these credentials. Only one of the following permissions may be set:
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input['MnqSnsCredentialsPermissionsArgs']]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _MnqSnsCredentialsState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input['MnqSnsCredentialsPermissionsArgs']] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MnqSnsCredentials resources.
        :param pulumi.Input[str] access_key: The ID of the key.
        :param pulumi.Input[str] name: The unique name of the SNS credentials.
        :param pulumi.Input['MnqSnsCredentialsPermissionsArgs'] permissions: . List of permissions associated with these credentials. Only one of the following permissions may be set:
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret value of the key.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the key.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the SNS credentials.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input['MnqSnsCredentialsPermissionsArgs']]:
        """
        . List of permissions associated with these credentials. Only one of the following permissions may be set:
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input['MnqSnsCredentialsPermissionsArgs']]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret value of the key.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)


class MnqSnsCredentials(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Union['MnqSnsCredentialsPermissionsArgs', 'MnqSnsCredentialsPermissionsArgsDict']]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Messaging and Queuing SNS credentials.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_mnq_sns = scaleway.MnqSns("mainMnqSns")
        main_mnq_sns_credentials = scaleway.MnqSnsCredentials("mainMnqSnsCredentials",
            project_id=main_mnq_sns.project_id,
            permissions={
                "can_manage": False,
                "can_receive": True,
                "can_publish": False,
            })
        ```

        ## Import

        SNS credentials can be imported using `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqSnsCredentials:MnqSnsCredentials main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The unique name of the SNS credentials.
        :param pulumi.Input[Union['MnqSnsCredentialsPermissionsArgs', 'MnqSnsCredentialsPermissionsArgsDict']] permissions: . List of permissions associated with these credentials. Only one of the following permissions may be set:
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region in which SNS is enabled.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MnqSnsCredentialsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Messaging and Queuing SNS credentials.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/)

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_mnq_sns = scaleway.MnqSns("mainMnqSns")
        main_mnq_sns_credentials = scaleway.MnqSnsCredentials("mainMnqSnsCredentials",
            project_id=main_mnq_sns.project_id,
            permissions={
                "can_manage": False,
                "can_receive": True,
                "can_publish": False,
            })
        ```

        ## Import

        SNS credentials can be imported using `{region}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqSnsCredentials:MnqSnsCredentials main fr-par/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param MnqSnsCredentialsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqSnsCredentialsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Union['MnqSnsCredentialsPermissionsArgs', 'MnqSnsCredentialsPermissionsArgsDict']]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqSnsCredentialsArgs.__new__(MnqSnsCredentialsArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["access_key"] = None
            __props__.__dict__["secret_key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MnqSnsCredentials, __self__).__init__(
            'scaleway:index/mnqSnsCredentials:MnqSnsCredentials',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Union['MnqSnsCredentialsPermissionsArgs', 'MnqSnsCredentialsPermissionsArgsDict']]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None) -> 'MnqSnsCredentials':
        """
        Get an existing MnqSnsCredentials resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The ID of the key.
        :param pulumi.Input[str] name: The unique name of the SNS credentials.
        :param pulumi.Input[Union['MnqSnsCredentialsPermissionsArgs', 'MnqSnsCredentialsPermissionsArgsDict']] permissions: . List of permissions associated with these credentials. Only one of the following permissions may be set:
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] region: `region`). The region in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret value of the key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqSnsCredentialsState.__new__(_MnqSnsCredentialsState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["name"] = name
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_key"] = secret_key
        return MnqSnsCredentials(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The ID of the key.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the SNS credentials.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output['outputs.MnqSnsCredentialsPermissions']:
        """
        . List of permissions associated with these credentials. Only one of the following permissions may be set:
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The secret value of the key.
        """
        return pulumi.get(self, "secret_key")

