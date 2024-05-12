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
    public sealed class GetCockpitPushUrlResult
    {
        /// <summary>
        /// Push URL for logs (Grafana Loki)
        /// </summary>
        public readonly string PushLogsUrl;
        /// <summary>
        /// Push URL for metrics (Grafana Mimir)
        /// </summary>
        public readonly string PushMetricsUrl;

        [OutputConstructor]
        private GetCockpitPushUrlResult(
            string pushLogsUrl,

            string pushMetricsUrl)
        {
            PushLogsUrl = pushLogsUrl;
            PushMetricsUrl = pushMetricsUrl;
        }
    }
}
