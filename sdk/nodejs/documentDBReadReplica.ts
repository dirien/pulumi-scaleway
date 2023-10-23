// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class DocumentDBReadReplica extends pulumi.CustomResource {
    /**
     * Get an existing DocumentDBReadReplica resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DocumentDBReadReplicaState, opts?: pulumi.CustomResourceOptions): DocumentDBReadReplica {
        return new DocumentDBReadReplica(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/documentDBReadReplica:DocumentDBReadReplica';

    /**
     * Returns true if the given object is an instance of DocumentDBReadReplica.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DocumentDBReadReplica {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DocumentDBReadReplica.__pulumiType;
    }

    /**
     * Direct access endpoint, it gives you an IP and a port to access your read-replica
     */
    public readonly directAccess!: pulumi.Output<outputs.DocumentDBReadReplicaDirectAccess | undefined>;
    /**
     * Id of the rdb instance to replicate
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Private network endpoints
     */
    public readonly privateNetwork!: pulumi.Output<outputs.DocumentDBReadReplicaPrivateNetwork | undefined>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a DocumentDBReadReplica resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DocumentDBReadReplicaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DocumentDBReadReplicaArgs | DocumentDBReadReplicaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DocumentDBReadReplicaState | undefined;
            resourceInputs["directAccess"] = state ? state.directAccess : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["privateNetwork"] = state ? state.privateNetwork : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as DocumentDBReadReplicaArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["directAccess"] = args ? args.directAccess : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["privateNetwork"] = args ? args.privateNetwork : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DocumentDBReadReplica.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DocumentDBReadReplica resources.
 */
export interface DocumentDBReadReplicaState {
    /**
     * Direct access endpoint, it gives you an IP and a port to access your read-replica
     */
    directAccess?: pulumi.Input<inputs.DocumentDBReadReplicaDirectAccess>;
    /**
     * Id of the rdb instance to replicate
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Private network endpoints
     */
    privateNetwork?: pulumi.Input<inputs.DocumentDBReadReplicaPrivateNetwork>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DocumentDBReadReplica resource.
 */
export interface DocumentDBReadReplicaArgs {
    /**
     * Direct access endpoint, it gives you an IP and a port to access your read-replica
     */
    directAccess?: pulumi.Input<inputs.DocumentDBReadReplicaDirectAccess>;
    /**
     * Id of the rdb instance to replicate
     */
    instanceId: pulumi.Input<string>;
    /**
     * Private network endpoints
     */
    privateNetwork?: pulumi.Input<inputs.DocumentDBReadReplicaPrivateNetwork>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
}