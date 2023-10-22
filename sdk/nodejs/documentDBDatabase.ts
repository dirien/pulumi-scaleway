// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class DocumentDBDatabase extends pulumi.CustomResource {
    /**
     * Get an existing DocumentDBDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentDBDatabaseState, opts?: pulumi.CustomResourceOptions): DocumentDBDatabase {
        return new DocumentDBDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/documentDBDatabase:DocumentDBDatabase';

    /**
     * Returns true if the given object is an instance of DocumentDBDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentDBDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentDBDatabase.__pulumiType;
    }

    /**
     * Instance on which the database is created
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Whether or not the database is managed
     */
    public /*out*/ readonly managed!: pulumi.Output<boolean>;
    /**
     * The database name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * User that own the database
     */
    public /*out*/ readonly owner!: pulumi.Output<string>;
    /**
     * The project_id you want to attach the resource to
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Size of the database
     */
    public /*out*/ readonly size!: pulumi.Output<string>;

    /**
     * Create a DocumentDBDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentDBDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentDBDatabaseArgs | DocumentDBDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentDBDatabaseState | undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["managed"] = state ? state.managed : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
        } else {
            const args = argsOrState as DocumentDBDatabaseArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["managed"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DocumentDBDatabase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentDBDatabase resources.
 */
export interface DocumentDBDatabaseState {
    /**
     * Instance on which the database is created
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Whether or not the database is managed
     */
    managed?: pulumi.Input<boolean>;
    /**
     * The database name
     */
    name?: pulumi.Input<string>;
    /**
     * User that own the database
     */
    owner?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * Size of the database
     */
    size?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentDBDatabase resource.
 */
export interface DocumentDBDatabaseArgs {
    /**
     * Instance on which the database is created
     */
    instanceId: pulumi.Input<string>;
    /**
     * The database name
     */
    name?: pulumi.Input<string>;
    /**
     * The project_id you want to attach the resource to
     */
    projectId?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
}
