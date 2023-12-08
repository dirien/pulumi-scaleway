# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ObjectBucketPolicyArgs', 'ObjectBucketPolicy']

@pulumi.input_type
class ObjectBucketPolicyArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 policy: pulumi.Input[str],
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ObjectBucketPolicy resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
               
               > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "policy", policy)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the bucket is associated with.

        > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ObjectBucketPolicyState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ObjectBucketPolicy resources.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
               
               > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def bucket(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the bucket is associated with.

        > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class ObjectBucketPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway object storage bucket policy.
        For more information, see [the documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import json

        bucket = scaleway.ObjectBucket("bucket")
        main = scaleway.IamApplication("main", description="a description")
        policy = scaleway.ObjectBucketPolicy("policy",
            bucket=bucket.id,
            policy=pulumi.Output.all(main.id, bucket.name, bucket.name).apply(lambda id, bucketName, bucketName1: json.dumps({
                "Version": "2023-04-17",
                "Id": "MyBucketPolicy",
                "Statement": [{
                    "Sid": "Delegate access",
                    "Effect": "Allow",
                    "Principal": {
                        "SCW": f"application_id:{id}",
                    },
                    "Action": "s3:ListBucket",
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```

        ## Import

        Buckets can be imported using the `{region}/{bucketName}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:index/objectBucketPolicy:ObjectBucketPolicy some_bucket fr-par/some-bucket
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
               
               > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ObjectBucketPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway object storage bucket policy.
        For more information, see [the documentation](https://www.scaleway.com/en/docs/storage/object/api-cli/bucket-policy/).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import json

        bucket = scaleway.ObjectBucket("bucket")
        main = scaleway.IamApplication("main", description="a description")
        policy = scaleway.ObjectBucketPolicy("policy",
            bucket=bucket.id,
            policy=pulumi.Output.all(main.id, bucket.name, bucket.name).apply(lambda id, bucketName, bucketName1: json.dumps({
                "Version": "2023-04-17",
                "Id": "MyBucketPolicy",
                "Statement": [{
                    "Sid": "Delegate access",
                    "Effect": "Allow",
                    "Principal": {
                        "SCW": f"application_id:{id}",
                    },
                    "Action": "s3:ListBucket",
                    "Resource": [
                        bucket_name,
                        f"{bucket_name1}/*",
                    ],
                }],
            })))
        ```

        ## Import

        Buckets can be imported using the `{region}/{bucketName}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:index/objectBucketPolicy:ObjectBucketPolicy some_bucket fr-par/some-bucket
        ```

        :param str resource_name: The name of the resource.
        :param ObjectBucketPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObjectBucketPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObjectBucketPolicyArgs.__new__(ObjectBucketPolicyArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
        super(ObjectBucketPolicy, __self__).__init__(
            'scaleway:index/objectBucketPolicy:ObjectBucketPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'ObjectBucketPolicy':
        """
        Get an existing ObjectBucketPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] policy: The text of the policy.
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
               
               > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObjectBucketPolicyState.__new__(_ObjectBucketPolicyState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["policy"] = policy
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        return ObjectBucketPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The text of the policy.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the bucket is associated with.

        > **Important:** The aws_iam_policy_document data source may be used, so long as it specifies a principal.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

