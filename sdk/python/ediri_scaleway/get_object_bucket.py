# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetObjectBucketResult',
    'AwaitableGetObjectBucketResult',
    'get_object_bucket',
    'get_object_bucket_output',
]

@pulumi.output_type
class GetObjectBucketResult:
    """
    A collection of values returned by getObjectBucket.
    """
    def __init__(__self__, acl=None, api_endpoint=None, cors_rules=None, endpoint=None, force_destroy=None, id=None, lifecycle_rules=None, name=None, object_lock_enabled=None, project_id=None, region=None, tags=None, versionings=None):
        if acl and not isinstance(acl, str):
            raise TypeError("Expected argument 'acl' to be a str")
        pulumi.set(__self__, "acl", acl)
        if api_endpoint and not isinstance(api_endpoint, str):
            raise TypeError("Expected argument 'api_endpoint' to be a str")
        pulumi.set(__self__, "api_endpoint", api_endpoint)
        if cors_rules and not isinstance(cors_rules, list):
            raise TypeError("Expected argument 'cors_rules' to be a list")
        pulumi.set(__self__, "cors_rules", cors_rules)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if force_destroy and not isinstance(force_destroy, bool):
            raise TypeError("Expected argument 'force_destroy' to be a bool")
        pulumi.set(__self__, "force_destroy", force_destroy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lifecycle_rules and not isinstance(lifecycle_rules, list):
            raise TypeError("Expected argument 'lifecycle_rules' to be a list")
        pulumi.set(__self__, "lifecycle_rules", lifecycle_rules)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if object_lock_enabled and not isinstance(object_lock_enabled, bool):
            raise TypeError("Expected argument 'object_lock_enabled' to be a bool")
        pulumi.set(__self__, "object_lock_enabled", object_lock_enabled)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if versionings and not isinstance(versionings, list):
            raise TypeError("Expected argument 'versionings' to be a list")
        pulumi.set(__self__, "versionings", versionings)

    @property
    @pulumi.getter
    def acl(self) -> str:
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter(name="apiEndpoint")
    def api_endpoint(self) -> str:
        return pulumi.get(self, "api_endpoint")

    @property
    @pulumi.getter(name="corsRules")
    def cors_rules(self) -> Sequence['outputs.GetObjectBucketCorsRuleResult']:
        return pulumi.get(self, "cors_rules")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The endpoint URL of the bucket
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> bool:
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lifecycleRules")
    def lifecycle_rules(self) -> Sequence['outputs.GetObjectBucketLifecycleRuleResult']:
        return pulumi.get(self, "lifecycle_rules")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectLockEnabled")
    def object_lock_enabled(self) -> bool:
        return pulumi.get(self, "object_lock_enabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def versionings(self) -> Sequence['outputs.GetObjectBucketVersioningResult']:
        return pulumi.get(self, "versionings")


class AwaitableGetObjectBucketResult(GetObjectBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetObjectBucketResult(
            acl=self.acl,
            api_endpoint=self.api_endpoint,
            cors_rules=self.cors_rules,
            endpoint=self.endpoint,
            force_destroy=self.force_destroy,
            id=self.id,
            lifecycle_rules=self.lifecycle_rules,
            name=self.name,
            object_lock_enabled=self.object_lock_enabled,
            project_id=self.project_id,
            region=self.region,
            tags=self.tags,
            versionings=self.versionings)


def get_object_bucket(name: Optional[str] = None,
                      project_id: Optional[str] = None,
                      region: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetObjectBucketResult:
    """
    Gets information about the Bucket.
    For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).

    ## Example Usage

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.ObjectBucket("main", tags={
        "foo": "bar",
    })
    selected = scaleway.get_object_bucket_output(name=main.id)
    ```

    ### Fetching the bucket from a specific project

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    selected = scaleway.get_object_bucket(name="bucket.test.com",
        project_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str project_id: `project_id`) The ID of the project the bucket is associated with.
    :param str region: `region`) The region in which the bucket exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getObjectBucket:getObjectBucket', __args__, opts=opts, typ=GetObjectBucketResult).value

    return AwaitableGetObjectBucketResult(
        acl=pulumi.get(__ret__, 'acl'),
        api_endpoint=pulumi.get(__ret__, 'api_endpoint'),
        cors_rules=pulumi.get(__ret__, 'cors_rules'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        force_destroy=pulumi.get(__ret__, 'force_destroy'),
        id=pulumi.get(__ret__, 'id'),
        lifecycle_rules=pulumi.get(__ret__, 'lifecycle_rules'),
        name=pulumi.get(__ret__, 'name'),
        object_lock_enabled=pulumi.get(__ret__, 'object_lock_enabled'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'),
        tags=pulumi.get(__ret__, 'tags'),
        versionings=pulumi.get(__ret__, 'versionings'))


@_utilities.lift_output_func(get_object_bucket)
def get_object_bucket_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                             project_id: Optional[pulumi.Input[Optional[str]]] = None,
                             region: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetObjectBucketResult]:
    """
    Gets information about the Bucket.
    For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).

    ## Example Usage

    ```python
    import pulumi
    import ediri_scaleway as scaleway
    import pulumi_scaleway as scaleway

    main = scaleway.ObjectBucket("main", tags={
        "foo": "bar",
    })
    selected = scaleway.get_object_bucket_output(name=main.id)
    ```

    ### Fetching the bucket from a specific project

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    selected = scaleway.get_object_bucket(name="bucket.test.com",
        project_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str project_id: `project_id`) The ID of the project the bucket is associated with.
    :param str region: `region`) The region in which the bucket exists.
    """
    ...
