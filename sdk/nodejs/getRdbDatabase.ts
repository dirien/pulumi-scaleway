// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a RDB database.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myDb = scaleway.getRdbDatabase({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 *     name: "foobar",
 * });
 * ```
 */
export function getRdbDatabase(args: GetRdbDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetRdbDatabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getRdbDatabase:getRdbDatabase", {
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdbDatabase.
 */
export interface GetRdbDatabaseArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: string;
    /**
     * The name of the RDB instance.
     */
    name: string;
    region?: string;
}

/**
 * A collection of values returned by getRdbDatabase.
 */
export interface GetRdbDatabaseResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Whether the database is managed or not.
     */
    readonly managed: boolean;
    readonly name: string;
    /**
     * The name of the owner of the database.
     */
    readonly owner: string;
    readonly region?: string;
    /**
     * Size of the database (in bytes).
     */
    readonly size: string;
}
/**
 * Gets information about a RDB database.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const myDb = scaleway.getRdbDatabase({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 *     name: "foobar",
 * });
 * ```
 */
export function getRdbDatabaseOutput(args: GetRdbDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRdbDatabaseResult> {
    return pulumi.output(args).apply((a: any) => getRdbDatabase(a, opts))
}

/**
 * A collection of arguments for invoking getRdbDatabase.
 */
export interface GetRdbDatabaseOutputArgs {
    /**
     * The RDB instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of the RDB instance.
     */
    name: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
