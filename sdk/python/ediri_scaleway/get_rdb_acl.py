# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetRdbAclResult',
    'AwaitableGetRdbAclResult',
    'get_rdb_acl',
    'get_rdb_acl_output',
]

@pulumi.output_type
class GetRdbAclResult:
    """
    A collection of values returned by getRdbAcl.
    """
    def __init__(__self__, acl_rules=None, id=None, instance_id=None, region=None):
        if acl_rules and not isinstance(acl_rules, list):
            raise TypeError("Expected argument 'acl_rules' to be a list")
        pulumi.set(__self__, "acl_rules", acl_rules)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> Sequence['outputs.GetRdbAclAclRuleResult']:
        """
        A list of ACLs rules (structure is described below)
        """
        return pulumi.get(self, "acl_rules")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")


class AwaitableGetRdbAclResult(GetRdbAclResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRdbAclResult(
            acl_rules=self.acl_rules,
            id=self.id,
            instance_id=self.instance_id,
            region=self.region)


def get_rdb_acl(instance_id: Optional[str] = None,
                region: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRdbAclResult:
    """
    Gets information about the RDB instance network Access Control List.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_acl = scaleway.get_rdb_acl(instance_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str instance_id: The RDB instance ID.
    :param str region: `region`) The region in which the Database Instance should be created.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getRdbAcl:getRdbAcl', __args__, opts=opts, typ=GetRdbAclResult).value

    return AwaitableGetRdbAclResult(
        acl_rules=pulumi.get(__ret__, 'acl_rules'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        region=pulumi.get(__ret__, 'region'))


@_utilities.lift_output_func(get_rdb_acl)
def get_rdb_acl_output(instance_id: Optional[pulumi.Input[str]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRdbAclResult]:
    """
    Gets information about the RDB instance network Access Control List.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_acl = scaleway.get_rdb_acl(instance_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str instance_id: The RDB instance ID.
    :param str region: `region`) The region in which the Database Instance should be created.
    """
    ...
