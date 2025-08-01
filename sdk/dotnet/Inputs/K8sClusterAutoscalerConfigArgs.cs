// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Scaleway.Inputs
{

    public sealed class K8sClusterAutoscalerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Detect similar node groups and balance the number of nodes between them.
        /// </summary>
        [Input("balanceSimilarNodeGroups")]
        public Input<bool>? BalanceSimilarNodeGroups { get; set; }

        /// <summary>
        /// Disables the scale down feature of the autoscaler.
        /// </summary>
        [Input("disableScaleDown")]
        public Input<bool>? DisableScaleDown { get; set; }

        /// <summary>
        /// Type of resource estimator to be used in scale up.
        /// </summary>
        [Input("estimator")]
        public Input<string>? Estimator { get; set; }

        /// <summary>
        /// Type of node group expander to be used in scale up.
        /// </summary>
        [Input("expander")]
        public Input<string>? Expander { get; set; }

        /// <summary>
        /// Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
        /// </summary>
        [Input("expendablePodsPriorityCutoff")]
        public Input<int>? ExpendablePodsPriorityCutoff { get; set; }

        /// <summary>
        /// Ignore DaemonSet pods when calculating resource utilization for scaling down.
        /// </summary>
        [Input("ignoreDaemonsetsUtilization")]
        public Input<bool>? IgnoreDaemonsetsUtilization { get; set; }

        /// <summary>
        /// Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
        /// </summary>
        [Input("maxGracefulTerminationSec")]
        public Input<int>? MaxGracefulTerminationSec { get; set; }

        /// <summary>
        /// How long after scale up that scale down evaluation resumes.
        /// </summary>
        [Input("scaleDownDelayAfterAdd")]
        public Input<string>? ScaleDownDelayAfterAdd { get; set; }

        /// <summary>
        /// How long a node should be unneeded before it is eligible for scale down.
        /// </summary>
        [Input("scaleDownUnneededTime")]
        public Input<string>? ScaleDownUnneededTime { get; set; }

        /// <summary>
        /// Node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
        /// </summary>
        [Input("scaleDownUtilizationThreshold")]
        public Input<double>? ScaleDownUtilizationThreshold { get; set; }

        public K8sClusterAutoscalerConfigArgs()
        {
        }
        public static new K8sClusterAutoscalerConfigArgs Empty => new K8sClusterAutoscalerConfigArgs();
    }
}
