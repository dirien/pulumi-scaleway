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
    /// Creates and manages Scaleway compute Instance User Data values.
    /// 
    /// User data is a key value store API you can use to provide data from and to your server without authentication. It is the mechanism by which a user can pass information contained in a local file to an Instance at launch time.
    /// 
    /// The typical use case is to pass something like a shell script or a configuration file as user data.
    /// 
    /// For more information about [user_data](https://www.scaleway.com/en/developers/api/instance/#path-user-data-list-user-data)  check our documentation guide [here](https://www.scaleway.com/en/docs/compute/instances/how-to/use-boot-modes/#how-to-use-cloud-init).
    /// 
    /// About cloud-init documentation please check this [link](https://cloudinit.readthedocs.io/en/latest/).
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
    ///     var config = new Config();
    ///     var userData = config.GetObject&lt;dynamic&gt;("userData") ?? 
    ///     {
    ///         { "cloud-init", @"#cloud-config
    /// apt-update: true
    /// apt-upgrade: true
    /// " },
    ///         { "foo", "bar" },
    ///     };
    ///     var mainInstanceServer = new Scaleway.InstanceServer("mainInstanceServer", new()
    ///     {
    ///         Image = "ubuntu_focal",
    ///         Type = "DEV1-S",
    ///     });
    /// 
    ///     // User data with a single value
    ///     var mainInstanceUserData = new Scaleway.InstanceUserData("mainInstanceUserData", new()
    ///     {
    ///         ServerId = mainInstanceServer.Id,
    ///         Key = "foo",
    ///         Value = "bar",
    ///     });
    /// 
    ///     // User Data with many keys.
    ///     var data = new List&lt;Scaleway.InstanceUserData&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; userData; rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         data.Add(new Scaleway.InstanceUserData($"data-{range.Value}", new()
    ///         {
    ///             ServerId = mainInstanceServer.Id,
    ///             Key = range.Key,
    ///             Value = range.Value,
    ///         }));
    ///     }
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// User data can be imported using the `{zone}/{key}/{server_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/instanceUserData:InstanceUserData main fr-par-1/cloud-init/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/instanceUserData:InstanceUserData")]
    public partial class InstanceUserData : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Key of the user data.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The ID of the server associated with.
        /// </summary>
        [Output("serverId")]
        public Output<string> ServerId { get; private set; } = null!;

        /// <summary>
        /// Value associated with your key
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// 
        /// &gt; **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        /// You can define values using:
        /// - string
        /// - UTF-8 encoded file content using file
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceUserData resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceUserData(string name, InstanceUserDataArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceUserData:InstanceUserData", name, args ?? new InstanceUserDataArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceUserData(string name, Input<string> id, InstanceUserDataState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceUserData:InstanceUserData", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceUserData resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceUserData Get(string name, Input<string> id, InstanceUserDataState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceUserData(name, id, state, options);
        }
    }

    public sealed class InstanceUserDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the user data.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The ID of the server associated with.
        /// </summary>
        [Input("serverId", required: true)]
        public Input<string> ServerId { get; set; } = null!;

        /// <summary>
        /// Value associated with your key
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// 
        /// &gt; **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        /// You can define values using:
        /// - string
        /// - UTF-8 encoded file content using file
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceUserDataArgs()
        {
        }
        public static new InstanceUserDataArgs Empty => new InstanceUserDataArgs();
    }

    public sealed class InstanceUserDataState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the user data.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The ID of the server associated with.
        /// </summary>
        [Input("serverId")]
        public Input<string>? ServerId { get; set; }

        /// <summary>
        /// Value associated with your key
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// 
        /// &gt; **Important:**   Use the `cloud-init` key to use [cloud-init](https://cloudinit.readthedocs.io/en/latest/) on your instance.
        /// You can define values using:
        /// - string
        /// - UTF-8 encoded file content using file
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceUserDataState()
        {
        }
        public static new InstanceUserDataState Empty => new InstanceUserDataState();
    }
}
