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

    public sealed class CockpitPushUrlGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Push URL for logs (Grafana Loki)
        /// </summary>
        [Input("pushLogsUrl")]
        public Input<string>? PushLogsUrl { get; set; }

        /// <summary>
        /// Push URL for metrics (Grafana Mimir)
        /// </summary>
        [Input("pushMetricsUrl")]
        public Input<string>? PushMetricsUrl { get; set; }

        public CockpitPushUrlGetArgs()
        {
        }
        public static new CockpitPushUrlGetArgs Empty => new CockpitPushUrlGetArgs();
    }
}