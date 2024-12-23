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
    public static class GetRdbInstance
    {
        /// <summary>
        /// Gets information about an Database Instance.
        /// 
        /// For further information refer the Managed Databases for PostgreSQL and MySQL [API documentation](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
        /// </summary>
        public static Task<GetRdbInstanceResult> InvokeAsync(GetRdbInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRdbInstanceResult>("scaleway:index/getRdbInstance:getRdbInstance", args ?? new GetRdbInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an Database Instance.
        /// 
        /// For further information refer the Managed Databases for PostgreSQL and MySQL [API documentation](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
        /// </summary>
        public static Output<GetRdbInstanceResult> Invoke(GetRdbInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdbInstanceResult>("scaleway:index/getRdbInstance:getRdbInstance", args ?? new GetRdbInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an Database Instance.
        /// 
        /// For further information refer the Managed Databases for PostgreSQL and MySQL [API documentation](https://developers.scaleway.com/en/products/rdb/api/#database-instance)
        /// </summary>
        public static Output<GetRdbInstanceResult> Invoke(GetRdbInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdbInstanceResult>("scaleway:index/getRdbInstance:getRdbInstance", args ?? new GetRdbInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdbInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `instance_id`.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The name of the RDB instance.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The ID of the project the Database Instance is in. Can be used to filter instances when using `name`.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance exists.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetRdbInstanceArgs()
        {
        }
        public static new GetRdbInstanceArgs Empty => new GetRdbInstanceArgs();
    }

    public sealed class GetRdbInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB instance ID.
        /// 
        /// &gt; **Note** You must specify at least one: `name` and/or `instance_id`.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the RDB instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the Database Instance is in. Can be used to filter instances when using `name`.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetRdbInstanceInvokeArgs()
        {
        }
        public static new GetRdbInstanceInvokeArgs Empty => new GetRdbInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdbInstanceResult
    {
        public readonly bool BackupSameRegion;
        public readonly int BackupScheduleFrequency;
        public readonly int BackupScheduleRetention;
        public readonly string Certificate;
        public readonly bool DisableBackup;
        public readonly bool EncryptionAtRest;
        public readonly string EndpointIp;
        public readonly int EndpointPort;
        public readonly string Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> InitSettings;
        public readonly string? InstanceId;
        public readonly bool IsHaCluster;
        public readonly ImmutableArray<Outputs.GetRdbInstanceLoadBalancerResult> LoadBalancers;
        public readonly ImmutableArray<Outputs.GetRdbInstanceLogsPolicyResult> LogsPolicies;
        public readonly string? Name;
        public readonly string NodeType;
        public readonly string OrganizationId;
        public readonly string Password;
        public readonly ImmutableArray<Outputs.GetRdbInstancePrivateNetworkResult> PrivateNetworks;
        public readonly string? ProjectId;
        public readonly ImmutableArray<Outputs.GetRdbInstanceReadReplicaResult> ReadReplicas;
        public readonly string? Region;
        public readonly ImmutableDictionary<string, string> Settings;
        public readonly ImmutableArray<string> Tags;
        public readonly string UserName;
        public readonly int VolumeSizeInGb;
        public readonly string VolumeType;

        [OutputConstructor]
        private GetRdbInstanceResult(
            bool backupSameRegion,

            int backupScheduleFrequency,

            int backupScheduleRetention,

            string certificate,

            bool disableBackup,

            bool encryptionAtRest,

            string endpointIp,

            int endpointPort,

            string engine,

            string id,

            ImmutableDictionary<string, string> initSettings,

            string? instanceId,

            bool isHaCluster,

            ImmutableArray<Outputs.GetRdbInstanceLoadBalancerResult> loadBalancers,

            ImmutableArray<Outputs.GetRdbInstanceLogsPolicyResult> logsPolicies,

            string? name,

            string nodeType,

            string organizationId,

            string password,

            ImmutableArray<Outputs.GetRdbInstancePrivateNetworkResult> privateNetworks,

            string? projectId,

            ImmutableArray<Outputs.GetRdbInstanceReadReplicaResult> readReplicas,

            string? region,

            ImmutableDictionary<string, string> settings,

            ImmutableArray<string> tags,

            string userName,

            int volumeSizeInGb,

            string volumeType)
        {
            BackupSameRegion = backupSameRegion;
            BackupScheduleFrequency = backupScheduleFrequency;
            BackupScheduleRetention = backupScheduleRetention;
            Certificate = certificate;
            DisableBackup = disableBackup;
            EncryptionAtRest = encryptionAtRest;
            EndpointIp = endpointIp;
            EndpointPort = endpointPort;
            Engine = engine;
            Id = id;
            InitSettings = initSettings;
            InstanceId = instanceId;
            IsHaCluster = isHaCluster;
            LoadBalancers = loadBalancers;
            LogsPolicies = logsPolicies;
            Name = name;
            NodeType = nodeType;
            OrganizationId = organizationId;
            Password = password;
            PrivateNetworks = privateNetworks;
            ProjectId = projectId;
            ReadReplicas = readReplicas;
            Region = region;
            Settings = settings;
            Tags = tags;
            UserName = userName;
            VolumeSizeInGb = volumeSizeInGb;
            VolumeType = volumeType;
        }
    }
}
