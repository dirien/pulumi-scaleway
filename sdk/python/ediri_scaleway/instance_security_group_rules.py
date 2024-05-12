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

__all__ = ['InstanceSecurityGroupRulesArgs', 'InstanceSecurityGroupRules']

@pulumi.input_type
class InstanceSecurityGroupRulesArgs:
    def __init__(__self__, *,
                 security_group_id: pulumi.Input[str],
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]]] = None):
        """
        The set of arguments for constructing a InstanceSecurityGroupRules resource.
        :param pulumi.Input[str] security_group_id: The ID of the security group.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        pulumi.set(__self__, "security_group_id", security_group_id)
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the security group.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]]]:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @inbound_rules.setter
    def inbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]]]):
        pulumi.set(self, "inbound_rules", value)

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]]]:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @outbound_rules.setter
    def outbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]]]):
        pulumi.set(self, "outbound_rules", value)


@pulumi.input_type
class _InstanceSecurityGroupRulesState:
    def __init__(__self__, *,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceSecurityGroupRules resources.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] security_group_id: The ID of the security group.
        """
        if inbound_rules is not None:
            pulumi.set(__self__, "inbound_rules", inbound_rules)
        if outbound_rules is not None:
            pulumi.set(__self__, "outbound_rules", outbound_rules)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]]]:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @inbound_rules.setter
    def inbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesInboundRuleArgs']]]]):
        pulumi.set(self, "inbound_rules", value)

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]]]:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @outbound_rules.setter
    def outbound_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceSecurityGroupRulesOutboundRuleArgs']]]]):
        pulumi.set(self, "outbound_rules", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the security group.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)


class InstanceSecurityGroupRules(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesInboundRuleArgs']]]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesOutboundRuleArgs']]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Compute Instance security group rules. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/instance/#security-groups-8d7f89).

        This resource can be used to externalize rules from a `InstanceSecurityGroup` to solve circular dependency problems. When using this resource do not forget to set `external_rules = true` on the security group.

        > **Warning:** In order to guaranty rules order in a given security group only one InstanceSecurityGroupRules is allowed per security group.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        sg01 = scaleway.InstanceSecurityGroup("sg01", external_rules=True)
        sgrs01 = scaleway.InstanceSecurityGroupRules("sgrs01",
            security_group_id=sg01.id,
            inbound_rules=[scaleway.InstanceSecurityGroupRulesInboundRuleArgs(
                action="accept",
                port=80,
                ip_range="0.0.0.0/0",
            )])
        ```

        ## Import

        Instance security group rules can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules web fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesInboundRuleArgs']]]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesOutboundRuleArgs']]]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] security_group_id: The ID of the security group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceSecurityGroupRulesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Compute Instance security group rules. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/instance/#security-groups-8d7f89).

        This resource can be used to externalize rules from a `InstanceSecurityGroup` to solve circular dependency problems. When using this resource do not forget to set `external_rules = true` on the security group.

        > **Warning:** In order to guaranty rules order in a given security group only one InstanceSecurityGroupRules is allowed per security group.

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        sg01 = scaleway.InstanceSecurityGroup("sg01", external_rules=True)
        sgrs01 = scaleway.InstanceSecurityGroupRules("sgrs01",
            security_group_id=sg01.id,
            inbound_rules=[scaleway.InstanceSecurityGroupRulesInboundRuleArgs(
                action="accept",
                port=80,
                ip_range="0.0.0.0/0",
            )])
        ```

        ## Import

        Instance security group rules can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules web fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param InstanceSecurityGroupRulesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceSecurityGroupRulesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesInboundRuleArgs']]]]] = None,
                 outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesOutboundRuleArgs']]]]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceSecurityGroupRulesArgs.__new__(InstanceSecurityGroupRulesArgs)

            __props__.__dict__["inbound_rules"] = inbound_rules
            __props__.__dict__["outbound_rules"] = outbound_rules
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
        super(InstanceSecurityGroupRules, __self__).__init__(
            'scaleway:index/instanceSecurityGroupRules:InstanceSecurityGroupRules',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            inbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesInboundRuleArgs']]]]] = None,
            outbound_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesOutboundRuleArgs']]]]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None) -> 'InstanceSecurityGroupRules':
        """
        Get an existing InstanceSecurityGroupRules resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesInboundRuleArgs']]]] inbound_rules: A list of inbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceSecurityGroupRulesOutboundRuleArgs']]]] outbound_rules: A list of outbound rule to add to the security group. (Structure is documented below.)
        :param pulumi.Input[str] security_group_id: The ID of the security group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceSecurityGroupRulesState.__new__(_InstanceSecurityGroupRulesState)

        __props__.__dict__["inbound_rules"] = inbound_rules
        __props__.__dict__["outbound_rules"] = outbound_rules
        __props__.__dict__["security_group_id"] = security_group_id
        return InstanceSecurityGroupRules(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceSecurityGroupRulesInboundRule']]]:
        """
        A list of inbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> pulumi.Output[Optional[Sequence['outputs.InstanceSecurityGroupRulesOutboundRule']]]:
        """
        A list of outbound rule to add to the security group. (Structure is documented below.)
        """
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the security group.
        """
        return pulumi.get(self, "security_group_id")

