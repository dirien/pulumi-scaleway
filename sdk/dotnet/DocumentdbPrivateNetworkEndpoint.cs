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
    /// Creates and manages Scaleway Database Private Network Endpoint.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Database Instance Endpoint can be imported using the `{region}/{endpoint_id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint end fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint")]
    public partial class DocumentdbPrivateNetworkEndpoint : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Hostname of the endpoint.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// UUID of the documentdb instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// IPv4 address on the network.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a
        /// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        /// service if not set.
        /// </summary>
        [Output("ipNet")]
        public Output<string> IpNet { get; private set; } = null!;

        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Output("privateNetworkId")]
        public Output<string> PrivateNetworkId { get; private set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentdbPrivateNetworkEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentdbPrivateNetworkEndpoint(string name, DocumentdbPrivateNetworkEndpointArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint", name, args ?? new DocumentdbPrivateNetworkEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentdbPrivateNetworkEndpoint(string name, Input<string> id, DocumentdbPrivateNetworkEndpointState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbPrivateNetworkEndpoint:DocumentdbPrivateNetworkEndpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DocumentdbPrivateNetworkEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentdbPrivateNetworkEndpoint Get(string name, Input<string> id, DocumentdbPrivateNetworkEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new DocumentdbPrivateNetworkEndpoint(name, id, state, options);
        }
    }

    public sealed class DocumentdbPrivateNetworkEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the documentdb instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a
        /// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        /// service if not set.
        /// </summary>
        [Input("ipNet")]
        public Input<string>? IpNet { get; set; }

        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Input("privateNetworkId", required: true)]
        public Input<string> PrivateNetworkId { get; set; } = null!;

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DocumentdbPrivateNetworkEndpointArgs()
        {
        }
        public static new DocumentdbPrivateNetworkEndpointArgs Empty => new DocumentdbPrivateNetworkEndpointArgs();
    }

    public sealed class DocumentdbPrivateNetworkEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of the endpoint.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// UUID of the documentdb instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// IPv4 address on the network.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The IP network address within the private subnet. This must be an IPv4 address with a
        /// CIDR notation. The IP network address within the private subnet is determined by the IP Address Management (IPAM)
        /// service if not set.
        /// </summary>
        [Input("ipNet")]
        public Input<string>? IpNet { get; set; }

        /// <summary>
        /// Name of the endpoint.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Port in the Private Network.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The ID of the private network.
        /// </summary>
        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        /// <summary>
        /// The region you want to attach the resource to
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The zone you want to attach the resource to
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public DocumentdbPrivateNetworkEndpointState()
        {
        }
        public static new DocumentdbPrivateNetworkEndpointState Empty => new DocumentdbPrivateNetworkEndpointState();
    }
}
