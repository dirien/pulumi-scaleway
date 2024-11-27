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
    public static class GetBlockVolume
    {
        /// <summary>
        /// The `scaleway.BlockVolume` data source is used to retrieve information about a Block Storage volume.
        /// Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/storage/block/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.
        /// 
        /// ## Retrieve a Block Storage volume
        /// 
        /// The following commands allow you to:
        /// 
        /// - retrieve a Block Storage volume specified by its name
        /// - retrieve a Block Storage volume specified by its ID
        /// 
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myVolume = Scaleway.GetBlockVolume.Invoke(new()
        ///     {
        ///         VolumeId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBlockVolumeResult> InvokeAsync(GetBlockVolumeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBlockVolumeResult>("scaleway:index/getBlockVolume:getBlockVolume", args ?? new GetBlockVolumeArgs(), options.WithDefaults());

        /// <summary>
        /// The `scaleway.BlockVolume` data source is used to retrieve information about a Block Storage volume.
        /// Refer to the Block Storage [product documentation](https://www.scaleway.com/en/docs/storage/block/) and [API documentation](https://www.scaleway.com/en/developers/api/block/) for more information.
        /// 
        /// ## Retrieve a Block Storage volume
        /// 
        /// The following commands allow you to:
        /// 
        /// - retrieve a Block Storage volume specified by its name
        /// - retrieve a Block Storage volume specified by its ID
        /// 
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Scaleway = Pulumi.Scaleway;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myVolume = Scaleway.GetBlockVolume.Invoke(new()
        ///     {
        ///         VolumeId = "11111111-1111-1111-1111-111111111111",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBlockVolumeResult> Invoke(GetBlockVolumeInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBlockVolumeResult>("scaleway:index/getBlockVolume:getBlockVolume", args ?? new GetBlockVolumeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBlockVolumeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the volume. Only one of `name` and `volume_id` should be specified.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The unique identifier of the Project to which the volume is associated.
        /// </summary>
        [Input("projectId")]
        public string? ProjectId { get; set; }

        /// <summary>
        /// The unique identifier of the volume. Only one of `name` and `volume_id` should be specified.
        /// </summary>
        [Input("volumeId")]
        public string? VolumeId { get; set; }

        /// <summary>
        /// ). The zone in which the volume exists.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetBlockVolumeArgs()
        {
        }
        public static new GetBlockVolumeArgs Empty => new GetBlockVolumeArgs();
    }

    public sealed class GetBlockVolumeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the volume. Only one of `name` and `volume_id` should be specified.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The unique identifier of the Project to which the volume is associated.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The unique identifier of the volume. Only one of `name` and `volume_id` should be specified.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        /// <summary>
        /// ). The zone in which the volume exists.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetBlockVolumeInvokeArgs()
        {
        }
        public static new GetBlockVolumeInvokeArgs Empty => new GetBlockVolumeInvokeArgs();
    }


    [OutputType]
    public sealed class GetBlockVolumeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int Iops;
        public readonly string? Name;
        public readonly string? ProjectId;
        public readonly int SizeInGb;
        public readonly string SnapshotId;
        public readonly ImmutableArray<string> Tags;
        public readonly string? VolumeId;
        public readonly string? Zone;

        [OutputConstructor]
        private GetBlockVolumeResult(
            string id,

            int iops,

            string? name,

            string? projectId,

            int sizeInGb,

            string snapshotId,

            ImmutableArray<string> tags,

            string? volumeId,

            string? zone)
        {
            Id = id;
            Iops = iops;
            Name = name;
            ProjectId = projectId;
            SizeInGb = sizeInGb;
            SnapshotId = snapshotId;
            Tags = tags;
            VolumeId = volumeId;
            Zone = zone;
        }
    }
}
