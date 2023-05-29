// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Function Triggers.
 * For more information see [the documentation](https://www.scaleway.com/en/developers/api/serverless-functions/#path-triggers).
 *
 * ## Examples
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.FunctionTrigger("main", {
 *     functionId: scaleway_function.main.id,
 *     sqs: {
 *         namespaceId: scaleway_mnq_namespace.main.id,
 *         queue: "MyQueue",
 *         projectId: scaleway_mnq_namespace.main.project_id,
 *         region: scaleway_mnq_namespace.main.region,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Function Triggers can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/functionTrigger:FunctionTrigger main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class FunctionTrigger extends pulumi.CustomResource {
    /**
     * Get an existing FunctionTrigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionTriggerState, opts?: pulumi.CustomResourceOptions): FunctionTrigger {
        return new FunctionTrigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/functionTrigger:FunctionTrigger';

    /**
     * Returns true if the given object is an instance of FunctionTrigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionTrigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionTrigger.__pulumiType;
    }

    /**
     * The description of the trigger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the function to create a trigger for
     */
    public readonly functionId!: pulumi.Output<string>;
    /**
     * The unique name of the trigger. Default to a generated name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The configuration of the Scaleway's SQS used by the trigger
     */
    public readonly sqs!: pulumi.Output<outputs.FunctionTriggerSqs | undefined>;

    /**
     * Create a FunctionTrigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionTriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionTriggerArgs | FunctionTriggerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionTriggerState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["functionId"] = state ? state.functionId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["sqs"] = state ? state.sqs : undefined;
        } else {
            const args = argsOrState as FunctionTriggerArgs | undefined;
            if ((!args || args.functionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["functionId"] = args ? args.functionId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["sqs"] = args ? args.sqs : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionTrigger.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionTrigger resources.
 */
export interface FunctionTriggerState {
    /**
     * The description of the trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the function to create a trigger for
     */
    functionId?: pulumi.Input<string>;
    /**
     * The unique name of the trigger. Default to a generated name.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The configuration of the Scaleway's SQS used by the trigger
     */
    sqs?: pulumi.Input<inputs.FunctionTriggerSqs>;
}

/**
 * The set of arguments for constructing a FunctionTrigger resource.
 */
export interface FunctionTriggerArgs {
    /**
     * The description of the trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the function to create a trigger for
     */
    functionId: pulumi.Input<string>;
    /**
     * The unique name of the trigger. Default to a generated name.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`). The region in which the namespace should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The configuration of the Scaleway's SQS used by the trigger
     */
    sqs?: pulumi.Input<inputs.FunctionTriggerSqs>;
}
