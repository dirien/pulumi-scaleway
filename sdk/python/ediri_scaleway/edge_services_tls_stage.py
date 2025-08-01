# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['EdgeServicesTlsStageArgs', 'EdgeServicesTlsStage']

@pulumi.input_type
class EdgeServicesTlsStageArgs:
    def __init__(__self__, *,
                 pipeline_id: pulumi.Input[builtins.str],
                 backend_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 cache_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 managed_certificate: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 secrets: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]]] = None,
                 waf_stage_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a EdgeServicesTlsStage resource.
        :param pulumi.Input[builtins.str] pipeline_id: The ID of the pipeline.
        :param pulumi.Input[builtins.str] backend_stage_id: The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.str] cache_stage_id: The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.bool] managed_certificate: Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the project the TLS stage is associated with.
        :param pulumi.Input[builtins.str] route_stage_id: The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]] secrets: The TLS secrets.
        :param pulumi.Input[builtins.str] waf_stage_id: The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        pulumi.set(__self__, "pipeline_id", pipeline_id)
        if backend_stage_id is not None:
            pulumi.set(__self__, "backend_stage_id", backend_stage_id)
        if cache_stage_id is not None:
            pulumi.set(__self__, "cache_stage_id", cache_stage_id)
        if managed_certificate is not None:
            pulumi.set(__self__, "managed_certificate", managed_certificate)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if route_stage_id is not None:
            pulumi.set(__self__, "route_stage_id", route_stage_id)
        if secrets is not None:
            pulumi.set(__self__, "secrets", secrets)
        if waf_stage_id is not None:
            pulumi.set(__self__, "waf_stage_id", waf_stage_id)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the pipeline.
        """
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter(name="backendStageId")
    def backend_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "backend_stage_id")

    @backend_stage_id.setter
    def backend_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend_stage_id", value)

    @property
    @pulumi.getter(name="cacheStageId")
    def cache_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "cache_stage_id")

    @cache_stage_id.setter
    def cache_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cache_stage_id", value)

    @property
    @pulumi.getter(name="managedCertificate")
    def managed_certificate(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        """
        return pulumi.get(self, "managed_certificate")

    @managed_certificate.setter
    def managed_certificate(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "managed_certificate", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `project_id`) The ID of the project the TLS stage is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="routeStageId")
    def route_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "route_stage_id")

    @route_stage_id.setter
    def route_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "route_stage_id", value)

    @property
    @pulumi.getter
    def secrets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]]]:
        """
        The TLS secrets.
        """
        return pulumi.get(self, "secrets")

    @secrets.setter
    def secrets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]]]):
        pulumi.set(self, "secrets", value)

    @property
    @pulumi.getter(name="wafStageId")
    def waf_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "waf_stage_id")

    @waf_stage_id.setter
    def waf_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "waf_stage_id", value)


