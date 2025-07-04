// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The `scaleway.DomainRegistration` resource allows you to purchase and manage domain registrations with Scaleway. Using this resource you can register one or more domains for a specified duration, configure auto-renewal and DNSSEC options, and set contact information. You can supply an owner contact either by providing an existing contact ID or by specifying the complete contact details. The resource automatically returns additional contact information (administrative and technical) as provided by the Scaleway API.
 *
 * Refer to the [Domains and DNS documentation](https://www.scaleway.com/en/docs/network/domains-and-dns/) and the [API documentation](https://developers.scaleway.com/) for more details.
 *
 * ## Example Usage
 *
 * ### Purchase a Single Domain
 *
 * The following example purchases a domain with a one-year registration period and specifies the owner contact details. Administrative and technical contacts are returned by the API.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const test = new scaleway.DomainRegistration("test", {
 *     domainNames: ["example.com"],
 *     durationInYears: 1,
 *     ownerContact: {
 *         addressLine1: "123 Main Street",
 *         city: "Paris",
 *         companyIdentificationCode: "123456789",
 *         country: "FR",
 *         email: "john.doe@example.com",
 *         firstname: "John",
 *         lastname: "DOE",
 *         legalForm: "individual",
 *         phoneNumber: "+1.23456789",
 *         vatIdentificationCode: "FR12345678901",
 *         zip: "75001",
 *     },
 * });
 * ```
 *
 * ### Update Domain Settings
 *
 * You can update the resource to enable auto-renewal and DNSSEC:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const test = new scaleway.DomainRegistration("test", {
 *     autoRenew: true,
 *     dnssec: true,
 *     domainNames: ["example.com"],
 *     durationInYears: 1,
 *     ownerContact: {
 *         addressLine1: "123 Main Street",
 *         city: "Paris",
 *         companyIdentificationCode: "123456789",
 *         country: "FR",
 *         email: "john.doe@example.com",
 *         firstname: "John",
 *         lastname: "DOE",
 *         legalForm: "individual",
 *         phoneNumber: "+1.23456789",
 *         vatIdentificationCode: "FR12345678901",
 *         zip: "75001",
 *     },
 * });
 * ```
 *
 * ### Purchase Multiple Domains
 *
 * The following example registers several domains in one go:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const multi = new scaleway.DomainRegistration("multi", {
 *     domainNames: [
 *         "domain1.com",
 *         "domain2.com",
 *         "domain3.com",
 *     ],
 *     durationInYears: 1,
 *     ownerContact: {
 *         addressLine1: "123 Main Street",
 *         city: "Paris",
 *         companyIdentificationCode: "123456789",
 *         country: "FR",
 *         email: "john.doe@example.com",
 *         firstname: "John",
 *         lastname: "DOE",
 *         legalForm: "individual",
 *         phoneNumber: "+1.23456789",
 *         vatIdentificationCode: "FR12345678901",
 *         zip: "75001",
 *     },
 * });
 * ```
 *
 * ## Contact Blocks
 *
 * Each contact block supports the following attributes:
 *
 * - `legalForm` (Required, String): Legal form of the contact.
 * - `firstname` (Required, String): First name.
 * - `lastname` (Required, String): Last name.
 * - `companyName` (Optional, String): Company name.
 * - `email` (Required, String): Primary email.
 * - `phoneNumber` (Required, String): Primary phone number.
 * - `addressLine1` (Required, String): Primary address.
 * - `zip` (Required, String): Postal code.
 * - `city` (Required, String): City.
 * - `country` (Required, String): Country code (ISO format).
 * - `vatIdentificationCode` (Required, String): VAT identification code.
 * - `companyIdentificationCode` (Required, String): Company identification code.
 *
 * ## Import
 *
 * To import an existing domain registration, use:
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/domainRegistration:DomainRegistration test <project_id>/<task_id>
 * ```
 */
export class DomainRegistration extends pulumi.CustomResource {
    /**
     * Get an existing DomainRegistration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainRegistrationState, opts?: pulumi.CustomResourceOptions): DomainRegistration {
        return new DomainRegistration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/domainRegistration:DomainRegistration';

    /**
     * Returns true if the given object is an instance of DomainRegistration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainRegistration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainRegistration.__pulumiType;
    }

    /**
     * : Administrative contact information.
     */
    public /*out*/ readonly administrativeContacts!: pulumi.Output<outputs.DomainRegistrationAdministrativeContact[]>;
    /**
     * : Enables or disables auto-renewal.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * : Enables or disables DNSSEC.
     */
    public readonly dnssec!: pulumi.Output<boolean | undefined>;
    /**
     * : A list of domain names to be registered.
     */
    public readonly domainNames!: pulumi.Output<string[]>;
    /**
     * DNSSEC DS record configuration.
     */
    public /*out*/ readonly dsRecords!: pulumi.Output<outputs.DomainRegistrationDsRecord[]>;
    /**
     * : The registration period in years.
     */
    public readonly durationInYears!: pulumi.Output<number | undefined>;
    /**
     * : Details of the owner contact.
     */
    public readonly ownerContact!: pulumi.Output<outputs.DomainRegistrationOwnerContact>;
    /**
     * : The ID of an existing owner contact.
     */
    public readonly ownerContactId!: pulumi.Output<string>;
    /**
     * : The Scaleway project ID.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * ID of the task that created the domain.
     */
    public /*out*/ readonly taskId!: pulumi.Output<string>;
    /**
     * : Technical contact information.
     */
    public /*out*/ readonly technicalContacts!: pulumi.Output<outputs.DomainRegistrationTechnicalContact[]>;

    /**
     * Create a DomainRegistration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainRegistrationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainRegistrationArgs | DomainRegistrationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainRegistrationState | undefined;
            resourceInputs["administrativeContacts"] = state ? state.administrativeContacts : undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["dnssec"] = state ? state.dnssec : undefined;
            resourceInputs["domainNames"] = state ? state.domainNames : undefined;
            resourceInputs["dsRecords"] = state ? state.dsRecords : undefined;
            resourceInputs["durationInYears"] = state ? state.durationInYears : undefined;
            resourceInputs["ownerContact"] = state ? state.ownerContact : undefined;
            resourceInputs["ownerContactId"] = state ? state.ownerContactId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["taskId"] = state ? state.taskId : undefined;
            resourceInputs["technicalContacts"] = state ? state.technicalContacts : undefined;
        } else {
            const args = argsOrState as DomainRegistrationArgs | undefined;
            if ((!args || args.domainNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainNames'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["dnssec"] = args ? args.dnssec : undefined;
            resourceInputs["domainNames"] = args ? args.domainNames : undefined;
            resourceInputs["durationInYears"] = args ? args.durationInYears : undefined;
            resourceInputs["ownerContact"] = args ? args.ownerContact : undefined;
            resourceInputs["ownerContactId"] = args ? args.ownerContactId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["administrativeContacts"] = undefined /*out*/;
            resourceInputs["dsRecords"] = undefined /*out*/;
            resourceInputs["taskId"] = undefined /*out*/;
            resourceInputs["technicalContacts"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainRegistration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainRegistration resources.
 */
export interface DomainRegistrationState {
    /**
     * : Administrative contact information.
     */
    administrativeContacts?: pulumi.Input<pulumi.Input<inputs.DomainRegistrationAdministrativeContact>[]>;
    /**
     * : Enables or disables auto-renewal.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * : Enables or disables DNSSEC.
     */
    dnssec?: pulumi.Input<boolean>;
    /**
     * : A list of domain names to be registered.
     */
    domainNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * DNSSEC DS record configuration.
     */
    dsRecords?: pulumi.Input<pulumi.Input<inputs.DomainRegistrationDsRecord>[]>;
    /**
     * : The registration period in years.
     */
    durationInYears?: pulumi.Input<number>;
    /**
     * : Details of the owner contact.
     */
    ownerContact?: pulumi.Input<inputs.DomainRegistrationOwnerContact>;
    /**
     * : The ID of an existing owner contact.
     */
    ownerContactId?: pulumi.Input<string>;
    /**
     * : The Scaleway project ID.
     */
    projectId?: pulumi.Input<string>;
    /**
     * ID of the task that created the domain.
     */
    taskId?: pulumi.Input<string>;
    /**
     * : Technical contact information.
     */
    technicalContacts?: pulumi.Input<pulumi.Input<inputs.DomainRegistrationTechnicalContact>[]>;
}

/**
 * The set of arguments for constructing a DomainRegistration resource.
 */
export interface DomainRegistrationArgs {
    /**
     * : Enables or disables auto-renewal.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * : Enables or disables DNSSEC.
     */
    dnssec?: pulumi.Input<boolean>;
    /**
     * : A list of domain names to be registered.
     */
    domainNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * : The registration period in years.
     */
    durationInYears?: pulumi.Input<number>;
    /**
     * : Details of the owner contact.
     */
    ownerContact?: pulumi.Input<inputs.DomainRegistrationOwnerContact>;
    /**
     * : The ID of an existing owner contact.
     */
    ownerContactId?: pulumi.Input<string>;
    /**
     * : The Scaleway project ID.
     */
    projectId?: pulumi.Input<string>;
}
