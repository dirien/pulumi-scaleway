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
    /// Creates and manages Scaleway Function Namespace.
    /// For more information see [the documentation](https://developers.scaleway.com/en/products/functions/api/).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.FunctionNamespace("main", new()
    ///     {
    ///         Description = "Main function namespace",
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
    ///  $ pulumi import scaleway:index/functionNamespace:FunctionNamespace main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/functionNamespace:FunctionNamespace")]
    public partial class FunctionNamespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the namespace.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The environment variables of the namespace.
        /// </summary>
        [Output("environmentVariables")]
        public Output<ImmutableDictionary<string, string>?> EnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// The unique name of the function namespace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the namespace is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the namespace is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The registry endpoint of the namespace.
        /// </summary>
        [Output("registryEndpoint")]
        public Output<string> RegistryEndpoint { get; private set; } = null!;

        /// <summary>
        /// The registry namespace ID of the namespace.
        /// </summary>
        [Output("registryNamespaceId")]
        public Output<string> RegistryNamespaceId { get; private set; } = null!;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the namespace.
        /// </summary>
        [Output("secretEnvironmentVariables")]
        public Output<ImmutableDictionary<string, string>?> SecretEnvironmentVariables { get; private set; } = null!;


        /// <summary>
        /// Create a FunctionNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FunctionNamespace(string name, FunctionNamespaceArgs? args = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionNamespace:FunctionNamespace", name, args ?? new FunctionNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FunctionNamespace(string name, Input<string> id, FunctionNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/functionNamespace:FunctionNamespace", name, state, MakeResourceOptions(options, id))
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
                    "secretEnvironmentVariables",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FunctionNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FunctionNamespace Get(string name, Input<string> id, FunctionNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new FunctionNamespace(name, id, state, options);
        }
    }

    public sealed class FunctionNamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the namespace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The environment variables of the namespace.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The unique name of the function namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the namespace is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretEnvironmentVariables")]
        private InputMap<string>? _secretEnvironmentVariables;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the namespace.
        /// </summary>
        public InputMap<string> SecretEnvironmentVariables
        {
            get => _secretEnvironmentVariables ?? (_secretEnvironmentVariables = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _secretEnvironmentVariables = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        public FunctionNamespaceArgs()
        {
        }
        public static new FunctionNamespaceArgs Empty => new FunctionNamespaceArgs();
    }

    public sealed class FunctionNamespaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the namespace.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The environment variables of the namespace.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// The unique name of the function namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the namespace is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// `project_id`) The ID of the project the namespace is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// `region`). The region in which the namespace should be created.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The registry endpoint of the namespace.
        /// </summary>
        [Input("registryEndpoint")]
        public Input<string>? RegistryEndpoint { get; set; }

        /// <summary>
        /// The registry namespace ID of the namespace.
        /// </summary>
        [Input("registryNamespaceId")]
        public Input<string>? RegistryNamespaceId { get; set; }

        [Input("secretEnvironmentVariables")]
        private InputMap<string>? _secretEnvironmentVariables;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/containers/concepts/#secrets) variables of the namespace.
        /// </summary>
        public InputMap<string> SecretEnvironmentVariables
        {
            get => _secretEnvironmentVariables ?? (_secretEnvironmentVariables = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _secretEnvironmentVariables = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        public FunctionNamespaceState()
        {
        }
        public static new FunctionNamespaceState Empty => new FunctionNamespaceState();
    }
}
