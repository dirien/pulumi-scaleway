// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `scaleway.CockpitSource` resource allows you to create and manage [data sources](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-sources) in Scaleway's Cockpit.
 *
 * Refer to Cockpit's [product documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/) and [API documentation](https://www.scaleway.com/en/developers/api/cockpit/regional-api) for more information.
 *
 * ## Example Usage
 *
 * ### Create a data source
 *
 * The following command allows you to create a [metrics](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#metric) data source named `my-data-source` in a given Project.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const project = new scaleway.AccountProject("project", {});
 * const main = new scaleway.CockpitSource("main", {
 *     projectId: project.id,
 *     type: "metrics",
 *     retentionDays: 6,
 * });
 * ```
 *
 * ## Import
 *
 * This section explains how to import a data source using the ID of the region it is located in, in the `{region}/{id}` format.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/cockpitSource:CockpitSource main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class CockpitSource extends pulumi.CustomResource {
    /**
     * Get an existing CockpitSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CockpitSourceState, opts?: pulumi.CustomResourceOptions): CockpitSource {
        return new CockpitSource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/cockpitSource:CockpitSource';

    /**
     * Returns true if the given object is an instance of CockpitSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CockpitSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CockpitSource.__pulumiType;
    }

    /**
     * The date and time the data source was created (in RFC 3339 format).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The name of the data source.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The origin of the Cockpit data source.
     */
    public /*out*/ readonly origin!: pulumi.Output<string>;
    /**
     * ) The ID of the Project the data source is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The URL endpoint used for pushing data to the Cockpit data source.
     */
    public /*out*/ readonly pushUrl!: pulumi.Output<string>;
    /**
     * ) The region where the data source is located.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The number of days to retain data in the data source. Must be a value between 1 and 365. Changes to this field will force the creation of a new resource.
     */
    public readonly retentionDays!: pulumi.Output<number>;
    /**
     * Indicates whether the data source is synchronized with Grafana.
     */
    public /*out*/ readonly synchronizedWithGrafana!: pulumi.Output<boolean>;
    /**
     * The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
     */
    public readonly type!: pulumi.Output<string | undefined>;
    /**
     * The date and time the data source was last updated (in RFC 3339 format).
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The URL of the Cockpit data source.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a CockpitSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CockpitSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CockpitSourceArgs | CockpitSourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CockpitSourceState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["origin"] = state ? state.origin : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["pushUrl"] = state ? state.pushUrl : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["retentionDays"] = state ? state.retentionDays : undefined;
            resourceInputs["synchronizedWithGrafana"] = state ? state.synchronizedWithGrafana : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as CockpitSourceArgs | undefined;
            if ((!args || args.retentionDays === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionDays'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["retentionDays"] = args ? args.retentionDays : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["origin"] = undefined /*out*/;
            resourceInputs["pushUrl"] = undefined /*out*/;
            resourceInputs["synchronizedWithGrafana"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CockpitSource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CockpitSource resources.
 */
export interface CockpitSourceState {
    /**
     * The date and time the data source was created (in RFC 3339 format).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The name of the data source.
     */
    name?: pulumi.Input<string>;
    /**
     * The origin of the Cockpit data source.
     */
    origin?: pulumi.Input<string>;
    /**
     * ) The ID of the Project the data source is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The URL endpoint used for pushing data to the Cockpit data source.
     */
    pushUrl?: pulumi.Input<string>;
    /**
     * ) The region where the data source is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The number of days to retain data in the data source. Must be a value between 1 and 365. Changes to this field will force the creation of a new resource.
     */
    retentionDays?: pulumi.Input<number>;
    /**
     * Indicates whether the data source is synchronized with Grafana.
     */
    synchronizedWithGrafana?: pulumi.Input<boolean>;
    /**
     * The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
     */
    type?: pulumi.Input<string>;
    /**
     * The date and time the data source was last updated (in RFC 3339 format).
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The URL of the Cockpit data source.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CockpitSource resource.
 */
export interface CockpitSourceArgs {
    /**
     * The name of the data source.
     */
    name?: pulumi.Input<string>;
    /**
     * ) The ID of the Project the data source is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * ) The region where the data source is located.
     */
    region?: pulumi.Input<string>;
    /**
     * The number of days to retain data in the data source. Must be a value between 1 and 365. Changes to this field will force the creation of a new resource.
     */
    retentionDays: pulumi.Input<number>;
    /**
     * The [type](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#data-types) of data source. Possible values are: `metrics`, `logs`, or `traces`.
     */
    type?: pulumi.Input<string>;
}
