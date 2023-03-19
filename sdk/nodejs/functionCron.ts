// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Function Triggers. For the moment, the feature is limited to CRON Schedule (time-based).
 *
 * For more information consult
 * the [documentation](https://www.scaleway.com/en/docs/compute/functions/api-cli/fun-uploading-with-serverless-framework/#configuring-events)
 * .
 *
 * For more details about the limitation
 * check [functions-limitations](https://www.scaleway.com/en/docs/compute/functions/reference-content/functions-limitations/).
 *
 * You can check also
 * our [functions cron api documentation](https://developers.scaleway.com/en/products/functions/api/#crons-942bf4).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const mainFunctionNamespace = new scaleway.FunctionNamespace("mainFunctionNamespace", {});
 * const mainFunction = new scaleway.Function("mainFunction", {
 *     namespaceId: mainFunctionNamespace.id,
 *     runtime: "node14",
 *     privacy: "private",
 *     handler: "handler.handle",
 * });
 * const mainFunctionCron = new scaleway.FunctionCron("mainFunctionCron", {
 *     functionId: mainFunction.id,
 *     schedule: "0 0 * * *",
 *     args: JSON.stringify({
 *         test: "scw",
 *     }),
 * });
 * const func = new scaleway.FunctionCron("func", {
 *     functionId: mainFunction.id,
 *     schedule: "0 1 * * *",
 *     args: JSON.stringify({
 *         my_var: "terraform",
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * Container Cron can be imported using the `{region}/{id}`, e.g. bash
 *
 * ```sh
 *  $ pulumi import scaleway:index/functionCron:FunctionCron main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class FunctionCron extends pulumi.CustomResource {
    /**
     * Get an existing FunctionCron resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionCronState, opts?: pulumi.CustomResourceOptions): FunctionCron {
        return new FunctionCron(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/functionCron:FunctionCron';

    /**
     * Returns true if the given object is an instance of FunctionCron.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionCron {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionCron.__pulumiType;
    }

    /**
     * The key-value mapping to define arguments that will be passed to your function’s event object
     * during
     */
    public readonly args!: pulumi.Output<string>;
    /**
     * The function ID to link with your cron.
     */
    public readonly functionId!: pulumi.Output<string>;
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     */
    public readonly schedule!: pulumi.Output<string>;
    /**
     * The cron status.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a FunctionCron resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionCronArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionCronArgs | FunctionCronState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionCronState | undefined;
            resourceInputs["args"] = state ? state.args : undefined;
            resourceInputs["functionId"] = state ? state.functionId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["schedule"] = state ? state.schedule : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as FunctionCronArgs | undefined;
            if ((!args || args.args === undefined) && !opts.urn) {
                throw new Error("Missing required property 'args'");
            }
            if ((!args || args.functionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionId'");
            }
            if ((!args || args.schedule === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schedule'");
            }
            resourceInputs["args"] = args ? args.args : undefined;
            resourceInputs["functionId"] = args ? args.functionId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["schedule"] = args ? args.schedule : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FunctionCron.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionCron resources.
 */
export interface FunctionCronState {
    /**
     * The key-value mapping to define arguments that will be passed to your function’s event object
     * during
     */
    args?: pulumi.Input<string>;
    /**
     * The function ID to link with your cron.
     */
    functionId?: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     */
    region?: pulumi.Input<string>;
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     */
    schedule?: pulumi.Input<string>;
    /**
     * The cron status.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionCron resource.
 */
export interface FunctionCronArgs {
    /**
     * The key-value mapping to define arguments that will be passed to your function’s event object
     * during
     */
    args: pulumi.Input<string>;
    /**
     * The function ID to link with your cron.
     */
    functionId: pulumi.Input<string>;
    /**
     * (Defaults to provider `region`) The region
     * in where the job was created.
     */
    region?: pulumi.Input<string>;
    /**
     * Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
     * executed.
     */
    schedule: pulumi.Input<string>;
}
