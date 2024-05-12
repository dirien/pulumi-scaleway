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

    public sealed class JobDefinitionCronArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cron format string.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<string> Schedule { get; set; } = null!;

        /// <summary>
        /// The timezone, must be a canonical TZ identifier as found in this [list](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
        /// </summary>
        [Input("timezone", required: true)]
        public Input<string> Timezone { get; set; } = null!;

        public JobDefinitionCronArgs()
        {
        }
        public static new JobDefinitionCronArgs Empty => new JobDefinitionCronArgs();
    }
}
