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
    /// The `scaleway.CockpitGrafanaUser` resource allows you to create and manage [Grafana users](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#grafana-users) in Scaleway Cockpit.
    /// 
    /// Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a Grafana user
    /// 
    /// The following command allows you to create a Grafana user within a specific Scaleway Project.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = new Scaleway.AccountProject("project");
    /// 
    ///     var main = new Scaleway.CockpitGrafanaUser("main", new()
    ///     {
    ///         ProjectId = project.Id,
    ///         Login = "my-awesome-user",
    ///         Role = "editor",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This section explains how to import Grafana users using the ID of the Project associated with Cockpit, and the Grafana user ID in the `{project_id}/{grafana_user_id}` format.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser main 11111111-1111-1111-1111-111111111111/2
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser")]
    public partial class CockpitGrafanaUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// URL for Grafana.
        /// </summary>
        [Output("grafanaUrl")]
        public Output<string> GrafanaUrl { get; private set; } = null!;

        /// <summary>
        /// The username of the Grafana user. The `admin` user is not yet available for creation. You need your Grafana username to log in to Grafana and access your dashboards.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// The password of the Grafana user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// ) The ID of the Project the Cockpit is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The role assigned to the Grafana user. Must be `editor` or `viewer`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a CockpitGrafanaUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CockpitGrafanaUser(string name, CockpitGrafanaUserArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser", name, args ?? new CockpitGrafanaUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CockpitGrafanaUser(string name, Input<string> id, CockpitGrafanaUserState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser", name, state, MakeResourceOptions(options, id))
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
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CockpitGrafanaUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CockpitGrafanaUser Get(string name, Input<string> id, CockpitGrafanaUserState? state = null, CustomResourceOptions? options = null)
        {
            return new CockpitGrafanaUser(name, id, state, options);
        }
    }

    public sealed class CockpitGrafanaUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The username of the Grafana user. The `admin` user is not yet available for creation. You need your Grafana username to log in to Grafana and access your dashboards.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        /// <summary>
        /// ) The ID of the Project the Cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The role assigned to the Grafana user. Must be `editor` or `viewer`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public CockpitGrafanaUserArgs()
        {
        }
        public static new CockpitGrafanaUserArgs Empty => new CockpitGrafanaUserArgs();
    }

    public sealed class CockpitGrafanaUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URL for Grafana.
        /// </summary>
        [Input("grafanaUrl")]
        public Input<string>? GrafanaUrl { get; set; }

        /// <summary>
        /// The username of the Grafana user. The `admin` user is not yet available for creation. You need your Grafana username to log in to Grafana and access your dashboards.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password of the Grafana user.
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

        /// <summary>
        /// ) The ID of the Project the Cockpit is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The role assigned to the Grafana user. Must be `editor` or `viewer`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public CockpitGrafanaUserState()
        {
        }
        public static new CockpitGrafanaUserState Empty => new CockpitGrafanaUserState();
    }
}
