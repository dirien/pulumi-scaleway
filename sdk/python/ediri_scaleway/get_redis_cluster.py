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
    'GetRedisClusterResult',
    'AwaitableGetRedisClusterResult',
    'get_redis_cluster',
    'get_redis_cluster_output',
]

@pulumi.output_type
class GetRedisClusterResult:
    """
    A collection of values returned by getRedisCluster.
    """
    def __init__(__self__, acls=None, certificate=None, cluster_id=None, cluster_size=None, created_at=None, id=None, name=None, node_type=None, password=None, private_networks=None, project_id=None, public_networks=None, settings=None, tags=None, tls_enabled=None, updated_at=None, user_name=None, version=None, zone=None):
        if acls and not isinstance(acls, list):
            raise TypeError("Expected argument 'acls' to be a list")
        pulumi.set(__self__, "acls", acls)
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_size and not isinstance(cluster_size, int):
            raise TypeError("Expected argument 'cluster_size' to be a int")
        pulumi.set(__self__, "cluster_size", cluster_size)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_type and not isinstance(node_type, str):
            raise TypeError("Expected argument 'node_type' to be a str")
        pulumi.set(__self__, "node_type", node_type)
        if password and not isinstance(password, str):
            raise TypeError("Expected argument 'password' to be a str")
        pulumi.set(__self__, "password", password)
        if private_networks and not isinstance(private_networks, list):
            raise TypeError("Expected argument 'private_networks' to be a list")
        pulumi.set(__self__, "private_networks", private_networks)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if public_networks and not isinstance(public_networks, list):
            raise TypeError("Expected argument 'public_networks' to be a list")
        pulumi.set(__self__, "public_networks", public_networks)
        if settings and not isinstance(settings, dict):
            raise TypeError("Expected argument 'settings' to be a dict")
        pulumi.set(__self__, "settings", settings)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if tls_enabled and not isinstance(tls_enabled, bool):
            raise TypeError("Expected argument 'tls_enabled' to be a bool")
        pulumi.set(__self__, "tls_enabled", tls_enabled)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def acls(self) -> Sequence['outputs.GetRedisClusterAclResult']:
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter
    def certificate(self) -> str:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[str]:
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> int:
        return pulumi.get(self, "cluster_size")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

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
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="privateNetworks")
    def private_networks(self) -> Sequence['outputs.GetRedisClusterPrivateNetworkResult']:
        return pulumi.get(self, "private_networks")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicNetworks")
    def public_networks(self) -> Sequence['outputs.GetRedisClusterPublicNetworkResult']:
        return pulumi.get(self, "public_networks")

    @property
    @pulumi.getter
    def settings(self) -> Mapping[str, str]:
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tlsEnabled")
    def tls_enabled(self) -> bool:
        return pulumi.get(self, "tls_enabled")

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
        return pulumi.get(self, "version")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetRedisClusterResult(GetRedisClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRedisClusterResult(
            acls=self.acls,
            certificate=self.certificate,
            cluster_id=self.cluster_id,
            cluster_size=self.cluster_size,
            created_at=self.created_at,
            id=self.id,
            name=self.name,
            node_type=self.node_type,
            password=self.password,
            private_networks=self.private_networks,
            project_id=self.project_id,
            public_networks=self.public_networks,
            settings=self.settings,
            tags=self.tags,
            tls_enabled=self.tls_enabled,
            updated_at=self.updated_at,
            user_name=self.user_name,
            version=self.version,
            zone=self.zone)


def get_redis_cluster(cluster_id: Optional[str] = None,
                      name: Optional[str] = None,
                      zone: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRedisClusterResult:
    """
    Gets information about a Redis cluster. For further information check our [api documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_cluster = scaleway.get_redis_cluster(cluster_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str cluster_id: The Redis cluster ID.
           Only one of the `name` and `cluster_id` should be specified.
    :param str name: The name of the Redis cluster.
           Only one of the `name` and `cluster_id` should be specified.
    :param str zone: `region`) The zone in which the server exists.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getRedisCluster:getRedisCluster', __args__, opts=opts, typ=GetRedisClusterResult).value

    return AwaitableGetRedisClusterResult(
        acls=pulumi.get(__ret__, 'acls'),
        certificate=pulumi.get(__ret__, 'certificate'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        cluster_size=pulumi.get(__ret__, 'cluster_size'),
        created_at=pulumi.get(__ret__, 'created_at'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        node_type=pulumi.get(__ret__, 'node_type'),
        password=pulumi.get(__ret__, 'password'),
        private_networks=pulumi.get(__ret__, 'private_networks'),
        project_id=pulumi.get(__ret__, 'project_id'),
        public_networks=pulumi.get(__ret__, 'public_networks'),
        settings=pulumi.get(__ret__, 'settings'),
        tags=pulumi.get(__ret__, 'tags'),
        tls_enabled=pulumi.get(__ret__, 'tls_enabled'),
        updated_at=pulumi.get(__ret__, 'updated_at'),
        user_name=pulumi.get(__ret__, 'user_name'),
        version=pulumi.get(__ret__, 'version'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_redis_cluster)
def get_redis_cluster_output(cluster_id: Optional[pulumi.Input[Optional[str]]] = None,
                             name: Optional[pulumi.Input[Optional[str]]] = None,
                             zone: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRedisClusterResult]:
    """
    Gets information about a Redis cluster. For further information check our [api documentation](https://developers.scaleway.com/en/products/redis/api/v1alpha1/#clusters-a85816)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_cluster = scaleway.get_redis_cluster(cluster_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str cluster_id: The Redis cluster ID.
           Only one of the `name` and `cluster_id` should be specified.
    :param str name: The name of the Redis cluster.
           Only one of the `name` and `cluster_id` should be specified.
    :param str zone: `region`) The zone in which the server exists.
    """
    ...
