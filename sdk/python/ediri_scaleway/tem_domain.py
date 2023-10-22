# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TemDomainArgs', 'TemDomain']

@pulumi.input_type
class TemDomainArgs:
    def __init__(__self__, *,
                 accept_tos: pulumi.Input[bool],
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TemDomain resource.
        :param pulumi.Input[bool] accept_tos: Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
               > **Important:**  This attribute must be set to `true`.
        :param pulumi.Input[str] name: The domain name, must not be used in another Transactional Email Domain.
               > **Important:** Updates to `name` will recreate the domain.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the domain is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the domain should be created.
        """
        TemDomainArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            accept_tos=accept_tos,
            name=name,
            project_id=project_id,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             accept_tos: pulumi.Input[bool],
             name: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'acceptTos' in kwargs:
            accept_tos = kwargs['acceptTos']
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']

        _setter("accept_tos", accept_tos)
        if name is not None:
            _setter("name", name)
        if project_id is not None:
            _setter("project_id", project_id)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="acceptTos")
    def accept_tos(self) -> pulumi.Input[bool]:
        """
        Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        > **Important:**  This attribute must be set to `true`.
        """
        return pulumi.get(self, "accept_tos")

    @accept_tos.setter
    def accept_tos(self, value: pulumi.Input[bool]):
        pulumi.set(self, "accept_tos", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name, must not be used in another Transactional Email Domain.
        > **Important:** Updates to `name` will recreate the domain.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the domain is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the domain should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _TemDomainState:
    def __init__(__self__, *,
                 accept_tos: Optional[pulumi.Input[bool]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 dkim_config: Optional[pulumi.Input[str]] = None,
                 last_error: Optional[pulumi.Input[str]] = None,
                 last_valid_at: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 next_check_at: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 revoked_at: Optional[pulumi.Input[str]] = None,
                 smtp_host: Optional[pulumi.Input[str]] = None,
                 smtp_port: Optional[pulumi.Input[int]] = None,
                 smtp_port_alternative: Optional[pulumi.Input[int]] = None,
                 smtp_port_unsecure: Optional[pulumi.Input[int]] = None,
                 smtps_port: Optional[pulumi.Input[int]] = None,
                 smtps_port_alternative: Optional[pulumi.Input[int]] = None,
                 spf_config: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TemDomain resources.
        :param pulumi.Input[bool] accept_tos: Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
               > **Important:**  This attribute must be set to `true`.
        :param pulumi.Input[str] created_at: The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        :param pulumi.Input[str] dkim_config: The DKIM public key, as should be recorded in the DNS zone.
        :param pulumi.Input[str] last_error: The error message if the last check failed.
        :param pulumi.Input[str] last_valid_at: The date and time the domain was last found to be valid (RFC 3339 format).
        :param pulumi.Input[str] name: The domain name, must not be used in another Transactional Email Domain.
               > **Important:** Updates to `name` will recreate the domain.
        :param pulumi.Input[str] next_check_at: The date and time of the next scheduled check (RFC 3339 format).
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the domain is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the domain should be created.
        :param pulumi.Input[str] revoked_at: The date and time of the revocation of the domain (RFC 3339 format).
        :param pulumi.Input[str] smtp_host: SMTP host to use to send emails
        :param pulumi.Input[int] smtp_port: SMTP port to use to send emails over TLS. (Port 587)
        :param pulumi.Input[int] smtp_port_alternative: SMTP port to use to send emails over TLS. (Port 2587)
        :param pulumi.Input[int] smtp_port_unsecure: SMTP port to use to send emails. (Port 25)
        :param pulumi.Input[int] smtps_port: SMTPS port to use to send emails over TLS Wrapper. (Port 465)
        :param pulumi.Input[int] smtps_port_alternative: SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
        :param pulumi.Input[str] spf_config: The snippet of the SPF record that should be registered in the DNS zone.
        :param pulumi.Input[str] status: The status of the Transaction Email Domain.
        """
        _TemDomainState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            accept_tos=accept_tos,
            created_at=created_at,
            dkim_config=dkim_config,
            last_error=last_error,
            last_valid_at=last_valid_at,
            name=name,
            next_check_at=next_check_at,
            project_id=project_id,
            region=region,
            revoked_at=revoked_at,
            smtp_host=smtp_host,
            smtp_port=smtp_port,
            smtp_port_alternative=smtp_port_alternative,
            smtp_port_unsecure=smtp_port_unsecure,
            smtps_port=smtps_port,
            smtps_port_alternative=smtps_port_alternative,
            spf_config=spf_config,
            status=status,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             accept_tos: Optional[pulumi.Input[bool]] = None,
             created_at: Optional[pulumi.Input[str]] = None,
             dkim_config: Optional[pulumi.Input[str]] = None,
             last_error: Optional[pulumi.Input[str]] = None,
             last_valid_at: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             next_check_at: Optional[pulumi.Input[str]] = None,
             project_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             revoked_at: Optional[pulumi.Input[str]] = None,
             smtp_host: Optional[pulumi.Input[str]] = None,
             smtp_port: Optional[pulumi.Input[int]] = None,
             smtp_port_alternative: Optional[pulumi.Input[int]] = None,
             smtp_port_unsecure: Optional[pulumi.Input[int]] = None,
             smtps_port: Optional[pulumi.Input[int]] = None,
             smtps_port_alternative: Optional[pulumi.Input[int]] = None,
             spf_config: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'acceptTos' in kwargs:
            accept_tos = kwargs['acceptTos']
        if 'createdAt' in kwargs:
            created_at = kwargs['createdAt']
        if 'dkimConfig' in kwargs:
            dkim_config = kwargs['dkimConfig']
        if 'lastError' in kwargs:
            last_error = kwargs['lastError']
        if 'lastValidAt' in kwargs:
            last_valid_at = kwargs['lastValidAt']
        if 'nextCheckAt' in kwargs:
            next_check_at = kwargs['nextCheckAt']
        if 'projectId' in kwargs:
            project_id = kwargs['projectId']
        if 'revokedAt' in kwargs:
            revoked_at = kwargs['revokedAt']
        if 'smtpHost' in kwargs:
            smtp_host = kwargs['smtpHost']
        if 'smtpPort' in kwargs:
            smtp_port = kwargs['smtpPort']
        if 'smtpPortAlternative' in kwargs:
            smtp_port_alternative = kwargs['smtpPortAlternative']
        if 'smtpPortUnsecure' in kwargs:
            smtp_port_unsecure = kwargs['smtpPortUnsecure']
        if 'smtpsPort' in kwargs:
            smtps_port = kwargs['smtpsPort']
        if 'smtpsPortAlternative' in kwargs:
            smtps_port_alternative = kwargs['smtpsPortAlternative']
        if 'spfConfig' in kwargs:
            spf_config = kwargs['spfConfig']

        if accept_tos is not None:
            _setter("accept_tos", accept_tos)
        if created_at is not None:
            _setter("created_at", created_at)
        if dkim_config is not None:
            _setter("dkim_config", dkim_config)
        if last_error is not None:
            _setter("last_error", last_error)
        if last_valid_at is not None:
            _setter("last_valid_at", last_valid_at)
        if name is not None:
            _setter("name", name)
        if next_check_at is not None:
            _setter("next_check_at", next_check_at)
        if project_id is not None:
            _setter("project_id", project_id)
        if region is not None:
            _setter("region", region)
        if revoked_at is not None:
            _setter("revoked_at", revoked_at)
        if smtp_host is not None:
            _setter("smtp_host", smtp_host)
        if smtp_port is not None:
            _setter("smtp_port", smtp_port)
        if smtp_port_alternative is not None:
            _setter("smtp_port_alternative", smtp_port_alternative)
        if smtp_port_unsecure is not None:
            _setter("smtp_port_unsecure", smtp_port_unsecure)
        if smtps_port is not None:
            _setter("smtps_port", smtps_port)
        if smtps_port_alternative is not None:
            _setter("smtps_port_alternative", smtps_port_alternative)
        if spf_config is not None:
            _setter("spf_config", spf_config)
        if status is not None:
            _setter("status", status)

    @property
    @pulumi.getter(name="acceptTos")
    def accept_tos(self) -> Optional[pulumi.Input[bool]]:
        """
        Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        > **Important:**  This attribute must be set to `true`.
        """
        return pulumi.get(self, "accept_tos")

    @accept_tos.setter
    def accept_tos(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "accept_tos", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="dkimConfig")
    def dkim_config(self) -> Optional[pulumi.Input[str]]:
        """
        The DKIM public key, as should be recorded in the DNS zone.
        """
        return pulumi.get(self, "dkim_config")

    @dkim_config.setter
    def dkim_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dkim_config", value)

    @property
    @pulumi.getter(name="lastError")
    def last_error(self) -> Optional[pulumi.Input[str]]:
        """
        The error message if the last check failed.
        """
        return pulumi.get(self, "last_error")

    @last_error.setter
    def last_error(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_error", value)

    @property
    @pulumi.getter(name="lastValidAt")
    def last_valid_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time the domain was last found to be valid (RFC 3339 format).
        """
        return pulumi.get(self, "last_valid_at")

    @last_valid_at.setter
    def last_valid_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_valid_at", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The domain name, must not be used in another Transactional Email Domain.
        > **Important:** Updates to `name` will recreate the domain.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nextCheckAt")
    def next_check_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the next scheduled check (RFC 3339 format).
        """
        return pulumi.get(self, "next_check_at")

    @next_check_at.setter
    def next_check_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_check_at", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the domain is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the domain should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="revokedAt")
    def revoked_at(self) -> Optional[pulumi.Input[str]]:
        """
        The date and time of the revocation of the domain (RFC 3339 format).
        """
        return pulumi.get(self, "revoked_at")

    @revoked_at.setter
    def revoked_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "revoked_at", value)

    @property
    @pulumi.getter(name="smtpHost")
    def smtp_host(self) -> Optional[pulumi.Input[str]]:
        """
        SMTP host to use to send emails
        """
        return pulumi.get(self, "smtp_host")

    @smtp_host.setter
    def smtp_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "smtp_host", value)

    @property
    @pulumi.getter(name="smtpPort")
    def smtp_port(self) -> Optional[pulumi.Input[int]]:
        """
        SMTP port to use to send emails over TLS. (Port 587)
        """
        return pulumi.get(self, "smtp_port")

    @smtp_port.setter
    def smtp_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "smtp_port", value)

    @property
    @pulumi.getter(name="smtpPortAlternative")
    def smtp_port_alternative(self) -> Optional[pulumi.Input[int]]:
        """
        SMTP port to use to send emails over TLS. (Port 2587)
        """
        return pulumi.get(self, "smtp_port_alternative")

    @smtp_port_alternative.setter
    def smtp_port_alternative(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "smtp_port_alternative", value)

    @property
    @pulumi.getter(name="smtpPortUnsecure")
    def smtp_port_unsecure(self) -> Optional[pulumi.Input[int]]:
        """
        SMTP port to use to send emails. (Port 25)
        """
        return pulumi.get(self, "smtp_port_unsecure")

    @smtp_port_unsecure.setter
    def smtp_port_unsecure(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "smtp_port_unsecure", value)

    @property
    @pulumi.getter(name="smtpsPort")
    def smtps_port(self) -> Optional[pulumi.Input[int]]:
        """
        SMTPS port to use to send emails over TLS Wrapper. (Port 465)
        """
        return pulumi.get(self, "smtps_port")

    @smtps_port.setter
    def smtps_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "smtps_port", value)

    @property
    @pulumi.getter(name="smtpsPortAlternative")
    def smtps_port_alternative(self) -> Optional[pulumi.Input[int]]:
        """
        SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
        """
        return pulumi.get(self, "smtps_port_alternative")

    @smtps_port_alternative.setter
    def smtps_port_alternative(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "smtps_port_alternative", value)

    @property
    @pulumi.getter(name="spfConfig")
    def spf_config(self) -> Optional[pulumi.Input[str]]:
        """
        The snippet of the SPF record that should be registered in the DNS zone.
        """
        return pulumi.get(self, "spf_config")

    @spf_config.setter
    def spf_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spf_config", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the Transaction Email Domain.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class TemDomain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_tos: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Transactional Email Domains.
        For more information see [the documentation](https://developers.scaleway.com/en/products/transactional_email/api/).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.TemDomain("main", accept_tos=True)
        ```

        ### Add the required records to your DNS zone

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        config = pulumi.Config()
        domain_name = config.require("domainName")
        main = scaleway.TemDomain("main", accept_tos=True)
        spf = scaleway.DomainRecord("spf",
            dns_zone=domain_name,
            type="TXT",
            data=main.spf_config.apply(lambda spf_config: f"v=spf1 {spf_config} -all"))
        dkim = scaleway.DomainRecord("dkim",
            dns_zone=domain_name,
            type="TXT",
            data=main.dkim_config)
        mx = scaleway.DomainRecord("mx",
            dns_zone=domain_name,
            type="MX",
            data=".")
        ```

        ## Import

        Domains can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/temDomain:TemDomain main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accept_tos: Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
               > **Important:**  This attribute must be set to `true`.
        :param pulumi.Input[str] name: The domain name, must not be used in another Transactional Email Domain.
               > **Important:** Updates to `name` will recreate the domain.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the domain is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the domain should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemDomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Transactional Email Domains.
        For more information see [the documentation](https://developers.scaleway.com/en/products/transactional_email/api/).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.TemDomain("main", accept_tos=True)
        ```

        ### Add the required records to your DNS zone

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        config = pulumi.Config()
        domain_name = config.require("domainName")
        main = scaleway.TemDomain("main", accept_tos=True)
        spf = scaleway.DomainRecord("spf",
            dns_zone=domain_name,
            type="TXT",
            data=main.spf_config.apply(lambda spf_config: f"v=spf1 {spf_config} -all"))
        dkim = scaleway.DomainRecord("dkim",
            dns_zone=domain_name,
            type="TXT",
            data=main.dkim_config)
        mx = scaleway.DomainRecord("mx",
            dns_zone=domain_name,
            type="MX",
            data=".")
        ```

        ## Import

        Domains can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/temDomain:TemDomain main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param TemDomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemDomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TemDomainArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accept_tos: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemDomainArgs.__new__(TemDomainArgs)

            if accept_tos is None and not opts.urn:
                raise TypeError("Missing required property 'accept_tos'")
            __props__.__dict__["accept_tos"] = accept_tos
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["created_at"] = None
            __props__.__dict__["dkim_config"] = None
            __props__.__dict__["last_error"] = None
            __props__.__dict__["last_valid_at"] = None
            __props__.__dict__["next_check_at"] = None
            __props__.__dict__["revoked_at"] = None
            __props__.__dict__["smtp_host"] = None
            __props__.__dict__["smtp_port"] = None
            __props__.__dict__["smtp_port_alternative"] = None
            __props__.__dict__["smtp_port_unsecure"] = None
            __props__.__dict__["smtps_port"] = None
            __props__.__dict__["smtps_port_alternative"] = None
            __props__.__dict__["spf_config"] = None
            __props__.__dict__["status"] = None
        super(TemDomain, __self__).__init__(
            'scaleway:index/temDomain:TemDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accept_tos: Optional[pulumi.Input[bool]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            dkim_config: Optional[pulumi.Input[str]] = None,
            last_error: Optional[pulumi.Input[str]] = None,
            last_valid_at: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            next_check_at: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            revoked_at: Optional[pulumi.Input[str]] = None,
            smtp_host: Optional[pulumi.Input[str]] = None,
            smtp_port: Optional[pulumi.Input[int]] = None,
            smtp_port_alternative: Optional[pulumi.Input[int]] = None,
            smtp_port_unsecure: Optional[pulumi.Input[int]] = None,
            smtps_port: Optional[pulumi.Input[int]] = None,
            smtps_port_alternative: Optional[pulumi.Input[int]] = None,
            spf_config: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'TemDomain':
        """
        Get an existing TemDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accept_tos: Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
               > **Important:**  This attribute must be set to `true`.
        :param pulumi.Input[str] created_at: The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        :param pulumi.Input[str] dkim_config: The DKIM public key, as should be recorded in the DNS zone.
        :param pulumi.Input[str] last_error: The error message if the last check failed.
        :param pulumi.Input[str] last_valid_at: The date and time the domain was last found to be valid (RFC 3339 format).
        :param pulumi.Input[str] name: The domain name, must not be used in another Transactional Email Domain.
               > **Important:** Updates to `name` will recreate the domain.
        :param pulumi.Input[str] next_check_at: The date and time of the next scheduled check (RFC 3339 format).
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the domain is associated with.
        :param pulumi.Input[str] region: `region`). The region in which the domain should be created.
        :param pulumi.Input[str] revoked_at: The date and time of the revocation of the domain (RFC 3339 format).
        :param pulumi.Input[str] smtp_host: SMTP host to use to send emails
        :param pulumi.Input[int] smtp_port: SMTP port to use to send emails over TLS. (Port 587)
        :param pulumi.Input[int] smtp_port_alternative: SMTP port to use to send emails over TLS. (Port 2587)
        :param pulumi.Input[int] smtp_port_unsecure: SMTP port to use to send emails. (Port 25)
        :param pulumi.Input[int] smtps_port: SMTPS port to use to send emails over TLS Wrapper. (Port 465)
        :param pulumi.Input[int] smtps_port_alternative: SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
        :param pulumi.Input[str] spf_config: The snippet of the SPF record that should be registered in the DNS zone.
        :param pulumi.Input[str] status: The status of the Transaction Email Domain.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemDomainState.__new__(_TemDomainState)

        __props__.__dict__["accept_tos"] = accept_tos
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["dkim_config"] = dkim_config
        __props__.__dict__["last_error"] = last_error
        __props__.__dict__["last_valid_at"] = last_valid_at
        __props__.__dict__["name"] = name
        __props__.__dict__["next_check_at"] = next_check_at
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["revoked_at"] = revoked_at
        __props__.__dict__["smtp_host"] = smtp_host
        __props__.__dict__["smtp_port"] = smtp_port
        __props__.__dict__["smtp_port_alternative"] = smtp_port_alternative
        __props__.__dict__["smtp_port_unsecure"] = smtp_port_unsecure
        __props__.__dict__["smtps_port"] = smtps_port
        __props__.__dict__["smtps_port_alternative"] = smtps_port_alternative
        __props__.__dict__["spf_config"] = spf_config
        __props__.__dict__["status"] = status
        return TemDomain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acceptTos")
    def accept_tos(self) -> pulumi.Output[bool]:
        """
        Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
        > **Important:**  This attribute must be set to `true`.
        """
        return pulumi.get(self, "accept_tos")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The date and time of the Transaction Email Domain's creation (RFC 3339 format).
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dkimConfig")
    def dkim_config(self) -> pulumi.Output[str]:
        """
        The DKIM public key, as should be recorded in the DNS zone.
        """
        return pulumi.get(self, "dkim_config")

    @property
    @pulumi.getter(name="lastError")
    def last_error(self) -> pulumi.Output[str]:
        """
        The error message if the last check failed.
        """
        return pulumi.get(self, "last_error")

    @property
    @pulumi.getter(name="lastValidAt")
    def last_valid_at(self) -> pulumi.Output[str]:
        """
        The date and time the domain was last found to be valid (RFC 3339 format).
        """
        return pulumi.get(self, "last_valid_at")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The domain name, must not be used in another Transactional Email Domain.
        > **Important:** Updates to `name` will recreate the domain.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextCheckAt")
    def next_check_at(self) -> pulumi.Output[str]:
        """
        The date and time of the next scheduled check (RFC 3339 format).
        """
        return pulumi.get(self, "next_check_at")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the domain is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the domain should be created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="revokedAt")
    def revoked_at(self) -> pulumi.Output[str]:
        """
        The date and time of the revocation of the domain (RFC 3339 format).
        """
        return pulumi.get(self, "revoked_at")

    @property
    @pulumi.getter(name="smtpHost")
    def smtp_host(self) -> pulumi.Output[str]:
        """
        SMTP host to use to send emails
        """
        return pulumi.get(self, "smtp_host")

    @property
    @pulumi.getter(name="smtpPort")
    def smtp_port(self) -> pulumi.Output[int]:
        """
        SMTP port to use to send emails over TLS. (Port 587)
        """
        return pulumi.get(self, "smtp_port")

    @property
    @pulumi.getter(name="smtpPortAlternative")
    def smtp_port_alternative(self) -> pulumi.Output[int]:
        """
        SMTP port to use to send emails over TLS. (Port 2587)
        """
        return pulumi.get(self, "smtp_port_alternative")

    @property
    @pulumi.getter(name="smtpPortUnsecure")
    def smtp_port_unsecure(self) -> pulumi.Output[int]:
        """
        SMTP port to use to send emails. (Port 25)
        """
        return pulumi.get(self, "smtp_port_unsecure")

    @property
    @pulumi.getter(name="smtpsPort")
    def smtps_port(self) -> pulumi.Output[int]:
        """
        SMTPS port to use to send emails over TLS Wrapper. (Port 465)
        """
        return pulumi.get(self, "smtps_port")

    @property
    @pulumi.getter(name="smtpsPortAlternative")
    def smtps_port_alternative(self) -> pulumi.Output[int]:
        """
        SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
        """
        return pulumi.get(self, "smtps_port_alternative")

    @property
    @pulumi.getter(name="spfConfig")
    def spf_config(self) -> pulumi.Output[str]:
        """
        The snippet of the SPF record that should be registered in the DNS zone.
        """
        return pulumi.get(self, "spf_config")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the Transaction Email Domain.
        """
        return pulumi.get(self, "status")

