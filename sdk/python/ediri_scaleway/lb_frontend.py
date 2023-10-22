# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['LbFrontendArgs', 'LbFrontend']

@pulumi.input_type
class LbFrontendArgs:
    def __init__(__self__, *,
                 backend_id: pulumi.Input[str],
                 inbound_port: pulumi.Input[int],
                 lb_id: pulumi.Input[str],
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LbFrontend resource.
        :param pulumi.Input[str] backend_id: The load-balancer backend ID this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen on the front side.
        :param pulumi.Input[str] lb_id: The load-balancer ID this frontend is attached to.
        :param pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]] acls: A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of Certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        LbFrontendArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            backend_id=backend_id,
            inbound_port=inbound_port,
            lb_id=lb_id,
            acls=acls,
            certificate_ids=certificate_ids,
            enable_http3=enable_http3,
            external_acls=external_acls,
            name=name,
            timeout_client=timeout_client,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             backend_id: pulumi.Input[str],
             inbound_port: pulumi.Input[int],
             lb_id: pulumi.Input[str],
             acls: Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]] = None,
             certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             enable_http3: Optional[pulumi.Input[bool]] = None,
             external_acls: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             timeout_client: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'backendId' in kwargs:
            backend_id = kwargs['backendId']
        if 'inboundPort' in kwargs:
            inbound_port = kwargs['inboundPort']
        if 'lbId' in kwargs:
            lb_id = kwargs['lbId']
        if 'certificateIds' in kwargs:
            certificate_ids = kwargs['certificateIds']
        if 'enableHttp3' in kwargs:
            enable_http3 = kwargs['enableHttp3']
        if 'externalAcls' in kwargs:
            external_acls = kwargs['externalAcls']
        if 'timeoutClient' in kwargs:
            timeout_client = kwargs['timeoutClient']

        _setter("backend_id", backend_id)
        _setter("inbound_port", inbound_port)
        _setter("lb_id", lb_id)
        if acls is not None:
            _setter("acls", acls)
        if certificate_ids is not None:
            _setter("certificate_ids", certificate_ids)
        if enable_http3 is not None:
            _setter("enable_http3", enable_http3)
        if external_acls is not None:
            _setter("external_acls", external_acls)
        if name is not None:
            _setter("name", name)
        if timeout_client is not None:
            _setter("timeout_client", timeout_client)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Input[str]:
        """
        The load-balancer backend ID this frontend is attached to.

        > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="inboundPort")
    def inbound_port(self) -> pulumi.Input[int]:
        """
        TCP port to listen on the front side.
        """
        return pulumi.get(self, "inbound_port")

    @inbound_port.setter
    def inbound_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "inbound_port", value)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Input[str]:
        """
        The load-balancer ID this frontend is attached to.
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "lb_id", value)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]]:
        """
        A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of Certificate IDs that should be used by the frontend.

        > **Important:** Certificates are not allowed on port 80.
        """
        return pulumi.get(self, "certificate_ids")

    @certificate_ids.setter
    def certificate_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "certificate_ids", value)

    @property
    @pulumi.getter(name="enableHttp3")
    def enable_http3(self) -> Optional[pulumi.Input[bool]]:
        """
        Activates HTTP/3 protocol.
        """
        return pulumi.get(self, "enable_http3")

    @enable_http3.setter
    def enable_http3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_http3", value)

    @property
    @pulumi.getter(name="externalAcls")
    def external_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether to use lb_acl.
        If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        """
        return pulumi.get(self, "external_acls")

    @external_acls.setter
    def external_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_acls", value)

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
    @pulumi.getter(name="timeoutClient")
    def timeout_client(self) -> Optional[pulumi.Input[str]]:
        """
        Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        return pulumi.get(self, "timeout_client")

    @timeout_client.setter
    def timeout_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout_client", value)


