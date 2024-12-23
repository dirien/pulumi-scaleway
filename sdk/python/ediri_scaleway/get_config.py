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
    'GetConfigResult',
    'AwaitableGetConfigResult',
    'get_config',
    'get_config_output',
]

@pulumi.output_type
class GetConfigResult:
    """
    A collection of values returned by getConfig.
    """
    def __init__(__self__, access_key=None, access_key_source=None, id=None, project_id=None, project_id_source=None, region=None, region_source=None, secret_key=None, secret_key_source=None, zone=None, zone_source=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if access_key_source and not isinstance(access_key_source, str):
            raise TypeError("Expected argument 'access_key_source' to be a str")
        pulumi.set(__self__, "access_key_source", access_key_source)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_id_source and not isinstance(project_id_source, str):
            raise TypeError("Expected argument 'project_id_source' to be a str")
        pulumi.set(__self__, "project_id_source", project_id_source)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if region_source and not isinstance(region_source, str):
            raise TypeError("Expected argument 'region_source' to be a str")
        pulumi.set(__self__, "region_source", region_source)
        if secret_key and not isinstance(secret_key, str):
            raise TypeError("Expected argument 'secret_key' to be a str")
        pulumi.set(__self__, "secret_key", secret_key)
        if secret_key_source and not isinstance(secret_key_source, str):
            raise TypeError("Expected argument 'secret_key_source' to be a str")
        pulumi.set(__self__, "secret_key_source", secret_key_source)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)
        if zone_source and not isinstance(zone_source, str):
            raise TypeError("Expected argument 'zone_source' to be a str")
        pulumi.set(__self__, "zone_source", zone_source)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> str:
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="accessKeySource")
    def access_key_source(self) -> str:
        return pulumi.get(self, "access_key_source")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectIdSource")
    def project_id_source(self) -> str:
        return pulumi.get(self, "project_id_source")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionSource")
    def region_source(self) -> str:
        return pulumi.get(self, "region_source")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> str:
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="secretKeySource")
    def secret_key_source(self) -> str:
        return pulumi.get(self, "secret_key_source")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")

    @property
    @pulumi.getter(name="zoneSource")
    def zone_source(self) -> str:
        return pulumi.get(self, "zone_source")


class AwaitableGetConfigResult(GetConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigResult(
            access_key=self.access_key,
            access_key_source=self.access_key_source,
            id=self.id,
            project_id=self.project_id,
            project_id_source=self.project_id_source,
            region=self.region,
            region_source=self.region_source,
            secret_key=self.secret_key,
            secret_key_source=self.secret_key_source,
            zone=self.zone,
            zone_source=self.zone_source)


def get_config(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getConfig:getConfig', __args__, opts=opts, typ=GetConfigResult).value

    return AwaitableGetConfigResult(
        access_key=pulumi.get(__ret__, 'access_key'),
        access_key_source=pulumi.get(__ret__, 'access_key_source'),
        id=pulumi.get(__ret__, 'id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        project_id_source=pulumi.get(__ret__, 'project_id_source'),
        region=pulumi.get(__ret__, 'region'),
        region_source=pulumi.get(__ret__, 'region_source'),
        secret_key=pulumi.get(__ret__, 'secret_key'),
        secret_key_source=pulumi.get(__ret__, 'secret_key_source'),
        zone=pulumi.get(__ret__, 'zone'),
        zone_source=pulumi.get(__ret__, 'zone_source'))
def get_config_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetConfigResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getConfig:getConfig', __args__, opts=opts, typ=GetConfigResult)
    return __ret__.apply(lambda __response__: GetConfigResult(
        access_key=pulumi.get(__response__, 'access_key'),
        access_key_source=pulumi.get(__response__, 'access_key_source'),
        id=pulumi.get(__response__, 'id'),
        project_id=pulumi.get(__response__, 'project_id'),
        project_id_source=pulumi.get(__response__, 'project_id_source'),
        region=pulumi.get(__response__, 'region'),
        region_source=pulumi.get(__response__, 'region_source'),
        secret_key=pulumi.get(__response__, 'secret_key'),
        secret_key_source=pulumi.get(__response__, 'secret_key_source'),
        zone=pulumi.get(__response__, 'zone'),
        zone_source=pulumi.get(__response__, 'zone_source')))
