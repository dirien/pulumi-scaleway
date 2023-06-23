# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ObjectItemArgs', 'ObjectItem']

@pulumi.input_type
class ObjectItemArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 key: pulumi.Input[str],
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 hash: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ObjectItem resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] key: The path of the object.
        :param pulumi.Input[str] content: The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] content_base64: The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] file: The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] hash: Hash of the file, used to trigger upload on file change
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map of metadata used for the object, keys must be lowercase
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        :param pulumi.Input[str] storage_class: Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags
        :param pulumi.Input[str] visibility: Visibility of the object, `public-read` or `private`
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_base64 is not None:
            pulumi.set(__self__, "content_base64", content_base64)
        if file is not None:
            pulumi.set(__self__, "file", file)
        if hash is not None:
            pulumi.set(__self__, "hash", hash)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if storage_class is not None:
            pulumi.set(__self__, "storage_class", storage_class)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

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
    def key(self) -> pulumi.Input[str]:
        """
        The path of the object.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def hash(self) -> Optional[pulumi.Input[str]]:
        """
        Hash of the file, used to trigger upload on file change
        """
        return pulumi.get(self, "hash")

    @hash.setter
    def hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hash", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of metadata used for the object, keys must be lowercase
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the bucket is associated with.
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

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_class", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Visibility of the object, `public-read` or `private`
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


@pulumi.input_type
class _ObjectItemState:
    def __init__(__self__, *,
                 bucket: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 hash: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 visibility: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ObjectItem resources.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] content: The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] content_base64: The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] file: The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] hash: Hash of the file, used to trigger upload on file change
        :param pulumi.Input[str] key: The path of the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map of metadata used for the object, keys must be lowercase
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        :param pulumi.Input[str] storage_class: Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags
        :param pulumi.Input[str] visibility: Visibility of the object, `public-read` or `private`
        """
        if bucket is not None:
            pulumi.set(__self__, "bucket", bucket)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_base64 is not None:
            pulumi.set(__self__, "content_base64", content_base64)
        if file is not None:
            pulumi.set(__self__, "file", file)
        if hash is not None:
            pulumi.set(__self__, "hash", hash)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if storage_class is not None:
            pulumi.set(__self__, "storage_class", storage_class)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

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
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "content_base64")

    @content_base64.setter
    def content_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_base64", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file", value)

    @property
    @pulumi.getter
    def hash(self) -> Optional[pulumi.Input[str]]:
        """
        Hash of the file, used to trigger upload on file change
        """
        return pulumi.get(self, "hash")

    @hash.setter
    def hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hash", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The path of the object.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of metadata used for the object, keys must be lowercase
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        `project_id`) The ID of the project the bucket is associated with.
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

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_class", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Map of tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Visibility of the object, `public-read` or `private`
        """
        return pulumi.get(self, "visibility")

    @visibility.setter
    def visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "visibility", value)


class ObjectItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 hash: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates and manages Scaleway object storage objects.
        For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).

        ## Import

        Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:index/objectItem:ObjectItem some_object fr-par/some-bucket/some-file
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] content: The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] content_base64: The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] file: The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] hash: Hash of the file, used to trigger upload on file change
        :param pulumi.Input[str] key: The path of the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map of metadata used for the object, keys must be lowercase
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        :param pulumi.Input[str] storage_class: Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags
        :param pulumi.Input[str] visibility: Visibility of the object, `public-read` or `private`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ObjectItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates and manages Scaleway object storage objects.
        For more information, see [the documentation](https://www.scaleway.com/en/docs/object-storage-feature/).

        ## Import

        Objects can be imported using the `{region}/{bucketName}/{objectKey}` identifier, e.g. bash

        ```sh
         $ pulumi import scaleway:index/objectItem:ObjectItem some_object fr-par/some-bucket/some-file
        ```

        :param str resource_name: The name of the resource.
        :param ObjectItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ObjectItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_base64: Optional[pulumi.Input[str]] = None,
                 file: Optional[pulumi.Input[str]] = None,
                 hash: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 visibility: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ObjectItemArgs.__new__(ObjectItemArgs)

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__.__dict__["bucket"] = bucket
            __props__.__dict__["content"] = content
            __props__.__dict__["content_base64"] = content_base64
            __props__.__dict__["file"] = file
            __props__.__dict__["hash"] = hash
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["storage_class"] = storage_class
            __props__.__dict__["tags"] = tags
            __props__.__dict__["visibility"] = visibility
        super(ObjectItem, __self__).__init__(
            'scaleway:index/objectItem:ObjectItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            content: Optional[pulumi.Input[str]] = None,
            content_base64: Optional[pulumi.Input[str]] = None,
            file: Optional[pulumi.Input[str]] = None,
            hash: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            storage_class: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            visibility: Optional[pulumi.Input[str]] = None) -> 'ObjectItem':
        """
        Get an existing ObjectItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket.
        :param pulumi.Input[str] content: The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] content_base64: The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] file: The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        :param pulumi.Input[str] hash: Hash of the file, used to trigger upload on file change
        :param pulumi.Input[str] key: The path of the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Map of metadata used for the object, keys must be lowercase
        :param pulumi.Input[str] project_id: `project_id`) The ID of the project the bucket is associated with.
        :param pulumi.Input[str] region: The Scaleway region this bucket resides in.
        :param pulumi.Input[str] storage_class: Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Map of tags
        :param pulumi.Input[str] visibility: Visibility of the object, `public-read` or `private`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ObjectItemState.__new__(_ObjectItemState)

        __props__.__dict__["bucket"] = bucket
        __props__.__dict__["content"] = content
        __props__.__dict__["content_base64"] = content_base64
        __props__.__dict__["file"] = file
        __props__.__dict__["hash"] = hash
        __props__.__dict__["key"] = key
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["storage_class"] = storage_class
        __props__.__dict__["tags"] = tags
        __props__.__dict__["visibility"] = visibility
        return ObjectItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[Optional[str]]:
        """
        The content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentBase64")
    def content_base64(self) -> pulumi.Output[Optional[str]]:
        """
        The base64-encoded content of the file to upload. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "content_base64")

    @property
    @pulumi.getter
    def file(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the file to upload, defaults to an empty file. Only one of `file`, `content` or `content_base64` can be defined.
        """
        return pulumi.get(self, "file")

    @property
    @pulumi.getter
    def hash(self) -> pulumi.Output[Optional[str]]:
        """
        Hash of the file, used to trigger upload on file change
        """
        return pulumi.get(self, "hash")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The path of the object.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of metadata used for the object, keys must be lowercase
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        `project_id`) The ID of the project the bucket is associated with.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The Scaleway region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Scaleway [storage class](https://www.scaleway.com/en/docs/storage/object/concepts/#storage-class) `STANDARD`, `GLACIER`, `ONEZONE_IA` used to store the object.
        """
        return pulumi.get(self, "storage_class")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Map of tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def visibility(self) -> pulumi.Output[str]:
        """
        Visibility of the object, `public-read` or `private`
        """
        return pulumi.get(self, "visibility")

