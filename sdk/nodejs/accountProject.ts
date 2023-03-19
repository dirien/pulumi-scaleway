// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages organization's projects on Scaleway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const project = new scaleway.AccountProject("project", {});
 * ```
 *
 * ## Import
 *
 * Projects can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/accountProject:AccountProject project 11111111-1111-1111-1111-111111111111
 * ```
 */
export class AccountProject extends pulumi.CustomResource {
    /**
     * Get an existing AccountProject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountProjectState, opts?: pulumi.CustomResourceOptions): AccountProject {
        return new AccountProject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/accountProject:AccountProject';

    /**
     * Returns true if the given object is an instance of AccountProject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountProject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountProject.__pulumiType;
    }

    /**
     * The Project creation time.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the Project.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the Project.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The Project last update time.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a AccountProject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AccountProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountProjectArgs | AccountProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountProjectState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as AccountProjectArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountProject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountProject resources.
 */
export interface AccountProjectState {
    /**
     * The Project creation time.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the Project.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the Project.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The Project last update time.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AccountProject resource.
 */
export interface AccountProjectArgs {
    /**
     * The description of the Project.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the Project.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`)The organization ID the Project is associated with. Please note that any change in `organizationId` will recreate the resource.
     */
    organizationId?: pulumi.Input<string>;
}
