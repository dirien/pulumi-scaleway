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

__all__ = ['K8sAclArgs', 'K8sAcl']

@pulumi.input_type
class K8sAclArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[builtins.str],
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]]] = None,
                 no_ip_allowed: Optional[pulumi.Input[builtins.bool]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a K8sAcl resource.
        :param pulumi.Input[builtins.str] cluster_id: UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.
               
               > **Important:** Updates to `cluster_id` will recreate the ACL.
        :param pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]] acl_rules: A list of ACLs (structure is described below)
               
               > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        :param pulumi.Input[builtins.bool] no_ip_allowed: If set to true, no IP will be allowed and the cluster will be in full-isolation.
               
               > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the ACL rule should be created.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if acl_rules is not None:
            pulumi.set(__self__, "acl_rules", acl_rules)
        if no_ip_allowed is not None:
            pulumi.set(__self__, "no_ip_allowed", no_ip_allowed)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[builtins.str]:
        """
        UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.

        > **Important:** Updates to `cluster_id` will recreate the ACL.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]]]:
        """
        A list of ACLs (structure is described below)

        > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        """
        return pulumi.get(self, "acl_rules")

    @acl_rules.setter
    def acl_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]]]):
        pulumi.set(self, "acl_rules", value)

    @property
    @pulumi.getter(name="noIpAllowed")
    def no_ip_allowed(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set to true, no IP will be allowed and the cluster will be in full-isolation.

        > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        """
        return pulumi.get(self, "no_ip_allowed")

    @no_ip_allowed.setter
    def no_ip_allowed(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "no_ip_allowed", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `region`) The region in which the ACL rule should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _K8sAclState:
    def __init__(__self__, *,
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]]] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 no_ip_allowed: Optional[pulumi.Input[builtins.bool]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering K8sAcl resources.
        :param pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]] acl_rules: A list of ACLs (structure is described below)
               
               > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        :param pulumi.Input[builtins.str] cluster_id: UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.
               
               > **Important:** Updates to `cluster_id` will recreate the ACL.
        :param pulumi.Input[builtins.bool] no_ip_allowed: If set to true, no IP will be allowed and the cluster will be in full-isolation.
               
               > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the ACL rule should be created.
        """
        if acl_rules is not None:
            pulumi.set(__self__, "acl_rules", acl_rules)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if no_ip_allowed is not None:
            pulumi.set(__self__, "no_ip_allowed", no_ip_allowed)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]]]:
        """
        A list of ACLs (structure is described below)

        > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        """
        return pulumi.get(self, "acl_rules")

    @acl_rules.setter
    def acl_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['K8sAclAclRuleArgs']]]]):
        pulumi.set(self, "acl_rules", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.

        > **Important:** Updates to `cluster_id` will recreate the ACL.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="noIpAllowed")
    def no_ip_allowed(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set to true, no IP will be allowed and the cluster will be in full-isolation.

        > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        """
        return pulumi.get(self, "no_ip_allowed")

    @no_ip_allowed.setter
    def no_ip_allowed(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "no_ip_allowed", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        `region`) The region in which the ACL rule should be created.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "region", value)


