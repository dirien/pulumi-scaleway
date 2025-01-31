// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `scaleway.CockpitSource` data source allows you to retrieve information about a specific [data source](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
 *
 * Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
 *
 * ## Example Usage
 *
 * ### Retrieve a specific data source by ID
 *
 * The following example retrieves a Cockpit data source by its unique ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const example = scaleway.getCockpitSource({
 *     id: "fr-par/11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getCockpitSource(args?: GetCockpitSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetCockpitSourceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getCockpitSource:getCockpitSource", {
        "id": args.id,
        "name": args.name,
        "origin": args.origin,
        "projectId": args.projectId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpitSource.
 */
export interface GetCockpitSourceArgs {
    /**
     * The unique identifier of the Cockpit data source in the `{region}/{id}` format. If specified, other filters are ignored.
     */
    id?: string;
    /**
     * The name of the data source.
     */
    name?: string;
    /**
     * The origin of the data source. Possible values are:
     */
    origin?: string;
    /**
     * The ID of the Project the data source is associated with. Defaults to the Project ID specified in the provider configuration.
     */
    projectId?: string;
    /**
     * The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
     */
    type?: string;
}

/**
 * A collection of values returned by getCockpitSource.
 */
export interface GetCockpitSourceResult {
    /**
     * The date and time the data source was created (in RFC 3339 format).
     */
    readonly createdAt: string;
    /**
     * The unique identifier of the data source in the `{region}/{id}` format.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The origin of the data source.
     */
    readonly origin: string;
    readonly projectId: string;
    readonly region: string;
    /**
     * The number of days the data is retained in the data source.
     */
    readonly retentionDays: number;
    /**
     * Indicates whether the data source is synchronized with Grafana.
     */
    readonly synchronizedWithGrafana: boolean;
    readonly type: string;
    /**
     * The date and time the data source was last updated (in RFC 3339 format).
     */
    readonly updatedAt: string;
    /**
     * The URL of the Cockpit data source.
     */
    readonly url: string;
}
/**
 * The `scaleway.CockpitSource` data source allows you to retrieve information about a specific [data source](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
 *
 * Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
 *
 * ## Example Usage
 *
 * ### Retrieve a specific data source by ID
 *
 * The following example retrieves a Cockpit data source by its unique ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const example = scaleway.getCockpitSource({
 *     id: "fr-par/11111111-1111-1111-1111-111111111111",
 * });
 * ```
 */
export function getCockpitSourceOutput(args?: GetCockpitSourceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCockpitSourceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("scaleway:index/getCockpitSource:getCockpitSource", {
        "id": args.id,
        "name": args.name,
        "origin": args.origin,
        "projectId": args.projectId,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getCockpitSource.
 */
export interface GetCockpitSourceOutputArgs {
    /**
     * The unique identifier of the Cockpit data source in the `{region}/{id}` format. If specified, other filters are ignored.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the data source.
     */
    name?: pulumi.Input<string>;
    /**
     * The origin of the data source. Possible values are:
     */
    origin?: pulumi.Input<string>;
    /**
     * The ID of the Project the data source is associated with. Defaults to the Project ID specified in the provider configuration.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
     */
    type?: pulumi.Input<string>;
}
