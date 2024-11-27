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

__all__ = [
    'GetIamUserResult',
    'AwaitableGetIamUserResult',
    'get_iam_user',
    'get_iam_user_output',
]

@pulumi.output_type
class GetIamUserResult:
    """
    A collection of values returned by getIamUser.
    """
    def __init__(__self__, email=None, id=None, organization_id=None, user_id=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[str]:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[str]:
        return pulumi.get(self, "user_id")


class AwaitableGetIamUserResult(GetIamUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamUserResult(
            email=self.email,
            id=self.id,
            organization_id=self.organization_id,
            user_id=self.user_id)


def get_iam_user(email: Optional[str] = None,
                 organization_id: Optional[str] = None,
                 user_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamUserResult:
    """
    Use this data source to get information on an existing IAM user based on its ID or email address.
    For more information refer to the [IAM API documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#users-06bdcf).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    find_by_id = scaleway.get_iam_user(user_id="11111111-1111-1111-1111-111111111111")
    find_by_email = scaleway.get_iam_user(email="foo@bar.com")
    ```


    :param str email: The email address of the IAM user.
    :param str organization_id: `organization_id`) The ID of the
           organization the user is associated with.
    :param str user_id: The ID of the IAM user.
           
           > **Note** You must specify at least one: `name` and/or `user_id`.
    """
    __args__ = dict()
    __args__['email'] = email
    __args__['organizationId'] = organization_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getIamUser:getIamUser', __args__, opts=opts, typ=GetIamUserResult).value

    return AwaitableGetIamUserResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_iam_user_output(email: Optional[pulumi.Input[Optional[str]]] = None,
                        organization_id: Optional[pulumi.Input[Optional[str]]] = None,
                        user_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIamUserResult]:
    """
    Use this data source to get information on an existing IAM user based on its ID or email address.
    For more information refer to the [IAM API documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#users-06bdcf).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    find_by_id = scaleway.get_iam_user(user_id="11111111-1111-1111-1111-111111111111")
    find_by_email = scaleway.get_iam_user(email="foo@bar.com")
    ```


    :param str email: The email address of the IAM user.
    :param str organization_id: `organization_id`) The ID of the
           organization the user is associated with.
    :param str user_id: The ID of the IAM user.
           
           > **Note** You must specify at least one: `name` and/or `user_id`.
    """
    __args__ = dict()
    __args__['email'] = email
    __args__['organizationId'] = organization_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getIamUser:getIamUser', __args__, opts=opts, typ=GetIamUserResult)
    return __ret__.apply(lambda __response__: GetIamUserResult(
        email=pulumi.get(__response__, 'email'),
        id=pulumi.get(__response__, 'id'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        user_id=pulumi.get(__response__, 'user_id')))
