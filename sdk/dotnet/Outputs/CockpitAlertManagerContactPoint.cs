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
    public sealed class CockpitAlertManagerContactPoint
    {
        /// <summary>
        /// Email addresses for the alert receivers
        /// </summary>
        public readonly string? Email;

        [OutputConstructor]
        private CockpitAlertManagerContactPoint(string? email)
        {
            Email = email;
        }
    }
}
