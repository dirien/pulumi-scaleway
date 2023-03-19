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
    /// Creates and manages Scaleway Functions.
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
    ///     var mainFunctionNamespace = new Scaleway.FunctionNamespace("mainFunctionNamespace", new()
    ///     {
    ///         Description = "Main function namespace",
    ///     });
    /// 
    ///     var mainFunction = new Scaleway.Function("mainFunction", new()
    ///     {
    ///         NamespaceId = mainFunctionNamespace.Id,
    ///         Runtime = "go118",
    ///         Handler = "Handle",
    ///         Privacy = "private",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Functions can be imported using the `{region}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/function:Function main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/function:Function")]
    public partial class Function : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CPU limit in mCPU for your function. More infos on resources [here](https://developers.scaleway.com/en/products/functions/api/#functions)
        /// </summary>
        [Output("cpuLimit")]
        public Output<int> CpuLimit { get; private set; } = null!;

        /// <summary>
        /// Define if the function should be deployed, terraform will wait for function to be deployed
        /// </summary>
        [Output("deploy")]
        public Output<bool?> Deploy { get; private set; } = null!;

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The native domain name of the function
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The environment variables of the function.
        /// </summary>
        [Output("environmentVariables")]
        public Output<ImmutableDictionary<string, string>?> EnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// HTTP traffic configuration
        /// </summary>
        [Output("httpOption")]
        public Output<string?> HttpOption { get; private set; } = null!;

        /// <summary>
        /// Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
        /// </summary>
        [Output("maxScale")]
        public Output<int?> MaxScale { get; private set; } = null!;

        /// <summary>
        /// Memory limit in MB for your function, defaults to 128MB
        /// </summary>
        [Output("memoryLimit")]
        public Output<int?> MemoryLimit { get; private set; } = null!;

        /// <summary>
        /// Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
        /// </summary>
        [Output("minScale")]
        public Output<int?> MinScale { get; private set; } = null!;

        /// <summary>
        /// The unique name of the function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace ID associated with this function
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// The organization ID the function is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
        /// </summary>
        [Output("privacy")]
        public Output<string> Privacy { get; private set; } = null!;

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
        /// Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
        /// </summary>
        [Output("secretEnvironmentVariables")]
        public Output<ImmutableDictionary<string, string>?> SecretEnvironmentVariables { get; private set; } = null!;

        /// <summary>
        /// Holds the max duration (in seconds) the function is allowed for responding to a request
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;

        /// <summary>
        /// Location of the zip file to upload containing your function sources
        /// </summary>
        [Output("zipFile")]
        public Output<string?> ZipFile { get; private set; } = null!;

        /// <summary>
        /// The hash of your source zip file, changing it will re-apply function. Can be any string
        /// </summary>
        [Output("zipHash")]
        public Output<string?> ZipHash { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/function:Function", name, args ?? new FunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/function:Function", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Define if the function should be deployed, terraform will wait for function to be deployed
        /// </summary>
        [Input("deploy")]
        public Input<bool>? Deploy { get; set; }

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The environment variables of the function.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
        /// </summary>
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        /// <summary>
        /// HTTP traffic configuration
        /// </summary>
        [Input("httpOption")]
        public Input<string>? HttpOption { get; set; }

        /// <summary>
        /// Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
        /// </summary>
        [Input("maxScale")]
        public Input<int>? MaxScale { get; set; }

        /// <summary>
        /// Memory limit in MB for your function, defaults to 128MB
        /// </summary>
        [Input("memoryLimit")]
        public Input<int>? MemoryLimit { get; set; }

        /// <summary>
        /// Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
        /// </summary>
        [Input("minScale")]
        public Input<int>? MinScale { get; set; }

        /// <summary>
        /// The unique name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace ID associated with this function
        /// </summary>
        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        /// <summary>
        /// Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
        /// </summary>
        [Input("privacy", required: true)]
        public Input<string> Privacy { get; set; } = null!;

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
        /// Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        [Input("secretEnvironmentVariables")]
        private InputMap<string>? _secretEnvironmentVariables;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
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

        /// <summary>
        /// Holds the max duration (in seconds) the function is allowed for responding to a request
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Location of the zip file to upload containing your function sources
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        /// <summary>
        /// The hash of your source zip file, changing it will re-apply function. Can be any string
        /// </summary>
        [Input("zipHash")]
        public Input<string>? ZipHash { get; set; }

        public FunctionArgs()
        {
        }
        public static new FunctionArgs Empty => new FunctionArgs();
    }

    public sealed class FunctionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPU limit in mCPU for your function. More infos on resources [here](https://developers.scaleway.com/en/products/functions/api/#functions)
        /// </summary>
        [Input("cpuLimit")]
        public Input<int>? CpuLimit { get; set; }

        /// <summary>
        /// Define if the function should be deployed, terraform will wait for function to be deployed
        /// </summary>
        [Input("deploy")]
        public Input<bool>? Deploy { get; set; }

        /// <summary>
        /// The description of the function.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The native domain name of the function
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// The environment variables of the function.
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// Handler of the function. Depends on the runtime ([function guide](https://developers.scaleway.com/en/products/functions/api/#create-a-function))
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// HTTP traffic configuration
        /// </summary>
        [Input("httpOption")]
        public Input<string>? HttpOption { get; set; }

        /// <summary>
        /// Maximum replicas for your function (defaults to 20), our system will scale your functions automatically based on incoming workload, but will never scale the number of replicas above the configured max_scale.
        /// </summary>
        [Input("maxScale")]
        public Input<int>? MaxScale { get; set; }

        /// <summary>
        /// Memory limit in MB for your function, defaults to 128MB
        /// </summary>
        [Input("memoryLimit")]
        public Input<int>? MemoryLimit { get; set; }

        /// <summary>
        /// Minimum replicas for your function, defaults to 0, Note that a function is billed when it gets executed, and using a min_scale greater than 0 will cause your function container to run constantly.
        /// </summary>
        [Input("minScale")]
        public Input<int>? MinScale { get; set; }

        /// <summary>
        /// The unique name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace ID associated with this function
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// The organization ID the function is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// Privacy of the function. Can be either `private` or `public`. Read more on [authentication](https://developers.scaleway.com/en/products/functions/api/#authentication)
        /// </summary>
        [Input("privacy")]
        public Input<string>? Privacy { get; set; }

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
        /// Runtime of the function. Runtimes can be fetched using [specific route](https://developers.scaleway.com/en/products/functions/api/#get-f7de6a)
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        [Input("secretEnvironmentVariables")]
        private InputMap<string>? _secretEnvironmentVariables;

        /// <summary>
        /// The [secret environment](https://www.scaleway.com/en/docs/compute/functions/concepts/#secrets) variables of the function.
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

        /// <summary>
        /// Holds the max duration (in seconds) the function is allowed for responding to a request
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Location of the zip file to upload containing your function sources
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        /// <summary>
        /// The hash of your source zip file, changing it will re-apply function. Can be any string
        /// </summary>
        [Input("zipHash")]
        public Input<string>? ZipHash { get; set; }

        public FunctionState()
        {
        }
        public static new FunctionState Empty => new FunctionState();
    }
}
