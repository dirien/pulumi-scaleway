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
    public static class GetRdbDatabaseBackup
    {
        /// <summary>
        /// Gets information about an RDB backup.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var findByName = Scaleway.GetRdbDatabaseBackup.Invoke(new()
        ///     {
        ///         Name = "mybackup",
        ///     });
        /// 
        ///     var findByNameAndInstance = Scaleway.GetRdbDatabaseBackup.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///         Name = "mybackup",
        ///     });
        /// 
        ///     var findById = Scaleway.GetRdbDatabaseBackup.Invoke(new()
        ///     {
        ///         BackupId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRdbDatabaseBackupResult> InvokeAsync(GetRdbDatabaseBackupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRdbDatabaseBackupResult>("scaleway:index/getRdbDatabaseBackup:getRdbDatabaseBackup", args ?? new GetRdbDatabaseBackupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an RDB backup.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var findByName = Scaleway.GetRdbDatabaseBackup.Invoke(new()
        ///     {
        ///         Name = "mybackup",
        ///     });
        /// 
        ///     var findByNameAndInstance = Scaleway.GetRdbDatabaseBackup.Invoke(new()
        ///     {
        ///         InstanceId = "11111111-1111-1111-1111-111111111111",
        ///         Name = "mybackup",
        ///     });
        /// 
        ///     var findById = Scaleway.GetRdbDatabaseBackup.Invoke(new()
        ///     {
        ///         BackupId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRdbDatabaseBackupResult> Invoke(GetRdbDatabaseBackupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdbDatabaseBackupResult>("scaleway:index/getRdbDatabaseBackup:getRdbDatabaseBackup", args ?? new GetRdbDatabaseBackupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdbDatabaseBackupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB backup ID.
        /// Only one of the `name` and `backup_id` should be specified.
        /// </summary>
        [Input("backupId")]
        public string? BackupId { get; set; }

        /// <summary>
        /// The RDB instance ID.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The name of the RDB instance.
        /// Only one of the `name` and `backup_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetRdbDatabaseBackupArgs()
        {
        }
        public static new GetRdbDatabaseBackupArgs Empty => new GetRdbDatabaseBackupArgs();
    }

    public sealed class GetRdbDatabaseBackupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The RDB backup ID.
        /// Only one of the `name` and `backup_id` should be specified.
        /// </summary>
        [Input("backupId")]
        public Input<string>? BackupId { get; set; }

        /// <summary>
        /// The RDB instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The name of the RDB instance.
        /// Only one of the `name` and `backup_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the Database Instance should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetRdbDatabaseBackupInvokeArgs()
        {
        }
        public static new GetRdbDatabaseBackupInvokeArgs Empty => new GetRdbDatabaseBackupInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdbDatabaseBackupResult
    {
        public readonly string? BackupId;
        public readonly string CreatedAt;
        public readonly string DatabaseName;
        public readonly string ExpiresAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? InstanceId;
        public readonly string InstanceName;
        public readonly string? Name;
        public readonly string? Region;
        public readonly int Size;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetRdbDatabaseBackupResult(
            string? backupId,

            string createdAt,

            string databaseName,

            string expiresAt,

            string id,

            string? instanceId,

            string instanceName,

            string? name,

            string? region,

            int size,

            string updatedAt)
        {
            BackupId = backupId;
            CreatedAt = createdAt;
            DatabaseName = databaseName;
            ExpiresAt = expiresAt;
            Id = id;
            InstanceId = instanceId;
            InstanceName = instanceName;
            Name = name;
            Region = region;
            Size = size;
            UpdatedAt = updatedAt;
        }
    }
}
