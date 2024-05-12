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

    public sealed class IotDeviceMessageFiltersPublishGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Publish message filter policy
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("topics")]
        private InputList<string>? _topics;

        /// <summary>
        /// List of topics in the set
        /// </summary>
        public InputList<string> Topics
        {
            get => _topics ?? (_topics = new InputList<string>());
            set => _topics = value;
        }

        public IotDeviceMessageFiltersPublishGetArgs()
        {
        }
        public static new IotDeviceMessageFiltersPublishGetArgs Empty => new IotDeviceMessageFiltersPublishGetArgs();
    }
}
