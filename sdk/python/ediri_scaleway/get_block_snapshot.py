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
    'GetBlockSnapshotResult',
    'AwaitableGetBlockSnapshotResult',
    'get_block_snapshot',
    'get_block_snapshot_output',
]

@pulumi.output_type
class GetBlockSnapshotResult:
    """
    A collection of values returned by getBlockSnapshot.
    """
    def __init__(__self__, id=None, name=None, project_id=None, snapshot_id=None, tags=None, volume_id=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[str]:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[str]:
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetBlockSnapshotResult(GetBlockSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBlockSnapshotResult(
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            volume_id=self.volume_id,
            zone=self.zone)


def get_block_snapshot(name: Optional[str] = None,
                       project_id: Optional[str] = None,
                       snapshot_id: Optional[str] = None,
                       volume_id: Optional[str] = None,
                       zone: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBlockSnapshotResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['snapshotId'] = snapshot_id
    __args__['volumeId'] = volume_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getBlockSnapshot:getBlockSnapshot', __args__, opts=opts, typ=GetBlockSnapshotResult).value

    return AwaitableGetBlockSnapshotResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        tags=pulumi.get(__ret__, 'tags'),
        volume_id=pulumi.get(__ret__, 'volume_id'),
        zone=pulumi.get(__ret__, 'zone'))
def get_block_snapshot_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                              project_id: Optional[pulumi.Input[Optional[str]]] = None,
                              snapshot_id: Optional[pulumi.Input[Optional[str]]] = None,
                              volume_id: Optional[pulumi.Input[Optional[str]]] = None,
                              zone: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBlockSnapshotResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['snapshotId'] = snapshot_id
    __args__['volumeId'] = volume_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getBlockSnapshot:getBlockSnapshot', __args__, opts=opts, typ=GetBlockSnapshotResult)
    return __ret__.apply(lambda __response__: GetBlockSnapshotResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        snapshot_id=pulumi.get(__response__, 'snapshot_id'),
        tags=pulumi.get(__response__, 'tags'),
        volume_id=pulumi.get(__response__, 'volume_id'),
        zone=pulumi.get(__response__, 'zone')))
