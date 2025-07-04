// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Get information about Scaleway Load Balancer frontends.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
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
 * const byID = scaleway.getLbFrontendOutput({
 *     frontendId: frt01.id,
 * });
 * const byName = scaleway.getLbFrontendOutput({
 *     name: frt01.name,
 *     lbId: lb01.id,
 * });
 * ```
 */
export function getLbFrontend(args?: GetLbFrontendArgs, opts?: pulumi.InvokeOptions): Promise<GetLbFrontendResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLbFrontend:getLbFrontend", {
        "frontendId": args.frontendId,
        "lbId": args.lbId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbFrontend.
 */
export interface GetLbFrontendArgs {
    /**
     * The frontend ID.
     * - Only one of `name` and `frontendId` should be specified.
     */
    frontendId?: string;
    /**
     * The Load Balancer ID this frontend is attached to.
     */
    lbId?: string;
    /**
     * The name of the frontend.
     * - When using the `name` you should specify the `lb-id`
     */
    name?: string;
}

/**
 * A collection of values returned by getLbFrontend.
 */
export interface GetLbFrontendResult {
    readonly acls: outputs.GetLbFrontendAcl[];
    readonly backendId: string;
    readonly certificateId: string;
    readonly certificateIds: string[];
    readonly connectionRateLimit: number;
    readonly enableHttp3: boolean;
    readonly externalAcls: boolean;
    readonly frontendId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly inboundPort: number;
    readonly lbId?: string;
    readonly name?: string;
    readonly timeoutClient: string;
}
/**
 * Get information about Scaleway Load Balancer frontends.
 *
 * For more information, see the [main documentation](https://www.scaleway.com/en/docs/load-balancer/reference-content/configuring-frontends/) or [API documentation](https://www.scaleway.com/en/developers/api/load-balancer/zoned-api/#path-frontends).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
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
 * const byID = scaleway.getLbFrontendOutput({
 *     frontendId: frt01.id,
 * });
 * const byName = scaleway.getLbFrontendOutput({
 *     name: frt01.name,
 *     lbId: lb01.id,
 * });
 * ```
 */
export function getLbFrontendOutput(args?: GetLbFrontendOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLbFrontendResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getLbFrontend:getLbFrontend", {
        "frontendId": args.frontendId,
        "lbId": args.lbId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbFrontend.
 */
export interface GetLbFrontendOutputArgs {
    /**
     * The frontend ID.
     * - Only one of `name` and `frontendId` should be specified.
     */
    frontendId?: pulumi.Input<string>;
    /**
     * The Load Balancer ID this frontend is attached to.
     */
    lbId?: pulumi.Input<string>;
    /**
     * The name of the frontend.
     * - When using the `name` you should specify the `lb-id`
     */
    name?: pulumi.Input<string>;
}
