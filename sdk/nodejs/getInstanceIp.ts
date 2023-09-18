// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an instance IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myIp = scaleway.getInstanceIp({
 *     id: "fr-par-1/11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceIp(args?: GetInstanceIpArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceIpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getInstanceIp:getInstanceIp", {
        "address": args.address,
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceIp.
 */
export interface GetInstanceIpArgs {
    /**
     * The IPv4 address to retrieve
     * Only one of `address` and `id` should be specified.
     */
    address?: string;
    /**
     * The ID of the IP address to retrieve
     * Only one of `address` and `id` should be specified.
     */
    id?: string;
}

/**
 * A collection of values returned by getInstanceIp.
 */
export interface GetInstanceIpResult {
    /**
     * The IP address.
     */
    readonly address?: string;
    /**
     * The ID of the IP.
     */
    readonly id?: string;
    /**
     * The organization ID the IP is associated with.
     */
    readonly organizationId: string;
    /**
     * The IP Prefix.
     */
    readonly prefix: string;
    readonly projectId: string;
    /**
     * The reverse dns attached to this IP
     */
    readonly reverse: string;
    readonly serverId: string;
    readonly tags: string[];
    /**
     * The type of the IP
     */
    readonly type: string;
    readonly zone: string;
}
/**
 * Gets information about an instance IP.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myIp = scaleway.getInstanceIp({
 *     id: "fr-par-1/11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getInstanceIpOutput(args?: GetInstanceIpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceIpResult> {
    return pulumi.output(args).apply((a: any) => getInstanceIp(a, opts))
}

/**
 * A collection of arguments for invoking getInstanceIp.
 */
export interface GetInstanceIpOutputArgs {
    /**
     * The IPv4 address to retrieve
     * Only one of `address` and `id` should be specified.
     */
    address?: pulumi.Input<string>;
    /**
     * The ID of the IP address to retrieve
     * Only one of `address` and `id` should be specified.
     */
    id?: pulumi.Input<string>;
}
