// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getDocumentDBLoadBalancerEndpoint(args?: GetDocumentDBLoadBalancerEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentDBLoadBalancerEndpointResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getDocumentDBLoadBalancerEndpoint:getDocumentDBLoadBalancerEndpoint", {
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDocumentDBLoadBalancerEndpoint.
 */
export interface GetDocumentDBLoadBalancerEndpointArgs {
    instanceId?: string;
    instanceName?: string;
    projectId?: string;
    region?: string;
}

/**
 * A collection of values returned by getDocumentDBLoadBalancerEndpoint.
 */
export interface GetDocumentDBLoadBalancerEndpointResult {
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly instanceName: string;
    readonly ip: string;
    readonly name: string;
    readonly port: number;
    readonly projectId: string;
    readonly region: string;
}
export function getDocumentDBLoadBalancerEndpointOutput(args?: GetDocumentDBLoadBalancerEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentDBLoadBalancerEndpointResult> {
    return pulumi.output(args).apply((a: any) => getDocumentDBLoadBalancerEndpoint(a, opts))
}

/**
 * A collection of arguments for invoking getDocumentDBLoadBalancerEndpoint.
 */
export interface GetDocumentDBLoadBalancerEndpointOutputArgs {
    instanceId?: pulumi.Input<string>;
    instanceName?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}