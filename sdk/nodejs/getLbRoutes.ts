// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about multiple Load Balancer routes.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const byFrontendID = scaleway.getLbRoutes({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 * });
 * const myKey = scaleway.getLbRoutes({
 *     frontendId: "11111111-1111-1111-1111-111111111111",
 *     zone: "fr-par-2",
 * });
 * ```
 */
export function getLbRoutes(args?: GetLbRoutesArgs, opts?: pulumi.InvokeOptions): Promise<GetLbRoutesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLbRoutes:getLbRoutes", {
        "frontendId": args.frontendId,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbRoutes.
 */
export interface GetLbRoutesArgs {
    /**
     * The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
     */
    frontendId?: string;
    projectId?: string;
    /**
     * `zone`) The zone in which the routes exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getLbRoutes.
 */
export interface GetLbRoutesResult {
    readonly frontendId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly organizationId: string;
    readonly projectId: string;
    /**
     * List of retrieved routes
     */
    readonly routes: outputs.GetLbRoutesRoute[];
    readonly zone: string;
}
/**
 * Gets information about multiple Load Balancer routes.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const byFrontendID = scaleway.getLbRoutes({
 *     frontendId: scaleway_lb_frontend.frt01.id,
 * });
 * const myKey = scaleway.getLbRoutes({
 *     frontendId: "11111111-1111-1111-1111-111111111111",
 *     zone: "fr-par-2",
 * });
 * ```
 */
export function getLbRoutesOutput(args?: GetLbRoutesOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLbRoutesResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getLbRoutes:getLbRoutes", {
        "frontendId": args.frontendId,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbRoutes.
 */
export interface GetLbRoutesOutputArgs {
    /**
     * The frontend ID (the origin of the redirection), to filter for. Routes with a matching frontend ID are listed.
     */
    frontendId?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which the routes exist.
     */
    zone?: pulumi.Input<string>;
}
