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
    /// Creates and manages Scaleway Compute Baremetal servers. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).
    /// 
    /// ## Examples
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = Pulumi.Scaleway;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = Scaleway.GetAccountSshKey.Invoke(new()
    ///     {
    ///         Name = "main",
    ///     });
    /// 
    ///     var @base = new Scaleway.BaremetalServer("base", new()
    ///     {
    ///         Zone = "fr-par-2",
    ///         Offer = "GP-BM1-S",
    ///         Os = "d17d6872-0412-45d9-a198-af82c34d3c5c",
    ///         SshKeyIds = new[]
    ///         {
    ///             main.Apply(getAccountSshKeyResult =&gt; getAccountSshKeyResult.Id),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Baremetal servers can be imported using the `{zone}/{id}`, e.g. bash
    /// 
    /// ```sh
    ///  $ pulumi import scaleway:index/baremetalServer:BaremetalServer web fr-par-2/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/baremetalServer:BaremetalServer")]
    public partial class BaremetalServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the server.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The domain of the server.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The hostname of the server.
        /// </summary>
        [Output("hostname")]
        public Output<string?> Hostname { get; private set; } = null!;

        /// <summary>
        /// (List of) The IPs of the server.
        /// </summary>
        [Output("ips")]
        public Output<ImmutableArray<Outputs.BaremetalServerIp>> Ips { get; private set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The offer name or UUID of the baremetal server.
        /// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
        /// </summary>
        [Output("offer")]
        public Output<string> Offer { get; private set; } = null!;

        /// <summary>
        /// The ID of the offer.
        /// </summary>
        [Output("offerId")]
        public Output<string> OfferId { get; private set; } = null!;

        /// <summary>
        /// The name of the offer.
        /// </summary>
        [Output("offerName")]
        public Output<string> OfferName { get; private set; } = null!;

        /// <summary>
        /// The options to enable on the server.
        /// &gt; The `options` block supports:
        /// </summary>
        [Output("options")]
        public Output<ImmutableArray<Outputs.BaremetalServerOption>> Options { get; private set; } = null!;

        /// <summary>
        /// The organization ID the server is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The UUID of the os to install on the server.
        /// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
        /// &gt; **Important:** Updates to `os` will reinstall the server.
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// The name of the os.
        /// </summary>
        [Output("osName")]
        public Output<string> OsName { get; private set; } = null!;

        /// <summary>
        /// Password used for the installation. May be required depending on used os.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
        /// </summary>
        [Output("privateNetworks")]
        public Output<ImmutableArray<Outputs.BaremetalServerPrivateNetwork>> PrivateNetworks { get; private set; } = null!;

        /// <summary>
        /// `project_id`) The ID of the project the server is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// If True, this boolean allows to reinstall the server on install config changes.
        /// &gt; **Important:** Updates to `ssh_key_ids`, `user`, `password`, `service_user` or `service_password` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
        /// </summary>
        [Output("reinstallOnConfigChanges")]
        public Output<bool?> ReinstallOnConfigChanges { get; private set; } = null!;

        /// <summary>
        /// Password used for the service to install. May be required depending on used os.
        /// </summary>
        [Output("servicePassword")]
        public Output<string?> ServicePassword { get; private set; } = null!;

        /// <summary>
        /// User used for the service to install.
        /// </summary>
        [Output("serviceUser")]
        public Output<string> ServiceUser { get; private set; } = null!;

        /// <summary>
        /// List of SSH keys allowed to connect to the server.
        /// </summary>
        [Output("sshKeyIds")]
        public Output<ImmutableArray<string>> SshKeyIds { get; private set; } = null!;

        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// User used for the installation.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a BaremetalServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BaremetalServer(string name, BaremetalServerArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/baremetalServer:BaremetalServer", name, args ?? new BaremetalServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BaremetalServer(string name, Input<string> id, BaremetalServerState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/baremetalServer:BaremetalServer", name, state, MakeResourceOptions(options, id))
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
                    "password",
                    "servicePassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BaremetalServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BaremetalServer Get(string name, Input<string> id, BaremetalServerState? state = null, CustomResourceOptions? options = null)
        {
            return new BaremetalServer(name, id, state, options);
        }
    }

    public sealed class BaremetalServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the server.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The hostname of the server.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The offer name or UUID of the baremetal server.
        /// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
        /// </summary>
        [Input("offer", required: true)]
        public Input<string> Offer { get; set; } = null!;

        [Input("options")]
        private InputList<Inputs.BaremetalServerOptionArgs>? _options;

        /// <summary>
        /// The options to enable on the server.
        /// &gt; The `options` block supports:
        /// </summary>
        public InputList<Inputs.BaremetalServerOptionArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.BaremetalServerOptionArgs>());
            set => _options = value;
        }

        /// <summary>
        /// The UUID of the os to install on the server.
        /// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
        /// &gt; **Important:** Updates to `os` will reinstall the server.
        /// </summary>
        [Input("os", required: true)]
        public Input<string> Os { get; set; } = null!;

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password used for the installation. May be required depending on used os.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateNetworks")]
        private InputList<Inputs.BaremetalServerPrivateNetworkArgs>? _privateNetworks;

        /// <summary>
        /// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
        /// </summary>
        public InputList<Inputs.BaremetalServerPrivateNetworkArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.BaremetalServerPrivateNetworkArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the server is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// If True, this boolean allows to reinstall the server on install config changes.
        /// &gt; **Important:** Updates to `ssh_key_ids`, `user`, `password`, `service_user` or `service_password` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
        /// </summary>
        [Input("reinstallOnConfigChanges")]
        public Input<bool>? ReinstallOnConfigChanges { get; set; }

        [Input("servicePassword")]
        private Input<string>? _servicePassword;

        /// <summary>
        /// Password used for the service to install. May be required depending on used os.
        /// </summary>
        public Input<string>? ServicePassword
        {
            get => _servicePassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _servicePassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// User used for the service to install.
        /// </summary>
        [Input("serviceUser")]
        public Input<string>? ServiceUser { get; set; }

        [Input("sshKeyIds", required: true)]
        private InputList<string>? _sshKeyIds;

        /// <summary>
        /// List of SSH keys allowed to connect to the server.
        /// </summary>
        public InputList<string> SshKeyIds
        {
            get => _sshKeyIds ?? (_sshKeyIds = new InputList<string>());
            set => _sshKeyIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// User used for the installation.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public BaremetalServerArgs()
        {
        }
        public static new BaremetalServerArgs Empty => new BaremetalServerArgs();
    }

    public sealed class BaremetalServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the server.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The domain of the server.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The hostname of the server.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("ips")]
        private InputList<Inputs.BaremetalServerIpGetArgs>? _ips;

        /// <summary>
        /// (List of) The IPs of the server.
        /// </summary>
        public InputList<Inputs.BaremetalServerIpGetArgs> Ips
        {
            get => _ips ?? (_ips = new InputList<Inputs.BaremetalServerIpGetArgs>());
            set => _ips = value;
        }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The offer name or UUID of the baremetal server.
        /// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.
        /// </summary>
        [Input("offer")]
        public Input<string>? Offer { get; set; }

        /// <summary>
        /// The ID of the offer.
        /// </summary>
        [Input("offerId")]
        public Input<string>? OfferId { get; set; }

        /// <summary>
        /// The name of the offer.
        /// </summary>
        [Input("offerName")]
        public Input<string>? OfferName { get; set; }

        [Input("options")]
        private InputList<Inputs.BaremetalServerOptionGetArgs>? _options;

        /// <summary>
        /// The options to enable on the server.
        /// &gt; The `options` block supports:
        /// </summary>
        public InputList<Inputs.BaremetalServerOptionGetArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.BaremetalServerOptionGetArgs>());
            set => _options = value;
        }

        /// <summary>
        /// The organization ID the server is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The UUID of the os to install on the server.
        /// Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.
        /// &gt; **Important:** Updates to `os` will reinstall the server.
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// The name of the os.
        /// </summary>
        [Input("osName")]
        public Input<string>? OsName { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password used for the installation. May be required depending on used os.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("privateNetworks")]
        private InputList<Inputs.BaremetalServerPrivateNetworkGetArgs>? _privateNetworks;

        /// <summary>
        /// The private networks to attach to the server. For more information, see [the documentation](https://www.scaleway.com/en/docs/compute/elastic-metal/how-to/use-private-networks/)
        /// </summary>
        public InputList<Inputs.BaremetalServerPrivateNetworkGetArgs> PrivateNetworks
        {
            get => _privateNetworks ?? (_privateNetworks = new InputList<Inputs.BaremetalServerPrivateNetworkGetArgs>());
            set => _privateNetworks = value;
        }

        /// <summary>
        /// `project_id`) The ID of the project the server is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// If True, this boolean allows to reinstall the server on install config changes.
        /// &gt; **Important:** Updates to `ssh_key_ids`, `user`, `password`, `service_user` or `service_password` will not take effect on the server, it requires to reinstall it. To do so please set 'reinstall_on_config_changes' argument to true.
        /// </summary>
        [Input("reinstallOnConfigChanges")]
        public Input<bool>? ReinstallOnConfigChanges { get; set; }

        [Input("servicePassword")]
        private Input<string>? _servicePassword;

        /// <summary>
        /// Password used for the service to install. May be required depending on used os.
        /// </summary>
        public Input<string>? ServicePassword
        {
            get => _servicePassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _servicePassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// User used for the service to install.
        /// </summary>
        [Input("serviceUser")]
        public Input<string>? ServiceUser { get; set; }

        [Input("sshKeyIds")]
        private InputList<string>? _sshKeyIds;

        /// <summary>
        /// List of SSH keys allowed to connect to the server.
        /// </summary>
        public InputList<string> SshKeyIds
        {
            get => _sshKeyIds ?? (_sshKeyIds = new InputList<string>());
            set => _sshKeyIds = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags associated with the server.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// User used for the installation.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// `zone`) The zone in which the server should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public BaremetalServerState()
        {
        }
        public static new BaremetalServerState Empty => new BaremetalServerState();
    }
}
