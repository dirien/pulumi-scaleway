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

__all__ = [
    'GetInstanceSecurityGroupResult',
    'AwaitableGetInstanceSecurityGroupResult',
    'get_instance_security_group',
    'get_instance_security_group_output',
]

@pulumi.output_type
class GetInstanceSecurityGroupResult:
    """
    A collection of values returned by getInstanceSecurityGroup.
    """
    def __init__(__self__, description=None, enable_default_security=None, external_rules=None, id=None, inbound_default_policy=None, inbound_rules=None, name=None, organization_id=None, outbound_default_policy=None, outbound_rules=None, project_id=None, security_group_id=None, stateful=None, tags=None, zone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_default_security and not isinstance(enable_default_security, bool):
            raise TypeError("Expected argument 'enable_default_security' to be a bool")
        pulumi.set(__self__, "enable_default_security", enable_default_security)
        if external_rules and not isinstance(external_rules, bool):
            raise TypeError("Expected argument 'external_rules' to be a bool")
        pulumi.set(__self__, "external_rules", external_rules)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inbound_default_policy and not isinstance(inbound_default_policy, str):
            raise TypeError("Expected argument 'inbound_default_policy' to be a str")
        pulumi.set(__self__, "inbound_default_policy", inbound_default_policy)
        if inbound_rules and not isinstance(inbound_rules, list):
            raise TypeError("Expected argument 'inbound_rules' to be a list")
        pulumi.set(__self__, "inbound_rules", inbound_rules)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if organization_id and not isinstance(organization_id, str):
            raise TypeError("Expected argument 'organization_id' to be a str")
        pulumi.set(__self__, "organization_id", organization_id)
        if outbound_default_policy and not isinstance(outbound_default_policy, str):
            raise TypeError("Expected argument 'outbound_default_policy' to be a str")
        pulumi.set(__self__, "outbound_default_policy", outbound_default_policy)
        if outbound_rules and not isinstance(outbound_rules, list):
            raise TypeError("Expected argument 'outbound_rules' to be a list")
        pulumi.set(__self__, "outbound_rules", outbound_rules)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if security_group_id and not isinstance(security_group_id, str):
            raise TypeError("Expected argument 'security_group_id' to be a str")
        pulumi.set(__self__, "security_group_id", security_group_id)
        if stateful and not isinstance(stateful, bool):
            raise TypeError("Expected argument 'stateful' to be a bool")
        pulumi.set(__self__, "stateful", stateful)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableDefaultSecurity")
    def enable_default_security(self) -> bool:
        return pulumi.get(self, "enable_default_security")

    @property
    @pulumi.getter(name="externalRules")
    def external_rules(self) -> bool:
        return pulumi.get(self, "external_rules")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inboundDefaultPolicy")
    def inbound_default_policy(self) -> str:
        """
        The default policy on incoming traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "inbound_default_policy")

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Sequence['outputs.GetInstanceSecurityGroupInboundRuleResult']:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> str:
        """
        The ID of the organization the security group is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="outboundDefaultPolicy")
    def outbound_default_policy(self) -> str:
        """
        The default policy on outgoing traffic. Possible values are: `accept` or `drop`.
        """
        return pulumi.get(self, "outbound_default_policy")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Sequence['outputs.GetInstanceSecurityGroupOutboundRuleResult']:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The ID of the project the security group is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[str]:
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter
    def stateful(self) -> bool:
        return pulumi.get(self, "stateful")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetInstanceSecurityGroupResult(GetInstanceSecurityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceSecurityGroupResult(
            description=self.description,
            enable_default_security=self.enable_default_security,
            external_rules=self.external_rules,
            id=self.id,
            inbound_default_policy=self.inbound_default_policy,
            inbound_rules=self.inbound_rules,
            name=self.name,
            organization_id=self.organization_id,
            outbound_default_policy=self.outbound_default_policy,
            outbound_rules=self.outbound_rules,
            project_id=self.project_id,
            security_group_id=self.security_group_id,
            stateful=self.stateful,
            tags=self.tags,
            zone=self.zone)


def get_instance_security_group(name: Optional[str] = None,
                                security_group_id: Optional[str] = None,
                                zone: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceSecurityGroupResult:
    """
    Gets information about a Security Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_instance_security_group(security_group_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The security group name. Only one of `name` and `security_group_id` should be specified.
    :param str security_group_id: The security group id. Only one of `name` and `security_group_id` should be specified.
    :param str zone: `zone`) The zone in which the security group exists.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['securityGroupId'] = security_group_id
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('scaleway:index/getInstanceSecurityGroup:getInstanceSecurityGroup', __args__, opts=opts, typ=GetInstanceSecurityGroupResult).value

    return AwaitableGetInstanceSecurityGroupResult(
        description=__ret__.description,
        enable_default_security=__ret__.enable_default_security,
        external_rules=__ret__.external_rules,
        id=__ret__.id,
        inbound_default_policy=__ret__.inbound_default_policy,
        inbound_rules=__ret__.inbound_rules,
        name=__ret__.name,
        organization_id=__ret__.organization_id,
        outbound_default_policy=__ret__.outbound_default_policy,
        outbound_rules=__ret__.outbound_rules,
        project_id=__ret__.project_id,
        security_group_id=__ret__.security_group_id,
        stateful=__ret__.stateful,
        tags=__ret__.tags,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_instance_security_group)
def get_instance_security_group_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                       security_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                       zone: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceSecurityGroupResult]:
    """
    Gets information about a Security Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_scaleway as scaleway

    my_key = scaleway.get_instance_security_group(security_group_id="11111111-1111-1111-1111-111111111111")
    ```


    :param str name: The security group name. Only one of `name` and `security_group_id` should be specified.
    :param str security_group_id: The security group id. Only one of `name` and `security_group_id` should be specified.
    :param str zone: `zone`) The zone in which the security group exists.
    """
    ...
