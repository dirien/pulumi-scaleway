// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLbIpPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLbIpPlainArgs Empty = new GetLbIpPlainArgs();

    /**
     * The IP address.
     * Only one of `ip_address` and `ip_id` should be specified.
     * 
     */
    @Import(name="ipAddress")
    private @Nullable String ipAddress;

    /**
     * @return The IP address.
     * Only one of `ip_address` and `ip_id` should be specified.
     * 
     */
    public Optional<String> ipAddress() {
        return Optional.ofNullable(this.ipAddress);
    }

    /**
     * The IP ID.
     * Only one of `ip_address` and `ip_id` should be specified.
     * 
     */
    @Import(name="ipId")
    private @Nullable String ipId;

    /**
     * @return The IP ID.
     * Only one of `ip_address` and `ip_id` should be specified.
     * 
     */
    public Optional<String> ipId() {
        return Optional.ofNullable(this.ipId);
    }

    private GetLbIpPlainArgs() {}

    private GetLbIpPlainArgs(GetLbIpPlainArgs $) {
        this.ipAddress = $.ipAddress;
        this.ipId = $.ipId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLbIpPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLbIpPlainArgs $;

        public Builder() {
            $ = new GetLbIpPlainArgs();
        }

        public Builder(GetLbIpPlainArgs defaults) {
            $ = new GetLbIpPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ipAddress The IP address.
         * Only one of `ip_address` and `ip_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder ipAddress(@Nullable String ipAddress) {
            $.ipAddress = ipAddress;
            return this;
        }

        /**
         * @param ipId The IP ID.
         * Only one of `ip_address` and `ip_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder ipId(@Nullable String ipId) {
            $.ipId = ipId;
            return this;
        }

        public GetLbIpPlainArgs build() {
            return $;
        }
    }

}
