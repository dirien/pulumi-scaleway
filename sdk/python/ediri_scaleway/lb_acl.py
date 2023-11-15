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

__all__ = ['LbAclArgs', 'LbAcl']

@pulumi.input_type
class LbAclArgs:
    def __init__(__self__, *,
                 action: pulumi.Input['LbAclActionArgs'],
                 frontend_id: pulumi.Input[str],
                 index: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input['LbAclMatchArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LbAcl resource.
        :param pulumi.Input['LbAclActionArgs'] action: Action to undertake when an ACL filter matches.
        :param pulumi.Input[str] frontend_id: The load-balancer Frontend ID to attach the ACL to.
        :param pulumi.Input[int] index: The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        :param pulumi.Input[str] description: The ACL description.
        :param pulumi.Input['LbAclMatchArgs'] match: The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "frontend_id", frontend_id)
        pulumi.set(__self__, "index", index)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if match is not None:
            pulumi.set(__self__, "match", match)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input['LbAclActionArgs']:
        """
        Action to undertake when an ACL filter matches.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input['LbAclActionArgs']):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> pulumi.Input[str]:
        """
        The load-balancer Frontend ID to attach the ACL to.
        """
        return pulumi.get(self, "frontend_id")

    @frontend_id.setter
    def frontend_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "frontend_id", value)

    @property
    @pulumi.getter
    def index(self) -> pulumi.Input[int]:
        """
        The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: pulumi.Input[int]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The ACL description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def match(self) -> Optional[pulumi.Input['LbAclMatchArgs']]:
        """
        The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: Optional[pulumi.Input['LbAclMatchArgs']]):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _LbAclState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input['LbAclActionArgs']] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 frontend_id: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 match: Optional[pulumi.Input['LbAclMatchArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LbAcl resources.
        :param pulumi.Input['LbAclActionArgs'] action: Action to undertake when an ACL filter matches.
        :param pulumi.Input[str] created_at: Date and time of ACL's creation (RFC 3339 format)
        :param pulumi.Input[str] description: The ACL description.
        :param pulumi.Input[str] frontend_id: The load-balancer Frontend ID to attach the ACL to.
        :param pulumi.Input[int] index: The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        :param pulumi.Input['LbAclMatchArgs'] match: The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] updated_at: Date and time of ACL's update (RFC 3339 format)
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if frontend_id is not None:
            pulumi.set(__self__, "frontend_id", frontend_id)
        if index is not None:
            pulumi.set(__self__, "index", index)
        if match is not None:
            pulumi.set(__self__, "match", match)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['LbAclActionArgs']]:
        """
        Action to undertake when an ACL filter matches.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['LbAclActionArgs']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of ACL's creation (RFC 3339 format)
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The ACL description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> Optional[pulumi.Input[str]]:
        """
        The load-balancer Frontend ID to attach the ACL to.
        """
        return pulumi.get(self, "frontend_id")

    @frontend_id.setter
    def frontend_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "frontend_id", value)

    @property
    @pulumi.getter
    def index(self) -> Optional[pulumi.Input[int]]:
        """
        The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        """
        return pulumi.get(self, "index")

    @index.setter
    def index(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "index", value)

    @property
    @pulumi.getter
    def match(self) -> Optional[pulumi.Input['LbAclMatchArgs']]:
        """
        The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: Optional[pulumi.Input['LbAclMatchArgs']]):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time of ACL's update (RFC 3339 format)
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class LbAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[pulumi.InputType['LbAclActionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 frontend_id: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 match: Optional[pulumi.Input[pulumi.InputType['LbAclMatchArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Load-Balancer ACLs. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls).

        ## Examples Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        acl01 = scaleway.LbAcl("acl01",
            frontend_id=scaleway_lb_frontend["frt01"]["id"],
            description="Exclude well-known IPs",
            index=0,
            action=scaleway.LbAclActionArgs(
                type="allow",
            ),
            match=scaleway.LbAclMatchArgs(
                ip_subnets=[
                    "192.168.0.1",
                    "192.168.0.2",
                    "192.168.10.0/24",
                ],
            ))
        ```

        ## Import

        Load-Balancer ACL can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/lbAcl:LbAcl acl01 fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LbAclActionArgs']] action: Action to undertake when an ACL filter matches.
        :param pulumi.Input[str] description: The ACL description.
        :param pulumi.Input[str] frontend_id: The load-balancer Frontend ID to attach the ACL to.
        :param pulumi.Input[int] index: The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        :param pulumi.Input[pulumi.InputType['LbAclMatchArgs']] match: The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Load-Balancer ACLs. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-acls).

        ## Examples Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        acl01 = scaleway.LbAcl("acl01",
            frontend_id=scaleway_lb_frontend["frt01"]["id"],
            description="Exclude well-known IPs",
            index=0,
            action=scaleway.LbAclActionArgs(
                type="allow",
            ),
            match=scaleway.LbAclMatchArgs(
                ip_subnets=[
                    "192.168.0.1",
                    "192.168.0.2",
                    "192.168.10.0/24",
                ],
            ))
        ```

        ## Import

        Load-Balancer ACL can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/lbAcl:LbAcl acl01 fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param LbAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[pulumi.InputType['LbAclActionArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 frontend_id: Optional[pulumi.Input[str]] = None,
                 index: Optional[pulumi.Input[int]] = None,
                 match: Optional[pulumi.Input[pulumi.InputType['LbAclMatchArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbAclArgs.__new__(LbAclArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["description"] = description
            if frontend_id is None and not opts.urn:
                raise TypeError("Missing required property 'frontend_id'")
            __props__.__dict__["frontend_id"] = frontend_id
            if index is None and not opts.urn:
                raise TypeError("Missing required property 'index'")
            __props__.__dict__["index"] = index
            __props__.__dict__["match"] = match
            __props__.__dict__["name"] = name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(LbAcl, __self__).__init__(
            'scaleway:index/lbAcl:LbAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[pulumi.InputType['LbAclActionArgs']]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            frontend_id: Optional[pulumi.Input[str]] = None,
            index: Optional[pulumi.Input[int]] = None,
            match: Optional[pulumi.Input[pulumi.InputType['LbAclMatchArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'LbAcl':
        """
        Get an existing LbAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LbAclActionArgs']] action: Action to undertake when an ACL filter matches.
        :param pulumi.Input[str] created_at: Date and time of ACL's creation (RFC 3339 format)
        :param pulumi.Input[str] description: The ACL description.
        :param pulumi.Input[str] frontend_id: The load-balancer Frontend ID to attach the ACL to.
        :param pulumi.Input[int] index: The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        :param pulumi.Input[pulumi.InputType['LbAclMatchArgs']] match: The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] updated_at: Date and time of ACL's update (RFC 3339 format)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbAclState.__new__(_LbAclState)

        __props__.__dict__["action"] = action
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["frontend_id"] = frontend_id
        __props__.__dict__["index"] = index
        __props__.__dict__["match"] = match
        __props__.__dict__["name"] = name
        __props__.__dict__["updated_at"] = updated_at
        return LbAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output['outputs.LbAclAction']:
        """
        Action to undertake when an ACL filter matches.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        Date and time of ACL's creation (RFC 3339 format)
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The ACL description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> pulumi.Output[str]:
        """
        The load-balancer Frontend ID to attach the ACL to.
        """
        return pulumi.get(self, "frontend_id")

    @property
    @pulumi.getter
    def index(self) -> pulumi.Output[int]:
        """
        The Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed).
        """
        return pulumi.get(self, "index")

    @property
    @pulumi.getter
    def match(self) -> pulumi.Output[Optional['outputs.LbAclMatch']]:
        """
        The ACL match rule. At least `ip_subnet` or `http_filter` and `http_filter_value` are required.
        """
        return pulumi.get(self, "match")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Date and time of ACL's update (RFC 3339 format)
        """
        return pulumi.get(self, "updated_at")

