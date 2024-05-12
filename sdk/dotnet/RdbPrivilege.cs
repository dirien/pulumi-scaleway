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
    /// Create and manage Scaleway RDB database privilege.
    /// For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api/#user-and-permissions).
    /// 
    /// ## Example Usage
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
    ///         Engine = "PostgreSQL-11",
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
    ///     var mainRdbUser = new Scaleway.RdbUser("mainRdbUser", new()
    ///     {
    ///         InstanceId = mainRdbInstance.Id,
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         IsAdmin = false,
    ///     });
    /// 
    ///     var mainRdbPrivilege = new Scaleway.RdbPrivilege("mainRdbPrivilege", new()
    ///     {
    ///         InstanceId = mainRdbInstance.Id,
    ///         UserName = "my-db-user",
    ///         DatabaseName = "my-db-name",
    ///         Permission = "all",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             mainRdbUser,
    ///             mainRdbDatabase,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// The user privileges can be imported using the `{region}/{instance_id}/{database_name}/{user_name}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/rdbPrivilege:RdbPrivilege o fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/rdbPrivilege:RdbPrivilege")]
    public partial class RdbPrivilege : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the database (e.g. `my-db-name`).
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        /// </summary>
        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Name of the user (e.g. `my-db-user`).
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a RdbPrivilege resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdbPrivilege(string name, RdbPrivilegeArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/rdbPrivilege:RdbPrivilege", name, args ?? new RdbPrivilegeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdbPrivilege(string name, Input<string> id, RdbPrivilegeState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/rdbPrivilege:RdbPrivilege", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdbPrivilege resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdbPrivilege Get(string name, Input<string> id, RdbPrivilegeState? state = null, CustomResourceOptions? options = null)
        {
            return new RdbPrivilege(name, id, state, options);
        }
    }

    public sealed class RdbPrivilegeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database (e.g. `my-db-name`).
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Name of the user (e.g. `my-db-user`).
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public RdbPrivilegeArgs()
        {
        }
        public static new RdbPrivilegeArgs Empty => new RdbPrivilegeArgs();
    }

    public sealed class RdbPrivilegeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the database (e.g. `my-db-name`).
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// UUID of the rdb instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Permission to set. Valid values are `readonly`, `readwrite`, `all`, `custom` and `none`.
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// `region`) The region in which the resource exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Name of the user (e.g. `my-db-user`).
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public RdbPrivilegeState()
        {
        }
        public static new RdbPrivilegeState Empty => new RdbPrivilegeState();
    }
}