@pulumi.input_type
class _EdgeServicesTlsStageState:
    def __init__(__self__, *,
                 backend_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 cache_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 certificate_expires_at: Optional[pulumi.Input[builtins.str]] = None,
                 created_at: Optional[pulumi.Input[builtins.str]] = None,
                 managed_certificate: Optional[pulumi.Input[builtins.bool]] = None,
                 pipeline_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 secrets: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]]] = None,
                 updated_at: Optional[pulumi.Input[builtins.str]] = None,
                 waf_stage_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering EdgeServicesTlsStage resources.
        :param pulumi.Input[builtins.str] backend_stage_id: The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.str] cache_stage_id: The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.str] certificate_expires_at: The expiration date of the certificate.
        :param pulumi.Input[builtins.str] created_at: The date and time of the creation of the TLS stage.
        :param pulumi.Input[builtins.bool] managed_certificate: Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        :param pulumi.Input[builtins.str] pipeline_id: The ID of the pipeline.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the project the TLS stage is associated with.
        :param pulumi.Input[builtins.str] route_stage_id: The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]] secrets: The TLS secrets.
        :param pulumi.Input[builtins.str] updated_at: The date and time of the last update of the TLS stage.
        :param pulumi.Input[builtins.str] waf_stage_id: The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        if backend_stage_id is not None:
            pulumi.set(__self__, "backend_stage_id", backend_stage_id)
        if cache_stage_id is not None:
            pulumi.set(__self__, "cache_stage_id", cache_stage_id)
        if certificate_expires_at is not None:
            pulumi.set(__self__, "certificate_expires_at", certificate_expires_at)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if managed_certificate is not None:
            pulumi.set(__self__, "managed_certificate", managed_certificate)
        if pipeline_id is not None:
            pulumi.set(__self__, "pipeline_id", pipeline_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if route_stage_id is not None:
            pulumi.set(__self__, "route_stage_id", route_stage_id)
        if secrets is not None:
            pulumi.set(__self__, "secrets", secrets)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if waf_stage_id is not None:
            pulumi.set(__self__, "waf_stage_id", waf_stage_id)

    @property
    @pulumi.getter(name="backendStageId")
    def backend_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "backend_stage_id")

    @backend_stage_id.setter
    def backend_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend_stage_id", value)

    @property
    @pulumi.getter(name="cacheStageId")
    def cache_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "cache_stage_id")

    @cache_stage_id.setter
    def cache_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cache_stage_id", value)

    @property
    @pulumi.getter(name="certificateExpiresAt")
    def certificate_expires_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The expiration date of the certificate.
        """
        return pulumi.get(self, "certificate_expires_at")

    @certificate_expires_at.setter
    def certificate_expires_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "certificate_expires_at", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time of the creation of the TLS stage.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="managedCertificate")
    def managed_certificate(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        """
        return pulumi.get(self, "managed_certificate")

    @managed_certificate.setter
    def managed_certificate(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "managed_certificate", value)

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the pipeline.
        """
        return pulumi.get(self, "pipeline_id")

    @pipeline_id.setter
    def pipeline_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "pipeline_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `project_id`) The ID of the project the TLS stage is associated with.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="routeStageId")
    def route_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "route_stage_id")

    @route_stage_id.setter
    def route_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "route_stage_id", value)

    @property
    @pulumi.getter
    def secrets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]]]:
        """
        The TLS secrets.
        """
        return pulumi.get(self, "secrets")

    @secrets.setter
    def secrets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeServicesTlsStageSecretArgs']]]]):
        pulumi.set(self, "secrets", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time of the last update of the TLS stage.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="wafStageId")
    def waf_stage_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "waf_stage_id")

    @waf_stage_id.setter
    def waf_stage_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "waf_stage_id", value)


@pulumi.type_token("scaleway:index/edgeServicesTlsStage:EdgeServicesTlsStage")
class EdgeServicesTlsStage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 cache_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 managed_certificate: Optional[pulumi.Input[builtins.bool]] = None,
                 pipeline_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 secrets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EdgeServicesTlsStageSecretArgs', 'EdgeServicesTlsStageSecretArgsDict']]]]] = None,
                 waf_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Edge Services TLS Stages.

        ## Example Usage

        ### Managed

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.EdgeServicesTlsStage("main",
            pipeline_id=scaleway_edge_services_pipeline["main"]["id"],
            managed_certificate=True)
        ```

        ### With a certificate stored in Scaleway Secret Manager

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.EdgeServicesTlsStage("main",
            pipeline_id=scaleway_edge_services_pipeline["main"]["id"],
            secrets=[{
                "secret_id": "11111111-1111-1111-1111-111111111111",
                "region": "fr-par",
            }])
        ```

        ## Import

        TLS stages can be imported using the `{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/edgeServicesTlsStage:EdgeServicesTlsStage basic 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend_stage_id: The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.str] cache_stage_id: The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.bool] managed_certificate: Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        :param pulumi.Input[builtins.str] pipeline_id: The ID of the pipeline.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the project the TLS stage is associated with.
        :param pulumi.Input[builtins.str] route_stage_id: The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EdgeServicesTlsStageSecretArgs', 'EdgeServicesTlsStageSecretArgsDict']]]] secrets: The TLS secrets.
        :param pulumi.Input[builtins.str] waf_stage_id: The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EdgeServicesTlsStageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Edge Services TLS Stages.

        ## Example Usage

        ### Managed

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.EdgeServicesTlsStage("main",
            pipeline_id=scaleway_edge_services_pipeline["main"]["id"],
            managed_certificate=True)
        ```

        ### With a certificate stored in Scaleway Secret Manager

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main = scaleway.EdgeServicesTlsStage("main",
            pipeline_id=scaleway_edge_services_pipeline["main"]["id"],
            secrets=[{
                "secret_id": "11111111-1111-1111-1111-111111111111",
                "region": "fr-par",
            }])
        ```

        ## Import

        TLS stages can be imported using the `{id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/edgeServicesTlsStage:EdgeServicesTlsStage basic 11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param EdgeServicesTlsStageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EdgeServicesTlsStageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 cache_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 managed_certificate: Optional[pulumi.Input[builtins.bool]] = None,
                 pipeline_id: Optional[pulumi.Input[builtins.str]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 route_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 secrets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EdgeServicesTlsStageSecretArgs', 'EdgeServicesTlsStageSecretArgsDict']]]]] = None,
                 waf_stage_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EdgeServicesTlsStageArgs.__new__(EdgeServicesTlsStageArgs)

            __props__.__dict__["backend_stage_id"] = backend_stage_id
            __props__.__dict__["cache_stage_id"] = cache_stage_id
            __props__.__dict__["managed_certificate"] = managed_certificate
            if pipeline_id is None and not opts.urn:
                raise TypeError("Missing required property 'pipeline_id'")
            __props__.__dict__["pipeline_id"] = pipeline_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["route_stage_id"] = route_stage_id
            __props__.__dict__["secrets"] = secrets
            __props__.__dict__["waf_stage_id"] = waf_stage_id
            __props__.__dict__["certificate_expires_at"] = None
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(EdgeServicesTlsStage, __self__).__init__(
            'scaleway:index/edgeServicesTlsStage:EdgeServicesTlsStage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend_stage_id: Optional[pulumi.Input[builtins.str]] = None,
            cache_stage_id: Optional[pulumi.Input[builtins.str]] = None,
            certificate_expires_at: Optional[pulumi.Input[builtins.str]] = None,
            created_at: Optional[pulumi.Input[builtins.str]] = None,
            managed_certificate: Optional[pulumi.Input[builtins.bool]] = None,
            pipeline_id: Optional[pulumi.Input[builtins.str]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            route_stage_id: Optional[pulumi.Input[builtins.str]] = None,
            secrets: Optional[pulumi.Input[Sequence[pulumi.Input[Union['EdgeServicesTlsStageSecretArgs', 'EdgeServicesTlsStageSecretArgsDict']]]]] = None,
            updated_at: Optional[pulumi.Input[builtins.str]] = None,
            waf_stage_id: Optional[pulumi.Input[builtins.str]] = None) -> 'EdgeServicesTlsStage':
        """
        Get an existing EdgeServicesTlsStage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend_stage_id: The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.str] cache_stage_id: The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[builtins.str] certificate_expires_at: The expiration date of the certificate.
        :param pulumi.Input[builtins.str] created_at: The date and time of the creation of the TLS stage.
        :param pulumi.Input[builtins.bool] managed_certificate: Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        :param pulumi.Input[builtins.str] pipeline_id: The ID of the pipeline.
        :param pulumi.Input[builtins.str] project_id: `project_id`) The ID of the project the TLS stage is associated with.
        :param pulumi.Input[builtins.str] route_stage_id: The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        :param pulumi.Input[Sequence[pulumi.Input[Union['EdgeServicesTlsStageSecretArgs', 'EdgeServicesTlsStageSecretArgsDict']]]] secrets: The TLS secrets.
        :param pulumi.Input[builtins.str] updated_at: The date and time of the last update of the TLS stage.
        :param pulumi.Input[builtins.str] waf_stage_id: The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EdgeServicesTlsStageState.__new__(_EdgeServicesTlsStageState)

        __props__.__dict__["backend_stage_id"] = backend_stage_id
        __props__.__dict__["cache_stage_id"] = cache_stage_id
        __props__.__dict__["certificate_expires_at"] = certificate_expires_at
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["managed_certificate"] = managed_certificate
        __props__.__dict__["pipeline_id"] = pipeline_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["route_stage_id"] = route_stage_id
        __props__.__dict__["secrets"] = secrets
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["waf_stage_id"] = waf_stage_id
        return EdgeServicesTlsStage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendStageId")
    def backend_stage_id(self) -> pulumi.Output[builtins.str]:
        """
        The backend stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "backend_stage_id")

    @property
    @pulumi.getter(name="cacheStageId")
    def cache_stage_id(self) -> pulumi.Output[builtins.str]:
        """
        The cache stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "cache_stage_id")

    @property
    @pulumi.getter(name="certificateExpiresAt")
    def certificate_expires_at(self) -> pulumi.Output[builtins.str]:
        """
        The expiration date of the certificate.
        """
        return pulumi.get(self, "certificate_expires_at")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time of the creation of the TLS stage.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="managedCertificate")
    def managed_certificate(self) -> pulumi.Output[builtins.bool]:
        """
        Set to true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
        """
        return pulumi.get(self, "managed_certificate")

    @property
    @pulumi.getter(name="pipelineId")
    def pipeline_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the pipeline.
        """
        return pulumi.get(self, "pipeline_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        `project_id`) The ID of the project the TLS stage is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="routeStageId")
    def route_stage_id(self) -> pulumi.Output[builtins.str]:
        """
        The route stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "route_stage_id")

    @property
    @pulumi.getter
    def secrets(self) -> pulumi.Output[Sequence['outputs.EdgeServicesTlsStageSecret']]:
        """
        The TLS secrets.
        """
        return pulumi.get(self, "secrets")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[builtins.str]:
        """
        The date and time of the last update of the TLS stage.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="wafStageId")
    def waf_stage_id(self) -> pulumi.Output[builtins.str]:
        """
        The WAF stage ID the TLS stage will be linked to. Only one of `backend_stage_id`, `cache_stage_id`, `route_stage_id` and `waf_stage_id` should be specified.
        """
        return pulumi.get(self, "waf_stage_id")

