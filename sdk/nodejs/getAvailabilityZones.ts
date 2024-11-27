// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `scaleway.getAvailabilityZones` data source is used to retrieve information about the available zones based on its Region.
 *
 * For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
 * you can choose the location that better fits your need (country, latency, etc.).
 *
 * Refer to the Account [documentation](https://www.scaleway.com/en/docs/console/account/reference-content/products-availability/) for more information.
 *
 * ## Retrieve the Availability Zones of a Region
 *
 * The following command allow you to retrieve a the AZs of a Region.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getAvailabilityZones({
 *     region: "nl-ams",
 * });
 * ```
 */
export function getAvailabilityZones(args?: GetAvailabilityZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailabilityZonesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getAvailabilityZones:getAvailabilityZones", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailabilityZones.
 */
export interface GetAvailabilityZonesArgs {
    /**
     * Region is represented as a Geographical area, such as France. Defaults to `fr-par`.
     */
    region?: string;
}

/**
 * A collection of values returned by getAvailabilityZones.
 */
export interface GetAvailabilityZonesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region?: string;
    /**
     * The list of availability zones in each Region
     */
    readonly zones: string[];
}
/**
 * The `scaleway.getAvailabilityZones` data source is used to retrieve information about the available zones based on its Region.
 *
 * For technical and legal reasons, some products are split by Region or by Availability Zones. When using such product,
 * you can choose the location that better fits your need (country, latency, etc.).
 *
 * Refer to the Account [documentation](https://www.scaleway.com/en/docs/console/account/reference-content/products-availability/) for more information.
 *
 * ## Retrieve the Availability Zones of a Region
 *
 * The following command allow you to retrieve a the AZs of a Region.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getAvailabilityZones({
 *     region: "nl-ams",
 * });
 * ```
 */
export function getAvailabilityZonesOutput(args?: GetAvailabilityZonesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAvailabilityZonesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getAvailabilityZones:getAvailabilityZones", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailabilityZones.
 */
export interface GetAvailabilityZonesOutputArgs {
    /**
     * Region is represented as a Geographical area, such as France. Defaults to `fr-par`.
     */
    region?: pulumi.Input<string>;
}
