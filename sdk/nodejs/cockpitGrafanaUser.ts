// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Cockpit Grafana Users.
 *
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#grafana-users).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 * import * as scaleway from "@pulumi/scaleway";
 *
 * const mainCockpit = scaleway.getCockpit({});
 * // Create an editor grafana user for the cockpit
 * const mainCockpitGrafanaUser = new scaleway.CockpitGrafanaUser("mainCockpitGrafanaUser", {
 *     projectId: mainCockpit.then(mainCockpit => mainCockpit.projectId),
 *     login: "my-awesome-user",
 *     role: "editor",
 * });
 * ```
 *
 * ## Import
 *
 * Cockpits Grafana Users can be imported using the project ID and the grafana user ID formatted `{project_id}/{grafana_user_id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser main 11111111-1111-1111-1111-111111111111/2
 * ```
 */
export class CockpitGrafanaUser extends pulumi.CustomResource {
    /**
     * Get an existing CockpitGrafanaUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CockpitGrafanaUserState, opts?: pulumi.CustomResourceOptions): CockpitGrafanaUser {
        return new CockpitGrafanaUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/cockpitGrafanaUser:CockpitGrafanaUser';

    /**
     * Returns true if the given object is an instance of CockpitGrafanaUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CockpitGrafanaUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CockpitGrafanaUser.__pulumiType;
    }

    /**
     * The login of the grafana user.
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * The password of the grafana user
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * The role of the grafana user. Must be `editor` or `viewer`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a CockpitGrafanaUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CockpitGrafanaUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CockpitGrafanaUserArgs | CockpitGrafanaUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CockpitGrafanaUserState | undefined;
            resourceInputs["login"] = state ? state.login : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as CockpitGrafanaUserArgs | undefined;
            if ((!args || args.login === undefined) && !opts.urn) {
                throw new Error("Missing required property 'login'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["password"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CockpitGrafanaUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CockpitGrafanaUser resources.
 */
export interface CockpitGrafanaUserState {
    /**
     * The login of the grafana user.
     */
    login?: pulumi.Input<string>;
    /**
     * The password of the grafana user
     */
    password?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The role of the grafana user. Must be `editor` or `viewer`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CockpitGrafanaUser resource.
 */
export interface CockpitGrafanaUserArgs {
    /**
     * The login of the grafana user.
     */
    login: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the cockpit is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The role of the grafana user. Must be `editor` or `viewer`.
     */
    role: pulumi.Input<string>;
}
