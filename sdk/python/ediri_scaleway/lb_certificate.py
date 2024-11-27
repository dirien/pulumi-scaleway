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
from ._inputs import *

__all__ = ['LbCertificateArgs', 'LbCertificate']

@pulumi.input_type
class LbCertificateArgs:
    def __init__(__self__, *,
                 lb_id: pulumi.Input[str],
                 custom_certificate: Optional[pulumi.Input['LbCertificateCustomCertificateArgs']] = None,
                 letsencrypt: Optional[pulumi.Input['LbCertificateLetsencryptArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LbCertificate resource.
        :param pulumi.Input[str] lb_id: The load-balancer ID
        :param pulumi.Input['LbCertificateCustomCertificateArgs'] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input['LbCertificateLetsencryptArgs'] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[str] name: The name of the load-balancer certificate
        """
        pulumi.set(__self__, "lb_id", lb_id)
        if custom_certificate is not None:
            pulumi.set(__self__, "custom_certificate", custom_certificate)
        if letsencrypt is not None:
            pulumi.set(__self__, "letsencrypt", letsencrypt)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Input[str]:
        """
        The load-balancer ID
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter(name="customCertificate")
    def custom_certificate(self) -> Optional[pulumi.Input['LbCertificateCustomCertificateArgs']]:
        """
        The custom type certificate type configuration
        """
        return pulumi.get(self, "custom_certificate")

    @custom_certificate.setter
    def custom_certificate(self, value: Optional[pulumi.Input['LbCertificateCustomCertificateArgs']]):
        pulumi.set(self, "custom_certificate", value)

    @property
    @pulumi.getter
    def letsencrypt(self) -> Optional[pulumi.Input['LbCertificateLetsencryptArgs']]:
        """
        The Let's Encrypt type certificate configuration
        """
        return pulumi.get(self, "letsencrypt")

    @letsencrypt.setter
    def letsencrypt(self, value: Optional[pulumi.Input['LbCertificateLetsencryptArgs']]):
        pulumi.set(self, "letsencrypt", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the load-balancer certificate
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _LbCertificateState:
    def __init__(__self__, *,
                 common_name: Optional[pulumi.Input[str]] = None,
                 custom_certificate: Optional[pulumi.Input['LbCertificateCustomCertificateArgs']] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 letsencrypt: Optional[pulumi.Input['LbCertificateLetsencryptArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_valid_after: Optional[pulumi.Input[str]] = None,
                 not_valid_before: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering LbCertificate resources.
        :param pulumi.Input[str] common_name: Main domain of the certificate
        :param pulumi.Input['LbCertificateCustomCertificateArgs'] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[str] fingerprint: The identifier (SHA-1) of the certificate
        :param pulumi.Input[str] lb_id: The load-balancer ID
        :param pulumi.Input['LbCertificateLetsencryptArgs'] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[str] name: The name of the load-balancer certificate
        :param pulumi.Input[str] not_valid_after: The not valid after validity bound timestamp
        :param pulumi.Input[str] not_valid_before: The not valid before validity bound timestamp
        :param pulumi.Input[str] status: Certificate status
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subject_alternative_names: The alternative domain names of the certificate
        """
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if custom_certificate is not None:
            pulumi.set(__self__, "custom_certificate", custom_certificate)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if lb_id is not None:
            pulumi.set(__self__, "lb_id", lb_id)
        if letsencrypt is not None:
            pulumi.set(__self__, "letsencrypt", letsencrypt)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if not_valid_after is not None:
            pulumi.set(__self__, "not_valid_after", not_valid_after)
        if not_valid_before is not None:
            pulumi.set(__self__, "not_valid_before", not_valid_before)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subject_alternative_names is not None:
            pulumi.set(__self__, "subject_alternative_names", subject_alternative_names)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[pulumi.Input[str]]:
        """
        Main domain of the certificate
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter(name="customCertificate")
    def custom_certificate(self) -> Optional[pulumi.Input['LbCertificateCustomCertificateArgs']]:
        """
        The custom type certificate type configuration
        """
        return pulumi.get(self, "custom_certificate")

    @custom_certificate.setter
    def custom_certificate(self, value: Optional[pulumi.Input['LbCertificateCustomCertificateArgs']]):
        pulumi.set(self, "custom_certificate", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier (SHA-1) of the certificate
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[pulumi.Input[str]]:
        """
        The load-balancer ID
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter
    def letsencrypt(self) -> Optional[pulumi.Input['LbCertificateLetsencryptArgs']]:
        """
        The Let's Encrypt type certificate configuration
        """
        return pulumi.get(self, "letsencrypt")

    @letsencrypt.setter
    def letsencrypt(self, value: Optional[pulumi.Input['LbCertificateLetsencryptArgs']]):
        pulumi.set(self, "letsencrypt", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the load-balancer certificate
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notValidAfter")
    def not_valid_after(self) -> Optional[pulumi.Input[str]]:
        """
        The not valid after validity bound timestamp
        """
        return pulumi.get(self, "not_valid_after")

    @not_valid_after.setter
    def not_valid_after(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_valid_after", value)

    @property
    @pulumi.getter(name="notValidBefore")
    def not_valid_before(self) -> Optional[pulumi.Input[str]]:
        """
        The not valid before validity bound timestamp
        """
        return pulumi.get(self, "not_valid_before")

    @not_valid_before.setter
    def not_valid_before(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_valid_before", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate status
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The alternative domain names of the certificate
        """
        return pulumi.get(self, "subject_alternative_names")

    @subject_alternative_names.setter
    def subject_alternative_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subject_alternative_names", value)


class LbCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_certificate: Optional[pulumi.Input[Union['LbCertificateCustomCertificateArgs', 'LbCertificateCustomCertificateArgsDict']]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 letsencrypt: Optional[pulumi.Input[Union['LbCertificateLetsencryptArgs', 'LbCertificateLetsencryptArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LbCertificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['LbCertificateCustomCertificateArgs', 'LbCertificateCustomCertificateArgsDict']] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[str] lb_id: The load-balancer ID
        :param pulumi.Input[Union['LbCertificateLetsencryptArgs', 'LbCertificateLetsencryptArgsDict']] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[str] name: The name of the load-balancer certificate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LbCertificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LbCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_certificate: Optional[pulumi.Input[Union['LbCertificateCustomCertificateArgs', 'LbCertificateCustomCertificateArgsDict']]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 letsencrypt: Optional[pulumi.Input[Union['LbCertificateLetsencryptArgs', 'LbCertificateLetsencryptArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbCertificateArgs.__new__(LbCertificateArgs)

            __props__.__dict__["custom_certificate"] = custom_certificate
            if lb_id is None and not opts.urn:
                raise TypeError("Missing required property 'lb_id'")
            __props__.__dict__["lb_id"] = lb_id
            __props__.__dict__["letsencrypt"] = letsencrypt
            __props__.__dict__["name"] = name
            __props__.__dict__["common_name"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["not_valid_after"] = None
            __props__.__dict__["not_valid_before"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["subject_alternative_names"] = None
        super(LbCertificate, __self__).__init__(
            'scaleway:index/lbCertificate:LbCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            common_name: Optional[pulumi.Input[str]] = None,
            custom_certificate: Optional[pulumi.Input[Union['LbCertificateCustomCertificateArgs', 'LbCertificateCustomCertificateArgsDict']]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            lb_id: Optional[pulumi.Input[str]] = None,
            letsencrypt: Optional[pulumi.Input[Union['LbCertificateLetsencryptArgs', 'LbCertificateLetsencryptArgsDict']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            not_valid_after: Optional[pulumi.Input[str]] = None,
            not_valid_before: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subject_alternative_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'LbCertificate':
        """
        Get an existing LbCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] common_name: Main domain of the certificate
        :param pulumi.Input[Union['LbCertificateCustomCertificateArgs', 'LbCertificateCustomCertificateArgsDict']] custom_certificate: The custom type certificate type configuration
        :param pulumi.Input[str] fingerprint: The identifier (SHA-1) of the certificate
        :param pulumi.Input[str] lb_id: The load-balancer ID
        :param pulumi.Input[Union['LbCertificateLetsencryptArgs', 'LbCertificateLetsencryptArgsDict']] letsencrypt: The Let's Encrypt type certificate configuration
        :param pulumi.Input[str] name: The name of the load-balancer certificate
        :param pulumi.Input[str] not_valid_after: The not valid after validity bound timestamp
        :param pulumi.Input[str] not_valid_before: The not valid before validity bound timestamp
        :param pulumi.Input[str] status: Certificate status
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subject_alternative_names: The alternative domain names of the certificate
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbCertificateState.__new__(_LbCertificateState)

        __props__.__dict__["common_name"] = common_name
        __props__.__dict__["custom_certificate"] = custom_certificate
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["lb_id"] = lb_id
        __props__.__dict__["letsencrypt"] = letsencrypt
        __props__.__dict__["name"] = name
        __props__.__dict__["not_valid_after"] = not_valid_after
        __props__.__dict__["not_valid_before"] = not_valid_before
        __props__.__dict__["status"] = status
        __props__.__dict__["subject_alternative_names"] = subject_alternative_names
        return LbCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Output[str]:
        """
        Main domain of the certificate
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="customCertificate")
    def custom_certificate(self) -> pulumi.Output[Optional['outputs.LbCertificateCustomCertificate']]:
        """
        The custom type certificate type configuration
        """
        return pulumi.get(self, "custom_certificate")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        The identifier (SHA-1) of the certificate
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Output[str]:
        """
        The load-balancer ID
        """
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def letsencrypt(self) -> pulumi.Output[Optional['outputs.LbCertificateLetsencrypt']]:
        """
        The Let's Encrypt type certificate configuration
        """
        return pulumi.get(self, "letsencrypt")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the load-balancer certificate
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notValidAfter")
    def not_valid_after(self) -> pulumi.Output[str]:
        """
        The not valid after validity bound timestamp
        """
        return pulumi.get(self, "not_valid_after")

    @property
    @pulumi.getter(name="notValidBefore")
    def not_valid_before(self) -> pulumi.Output[str]:
        """
        The not valid before validity bound timestamp
        """
        return pulumi.get(self, "not_valid_before")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Certificate status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subjectAlternativeNames")
    def subject_alternative_names(self) -> pulumi.Output[Sequence[str]]:
        """
        The alternative domain names of the certificate
        """
        return pulumi.get(self, "subject_alternative_names")

