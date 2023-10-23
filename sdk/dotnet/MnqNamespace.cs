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
    /// Creates and manages Scaleway Messaging and queuing Namespace.
    /// For further information please check
    /// our [documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-index)
    /// 
    /// &gt; NOTE: This resource refers to the old version of the MNQ API. You should use new resources dedicated to your protocol. SQS, NATS.
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
    ///     var main = new Scaleway.MnqNamespace("main", new()
    ///     {
    ///         Protocol = "nats",
    ///         Region = "fr-par",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Namespaces can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/mnqNamespace:MnqNamespace main fr-par/11111111111111111111111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/mnqNamespace:MnqNamespace")]
    public partial class MnqNamespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The date and time the Namespace was created.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// The endpoint of the service matching the Namespace protocol.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// The unique name of the namespace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the
        /// namespace is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The protocol to apply on your namespace. Please check our
        /// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which the namespace should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The date and time the Namespace was updated.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a MnqNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MnqNamespace(string name, MnqNamespaceArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqNamespace:MnqNamespace", name, args ?? new MnqNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MnqNamespace(string name, Input<string> id, MnqNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/mnqNamespace:MnqNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MnqNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MnqNamespace Get(string name, Input<string> id, MnqNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new MnqNamespace(name, id, state, options);
        }
    }

    public sealed class MnqNamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique name of the namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the
        /// namespace is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The protocol to apply on your namespace. Please check our
        /// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// `region`). The region
        /// in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public MnqNamespaceArgs()
        {
        }
        public static new MnqNamespaceArgs Empty => new MnqNamespaceArgs();
    }

    public sealed class MnqNamespaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The date and time the Namespace was created.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// The endpoint of the service matching the Namespace protocol.
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// The unique name of the namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the
        /// namespace is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The protocol to apply on your namespace. Please check our
        /// supported [protocols](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/mnq/v1alpha1#pkg-constants).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// `region`). The region
        /// in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The date and time the Namespace was updated.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public MnqNamespaceState()
        {
        }
        public static new MnqNamespaceState Empty => new MnqNamespaceState();
    }
}
