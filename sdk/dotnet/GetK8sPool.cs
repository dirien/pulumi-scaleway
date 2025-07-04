// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway
{
    public static class GetK8sPool
    {
        /// <summary>
        /// Gets information about a Kubernetes Cluster's Pool.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetK8sPool.Invoke(new()
        ///     {
        ///         PoolId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetK8sPoolResult> InvokeAsync(GetK8sPoolArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetK8sPoolResult>("scaleway:index/getK8sPool:getK8sPool", args ?? new GetK8sPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Kubernetes Cluster's Pool.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetK8sPool.Invoke(new()
        ///     {
        ///         PoolId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetK8sPoolResult> Invoke(GetK8sPoolInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetK8sPoolResult>("scaleway:index/getK8sPool:getK8sPool", args ?? new GetK8sPoolInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a Kubernetes Cluster's Pool.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKey = Scaleway.GetK8sPool.Invoke(new()
        ///     {
        ///         PoolId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetK8sPoolResult> Invoke(GetK8sPoolInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetK8sPoolResult>("scaleway:index/getK8sPool:getK8sPool", args ?? new GetK8sPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetK8sPoolArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster ID. Required when `name` is set.
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// The pool name. Only one of `name` and `pool_id` should be specified. `cluster_id` should be specified with `name`.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The pool's ID. Only one of `name` and `pool_id` should be specified.
        /// </summary>
        [Input("poolId")]
        public string? PoolId { get; set; }

        /// <summary>
        /// `region`) The region in which the pool exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The size of the pool.
        /// </summary>
        [Input("size")]
        public int? Size { get; set; }

        public GetK8sPoolArgs()
        {
        }
        public static new GetK8sPoolArgs Empty => new GetK8sPoolArgs();
    }

    public sealed class GetK8sPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The cluster ID. Required when `name` is set.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The pool name. Only one of `name` and `pool_id` should be specified. `cluster_id` should be specified with `name`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The pool's ID. Only one of `name` and `pool_id` should be specified.
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// `region`) The region in which the pool exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The size of the pool.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        public GetK8sPoolInvokeArgs()
        {
        }
        public static new GetK8sPoolInvokeArgs Empty => new GetK8sPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetK8sPoolResult
    {
        /// <summary>
        /// True if the autohealing feature is enabled for this pool.
        /// </summary>
        public readonly bool Autohealing;
        /// <summary>
        /// True if the autoscaling feature is enabled for this pool.
        /// </summary>
        public readonly bool Autoscaling;
        public readonly string? ClusterId;
        /// <summary>
        /// The container runtime of the pool.
        /// </summary>
        public readonly string ContainerRuntime;
        /// <summary>
        /// The creation date of the pool.
        /// </summary>
        public readonly string CreatedAt;
        public readonly int CurrentSize;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> KubeletArgs;
        /// <summary>
        /// The maximum size of the pool, used by the autoscaling feature.
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// The minimum size of the pool, used by the autoscaling feature.
        /// </summary>
        public readonly int MinSize;
        /// <summary>
        /// The name of the node.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The commercial type of the pool instances.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// (List of) The nodes in the default pool.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetK8sPoolNodeResult> Nodes;
        /// <summary>
        /// [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the nodes of the pool are attached to.
        /// </summary>
        public readonly string PlacementGroupId;
        public readonly string? PoolId;
        public readonly bool PublicIpDisabled;
        public readonly string? Region;
        public readonly int RootVolumeSizeInGb;
        public readonly string RootVolumeType;
        /// <summary>
        /// The size of the pool.
        /// </summary>
        public readonly int? Size;
        /// <summary>
        /// The status of the node.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tags associated with the pool.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The last update date of the pool.
        /// </summary>
        public readonly string UpdatedAt;
        public readonly ImmutableArray<Outputs.GetK8sPoolUpgradePolicyResult> UpgradePolicies;
        /// <summary>
        /// The version of the pool.
        /// </summary>
        public readonly string Version;
        public readonly bool WaitForPoolReady;
        public readonly string Zone;

        [OutputConstructor]
        private GetK8sPoolResult(
            bool autohealing,

            bool autoscaling,

            string? clusterId,

            string containerRuntime,

            string createdAt,

            int currentSize,

            string id,

            ImmutableDictionary<string, string> kubeletArgs,

            int maxSize,

            int minSize,

            string? name,

            string nodeType,

            ImmutableArray<Outputs.GetK8sPoolNodeResult> nodes,

            string placementGroupId,

            string? poolId,

            bool publicIpDisabled,

            string? region,

            int rootVolumeSizeInGb,

            string rootVolumeType,

            int? size,

            string status,

            ImmutableArray<string> tags,

            string updatedAt,

            ImmutableArray<Outputs.GetK8sPoolUpgradePolicyResult> upgradePolicies,

            string version,

            bool waitForPoolReady,

            string zone)
        {
            Autohealing = autohealing;
            Autoscaling = autoscaling;
            ClusterId = clusterId;
            ContainerRuntime = containerRuntime;
            CreatedAt = createdAt;
            CurrentSize = currentSize;
            Id = id;
            KubeletArgs = kubeletArgs;
            MaxSize = maxSize;
            MinSize = minSize;
            Name = name;
            NodeType = nodeType;
            Nodes = nodes;
            PlacementGroupId = placementGroupId;
            PoolId = poolId;
            PublicIpDisabled = publicIpDisabled;
            Region = region;
            RootVolumeSizeInGb = rootVolumeSizeInGb;
            RootVolumeType = rootVolumeType;
            Size = size;
            Status = status;
            Tags = tags;
            UpdatedAt = updatedAt;
            UpgradePolicies = upgradePolicies;
            Version = version;
            WaitForPoolReady = waitForPoolReady;
            Zone = zone;
        }
    }
}
