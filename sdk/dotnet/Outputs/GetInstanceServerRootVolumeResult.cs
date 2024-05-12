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
    public sealed class GetInstanceServerRootVolumeResult
    {
        /// <summary>
        /// Set the volume where the boot the server
        /// </summary>
        public readonly bool Boot;
        /// <summary>
        /// Forces deletion of the root volume on instance termination.
        /// </summary>
        public readonly bool DeleteOnTermination;
        /// <summary>
        /// The server name. Only one of `name` and `server_id` should be specified.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Size of the root volume in gigabytes.
        /// </summary>
        public readonly int SizeInGb;
        /// <summary>
        /// The volume ID of the root volume of the server.
        /// </summary>
        public readonly string VolumeId;
        /// <summary>
        /// Volume type of the root volume
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private GetInstanceServerRootVolumeResult(
            bool boot,

            bool deleteOnTermination,

            string name,

            int sizeInGb,

            string volumeId,

            string volumeType)
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
