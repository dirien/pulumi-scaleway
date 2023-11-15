# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IamGroupArgs', 'IamGroup']

@pulumi.input_type
class IamGroupArgs:
    def __init__(__self__, *,
                 application_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_membership: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a IamGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] application_ids: The list of IDs of the applications attached to the group.
        :param pulumi.Input[str] description: The description of the IAM group.
        :param pulumi.Input[bool] external_membership: Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        :param pulumi.Input[str] name: The name of the IAM group.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the group is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: The list of IDs of the users attached to the group.
        """
        if application_ids is not None:
            pulumi.set(__self__, "application_ids", application_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_membership is not None:
            pulumi.set(__self__, "external_membership", external_membership)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)

    @property
    @pulumi.getter(name="applicationIds")
    def application_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of IDs of the applications attached to the group.
        """
        return pulumi.get(self, "application_ids")

    @application_ids.setter
    def application_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "application_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IAM group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalMembership")
    def external_membership(self) -> Optional[pulumi.Input[bool]]:
        """
        Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        """
        return pulumi.get(self, "external_membership")

    @external_membership.setter
    def external_membership(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_membership", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IAM group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        `organization_id`) The ID of the organization the group is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of IDs of the users attached to the group.
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_ids", value)


@pulumi.input_type
class _IamGroupState:
    def __init__(__self__, *,
                 application_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_membership: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering IamGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] application_ids: The list of IDs of the applications attached to the group.
        :param pulumi.Input[str] created_at: The date and time of the creation of the group
        :param pulumi.Input[str] description: The description of the IAM group.
        :param pulumi.Input[bool] external_membership: Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        :param pulumi.Input[str] name: The name of the IAM group.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the group is associated with.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: The list of IDs of the users attached to the group.
        """
        if application_ids is not None:
            pulumi.set(__self__, "application_ids", application_ids)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_membership is not None:
            pulumi.set(__self__, "external_membership", external_membership)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if user_ids is not None:
            pulumi.set(__self__, "user_ids", user_ids)

    @property
    @pulumi.getter(name="applicationIds")
    def application_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of IDs of the applications attached to the group.
        """
        return pulumi.get(self, "application_ids")

    @application_ids.setter
    def application_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "application_ids", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the group
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IAM group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalMembership")
    def external_membership(self) -> Optional[pulumi.Input[bool]]:
        """
        Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        """
        return pulumi.get(self, "external_membership")

    @external_membership.setter
    def external_membership(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_membership", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IAM group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        `organization_id`) The ID of the organization the group is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the group
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of IDs of the users attached to the group.
        """
        return pulumi.get(self, "user_ids")

    @user_ids.setter
    def user_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "user_ids", value)


class IamGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_membership: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway IAM Groups.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        basic = scaleway.IamGroup("basic",
            application_ids=[],
            description="basic description",
            user_ids=[])
        ```

        ### With applications

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        app = scaleway.IamApplication("app")
        with_app = scaleway.IamGroup("withApp",
            application_ids=[app.id],
            user_ids=[])
        ```

        ## Import

        IAM groups can be imported using the `{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/iamGroup:IamGroup basic 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] application_ids: The list of IDs of the applications attached to the group.
        :param pulumi.Input[str] description: The description of the IAM group.
        :param pulumi.Input[bool] external_membership: Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        :param pulumi.Input[str] name: The name of the IAM group.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the group is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: The list of IDs of the users attached to the group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[IamGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway IAM Groups.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        basic = scaleway.IamGroup("basic",
            application_ids=[],
            description="basic description",
            user_ids=[])
        ```

        ### With applications

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        app = scaleway.IamApplication("app")
        with_app = scaleway.IamGroup("withApp",
            application_ids=[app.id],
            user_ids=[])
        ```

        ## Import

        IAM groups can be imported using the `{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/iamGroup:IamGroup basic 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IamGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_membership: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamGroupArgs.__new__(IamGroupArgs)

            __props__.__dict__["application_ids"] = application_ids
            __props__.__dict__["description"] = description
            __props__.__dict__["external_membership"] = external_membership
            __props__.__dict__["name"] = name
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["user_ids"] = user_ids
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(IamGroup, __self__).__init__(
            'scaleway:index/iamGroup:IamGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_membership: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            user_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'IamGroup':
        """
        Get an existing IamGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] application_ids: The list of IDs of the applications attached to the group.
        :param pulumi.Input[str] created_at: The date and time of the creation of the group
        :param pulumi.Input[str] description: The description of the IAM group.
        :param pulumi.Input[bool] external_membership: Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        :param pulumi.Input[str] name: The name of the IAM group.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the group is associated with.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] user_ids: The list of IDs of the users attached to the group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamGroupState.__new__(_IamGroupState)

        __props__.__dict__["application_ids"] = application_ids
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["external_membership"] = external_membership
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["user_ids"] = user_ids
        return IamGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationIds")
    def application_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of IDs of the applications attached to the group.
        """
        return pulumi.get(self, "application_ids")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the group
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the IAM group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalMembership")
    def external_membership(self) -> pulumi.Output[Optional[bool]]:
        """
        Manage membership externally. This make the resource ignore user_ids and application_ids. Should be used when using iam_group_membership
        """
        return pulumi.get(self, "external_membership")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the IAM group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        `organization_id`) The ID of the organization the group is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the group
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userIds")
    def user_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of IDs of the users attached to the group.
        """
        return pulumi.get(self, "user_ids")

