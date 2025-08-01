// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
    public sealed class GetLbIpsIpResult
    {
        /// <summary>
        /// The ID of the associated IP.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP address
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The ID of the associated Load BalancerD, if any
        /// </summary>
        public readonly string LbId;
        /// <summary>
        /// The ID of the Organization the Load Balancer is associated with.
        /// </summary>
        public readonly string OrganizationId;
        /// <summary>
        /// The ID of the Project the Load Balancer is associated with.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The reverse domain associated with this IP.
        /// </summary>
        public readonly string Reverse;
        /// <summary>
        /// List of tags used as filter. IPs with these exact tags are listed.
        /// </summary>
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// `zone`) The zone in which the IPs exist.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetLbIpsIpResult(
            string id,

            string ipAddress,

            string lbId,

            string organizationId,

            string projectId,

            string reverse,

            ImmutableArray<string> tags,

            string zone)
        {
            Id = id;
            IpAddress = ipAddress;
            LbId = lbId;
            OrganizationId = organizationId;
            ProjectId = projectId;
            Reverse = reverse;
            Tags = tags;
            Zone = zone;
        }
    }
}
