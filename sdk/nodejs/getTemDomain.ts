// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Gets information about a transactional email domain.
 */
export function getTemDomain(args?: GetTemDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetTemDomainResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("scaleway:index/getTemDomain:getTemDomain", {
        "domainId": args.domainId,
        "name": args.name,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getTemDomain.
 */
export interface GetTemDomainArgs {
    domainId?: string;
    /**
     * The domain name.
     * Only one of `name` and `id` should be specified.
     */
    name?: string;
    /**
     * `region`) The region in which the domain exists.
     */
    region?: string;
}

/**
 * A collection of values returned by getTemDomain.
 */
export interface GetTemDomainResult {
    readonly acceptTos: boolean;
    /**
     * The date and time of the Transaction Email Domain's creation (RFC 3339 format).
     */
    readonly createdAt: string;
    /**
     * The DKIM public key, as should be recorded in the DNS zone.
     */
    readonly dkimConfig: string;
    readonly domainId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The error message if the last check failed.
     */
    readonly lastError: string;
    /**
     * The date and time the domain was last found to be valid (RFC 3339 format).
     */
    readonly lastValidAt: string;
    readonly name?: string;
    /**
     * The date and time of the next scheduled check (RFC 3339 format).
     */
    readonly nextCheckAt: string;
    readonly projectId: string;
    readonly region?: string;
    /**
     * The date and time of the revocation of the domain (RFC 3339 format).
     */
    readonly revokedAt: string;
    /**
     * The SMTP host to use to send emails.
     */
    readonly smtpHost: string;
    /**
     * The SMTP port to use to send emails over TLS.
     */
    readonly smtpPort: number;
    /**
     * The SMTP port to use to send emails over TLS.
     */
    readonly smtpPortAlternative: number;
    /**
     * The SMTP port to use to send emails.
     */
    readonly smtpPortUnsecure: number;
    /**
     * The SMTPS port to use to send emails over TLS Wrapper.
     */
    readonly smtpsPort: number;
    /**
     * The SMTPS port to use to send emails over TLS Wrapper.
     */
    readonly smtpsPortAlternative: number;
    /**
     * The snippet of the SPF record that should be registered in the DNS zone.
     */
    readonly spfConfig: string;
    /**
     * The status of the Transaction Email Domain.
     */
    readonly status: string;
}
/**
 * Gets information about a transactional email domain.
 */
export function getTemDomainOutput(args?: GetTemDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTemDomainResult> {
    return pulumi.output(args).apply((a: any) => getTemDomain(a, opts))
}

/**
 * A collection of arguments for invoking getTemDomain.
 */
export interface GetTemDomainOutputArgs {
    domainId?: pulumi.Input<string>;
    /**
     * The domain name.
     * Only one of `name` and `id` should be specified.
     */
    name?: pulumi.Input<string>;
    /**
     * `region`) The region in which the domain exists.
     */
    region?: pulumi.Input<string>;
}
