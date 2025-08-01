// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const mainIotHub = new scaleway.IotHub("mainIotHub", {productPlan: "plan_shared"});
 * const mainIotDevice = new scaleway.IotDevice("mainIotDevice", {hubId: mainIotHub.id});
 * ```
 *
 * ### With custom certificate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as local from "@pulumi/local";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const mainIotHub = new scaleway.IotHub("mainIotHub", {productPlan: "plan_shared"});
 * const deviceCert = local.getFile({
 *     filename: "device-certificate.pem",
 * });
 * const mainIotDevice = new scaleway.IotDevice("mainIotDevice", {
 *     hubId: mainIotHub.id,
 *     certificate: {
 *         crt: deviceCert.then(deviceCert => deviceCert.content),
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * IoT devices can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/iotDevice:IotDevice device01 fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class IotDevice extends pulumi.CustomResource {
    /**
     * Get an existing IotDevice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IotDeviceState, opts?: pulumi.CustomResourceOptions): IotDevice {
        return new IotDevice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/iotDevice:IotDevice';

    /**
     * Returns true if the given object is an instance of IotDevice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IotDevice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IotDevice.__pulumiType;
    }

    /**
     * Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
     *
     * > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
     */
    public readonly allowInsecure!: pulumi.Output<boolean | undefined>;
    /**
     * Allow more than one simultaneous connection using the same device credentials.
     *
     * > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
     */
    public readonly allowMultipleConnections!: pulumi.Output<boolean | undefined>;
    /**
     * The certificate bundle of the device.
     */
    public readonly certificate!: pulumi.Output<outputs.IotDeviceCertificate>;
    /**
     * The date and time the device was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The description of the IoT device (e.g. `living room`).
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the hub on which this device will be created.
     */
    public readonly hubId!: pulumi.Output<string>;
    /**
     * The current connection status of the device.
     */
    public /*out*/ readonly isConnected!: pulumi.Output<boolean>;
    /**
     * The last MQTT activity of the device.
     */
    public /*out*/ readonly lastActivityAt!: pulumi.Output<string>;
    /**
     * Rules that define which messages are authorized or denied based on their topic.
     */
    public readonly messageFilters!: pulumi.Output<outputs.IotDeviceMessageFilters | undefined>;
    /**
     * The name of the IoT device you want to create (e.g. `my-device`).
     *
     * > **Important:** Updates to `name` will destroy and recreate a new resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region you want to attach the resource to
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The current status of the device.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The date and time the device resource was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a IotDevice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IotDeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IotDeviceArgs | IotDeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IotDeviceState | undefined;
            resourceInputs["allowInsecure"] = state ? state.allowInsecure : undefined;
            resourceInputs["allowMultipleConnections"] = state ? state.allowMultipleConnections : undefined;
            resourceInputs["certificate"] = state ? state.certificate : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["hubId"] = state ? state.hubId : undefined;
            resourceInputs["isConnected"] = state ? state.isConnected : undefined;
            resourceInputs["lastActivityAt"] = state ? state.lastActivityAt : undefined;
            resourceInputs["messageFilters"] = state ? state.messageFilters : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as IotDeviceArgs | undefined;
            if ((!args || args.hubId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hubId'");
            }
            resourceInputs["allowInsecure"] = args ? args.allowInsecure : undefined;
            resourceInputs["allowMultipleConnections"] = args ? args.allowMultipleConnections : undefined;
            resourceInputs["certificate"] = args ? args.certificate : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hubId"] = args ? args.hubId : undefined;
            resourceInputs["messageFilters"] = args ? args.messageFilters : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["isConnected"] = undefined /*out*/;
            resourceInputs["lastActivityAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IotDevice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IotDevice resources.
 */
export interface IotDeviceState {
    /**
     * Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
     *
     * > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
     */
    allowInsecure?: pulumi.Input<boolean>;
    /**
     * Allow more than one simultaneous connection using the same device credentials.
     *
     * > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
     */
    allowMultipleConnections?: pulumi.Input<boolean>;
    /**
     * The certificate bundle of the device.
     */
    certificate?: pulumi.Input<inputs.IotDeviceCertificate>;
    /**
     * The date and time the device was created.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The description of the IoT device (e.g. `living room`).
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the hub on which this device will be created.
     */
    hubId?: pulumi.Input<string>;
    /**
     * The current connection status of the device.
     */
    isConnected?: pulumi.Input<boolean>;
    /**
     * The last MQTT activity of the device.
     */
    lastActivityAt?: pulumi.Input<string>;
    /**
     * Rules that define which messages are authorized or denied based on their topic.
     */
    messageFilters?: pulumi.Input<inputs.IotDeviceMessageFilters>;
    /**
     * The name of the IoT device you want to create (e.g. `my-device`).
     *
     * > **Important:** Updates to `name` will destroy and recreate a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
    /**
     * The current status of the device.
     */
    status?: pulumi.Input<string>;
    /**
     * The date and time the device resource was updated.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IotDevice resource.
 */
export interface IotDeviceArgs {
    /**
     * Allow plain and server-authenticated TLS connections in addition to mutually-authenticated ones.
     *
     * > **Important:** Updates to `allowInsecure` can disconnect eventually connected devices.
     */
    allowInsecure?: pulumi.Input<boolean>;
    /**
     * Allow more than one simultaneous connection using the same device credentials.
     *
     * > **Important:** Updates to `allowMultipleConnections` can disconnect eventually connected devices.
     */
    allowMultipleConnections?: pulumi.Input<boolean>;
    /**
     * The certificate bundle of the device.
     */
    certificate?: pulumi.Input<inputs.IotDeviceCertificate>;
    /**
     * The description of the IoT device (e.g. `living room`).
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the hub on which this device will be created.
     */
    hubId: pulumi.Input<string>;
    /**
     * Rules that define which messages are authorized or denied based on their topic.
     */
    messageFilters?: pulumi.Input<inputs.IotDeviceMessageFilters>;
    /**
     * The name of the IoT device you want to create (e.g. `my-device`).
     *
     * > **Important:** Updates to `name` will destroy and recreate a new resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The region you want to attach the resource to
     */
    region?: pulumi.Input<string>;
}
