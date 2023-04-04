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
    /// Creates and manages Scaleway RDB database backup.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.RdbDatabaseBackup("main", new()
    ///     {
    ///         InstanceId = data.Scaleway_rdb_instance.Main.Id,
    ///         DatabaseName = data.Scaleway_rdb_database.Main.Name,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### With expiration
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.RdbDatabaseBackup("main", new()
    ///     {
    ///         InstanceId = data.Scaleway_rdb_instance.Main.Id,
    ///         DatabaseName = data.Scaleway_rdb_database.Main.Name,
    ///         ExpiresAt = "2022-06-16T07:48:44Z",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDB Database can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup mybackup fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup")]
    public partial class RdbDatabaseBackup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation date (Format ISO 8601).
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the database of this backup.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Expiration date (Format ISO 8601).
        /// </summary>
        [Output("expiresAt")]
        public Output<string?> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Name of the instance of the backup.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// Name of the database (e.g. `my-database`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Size of the backup (in bytes).
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Updated date (Format ISO 8601).
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a RdbDatabaseBackup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdbDatabaseBackup(string name, RdbDatabaseBackupArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup", name, args ?? new RdbDatabaseBackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdbDatabaseBackup(string name, Input<string> id, RdbDatabaseBackupState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/rdbDatabaseBackup:RdbDatabaseBackup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdbDatabaseBackup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdbDatabaseBackup Get(string name, Input<string> id, RdbDatabaseBackupState? state = null, CustomResourceOptions? options = null)
        {
            return new RdbDatabaseBackup(name, id, state, options);
        }
    }

    public sealed class RdbDatabaseBackupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database of this backup.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// Expiration date (Format ISO 8601).
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Name of the database (e.g. `my-database`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdbDatabaseBackupArgs()
        {
        }
        public static new RdbDatabaseBackupArgs Empty => new RdbDatabaseBackupArgs();
    }

    public sealed class RdbDatabaseBackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation date (Format ISO 8601).
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the database of this backup.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Expiration date (Format ISO 8601).
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of the instance of the backup.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Name of the database (e.g. `my-database`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Size of the backup (in bytes).
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Updated date (Format ISO 8601).
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public RdbDatabaseBackupState()
        {
        }
        public static new RdbDatabaseBackupState Empty => new RdbDatabaseBackupState();
    }
}
