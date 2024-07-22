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
    /// Creates and manages Scaleway Database DocumentDB Users.
    /// 
    /// ## Example Usage
    /// 
    /// ### Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Random = Pulumi.Random;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Scaleway.DocumentdbInstance("instance", new()
    ///     {
    ///         NodeType = "docdb-play2-pico",
    ///         Engine = "FerretDB-1",
    ///         UserName = "my_initial_user",
    ///         Password = "thiZ_is_v&amp;ry_s3cret",
    ///         VolumeSizeInGb = 20,
    ///     });
    /// 
    ///     var dbPassword = new Random.RandomPassword("dbPassword", new()
    ///     {
    ///         Length = 16,
    ///         Special = true,
    ///     });
    /// 
    ///     var dbAdmin = new Scaleway.DocumentdbUser("dbAdmin", new()
    ///     {
    ///         InstanceId = instance.Id,
    ///         Password = dbPassword.Result,
    ///         IsAdmin = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/documentdbUser:DocumentdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/documentdbUser:DocumentdbUser")]
    public partial class DocumentdbUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// UUID of the documentDB instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database User.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Grant admin permissions to the Database User.
        /// </summary>
        [Output("isAdmin")]
        public Output<bool?> IsAdmin { get; private set; } = null!;

        /// <summary>
        /// Database Username.
        /// 
        /// &gt; **Important:** Updates to `name` will recreate the Database User.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Database User password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The Scaleway region this resource resides in.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentdbUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentdbUser(string name, DocumentdbUserArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbUser:DocumentdbUser", name, args ?? new DocumentdbUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentdbUser(string name, Input<string> id, DocumentdbUserState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/documentdbUser:DocumentdbUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DocumentdbUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentdbUser Get(string name, Input<string> id, DocumentdbUserState? state = null, CustomResourceOptions? options = null)
        {
            return new DocumentdbUser(name, id, state, options);
        }
    }

    public sealed class DocumentdbUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the documentDB instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database User.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Grant admin permissions to the Database User.
        /// </summary>
        [Input("isAdmin")]
        public Input<bool>? IsAdmin { get; set; }

        /// <summary>
        /// Database Username.
        /// 
        /// &gt; **Important:** Updates to `name` will recreate the Database User.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Database User password.
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
        /// The Scaleway region this resource resides in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentdbUserArgs()
        {
        }
        public static new DocumentdbUserArgs Empty => new DocumentdbUserArgs();
    }

    public sealed class DocumentdbUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// UUID of the documentDB instance.
        /// 
        /// &gt; **Important:** Updates to `instance_id` will recreate the Database User.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Grant admin permissions to the Database User.
        /// </summary>
        [Input("isAdmin")]
        public Input<bool>? IsAdmin { get; set; }

        /// <summary>
        /// Database Username.
        /// 
        /// &gt; **Important:** Updates to `name` will recreate the Database User.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Database User password.
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
        /// The Scaleway region this resource resides in.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public DocumentdbUserState()
        {
        }
        public static new DocumentdbUserState Empty => new DocumentdbUserState();
    }
}
