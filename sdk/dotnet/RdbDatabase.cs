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
    /// Creates and manages databases.
    /// For more information, refer to [the API documentation](https://www.scaleway.com/en/developers/api/managed-database-postgre-mysql/).
    /// 
    /// ## Example Usage
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
    ///     var mainRdbInstance = new Scaleway.RdbInstance("mainRdbInstance", new()
    ///     {
    ///         NodeType = "DB-DEV-S",
    ///         Engine = "PostgreSQL-15",
    ///         IsHaCluster = true,
    ///         DisableBackup = true,
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///     });
    /// 
    ///     var mainRdbDatabase = new Scaleway.RdbDatabase("mainRdbDatabase", new()
    ///     {
    ///         InstanceId = mainRdbInstance.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/rdbDatabase:RdbDatabase rdb01_mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/rdbDatabase:RdbDatabase")]
    public partial class RdbDatabase : global::Pulumi.CustomResource
    {
        /// <summary>
        /// UUID of the Database Instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the database.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Whether the database is managed or not.
        /// </summary>
        [Output("managed")]
        public Output<bool> Managed { get; private set; } = null!;

        /// <summary>
        /// Name of the database (e.g. `my-new-database`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the owner of the database.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Size of the database (in bytes).
        /// </summary>
        [Output("size")]
        public Output<string> Size { get; private set; } = null!;


        /// <summary>
        /// Create a RdbDatabase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdbDatabase(string name, RdbDatabaseArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/rdbDatabase:RdbDatabase", name, args ?? new RdbDatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdbDatabase(string name, Input<string> id, RdbDatabaseState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/rdbDatabase:RdbDatabase", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdbDatabase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdbDatabase Get(string name, Input<string> id, RdbDatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new RdbDatabase(name, id, state, options);
        }
    }

    public sealed class RdbDatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the Database Instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the database.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Name of the database (e.g. `my-new-database`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RdbDatabaseArgs()
        {
        }
        public static new RdbDatabaseArgs Empty => new RdbDatabaseArgs();
    }

    public sealed class RdbDatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the Database Instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the database.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Whether the database is managed or not.
        /// </summary>
        [Input("managed")]
        public Input<bool>? Managed { get; set; }

        /// <summary>
        /// Name of the database (e.g. `my-new-database`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the owner of the database.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Size of the database (in bytes).
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        public RdbDatabaseState()
        {
        }
        public static new RdbDatabaseState Empty => new RdbDatabaseState();
    }
}
