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

__all__ = ['ObjectBucketWebsiteConfigurationArgs', 'ObjectBucketWebsiteConfiguration']

@pulumi.input_type
class ObjectBucketWebsiteConfigurationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 index_document: pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs'],
                 error_document: Optional[pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs']] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ObjectBucketWebsiteConfiguration resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs'] index_document: The name of the index document for the website detailed below.
        :param pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs'] error_document: The name of the error document for the website detailed below.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "index_document", index_document)
        if error_document is not None:
            pulumi.set(__self__, "error_document", error_document)
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
    @pulumi.getter(name="indexDocument")
    def index_document(self) -> pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs']:
        """
        The name of the index document for the website detailed below.
        """
        return pulumi.get(self, "index_document")

    @index_document.setter
    def index_document(self, value: pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs']):
        pulumi.set(self, "index_document", value)

    @property
    @pulumi.getter(name="errorDocument")
    def error_document(self) -> Optional[pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]:
        """
        The name of the error document for the website detailed below.
        """
        return pulumi.get(self, "error_document")

    @error_document.setter
    def error_document(self, value: Optional[pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]):
        pulumi.set(self, "error_document", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ObjectBucketWebsiteConfigurationState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 error_document: Optional[pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs']] = None,
                 index_document: Optional[pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs']] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 website_domain: Optional[pulumi.Input[str]] = None,
                 website_endpoint: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ObjectBucketWebsiteConfiguration resources.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs'] error_document: The name of the error document for the website detailed below.
        :param pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs'] index_document: The name of the index document for the website detailed below.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] website_domain: The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
        :param pulumi.Input[str] website_endpoint: The website endpoint.
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if error_document is not None:
            pulumi.set(__self__, "error_document", error_document)
        if index_document is not None:
            pulumi.set(__self__, "index_document", index_document)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if website_domain is not None:
            pulumi.set(__self__, "website_domain", website_domain)
        if website_endpoint is not None:
            pulumi.set(__self__, "website_endpoint", website_endpoint)

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
    @pulumi.getter(name="errorDocument")
    def error_document(self) -> Optional[pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]:
        """
        The name of the error document for the website detailed below.
        """
        return pulumi.get(self, "error_document")

    @error_document.setter
    def error_document(self, value: Optional[pulumi.Input['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]):
        pulumi.set(self, "error_document", value)

    @property
    @pulumi.getter(name="indexDocument")
    def index_document(self) -> Optional[pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs']]:
        """
        The name of the index document for the website detailed below.
        """
        return pulumi.get(self, "index_document")

    @index_document.setter
    def index_document(self, value: Optional[pulumi.Input['ObjectBucketWebsiteConfigurationIndexDocumentArgs']]):
        pulumi.set(self, "index_document", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="websiteDomain")
    def website_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
        """
        return pulumi.get(self, "website_domain")

    @website_domain.setter
    def website_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website_domain", value)

    @property
    @pulumi.getter(name="websiteEndpoint")
    def website_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The website endpoint.
        """
        return pulumi.get(self, "website_endpoint")

    @website_endpoint.setter
    def website_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "website_endpoint", value)


class ObjectBucketWebsiteConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 error_document: Optional[pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]] = None,
                 index_document: Optional[pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationIndexDocumentArgs']]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an Object bucket website configuration resource.
        For more information, see [Hosting Websites on Object bucket](https://www.scaleway.com/en/docs/storage/object/how-to/use-bucket-website/).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_object_bucket = scaleway.ObjectBucket("mainObjectBucket", acl="public-read")
        main_object_bucket_website_configuration = scaleway.ObjectBucketWebsiteConfiguration("mainObjectBucketWebsiteConfiguration",
            bucket=main_object_bucket.id,
            index_document=scaleway.ObjectBucketWebsiteConfigurationIndexDocumentArgs(
                suffix="index.html",
            ))
        ```

        ### With `Policy`

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import json

        main_object_bucket = scaleway.ObjectBucket("mainObjectBucket", acl="public-read")
        main_object_bucket_policy = scaleway.ObjectBucketPolicy("mainObjectBucketPolicy",
            bucket=main_object_bucket.id,
            policy=json.dumps({
                "Version": "2012-10-17",
                "Id": "MyPolicy",
                "Statement": [{
                    "Sid": "GrantToEveryone",
                    "Effect": "Allow",
                    "Principal": "*",
                    "Action": ["s3:GetObject"],
                    "Resource": ["<bucket-name>/*"],
                }],
            }))
        main_object_bucket_website_configuration = scaleway.ObjectBucketWebsiteConfiguration("mainObjectBucketWebsiteConfiguration",
            bucket=main_object_bucket.id,
            index_document=scaleway.ObjectBucketWebsiteConfigurationIndexDocumentArgs(
                suffix="index.html",
            ))
        ```

        ## Import

        Bucket website configurations can be imported using the `{region}/{bucketName}` identifier, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket
        ```

        ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.

        If you are using a project different from the default one, you have to specify the project ID at the end of the import command.

        bash

        ```sh
        $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationErrorDocumentArgs']] error_document: The name of the error document for the website detailed below.
        :param pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationIndexDocumentArgs']] index_document: The name of the index document for the website detailed below.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ObjectBucketWebsiteConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Object bucket website configuration resource.
        For more information, see [Hosting Websites on Object bucket](https://www.scaleway.com/en/docs/storage/object/how-to/use-bucket-website/).

        ## Example Usage

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        main_object_bucket = scaleway.ObjectBucket("mainObjectBucket", acl="public-read")
        main_object_bucket_website_configuration = scaleway.ObjectBucketWebsiteConfiguration("mainObjectBucketWebsiteConfiguration",
            bucket=main_object_bucket.id,
            index_document=scaleway.ObjectBucketWebsiteConfigurationIndexDocumentArgs(
                suffix="index.html",
            ))
        ```

        ### With `Policy`

        ```python
        import pulumi
        import ediri_scaleway as scaleway
        import json

        main_object_bucket = scaleway.ObjectBucket("mainObjectBucket", acl="public-read")
        main_object_bucket_policy = scaleway.ObjectBucketPolicy("mainObjectBucketPolicy",
            bucket=main_object_bucket.id,
            policy=json.dumps({
                "Version": "2012-10-17",
                "Id": "MyPolicy",
                "Statement": [{
                    "Sid": "GrantToEveryone",
                    "Effect": "Allow",
                    "Principal": "*",
                    "Action": ["s3:GetObject"],
                    "Resource": ["<bucket-name>/*"],
                }],
            }))
        main_object_bucket_website_configuration = scaleway.ObjectBucketWebsiteConfiguration("mainObjectBucketWebsiteConfiguration",
            bucket=main_object_bucket.id,
            index_document=scaleway.ObjectBucketWebsiteConfigurationIndexDocumentArgs(
                suffix="index.html",
            ))
        ```

        ## Import

        Bucket website configurations can be imported using the `{region}/{bucketName}` identifier, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket
        ```

        ~> **Important:** The `project_id` attribute has a particular behavior with s3 products because the s3 API is scoped by project.

        If you are using a project different from the default one, you have to specify the project ID at the end of the import command.

        bash

        ```sh
        $ pulumi import scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration some_bucket fr-par/some-bucket@xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param ObjectBucketWebsiteConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObjectBucketWebsiteConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 error_document: Optional[pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]] = None,
                 index_document: Optional[pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationIndexDocumentArgs']]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObjectBucketWebsiteConfigurationArgs.__new__(ObjectBucketWebsiteConfigurationArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["error_document"] = error_document
            if index_document is None and not opts.urn:
                raise TypeError("Missing required property 'index_document'")
            __props__.__dict__["index_document"] = index_document
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["website_domain"] = None
            __props__.__dict__["website_endpoint"] = None
        super(ObjectBucketWebsiteConfiguration, __self__).__init__(
            'scaleway:index/objectBucketWebsiteConfiguration:ObjectBucketWebsiteConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            error_document: Optional[pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationErrorDocumentArgs']]] = None,
            index_document: Optional[pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationIndexDocumentArgs']]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            website_domain: Optional[pulumi.Input[str]] = None,
            website_endpoint: Optional[pulumi.Input[str]] = None) -> 'ObjectBucketWebsiteConfiguration':
        """
        Get an existing ObjectBucketWebsiteConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationErrorDocumentArgs']] error_document: The name of the error document for the website detailed below.
        :param pulumi.Input[pulumi.InputType['ObjectBucketWebsiteConfigurationIndexDocumentArgs']] index_document: The name of the index document for the website detailed below.
        :param pulumi.Input[str] project_id: The project_id you want to attach the resource to
        :param pulumi.Input[str] region: The region you want to attach the resource to
        :param pulumi.Input[str] website_domain: The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
        :param pulumi.Input[str] website_endpoint: The website endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObjectBucketWebsiteConfigurationState.__new__(_ObjectBucketWebsiteConfigurationState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["error_document"] = error_document
        __props__.__dict__["index_document"] = index_document
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["website_domain"] = website_domain
        __props__.__dict__["website_endpoint"] = website_endpoint
        return ObjectBucketWebsiteConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="errorDocument")
    def error_document(self) -> pulumi.Output[Optional['outputs.ObjectBucketWebsiteConfigurationErrorDocument']]:
        """
        The name of the error document for the website detailed below.
        """
        return pulumi.get(self, "error_document")

    @property
    @pulumi.getter(name="indexDocument")
    def index_document(self) -> pulumi.Output['outputs.ObjectBucketWebsiteConfigurationIndexDocument']:
        """
        The name of the index document for the website detailed below.
        """
        return pulumi.get(self, "index_document")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project_id you want to attach the resource to
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region you want to attach the resource to
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="websiteDomain")
    def website_domain(self) -> pulumi.Output[str]:
        """
        The domain of the website endpoint. This is used to create DNS alias [records](https://www.scaleway.com/en/docs/network/domains-and-dns/how-to/manage-dns-records/).
        """
        return pulumi.get(self, "website_domain")

    @property
    @pulumi.getter(name="websiteEndpoint")
    def website_endpoint(self) -> pulumi.Output[str]:
        """
        The website endpoint.
        """
        return pulumi.get(self, "website_endpoint")

