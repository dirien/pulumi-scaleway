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

    public sealed class IotDeviceMessageFiltersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rules used to restrict topics the device can publish to.
        /// </summary>
        [Input("publish")]
        public Input<Inputs.IotDeviceMessageFiltersPublishArgs>? Publish { get; set; }

        /// <summary>
        /// Rules used to restrict topics the device can subscribe to.
        /// </summary>
        [Input("subscribe")]
        public Input<Inputs.IotDeviceMessageFiltersSubscribeArgs>? Subscribe { get; set; }

        public IotDeviceMessageFiltersArgs()
        {
        }
        public static new IotDeviceMessageFiltersArgs Empty => new IotDeviceMessageFiltersArgs();
    }
}
