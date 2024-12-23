// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about an IOT Hub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myHub = scaleway.getIotHub({
 *     hubId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getIotHub(args?: GetIotHubArgs, opts?: pulumi.InvokeOptions): Promise<GetIotHubResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getIotHub:getIotHub", {
        "hubId": args.hubId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getIotHub.
 */
export interface GetIotHubArgs {
    /**
     * The Hub ID.
     * Only one of the `name` and `hubId` should be specified.
     */
    hubId?: string;
    /**
     * The name of the Hub.
     * Only one of the `name` and `hubId` should be specified.
     */
    name?: string;
    /**
     * The ID of the project the hub is associated with.
     */
    projectId?: string;
    /**
     * `region`) The region in which the hub exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getIotHub.
 */
export interface GetIotHubResult {
    readonly connectedDeviceCount: number;
    readonly createdAt: string;
    readonly deviceAutoProvisioning: boolean;
    readonly deviceCount: number;
    readonly disableEvents: boolean;
    readonly enabled: boolean;
    readonly endpoint: string;
    readonly eventsTopicPrefix: string;
    readonly hubCa: string;
    readonly hubCaChallenge: string;
    readonly hubId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mqttCa: string;
    readonly mqttCaUrl: string;
    readonly name?: string;
    readonly organizationId: string;
    readonly productPlan: string;
    readonly projectId?: string;
    readonly region?: string;
    readonly status: string;
    readonly updatedAt: string;
}
/**
 * Gets information about an IOT Hub.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myHub = scaleway.getIotHub({
 *     hubId: "11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getIotHubOutput(args?: GetIotHubOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIotHubResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getIotHub:getIotHub", {
        "hubId": args.hubId,
        "name": args.name,
        "projectId": args.projectId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getIotHub.
 */
export interface GetIotHubOutputArgs {
    /**
     * The Hub ID.
     * Only one of the `name` and `hubId` should be specified.
     */
    hubId?: pulumi.Input<string>;
    /**
     * The name of the Hub.
     * Only one of the `name` and `hubId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project the hub is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`) The region in which the hub exists.
     */
    region?: pulumi.Input<string>;
}
