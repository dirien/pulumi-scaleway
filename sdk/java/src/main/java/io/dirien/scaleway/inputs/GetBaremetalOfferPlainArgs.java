// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBaremetalOfferPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBaremetalOfferPlainArgs Empty = new GetBaremetalOfferPlainArgs();

    @Import(name="includeDisabled")
    private @Nullable Boolean includeDisabled;

    public Optional<Boolean> includeDisabled() {
        return Optional.ofNullable(this.includeDisabled);
    }

    /**
     * The offer name. Only one of `name` and `offer_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The offer name. Only one of `name` and `offer_id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The offer id. Only one of `name` and `offer_id` should be specified.
     * 
     */
    @Import(name="offerId")
    private @Nullable String offerId;

    /**
     * @return The offer id. Only one of `name` and `offer_id` should be specified.
     * 
     */
    public Optional<String> offerId() {
        return Optional.ofNullable(this.offerId);
    }

    /**
     * Period of subscription the desired offer. Should be `hourly` or `monthly`.
     * 
     */
    @Import(name="subscriptionPeriod")
    private @Nullable String subscriptionPeriod;

    /**
     * @return Period of subscription the desired offer. Should be `hourly` or `monthly`.
     * 
     */
    public Optional<String> subscriptionPeriod() {
        return Optional.ofNullable(this.subscriptionPeriod);
    }

    /**
     * `zone`) The zone in which the offer should be created.
     * 
     */
    @Import(name="zone")
    private @Nullable String zone;

    /**
     * @return `zone`) The zone in which the offer should be created.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetBaremetalOfferPlainArgs() {}

    private GetBaremetalOfferPlainArgs(GetBaremetalOfferPlainArgs $) {
        this.includeDisabled = $.includeDisabled;
        this.name = $.name;
        this.offerId = $.offerId;
        this.subscriptionPeriod = $.subscriptionPeriod;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBaremetalOfferPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBaremetalOfferPlainArgs $;

        public Builder() {
            $ = new GetBaremetalOfferPlainArgs();
        }

        public Builder(GetBaremetalOfferPlainArgs defaults) {
            $ = new GetBaremetalOfferPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder includeDisabled(@Nullable Boolean includeDisabled) {
            $.includeDisabled = includeDisabled;
            return this;
        }

        /**
         * @param name The offer name. Only one of `name` and `offer_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param offerId The offer id. Only one of `name` and `offer_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder offerId(@Nullable String offerId) {
            $.offerId = offerId;
            return this;
        }

        /**
         * @param subscriptionPeriod Period of subscription the desired offer. Should be `hourly` or `monthly`.
         * 
         * @return builder
         * 
         */
        public Builder subscriptionPeriod(@Nullable String subscriptionPeriod) {
            $.subscriptionPeriod = subscriptionPeriod;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which the offer should be created.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable String zone) {
            $.zone = zone;
            return this;
        }

        public GetBaremetalOfferPlainArgs build() {
            return $;
        }
    }

}
