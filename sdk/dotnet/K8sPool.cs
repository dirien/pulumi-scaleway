// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway
{
    /// <summary>
    /// ## Import
    /// 
    /// Kubernetes pools can be imported using the `{region}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/k8sPool:K8sPool mypool fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/k8sPool:K8sPool")]
    public partial class K8sPool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enables the autohealing feature for this pool.
        /// </summary>
        [Output("autohealing")]
        public Output<bool?> Autohealing { get; private set; } = null!;

        /// <summary>
        /// Enables the autoscaling feature for this pool.
        /// &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
        /// </summary>
        [Output("autoscaling")]
        public Output<bool?> Autoscaling { get; private set; } = null!;

        /// <summary>
        /// The ID of the Kubernetes cluster on which this pool will be created.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The container runtime of the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Output("containerRuntime")]
        public Output<string?> ContainerRuntime { get; private set; } = null!;

        /// <summary>
        /// The creation date of the pool.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The actual size of the pool
        /// </summary>
        [Output("currentSize")]
        public Output<int> CurrentSize { get; private set; } = null!;

        /// <summary>
        /// The Kubelet arguments to be used by this pool
        /// </summary>
        [Output("kubeletArgs")]
        public Output<ImmutableDictionary<string, string>?> KubeletArgs { get; private set; } = null!;

        /// <summary>
        /// The maximum size of the pool, used by the autoscaling feature.
        /// </summary>
        [Output("maxSize")]
        public Output<int> MaxSize { get; private set; } = null!;

        /// <summary>
        /// The minimum size of the pool, used by the autoscaling feature.
        /// </summary>
        [Output("minSize")]
        public Output<int?> MinSize { get; private set; } = null!;

        /// <summary>
        /// The name for the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
        /// 
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// (List of) The nodes in the default pool.
        /// </summary>
        [Output("nodes")]
        public Output<ImmutableArray<Outputs.K8sPoolNode>> Nodes { get; private set; } = null!;

        /// <summary>
        /// The [placement group](https://www.scaleway.com/en/developers/api/instance/#placement-groups-d8f653) the nodes of the pool will be attached to.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Output("placementGroupId")]
        public Output<string?> PlacementGroupId { get; private set; } = null!;

        /// <summary>
        /// Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Output("publicIpDisabled")]
        public Output<bool?> PublicIpDisabled { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the pool should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The size of the system volume of the nodes in gigabyte
        /// </summary>
        [Output("rootVolumeSizeInGb")]
        public Output<int> RootVolumeSizeInGb { get; private set; } = null!;

        /// <summary>
        /// System volume type of the nodes composing the pool
        /// </summary>
        [Output("rootVolumeType")]
        public Output<string> RootVolumeType { get; private set; } = null!;

        /// <summary>
        /// The size of the pool.
        /// &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// The status of the node.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the pool.
        /// &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool's nodes are applied using tags. (Example: "taint=taintName=taineValue:Effect")
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The last update date of the pool.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// The Pool upgrade policy
        /// </summary>
        [Output("upgradePolicy")]
        public Output<Outputs.K8sPoolUpgradePolicy> UpgradePolicy { get; private set; } = null!;

        /// <summary>
        /// The version of the pool.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Whether to wait for the pool to be ready.
        /// </summary>
        [Output("waitForPoolReady")]
        public Output<bool?> WaitForPoolReady { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the pool should be created.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a K8sPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public K8sPool(string name, K8sPoolArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/k8sPool:K8sPool", name, args ?? new K8sPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private K8sPool(string name, Input<string> id, K8sPoolState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/k8sPool:K8sPool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-scaleway",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing K8sPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static K8sPool Get(string name, Input<string> id, K8sPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new K8sPool(name, id, state, options);
        }
    }

    public sealed class K8sPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables the autohealing feature for this pool.
        /// </summary>
        [Input("autohealing")]
        public Input<bool>? Autohealing { get; set; }

        /// <summary>
        /// Enables the autoscaling feature for this pool.
        /// &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
        /// </summary>
        [Input("autoscaling")]
        public Input<bool>? Autoscaling { get; set; }

        /// <summary>
        /// The ID of the Kubernetes cluster on which this pool will be created.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The container runtime of the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("containerRuntime")]
        public Input<string>? ContainerRuntime { get; set; }

        [Input("kubeletArgs")]
        private InputMap<string>? _kubeletArgs;

        /// <summary>
        /// The Kubelet arguments to be used by this pool
        /// </summary>
        public InputMap<string> KubeletArgs
        {
            get => _kubeletArgs ?? (_kubeletArgs = new InputMap<string>());
            set => _kubeletArgs = value;
        }

        /// <summary>
        /// The maximum size of the pool, used by the autoscaling feature.
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// The minimum size of the pool, used by the autoscaling feature.
        /// </summary>
        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        /// <summary>
        /// The name for the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
        /// 
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        /// <summary>
        /// The [placement group](https://www.scaleway.com/en/developers/api/instance/#placement-groups-d8f653) the nodes of the pool will be attached to.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("placementGroupId")]
        public Input<string>? PlacementGroupId { get; set; }

        /// <summary>
        /// Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("publicIpDisabled")]
        public Input<bool>? PublicIpDisabled { get; set; }

        /// <summary>
        /// `region`) The region in which the pool should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The size of the system volume of the nodes in gigabyte
        /// </summary>
        [Input("rootVolumeSizeInGb")]
        public Input<int>? RootVolumeSizeInGb { get; set; }

        /// <summary>
        /// System volume type of the nodes composing the pool
        /// </summary>
        [Input("rootVolumeType")]
        public Input<string>? RootVolumeType { get; set; }

        /// <summary>
        /// The size of the pool.
        /// &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the pool.
        /// &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool's nodes are applied using tags. (Example: "taint=taintName=taineValue:Effect")
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The Pool upgrade policy
        /// </summary>
        [Input("upgradePolicy")]
        public Input<Inputs.K8sPoolUpgradePolicyArgs>? UpgradePolicy { get; set; }

        /// <summary>
        /// Whether to wait for the pool to be ready.
        /// </summary>
        [Input("waitForPoolReady")]
        public Input<bool>? WaitForPoolReady { get; set; }

        /// <summary>
        /// `zone`) The zone in which the pool should be created.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public K8sPoolArgs()
        {
        }
        public static new K8sPoolArgs Empty => new K8sPoolArgs();
    }

    public sealed class K8sPoolState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables the autohealing feature for this pool.
        /// </summary>
        [Input("autohealing")]
        public Input<bool>? Autohealing { get; set; }

        /// <summary>
        /// Enables the autoscaling feature for this pool.
        /// &gt; **Important:** When enabled, an update of the `size` will not be taken into account.
        /// </summary>
        [Input("autoscaling")]
        public Input<bool>? Autoscaling { get; set; }

        /// <summary>
        /// The ID of the Kubernetes cluster on which this pool will be created.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The container runtime of the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("containerRuntime")]
        public Input<string>? ContainerRuntime { get; set; }

        /// <summary>
        /// The creation date of the pool.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The actual size of the pool
        /// </summary>
        [Input("currentSize")]
        public Input<int>? CurrentSize { get; set; }

        [Input("kubeletArgs")]
        private InputMap<string>? _kubeletArgs;

        /// <summary>
        /// The Kubelet arguments to be used by this pool
        /// </summary>
        public InputMap<string> KubeletArgs
        {
            get => _kubeletArgs ?? (_kubeletArgs = new InputMap<string>());
            set => _kubeletArgs = value;
        }

        /// <summary>
        /// The maximum size of the pool, used by the autoscaling feature.
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// The minimum size of the pool, used by the autoscaling feature.
        /// </summary>
        [Input("minSize")]
        public Input<int>? MinSize { get; set; }

        /// <summary>
        /// The name for the pool.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The commercial type of the pool instances. Instances with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). `external` is a special node type used to provision from other Cloud providers.
        /// 
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("nodes")]
        private InputList<Inputs.K8sPoolNodeGetArgs>? _nodes;

        /// <summary>
        /// (List of) The nodes in the default pool.
        /// </summary>
        public InputList<Inputs.K8sPoolNodeGetArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.K8sPoolNodeGetArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// The [placement group](https://www.scaleway.com/en/developers/api/instance/#placement-groups-d8f653) the nodes of the pool will be attached to.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("placementGroupId")]
        public Input<string>? PlacementGroupId { get; set; }

        /// <summary>
        /// Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("publicIpDisabled")]
        public Input<bool>? PublicIpDisabled { get; set; }

        /// <summary>
        /// `region`) The region in which the pool should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The size of the system volume of the nodes in gigabyte
        /// </summary>
        [Input("rootVolumeSizeInGb")]
        public Input<int>? RootVolumeSizeInGb { get; set; }

        /// <summary>
        /// System volume type of the nodes composing the pool
        /// </summary>
        [Input("rootVolumeType")]
        public Input<string>? RootVolumeType { get; set; }

        /// <summary>
        /// The size of the pool.
        /// &gt; **Important:** This field will only be used at creation if autoscaling is enabled.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// The status of the node.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the pool.
        /// &gt; Note: As mentionned in [this document](https://github.com/scaleway/scaleway-cloud-controller-manager/blob/master/docs/tags.md#taints), taints of a pool's nodes are applied using tags. (Example: "taint=taintName=taineValue:Effect")
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The last update date of the pool.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// The Pool upgrade policy
        /// </summary>
        [Input("upgradePolicy")]
        public Input<Inputs.K8sPoolUpgradePolicyGetArgs>? UpgradePolicy { get; set; }

        /// <summary>
        /// The version of the pool.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Whether to wait for the pool to be ready.
        /// </summary>
        [Input("waitForPoolReady")]
        public Input<bool>? WaitForPoolReady { get; set; }

        /// <summary>
        /// `zone`) The zone in which the pool should be created.
        /// &gt; **Important:** Updates to this field will recreate a new resource.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public K8sPoolState()
        {
        }
        public static new K8sPoolState Empty => new K8sPoolState();
    }
}
