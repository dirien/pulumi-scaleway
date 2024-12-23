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

__all__ = [
    'GetBillingConsumptionsResult',
    'AwaitableGetBillingConsumptionsResult',
    'get_billing_consumptions',
    'get_billing_consumptions_output',
]

@pulumi.output_type
class GetBillingConsumptionsResult:
    """
    A collection of values returned by getBillingConsumptions.
    """
    def __init__(__self__, consumptions=None, id=None, organization_id=None, project_id=None, updated_at=None):
        if consumptions and not isinstance(consumptions, list):
            raise TypeError("Expected argument 'consumptions' to be a list")
        pulumi.set(__self__, "consumptions", consumptions)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def consumptions(self) -> Sequence['outputs.GetBillingConsumptionsConsumptionResult']:
        return pulumi.get(self, "consumptions")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")


class AwaitableGetBillingConsumptionsResult(GetBillingConsumptionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBillingConsumptionsResult(
            consumptions=self.consumptions,
            id=self.id,
            organization_id=self.organization_id,
            project_id=self.project_id,
            updated_at=self.updated_at)


def get_billing_consumptions(project_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBillingConsumptionsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getBillingConsumptions:getBillingConsumptions', __args__, opts=opts, typ=GetBillingConsumptionsResult).value

    return AwaitableGetBillingConsumptionsResult(
        consumptions=pulumi.get(__ret__, 'consumptions'),
        id=pulumi.get(__ret__, 'id'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_billing_consumptions_output(project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBillingConsumptionsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getBillingConsumptions:getBillingConsumptions', __args__, opts=opts, typ=GetBillingConsumptionsResult)
    return __ret__.apply(lambda __response__: GetBillingConsumptionsResult(
        consumptions=pulumi.get(__response__, 'consumptions'),
        id=pulumi.get(__response__, 'id'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        updated_at=pulumi.get(__response__, 'updated_at')))
