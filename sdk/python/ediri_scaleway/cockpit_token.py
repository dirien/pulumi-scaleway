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

__all__ = ['CockpitTokenArgs', 'CockpitToken']

@pulumi.input_type
class CockpitTokenArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input['CockpitTokenScopesArgs']] = None):
        """
        The set of arguments for constructing a CockpitToken resource.
        :param pulumi.Input[str] name: The name of the token.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input['CockpitTokenScopesArgs'] scopes: Allowed scopes.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input['CockpitTokenScopesArgs']]:
        """
        Allowed scopes.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input['CockpitTokenScopesArgs']]):
        pulumi.set(self, "scopes", value)


@pulumi.input_type
class _CockpitTokenState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input['CockpitTokenScopesArgs']] = None,
                 secret_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CockpitToken resources.
        :param pulumi.Input[str] name: The name of the token.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input['CockpitTokenScopesArgs'] scopes: Allowed scopes.
        :param pulumi.Input[str] secret_key: The secret key of the token.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the token.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input['CockpitTokenScopesArgs']]:
        """
        Allowed scopes.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input['CockpitTokenScopesArgs']]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key of the token.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)


class CockpitToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[pulumi.InputType['CockpitTokenScopesArgs']]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Cockpit Tokens.

        For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#tokens).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        main_cockpit = scaleway.get_cockpit()
        # Create a token for the cockpit that can write metrics and logs
        main_cockpit_token = scaleway.CockpitToken("mainCockpitToken", project_id=main_cockpit.project_id)
        ```

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        main_cockpit = scaleway.get_cockpit()
        # Create a token for the cockpit that can read metrics and logs but not write
        main_cockpit_token = scaleway.CockpitToken("mainCockpitToken",
            project_id=main_cockpit.project_id,
            scopes=scaleway.CockpitTokenScopesArgs(
                query_metrics=True,
                write_metrics=False,
                query_logs=True,
                write_logs=False,
            ))
        ```

        ## Import

        Cockpits can be imported using the token ID, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/cockpitToken:CockpitToken main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the token.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input[pulumi.InputType['CockpitTokenScopesArgs']] scopes: Allowed scopes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CockpitTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Cockpit Tokens.

        For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#tokens).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        main_cockpit = scaleway.get_cockpit()
        # Create a token for the cockpit that can write metrics and logs
        main_cockpit_token = scaleway.CockpitToken("mainCockpitToken", project_id=main_cockpit.project_id)
        ```

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import pulumi_scaleway as scaleway

        main_cockpit = scaleway.get_cockpit()
        # Create a token for the cockpit that can read metrics and logs but not write
        main_cockpit_token = scaleway.CockpitToken("mainCockpitToken",
            project_id=main_cockpit.project_id,
            scopes=scaleway.CockpitTokenScopesArgs(
                query_metrics=True,
                write_metrics=False,
                query_logs=True,
                write_logs=False,
            ))
        ```

        ## Import

        Cockpits can be imported using the token ID, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/cockpitToken:CockpitToken main 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param CockpitTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CockpitTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[pulumi.InputType['CockpitTokenScopesArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CockpitTokenArgs.__new__(CockpitTokenArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["scopes"] = scopes
            __props__.__dict__["secret_key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(CockpitToken, __self__).__init__(
            'scaleway:index/cockpitToken:CockpitToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[pulumi.InputType['CockpitTokenScopesArgs']]] = None,
            secret_key: Optional[pulumi.Input[str]] = None) -> 'CockpitToken':
        """
        Get an existing CockpitToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the token.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the cockpit is associated with.
        :param pulumi.Input[pulumi.InputType['CockpitTokenScopesArgs']] scopes: Allowed scopes.
        :param pulumi.Input[str] secret_key: The secret key of the token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CockpitTokenState.__new__(_CockpitTokenState)

        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["secret_key"] = secret_key
        return CockpitToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the token.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the cockpit is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output['outputs.CockpitTokenScopes']:
        """
        Allowed scopes.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The secret key of the token.
        """
        return pulumi.get(self, "secret_key")

