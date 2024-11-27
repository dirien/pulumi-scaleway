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

__all__ = ['IamPolicyArgs', 'IamPolicy']

@pulumi.input_type
class IamPolicyArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]],
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_principal: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamPolicy resource.
        :param pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]] rules: List of rules in the policy.
        :param pulumi.Input[str] application_id: ID of the application the policy will be linked to
        :param pulumi.Input[str] description: The description of the IAM policy.
        :param pulumi.Input[str] group_id: ID of the group the policy will be linked to
        :param pulumi.Input[str] name: The name of the IAM policy.
        :param pulumi.Input[bool] no_principal: If the policy doesn't apply to a principal.
               
               > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the policy is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IAM policy.
        :param pulumi.Input[str] user_id: ID of the user the policy will be linked to
        """
        pulumi.set(__self__, "rules", rules)
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_principal is not None:
            pulumi.set(__self__, "no_principal", no_principal)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]]:
        """
        List of rules in the policy.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the application the policy will be linked to
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IAM policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the group the policy will be linked to
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IAM policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noPrincipal")
    def no_principal(self) -> Optional[pulumi.Input[bool]]:
        """
        If the policy doesn't apply to a principal.

        > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        """
        return pulumi.get(self, "no_principal")

    @no_principal.setter
    def no_principal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_principal", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        `organization_id`) The ID of the organization the policy is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the IAM policy.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the user the policy will be linked to
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _IamPolicyState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 editable: Optional[pulumi.Input[bool]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_principal: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamPolicy resources.
        :param pulumi.Input[str] application_id: ID of the application the policy will be linked to
        :param pulumi.Input[str] created_at: The date and time of the creation of the policy.
        :param pulumi.Input[str] description: The description of the IAM policy.
        :param pulumi.Input[bool] editable: Whether the policy is editable.
        :param pulumi.Input[str] group_id: ID of the group the policy will be linked to
        :param pulumi.Input[str] name: The name of the IAM policy.
        :param pulumi.Input[bool] no_principal: If the policy doesn't apply to a principal.
               
               > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the policy is associated with.
        :param pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]] rules: List of rules in the policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IAM policy.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the policy.
        :param pulumi.Input[str] user_id: ID of the user the policy will be linked to
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if editable is not None:
            pulumi.set(__self__, "editable", editable)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if no_principal is not None:
            pulumi.set(__self__, "no_principal", no_principal)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the application the policy will be linked to
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the policy.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the IAM policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def editable(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the policy is editable.
        """
        return pulumi.get(self, "editable")

    @editable.setter
    def editable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "editable", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the group the policy will be linked to
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IAM policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noPrincipal")
    def no_principal(self) -> Optional[pulumi.Input[bool]]:
        """
        If the policy doesn't apply to a principal.

        > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        """
        return pulumi.get(self, "no_principal")

    @no_principal.setter
    def no_principal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_principal", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        `organization_id`) The ID of the organization the policy is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]]]:
        """
        List of rules in the policy.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IamPolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the IAM policy.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the policy.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the user the policy will be linked to
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class IamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_principal: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IamPolicyRuleArgs', 'IamPolicyRuleArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway IAM Policies. For more information refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#path-policies-create-a-new-policy).

        > You can find a detailed list of all permission sets available at Scaleway in the permission sets [reference page](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/).

        ## Example Usage

        ### Create a policy for an organization's project

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        default = scaleway.get_account_project(name="default")
        app = scaleway.IamApplication("app")
        object_read_only = scaleway.IamPolicy("objectReadOnly",
            description="gives app readonly access to object storage in project",
            application_id=app.id,
            rules=[{
                "project_ids": [default.id],
                "permission_set_names": ["ObjectStorageReadOnly"],
            }])
        ```

        ### Create a policy for all current and future projects in an organization

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        app = scaleway.IamApplication("app")
        object_read_only = scaleway.IamPolicy("objectReadOnly",
            description="gives app readonly access to object storage in project",
            application_id=app.id,
            rules=[{
                "organization_id": app.organization_id,
                "permission_set_names": ["ObjectStorageReadOnly"],
            }])
        ```

        ### Create a policy with a particular condition

        IAM policy rule can use a condition to be applied.
        The following variables are available:

        - `request.ip`
        - `request.user_agent`
        - `request.time`

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.IamPolicy("main",
            no_principal=True,
            rules=[{
                "condition": "request.user_agent == 'My User Agent'",
                "organization_id": "%s",
                "permission_set_names": ["AllProductsFullAccess"],
            }])
        ```

        ## Import

        Policies can be imported using the `{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iamPolicy:IamPolicy main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: ID of the application the policy will be linked to
        :param pulumi.Input[str] description: The description of the IAM policy.
        :param pulumi.Input[str] group_id: ID of the group the policy will be linked to
        :param pulumi.Input[str] name: The name of the IAM policy.
        :param pulumi.Input[bool] no_principal: If the policy doesn't apply to a principal.
               
               > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the policy is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[Union['IamPolicyRuleArgs', 'IamPolicyRuleArgsDict']]]] rules: List of rules in the policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IAM policy.
        :param pulumi.Input[str] user_id: ID of the user the policy will be linked to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway IAM Policies. For more information refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#path-policies-create-a-new-policy).

        > You can find a detailed list of all permission sets available at Scaleway in the permission sets [reference page](https://www.scaleway.com/en/docs/identity-and-access-management/iam/reference-content/permission-sets/).

        ## Example Usage

        ### Create a policy for an organization's project

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        default = scaleway.get_account_project(name="default")
        app = scaleway.IamApplication("app")
        object_read_only = scaleway.IamPolicy("objectReadOnly",
            description="gives app readonly access to object storage in project",
            application_id=app.id,
            rules=[{
                "project_ids": [default.id],
                "permission_set_names": ["ObjectStorageReadOnly"],
            }])
        ```

        ### Create a policy for all current and future projects in an organization

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        app = scaleway.IamApplication("app")
        object_read_only = scaleway.IamPolicy("objectReadOnly",
            description="gives app readonly access to object storage in project",
            application_id=app.id,
            rules=[{
                "organization_id": app.organization_id,
                "permission_set_names": ["ObjectStorageReadOnly"],
            }])
        ```

        ### Create a policy with a particular condition

        IAM policy rule can use a condition to be applied.
        The following variables are available:

        - `request.ip`
        - `request.user_agent`
        - `request.time`

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.IamPolicy("main",
            no_principal=True,
            rules=[{
                "condition": "request.user_agent == 'My User Agent'",
                "organization_id": "%s",
                "permission_set_names": ["AllProductsFullAccess"],
            }])
        ```

        ## Import

        Policies can be imported using the `{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/iamPolicy:IamPolicy main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param IamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_principal: Optional[pulumi.Input[bool]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IamPolicyRuleArgs', 'IamPolicyRuleArgsDict']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamPolicyArgs.__new__(IamPolicyArgs)

            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["description"] = description
            __props__.__dict__["group_id"] = group_id
            __props__.__dict__["name"] = name
            __props__.__dict__["no_principal"] = no_principal
            __props__.__dict__["organization_id"] = organization_id
            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            __props__.__dict__["tags"] = tags
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["created_at"] = None
            __props__.__dict__["editable"] = None
            __props__.__dict__["updated_at"] = None
        super(IamPolicy, __self__).__init__(
            'scaleway:index/iamPolicy:IamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            editable: Optional[pulumi.Input[bool]] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            no_principal: Optional[pulumi.Input[bool]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IamPolicyRuleArgs', 'IamPolicyRuleArgsDict']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'IamPolicy':
        """
        Get an existing IamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: ID of the application the policy will be linked to
        :param pulumi.Input[str] created_at: The date and time of the creation of the policy.
        :param pulumi.Input[str] description: The description of the IAM policy.
        :param pulumi.Input[bool] editable: Whether the policy is editable.
        :param pulumi.Input[str] group_id: ID of the group the policy will be linked to
        :param pulumi.Input[str] name: The name of the IAM policy.
        :param pulumi.Input[bool] no_principal: If the policy doesn't apply to a principal.
               
               > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        :param pulumi.Input[str] organization_id: `organization_id`) The ID of the organization the policy is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[Union['IamPolicyRuleArgs', 'IamPolicyRuleArgsDict']]]] rules: List of rules in the policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the IAM policy.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the policy.
        :param pulumi.Input[str] user_id: ID of the user the policy will be linked to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamPolicyState.__new__(_IamPolicyState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["editable"] = editable
        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["name"] = name
        __props__.__dict__["no_principal"] = no_principal
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["rules"] = rules
        __props__.__dict__["tags"] = tags
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["user_id"] = user_id
        return IamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the application the policy will be linked to
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the policy.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the IAM policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def editable(self) -> pulumi.Output[bool]:
        """
        Whether the policy is editable.
        """
        return pulumi.get(self, "editable")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the group the policy will be linked to
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the IAM policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noPrincipal")
    def no_principal(self) -> pulumi.Output[Optional[bool]]:
        """
        If the policy doesn't apply to a principal.

        > **Important** Only one of `user_id`, `group_id`, `application_id` and `no_principal` may be set.
        """
        return pulumi.get(self, "no_principal")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        `organization_id`) The ID of the organization the policy is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.IamPolicyRule']]:
        """
        List of rules in the policy.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the IAM policy.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the policy.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the user the policy will be linked to
        """
        return pulumi.get(self, "user_id")

