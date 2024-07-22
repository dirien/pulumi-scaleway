// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Get information about Scaleway Load Balancer routes.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const ip01 = new scaleway.LbIp("ip01", {});
 * const lb01 = new scaleway.Lb("lb01", {
 *     ipId: ip01.id,
 *     type: "lb-s",
 * });
 * const bkd01 = new scaleway.LbBackend("bkd01", {
 *     lbId: lb01.id,
 *     forwardProtocol: "tcp",
 *     forwardPort: 80,
 *     proxyProtocol: "none",
 * });
 * const frt01 = new scaleway.LbFrontend("frt01", {
 *     lbId: lb01.id,
 *     backendId: bkd01.id,
 *     inboundPort: 80,
 * });
 * const rt01 = new scaleway.LbRoute("rt01", {
 *     frontendId: frt01.id,
 *     backendId: bkd01.id,
 *     matchSni: "sni.scaleway.com",
 * });
 * const byID = scaleway.getLbRouteOutput({
 *     routeId: rt01.id,
 * });
 * ```
 */
export function getLbRoute(args: GetLbRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetLbRouteResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLbRoute:getLbRoute", {
        "routeId": args.routeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbRoute.
 */
export interface GetLbRouteArgs {
    /**
     * The route ID.
     */
    routeId: string;
}

/**
 * A collection of values returned by getLbRoute.
 */
export interface GetLbRouteResult {
    readonly backendId: string;
    readonly createdAt: string;
    readonly frontendId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly matchHostHeader: string;
    readonly matchSni: string;
    readonly routeId: string;
    readonly updatedAt: string;
}
/**
 * Get information about Scaleway Load Balancer routes.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/network/load-balancer/how-to/create-manage-routes/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-route).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const ip01 = new scaleway.LbIp("ip01", {});
 * const lb01 = new scaleway.Lb("lb01", {
 *     ipId: ip01.id,
 *     type: "lb-s",
 * });
 * const bkd01 = new scaleway.LbBackend("bkd01", {
 *     lbId: lb01.id,
 *     forwardProtocol: "tcp",
 *     forwardPort: 80,
 *     proxyProtocol: "none",
 * });
 * const frt01 = new scaleway.LbFrontend("frt01", {
 *     lbId: lb01.id,
 *     backendId: bkd01.id,
 *     inboundPort: 80,
 * });
 * const rt01 = new scaleway.LbRoute("rt01", {
 *     frontendId: frt01.id,
 *     backendId: bkd01.id,
 *     matchSni: "sni.scaleway.com",
 * });
 * const byID = scaleway.getLbRouteOutput({
 *     routeId: rt01.id,
 * });
 * ```
 */
export function getLbRouteOutput(args: GetLbRouteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbRouteResult> {
    return pulumi.output(args).apply((a: any) => getLbRoute(a, opts))
}

/**
 * A collection of arguments for invoking getLbRoute.
 */
export interface GetLbRouteOutputArgs {
    /**
     * The route ID.
     */
    routeId: pulumi.Input<string>;
}
