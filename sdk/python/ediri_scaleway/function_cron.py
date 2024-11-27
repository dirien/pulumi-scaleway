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

__all__ = ['FunctionCronArgs', 'FunctionCron']

@pulumi.input_type
class FunctionCronArgs:
    def __init__(__self__, *,
                 args: pulumi.Input[str],
                 function_id: pulumi.Input[str],
                 schedule: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FunctionCron resource.
        :param pulumi.Input[str] args: The key-value mapping to define arguments that will be passed to your function’s event object
        :param pulumi.Input[str] function_id: The unique identifier of the function to link to your CRON trigger.
        :param pulumi.Input[str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        :param pulumi.Input[str] name: The name of the function CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[str] region: `region`) The region
               in which the function was created.
        """
        pulumi.set(__self__, "args", args)
        pulumi.set(__self__, "function_id", function_id)
        pulumi.set(__self__, "schedule", schedule)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Input[str]:
        """
        The key-value mapping to define arguments that will be passed to your function’s event object
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: pulumi.Input[str]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Input[str]:
        """
        The unique identifier of the function to link to your CRON trigger.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Input[str]:
        """
        CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the function CRON trigger. If not provided, a random name is generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the function was created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _FunctionCronState:
    def __init__(__self__, *,
                 args: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FunctionCron resources.
        :param pulumi.Input[str] args: The key-value mapping to define arguments that will be passed to your function’s event object
        :param pulumi.Input[str] function_id: The unique identifier of the function to link to your CRON trigger.
        :param pulumi.Input[str] name: The name of the function CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[str] region: `region`) The region
               in which the function was created.
        :param pulumi.Input[str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        :param pulumi.Input[str] status: The CRON status.
        """
        if args is not None:
            pulumi.set(__self__, "args", args)
        if function_id is not None:
            pulumi.set(__self__, "function_id", function_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[str]]:
        """
        The key-value mapping to define arguments that will be passed to your function’s event object
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the function to link to your CRON trigger.
        """
        return pulumi.get(self, "function_id")

    @function_id.setter
    def function_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "function_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the function CRON trigger. If not provided, a random name is generated.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`) The region
        in which the function was created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The CRON status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class FunctionCron(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The `FunctionCron` resource allows you to create and manage CRON triggers for Scaleway [Serverless Functions](https://www.scaleway.com/en/docs/serverless/functions/).

        Refer to the Functions CRON triggers [documentation](https://www.scaleway.com/en/docs/serverless/functions/how-to/add-trigger-to-a-function/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers-list-all-triggers) for more information.

        ## Example Usage

        The following command allows you to add a CRON trigger to a Serverless Function.

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import json

        main_function_namespace = scaleway.FunctionNamespace("mainFunctionNamespace")
        main_function = scaleway.Function("mainFunction",
            namespace_id=main_function_namespace.id,
            runtime="node14",
            privacy="private",
            handler="handler.handle")
        main_function_cron = scaleway.FunctionCron("mainFunctionCron",
            function_id=main_function.id,
            schedule="0 0 * * *",
            args=json.dumps({
                "test": "scw",
            }))
        func = scaleway.FunctionCron("func",
            function_id=main_function.id,
            schedule="0 1 * * *",
            args=json.dumps({
                "my_var": "terraform",
            }))
        ```

        ## Import

        Function Cron can be imported using `{region}/{id}`, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:index/functionCron:FunctionCron main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] args: The key-value mapping to define arguments that will be passed to your function’s event object
        :param pulumi.Input[str] function_id: The unique identifier of the function to link to your CRON trigger.
        :param pulumi.Input[str] name: The name of the function CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[str] region: `region`) The region
               in which the function was created.
        :param pulumi.Input[str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionCronArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The `FunctionCron` resource allows you to create and manage CRON triggers for Scaleway [Serverless Functions](https://www.scaleway.com/en/docs/serverless/functions/).

        Refer to the Functions CRON triggers [documentation](https://www.scaleway.com/en/docs/serverless/functions/how-to/add-trigger-to-a-function/) and [API documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers-list-all-triggers) for more information.

        ## Example Usage

        The following command allows you to add a CRON trigger to a Serverless Function.

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import json

        main_function_namespace = scaleway.FunctionNamespace("mainFunctionNamespace")
        main_function = scaleway.Function("mainFunction",
            namespace_id=main_function_namespace.id,
            runtime="node14",
            privacy="private",
            handler="handler.handle")
        main_function_cron = scaleway.FunctionCron("mainFunctionCron",
            function_id=main_function.id,
            schedule="0 0 * * *",
            args=json.dumps({
                "test": "scw",
            }))
        func = scaleway.FunctionCron("func",
            function_id=main_function.id,
            schedule="0 1 * * *",
            args=json.dumps({
                "my_var": "terraform",
            }))
        ```

        ## Import

        Function Cron can be imported using `{region}/{id}`, as shown below:

        bash

        ```sh
        $ pulumi import scaleway:index/functionCron:FunctionCron main fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param FunctionCronArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionCronArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[str]] = None,
                 function_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionCronArgs.__new__(FunctionCronArgs)

            if args is None and not opts.urn:
                raise TypeError("Missing required property 'args'")
            __props__.__dict__["args"] = args
            if function_id is None and not opts.urn:
                raise TypeError("Missing required property 'function_id'")
            __props__.__dict__["function_id"] = function_id
            __props__.__dict__["name"] = name
            __props__.__dict__["region"] = region
            if schedule is None and not opts.urn:
                raise TypeError("Missing required property 'schedule'")
            __props__.__dict__["schedule"] = schedule
            __props__.__dict__["status"] = None
        super(FunctionCron, __self__).__init__(
            'scaleway:index/functionCron:FunctionCron',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            args: Optional[pulumi.Input[str]] = None,
            function_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'FunctionCron':
        """
        Get an existing FunctionCron resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] args: The key-value mapping to define arguments that will be passed to your function’s event object
        :param pulumi.Input[str] function_id: The unique identifier of the function to link to your CRON trigger.
        :param pulumi.Input[str] name: The name of the function CRON trigger. If not provided, a random name is generated.
        :param pulumi.Input[str] region: `region`) The region
               in which the function was created.
        :param pulumi.Input[str] schedule: CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        :param pulumi.Input[str] status: The CRON status.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionCronState.__new__(_FunctionCronState)

        __props__.__dict__["args"] = args
        __props__.__dict__["function_id"] = function_id
        __props__.__dict__["name"] = name
        __props__.__dict__["region"] = region
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["status"] = status
        return FunctionCron(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output[str]:
        """
        The key-value mapping to define arguments that will be passed to your function’s event object
        """
        return pulumi.get(self, "args")

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> pulumi.Output[str]:
        """
        The unique identifier of the function to link to your CRON trigger.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the function CRON trigger. If not provided, a random name is generated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`) The region
        in which the function was created.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[str]:
        """
        CRON format string (refer to the [CRON schedule reference](https://www.scaleway.com/en/docs/serverless/functions/reference-content/cron-schedules/) for more information).
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The CRON status.
        """
        return pulumi.get(self, "status")

