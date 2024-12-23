// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about a webhosting offer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byName = scaleway.getWebhostingOffer({
 *     name: "performance",
 * });
 * const byId = scaleway.getWebhostingOffer({
 *     offerId: "de2426b4-a9e9-11ec-b909-0242ac120002",
 * });
 * ```
 */
export function getWebhostingOffer(args?: GetWebhostingOfferArgs, opts?: pulumi.InvokeOptions): Promise<GetWebhostingOfferResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getWebhostingOffer:getWebhostingOffer", {
        "name": args.name,
        "offerId": args.offerId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebhostingOffer.
 */
export interface GetWebhostingOfferArgs {
    /**
     * The offer name. Only one of `name` and `offerId` should be specified.
     */
    name?: string;
    /**
     * The offer id. Only one of `name` and `offerId` should be specified.
     */
    offerId?: string;
    /**
     * `region`) The region in which offer exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getWebhostingOffer.
 */
export interface GetWebhostingOfferResult {
    /**
     * The unique identifier used for billing.
     */
    readonly billingOperationPath: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name?: string;
    readonly offerId?: string;
    /**
     * The offer price.
     */
    readonly price: string;
    /**
     * The offer product.
     */
    readonly products: outputs.GetWebhostingOfferProduct[];
    readonly region: string;
}
/**
 * Gets information about a webhosting offer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byName = scaleway.getWebhostingOffer({
 *     name: "performance",
 * });
 * const byId = scaleway.getWebhostingOffer({
 *     offerId: "de2426b4-a9e9-11ec-b909-0242ac120002",
 * });
 * ```
 */
export function getWebhostingOfferOutput(args?: GetWebhostingOfferOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetWebhostingOfferResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getWebhostingOffer:getWebhostingOffer", {
        "name": args.name,
        "offerId": args.offerId,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebhostingOffer.
 */
export interface GetWebhostingOfferOutputArgs {
    /**
     * The offer name. Only one of `name` and `offerId` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * The offer id. Only one of `name` and `offerId` should be specified.
     */
    offerId?: pulumi.Input<string>;
    /**
     * `region`) The region in which offer exists.
     */
    region?: pulumi.Input<string>;
}
