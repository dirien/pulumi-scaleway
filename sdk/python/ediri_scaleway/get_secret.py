# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSecretResult',
    'AwaitableGetSecretResult',
    'get_secret',
    'get_secret_output',
]

@pulumi.output_type
class GetSecretResult:
    """
    A collection of values returned by getSecret.
    """
    def __init__(__self__, created_at=None, description=None, id=None, name=None, organization_id=None, project_id=None, region=None, secret_id=None, status=None, tags=None, updated_at=None, version_count=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if version_count and not isinstance(version_count, int):
            raise TypeError("Expected argument 'version_count' to be a int")
        pulumi.set(__self__, "version_count", version_count)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[str]:
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="versionCount")
    def version_count(self) -> int:
        return pulumi.get(self, "version_count")


class AwaitableGetSecretResult(GetSecretResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecretResult(
            created_at=self.created_at,
            description=self.description,
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            region=self.region,
            secret_id=self.secret_id,
            status=self.status,
            tags=self.tags,
            updated_at=self.updated_at,
            version_count=self.version_count)


def get_secret(name: Optional[str] = None,
               organization_id: Optional[str] = None,
               project_id: Optional[str] = None,
               region: Optional[str] = None,
               secret_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecretResult:
    """
    Gets information about Scaleway Secrets.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/).

    ## Examples

    ### Basic

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.Secret("main", description="barr")
    my_secret = scaleway.get_secret(secret_id="11111111-1111-1111-1111-111111111111")
    by_name = scaleway.get_secret(name="your_secret_name")
    ```


    :param str name: The secret name.
           Only one of `name` and `secret_id` should be specified.
    :param str organization_id: The organization ID the Project is associated with.
           If no default organization_id is set, one must be set explicitly in this datasource
    :param str project_id: `project_id`) The ID of the
           project the secret is associated with.
    :param str region: `region`) The region in which the secret exists.
    :param str secret_id: The secret id.
           Only one of `name` and `secret_id` should be specified.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['organizationId'] = organization_id
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['secretId'] = secret_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getSecret:getSecret', __args__, opts=opts, typ=GetSecretResult).value

    return AwaitableGetSecretResult(
        created_at=__ret__.created_at,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        organization_id=__ret__.organization_id,
        project_id=__ret__.project_id,
        region=__ret__.region,
        secret_id=__ret__.secret_id,
        status=__ret__.status,
        tags=__ret__.tags,
        updated_at=__ret__.updated_at,
        version_count=__ret__.version_count)


@_utilities.lift_output_func(get_secret)
def get_secret_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                      organization_id: Optional[pulumi.Input[Optional[str]]] = None,
                      project_id: Optional[pulumi.Input[Optional[str]]] = None,
                      region: Optional[pulumi.Input[Optional[str]]] = None,
                      secret_id: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecretResult]:
    """
    Gets information about Scaleway Secrets.
    For more information, see [the documentation](https://developers.scaleway.com/en/products/secret_manager/api/v1alpha1/).

    ## Examples

    ### Basic

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.Secret("main", description="barr")
    my_secret = scaleway.get_secret(secret_id="11111111-1111-1111-1111-111111111111")
    by_name = scaleway.get_secret(name="your_secret_name")
    ```


    :param str name: The secret name.
           Only one of `name` and `secret_id` should be specified.
    :param str organization_id: The organization ID the Project is associated with.
           If no default organization_id is set, one must be set explicitly in this datasource
    :param str project_id: `project_id`) The ID of the
           project the secret is associated with.
    :param str region: `region`) The region in which the secret exists.
    :param str secret_id: The secret id.
           Only one of `name` and `secret_id` should be specified.
    """
    ...
