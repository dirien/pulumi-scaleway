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
    public sealed class WebHostingCpanelUrl
    {
        /// <summary>
        /// The URL of the Dashboard.
        /// </summary>
        public readonly string? Dashboard;
        /// <summary>
        /// The URL of the Webmail interface.
        /// </summary>
        public readonly string? Webmail;

        [OutputConstructor]
        private WebHostingCpanelUrl(
            string? dashboard,

            string? webmail)
        {
            Dashboard = dashboard;
            Webmail = webmail;
        }
    }
}
