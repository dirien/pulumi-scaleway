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
    /// Creates and manages Scaleway MongoDB® instance.
    /// For more information refer to [the API documentation](https://www.scaleway.com/en/docs/managed-databases/mongodb/).
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
    ///     var main = new Scaleway.MongodbInstance("main", new()
    ///     {
    ///         NodeNumber = 1,
    ///         NodeType = "MGDB-PLAY2-NANO",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         UserName = "my_initial_user",
    ///         Version = "7.0.12",
    ///         VolumeSizeInGb = 5,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Restore From Snapshot
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var restoredInstance = new Scaleway.MongodbInstance("restoredInstance", new()
    ///     {
    ///         NodeNumber = 1,
    ///         NodeType = "MGDB-PLAY2-NANO",
    ///         SnapshotId = scaleway_vpc_private_network.Pn.Idscaleway_mongodb_snapshot.Main_snapshot.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MongoDB® instance can be imported using the `id`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/mongodbInstance:MongodbInstance main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mongodbInstance:MongodbInstance")]
    public partial class MongodbInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time of the creation of the MongoDB® instance.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the MongoDB® instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of nodes in the instance
        /// </summary>
        [Output("nodeNumber")]
        public Output<int> NodeNumber { get; private set; } = null!;

        /// <summary>
        /// The type of MongoDB® intance to create.
        /// </summary>
        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        /// <summary>
        /// Password of the user.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Public network specs details.
        /// </summary>
        [Output("publicNetwork")]
        public Output<Outputs.MongodbInstancePublicNetwork> PublicNetwork { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Map of settings to define for the instance.
        /// </summary>
        [Output("settings")]
        public Output<ImmutableDictionary<string, string>?> Settings { get; private set; } = null!;

        /// <summary>
        /// Snapshot ID to restore the MongoDB® instance from.
        /// </summary>
        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// List of tags attached to the MongoDB® instance.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The date and time of the last update of the MongoDB® instance.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the user created when the intance is created.
        /// </summary>
        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;

        /// <summary>
        /// MongoDB® version of the instance.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;

        /// <summary>
        /// Volume size in GB.
        /// </summary>
        [Output("volumeSizeInGb")]
        public Output<int> VolumeSizeInGb { get; private set; } = null!;

        /// <summary>
        /// Volume type of the instance.
        /// </summary>
        [Output("volumeType")]
        public Output<string?> VolumeType { get; private set; } = null!;


        /// <summary>
        /// Create a MongodbInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MongodbInstance(string name, MongodbInstanceArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/mongodbInstance:MongodbInstance", name, args ?? new MongodbInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MongodbInstance(string name, Input<string> id, MongodbInstanceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mongodbInstance:MongodbInstance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/dirien/pulumi-scaleway",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MongodbInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MongodbInstance Get(string name, Input<string> id, MongodbInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new MongodbInstance(name, id, state, options);
        }
    }

    public sealed class MongodbInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the MongoDB® instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of nodes in the instance
        /// </summary>
        [Input("nodeNumber", required: true)]
        public Input<int> NodeNumber { get; set; } = null!;

        /// <summary>
        /// The type of MongoDB® intance to create.
        /// </summary>
        [Input("nodeType", required: true)]
        public Input<string> NodeType { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password of the user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Public network specs details.
        /// </summary>
        [Input("publicNetwork")]
        public Input<Inputs.MongodbInstancePublicNetworkArgs>? PublicNetwork { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of settings to define for the instance.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        /// <summary>
        /// Snapshot ID to restore the MongoDB® instance from.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags attached to the MongoDB® instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Name of the user created when the intance is created.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// MongoDB® version of the instance.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Volume size in GB.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Volume type of the instance.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public MongodbInstanceArgs()
        {
        }
        public static new MongodbInstanceArgs Empty => new MongodbInstanceArgs();
    }

    public sealed class MongodbInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time of the creation of the MongoDB® instance.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the MongoDB® instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of nodes in the instance
        /// </summary>
        [Input("nodeNumber")]
        public Input<int>? NodeNumber { get; set; }

        /// <summary>
        /// The type of MongoDB® intance to create.
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password of the user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The project_id you want to attach the resource to
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Public network specs details.
        /// </summary>
        [Input("publicNetwork")]
        public Input<Inputs.MongodbInstancePublicNetworkGetArgs>? PublicNetwork { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("settings")]
        private InputMap<string>? _settings;

        /// <summary>
        /// Map of settings to define for the instance.
        /// </summary>
        public InputMap<string> Settings
        {
            get => _settings ?? (_settings = new InputMap<string>());
            set => _settings = value;
        }

        /// <summary>
        /// Snapshot ID to restore the MongoDB® instance from.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags attached to the MongoDB® instance.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The date and time of the last update of the MongoDB® instance.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// Name of the user created when the intance is created.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        /// <summary>
        /// MongoDB® version of the instance.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// Volume size in GB.
        /// </summary>
        [Input("volumeSizeInGb")]
        public Input<int>? VolumeSizeInGb { get; set; }

        /// <summary>
        /// Volume type of the instance.
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public MongodbInstanceState()
        {
        }
        public static new MongodbInstanceState Empty => new MongodbInstanceState();
    }
}
