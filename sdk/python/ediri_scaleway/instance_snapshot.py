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
from ._inputs import *

__all__ = ['InstanceSnapshotArgs', 'InstanceSnapshot']

@pulumi.input_type
class InstanceSnapshotArgs:
    def __init__(__self__, *,
                 import_: Optional[pulumi.Input['InstanceSnapshotImportArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceSnapshot resource.
        :param pulumi.Input['InstanceSnapshotImportArgs'] import_: Import a snapshot from a qcow2 file located in a bucket
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is
               associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] type: The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
               Updates to this field will recreate a new resource.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which
               the snapshot should be created.
        """
        if import_ is not None:
            pulumi.set(__self__, "import_", import_)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="import")
    def import_(self) -> Optional[pulumi.Input['InstanceSnapshotImportArgs']]:
        """
        Import a snapshot from a qcow2 file located in a bucket
        """
        return pulumi.get(self, "import_")

    @import_.setter
    def import_(self, value: Optional[pulumi.Input['InstanceSnapshotImportArgs']]):
        pulumi.set(self, "import_", value)

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
        `project_id`) The ID of the project the snapshot is
        associated with.
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
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
        Updates to this field will recreate a new resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
        `zone`) The zone in which
        the snapshot should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceSnapshotState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 import_: Optional[pulumi.Input['InstanceSnapshotImportArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 size_in_gb: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceSnapshot resources.
        :param pulumi.Input[str] created_at: The snapshot creation time.
        :param pulumi.Input['InstanceSnapshotImportArgs'] import_: Import a snapshot from a qcow2 file located in a bucket
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the snapshot is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is
               associated with.
        :param pulumi.Input[int] size_in_gb: (Optional) The size of the snapshot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] type: The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
               Updates to this field will recreate a new resource.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which
               the snapshot should be created.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if import_ is not None:
            pulumi.set(__self__, "import_", import_)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if size_in_gb is not None:
            pulumi.set(__self__, "size_in_gb", size_in_gb)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot creation time.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="import")
    def import_(self) -> Optional[pulumi.Input['InstanceSnapshotImportArgs']]:
        """
        Import a snapshot from a qcow2 file located in a bucket
        """
        return pulumi.get(self, "import_")

    @import_.setter
    def import_(self, value: Optional[pulumi.Input['InstanceSnapshotImportArgs']]):
        pulumi.set(self, "import_", value)

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
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the snapshot is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the snapshot is
        associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> Optional[pulumi.Input[int]]:
        """
        (Optional) The size of the snapshot.
        """
        return pulumi.get(self, "size_in_gb")

    @size_in_gb.setter
    def size_in_gb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_gb", value)

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
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
        Updates to this field will recreate a new resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

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
        `zone`) The zone in which
        the snapshot should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceSnapshot(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 import_: Optional[pulumi.Input[pulumi.InputType['InstanceSnapshotImportArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Compute Snapshots.
        For more information,
        see [the documentation](https://www.scaleway.com/en/developers/api/instance/#path-snapshots-list-snapshots).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.InstanceSnapshot("main", volume_id="11111111-1111-1111-1111-111111111111")
        ```

        ### Example with Unified type snapshot

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_instance_volume = scaleway.InstanceVolume("mainInstanceVolume",
            type="l_ssd",
            size_in_gb=10)
        main_instance_server = scaleway.InstanceServer("mainInstanceServer",
            image="ubuntu_jammy",
            type="DEV1-S",
            root_volume=scaleway.InstanceServerRootVolumeArgs(
                size_in_gb=10,
                volume_type="l_ssd",
            ),
            additional_volume_ids=[main_instance_volume.id])
        main_instance_snapshot = scaleway.InstanceSnapshot("mainInstanceSnapshot",
            volume_id=main_instance_volume.id,
            type="unified",
            opts=pulumi.ResourceOptions(depends_on=[main_instance_server]))
        ```

        ### Example importing a local qcow2 file

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        bucket = scaleway.ObjectBucket("bucket")
        qcow = scaleway.ObjectItem("qcow",
            bucket=bucket.name,
            key="server.qcow2",
            file="myqcow.qcow2")
        snapshot = scaleway.InstanceSnapshot("snapshot",
            type="unified",
            import_=scaleway.InstanceSnapshotImportArgs(
                bucket=qcow.bucket,
                key=qcow.key,
            ))
        ```

        ## Import

        Snapshots can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceSnapshot:InstanceSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InstanceSnapshotImportArgs']] import_: Import a snapshot from a qcow2 file located in a bucket
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is
               associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] type: The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
               Updates to this field will recreate a new resource.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which
               the snapshot should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[InstanceSnapshotArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Compute Snapshots.
        For more information,
        see [the documentation](https://www.scaleway.com/en/developers/api/instance/#path-snapshots-list-snapshots).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.InstanceSnapshot("main", volume_id="11111111-1111-1111-1111-111111111111")
        ```

        ### Example with Unified type snapshot

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_instance_volume = scaleway.InstanceVolume("mainInstanceVolume",
            type="l_ssd",
            size_in_gb=10)
        main_instance_server = scaleway.InstanceServer("mainInstanceServer",
            image="ubuntu_jammy",
            type="DEV1-S",
            root_volume=scaleway.InstanceServerRootVolumeArgs(
                size_in_gb=10,
                volume_type="l_ssd",
            ),
            additional_volume_ids=[main_instance_volume.id])
        main_instance_snapshot = scaleway.InstanceSnapshot("mainInstanceSnapshot",
            volume_id=main_instance_volume.id,
            type="unified",
            opts=pulumi.ResourceOptions(depends_on=[main_instance_server]))
        ```

        ### Example importing a local qcow2 file

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        bucket = scaleway.ObjectBucket("bucket")
        qcow = scaleway.ObjectItem("qcow",
            bucket=bucket.name,
            key="server.qcow2",
            file="myqcow.qcow2")
        snapshot = scaleway.InstanceSnapshot("snapshot",
            type="unified",
            import_=scaleway.InstanceSnapshotImportArgs(
                bucket=qcow.bucket,
                key=qcow.key,
            ))
        ```

        ## Import

        Snapshots can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceSnapshot:InstanceSnapshot main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceSnapshotArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceSnapshotArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 import_: Optional[pulumi.Input[pulumi.InputType['InstanceSnapshotImportArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceSnapshotArgs.__new__(InstanceSnapshotArgs)

            __props__.__dict__["import_"] = import_
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["type"] = type
            __props__.__dict__["volume_id"] = volume_id
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["size_in_gb"] = None
        super(InstanceSnapshot, __self__).__init__(
            'scaleway:index/instanceSnapshot:InstanceSnapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            import_: Optional[pulumi.Input[pulumi.InputType['InstanceSnapshotImportArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            size_in_gb: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            volume_id: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceSnapshot':
        """
        Get an existing InstanceSnapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The snapshot creation time.
        :param pulumi.Input[pulumi.InputType['InstanceSnapshotImportArgs']] import_: Import a snapshot from a qcow2 file located in a bucket
        :param pulumi.Input[str] name: The name of the snapshot. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the snapshot is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the snapshot is
               associated with.
        :param pulumi.Input[int] size_in_gb: (Optional) The size of the snapshot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the snapshot.
        :param pulumi.Input[str] type: The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
               Updates to this field will recreate a new resource.
        :param pulumi.Input[str] volume_id: The ID of the volume to take a snapshot from.
        :param pulumi.Input[str] zone: `zone`) The zone in which
               the snapshot should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceSnapshotState.__new__(_InstanceSnapshotState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["import_"] = import_
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["size_in_gb"] = size_in_gb
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["volume_id"] = volume_id
        __props__.__dict__["zone"] = zone
        return InstanceSnapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The snapshot creation time.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="import")
    def import_(self) -> pulumi.Output[Optional['outputs.InstanceSnapshotImport']]:
        """
        Import a snapshot from a qcow2 file located in a bucket
        """
        return pulumi.get(self, "import_")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the snapshot. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the snapshot is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the snapshot is
        associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="sizeInGb")
    def size_in_gb(self) -> pulumi.Output[int]:
        """
        (Optional) The size of the snapshot.
        """
        return pulumi.get(self, "size_in_gb")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tags to apply to the snapshot.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The snapshot's volume type.  The possible values are: `b_ssd` (Block SSD), `l_ssd` (Local SSD) and `unified`.
        Updates to this field will recreate a new resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the volume to take a snapshot from.
        """
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which
        the snapshot should be created.
        """
        return pulumi.get(self, "zone")

