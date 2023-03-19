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

    public sealed class DomainRecordWeightedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The weighted IP
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// The weight of the IP as an integer UInt32.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public DomainRecordWeightedArgs()
        {
        }
        public static new DomainRecordWeightedArgs Empty => new DomainRecordWeightedArgs();
    }
}
