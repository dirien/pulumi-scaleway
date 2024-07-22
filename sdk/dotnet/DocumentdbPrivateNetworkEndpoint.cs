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
    /// ### Example Basic
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
    ///     var instance = new Scaleway.DocumentdbInstance("instance", new()
    ///     {
    ///         NodeType = "docdb-play2-pico",
    ///         Engine = "FerretDB-1",
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         VolumeSizeInGb = 20,
    ///     });
    /// 
    ///     var main = new Scaleway.DocumentdbPrivateNetworkEndpoint("main", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         PrivateNetwork = new Scaleway.Inputs.DocumentdbPrivateNetworkEndpointPrivateNetworkArgs
    ///         {
    ///             IpNet = "172.16.32.3/22",
    ///             Id = pn.Id,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn =
    ///         {
    ///             pn,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
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
        /// UUID of the documentdb instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The private network specs details. This is a list with maximum one element and supports the following attributes:
        /// </summary>
        [Output("privateNetwork")]
        public Output<Outputs.DocumentdbPrivateNetworkEndpointPrivateNetwork?> PrivateNetwork { get; private set; } = null!;

        /// <summary>
        /// The region of the endpoint.
        /// 
        /// 
        /// &gt; **NOTE:** Please calculate your host IP.
        /// using cirhost. Otherwise, lets IPAM service
        /// handle the host IP on the network.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


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
        /// The private network specs details. This is a list with maximum one element and supports the following attributes:
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentdbPrivateNetworkEndpointPrivateNetworkArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// The region of the endpoint.
        /// 
        /// 
        /// &gt; **NOTE:** Please calculate your host IP.
        /// using cirhost. Otherwise, lets IPAM service
        /// handle the host IP on the network.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentdbPrivateNetworkEndpointArgs()
        {
        }
        public static new DocumentdbPrivateNetworkEndpointArgs Empty => new DocumentdbPrivateNetworkEndpointArgs();
    }

    public sealed class DocumentdbPrivateNetworkEndpointState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the documentdb instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The private network specs details. This is a list with maximum one element and supports the following attributes:
        /// </summary>
        [Input("privateNetwork")]
        public Input<Inputs.DocumentdbPrivateNetworkEndpointPrivateNetworkGetArgs>? PrivateNetwork { get; set; }

        /// <summary>
        /// The region of the endpoint.
        /// 
        /// 
        /// &gt; **NOTE:** Please calculate your host IP.
        /// using cirhost. Otherwise, lets IPAM service
        /// handle the host IP on the network.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentdbPrivateNetworkEndpointState()
        {
        }
        public static new DocumentdbPrivateNetworkEndpointState Empty => new DocumentdbPrivateNetworkEndpointState();
    }
}
