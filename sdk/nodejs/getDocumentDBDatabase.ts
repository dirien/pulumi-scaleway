// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about DocumentDB database. More on our official [site](https://www.scaleway.com/en/developers/api/document_db/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getDocumentDBDatabase({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 *     name: "foobar",
 * });
 * ```
 */
export function getDocumentDBDatabase(args: GetDocumentDBDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentDBDatabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getDocumentDBDatabase:getDocumentDBDatabase", {
        "instanceId": args.instanceId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getDocumentDBDatabase.
 */
export interface GetDocumentDBDatabaseArgs {
    /**
     * The DocumentDB instance ID.
     */
    instanceId: string;
    /**
     * The name of the DocumentDB instance.
     */
    name?: string;
    region?: string;
}

/**
 * A collection of values returned by getDocumentDBDatabase.
 */
export interface GetDocumentDBDatabaseResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    /**
     * Whether the database is managed or not.
     */
    readonly managed: boolean;
    readonly name?: string;
    /**
     * The name of the owner of the database.
     */
    readonly owner: string;
    readonly projectId: string;
    readonly region?: string;
    /**
     * Size of the database (in bytes).
     */
    readonly size: string;
}
/**
 * Gets information about DocumentDB database. More on our official [site](https://www.scaleway.com/en/developers/api/document_db/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const main = scaleway.getDocumentDBDatabase({
 *     instanceId: "11111111-1111-1111-1111-111111111111",
 *     name: "foobar",
 * });
 * ```
 */
export function getDocumentDBDatabaseOutput(args: GetDocumentDBDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentDBDatabaseResult> {
    return pulumi.output(args).apply((a: any) => getDocumentDBDatabase(a, opts))
}

/**
 * A collection of arguments for invoking getDocumentDBDatabase.
 */
export interface GetDocumentDBDatabaseOutputArgs {
    /**
     * The DocumentDB instance ID.
     */
    instanceId: pulumi.Input<string>;
    /**
     * The name of the DocumentDB instance.
     */
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
