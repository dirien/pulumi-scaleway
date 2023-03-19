// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway IAM Groups.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#groups-f592eb).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const basic = new scaleway.IamGroup("basic", {
 *     applicationIds: [],
 *     description: "basic description",
 *     userIds: [],
 * });
 * ```
 *
 * ### With applications
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const app = new scaleway.IamApplication("app", {});
 * const withApp = new scaleway.IamGroup("withApp", {
 *     applicationIds: [app.id],
 *     userIds: [],
 * });
 * ```
 *
 * ## Import
 *
 * IAM groups can be imported using the `{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/iamGroup:IamGroup basic 11111111-1111-1111-1111-111111111111
 * ```
 */
export class IamGroup extends pulumi.CustomResource {
    /**
     * Get an existing IamGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamGroupState, opts?: pulumi.CustomResourceOptions): IamGroup {
        return new IamGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iamGroup:IamGroup';

    /**
     * Returns true if the given object is an instance of IamGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamGroup.__pulumiType;
    }

    /**
     * The list of IDs of the applications attached to the group.
     */
    public readonly applicationIds!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time of the creation of the group
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the IAM group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the IAM group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `organizationId`) The ID of the organization the group is associated with.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The date and time of the last update of the group
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The list of IDs of the users attached to the group.
     */
    public readonly userIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a IamGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IamGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamGroupArgs | IamGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamGroupState | undefined;
            resourceInputs["applicationIds"] = state ? state.applicationIds : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userIds"] = state ? state.userIds : undefined;
        } else {
            const args = argsOrState as IamGroupArgs | undefined;
            resourceInputs["applicationIds"] = args ? args.applicationIds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["userIds"] = args ? args.userIds : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamGroup resources.
 */
export interface IamGroupState {
    /**
     * The list of IDs of the applications attached to the group.
     */
    applicationIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time of the creation of the group
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the IAM group.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the IAM group.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`) The ID of the organization the group is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The date and time of the last update of the group
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The list of IDs of the users attached to the group.
     */
    userIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a IamGroup resource.
 */
export interface IamGroupArgs {
    /**
     * The list of IDs of the applications attached to the group.
     */
    applicationIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the IAM group.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the IAM group.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`) The ID of the organization the group is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The list of IDs of the users attached to the group.
     */
    userIds?: pulumi.Input<pulumi.Input<string>[]>;
}
