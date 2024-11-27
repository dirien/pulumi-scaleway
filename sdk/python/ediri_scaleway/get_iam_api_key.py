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
    'GetIamApiKeyResult',
    'AwaitableGetIamApiKeyResult',
    'get_iam_api_key',
    'get_iam_api_key_output',
]

@pulumi.output_type
class GetIamApiKeyResult:
    """
    A collection of values returned by getIamApiKey.
    """
    def __init__(__self__, access_key=None, application_id=None, created_at=None, creation_ip=None, default_project_id=None, description=None, editable=None, expires_at=None, id=None, updated_at=None, user_id=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if creation_ip and not isinstance(creation_ip, str):
            raise TypeError("Expected argument 'creation_ip' to be a str")
        pulumi.set(__self__, "creation_ip", creation_ip)
        if default_project_id and not isinstance(default_project_id, str):
            raise TypeError("Expected argument 'default_project_id' to be a str")
        pulumi.set(__self__, "default_project_id", default_project_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if editable and not isinstance(editable, bool):
            raise TypeError("Expected argument 'editable' to be a bool")
        pulumi.set(__self__, "editable", editable)
        if expires_at and not isinstance(expires_at, str):
            raise TypeError("Expected argument 'expires_at' to be a str")
        pulumi.set(__self__, "expires_at", expires_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> str:
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="creationIp")
    def creation_ip(self) -> str:
        return pulumi.get(self, "creation_ip")

    @property
    @pulumi.getter(name="defaultProjectId")
    def default_project_id(self) -> str:
        return pulumi.get(self, "default_project_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def editable(self) -> bool:
        return pulumi.get(self, "editable")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> str:
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")


class AwaitableGetIamApiKeyResult(GetIamApiKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamApiKeyResult(
            access_key=self.access_key,
            application_id=self.application_id,
            created_at=self.created_at,
            creation_ip=self.creation_ip,
            default_project_id=self.default_project_id,
            description=self.description,
            editable=self.editable,
            expires_at=self.expires_at,
            id=self.id,
            updated_at=self.updated_at,
            user_id=self.user_id)


def get_iam_api_key(access_key: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamApiKeyResult:
    """
    Gets information about an existing IAM API key. For more information, refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#api-keys-3665ae).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_iam_api_key(access_key="SCWABCDEFGHIJKLMNOPQ")
    ```


    :param str access_key: The access key of the IAM API key which is also the ID of the API key.
    """
    __args__ = dict()
    __args__['accessKey'] = access_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getIamApiKey:getIamApiKey', __args__, opts=opts, typ=GetIamApiKeyResult).value

    return AwaitableGetIamApiKeyResult(
        access_key=pulumi.get(__ret__, 'access_key'),
        application_id=pulumi.get(__ret__, 'application_id'),
        created_at=pulumi.get(__ret__, 'created_at'),
        creation_ip=pulumi.get(__ret__, 'creation_ip'),
        default_project_id=pulumi.get(__ret__, 'default_project_id'),
        description=pulumi.get(__ret__, 'description'),
        editable=pulumi.get(__ret__, 'editable'),
        expires_at=pulumi.get(__ret__, 'expires_at'),
        id=pulumi.get(__ret__, 'id'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_iam_api_key_output(access_key: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIamApiKeyResult]:
    """
    Gets information about an existing IAM API key. For more information, refer to the [IAM API documentation](https://www.scaleway.com/en/developers/api/iam/#api-keys-3665ae).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_iam_api_key(access_key="SCWABCDEFGHIJKLMNOPQ")
    ```


    :param str access_key: The access key of the IAM API key which is also the ID of the API key.
    """
    __args__ = dict()
    __args__['accessKey'] = access_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getIamApiKey:getIamApiKey', __args__, opts=opts, typ=GetIamApiKeyResult)
    return __ret__.apply(lambda __response__: GetIamApiKeyResult(
        access_key=pulumi.get(__response__, 'access_key'),
        application_id=pulumi.get(__response__, 'application_id'),
        created_at=pulumi.get(__response__, 'created_at'),
        creation_ip=pulumi.get(__response__, 'creation_ip'),
        default_project_id=pulumi.get(__response__, 'default_project_id'),
        description=pulumi.get(__response__, 'description'),
        editable=pulumi.get(__response__, 'editable'),
        expires_at=pulumi.get(__response__, 'expires_at'),
        id=pulumi.get(__response__, 'id'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        user_id=pulumi.get(__response__, 'user_id')))
