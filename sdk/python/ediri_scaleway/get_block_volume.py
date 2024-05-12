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
    'GetBlockVolumeResult',
    'AwaitableGetBlockVolumeResult',
    'get_block_volume',
    'get_block_volume_output',
]

@pulumi.output_type
class GetBlockVolumeResult:
    """
    A collection of values returned by getBlockVolume.
    """
    def __init__(__self__, id=None, iops=None, name=None, project_id=None, size_in_gb=None, snapshot_id=None, tags=None, volume_id=None, zone=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iops and not isinstance(iops, int):
            raise TypeError("Expected argument 'iops' to be a int")
        pulumi.set(__self__, "iops", iops)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if size_in_gb and not isinstance(size_in_gb, int):
            raise TypeError("Expected argument 'size_in_gb' to be a int")
        pulumi.set(__self__, "size_in_gb", size_in_gb)
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
    def iops(self) -> int:
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> int:
        return pulumi.get(self, "size_in_gb")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> str:
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


class AwaitableGetBlockVolumeResult(GetBlockVolumeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBlockVolumeResult(
            id=self.id,
            iops=self.iops,
            name=self.name,
            project_id=self.project_id,
            size_in_gb=self.size_in_gb,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            volume_id=self.volume_id,
            zone=self.zone)


def get_block_volume(name: Optional[str] = None,
                     project_id: Optional[str] = None,
                     volume_id: Optional[str] = None,
                     zone: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBlockVolumeResult:
    """
    Gets information about a Block Volume.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_volume = scaleway.get_block_volume(volume_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The name of the volume. Only one of `name` and `volume_id` should be specified.
    :param str project_id: The ID of the project the volume is associated with.
    :param str volume_id: The ID of the volume. Only one of `name` and `volume_id` should be specified.
    :param str zone: `zone`) The zone in which the volume exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['volumeId'] = volume_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getBlockVolume:getBlockVolume', __args__, opts=opts, typ=GetBlockVolumeResult).value

    return AwaitableGetBlockVolumeResult(
        id=pulumi.get(__ret__, 'id'),
        iops=pulumi.get(__ret__, 'iops'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        size_in_gb=pulumi.get(__ret__, 'size_in_gb'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        tags=pulumi.get(__ret__, 'tags'),
        volume_id=pulumi.get(__ret__, 'volume_id'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_block_volume)
def get_block_volume_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                            project_id: Optional[pulumi.Input[Optional[str]]] = None,
                            volume_id: Optional[pulumi.Input[Optional[str]]] = None,
                            zone: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBlockVolumeResult]:
    """
    Gets information about a Block Volume.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_volume = scaleway.get_block_volume(volume_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The name of the volume. Only one of `name` and `volume_id` should be specified.
    :param str project_id: The ID of the project the volume is associated with.
    :param str volume_id: The ID of the volume. Only one of `name` and `volume_id` should be specified.
    :param str zone: `zone`) The zone in which the volume exists.
    """
    ...
