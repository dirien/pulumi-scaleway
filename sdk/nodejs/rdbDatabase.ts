// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway RDB database.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.RdbDatabase("main", {instanceId: scaleway_rdb_instance.main.id});
 * ```
 *
 * ## Import
 *
 * RDB Database can be imported using the `{region}/{id}/{DBNAME}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/rdbDatabase:RdbDatabase rdb01_mydb fr-par/11111111-1111-1111-1111-111111111111/mydb
 * ```
 */
export class RdbDatabase extends pulumi.CustomResource {
    /**
     * Get an existing RdbDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdbDatabaseState, opts?: pulumi.CustomResourceOptions): RdbDatabase {
        return new RdbDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/rdbDatabase:RdbDatabase';

    /**
     * Returns true if the given object is an instance of RdbDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdbDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdbDatabase.__pulumiType;
    }

    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether the database is managed or not.
     */
    public /*out*/ readonly managed!: pulumi.Output<boolean>;
    /**
     * Name of the database (e.g. `my-new-database`).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the owner of the database.
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Size of the database (in bytes).
     */
    public /*out*/ readonly size!: pulumi.Output<string>;

    /**
     * Create a RdbDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdbDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdbDatabaseArgs | RdbDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdbDatabaseState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
        } else {
            const args = argsOrState as RdbDatabaseArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["managed"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdbDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdbDatabase resources.
 */
export interface RdbDatabaseState {
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether the database is managed or not.
     */
    managed?: pulumi.Input<boolean>;
    /**
     * Name of the database (e.g. `my-new-database`).
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the owner of the database.
     */
    owner?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
    /**
     * Size of the database (in bytes).
     */
    size?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdbDatabase resource.
 */
export interface RdbDatabaseArgs {
    /**
     * UUID of the rdb instance.
     *
     * > **Important:** Updates to `instanceId` will recreate the Database.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Name of the database (e.g. `my-new-database`).
     */
    name?: pulumi.Input<string>;
    /**
     * `region`) The region in which the resource exists.
     */
    region?: pulumi.Input<string>;
}
