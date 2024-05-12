# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['VpcPublicGatewayArgs', 'VpcPublicGateway']

@pulumi.input_type
class VpcPublicGatewayArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 bastion_enabled: Optional[pulumi.Input[bool]] = None,
                 bastion_port: Optional[pulumi.Input[int]] = None,
                 enable_smtp: Optional[pulumi.Input[bool]] = None,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 upstream_dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcPublicGateway resource.
        :param pulumi.Input[str] type: The gateway type.
        :param pulumi.Input[bool] bastion_enabled: Enable SSH bastion on the gateway
        :param pulumi.Input[int] bastion_port: The port on which the SSH bastion will listen.
        :param pulumi.Input[bool] enable_smtp: Enable SMTP on the gateway
        :param pulumi.Input[str] ip_id: attach an existing flexible IP to the gateway
        :param pulumi.Input[str] name: The name of the public gateway. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the public gateway is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the public gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] upstream_dns_servers: override the gateway's default recursive DNS servers, if DNS features are enabled.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway should be created.
        """
        pulumi.set(__self__, "type", type)
        if bastion_enabled is not None:
            pulumi.set(__self__, "bastion_enabled", bastion_enabled)
        if bastion_port is not None:
            pulumi.set(__self__, "bastion_port", bastion_port)
        if enable_smtp is not None:
            pulumi.set(__self__, "enable_smtp", enable_smtp)
        if ip_id is not None:
            pulumi.set(__self__, "ip_id", ip_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if upstream_dns_servers is not None:
            pulumi.set(__self__, "upstream_dns_servers", upstream_dns_servers)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The gateway type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="bastionEnabled")
    def bastion_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SSH bastion on the gateway
        """
        return pulumi.get(self, "bastion_enabled")

    @bastion_enabled.setter
    def bastion_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bastion_enabled", value)

    @property
    @pulumi.getter(name="bastionPort")
    def bastion_port(self) -> Optional[pulumi.Input[int]]:
        """
        The port on which the SSH bastion will listen.
        """
        return pulumi.get(self, "bastion_port")

    @bastion_port.setter
    def bastion_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bastion_port", value)

    @property
    @pulumi.getter(name="enableSmtp")
    def enable_smtp(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SMTP on the gateway
        """
        return pulumi.get(self, "enable_smtp")

    @enable_smtp.setter
    def enable_smtp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_smtp", value)

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        attach an existing flexible IP to the gateway
        """
        return pulumi.get(self, "ip_id")

    @ip_id.setter
    def ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the public gateway. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the public gateway is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the public gateway.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="upstreamDnsServers")
    def upstream_dns_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        override the gateway's default recursive DNS servers, if DNS features are enabled.
        """
        return pulumi.get(self, "upstream_dns_servers")

    @upstream_dns_servers.setter
    def upstream_dns_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "upstream_dns_servers", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the public gateway should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _VpcPublicGatewayState:
    def __init__(__self__, *,
                 bastion_enabled: Optional[pulumi.Input[bool]] = None,
                 bastion_port: Optional[pulumi.Input[int]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 enable_smtp: Optional[pulumi.Input[bool]] = None,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 upstream_dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcPublicGateway resources.
        :param pulumi.Input[bool] bastion_enabled: Enable SSH bastion on the gateway
        :param pulumi.Input[int] bastion_port: The port on which the SSH bastion will listen.
        :param pulumi.Input[str] created_at: The date and time of the creation of the public gateway.
        :param pulumi.Input[bool] enable_smtp: Enable SMTP on the gateway
        :param pulumi.Input[str] ip_id: attach an existing flexible IP to the gateway
        :param pulumi.Input[str] name: The name of the public gateway. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the public gateway is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the public gateway is associated with.
        :param pulumi.Input[str] status: The status of the public gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the public gateway.
        :param pulumi.Input[str] type: The gateway type.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the public gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] upstream_dns_servers: override the gateway's default recursive DNS servers, if DNS features are enabled.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway should be created.
        """
        if bastion_enabled is not None:
            pulumi.set(__self__, "bastion_enabled", bastion_enabled)
        if bastion_port is not None:
            pulumi.set(__self__, "bastion_port", bastion_port)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if enable_smtp is not None:
            pulumi.set(__self__, "enable_smtp", enable_smtp)
        if ip_id is not None:
            pulumi.set(__self__, "ip_id", ip_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if upstream_dns_servers is not None:
            pulumi.set(__self__, "upstream_dns_servers", upstream_dns_servers)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="bastionEnabled")
    def bastion_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SSH bastion on the gateway
        """
        return pulumi.get(self, "bastion_enabled")

    @bastion_enabled.setter
    def bastion_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bastion_enabled", value)

    @property
    @pulumi.getter(name="bastionPort")
    def bastion_port(self) -> Optional[pulumi.Input[int]]:
        """
        The port on which the SSH bastion will listen.
        """
        return pulumi.get(self, "bastion_port")

    @bastion_port.setter
    def bastion_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bastion_port", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the creation of the public gateway.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="enableSmtp")
    def enable_smtp(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SMTP on the gateway
        """
        return pulumi.get(self, "enable_smtp")

    @enable_smtp.setter
    def enable_smtp(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_smtp", value)

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> Optional[pulumi.Input[str]]:
        """
        attach an existing flexible IP to the gateway
        """
        return pulumi.get(self, "ip_id")

    @ip_id.setter
    def ip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the public gateway. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization ID the public gateway is associated with.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the public gateway is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the public gateway.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags associated with the public gateway.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The gateway type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the last update of the public gateway.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="upstreamDnsServers")
    def upstream_dns_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        override the gateway's default recursive DNS servers, if DNS features are enabled.
        """
        return pulumi.get(self, "upstream_dns_servers")

    @upstream_dns_servers.setter
    def upstream_dns_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "upstream_dns_servers", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        `zone`) The zone in which the public gateway should be created.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class VpcPublicGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bastion_enabled: Optional[pulumi.Input[bool]] = None,
                 bastion_port: Optional[pulumi.Input[int]] = None,
                 enable_smtp: Optional[pulumi.Input[bool]] = None,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 upstream_dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway VPC Public Gateway.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/public-gateway).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.VpcPublicGateway("main",
            tags=[
                "demo",
                "terraform",
            ],
            type="VPC-GW-S")
        ```

        ## Import

        Public gateway can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/vpcPublicGateway:VpcPublicGateway main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bastion_enabled: Enable SSH bastion on the gateway
        :param pulumi.Input[int] bastion_port: The port on which the SSH bastion will listen.
        :param pulumi.Input[bool] enable_smtp: Enable SMTP on the gateway
        :param pulumi.Input[str] ip_id: attach an existing flexible IP to the gateway
        :param pulumi.Input[str] name: The name of the public gateway. If not provided it will be randomly generated.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the public gateway is associated with.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the public gateway.
        :param pulumi.Input[str] type: The gateway type.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] upstream_dns_servers: override the gateway's default recursive DNS servers, if DNS features are enabled.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcPublicGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway VPC Public Gateway.
        For more information, see [the documentation](https://www.scaleway.com/en/developers/api/public-gateway).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.VpcPublicGateway("main",
            tags=[
                "demo",
                "terraform",
            ],
            type="VPC-GW-S")
        ```

        ## Import

        Public gateway can be imported using the `{zone}/{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/vpcPublicGateway:VpcPublicGateway main fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param VpcPublicGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcPublicGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bastion_enabled: Optional[pulumi.Input[bool]] = None,
                 bastion_port: Optional[pulumi.Input[int]] = None,
                 enable_smtp: Optional[pulumi.Input[bool]] = None,
                 ip_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 upstream_dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcPublicGatewayArgs.__new__(VpcPublicGatewayArgs)

            __props__.__dict__["bastion_enabled"] = bastion_enabled
            __props__.__dict__["bastion_port"] = bastion_port
            __props__.__dict__["enable_smtp"] = enable_smtp
            __props__.__dict__["ip_id"] = ip_id
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["tags"] = tags
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["upstream_dns_servers"] = upstream_dns_servers
            __props__.__dict__["zone"] = zone
            __props__.__dict__["created_at"] = None
            __props__.__dict__["organization_id"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(VpcPublicGateway, __self__).__init__(
            'scaleway:index/vpcPublicGateway:VpcPublicGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bastion_enabled: Optional[pulumi.Input[bool]] = None,
            bastion_port: Optional[pulumi.Input[int]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            enable_smtp: Optional[pulumi.Input[bool]] = None,
            ip_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            upstream_dns_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'VpcPublicGateway':
        """
        Get an existing VpcPublicGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bastion_enabled: Enable SSH bastion on the gateway
        :param pulumi.Input[int] bastion_port: The port on which the SSH bastion will listen.
        :param pulumi.Input[str] created_at: The date and time of the creation of the public gateway.
        :param pulumi.Input[bool] enable_smtp: Enable SMTP on the gateway
        :param pulumi.Input[str] ip_id: attach an existing flexible IP to the gateway
        :param pulumi.Input[str] name: The name of the public gateway. If not provided it will be randomly generated.
        :param pulumi.Input[str] organization_id: The organization ID the public gateway is associated with.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the public gateway is associated with.
        :param pulumi.Input[str] status: The status of the public gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags associated with the public gateway.
        :param pulumi.Input[str] type: The gateway type.
        :param pulumi.Input[str] updated_at: The date and time of the last update of the public gateway.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] upstream_dns_servers: override the gateway's default recursive DNS servers, if DNS features are enabled.
        :param pulumi.Input[str] zone: `zone`) The zone in which the public gateway should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcPublicGatewayState.__new__(_VpcPublicGatewayState)

        __props__.__dict__["bastion_enabled"] = bastion_enabled
        __props__.__dict__["bastion_port"] = bastion_port
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["enable_smtp"] = enable_smtp
        __props__.__dict__["ip_id"] = ip_id
        __props__.__dict__["name"] = name
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["upstream_dns_servers"] = upstream_dns_servers
        __props__.__dict__["zone"] = zone
        return VpcPublicGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bastionEnabled")
    def bastion_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable SSH bastion on the gateway
        """
        return pulumi.get(self, "bastion_enabled")

    @property
    @pulumi.getter(name="bastionPort")
    def bastion_port(self) -> pulumi.Output[int]:
        """
        The port on which the SSH bastion will listen.
        """
        return pulumi.get(self, "bastion_port")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the creation of the public gateway.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="enableSmtp")
    def enable_smtp(self) -> pulumi.Output[bool]:
        """
        Enable SMTP on the gateway
        """
        return pulumi.get(self, "enable_smtp")

    @property
    @pulumi.getter(name="ipId")
    def ip_id(self) -> pulumi.Output[str]:
        """
        attach an existing flexible IP to the gateway
        """
        return pulumi.get(self, "ip_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the public gateway. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        """
        The organization ID the public gateway is associated with.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the public gateway is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the public gateway.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The tags associated with the public gateway.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The gateway type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The date and time of the last update of the public gateway.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="upstreamDnsServers")
    def upstream_dns_servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        override the gateway's default recursive DNS servers, if DNS features are enabled.
        """
        return pulumi.get(self, "upstream_dns_servers")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        `zone`) The zone in which the public gateway should be created.
        """
        return pulumi.get(self, "zone")

