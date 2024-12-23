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
    'GetObjectBucketPolicyResult',
    'AwaitableGetObjectBucketPolicyResult',
    'get_object_bucket_policy',
    'get_object_bucket_policy_output',
]

@pulumi.output_type
class GetObjectBucketPolicyResult:
    """
    A collection of values returned by getObjectBucketPolicy.
    """
    def __init__(__self__, bucket=None, id=None, policy=None, project_id=None, region=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy and not isinstance(policy, str):
            raise TypeError("Expected argument 'policy' to be a str")
        pulumi.set(__self__, "policy", policy)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def bucket(self) -> str:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def policy(self) -> str:
        """
        The content of the bucket policy in JSON format.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")


class AwaitableGetObjectBucketPolicyResult(GetObjectBucketPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetObjectBucketPolicyResult(
            bucket=self.bucket,
            id=self.id,
            policy=self.policy,
            project_id=self.project_id,
            region=self.region)


def get_object_bucket_policy(bucket: Optional[str] = None,
                             project_id: Optional[str] = None,
                             region: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetObjectBucketPolicyResult:
    """
    The `ObjectBucketPolicy` data source is used to retrieve information about the bucket policy of an Object Storage bucket.

    Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/) for more information.

    ## Retrieve the bucket policy of a bucket

    The following command allows you to retrieve a bucket policy by its bucket.

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_object_bucket_policy(bucket="bucket.test.com")
    ```


    :param str bucket: The name of the bucket.
    :param str region: `region`) The region in which the Object Storage exists.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getObjectBucketPolicy:getObjectBucketPolicy', __args__, opts=opts, typ=GetObjectBucketPolicyResult).value

    return AwaitableGetObjectBucketPolicyResult(
        bucket=pulumi.get(__ret__, 'bucket'),
        id=pulumi.get(__ret__, 'id'),
        policy=pulumi.get(__ret__, 'policy'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'))
def get_object_bucket_policy_output(bucket: Optional[pulumi.Input[str]] = None,
                                    project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    region: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetObjectBucketPolicyResult]:
    """
    The `ObjectBucketPolicy` data source is used to retrieve information about the bucket policy of an Object Storage bucket.

    Refer to the Object Storage [documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/) for more information.

    ## Retrieve the bucket policy of a bucket

    The following command allows you to retrieve a bucket policy by its bucket.

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_object_bucket_policy(bucket="bucket.test.com")
    ```


    :param str bucket: The name of the bucket.
    :param str region: `region`) The region in which the Object Storage exists.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getObjectBucketPolicy:getObjectBucketPolicy', __args__, opts=opts, typ=GetObjectBucketPolicyResult)
    return __ret__.apply(lambda __response__: GetObjectBucketPolicyResult(
        bucket=pulumi.get(__response__, 'bucket'),
        id=pulumi.get(__response__, 'id'),
        policy=pulumi.get(__response__, 'policy'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region')))
