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
    /// Creates and manages Scaleway DocumentDB Database read replicas.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/document_db/).
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
    ///     var replica = new Scaleway.DocumentDBReadReplica("replica", new()
    ///     {
    ///         DirectAccess = null,
    ///         InstanceId = "11111111-1111-1111-1111-111111111111",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Private network
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pn = new Scaleway.VpcPrivateNetwork("pn");
    /// 
    ///     var replica = new Scaleway.DocumentDBReadReplica("replica", new()
    ///     {
    ///         InstanceId = scaleway_rdb_instance.Instance.Id,
    ///         PrivateNetwork = new Scaleway.Inputs.DocumentDBReadReplicaPrivateNetworkArgs
    ///         {
    ///             PrivateNetworkId = pn.Id,
    ///             ServiceIp = "192.168.1.254/24",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database Read replica can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/documentDBReadReplica:DocumentDBReadReplica rr fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/documentDBReadReplica:DocumentDBReadReplica")]
    public partial class DocumentDBReadReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creates a direct access endpoint to documentdb replica.
        /// </summary>
        [Output("directAccess")]
        public Output<Outputs.DocumentDBReadReplicaDirectAccess?> DirectAccess { get; private set; } = null!;

        /// <summary>
        /// UUID of the documentdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.DocumentDBReadReplicaPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentDBReadReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentDBReadReplica(string name, DocumentDBReadReplicaArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/documentDBReadReplica:DocumentDBReadReplica", name, args ?? new DocumentDBReadReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentDBReadReplica(string name, Input<string> id, DocumentDBReadReplicaState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/documentDBReadReplica:DocumentDBReadReplica", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DocumentDBReadReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentDBReadReplica Get(string name, Input<string> id, DocumentDBReadReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new DocumentDBReadReplica(name, id, state, options);
        }
    }

    public sealed class DocumentDBReadReplicaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to documentdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DocumentDBReadReplicaDirectAccessArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the documentdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentDBReadReplicaPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentDBReadReplicaArgs()
        {
        }
        public static new DocumentDBReadReplicaArgs Empty => new DocumentDBReadReplicaArgs();
    }

    public sealed class DocumentDBReadReplicaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creates a direct access endpoint to documentdb replica.
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DocumentDBReadReplicaDirectAccessGetArgs>? DirectAccess { get; set; }

        /// <summary>
        /// UUID of the documentdb instance.
        /// 
        /// &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Create an endpoint in a private network.
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentDBReadReplicaPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// `region`) The region
        /// in which the Database read replica should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentDBReadReplicaState()
        {
        }
        public static new DocumentDBReadReplicaState Empty => new DocumentDBReadReplicaState();
    }
}
