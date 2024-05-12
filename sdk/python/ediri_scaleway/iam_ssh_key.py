# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamSshKeyArgs', 'IamSshKey']

@pulumi.input_type
class IamSshKeyArgs:
    def __init__(__self__, *,
                 public_key: pulumi.Input[str],
                 disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamSshKey resource.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        :param pulumi.Input[bool] disabled: The SSH key status.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is
               associated with.
        """
        pulumi.set(__self__, "public_key", public_key)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Input[str]:
        """
        The public SSH key to be added.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The SSH key status.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the SSH key is
        associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)


@pulumi.input_type
class _IamSshKeyState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamSshKey resources.
        :param pulumi.Input[str] created_at: The date and time of the creation of the SSH key.
        :param pulumi.Input[bool] disabled: The SSH key status.
        :param pulumi.Input[str] fingerprint: The fingerprint of the iam SSH key.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] organization_id: The ID of the organization the SSH key is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is
               associated with.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the SSH key.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the SSH key.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The SSH key status.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The fingerprint of the iam SSH key.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the organization the SSH key is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the SSH key is
        associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public SSH key to be added.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the SSH key.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class IamSshKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway IAM SSH Keys.
        For more information,
        see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#ssh-keys-d8ccd4).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.IamSshKey("main", public_key="<YOUR-PUBLIC-SSH-KEY>")
        ```

        ## Import

        SSH keys can be imported using the `id`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iamSshKey:IamSshKey main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: The SSH key status.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is
               associated with.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamSshKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway IAM SSH Keys.
        For more information,
        see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#ssh-keys-d8ccd4).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.IamSshKey("main", public_key="<YOUR-PUBLIC-SSH-KEY>")
        ```

        ## Import

        SSH keys can be imported using the `id`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iamSshKey:IamSshKey main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IamSshKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamSshKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamSshKeyArgs.__new__(IamSshKeyArgs)

            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            if public_key is None and not opts.urn:
                raise TypeError("Missing required property 'public_key'")
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["created_at"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["updated_at"] = None
        super(IamSshKey, __self__).__init__(
            'scaleway:index/iamSshKey:IamSshKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'IamSshKey':
        """
        Get an existing IamSshKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The date and time of the creation of the SSH key.
        :param pulumi.Input[bool] disabled: The SSH key status.
        :param pulumi.Input[str] fingerprint: The fingerprint of the iam SSH key.
        :param pulumi.Input[str] name: The name of the SSH key.
        :param pulumi.Input[str] organization_id: The ID of the organization the SSH key is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the SSH key is
               associated with.
        :param pulumi.Input[str] public_key: The public SSH key to be added.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the SSH key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamSshKeyState.__new__(_IamSshKeyState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["updated_at"] = updated_at
        return IamSshKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the SSH key.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        The SSH key status.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint of the iam SSH key.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the SSH key.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The ID of the organization the SSH key is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the SSH key is
        associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        The public SSH key to be added.
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the SSH key.
        """
        return pulumi.get(self, "updated_at")

