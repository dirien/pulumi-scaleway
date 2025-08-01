# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['BlockVolumeArgs', 'BlockVolume']

@pulumi.input_type
class BlockVolumeArgs:
    def __init__(__self__, *,
                 iops: pulumi.Input[builtins.int],
                 instance_volume_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 size_in_gb: Optional[pulumi.Input[builtins.int]] = None,
                 snapshot_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a BlockVolume resource.
        :param pulumi.Input[builtins.int] iops: The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        :param pulumi.Input[builtins.str] instance_volume_id: The instance volume to create the block volume from
        :param pulumi.Input[builtins.str] name: The name of the volume. If not provided, a name will be randomly generated.
        :param pulumi.Input[builtins.str] project_id: ). The ID of the Project the volume is associated with.
        :param pulumi.Input[builtins.int] size_in_gb: The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[builtins.str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[builtins.str] zone: ). The zone in which the volume should be created.
        """
        pulumi.set(__self__, "iops", iops)
        if instance_volume_id is not None:
            pulumi.set(__self__, "instance_volume_id", instance_volume_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def iops(self) -> pulumi.Input[builtins.int]:
        """
        The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter(name="instanceVolumeId")
    def instance_volume_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The instance volume to create the block volume from
        """
        return pulumi.get(self, "instance_volume_id")

    @instance_volume_id.setter
    def instance_volume_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_volume_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the volume. If not provided, a name will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ). The ID of the Project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "size_in_gb", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ). The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _BlockVolumeState:
    def __init__(__self__, *,
                 instance_volume_id: Optional[pulumi.Input[builtins.str]] = None,
                 iops: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 size_in_gb: Optional[pulumi.Input[builtins.int]] = None,
                 snapshot_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering BlockVolume resources.
        :param pulumi.Input[builtins.str] instance_volume_id: The instance volume to create the block volume from
        :param pulumi.Input[builtins.int] iops: The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        :param pulumi.Input[builtins.str] name: The name of the volume. If not provided, a name will be randomly generated.
        :param pulumi.Input[builtins.str] project_id: ). The ID of the Project the volume is associated with.
        :param pulumi.Input[builtins.int] size_in_gb: The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[builtins.str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[builtins.str] zone: ). The zone in which the volume should be created.
        """
        if instance_volume_id is not None:
            pulumi.set(__self__, "instance_volume_id", instance_volume_id)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="instanceVolumeId")
    def instance_volume_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The instance volume to create the block volume from
        """
        return pulumi.get(self, "instance_volume_id")

    @instance_volume_id.setter
    def instance_volume_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "instance_volume_id", value)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the volume. If not provided, a name will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ). The ID of the Project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "size_in_gb", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        ). The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "zone", value)


@pulumi.type_token("scaleway:index/blockVolume:BlockVolume")
class BlockVolume(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_volume_id: Optional[pulumi.Input[builtins.str]] = None,
                 iops: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 size_in_gb: Optional[pulumi.Input[builtins.int]] = None,
                 snapshot_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        The `BlockVolume` resource is used to create and manage Scaleway Block Storage volumes.

        Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/block-storage/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.

        ## Example Usage

        ### Create a Block Storage volume

        The following command allows you to create a Block Storage volume of 20 GB with a 5000 [IOPS](https://www.scaleway.com/en/docs/block-storage/concepts/#iops).

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        block_volume = scaleway.BlockVolume("blockVolume",
            iops=5000,
            size_in_gb=20)
        ```

        ### With snapshot

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        base = scaleway.BlockVolume("base",
            iops=5000,
            size_in_gb=20)
        main_block_snapshot = scaleway.BlockSnapshot("mainBlockSnapshot", volume_id=base.id)
        main_block_volume = scaleway.BlockVolume("mainBlockVolume",
            iops=5000,
            snapshot_id=main_block_snapshot.id)
        ```

        ## Import

        This section explains how to import a Block Storage volume using the zoned ID (`{zone}/{id}`) format.

        bash

        ```sh
        $ pulumi import scaleway:index/blockVolume:BlockVolume block_volume fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] instance_volume_id: The instance volume to create the block volume from
        :param pulumi.Input[builtins.int] iops: The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        :param pulumi.Input[builtins.str] name: The name of the volume. If not provided, a name will be randomly generated.
        :param pulumi.Input[builtins.str] project_id: ). The ID of the Project the volume is associated with.
        :param pulumi.Input[builtins.int] size_in_gb: The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[builtins.str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[builtins.str] zone: ). The zone in which the volume should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BlockVolumeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `BlockVolume` resource is used to create and manage Scaleway Block Storage volumes.

        Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/block-storage/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.

        ## Example Usage

        ### Create a Block Storage volume

        The following command allows you to create a Block Storage volume of 20 GB with a 5000 [IOPS](https://www.scaleway.com/en/docs/block-storage/concepts/#iops).

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        block_volume = scaleway.BlockVolume("blockVolume",
            iops=5000,
            size_in_gb=20)
        ```

        ### With snapshot

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        base = scaleway.BlockVolume("base",
            iops=5000,
            size_in_gb=20)
        main_block_snapshot = scaleway.BlockSnapshot("mainBlockSnapshot", volume_id=base.id)
        main_block_volume = scaleway.BlockVolume("mainBlockVolume",
            iops=5000,
            snapshot_id=main_block_snapshot.id)
        ```

        ## Import

        This section explains how to import a Block Storage volume using the zoned ID (`{zone}/{id}`) format.

        bash

        ```sh
        $ pulumi import scaleway:index/blockVolume:BlockVolume block_volume fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param BlockVolumeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BlockVolumeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_volume_id: Optional[pulumi.Input[builtins.str]] = None,
                 iops: Optional[pulumi.Input[builtins.int]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 size_in_gb: Optional[pulumi.Input[builtins.int]] = None,
                 snapshot_id: Optional[pulumi.Input[builtins.str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 zone: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BlockVolumeArgs.__new__(BlockVolumeArgs)

            __props__.__dict__["instance_volume_id"] = instance_volume_id
            if iops is None and not opts.urn:
                raise TypeError("Missing required property 'iops'")
            __props__.__dict__["iops"] = iops
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["size_in_gb"] = size_in_gb
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
        super(BlockVolume, __self__).__init__(
            'scaleway:index/blockVolume:BlockVolume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_volume_id: Optional[pulumi.Input[builtins.str]] = None,
            iops: Optional[pulumi.Input[builtins.int]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            size_in_gb: Optional[pulumi.Input[builtins.int]] = None,
            snapshot_id: Optional[pulumi.Input[builtins.str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            zone: Optional[pulumi.Input[builtins.str]] = None) -> 'BlockVolume':
        """
        Get an existing BlockVolume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] instance_volume_id: The instance volume to create the block volume from
        :param pulumi.Input[builtins.int] iops: The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        :param pulumi.Input[builtins.str] name: The name of the volume. If not provided, a name will be randomly generated.
        :param pulumi.Input[builtins.str] project_id: ). The ID of the Project the volume is associated with.
        :param pulumi.Input[builtins.int] size_in_gb: The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        :param pulumi.Input[builtins.str] snapshot_id: If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] tags: A list of tags to apply to the volume.
        :param pulumi.Input[builtins.str] zone: ). The zone in which the volume should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BlockVolumeState.__new__(_BlockVolumeState)

        __props__.__dict__["instance_volume_id"] = instance_volume_id
        __props__.__dict__["iops"] = iops
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["size_in_gb"] = size_in_gb
        __props__.__dict__["snapshot_id"] = snapshot_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone"] = zone
        return BlockVolume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceVolumeId")
    def instance_volume_id(self) -> pulumi.Output[builtins.str]:
        """
        The instance volume to create the block volume from
        """
        return pulumi.get(self, "instance_volume_id")

    @property
    @pulumi.getter
    def iops(self) -> pulumi.Output[builtins.int]:
        """
        The maximum [IOPs](https://www.scaleway.com/en/docs/block-storage/concepts/#iops) expected, must match available options.
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the volume. If not provided, a name will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        ). The ID of the Project the volume is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> pulumi.Output[builtins.int]:
        """
        The size of the volume in gigabytes. Only one of `size_in_gb`, and `snapshot_id` should be specified.
        """
        return pulumi.get(self, "size_in_gb")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        If set, the new volume will be created from this snapshot. Only one of `size_in_gb`, `snapshot_id` should be specified.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A list of tags to apply to the volume.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[builtins.str]:
        """
        ). The zone in which the volume should be created.
        """
        return pulumi.get(self, "zone")

