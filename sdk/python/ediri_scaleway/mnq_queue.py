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

__all__ = ['MnqQueueArgs', 'MnqQueue']

@pulumi.input_type
class MnqQueueArgs:
    def __init__(__self__, *,
                 namespace_id: pulumi.Input[str],
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input['MnqQueueNatsArgs']] = None,
                 sqs: Optional[pulumi.Input['MnqQueueSqsArgs']] = None):
        """
        The set of arguments for constructing a MnqQueue resource.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace associated to
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes.
        :param pulumi.Input[str] name: The name of the queue. Conflicts with name_prefix.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with name.
        :param pulumi.Input['MnqQueueNatsArgs'] nats: The NATS attributes of the queue
        :param pulumi.Input['MnqQueueSqsArgs'] sqs: The SQS attributes of the queue
        """
        pulumi.set(__self__, "namespace_id", namespace_id)
        if message_max_age is not None:
            pulumi.set(__self__, "message_max_age", message_max_age)
        if message_max_size is not None:
            pulumi.set(__self__, "message_max_size", message_max_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if nats is not None:
            pulumi.set(__self__, "nats", nats)
        if sqs is not None:
            pulumi.set(__self__, "sqs", sqs)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Input[str]:
        """
        The ID of the Namespace associated to
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="messageMaxAge")
    def message_max_age(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds the queue retains a message.
        """
        return pulumi.get(self, "message_max_age")

    @message_max_age.setter
    def message_max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_age", value)

    @property
    @pulumi.getter(name="messageMaxSize")
    def message_max_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of a message. Should be in bytes.
        """
        return pulumi.get(self, "message_max_size")

    @message_max_size.setter
    def message_max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the queue. Conflicts with name_prefix.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with name.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter
    def nats(self) -> Optional[pulumi.Input['MnqQueueNatsArgs']]:
        """
        The NATS attributes of the queue
        """
        return pulumi.get(self, "nats")

    @nats.setter
    def nats(self, value: Optional[pulumi.Input['MnqQueueNatsArgs']]):
        pulumi.set(self, "nats", value)

    @property
    @pulumi.getter
    def sqs(self) -> Optional[pulumi.Input['MnqQueueSqsArgs']]:
        """
        The SQS attributes of the queue
        """
        return pulumi.get(self, "sqs")

    @sqs.setter
    def sqs(self, value: Optional[pulumi.Input['MnqQueueSqsArgs']]):
        pulumi.set(self, "sqs", value)


@pulumi.input_type
class _MnqQueueState:
    def __init__(__self__, *,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input['MnqQueueNatsArgs']] = None,
                 sqs: Optional[pulumi.Input['MnqQueueSqsArgs']] = None):
        """
        Input properties used for looking up and filtering MnqQueue resources.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes.
        :param pulumi.Input[str] name: The name of the queue. Conflicts with name_prefix.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with name.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace associated to
        :param pulumi.Input['MnqQueueNatsArgs'] nats: The NATS attributes of the queue
        :param pulumi.Input['MnqQueueSqsArgs'] sqs: The SQS attributes of the queue
        """
        if message_max_age is not None:
            pulumi.set(__self__, "message_max_age", message_max_age)
        if message_max_size is not None:
            pulumi.set(__self__, "message_max_size", message_max_size)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if nats is not None:
            pulumi.set(__self__, "nats", nats)
        if sqs is not None:
            pulumi.set(__self__, "sqs", sqs)

    @property
    @pulumi.getter(name="messageMaxAge")
    def message_max_age(self) -> Optional[pulumi.Input[int]]:
        """
        The number of seconds the queue retains a message.
        """
        return pulumi.get(self, "message_max_age")

    @message_max_age.setter
    def message_max_age(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_age", value)

    @property
    @pulumi.getter(name="messageMaxSize")
    def message_max_size(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum size of a message. Should be in bytes.
        """
        return pulumi.get(self, "message_max_size")

    @message_max_size.setter
    def message_max_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "message_max_size", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the queue. Conflicts with name_prefix.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with name.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Namespace associated to
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter
    def nats(self) -> Optional[pulumi.Input['MnqQueueNatsArgs']]:
        """
        The NATS attributes of the queue
        """
        return pulumi.get(self, "nats")

    @nats.setter
    def nats(self, value: Optional[pulumi.Input['MnqQueueNatsArgs']]):
        pulumi.set(self, "nats", value)

    @property
    @pulumi.getter
    def sqs(self) -> Optional[pulumi.Input['MnqQueueSqsArgs']]:
        """
        The SQS attributes of the queue
        """
        return pulumi.get(self, "sqs")

    @sqs.setter
    def sqs(self, value: Optional[pulumi.Input['MnqQueueSqsArgs']]):
        pulumi.set(self, "sqs", value)


class MnqQueue(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input[pulumi.InputType['MnqQueueNatsArgs']]] = None,
                 sqs: Optional[pulumi.Input[pulumi.InputType['MnqQueueSqsArgs']]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway Messaging and Queuing queues.

        For more information about MNQ, see [the documentation](https://www.scaleway.com/en/developers/api/messaging-and-queuing/).

        ## Examples

        ### NATS

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_mnq_namespace = scaleway.MnqNamespace("mainMnqNamespace", protocol="nats")
        main_mnq_credential = scaleway.MnqCredential("mainMnqCredential", namespace_id=main_mnq_namespace.id)
        my_queue = scaleway.MnqQueue("myQueue",
            namespace_id=main_mnq_namespace.id,
            nats=scaleway.MnqQueueNatsArgs(
                credentials=main_mnq_credential.nats_credentials.content,
            ))
        ```

        ### SQS

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_mnq_namespace = scaleway.MnqNamespace("mainMnqNamespace", protocol="sqs_sns")
        main_mnq_credential = scaleway.MnqCredential("mainMnqCredential",
            namespace_id=main_mnq_namespace.id,
            sqs_sns_credentials=scaleway.MnqCredentialSqsSnsCredentialsArgs(
                permissions=scaleway.MnqCredentialSqsSnsCredentialsPermissionsArgs(
                    can_publish=True,
                    can_receive=True,
                    can_manage=True,
                ),
            ))
        my_queue = scaleway.MnqQueue("myQueue",
            namespace_id=main_mnq_namespace.id,
            sqs=scaleway.MnqQueueSqsArgs(
                access_key=main_mnq_credential.sqs_sns_credentials.access_key,
                secret_key=main_mnq_credential.sqs_sns_credentials.secret_key,
            ))
        ```

        ### Argument Reference

        The following arguments are supported:

        * `namespace_id` - (Required) The ID of the Namespace associated to.

        * `name` - (Optional) The name of the queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.

        * `name_prefix` - (Optional) Creates a unique name beginning with the specified prefix. Conflicts with `name`.

        * `message_max_age` - (Optional) The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.

        * `message_max_size` - (Optional) The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.

        * `sqs` - (Optional) The SQS attributes of the queue. Conflicts with `nats`.
            - `endpoint` - (Optional) The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
            - `access_key` - (Required) The access key of the SQS queue.
            - `secret_key` - (Required) The secret key of the SQS queue.
            - `fifo_queue` - (Optional) Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
            - `content_based_deduplication` - (Optional) Specifies whether to enable content-based deduplication. Defaults to `false`.
            - `receive_wait_time_seconds` - (Optional) The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
            - `visibility_timeout_seconds` - (Optional) The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
            - For more information about the SQS limitations, see [the documentation](https://www.scaleway.com/en/developers/api/messaging-and-queuing/#technical-limitations).

        * `nats` - (Optional) The NATS attributes of the queue. Conflicts with `sqs`.
            - `endpoint` - (Optional) The endpoint of the NATS queue. Can contain a {region} placeholder. Defaults to `nats://nats.mnq.{region}.scw.cloud:4222`.
            - `credentials` - (Required) Line jump separated key and seed.
            - `retention_policy` - (Optional) The retention policy of the queue. See https://docs.nats.io/nats-concepts/jetstream/streams#retentionpolicy for more information. Defaults to `workqueue`.

        ### Attribute Reference

        In addition to all arguments above, the following attributes are exported:

        * `sqs` - The SQS attributes of the queue.
          ~ `url` - The URL of the queue.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes.
        :param pulumi.Input[str] name: The name of the queue. Conflicts with name_prefix.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with name.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace associated to
        :param pulumi.Input[pulumi.InputType['MnqQueueNatsArgs']] nats: The NATS attributes of the queue
        :param pulumi.Input[pulumi.InputType['MnqQueueSqsArgs']] sqs: The SQS attributes of the queue
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqQueueArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway Messaging and Queuing queues.

        For more information about MNQ, see [the documentation](https://www.scaleway.com/en/developers/api/messaging-and-queuing/).

        ## Examples

        ### NATS

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_mnq_namespace = scaleway.MnqNamespace("mainMnqNamespace", protocol="nats")
        main_mnq_credential = scaleway.MnqCredential("mainMnqCredential", namespace_id=main_mnq_namespace.id)
        my_queue = scaleway.MnqQueue("myQueue",
            namespace_id=main_mnq_namespace.id,
            nats=scaleway.MnqQueueNatsArgs(
                credentials=main_mnq_credential.nats_credentials.content,
            ))
        ```

        ### SQS

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_mnq_namespace = scaleway.MnqNamespace("mainMnqNamespace", protocol="sqs_sns")
        main_mnq_credential = scaleway.MnqCredential("mainMnqCredential",
            namespace_id=main_mnq_namespace.id,
            sqs_sns_credentials=scaleway.MnqCredentialSqsSnsCredentialsArgs(
                permissions=scaleway.MnqCredentialSqsSnsCredentialsPermissionsArgs(
                    can_publish=True,
                    can_receive=True,
                    can_manage=True,
                ),
            ))
        my_queue = scaleway.MnqQueue("myQueue",
            namespace_id=main_mnq_namespace.id,
            sqs=scaleway.MnqQueueSqsArgs(
                access_key=main_mnq_credential.sqs_sns_credentials.access_key,
                secret_key=main_mnq_credential.sqs_sns_credentials.secret_key,
            ))
        ```

        ### Argument Reference

        The following arguments are supported:

        * `namespace_id` - (Required) The ID of the Namespace associated to.

        * `name` - (Optional) The name of the queue. Either `name` or `name_prefix` is required. Conflicts with `name_prefix`.

        * `name_prefix` - (Optional) Creates a unique name beginning with the specified prefix. Conflicts with `name`.

        * `message_max_age` - (Optional) The number of seconds the queue retains a message. Must be between 60 and 1_209_600. Defaults to 345_600.

        * `message_max_size` - (Optional) The maximum size of a message. Should be in bytes. Must be between 1024 and 262_144. Defaults to 262_144.

        * `sqs` - (Optional) The SQS attributes of the queue. Conflicts with `nats`.
            - `endpoint` - (Optional) The endpoint of the SQS queue. Can contain a {region} placeholder. Defaults to `http://sqs-sns.mnq.{region}.scw.cloud`.
            - `access_key` - (Required) The access key of the SQS queue.
            - `secret_key` - (Required) The secret key of the SQS queue.
            - `fifo_queue` - (Optional) Whether the queue is a FIFO queue. If true, the queue name must end with .fifo. Defaults to `false`.
            - `content_based_deduplication` - (Optional) Specifies whether to enable content-based deduplication. Defaults to `false`.
            - `receive_wait_time_seconds` - (Optional) The number of seconds to wait for a message to arrive in the queue before returning. Must be between 0 and 20. Defaults to 0.
            - `visibility_timeout_seconds` - (Optional) The number of seconds a message is hidden from other consumers. Must be between 0 and 43_200. Defaults to 30.
            - For more information about the SQS limitations, see [the documentation](https://www.scaleway.com/en/developers/api/messaging-and-queuing/#technical-limitations).

        * `nats` - (Optional) The NATS attributes of the queue. Conflicts with `sqs`.
            - `endpoint` - (Optional) The endpoint of the NATS queue. Can contain a {region} placeholder. Defaults to `nats://nats.mnq.{region}.scw.cloud:4222`.
            - `credentials` - (Required) Line jump separated key and seed.
            - `retention_policy` - (Optional) The retention policy of the queue. See https://docs.nats.io/nats-concepts/jetstream/streams#retentionpolicy for more information. Defaults to `workqueue`.

        ### Attribute Reference

        In addition to all arguments above, the following attributes are exported:

        * `sqs` - The SQS attributes of the queue.
          ~ `url` - The URL of the queue.

        :param str resource_name: The name of the resource.
        :param MnqQueueArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqQueueArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 message_max_age: Optional[pulumi.Input[int]] = None,
                 message_max_size: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 nats: Optional[pulumi.Input[pulumi.InputType['MnqQueueNatsArgs']]] = None,
                 sqs: Optional[pulumi.Input[pulumi.InputType['MnqQueueSqsArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqQueueArgs.__new__(MnqQueueArgs)

            __props__.__dict__["message_max_age"] = message_max_age
            __props__.__dict__["message_max_size"] = message_max_size
            __props__.__dict__["name"] = name
            __props__.__dict__["name_prefix"] = name_prefix
            if namespace_id is None and not opts.urn:
                raise TypeError("Missing required property 'namespace_id'")
            __props__.__dict__["namespace_id"] = namespace_id
            __props__.__dict__["nats"] = nats
            __props__.__dict__["sqs"] = sqs
        super(MnqQueue, __self__).__init__(
            'scaleway:index/mnqQueue:MnqQueue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            message_max_age: Optional[pulumi.Input[int]] = None,
            message_max_size: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            nats: Optional[pulumi.Input[pulumi.InputType['MnqQueueNatsArgs']]] = None,
            sqs: Optional[pulumi.Input[pulumi.InputType['MnqQueueSqsArgs']]] = None) -> 'MnqQueue':
        """
        Get an existing MnqQueue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] message_max_age: The number of seconds the queue retains a message.
        :param pulumi.Input[int] message_max_size: The maximum size of a message. Should be in bytes.
        :param pulumi.Input[str] name: The name of the queue. Conflicts with name_prefix.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with name.
        :param pulumi.Input[str] namespace_id: The ID of the Namespace associated to
        :param pulumi.Input[pulumi.InputType['MnqQueueNatsArgs']] nats: The NATS attributes of the queue
        :param pulumi.Input[pulumi.InputType['MnqQueueSqsArgs']] sqs: The SQS attributes of the queue
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqQueueState.__new__(_MnqQueueState)

        __props__.__dict__["message_max_age"] = message_max_age
        __props__.__dict__["message_max_size"] = message_max_size
        __props__.__dict__["name"] = name
        __props__.__dict__["name_prefix"] = name_prefix
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["nats"] = nats
        __props__.__dict__["sqs"] = sqs
        return MnqQueue(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="messageMaxAge")
    def message_max_age(self) -> pulumi.Output[Optional[int]]:
        """
        The number of seconds the queue retains a message.
        """
        return pulumi.get(self, "message_max_age")

    @property
    @pulumi.getter(name="messageMaxSize")
    def message_max_size(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum size of a message. Should be in bytes.
        """
        return pulumi.get(self, "message_max_size")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the queue. Conflicts with name_prefix.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the specified prefix. Conflicts with name.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        The ID of the Namespace associated to
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def nats(self) -> pulumi.Output[Optional['outputs.MnqQueueNats']]:
        """
        The NATS attributes of the queue
        """
        return pulumi.get(self, "nats")

    @property
    @pulumi.getter
    def sqs(self) -> pulumi.Output[Optional['outputs.MnqQueueSqs']]:
        """
        The SQS attributes of the queue
        """
        return pulumi.get(self, "sqs")
