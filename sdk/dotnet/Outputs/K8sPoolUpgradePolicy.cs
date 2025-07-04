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
    public sealed class K8sPoolUpgradePolicy
    {
        /// <summary>
        /// The maximum number of nodes to be created during the upgrade
        /// </summary>
        public readonly int? MaxSurge;
        /// <summary>
        /// The maximum number of nodes that can be not ready at the same time
        /// </summary>
        public readonly int? MaxUnavailable;

        [OutputConstructor]
        private K8sPoolUpgradePolicy(
            int? maxSurge,

            int? maxUnavailable)
        {
            MaxSurge = maxSurge;
            MaxUnavailable = maxUnavailable;
        }
    }
}
