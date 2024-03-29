// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Gets information about multiple Load Balancer Backends.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byLBID = scaleway.getLbBackends({
 *     lbId: scaleway_lb.lb01.id,
 * });
 * const byLBIDAndName = scaleway.getLbBackends({
 *     lbId: scaleway_lb.lb01.id,
 *     name: "tf-backend-datasource",
 * });
 * ```
 */
export function getLbBackends(args: GetLbBackendsArgs, opts?: pulumi.InvokeOptions): Promise<GetLbBackendsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getLbBackends:getLbBackends", {
        "lbId": args.lbId,
        "name": args.name,
        "projectId": args.projectId,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbBackends.
 */
export interface GetLbBackendsArgs {
    /**
     * The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
     */
    lbId: string;
    /**
     * The backend name used as filter. Backends with a name like it are listed.
     */
    name?: string;
    projectId?: string;
    /**
     * `zone`) The zone in which backends exist.
     */
    zone?: string;
}

/**
 * A collection of values returned by getLbBackends.
 */
export interface GetLbBackendsResult {
    /**
     * List of found backends
     */
    readonly backends: outputs.GetLbBackendsBackend[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lbId: string;
    readonly name?: string;
    readonly organizationId: string;
    readonly projectId: string;
    readonly zone: string;
}
/**
 * Gets information about multiple Load Balancer Backends.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const byLBID = scaleway.getLbBackends({
 *     lbId: scaleway_lb.lb01.id,
 * });
 * const byLBIDAndName = scaleway.getLbBackends({
 *     lbId: scaleway_lb.lb01.id,
 *     name: "tf-backend-datasource",
 * });
 * ```
 */
export function getLbBackendsOutput(args: GetLbBackendsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLbBackendsResult> {
    return pulumi.output(args).apply((a: any) => getLbBackends(a, opts))
}

/**
 * A collection of arguments for invoking getLbBackends.
 */
export interface GetLbBackendsOutputArgs {
    /**
     * The load-balancer ID this backend is attached to. backends with a LB ID like it are listed.
     */
    lbId: pulumi.Input<string>;
    /**
     * The backend name used as filter. Backends with a name like it are listed.
     */
    name?: pulumi.Input<string>;
    projectId?: pulumi.Input<string>;
    /**
     * `zone`) The zone in which backends exist.
     */
    zone?: pulumi.Input<string>;
}
