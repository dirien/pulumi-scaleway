// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a Load Balancer IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myIp = scaleway.getLbIp({
 *     ipId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getLbIp(args?: GetLbIpArgs, opts?: pulumi.InvokeOptions): Promise<GetLbIpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLbIp:getLbIp", {
        "ipAddress": args.ipAddress,
        "ipId": args.ipId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbIp.
 */
export interface GetLbIpArgs {
    /**
     * The IP address.
     * Only one of `ipAddress` and `ipId` should be specified.
     */
    ipAddress?: string;
    /**
     * The IP ID.
     * Only one of `ipAddress` and `ipId` should be specified.
     */
    ipId?: string;
    /**
     * The ID of the project the LB IP associated with.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getLbIp.
 */
export interface GetLbIpResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddress?: string;
    readonly ipId?: string;
    readonly isIpv6: boolean;
    /**
     * The associated load-balancer ID if any
     */
    readonly lbId: string;
    /**
     * (Defaults to provider `organizationId`) The ID of the organization the LB IP is associated with.
     */
    readonly organizationId: string;
    readonly projectId: string;
    readonly region: string;
    /**
     * The reverse domain associated with this IP.
     */
    readonly reverse: string;
    readonly zone: string;
}
/**
 * Gets information about a Load Balancer IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myIp = scaleway.getLbIp({
 *     ipId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getLbIpOutput(args?: GetLbIpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbIpResult> {
    return pulumi.output(args).apply((a: any) => getLbIp(a, opts))
}

/**
 * A collection of arguments for invoking getLbIp.
 */
export interface GetLbIpOutputArgs {
    /**
     * The IP address.
     * Only one of `ipAddress` and `ipId` should be specified.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The IP ID.
     * Only one of `ipAddress` and `ipId` should be specified.
     */
    ipId?: pulumi.Input<string>;
    /**
     * The ID of the project the LB IP associated with.
     */
    projectId?: pulumi.Input<string>;
}
