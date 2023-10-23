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
    [ScalewayResourceType("scaleway:index/documentDBReadReplica:DocumentDBReadReplica")]
    public partial class DocumentDBReadReplica : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Direct access endpoint, it gives you an IP and a port to access your read-replica
        /// </summary>
        [Output("directAccess")]
        public Output<Outputs.DocumentDBReadReplicaDirectAccess?> DirectAccess { get; private set; } = null!;

        /// <summary>
        /// Id of the rdb instance to replicate
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Private network endpoints
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.DocumentDBReadReplicaPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
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
        /// Direct access endpoint, it gives you an IP and a port to access your read-replica
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DocumentDBReadReplicaDirectAccessArgs>? DirectAccess { get; set; }

        /// <summary>
        /// Id of the rdb instance to replicate
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Private network endpoints
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentDBReadReplicaPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
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
        /// Direct access endpoint, it gives you an IP and a port to access your read-replica
        /// </summary>
        [Input("directAccess")]
        public Input<Inputs.DocumentDBReadReplicaDirectAccessGetArgs>? DirectAccess { get; set; }

        /// <summary>
        /// Id of the rdb instance to replicate
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Private network endpoints
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentDBReadReplicaPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentDBReadReplicaState()
        {
        }
        public static new DocumentDBReadReplicaState Empty => new DocumentDBReadReplicaState();
    }
}