@pulumi.input_type
class _LbFrontendState:
    def __init__(__self__, *,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 inbound_port: Optional[pulumi.Input[int]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LbFrontend resources.
        :param pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]] acls: A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        :param pulumi.Input[str] backend_id: The load-balancer backend ID this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[str] certificate_id: (Deprecated) first certificate ID used by the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of Certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen on the front side.
        :param pulumi.Input[str] lb_id: The load-balancer ID this frontend is attached to.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        _LbFrontendState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            acls=acls,
            backend_id=backend_id,
            certificate_id=certificate_id,
            certificate_ids=certificate_ids,
            enable_http3=enable_http3,
            external_acls=external_acls,
            inbound_port=inbound_port,
            lb_id=lb_id,
            name=name,
            timeout_client=timeout_client,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             acls: Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]] = None,
             backend_id: Optional[pulumi.Input[str]] = None,
             certificate_id: Optional[pulumi.Input[str]] = None,
             certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             enable_http3: Optional[pulumi.Input[bool]] = None,
             external_acls: Optional[pulumi.Input[bool]] = None,
             inbound_port: Optional[pulumi.Input[int]] = None,
             lb_id: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             timeout_client: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'backendId' in kwargs:
            backend_id = kwargs['backendId']
        if 'certificateId' in kwargs:
            certificate_id = kwargs['certificateId']
        if 'certificateIds' in kwargs:
            certificate_ids = kwargs['certificateIds']
        if 'enableHttp3' in kwargs:
            enable_http3 = kwargs['enableHttp3']
        if 'externalAcls' in kwargs:
            external_acls = kwargs['externalAcls']
        if 'inboundPort' in kwargs:
            inbound_port = kwargs['inboundPort']
        if 'lbId' in kwargs:
            lb_id = kwargs['lbId']
        if 'timeoutClient' in kwargs:
            timeout_client = kwargs['timeoutClient']

        if acls is not None:
            _setter("acls", acls)
        if backend_id is not None:
            _setter("backend_id", backend_id)
        if certificate_id is not None:
            warnings.warn("""Please use certificate_ids""", DeprecationWarning)
            pulumi.log.warn("""certificate_id is deprecated: Please use certificate_ids""")
        if certificate_id is not None:
            _setter("certificate_id", certificate_id)
        if certificate_ids is not None:
            _setter("certificate_ids", certificate_ids)
        if enable_http3 is not None:
            _setter("enable_http3", enable_http3)
        if external_acls is not None:
            _setter("external_acls", external_acls)
        if inbound_port is not None:
            _setter("inbound_port", inbound_port)
        if lb_id is not None:
            _setter("lb_id", lb_id)
        if name is not None:
            _setter("name", name)
        if timeout_client is not None:
            _setter("timeout_client", timeout_client)

    @property
    @pulumi.getter
    def acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]]:
        """
        A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        """
        return pulumi.get(self, "acls")

    @acls.setter
    def acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['LbFrontendAclArgs']]]]):
        pulumi.set(self, "acls", value)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> Optional[pulumi.Input[str]]:
        """
        The load-balancer backend ID this frontend is attached to.

        > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Deprecated) first certificate ID used by the frontend.
        """
        warnings.warn("""Please use certificate_ids""", DeprecationWarning)
        pulumi.log.warn("""certificate_id is deprecated: Please use certificate_ids""")

        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of Certificate IDs that should be used by the frontend.

        > **Important:** Certificates are not allowed on port 80.
        """
        return pulumi.get(self, "certificate_ids")

    @certificate_ids.setter
    def certificate_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "certificate_ids", value)

    @property
    @pulumi.getter(name="enableHttp3")
    def enable_http3(self) -> Optional[pulumi.Input[bool]]:
        """
        Activates HTTP/3 protocol.
        """
        return pulumi.get(self, "enable_http3")

    @enable_http3.setter
    def enable_http3(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_http3", value)

    @property
    @pulumi.getter(name="externalAcls")
    def external_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean to specify whether to use lb_acl.
        If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        """
        return pulumi.get(self, "external_acls")

    @external_acls.setter
    def external_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_acls", value)

    @property
    @pulumi.getter(name="inboundPort")
    def inbound_port(self) -> Optional[pulumi.Input[int]]:
        """
        TCP port to listen on the front side.
        """
        return pulumi.get(self, "inbound_port")

    @inbound_port.setter
    def inbound_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "inbound_port", value)

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> Optional[pulumi.Input[str]]:
        """
        The load-balancer ID this frontend is attached to.
        """
        return pulumi.get(self, "lb_id")

    @lb_id.setter
    def lb_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lb_id", value)

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
    @pulumi.getter(name="timeoutClient")
    def timeout_client(self) -> Optional[pulumi.Input[str]]:
        """
        Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        return pulumi.get(self, "timeout_client")

    @timeout_client.setter
    def timeout_client(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout_client", value)


class LbFrontend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbFrontendAclArgs']]]]] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 inbound_port: Optional[pulumi.Input[int]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Load-Balancer Frontends. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).

        ## Examples Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        frontend01 = scaleway.LbFrontend("frontend01",
            lb_id=scaleway_lb["lb01"]["id"],
            backend_id=scaleway_lb_backend["backend01"]["id"],
            inbound_port=80)
        ```

        ## Import

        Load-Balancer frontend can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/lbFrontend:LbFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbFrontendAclArgs']]]] acls: A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        :param pulumi.Input[str] backend_id: The load-balancer backend ID this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of Certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen on the front side.
        :param pulumi.Input[str] lb_id: The load-balancer ID this frontend is attached to.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LbFrontendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Load-Balancer Frontends. For more information, see [the documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).

        ## Examples Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        frontend01 = scaleway.LbFrontend("frontend01",
            lb_id=scaleway_lb["lb01"]["id"],
            backend_id=scaleway_lb_backend["backend01"]["id"],
            inbound_port=80)
        ```

        ## Import

        Load-Balancer frontend can be imported using the `{zone}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/lbFrontend:LbFrontend frontend01 fr-par-1/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param LbFrontendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LbFrontendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            LbFrontendArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbFrontendAclArgs']]]]] = None,
                 backend_id: Optional[pulumi.Input[str]] = None,
                 certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_http3: Optional[pulumi.Input[bool]] = None,
                 external_acls: Optional[pulumi.Input[bool]] = None,
                 inbound_port: Optional[pulumi.Input[int]] = None,
                 lb_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 timeout_client: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LbFrontendArgs.__new__(LbFrontendArgs)

            __props__.__dict__["acls"] = acls
            if backend_id is None and not opts.urn:
                raise TypeError("Missing required property 'backend_id'")
            __props__.__dict__["backend_id"] = backend_id
            __props__.__dict__["certificate_ids"] = certificate_ids
            __props__.__dict__["enable_http3"] = enable_http3
            __props__.__dict__["external_acls"] = external_acls
            if inbound_port is None and not opts.urn:
                raise TypeError("Missing required property 'inbound_port'")
            __props__.__dict__["inbound_port"] = inbound_port
            if lb_id is None and not opts.urn:
                raise TypeError("Missing required property 'lb_id'")
            __props__.__dict__["lb_id"] = lb_id
            __props__.__dict__["name"] = name
            __props__.__dict__["timeout_client"] = timeout_client
            __props__.__dict__["certificate_id"] = None
        super(LbFrontend, __self__).__init__(
            'scaleway:index/lbFrontend:LbFrontend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbFrontendAclArgs']]]]] = None,
            backend_id: Optional[pulumi.Input[str]] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            certificate_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enable_http3: Optional[pulumi.Input[bool]] = None,
            external_acls: Optional[pulumi.Input[bool]] = None,
            inbound_port: Optional[pulumi.Input[int]] = None,
            lb_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            timeout_client: Optional[pulumi.Input[str]] = None) -> 'LbFrontend':
        """
        Get an existing LbFrontend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['LbFrontendAclArgs']]]] acls: A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        :param pulumi.Input[str] backend_id: The load-balancer backend ID this frontend is attached to.
               
               > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        :param pulumi.Input[str] certificate_id: (Deprecated) first certificate ID used by the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] certificate_ids: List of Certificate IDs that should be used by the frontend.
               
               > **Important:** Certificates are not allowed on port 80.
        :param pulumi.Input[bool] enable_http3: Activates HTTP/3 protocol.
        :param pulumi.Input[bool] external_acls: A boolean to specify whether to use lb_acl.
               If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        :param pulumi.Input[int] inbound_port: TCP port to listen on the front side.
        :param pulumi.Input[str] lb_id: The load-balancer ID this frontend is attached to.
        :param pulumi.Input[str] name: The ACL name. If not provided it will be randomly generated.
        :param pulumi.Input[str] timeout_client: Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LbFrontendState.__new__(_LbFrontendState)

        __props__.__dict__["acls"] = acls
        __props__.__dict__["backend_id"] = backend_id
        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["certificate_ids"] = certificate_ids
        __props__.__dict__["enable_http3"] = enable_http3
        __props__.__dict__["external_acls"] = external_acls
        __props__.__dict__["inbound_port"] = inbound_port
        __props__.__dict__["lb_id"] = lb_id
        __props__.__dict__["name"] = name
        __props__.__dict__["timeout_client"] = timeout_client
        return LbFrontend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def acls(self) -> pulumi.Output[Optional[Sequence['outputs.LbFrontendAcl']]]:
        """
        A list of ACL rules to apply to the load-balancer frontend.  Defined below.
        """
        return pulumi.get(self, "acls")

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Output[str]:
        """
        The load-balancer backend ID this frontend is attached to.

        > **Important:** Updates to `lb_id` or `backend_id` will recreate the frontend.
        """
        return pulumi.get(self, "backend_id")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Output[str]:
        """
        (Deprecated) first certificate ID used by the frontend.
        """
        warnings.warn("""Please use certificate_ids""", DeprecationWarning)
        pulumi.log.warn("""certificate_id is deprecated: Please use certificate_ids""")

        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="certificateIds")
    def certificate_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of Certificate IDs that should be used by the frontend.

        > **Important:** Certificates are not allowed on port 80.
        """
        return pulumi.get(self, "certificate_ids")

    @property
    @pulumi.getter(name="enableHttp3")
    def enable_http3(self) -> pulumi.Output[Optional[bool]]:
        """
        Activates HTTP/3 protocol.
        """
        return pulumi.get(self, "enable_http3")

    @property
    @pulumi.getter(name="externalAcls")
    def external_acls(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean to specify whether to use lb_acl.
        If `external_acls` is set to `true`, `acl` can not be set directly in the lb frontend.
        """
        return pulumi.get(self, "external_acls")

    @property
    @pulumi.getter(name="inboundPort")
    def inbound_port(self) -> pulumi.Output[int]:
        """
        TCP port to listen on the front side.
        """
        return pulumi.get(self, "inbound_port")

    @property
    @pulumi.getter(name="lbId")
    def lb_id(self) -> pulumi.Output[str]:
        """
        The load-balancer ID this frontend is attached to.
        """
        return pulumi.get(self, "lb_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The ACL name. If not provided it will be randomly generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="timeoutClient")
    def timeout_client(self) -> pulumi.Output[Optional[str]]:
        """
        Maximum inactivity time on the client side. (e.g.: `1s`)
        """
        return pulumi.get(self, "timeout_client")

