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
    /// Creates and manages Scaleway Compute Images.
    /// For more information, see [the documentation](https://www.scaleway.com/en/developers/api/instance/#images-41389b).
    /// 
    /// ## Example Usage
    /// 
    /// ### From a volume
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var volume = new Scaleway.InstanceVolume("volume", new()
    ///     {
    ///         Type = "b_ssd",
    ///         SizeInGb = 20,
    ///     });
    /// 
    ///     var volumeSnapshot = new Scaleway.InstanceSnapshot("volumeSnapshot", new()
    ///     {
    ///         VolumeId = volume.Id,
    ///     });
    /// 
    ///     var volumeImage = new Scaleway.InstanceImage("volumeImage", new()
    ///     {
    ///         RootVolumeId = volumeSnapshot.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### From a server
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Scaleway = ediri.Scaleway;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var server = new Scaleway.InstanceServer("server", new()
    ///     {
    ///         Image = "ubuntu_jammy",
    ///         Type = "DEV1-S",
    ///     });
    /// 
    ///     var serverSnapshot = new Scaleway.InstanceSnapshot("serverSnapshot", new()
    ///     {
    ///         VolumeId = scaleway_instance_server.Main.Root_volume[0].Volume_id,
    ///     });
    /// 
    ///     var serverImage = new Scaleway.InstanceImage("serverImage", new()
    ///     {
    ///         RootVolumeId = serverSnapshot.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Images can be imported using the `{zone}/{id}`, e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import scaleway:index/instanceImage:InstanceImage main fr-par-1/11111111-1111-1111-1111-111111111111
    /// ```
    /// </summary>
    [ScalewayResourceType("scaleway:index/instanceImage:InstanceImage")]
    public partial class InstanceImage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of IDs of the snapshots of the additional volumes to be attached to the image.
        /// 
        /// &gt; **Important:** For now it is only possible to have 1 additional_volume.
        /// </summary>
        [Output("additionalVolumeIds")]
        public Output<string?> AdditionalVolumeIds { get; private set; } = null!;

        /// <summary>
        /// The description of the extra volumes attached to the image.
        /// </summary>
        [Output("additionalVolumes")]
        public Output<ImmutableArray<Outputs.InstanceImageAdditionalVolume>> AdditionalVolumes { get; private set; } = null!;

        /// <summary>
        /// The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        /// </summary>
        [Output("architecture")]
        public Output<string?> Architecture { get; private set; } = null!;

        /// <summary>
        /// Date of the volume creation.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// ID of the server the image is based on (in case it is a backup).
        /// </summary>
        [Output("fromServerId")]
        public Output<string> FromServerId { get; private set; } = null!;

        /// <summary>
        /// Date of volume latest update.
        /// </summary>
        [Output("modificationDate")]
        public Output<string> ModificationDate { get; private set; } = null!;

        /// <summary>
        /// The name of the image. If not provided it will be randomly generated.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The organization ID the image is associated with.
        /// </summary>
        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project the image is associated with.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if the image is public.
        /// </summary>
        [Output("public")]
        public Output<bool?> Public { get; private set; } = null!;

        /// <summary>
        /// The ID of the snapshot of the volume to be used as root in the image.
        /// </summary>
        [Output("rootVolumeId")]
        public Output<string> RootVolumeId { get; private set; } = null!;

        /// <summary>
        /// State of the volume.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A list of tags to apply to the image.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The zone in which the image should be created.
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceImage(string name, InstanceImageArgs args, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceImage:InstanceImage", name, args ?? new InstanceImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceImage(string name, Input<string> id, InstanceImageState? state = null, CustomResourceOptions? options = null)
            : base("scaleway:index/instanceImage:InstanceImage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceImage Get(string name, Input<string> id, InstanceImageState? state = null, CustomResourceOptions? options = null)
        {
            return new InstanceImage(name, id, state, options);
        }
    }

    public sealed class InstanceImageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of IDs of the snapshots of the additional volumes to be attached to the image.
        /// 
        /// &gt; **Important:** For now it is only possible to have 1 additional_volume.
        /// </summary>
        [Input("additionalVolumeIds")]
        public Input<string>? AdditionalVolumeIds { get; set; }

        /// <summary>
        /// The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The name of the image. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project the image is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Set to `true` if the image is public.
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// The ID of the snapshot of the volume to be used as root in the image.
        /// </summary>
        [Input("rootVolumeId", required: true)]
        public Input<string> RootVolumeId { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the image.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone in which the image should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceImageArgs()
        {
        }
        public static new InstanceImageArgs Empty => new InstanceImageArgs();
    }

    public sealed class InstanceImageState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// List of IDs of the snapshots of the additional volumes to be attached to the image.
        /// 
        /// &gt; **Important:** For now it is only possible to have 1 additional_volume.
        /// </summary>
        [Input("additionalVolumeIds")]
        public Input<string>? AdditionalVolumeIds { get; set; }

        [Input("additionalVolumes")]
        private InputList<Inputs.InstanceImageAdditionalVolumeGetArgs>? _additionalVolumes;

        /// <summary>
        /// The description of the extra volumes attached to the image.
        /// </summary>
        public InputList<Inputs.InstanceImageAdditionalVolumeGetArgs> AdditionalVolumes
        {
            get => _additionalVolumes ?? (_additionalVolumes = new InputList<Inputs.InstanceImageAdditionalVolumeGetArgs>());
            set => _additionalVolumes = value;
        }

        /// <summary>
        /// The architecture the image is compatible with. Possible values are: `x86_64` or `arm`.
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// Date of the volume creation.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// ID of the server the image is based on (in case it is a backup).
        /// </summary>
        [Input("fromServerId")]
        public Input<string>? FromServerId { get; set; }

        /// <summary>
        /// Date of volume latest update.
        /// </summary>
        [Input("modificationDate")]
        public Input<string>? ModificationDate { get; set; }

        /// <summary>
        /// The name of the image. If not provided it will be randomly generated.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The organization ID the image is associated with.
        /// </summary>
        [Input("organizationId")]
        public Input<string>? OrganizationId { get; set; }

        /// <summary>
        /// The ID of the project the image is associated with.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Set to `true` if the image is public.
        /// </summary>
        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// The ID of the snapshot of the volume to be used as root in the image.
        /// </summary>
        [Input("rootVolumeId")]
        public Input<string>? RootVolumeId { get; set; }

        /// <summary>
        /// State of the volume.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tags to apply to the image.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The zone in which the image should be created.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public InstanceImageState()
        {
        }
        public static new InstanceImageState Empty => new InstanceImageState();
    }
}
