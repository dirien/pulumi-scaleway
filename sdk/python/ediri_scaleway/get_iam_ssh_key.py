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
    'GetIamSshKeyResult',
    'AwaitableGetIamSshKeyResult',
    'get_iam_ssh_key',
    'get_iam_ssh_key_output',
]

@pulumi.output_type
class GetIamSshKeyResult:
    """
    A collection of values returned by getIamSshKey.
    """
    def __init__(__self__, created_at=None, disabled=None, fingerprint=None, id=None, name=None, organization_id=None, project_id=None, public_key=None, ssh_key_id=None, updated_at=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if public_key and not isinstance(public_key, str):
            raise TypeError("Expected argument 'public_key' to be a str")
        pulumi.set(__self__, "public_key", public_key)
        if ssh_key_id and not isinstance(ssh_key_id, str):
            raise TypeError("Expected argument 'ssh_key_id' to be a str")
        pulumi.set(__self__, "ssh_key_id", ssh_key_id)
        if updated_at and not isinstance(updated_at, str):
            raise TypeError("Expected argument 'updated_at' to be a str")
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        The date and time of the creation of the SSH key.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        The SSH key status.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        return pulumi.get(self, "fingerprint")

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
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The ID of the organization the SSH key is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> str:
        """
        The SSH public key string
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="sshKeyId")
    def ssh_key_id(self) -> Optional[str]:
        return pulumi.get(self, "ssh_key_id")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The date and time of the last update of the SSH key.
        """
        return pulumi.get(self, "updated_at")


class AwaitableGetIamSshKeyResult(GetIamSshKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIamSshKeyResult(
            created_at=self.created_at,
            disabled=self.disabled,
            fingerprint=self.fingerprint,
            id=self.id,
            name=self.name,
            organization_id=self.organization_id,
            project_id=self.project_id,
            public_key=self.public_key,
            ssh_key_id=self.ssh_key_id,
            updated_at=self.updated_at)


def get_iam_ssh_key(name: Optional[str] = None,
                    project_id: Optional[str] = None,
                    ssh_key_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIamSshKeyResult:
    """
    Use this data source to get SSH key information based on its ID or name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_iam_ssh_key(ssh_key_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The SSH key name.
    :param str project_id: `project_id`) The ID of the project the SSH
           key is associated with.
    :param str ssh_key_id: The SSH key id.
           
           > **Note** You must specify at least one: `name` and/or `ssh_key_id`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['sshKeyId'] = ssh_key_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getIamSshKey:getIamSshKey', __args__, opts=opts, typ=GetIamSshKeyResult).value

    return AwaitableGetIamSshKeyResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        disabled=pulumi.get(__ret__, 'disabled'),
        fingerprint=pulumi.get(__ret__, 'fingerprint'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        organization_id=pulumi.get(__ret__, 'organization_id'),
        project_id=pulumi.get(__ret__, 'project_id'),
        public_key=pulumi.get(__ret__, 'public_key'),
        ssh_key_id=pulumi.get(__ret__, 'ssh_key_id'),
        updated_at=pulumi.get(__ret__, 'updated_at'))
def get_iam_ssh_key_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                           project_id: Optional[pulumi.Input[Optional[str]]] = None,
                           ssh_key_id: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIamSshKeyResult]:
    """
    Use this data source to get SSH key information based on its ID or name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_iam_ssh_key(ssh_key_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The SSH key name.
    :param str project_id: `project_id`) The ID of the project the SSH
           key is associated with.
    :param str ssh_key_id: The SSH key id.
           
           > **Note** You must specify at least one: `name` and/or `ssh_key_id`.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['projectId'] = project_id
    __args__['sshKeyId'] = ssh_key_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('scaleway:index/getIamSshKey:getIamSshKey', __args__, opts=opts, typ=GetIamSshKeyResult)
    return __ret__.apply(lambda __response__: GetIamSshKeyResult(
        created_at=pulumi.get(__response__, 'created_at'),
        disabled=pulumi.get(__response__, 'disabled'),
        fingerprint=pulumi.get(__response__, 'fingerprint'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        organization_id=pulumi.get(__response__, 'organization_id'),
        project_id=pulumi.get(__response__, 'project_id'),
        public_key=pulumi.get(__response__, 'public_key'),
        ssh_key_id=pulumi.get(__response__, 'ssh_key_id'),
        updated_at=pulumi.get(__response__, 'updated_at')))
