// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway IAM Applications. For more information, see [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#applications-83ce5e).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.IamApplication("main", {description: "a description"});
 * ```
 *
 * ## Import
 *
 * Applications can be imported using the `{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/iamApplication:IamApplication main 11111111-1111-1111-1111-111111111111
 * ```
 */
export class IamApplication extends pulumi.CustomResource {
    /**
     * Get an existing IamApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamApplicationState, opts?: pulumi.CustomResourceOptions): IamApplication {
        return new IamApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iamApplication:IamApplication';

    /**
     * Returns true if the given object is an instance of IamApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamApplication.__pulumiType;
    }

    /**
     * The date and time of the creation of the application.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the iam application.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the application is editable.
     */
    public /*out*/ readonly editable!: pulumi.Output<boolean>;
    /**
     * The name of the iam application.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `organizationId`) The ID of the organization the application is associated with.
     */
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The tags associated with the application.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time of the last update of the application.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IamApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IamApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamApplicationArgs | IamApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamApplicationState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["editable"] = state ? state.editable : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organizationId"] = state ? state.organizationId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IamApplicationArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["editable"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamApplication resources.
 */
export interface IamApplicationState {
    /**
     * The date and time of the creation of the application.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the iam application.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the application is editable.
     */
    editable?: pulumi.Input<boolean>;
    /**
     * The name of the iam application.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`) The ID of the organization the application is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The tags associated with the application.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time of the last update of the application.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamApplication resource.
 */
export interface IamApplicationArgs {
    /**
     * The description of the iam application.
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the iam application.
     */
    name?: pulumi.Input<string>;
    /**
     * `organizationId`) The ID of the organization the application is associated with.
     */
    organizationId?: pulumi.Input<string>;
    /**
     * The tags associated with the application.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