@pulumi.type_token("scaleway:index/k8sAcl:K8sAcl")
class K8sAcl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['K8sAclAclRuleArgs', 'K8sAclAclRuleArgsDict']]]]] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 no_ip_allowed: Optional[pulumi.Input[builtins.bool]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        acl_basic_vpc_private_network = scaleway.VpcPrivateNetwork("aclBasicVpcPrivateNetwork")
        acl_basic_k8s_cluster = scaleway.K8sCluster("aclBasicK8sCluster",
            version="1.32.2",
            cni="cilium",
            delete_additional_resources=True,
            private_network_id=acl_basic_vpc_private_network.id)
        acl_basic_k8s_acl = scaleway.K8sAcl("aclBasicK8sAcl",
            cluster_id=acl_basic_k8s_cluster.id,
            acl_rules=[
                {
                    "ip": "1.2.3.4/32",
                    "description": "Allow 1.2.3.4",
                },
                {
                    "scaleway_ranges": True,
                    "description": "Allow all Scaleway ranges",
                },
            ])
        ```

        ### Full-isolation

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        acl_basic_vpc_private_network = scaleway.VpcPrivateNetwork("aclBasicVpcPrivateNetwork")
        acl_basic_k8s_cluster = scaleway.K8sCluster("aclBasicK8sCluster",
            version="1.32.2",
            cni="cilium",
            delete_additional_resources=True,
            private_network_id=acl_basic_vpc_private_network.id)
        acl_basic_k8s_acl = scaleway.K8sAcl("aclBasicK8sAcl",
            cluster_id=acl_basic_k8s_cluster.id,
            no_ip_allowed=True)
        ```

        ## Import

        Kubernetes ACLs can be imported using the `{region}/{cluster-id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/k8sAcl:K8sAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['K8sAclAclRuleArgs', 'K8sAclAclRuleArgsDict']]]] acl_rules: A list of ACLs (structure is described below)
               
               > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        :param pulumi.Input[builtins.str] cluster_id: UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.
               
               > **Important:** Updates to `cluster_id` will recreate the ACL.
        :param pulumi.Input[builtins.bool] no_ip_allowed: If set to true, no IP will be allowed and the cluster will be in full-isolation.
               
               > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the ACL rule should be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: K8sAclArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Basic

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        acl_basic_vpc_private_network = scaleway.VpcPrivateNetwork("aclBasicVpcPrivateNetwork")
        acl_basic_k8s_cluster = scaleway.K8sCluster("aclBasicK8sCluster",
            version="1.32.2",
            cni="cilium",
            delete_additional_resources=True,
            private_network_id=acl_basic_vpc_private_network.id)
        acl_basic_k8s_acl = scaleway.K8sAcl("aclBasicK8sAcl",
            cluster_id=acl_basic_k8s_cluster.id,
            acl_rules=[
                {
                    "ip": "1.2.3.4/32",
                    "description": "Allow 1.2.3.4",
                },
                {
                    "scaleway_ranges": True,
                    "description": "Allow all Scaleway ranges",
                },
            ])
        ```

        ### Full-isolation

        ```python
        import pulumi
        import ediri_scaleway as scaleway

        acl_basic_vpc_private_network = scaleway.VpcPrivateNetwork("aclBasicVpcPrivateNetwork")
        acl_basic_k8s_cluster = scaleway.K8sCluster("aclBasicK8sCluster",
            version="1.32.2",
            cni="cilium",
            delete_additional_resources=True,
            private_network_id=acl_basic_vpc_private_network.id)
        acl_basic_k8s_acl = scaleway.K8sAcl("aclBasicK8sAcl",
            cluster_id=acl_basic_k8s_cluster.id,
            no_ip_allowed=True)
        ```

        ## Import

        Kubernetes ACLs can be imported using the `{region}/{cluster-id}`, e.g.

        bash

        ```sh
        $ pulumi import scaleway:index/k8sAcl:K8sAcl acl01 fr-par/11111111-1111-1111-1111-111111111111
        ```

        :param str resource_name: The name of the resource.
        :param K8sAclArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(K8sAclArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['K8sAclAclRuleArgs', 'K8sAclAclRuleArgsDict']]]]] = None,
                 cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                 no_ip_allowed: Optional[pulumi.Input[builtins.bool]] = None,
                 region: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = K8sAclArgs.__new__(K8sAclArgs)

            __props__.__dict__["acl_rules"] = acl_rules
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["no_ip_allowed"] = no_ip_allowed
            __props__.__dict__["region"] = region
        super(K8sAcl, __self__).__init__(
            'scaleway:index/k8sAcl:K8sAcl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl_rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['K8sAclAclRuleArgs', 'K8sAclAclRuleArgsDict']]]]] = None,
            cluster_id: Optional[pulumi.Input[builtins.str]] = None,
            no_ip_allowed: Optional[pulumi.Input[builtins.bool]] = None,
            region: Optional[pulumi.Input[builtins.str]] = None) -> 'K8sAcl':
        """
        Get an existing K8sAcl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['K8sAclAclRuleArgs', 'K8sAclAclRuleArgsDict']]]] acl_rules: A list of ACLs (structure is described below)
               
               > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        :param pulumi.Input[builtins.str] cluster_id: UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.
               
               > **Important:** Updates to `cluster_id` will recreate the ACL.
        :param pulumi.Input[builtins.bool] no_ip_allowed: If set to true, no IP will be allowed and the cluster will be in full-isolation.
               
               > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        :param pulumi.Input[builtins.str] region: `region`) The region in which the ACL rule should be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _K8sAclState.__new__(_K8sAclState)

        __props__.__dict__["acl_rules"] = acl_rules
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["no_ip_allowed"] = no_ip_allowed
        __props__.__dict__["region"] = region
        return K8sAcl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aclRules")
    def acl_rules(self) -> pulumi.Output[Optional[Sequence['outputs.K8sAclAclRule']]]:
        """
        A list of ACLs (structure is described below)

        > **Important:** This block cannot be defined if the `no_ip_allowed` field is set to true.
        """
        return pulumi.get(self, "acl_rules")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[builtins.str]:
        """
        UUID of the cluster. The ID of the cluster is also the ID of the ACL resource, as there can only be one per cluster.

        > **Important:** Updates to `cluster_id` will recreate the ACL.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="noIpAllowed")
    def no_ip_allowed(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If set to true, no IP will be allowed and the cluster will be in full-isolation.

        > **Important:** This field cannot be set to true if the `acl_rules` block is defined.
        """
        return pulumi.get(self, "no_ip_allowed")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[builtins.str]:
        """
        `region`) The region in which the ACL rule should be created.
        """
        return pulumi.get(self, "region")

