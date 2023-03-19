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

    public sealed class MnqCredentialSqsSnsCredentialsPermissionsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// . Defines if user can manage the associated resource(s).
        /// </summary>
        [Input("canManage")]
        public Input<bool>? CanManage { get; set; }

        /// <summary>
        /// . Defines if user can publish messages to the service.
        /// </summary>
        [Input("canPublish")]
        public Input<bool>? CanPublish { get; set; }

        /// <summary>
        /// . Defines if user can receive messages from the service.
        /// </summary>
        [Input("canReceive")]
        public Input<bool>? CanReceive { get; set; }

        public MnqCredentialSqsSnsCredentialsPermissionsGetArgs()
        {
        }
        public static new MnqCredentialSqsSnsCredentialsPermissionsGetArgs Empty => new MnqCredentialSqsSnsCredentialsPermissionsGetArgs();
    }
}
