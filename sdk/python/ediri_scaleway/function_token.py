# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FunctionTokenArgs', 'FunctionToken']

@pulumi.input_type
class FunctionTokenArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FunctionToken resource.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] function_id: The ID of the function.
               
               > Only one of `namespace_id` or `function_id` must be set.
        :param pulumi.Input[str] namespace_id: The ID of the function namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        """
        FunctionTokenArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            expires_at=expires_at,
            function_id=function_id,
            namespace_id=namespace_id,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             expires_at: Optional[pulumi.Input[str]] = None,
             function_id: Optional[pulumi.Input[str]] = None,
             namespace_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'expiresAt' in kwargs:
            expires_at = kwargs['expiresAt']
        if 'functionId' in kwargs:
            function_id = kwargs['functionId']
        if 'namespaceId' in kwargs:
            namespace_id = kwargs['namespaceId']

        if description is not None:
            _setter("description", description)
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if function_id is not None:
            _setter("function_id", function_id)
        if namespace_id is not None:
            _setter("namespace_id", namespace_id)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the token.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The expiration date of the token.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the function.

        > Only one of `namespace_id` or `function_id` must be set.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the function namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.

        > **Important** Updates to any fields will recreate the token.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _FunctionTokenState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FunctionToken resources.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] function_id: The ID of the function.
               
               > Only one of `namespace_id` or `function_id` must be set.
        :param pulumi.Input[str] namespace_id: The ID of the function namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        :param pulumi.Input[str] token: The token.
        """
        _FunctionTokenState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            expires_at=expires_at,
            function_id=function_id,
            namespace_id=namespace_id,
            region=region,
            token=token,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             expires_at: Optional[pulumi.Input[str]] = None,
             function_id: Optional[pulumi.Input[str]] = None,
             namespace_id: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             token: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'expiresAt' in kwargs:
            expires_at = kwargs['expiresAt']
        if 'functionId' in kwargs:
            function_id = kwargs['functionId']
        if 'namespaceId' in kwargs:
            namespace_id = kwargs['namespaceId']

        if description is not None:
            _setter("description", description)
        if expires_at is not None:
            _setter("expires_at", expires_at)
        if function_id is not None:
            _setter("function_id", function_id)
        if namespace_id is not None:
            _setter("namespace_id", namespace_id)
        if region is not None:
            _setter("region", region)
        if token is not None:
            _setter("token", token)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the token.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The expiration date of the token.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the function.

        > Only one of `namespace_id` or `function_id` must be set.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the function namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region in which the namespace should be created.

        > **Important** Updates to any fields will recreate the token.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class FunctionToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Function Token.
        For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/#tokens-26b085).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_function_namespace = scaleway.FunctionNamespace("mainFunctionNamespace")
        main_function = scaleway.Function("mainFunction",
            namespace_id=main_function_namespace.id,
            runtime="go118",
            handler="Handle",
            privacy="private")
        # Namespace Token
        namespace = scaleway.FunctionToken("namespace",
            namespace_id=main_function_namespace.id,
            expires_at="2022-10-18T11:35:15+02:00")
        # Function Token
        function = scaleway.FunctionToken("function", function_id=main_function.id)
        ```

        ## Import

        Tokens can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/functionToken:FunctionToken main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] function_id: The ID of the function.
               
               > Only one of `namespace_id` or `function_id` must be set.
        :param pulumi.Input[str] namespace_id: The ID of the function namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FunctionTokenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Function Token.
        For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/#tokens-26b085).

        ## Examples

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_function_namespace = scaleway.FunctionNamespace("mainFunctionNamespace")
        main_function = scaleway.Function("mainFunction",
            namespace_id=main_function_namespace.id,
            runtime="go118",
            handler="Handle",
            privacy="private")
        # Namespace Token
        namespace = scaleway.FunctionToken("namespace",
            namespace_id=main_function_namespace.id,
            expires_at="2022-10-18T11:35:15+02:00")
        # Function Token
        function = scaleway.FunctionToken("function", function_id=main_function.id)
        ```

        ## Import

        Tokens can be imported using the `{region}/{id}`, e.g. bash

        ```sh
         $ pulumi import scaleway:index/functionToken:FunctionToken main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param FunctionTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            FunctionTokenArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionTokenArgs.__new__(FunctionTokenArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["function_id"] = function_id
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["region"] = region
            __props__.__dict__["token"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(FunctionToken, __self__).__init__(
            'scaleway:index/functionToken:FunctionToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            function_id: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None) -> 'FunctionToken':
        """
        Get an existing FunctionToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the token.
        :param pulumi.Input[str] expires_at: The expiration date of the token.
        :param pulumi.Input[str] function_id: The ID of the function.
               
               > Only one of `namespace_id` or `function_id` must be set.
        :param pulumi.Input[str] namespace_id: The ID of the function namespace.
        :param pulumi.Input[str] region: `region`). The region in which the namespace should be created.
               
               > **Important** Updates to any fields will recreate the token.
        :param pulumi.Input[str] token: The token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionTokenState.__new__(_FunctionTokenState)

        __props__.__dict__["description"] = description
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["region"] = region
        __props__.__dict__["token"] = token
        return FunctionToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the token.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[str]]:
        """
        The expiration date of the token.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the function.

        > Only one of `namespace_id` or `function_id` must be set.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the function namespace.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region in which the namespace should be created.

        > **Important** Updates to any fields will recreate the token.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The token.
        """
        return pulumi.get(self, "token")

