// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Messaging and Queuing SQS credentials.
 * For further information, see
 * our [main documentation](https://www.scaleway.com/en/docs/serverless/messaging/reference-content/sqs-overview/)
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const mainMnqSqs = new scaleway.MnqSqs("mainMnqSqs", {});
 * const mainMnqSqsCredentials = new scaleway.MnqSqsCredentials("mainMnqSqsCredentials", {
 *     projectId: mainMnqSqs.projectId,
 *     permissions: {
 *         canManage: false,
 *         canReceive: true,
 *         canPublish: false,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * SQS credentials can be imported using `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/mnqSqsCredentials:MnqSqsCredentials main fr-par/11111111111111111111111111111111
 * ```
 */
export class MnqSqsCredentials extends pulumi.CustomResource {
    /**
     * Get an existing MnqSqsCredentials resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MnqSqsCredentialsState, opts?: pulumi.CustomResourceOptions): MnqSqsCredentials {
        return new MnqSqsCredentials(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/mnqSqsCredentials:MnqSqsCredentials';

    /**
     * Returns true if the given object is an instance of MnqSqsCredentials.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MnqSqsCredentials {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MnqSqsCredentials.__pulumiType;
    }

    /**
     * The ID of the key.
     */
    public /*out*/ readonly accessKey!: pulumi.Output<string>;
    /**
     * The unique name of the SQS credentials.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * . List of permissions associated with these credentials. Only one of the following permissions may be set:
     */
    public readonly permissions!: pulumi.Output<outputs.MnqSqsCredentialsPermissions>;
    /**
     * `projectId`) The ID of the Project in which SQS is enabled.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which SQS is enabled.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The secret value of the key.
     */
    public /*out*/ readonly secretKey!: pulumi.Output<string>;

    /**
     * Create a MnqSqsCredentials resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MnqSqsCredentialsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MnqSqsCredentialsArgs | MnqSqsCredentialsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MnqSqsCredentialsState | undefined;
            resourceInputs["accessKey"] = state ? state.accessKey : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["secretKey"] = state ? state.secretKey : undefined;
        } else {
            const args = argsOrState as MnqSqsCredentialsArgs | undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["accessKey"] = undefined /*out*/;
            resourceInputs["secretKey"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["accessKey", "secretKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(MnqSqsCredentials.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MnqSqsCredentials resources.
 */
export interface MnqSqsCredentialsState {
    /**
     * The ID of the key.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The unique name of the SQS credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated with these credentials. Only one of the following permissions may be set:
     */
    permissions?: pulumi.Input<inputs.MnqSqsCredentialsPermissions>;
    /**
     * `projectId`) The ID of the Project in which SQS is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which SQS is enabled.
     */
    region?: pulumi.Input<string>;
    /**
     * The secret value of the key.
     */
    secretKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MnqSqsCredentials resource.
 */
export interface MnqSqsCredentialsArgs {
    /**
     * The unique name of the SQS credentials.
     */
    name?: pulumi.Input<string>;
    /**
     * . List of permissions associated with these credentials. Only one of the following permissions may be set:
     */
    permissions?: pulumi.Input<inputs.MnqSqsCredentialsPermissions>;
    /**
     * `projectId`) The ID of the Project in which SQS is enabled.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which SQS is enabled.
     */
    region?: pulumi.Input<string>;
}
