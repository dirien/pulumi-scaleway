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
    /// ## Import
    /// 
    /// IoT Networks can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/iotNetwork:IotNetwork net01 fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/iotNetwork:IotNetwork")]
    public partial class IotNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time the Network was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The endpoint to use when interacting with the network.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The hub ID to which the Network will be attached to.
        /// </summary>
        [Output("hubId")]
        public Output<string> HubId { get; private set; } = null!;

        /// <summary>
        /// The name of the IoT Network you want to create (e.g. `my-net`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The endpoint key to keep secret.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// The prefix that will be prepended to all topics for this Network.
        /// </summary>
        [Output("topicPrefix")]
        public Output<string?> TopicPrefix { get; private set; } = null!;

        /// <summary>
        /// The network type to create (e.g. `sigfox`).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IotNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IotNetwork(string name, IotNetworkArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/iotNetwork:IotNetwork", name, args ?? new IotNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IotNetwork(string name, Input<string> id, IotNetworkState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/iotNetwork:IotNetwork", name, state, MakeResourceOptions(options, id))
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
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IotNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IotNetwork Get(string name, Input<string> id, IotNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new IotNetwork(name, id, state, options);
        }
    }

    public sealed class IotNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The hub ID to which the Network will be attached to.
        /// </summary>
        [Input("hubId", required: true)]
        public Input<string> HubId { get; set; } = null!;

        /// <summary>
        /// The name of the IoT Network you want to create (e.g. `my-net`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The prefix that will be prepended to all topics for this Network.
        /// </summary>
        [Input("topicPrefix")]
        public Input<string>? TopicPrefix { get; set; }

        /// <summary>
        /// The network type to create (e.g. `sigfox`).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IotNetworkArgs()
        {
        }
        public static new IotNetworkArgs Empty => new IotNetworkArgs();
    }

    public sealed class IotNetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time the Network was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The endpoint to use when interacting with the network.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The hub ID to which the Network will be attached to.
        /// </summary>
        [Input("hubId")]
        public Input<string>? HubId { get; set; }

        /// <summary>
        /// The name of the IoT Network you want to create (e.g. `my-net`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("secret")]
        private Input<string>? _secret;

        /// <summary>
        /// The endpoint key to keep secret.
        /// </summary>
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The prefix that will be prepended to all topics for this Network.
        /// </summary>
        [Input("topicPrefix")]
        public Input<string>? TopicPrefix { get; set; }

        /// <summary>
        /// The network type to create (e.g. `sigfox`).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IotNetworkState()
        {
        }
        public static new IotNetworkState Empty => new IotNetworkState();
    }
}
