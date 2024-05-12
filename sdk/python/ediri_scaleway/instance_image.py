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

__all__ = ['InstanceImageArgs', 'InstanceImage']

@pulumi.input_type
class InstanceImageArgs:
    def __init__(__self__, *,
                 root_volume_id: pulumi.Input[str],
                 additional_volume_ids: Optional[pulumi.Input[str]] = None,
                 architecture: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceImage resource.
        :param pulumi.Input[str] root_volume_id: The ID of the snapshot of the volume to be used as root in the image.
        :param pulumi.Input[str] additional_volume_ids: List of IDs of the snapshots of the additional volumes to be attached to the image.
               
               > **Important:** For now it is only possible to have 1 additional_volume.
        :param pulumi.Input[str] architecture: The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        :param pulumi.Input[str] name: The name of the image. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: The ID of the project the image is associated with.
        :param pulumi.Input[bool] public: Set to `true` if the image is public.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the image.
        :param pulumi.Input[str] zone: The zone in which the image should be created.
        """
        pulumi.set(__self__, "root_volume_id", root_volume_id)
        if additional_volume_ids is not None:
            pulumi.set(__self__, "additional_volume_ids", additional_volume_ids)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="rootVolumeId")
    def root_volume_id(self) -> pulumi.Input[str]:
        """
        The ID of the snapshot of the volume to be used as root in the image.
        """
        return pulumi.get(self, "root_volume_id")

    @root_volume_id.setter
    def root_volume_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_volume_id", value)

    @property
    @pulumi.getter(name="additionalVolumeIds")
    def additional_volume_ids(self) -> Optional[pulumi.Input[str]]:
        """
        List of IDs of the snapshots of the additional volumes to be attached to the image.

        > **Important:** For now it is only possible to have 1 additional_volume.
        """
        return pulumi.get(self, "additional_volume_ids")

    @additional_volume_ids.setter
    def additional_volume_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_volume_ids", value)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input[str]]:
        """
        The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the image. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project the image is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` if the image is public.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tags to apply to the image.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone in which the image should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _InstanceImageState:
    def __init__(__self__, *,
                 additional_volume_ids: Optional[pulumi.Input[str]] = None,
                 additional_volumes: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceImageAdditionalVolumeArgs']]]] = None,
                 architecture: Optional[pulumi.Input[str]] = None,
                 creation_date: Optional[pulumi.Input[str]] = None,
                 from_server_id: Optional[pulumi.Input[str]] = None,
                 modification_date: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 root_volume_id: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceImage resources.
        :param pulumi.Input[str] additional_volume_ids: List of IDs of the snapshots of the additional volumes to be attached to the image.
               
               > **Important:** For now it is only possible to have 1 additional_volume.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceImageAdditionalVolumeArgs']]] additional_volumes: The description of the extra volumes attached to the image.
        :param pulumi.Input[str] architecture: The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        :param pulumi.Input[str] creation_date: Date of the volume creation.
        :param pulumi.Input[str] from_server_id: ID of the server the image is based on (in case it is a backup).
        :param pulumi.Input[str] modification_date: Date of volume latest update.
        :param pulumi.Input[str] name: The name of the image. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the image is associated with.
        :param pulumi.Input[str] project_id: The ID of the project the image is associated with.
        :param pulumi.Input[bool] public: Set to `true` if the image is public.
        :param pulumi.Input[str] root_volume_id: The ID of the snapshot of the volume to be used as root in the image.
        :param pulumi.Input[str] state: State of the volume.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the image.
        :param pulumi.Input[str] zone: The zone in which the image should be created.
        """
        if additional_volume_ids is not None:
            pulumi.set(__self__, "additional_volume_ids", additional_volume_ids)
        if additional_volumes is not None:
            pulumi.set(__self__, "additional_volumes", additional_volumes)
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if creation_date is not None:
            pulumi.set(__self__, "creation_date", creation_date)
        if from_server_id is not None:
            pulumi.set(__self__, "from_server_id", from_server_id)
        if modification_date is not None:
            pulumi.set(__self__, "modification_date", modification_date)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if public is not None:
            pulumi.set(__self__, "public", public)
        if root_volume_id is not None:
            pulumi.set(__self__, "root_volume_id", root_volume_id)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="additionalVolumeIds")
    def additional_volume_ids(self) -> Optional[pulumi.Input[str]]:
        """
        List of IDs of the snapshots of the additional volumes to be attached to the image.

        > **Important:** For now it is only possible to have 1 additional_volume.
        """
        return pulumi.get(self, "additional_volume_ids")

    @additional_volume_ids.setter
    def additional_volume_ids(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "additional_volume_ids", value)

    @property
    @pulumi.getter(name="additionalVolumes")
    def additional_volumes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceImageAdditionalVolumeArgs']]]]:
        """
        The description of the extra volumes attached to the image.
        """
        return pulumi.get(self, "additional_volumes")

    @additional_volumes.setter
    def additional_volumes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceImageAdditionalVolumeArgs']]]]):
        pulumi.set(self, "additional_volumes", value)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[pulumi.Input[str]]:
        """
        The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        """
        return pulumi.get(self, "architecture")

    @architecture.setter
    def architecture(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "architecture", value)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date of the volume creation.
        """
        return pulumi.get(self, "creation_date")

    @creation_date.setter
    def creation_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_date", value)

    @property
    @pulumi.getter(name="fromServerId")
    def from_server_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the server the image is based on (in case it is a backup).
        """
        return pulumi.get(self, "from_server_id")

    @from_server_id.setter
    def from_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_server_id", value)

    @property
    @pulumi.getter(name="modificationDate")
    def modification_date(self) -> Optional[pulumi.Input[str]]:
        """
        Date of volume latest update.
        """
        return pulumi.get(self, "modification_date")

    @modification_date.setter
    def modification_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modification_date", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the image. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the image is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project the image is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` if the image is public.
        """
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)

    @property
    @pulumi.getter(name="rootVolumeId")
    def root_volume_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the snapshot of the volume to be used as root in the image.
        """
        return pulumi.get(self, "root_volume_id")

    @root_volume_id.setter
    def root_volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_volume_id", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the volume.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of tags to apply to the image.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone in which the image should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class InstanceImage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_volume_ids: Optional[pulumi.Input[str]] = None,
                 architecture: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 root_volume_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Compute Images.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#images-41389b).

        ## Example Usage

        ### From a volume

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        volume = scaleway.InstanceVolume("volume",
            type="b_ssd",
            size_in_gb=20)
        volume_snapshot = scaleway.InstanceSnapshot("volumeSnapshot", volume_id=volume.id)
        volume_image = scaleway.InstanceImage("volumeImage", root_volume_id=volume_snapshot.id)
        ```

        ### From a server

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        server = scaleway.InstanceServer("server",
            image="ubuntu_jammy",
            type="DEV1-S")
        server_snapshot = scaleway.InstanceSnapshot("serverSnapshot", volume_id=scaleway_instance_server["main"]["root_volume"][0]["volume_id"])
        server_image = scaleway.InstanceImage("serverImage", root_volume_id=server_snapshot.id)
        ```

        ## Import

        Images can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceImage:InstanceImage main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] additional_volume_ids: List of IDs of the snapshots of the additional volumes to be attached to the image.
               
               > **Important:** For now it is only possible to have 1 additional_volume.
        :param pulumi.Input[str] architecture: The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        :param pulumi.Input[str] name: The name of the image. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: The ID of the project the image is associated with.
        :param pulumi.Input[bool] public: Set to `true` if the image is public.
        :param pulumi.Input[str] root_volume_id: The ID of the snapshot of the volume to be used as root in the image.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the image.
        :param pulumi.Input[str] zone: The zone in which the image should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceImageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Compute Images.
        For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#images-41389b).

        ## Example Usage

        ### From a volume

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        volume = scaleway.InstanceVolume("volume",
            type="b_ssd",
            size_in_gb=20)
        volume_snapshot = scaleway.InstanceSnapshot("volumeSnapshot", volume_id=volume.id)
        volume_image = scaleway.InstanceImage("volumeImage", root_volume_id=volume_snapshot.id)
        ```

        ### From a server

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        server = scaleway.InstanceServer("server",
            image="ubuntu_jammy",
            type="DEV1-S")
        server_snapshot = scaleway.InstanceSnapshot("serverSnapshot", volume_id=scaleway_instance_server["main"]["root_volume"][0]["volume_id"])
        server_image = scaleway.InstanceImage("serverImage", root_volume_id=server_snapshot.id)
        ```

        ## Import

        Images can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceImage:InstanceImage main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceImageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceImageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_volume_ids: Optional[pulumi.Input[str]] = None,
                 architecture: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 root_volume_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceImageArgs.__new__(InstanceImageArgs)

            __props__.__dict__["additional_volume_ids"] = additional_volume_ids
            __props__.__dict__["architecture"] = architecture
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["public"] = public
            if root_volume_id is None and not opts.urn:
                raise TypeError("Missing required property 'root_volume_id'")
            __props__.__dict__["root_volume_id"] = root_volume_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["zone"] = zone
            __props__.__dict__["additional_volumes"] = None
            __props__.__dict__["creation_date"] = None
            __props__.__dict__["from_server_id"] = None
            __props__.__dict__["modification_date"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["state"] = None
        super(InstanceImage, __self__).__init__(
            'scaleway:index/instanceImage:InstanceImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_volume_ids: Optional[pulumi.Input[str]] = None,
            additional_volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceImageAdditionalVolumeArgs']]]]] = None,
            architecture: Optional[pulumi.Input[str]] = None,
            creation_date: Optional[pulumi.Input[str]] = None,
            from_server_id: Optional[pulumi.Input[str]] = None,
            modification_date: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            public: Optional[pulumi.Input[bool]] = None,
            root_volume_id: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'InstanceImage':
        """
        Get an existing InstanceImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] additional_volume_ids: List of IDs of the snapshots of the additional volumes to be attached to the image.
               
               > **Important:** For now it is only possible to have 1 additional_volume.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceImageAdditionalVolumeArgs']]]] additional_volumes: The description of the extra volumes attached to the image.
        :param pulumi.Input[str] architecture: The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        :param pulumi.Input[str] creation_date: Date of the volume creation.
        :param pulumi.Input[str] from_server_id: ID of the server the image is based on (in case it is a backup).
        :param pulumi.Input[str] modification_date: Date of volume latest update.
        :param pulumi.Input[str] name: The name of the image. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the image is associated with.
        :param pulumi.Input[str] project_id: The ID of the project the image is associated with.
        :param pulumi.Input[bool] public: Set to `true` if the image is public.
        :param pulumi.Input[str] root_volume_id: The ID of the snapshot of the volume to be used as root in the image.
        :param pulumi.Input[str] state: State of the volume.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of tags to apply to the image.
        :param pulumi.Input[str] zone: The zone in which the image should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceImageState.__new__(_InstanceImageState)

        __props__.__dict__["additional_volume_ids"] = additional_volume_ids
        __props__.__dict__["additional_volumes"] = additional_volumes
        __props__.__dict__["architecture"] = architecture
        __props__.__dict__["creation_date"] = creation_date
        __props__.__dict__["from_server_id"] = from_server_id
        __props__.__dict__["modification_date"] = modification_date
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["public"] = public
        __props__.__dict__["root_volume_id"] = root_volume_id
        __props__.__dict__["state"] = state
        __props__.__dict__["tags"] = tags
        __props__.__dict__["zone"] = zone
        return InstanceImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalVolumeIds")
    def additional_volume_ids(self) -> pulumi.Output[Optional[str]]:
        """
        List of IDs of the snapshots of the additional volumes to be attached to the image.

        > **Important:** For now it is only possible to have 1 additional_volume.
        """
        return pulumi.get(self, "additional_volume_ids")

    @property
    @pulumi.getter(name="additionalVolumes")
    def additional_volumes(self) -> pulumi.Output[Sequence['outputs.InstanceImageAdditionalVolume']]:
        """
        The description of the extra volumes attached to the image.
        """
        return pulumi.get(self, "additional_volumes")

    @property
    @pulumi.getter
    def architecture(self) -> pulumi.Output[Optional[str]]:
        """
        The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        Date of the volume creation.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="fromServerId")
    def from_server_id(self) -> pulumi.Output[str]:
        """
        ID of the server the image is based on (in case it is a backup).
        """
        return pulumi.get(self, "from_server_id")

    @property
    @pulumi.getter(name="modificationDate")
    def modification_date(self) -> pulumi.Output[str]:
        """
        Date of volume latest update.
        """
        return pulumi.get(self, "modification_date")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the image. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the image is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The ID of the project the image is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def public(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` if the image is public.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter(name="rootVolumeId")
    def root_volume_id(self) -> pulumi.Output[str]:
        """
        The ID of the snapshot of the volume to be used as root in the image.
        """
        return pulumi.get(self, "root_volume_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the volume.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of tags to apply to the image.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone in which the image should be created.
        """
        return pulumi.get(self, "zone")

