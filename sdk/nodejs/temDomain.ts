// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates and manages Scaleway Transactional Email Domains.
 * For more information see [the documentation](https://www.scaleway.com/en/developers/api/transactional-email).
 *
 * ## Example Usage
 *
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const main = new scaleway.TemDomain("main", {acceptTos: true});
 * ```
 *
 * ### Add the required records to your DNS zone
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as scaleway from "@ediri/scaleway";
 *
 * const config = new pulumi.Config();
 * const domainName = config.require("domainName");
 * const main = new scaleway.TemDomain("main", {acceptTos: true});
 * const spf = new scaleway.DomainRecord("spf", {
 *     dnsZone: domainName,
 *     type: "TXT",
 *     data: pulumi.interpolate`v=spf1 ${main.spfConfig} -all`,
 * });
 * const dkim = new scaleway.DomainRecord("dkim", {
 *     dnsZone: domainName,
 *     type: "TXT",
 *     data: main.dkimConfig,
 * });
 * const mx = new scaleway.DomainRecord("mx", {
 *     dnsZone: domainName,
 *     type: "MX",
 *     data: ".",
 * });
 * const dmarc = new scaleway.DomainRecord("dmarc", {
 *     dnsZone: domainName,
 *     type: "TXT",
 *     data: main.dmarcConfig,
 * });
 * ```
 *
 * ## Import
 *
 * Domains can be imported using the `{region}/{id}`, e.g.
 *
 * bash
 *
 * ```sh
 * $ pulumi import scaleway:index/temDomain:TemDomain main fr-par/11111111-1111-1111-1111-111111111111
 * ```
 */
export class TemDomain extends pulumi.CustomResource {
    /**
     * Get an existing TemDomain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TemDomainState, opts?: pulumi.CustomResourceOptions): TemDomain {
        return new TemDomain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'scaleway:index/temDomain:TemDomain';

    /**
     * Returns true if the given object is an instance of TemDomain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TemDomain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TemDomain.__pulumiType;
    }

    /**
     * Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
     * > **Important:**  This attribute must be set to `true`.
     */
    public readonly acceptTos!: pulumi.Output<boolean>;
    /**
     * The date and time of the Transaction Email Domain's creation (RFC 3339 format).
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The DKIM public key, as should be recorded in the DNS zone.
     */
    public /*out*/ readonly dkimConfig!: pulumi.Output<string>;
    /**
     * DMARC record for the domain, as should be recorded in the DNS zone.
     */
    public /*out*/ readonly dmarcConfig!: pulumi.Output<string>;
    /**
     * DMARC name for the domain, as should be recorded in the DNS zone.
     */
    public /*out*/ readonly dmarcName!: pulumi.Output<string>;
    /**
     * The error message if the last check failed.
     */
    public /*out*/ readonly lastError!: pulumi.Output<string>;
    /**
     * The date and time the domain was last found to be valid (RFC 3339 format).
     */
    public /*out*/ readonly lastValidAt!: pulumi.Output<string>;
    /**
     * The Scaleway's blackhole MX server to use if you do not have one.
     */
    public /*out*/ readonly mxBlackhole!: pulumi.Output<string>;
    /**
     * The domain name, must not be used in another Transactional Email Domain.
     * > **Important:** Updates to `name` will recreate the domain.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The date and time of the next scheduled check (RFC 3339 format).
     */
    public /*out*/ readonly nextCheckAt!: pulumi.Output<string>;
    /**
     * `projectId`) The ID of the project the domain is associated with.
     */
    public readonly projectId!: pulumi.Output<string>;
    /**
     * `region`). The region in which the domain should be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The domain's reputation.
     */
    public /*out*/ readonly reputations!: pulumi.Output<outputs.TemDomainReputation[]>;
    /**
     * The date and time of the revocation of the domain (RFC 3339 format).
     */
    public /*out*/ readonly revokedAt!: pulumi.Output<string>;
    /**
     * The SMTP host to use to send emails.
     */
    public /*out*/ readonly smtpHost!: pulumi.Output<string>;
    /**
     * The SMTP port to use to send emails over TLS.
     */
    public /*out*/ readonly smtpPort!: pulumi.Output<number>;
    /**
     * The SMTP port to use to send emails over TLS.
     */
    public /*out*/ readonly smtpPortAlternative!: pulumi.Output<number>;
    /**
     * The SMTP port to use to send emails.
     */
    public /*out*/ readonly smtpPortUnsecure!: pulumi.Output<number>;
    /**
     * SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
     */
    public /*out*/ readonly smtpsAuthUser!: pulumi.Output<string>;
    /**
     * The SMTPS port to use to send emails over TLS Wrapper.
     */
    public /*out*/ readonly smtpsPort!: pulumi.Output<number>;
    /**
     * The SMTPS port to use to send emails over TLS Wrapper.
     */
    public /*out*/ readonly smtpsPortAlternative!: pulumi.Output<number>;
    /**
     * The snippet of the SPF record that should be registered in the DNS zone.
     */
    public /*out*/ readonly spfConfig!: pulumi.Output<string>;
    /**
     * The status of the domain's reputation.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a TemDomain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TemDomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TemDomainArgs | TemDomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TemDomainState | undefined;
            resourceInputs["acceptTos"] = state ? state.acceptTos : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dkimConfig"] = state ? state.dkimConfig : undefined;
            resourceInputs["dmarcConfig"] = state ? state.dmarcConfig : undefined;
            resourceInputs["dmarcName"] = state ? state.dmarcName : undefined;
            resourceInputs["lastError"] = state ? state.lastError : undefined;
            resourceInputs["lastValidAt"] = state ? state.lastValidAt : undefined;
            resourceInputs["mxBlackhole"] = state ? state.mxBlackhole : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nextCheckAt"] = state ? state.nextCheckAt : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["reputations"] = state ? state.reputations : undefined;
            resourceInputs["revokedAt"] = state ? state.revokedAt : undefined;
            resourceInputs["smtpHost"] = state ? state.smtpHost : undefined;
            resourceInputs["smtpPort"] = state ? state.smtpPort : undefined;
            resourceInputs["smtpPortAlternative"] = state ? state.smtpPortAlternative : undefined;
            resourceInputs["smtpPortUnsecure"] = state ? state.smtpPortUnsecure : undefined;
            resourceInputs["smtpsAuthUser"] = state ? state.smtpsAuthUser : undefined;
            resourceInputs["smtpsPort"] = state ? state.smtpsPort : undefined;
            resourceInputs["smtpsPortAlternative"] = state ? state.smtpsPortAlternative : undefined;
            resourceInputs["spfConfig"] = state ? state.spfConfig : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as TemDomainArgs | undefined;
            if ((!args || args.acceptTos === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acceptTos'");
            }
            resourceInputs["acceptTos"] = args ? args.acceptTos : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dkimConfig"] = undefined /*out*/;
            resourceInputs["dmarcConfig"] = undefined /*out*/;
            resourceInputs["dmarcName"] = undefined /*out*/;
            resourceInputs["lastError"] = undefined /*out*/;
            resourceInputs["lastValidAt"] = undefined /*out*/;
            resourceInputs["mxBlackhole"] = undefined /*out*/;
            resourceInputs["nextCheckAt"] = undefined /*out*/;
            resourceInputs["reputations"] = undefined /*out*/;
            resourceInputs["revokedAt"] = undefined /*out*/;
            resourceInputs["smtpHost"] = undefined /*out*/;
            resourceInputs["smtpPort"] = undefined /*out*/;
            resourceInputs["smtpPortAlternative"] = undefined /*out*/;
            resourceInputs["smtpPortUnsecure"] = undefined /*out*/;
            resourceInputs["smtpsAuthUser"] = undefined /*out*/;
            resourceInputs["smtpsPort"] = undefined /*out*/;
            resourceInputs["smtpsPortAlternative"] = undefined /*out*/;
            resourceInputs["spfConfig"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TemDomain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TemDomain resources.
 */
export interface TemDomainState {
    /**
     * Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
     * > **Important:**  This attribute must be set to `true`.
     */
    acceptTos?: pulumi.Input<boolean>;
    /**
     * The date and time of the Transaction Email Domain's creation (RFC 3339 format).
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The DKIM public key, as should be recorded in the DNS zone.
     */
    dkimConfig?: pulumi.Input<string>;
    /**
     * DMARC record for the domain, as should be recorded in the DNS zone.
     */
    dmarcConfig?: pulumi.Input<string>;
    /**
     * DMARC name for the domain, as should be recorded in the DNS zone.
     */
    dmarcName?: pulumi.Input<string>;
    /**
     * The error message if the last check failed.
     */
    lastError?: pulumi.Input<string>;
    /**
     * The date and time the domain was last found to be valid (RFC 3339 format).
     */
    lastValidAt?: pulumi.Input<string>;
    /**
     * The Scaleway's blackhole MX server to use if you do not have one.
     */
    mxBlackhole?: pulumi.Input<string>;
    /**
     * The domain name, must not be used in another Transactional Email Domain.
     * > **Important:** Updates to `name` will recreate the domain.
     */
    name?: pulumi.Input<string>;
    /**
     * The date and time of the next scheduled check (RFC 3339 format).
     */
    nextCheckAt?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the domain is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the domain should be created.
     */
    region?: pulumi.Input<string>;
    /**
     * The domain's reputation.
     */
    reputations?: pulumi.Input<pulumi.Input<inputs.TemDomainReputation>[]>;
    /**
     * The date and time of the revocation of the domain (RFC 3339 format).
     */
    revokedAt?: pulumi.Input<string>;
    /**
     * The SMTP host to use to send emails.
     */
    smtpHost?: pulumi.Input<string>;
    /**
     * The SMTP port to use to send emails over TLS.
     */
    smtpPort?: pulumi.Input<number>;
    /**
     * The SMTP port to use to send emails over TLS.
     */
    smtpPortAlternative?: pulumi.Input<number>;
    /**
     * The SMTP port to use to send emails.
     */
    smtpPortUnsecure?: pulumi.Input<number>;
    /**
     * SMTPS auth user refers to the identifier for a user authorized to send emails via SMTPS, ensuring secure email transmission.
     */
    smtpsAuthUser?: pulumi.Input<string>;
    /**
     * The SMTPS port to use to send emails over TLS Wrapper.
     */
    smtpsPort?: pulumi.Input<number>;
    /**
     * The SMTPS port to use to send emails over TLS Wrapper.
     */
    smtpsPortAlternative?: pulumi.Input<number>;
    /**
     * The snippet of the SPF record that should be registered in the DNS zone.
     */
    spfConfig?: pulumi.Input<string>;
    /**
     * The status of the domain's reputation.
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TemDomain resource.
 */
export interface TemDomainArgs {
    /**
     * Acceptation of the [Term of Service](https://tem.s3.fr-par.scw.cloud/antispam_policy.pdf).
     * > **Important:**  This attribute must be set to `true`.
     */
    acceptTos: pulumi.Input<boolean>;
    /**
     * The domain name, must not be used in another Transactional Email Domain.
     * > **Important:** Updates to `name` will recreate the domain.
     */
    name?: pulumi.Input<string>;
    /**
     * `projectId`) The ID of the project the domain is associated with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * `region`). The region in which the domain should be created.
     */
    region?: pulumi.Input<string>;
}
