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
    'GetMnqSnsResult',
    'AwaitableGetMnqSnsResult',
    'get_mnq_sns',
    'get_mnq_sns_output',
]

@pulumi.output_type
class GetMnqSnsResult:
    """
    A collection of values returned by getMnqSns.
    """
    def __init__(__self__, endpoint=None, id=None, project_id=None, region=None):
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The endpoint of the SNS service for this Project.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")


class AwaitableGetMnqSnsResult(GetMnqSnsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMnqSnsResult(
            endpoint=self.endpoint,
            id=self.id,
            project_id=self.project_id,
            region=self.region)


def get_mnq_sns(project_id: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMnqSnsResult:
    """
    Gets information about SNS for a Project

    ## Examples

    ### Basic

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_mnq_sns()
    for_project = scaleway.get_mnq_sns(project_id=scaleway_account_project["main"]["id"])
    ```


    :param str project_id: `project_id`) The ID of the Project in which sns is enabled.
    :param str region: `region`). The region in which sns is enabled.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getMnqSns:getMnqSns', __args__, opts=opts, typ=GetMnqSnsResult).value

    return AwaitableGetMnqSnsResult(
        endpoint=pulumi.get(__ret__, 'endpoint'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        region=pulumi.get(__ret__, 'region'))
def get_mnq_sns_output(project_id: Optional[pulumi.Input[Optional[str]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMnqSnsResult]:
    """
    Gets information about SNS for a Project

    ## Examples

    ### Basic

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    main = scaleway.get_mnq_sns()
    for_project = scaleway.get_mnq_sns(project_id=scaleway_account_project["main"]["id"])
    ```


    :param str project_id: `project_id`) The ID of the Project in which sns is enabled.
    :param str region: `region`). The region in which sns is enabled.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getMnqSns:getMnqSns', __args__, opts=opts, typ=GetMnqSnsResult)
    return __ret__.apply(lambda __response__: GetMnqSnsResult(
        endpoint=pulumi.get(__response__, 'endpoint'),
        id=pulumi.get(__response__, 'id'),
        project_id=pulumi.get(__response__, 'project_id'),
        region=pulumi.get(__response__, 'region')))
