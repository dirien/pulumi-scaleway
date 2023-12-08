# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BlockSnapshotArgs', 'BlockSnapshot']

@pulumi.input_type
class BlockSnapshotArgs:
    def __init__(__self__, *,
                 volume_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BlockSnapshot resource.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] zone: `zone`) The zone in which the snapshot should be created.
        """
        pulumi.set(__self__, "volume_id", volume_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Input[str]:
        """
        The ID of the volume to take a snapshot from.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "volume_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the snapshot is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tags to apply to the snapshot.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the snapshot should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _BlockSnapshotState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BlockSnapshot resources.
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which the snapshot should be created.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the snapshot is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tags to apply to the snapshot.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the volume to take a snapshot from.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_id", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the snapshot should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class BlockSnapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Block Snapshots.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/block/).

        ## Example

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        block_snapshot = scaleway.BlockSnapshot("blockSnapshot", volume_id="11111111-1111-1111-1111-111111111111")
        ```

        ## Import

        Block Snapshots can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/blockSnapshot:BlockSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which the snapshot should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BlockSnapshotArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Block Snapshots.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/block/).

        ## Example

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        block_snapshot = scaleway.BlockSnapshot("blockSnapshot", volume_id="11111111-1111-1111-1111-111111111111")
        ```

        ## Import

        Block Snapshots can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/blockSnapshot:BlockSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param BlockSnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BlockSnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BlockSnapshotArgs.__new__(BlockSnapshotArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["tags"] = tags
            if volume_id is None and not opts.urn:
                raise TypeError("Missing required property 'volume_id'")
            __props__.__dict__["volume_id"] = volume_id
            __props__.__dict__["zone"] = zone
        super(BlockSnapshot, __self__).__init__(
            'scaleway:index/blockSnapshot:BlockSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            volume_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'BlockSnapshot':
        """
        Get an existing BlockSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which the snapshot should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BlockSnapshotState.__new__(_BlockSnapshotState)

        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["volume_id"] = volume_id
        __props__.__dict__["zone"] = zone
        return BlockSnapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the snapshot. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the snapshot is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tags to apply to the snapshot.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[str]:
        """
        The ID of the volume to take a snapshot from.
        """
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the snapshot should be created.
        """
        return pulumi.get(self, "zone")

