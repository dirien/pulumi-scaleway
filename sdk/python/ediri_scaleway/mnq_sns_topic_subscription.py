# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MnqSnsTopicSubscriptionArgs', 'MnqSnsTopicSubscription']

@pulumi.input_type
class MnqSnsTopicSubscriptionArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 endpoint: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 redrive_policy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MnqSnsTopicSubscription resource.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] protocol: Protocol of the SNS topic subscription.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] endpoint: Endpoint of the subscription
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[bool] redrive_policy: Activate JSON redrive policy.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        :param pulumi.Input[str] topic_arn: The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        :param pulumi.Input[str] topic_id: The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "secret_key", secret_key)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if redrive_policy is not None:
            pulumi.set(__self__, "redrive_policy", redrive_policy)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if sns_endpoint is not None:
            pulumi.set(__self__, "sns_endpoint", sns_endpoint)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)
        if topic_id is not None:
            pulumi.set(__self__, "topic_id", topic_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The access key of the SNS credentials.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        """
        Protocol of the SNS topic subscription.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        The secret key of the SNS credentials.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint of the subscription
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="redrivePolicy")
    def redrive_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Activate JSON redrive policy.
        """
        return pulumi.get(self, "redrive_policy")

    @redrive_policy.setter
    def redrive_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "redrive_policy", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="snsEndpoint")
    def sns_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sns_endpoint")

    @sns_endpoint.setter
    def sns_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_endpoint", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_id", value)


