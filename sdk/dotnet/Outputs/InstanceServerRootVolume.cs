// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Outputs
{

    [OutputType]
    public sealed class InstanceServerRootVolume
    {
        public readonly bool? Boot;
        /// <summary>
        /// Forces deletion of the root volume on instance termination.
        /// </summary>
        public readonly bool? DeleteOnTermination;
        /// <summary>
        /// The name of the server.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Size of the root volume in gigabytes.
        /// To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
        /// check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
        /// Updates to this field will recreate a new resource.
        /// </summary>
        public readonly int? SizeInGb;
        /// <summary>
        /// The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
        /// </summary>
        public readonly string? VolumeId;
        /// <summary>
        /// Volume type of root volume, can be `b_ssd` or `l_ssd`, default value depends on server type
        /// </summary>
        public readonly string? VolumeType;

        [OutputConstructor]
        private InstanceServerRootVolume(
            bool? boot,

            bool? deleteOnTermination,

            string? name,

            int? sizeInGb,

            string? volumeId,

            string? volumeType)
        {
            Boot = boot;
            DeleteOnTermination = deleteOnTermination;
            Name = name;
            SizeInGb = sizeInGb;
            VolumeId = volumeId;
            VolumeType = volumeType;
        }
    }
}
