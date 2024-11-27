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
    'GetRegistryImageTagResult',
    'AwaitableGetRegistryImageTagResult',
    'get_registry_image_tag',
    'get_registry_image_tag_output',
]

@pulumi.output_type
class GetRegistryImageTagResult:
    """
    A collection of values returned by getRegistryImageTag.
    """
    def __init__(__self__, created_at=None, digest=None, id=None, image_id=None, name=None, organization_id=None, project_id=None, region=None, status=None, tag_id=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if digest and not isinstance(digest, str):
            raise TypeError("Expected argument 'digest' to be a str")
        pulumi.set(__self__, "digest", digest)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
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
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tag_id and not isinstance(tag_id, str):
            raise TypeError("Expected argument 'tag_id' to be a str")
        pulumi.set(__self__, "tag_id", tag_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time when the registry image tag was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def digest(self) -> str:
        """
        Hash of the tag content. Several tags of the same image may have the same digest.
        """
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The organization ID the image tag is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the registry image tag.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tagId")
    def tag_id(self) -> Optional[str]:
        return pulumi.get(self, "tag_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time of the last update to the registry image tag.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetRegistryImageTagResult(GetRegistryImageTagResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistryImageTagResult(
            created_at=self.created_at,
            digest=self.digest,
            id=self.id,
            image_id=self.image_id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            region=self.region,
            status=self.status,
            tag_id=self.tag_id,
            updated_at=self.updated_at)


def get_registry_image_tag(image_id: Optional[str] = None,
                           name: Optional[str] = None,
                           project_id: Optional[str] = None,
                           region: Optional[str] = None,
                           tag_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistryImageTagResult:
    """
    Gets information about a specific tag of a Container Registry image.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_image_tag = scaleway.get_registry_image_tag(image_id="22222222-2222-2222-2222-222222222222",
        name="my-tag-name")
    ```


    :param str image_id: The ID of the registry image.
    :param str name: The name of the registry image tag.
    :param str project_id: The ID of the project the image tag is associated with.
    :param str region: The region in which the registry image tag exists.
    :param str tag_id: The ID of the registry image tag.
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getRegistryImageTag:getRegistryImageTag', __args__, opts=opts, typ=GetRegistryImageTagResult).value

    return AwaitableGetRegistryImageTagResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        digest=pulumi.get(__ret__, 'digest'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        status=pulumi.get(__ret__, 'status'),
        tag_id=pulumi.get(__ret__, 'tag_id'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_registry_image_tag_output(image_id: Optional[pulumi.Input[str]] = None,
                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                  project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  region: Optional[pulumi.Input[Optional[str]]] = None,
                                  tag_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegistryImageTagResult]:
    """
    Gets information about a specific tag of a Container Registry image.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_image_tag = scaleway.get_registry_image_tag(image_id="22222222-2222-2222-2222-222222222222",
        name="my-tag-name")
    ```


    :param str image_id: The ID of the registry image.
    :param str name: The name of the registry image tag.
    :param str project_id: The ID of the project the image tag is associated with.
    :param str region: The region in which the registry image tag exists.
    :param str tag_id: The ID of the registry image tag.
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getRegistryImageTag:getRegistryImageTag', __args__, opts=opts, typ=GetRegistryImageTagResult)
    return __ret__.apply(lambda __response__: GetRegistryImageTagResult(
        created_at=pulumi.get(__response__, 'created_at'),
        digest=pulumi.get(__response__, 'digest'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region'),
        status=pulumi.get(__response__, 'status'),
        tag_id=pulumi.get(__response__, 'tag_id'),
        updated_at=pulumi.get(__response__, 'updated_at')))
