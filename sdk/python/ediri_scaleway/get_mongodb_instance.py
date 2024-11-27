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
    'GetMongodbInstanceResult',
    'AwaitableGetMongodbInstanceResult',
    'get_mongodb_instance',
    'get_mongodb_instance_output',
]

@pulumi.output_type
class GetMongodbInstanceResult:
    """
    A collection of values returned by getMongodbInstance.
    """
    def __init__(__self__, created_at=None, id=None, instance_id=None, name=None, node_number=None, node_type=None, password=None, project_id=None, public_networks=None, region=None, settings=None, snapshot_id=None, tags=None, updated_at=None, user_name=None, version=None, volume_size_in_gb=None, volume_type=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_number and not isinstance(node_number, int):
            raise TypeError("Expected argument 'node_number' to be a int")
        pulumi.set(__self__, "node_number", node_number)
        if node_type and not isinstance(node_type, str):
            raise TypeError("Expected argument 'node_type' to be a str")
        pulumi.set(__self__, "node_type", node_type)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if public_networks and not isinstance(public_networks, list):
            raise TypeError("Expected argument 'public_networks' to be a list")
        pulumi.set(__self__, "public_networks", public_networks)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if settings and not isinstance(settings, dict):
            raise TypeError("Expected argument 'settings' to be a dict")
        pulumi.set(__self__, "settings", settings)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if volume_size_in_gb and not isinstance(volume_size_in_gb, int):
            raise TypeError("Expected argument 'volume_size_in_gb' to be a int")
        pulumi.set(__self__, "volume_size_in_gb", volume_size_in_gb)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time the MongoDB® instance was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the MongoDB® instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeNumber")
    def node_number(self) -> int:
        """
        The number of nodes in the MongoDB® cluster.
        """
        return pulumi.get(self, "node_number")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        """
        The type of MongoDB® node.
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        The ID of the project the instance belongs to.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicNetworks")
    def public_networks(self) -> Sequence['outputs.GetMongodbInstancePublicNetworkResult']:
        """
        The details of the public network configuration, if applicable.
        """
        return pulumi.get(self, "public_networks")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def settings(self) -> Mapping[str, str]:
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> str:
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        A list of tags attached to the MongoDB® instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of MongoDB® running on the instance.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="volumeSizeInGb")
    def volume_size_in_gb(self) -> int:
        """
        The size of the attached volume, in GB.
        """
        return pulumi.get(self, "volume_size_in_gb")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        """
        The type of volume attached to the MongoDB® instance.
        """
        return pulumi.get(self, "volume_type")


class AwaitableGetMongodbInstanceResult(GetMongodbInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMongodbInstanceResult(
            created_at=self.created_at,
            id=self.id,
            instance_id=self.instance_id,
            name=self.name,
            node_number=self.node_number,
            node_type=self.node_type,
            password=self.password,
            project_id=self.project_id,
            public_networks=self.public_networks,
            region=self.region,
            settings=self.settings,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            updated_at=self.updated_at,
            user_name=self.user_name,
            version=self.version,
            volume_size_in_gb=self.volume_size_in_gb,
            volume_type=self.volume_type)


def get_mongodb_instance(instance_id: Optional[str] = None,
                         name: Optional[str] = None,
                         project_id: Optional[str] = None,
                         region: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMongodbInstanceResult:
    """
    Gets information about a MongoDB® Instance.

    For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)


    :param str instance_id: The MongoDB® instance ID.
           
           > **Note** You must specify at least one: `name` or `instance_id`.
    :param str name: The name of the MongoDB® instance.
    :param str project_id: The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
    :param str region: `region`) The region in which the MongoDB® Instance exists.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getMongodbInstance:getMongodbInstance', __args__, opts=opts, typ=GetMongodbInstanceResult).value

    return AwaitableGetMongodbInstanceResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        node_number=pulumi.get(__ret__, 'node_number'),
        node_type=pulumi.get(__ret__, 'node_type'),
        password=pulumi.get(__ret__, 'password'),
        project_id=pulumi.get(__ret__, 'project_id'),
        public_networks=pulumi.get(__ret__, 'public_networks'),
        region=pulumi.get(__ret__, 'region'),
        settings=pulumi.get(__ret__, 'settings'),
        snapshot_id=pulumi.get(__ret__, 'snapshot_id'),
        tags=pulumi.get(__ret__, 'tags'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        user_name=pulumi.get(__ret__, 'user_name'),
        version=pulumi.get(__ret__, 'version'),
        volume_size_in_gb=pulumi.get(__ret__, 'volume_size_in_gb'),
        volume_type=pulumi.get(__ret__, 'volume_type'))
def get_mongodb_instance_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                project_id: Optional[pulumi.Input[Optional[str]]] = None,
                                region: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMongodbInstanceResult]:
    """
    Gets information about a MongoDB® Instance.

    For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)


    :param str instance_id: The MongoDB® instance ID.
           
           > **Note** You must specify at least one: `name` or `instance_id`.
    :param str name: The name of the MongoDB® instance.
    :param str project_id: The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
    :param str region: `region`) The region in which the MongoDB® Instance exists.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getMongodbInstance:getMongodbInstance', __args__, opts=opts, typ=GetMongodbInstanceResult)
    return __ret__.apply(lambda __response__: GetMongodbInstanceResult(
        created_at=pulumi.get(__response__, 'created_at'),
        id=pulumi.get(__response__, 'id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        name=pulumi.get(__response__, 'name'),
        node_number=pulumi.get(__response__, 'node_number'),
        node_type=pulumi.get(__response__, 'node_type'),
        password=pulumi.get(__response__, 'password'),
        project_id=pulumi.get(__response__, 'project_id'),
        public_networks=pulumi.get(__response__, 'public_networks'),
        region=pulumi.get(__response__, 'region'),
        settings=pulumi.get(__response__, 'settings'),
        snapshot_id=pulumi.get(__response__, 'snapshot_id'),
        tags=pulumi.get(__response__, 'tags'),
        updated_at=pulumi.get(__response__, 'updated_at'),
        user_name=pulumi.get(__response__, 'user_name'),
        version=pulumi.get(__response__, 'version'),
        volume_size_in_gb=pulumi.get(__response__, 'volume_size_in_gb'),
        volume_type=pulumi.get(__response__, 'volume_type')))
