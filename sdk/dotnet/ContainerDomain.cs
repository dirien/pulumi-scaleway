// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    /// The `scaleway.ContainerDomain` resource allows you to create and manage domain name bindings for Scaleway [Serverless Containers](https://www.scaleway.com/en/docs/serverless/containers/).
    /// 
    /// Refer to the Containers domain [documentation](https://www.scaleway.com/en/docs/serverless-containers/how-to/add-a-custom-domain-to-a-container/) and the [API documentation](https://www.scaleway.com/en/developers/api/serverless-containers/#path-domains-list-all-domain-name-bindings) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// The commands below shows how to bind a custom domain name to a container.
    /// 
    /// ### Simple
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var appContainer = new Scaleway.Container("appContainer");
    /// 
    ///     var appContainerDomain = new Scaleway.ContainerDomain("appContainerDomain", new()
    ///     {
    ///         ContainerId = appContainer.Id,
    ///         Hostname = "container.domain.tld",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Complete example with domain
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Scaleway.ContainerNamespace("main", new()
    ///     {
    ///         Description = "test container",
    ///     });
    /// 
    ///     var appContainer = new Scaleway.Container("appContainer", new()
    ///     {
    ///         NamespaceId = main.Id,
    ///         RegistryImage = main.RegistryEndpoint.Apply(registryEndpoint =&gt; $"{registryEndpoint}/nginx:alpine"),
    ///         Port = 80,
    ///         CpuLimit = 140,
    ///         MemoryLimit = 256,
    ///         MinScale = 1,
    ///         MaxScale = 1,
    ///         Timeout = 600,
    ///         MaxConcurrency = 80,
    ///         Privacy = "public",
    ///         Protocol = "http1",
    ///         Deploy = true,
    ///     });
    /// 
    ///     var appDomainRecord = new Scaleway.DomainRecord("appDomainRecord", new()
    ///     {
    ///         DnsZone = "domain.tld",
    ///         Type = "CNAME",
    ///         Data = appContainer.DomainName.Apply(domainName =&gt; $"{domainName}."),
    ///         Ttl = 3600,
    ///     });
    /// 
    ///     var appContainerDomain = new Scaleway.ContainerDomain("appContainerDomain", new()
    ///     {
    ///         ContainerId = appContainer.Id,
    ///         Hostname = Output.Tuple(appDomainRecord.Name, appDomainRecord.DnsZone).Apply(values =&gt;
    ///         {
    ///             var name = values.Item1;
    ///             var dnsZone = values.Item2;
    ///             return $"{name}.{dnsZone}";
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Container domain binding can be imported using `{region}/{id}`, as shown below:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/containerDomain:ContainerDomain main fr-par/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/containerDomain:ContainerDomain")]
    public partial class ContainerDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unique identifier of the container.
        /// </summary>
        [Output("containerId")]
        public Output<string> ContainerId { get; private set; } = null!;

        /// <summary>
        /// The hostname with a CNAME record.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// `region`) The region in which the container exists.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URL used to query the container.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerDomain(string name, ContainerDomainArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/containerDomain:ContainerDomain", name, args ?? new ContainerDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerDomain(string name, Input<string> id, ContainerDomainState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/containerDomain:ContainerDomain", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerDomain Get(string name, Input<string> id, ContainerDomainState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerDomain(name, id, state, options);
        }
    }

    public sealed class ContainerDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the container.
        /// </summary>
        [Input("containerId", required: true)]
        public Input<string> ContainerId { get; set; } = null!;

        /// <summary>
        /// The hostname with a CNAME record.
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// `region`) The region in which the container exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ContainerDomainArgs()
        {
        }
        public static new ContainerDomainArgs Empty => new ContainerDomainArgs();
    }

    public sealed class ContainerDomainState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique identifier of the container.
        /// </summary>
        [Input("containerId")]
        public Input<string>? ContainerId { get; set; }

        /// <summary>
        /// The hostname with a CNAME record.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// `region`) The region in which the container exists.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URL used to query the container.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public ContainerDomainState()
        {
        }
        public static new ContainerDomainState Empty => new ContainerDomainState();
    }
}