@pulumi.input_type
class _MnqSnsTopicSubscriptionState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 redrive_policy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MnqSnsTopicSubscription resources.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] arn: The ARN of the topic subscription
        :param pulumi.Input[str] endpoint: Endpoint of the subscription
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] protocol: Protocol of the SNS topic subscription.
        :param pulumi.Input[bool] redrive_policy: Activate JSON redrive policy.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        :param pulumi.Input[str] topic_arn: The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        :param pulumi.Input[str] topic_id: The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if redrive_policy is not None:
            pulumi.set(__self__, "redrive_policy", redrive_policy)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sns_endpoint is not None:
            pulumi.set(__self__, "sns_endpoint", sns_endpoint)
        if topic_arn is not None:
            pulumi.set(__self__, "topic_arn", topic_arn)
        if topic_id is not None:
            pulumi.set(__self__, "topic_id", topic_id)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key of the SNS credentials.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the topic subscription
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Endpoint of the subscription
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Protocol of the SNS topic subscription.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="redrivePolicy")
    def redrive_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Activate JSON redrive policy.
        """
        return pulumi.get(self, "redrive_policy")

    @redrive_policy.setter
    def redrive_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "redrive_policy", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        `region`). The region
        in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key of the SNS credentials.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="snsEndpoint")
    def sns_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sns_endpoint")

    @sns_endpoint.setter
    def sns_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sns_endpoint", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_arn", value)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_id", value)


class MnqSnsTopicSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 redrive_policy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages Scaleway Messaging and Queuing SNS topic subscriptions.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        # For default project in default region
        main_mnq_sns = scaleway.MnqSns("mainMnqSns")
        main_mnq_sns_credentials = scaleway.MnqSnsCredentials("mainMnqSnsCredentials",
            project_id=main_mnq_sns.project_id,
            permissions=scaleway.MnqSnsCredentialsPermissionsArgs(
                can_manage=True,
                can_publish=True,
                can_receive=True,
            ))
        topic = scaleway.MnqSnsTopic("topic",
            project_id=main_mnq_sns.project_id,
            access_key=main_mnq_sns_credentials.access_key,
            secret_key=main_mnq_sns_credentials.secret_key)
        main_mnq_sns_topic_subscription = scaleway.MnqSnsTopicSubscription("mainMnqSnsTopicSubscription",
            project_id=main_mnq_sns.project_id,
            access_key=main_mnq_sns_credentials.access_key,
            secret_key=main_mnq_sns_credentials.secret_key,
            topic_id=topic.id,
            protocol="http",
            endpoint="http://example.com")
        ```

        ## Import

        SNS topic subscriptions can be imported using `{region}/{project-id}/{topic-name}/{subscription-id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription main fr-par/11111111111111111111111111111111/my-topic/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] endpoint: Endpoint of the subscription
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] protocol: Protocol of the SNS topic subscription.
        :param pulumi.Input[bool] redrive_policy: Activate JSON redrive policy.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        :param pulumi.Input[str] topic_arn: The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        :param pulumi.Input[str] topic_id: The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MnqSnsTopicSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages Scaleway Messaging and Queuing SNS topic subscriptions.
        For further information, see
        our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sns-overview/).

        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        # For default project in default region
        main_mnq_sns = scaleway.MnqSns("mainMnqSns")
        main_mnq_sns_credentials = scaleway.MnqSnsCredentials("mainMnqSnsCredentials",
            project_id=main_mnq_sns.project_id,
            permissions=scaleway.MnqSnsCredentialsPermissionsArgs(
                can_manage=True,
                can_publish=True,
                can_receive=True,
            ))
        topic = scaleway.MnqSnsTopic("topic",
            project_id=main_mnq_sns.project_id,
            access_key=main_mnq_sns_credentials.access_key,
            secret_key=main_mnq_sns_credentials.secret_key)
        main_mnq_sns_topic_subscription = scaleway.MnqSnsTopicSubscription("mainMnqSnsTopicSubscription",
            project_id=main_mnq_sns.project_id,
            access_key=main_mnq_sns_credentials.access_key,
            secret_key=main_mnq_sns_credentials.secret_key,
            topic_id=topic.id,
            protocol="http",
            endpoint="http://example.com")
        ```

        ## Import

        SNS topic subscriptions can be imported using `{region}/{project-id}/{topic-name}/{subscription-id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription main fr-par/11111111111111111111111111111111/my-topic/11111111111111111111111111111111
        ```

        :param str resource_name: The name of the resource.
        :param MnqSnsTopicSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MnqSnsTopicSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 redrive_policy: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sns_endpoint: Optional[pulumi.Input[str]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MnqSnsTopicSubscriptionArgs.__new__(MnqSnsTopicSubscriptionArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = None if access_key is None else pulumi.Output.secret(access_key)
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["project_id"] = project_id
            if protocol is None and not opts.urn:
                raise TypeError("Missing required property 'protocol'")
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["redrive_policy"] = redrive_policy
            __props__.__dict__["region"] = region
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["sns_endpoint"] = sns_endpoint
            __props__.__dict__["topic_arn"] = topic_arn
            __props__.__dict__["topic_id"] = topic_id
            __props__.__dict__["arn"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MnqSnsTopicSubscription, __self__).__init__(
            'scaleway:index/mnqSnsTopicSubscription:MnqSnsTopicSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            redrive_policy: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            sns_endpoint: Optional[pulumi.Input[str]] = None,
            topic_arn: Optional[pulumi.Input[str]] = None,
            topic_id: Optional[pulumi.Input[str]] = None) -> 'MnqSnsTopicSubscription':
        """
        Get an existing MnqSnsTopicSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key of the SNS credentials.
        :param pulumi.Input[str] arn: The ARN of the topic subscription
        :param pulumi.Input[str] endpoint: Endpoint of the subscription
        :param pulumi.Input[str] project_id: `project_id`) The ID of the Project in which SNS is enabled.
        :param pulumi.Input[str] protocol: Protocol of the SNS topic subscription.
        :param pulumi.Input[bool] redrive_policy: Activate JSON redrive policy.
        :param pulumi.Input[str] region: `region`). The region
               in which SNS is enabled.
        :param pulumi.Input[str] secret_key: The secret key of the SNS credentials.
        :param pulumi.Input[str] sns_endpoint: The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        :param pulumi.Input[str] topic_arn: The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        :param pulumi.Input[str] topic_id: The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MnqSnsTopicSubscriptionState.__new__(_MnqSnsTopicSubscriptionState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["arn"] = arn
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["redrive_policy"] = redrive_policy
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["sns_endpoint"] = sns_endpoint
        __props__.__dict__["topic_arn"] = topic_arn
        __props__.__dict__["topic_id"] = topic_id
        return MnqSnsTopicSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The access key of the SNS credentials.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the topic subscription
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Endpoint of the subscription
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the Project in which SNS is enabled.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Protocol of the SNS topic subscription.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="redrivePolicy")
    def redrive_policy(self) -> pulumi.Output[bool]:
        """
        Activate JSON redrive policy.
        """
        return pulumi.get(self, "redrive_policy")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        `region`). The region
        in which SNS is enabled.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The secret key of the SNS credentials.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="snsEndpoint")
    def sns_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The endpoint of the SNS service. Can contain a {region} placeholder. Defaults to `https://sns.mnq.{region}.scaleway.com`.
        """
        return pulumi.get(self, "sns_endpoint")

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the topic. Either `topic_id` or `topic_arn` is required.
        """
        return pulumi.get(self, "topic_arn")

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the topic. Either `topic_id` or `topic_arn` is required. Conflicts with `topic_arn`.
        """
        return pulumi.get(self, "topic_id")

