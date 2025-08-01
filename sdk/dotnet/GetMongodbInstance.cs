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
    public static class GetMongodbInstance
    {
        /// <summary>
        /// Gets information about a MongoDB® Instance.
        /// 
        /// For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)
        /// </summary>
        public static Task<GetMongodbInstanceResult> InvokeAsync(GetMongodbInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMongodbInstanceResult>("scaleway:index/getMongodbInstance:getMongodbInstance", args ?? new GetMongodbInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a MongoDB® Instance.
        /// 
        /// For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)
        /// </summary>
        public static Output<GetMongodbInstanceResult> Invoke(GetMongodbInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongodbInstanceResult>("scaleway:index/getMongodbInstance:getMongodbInstance", args ?? new GetMongodbInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about a MongoDB® Instance.
        /// 
        /// For further information refer to the Managed Databases for MongoDB® [API documentation](https://developers.scaleway.com/en/products/mongodb/api/)
        /// </summary>
        public static Output<GetMongodbInstanceResult> Invoke(GetMongodbInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMongodbInstanceResult>("scaleway:index/getMongodbInstance:getMongodbInstance", args ?? new GetMongodbInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMongodbInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The MongoDB® instance ID.
        /// 
        /// &gt; **Note** You must specify at least one: `name` or `instance_id`.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The name of the MongoDB® instance.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the MongoDB® Instance exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetMongodbInstanceArgs()
        {
        }
        public static new GetMongodbInstanceArgs Empty => new GetMongodbInstanceArgs();
    }

    public sealed class GetMongodbInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The MongoDB® instance ID.
        /// 
        /// &gt; **Note** You must specify at least one: `name` or `instance_id`.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the MongoDB® instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the MongoDB® instance is in. Can be used to filter instances when using `name`.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the MongoDB® Instance exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetMongodbInstanceInvokeArgs()
        {
        }
        public static new GetMongodbInstanceInvokeArgs Empty => new GetMongodbInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetMongodbInstanceResult
    {
        /// <summary>
        /// The date and time the MongoDB® instance was created.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        /// <summary>
        /// The name of the MongoDB® instance.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The number of nodes in the MongoDB® cluster.
        /// </summary>
        public readonly int NodeNumber;
        /// <summary>
        /// The type of MongoDB® node.
        /// </summary>
        public readonly string NodeType;
        public readonly string Password;
        public readonly ImmutableArray<Outputs.GetMongodbInstancePrivateIpResult> PrivateIps;
        public readonly ImmutableArray<Outputs.GetMongodbInstancePrivateNetworkResult> PrivateNetworks;
        /// <summary>
        /// The ID of the project the instance belongs to.
        /// </summary>
        public readonly string? ProjectId;
        /// <summary>
        /// The details of the public network configuration, if applicable.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetMongodbInstancePublicNetworkResult> PublicNetworks;
        public readonly string? Region;
        public readonly ImmutableDictionary<string, string> Settings;
        public readonly string SnapshotId;
        /// <summary>
        /// A list of tags attached to the MongoDB® instance.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        public readonly string UpdatedAt;
        public readonly string UserName;
        /// <summary>
        /// The version of MongoDB® running on the instance.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// The size of the attached volume, in GB.
        /// </summary>
        public readonly int VolumeSizeInGb;
        /// <summary>
        /// The type of volume attached to the MongoDB® instance.
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private GetMongodbInstanceResult(
            string createdAt,

            string id,

            string? instanceId,

            string? name,

            int nodeNumber,

            string nodeType,

            string password,

            ImmutableArray<Outputs.GetMongodbInstancePrivateIpResult> privateIps,

            ImmutableArray<Outputs.GetMongodbInstancePrivateNetworkResult> privateNetworks,

            string? projectId,

            ImmutableArray<Outputs.GetMongodbInstancePublicNetworkResult> publicNetworks,

            string? region,

            ImmutableDictionary<string, string> settings,

            string snapshotId,

            ImmutableArray<string> tags,

            string updatedAt,

            string userName,

            string version,

            int volumeSizeInGb,

            string volumeType)
        {
            CreatedAt = createdAt;
            Id = id;
            InstanceId = instanceId;
            Name = name;
            NodeNumber = nodeNumber;
            NodeType = nodeType;
            Password = password;
            PrivateIps = privateIps;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            PublicNetworks = publicNetworks;
            Region = region;
            Settings = settings;
            SnapshotId = snapshotId;
            Tags = tags;
            UpdatedAt = updatedAt;
            UserName = userName;
            Version = version;
            VolumeSizeInGb = volumeSizeInGb;
            VolumeType = volumeType;
        }
    }
}
