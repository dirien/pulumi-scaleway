// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class InstanceServerRootVolumeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the volume where the boot the server
        /// </summary>
        [Input("boot")]
        public Input<bool>? Boot { get; set; }

        /// <summary>
        /// Forces deletion of the root volume on instance termination.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Choose IOPS of your sbs volume, has to be used with `sbs_volume` for root volume type.
        /// 
        /// &gt; **Important:** Updates to `root_volume.size_in_gb` will be ignored after the creation of the server.
        /// </summary>
        [Input("sbsIops")]
        public Input<int>? SbsIops { get; set; }

        /// <summary>
        /// Size of the root volume in gigabytes.
        /// To find the right size use [this endpoint](https://www.scaleway.com/en/developers/api/instance/#path-instances-list-all-instances) and
        /// check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
        /// Updates to this field will recreate a new resource.
        /// </summary>
        [Input("sizeInGb")]
        public Input<int>? SizeInGb { get; set; }

        /// <summary>
        /// The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        /// <summary>
        /// Volume type of root volume, can be `b_ssd`, `l_ssd` or `sbs_volume`, default value depends on server type
        /// </summary>
        [Input("volumeType")]
        public Input<string>? VolumeType { get; set; }

        public InstanceServerRootVolumeArgs()
        {
        }
        public static new InstanceServerRootVolumeArgs Empty => new InstanceServerRootVolumeArgs();
    }
}
