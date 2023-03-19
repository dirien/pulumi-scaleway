// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LbRouteState extends com.pulumi.resources.ResourceArgs {

    public static final LbRouteState Empty = new LbRouteState();

    /**
     * The ID of the backend to which the route is associated.
     * 
     */
    @Import(name="backendId")
    private @Nullable Output<String> backendId;

    /**
     * @return The ID of the backend to which the route is associated.
     * 
     */
    public Optional<Output<String>> backendId() {
        return Optional.ofNullable(this.backendId);
    }

    /**
     * The date at which the route was created.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The date at which the route was created.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The ID of the frontend to which the route is associated.
     * 
     */
    @Import(name="frontendId")
    private @Nullable Output<String> frontendId;

    /**
     * @return The ID of the frontend to which the route is associated.
     * 
     */
    public Optional<Output<String>> frontendId() {
        return Optional.ofNullable(this.frontendId);
    }

    /**
     * The Host request header specifies the host of the server to which the request is being sent.
     * Only one of `match_sni` and `match_host_header` should be specified.
     * 
     */
    @Import(name="matchHostHeader")
    private @Nullable Output<String> matchHostHeader;

    /**
     * @return The Host request header specifies the host of the server to which the request is being sent.
     * Only one of `match_sni` and `match_host_header` should be specified.
     * 
     */
    public Optional<Output<String>> matchHostHeader() {
        return Optional.ofNullable(this.matchHostHeader);
    }

    /**
     * The Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer.
     * Only one of `match_sni` and `match_host_header` should be specified.
     * 
     */
    @Import(name="matchSni")
    private @Nullable Output<String> matchSni;

    /**
     * @return The Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer.
     * Only one of `match_sni` and `match_host_header` should be specified.
     * 
     */
    public Optional<Output<String>> matchSni() {
        return Optional.ofNullable(this.matchSni);
    }

    /**
     * The date at which the route was last updated.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The date at which the route was last updated.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private LbRouteState() {}

    private LbRouteState(LbRouteState $) {
        this.backendId = $.backendId;
        this.createdAt = $.createdAt;
        this.frontendId = $.frontendId;
        this.matchHostHeader = $.matchHostHeader;
        this.matchSni = $.matchSni;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LbRouteState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LbRouteState $;

        public Builder() {
            $ = new LbRouteState();
        }

        public Builder(LbRouteState defaults) {
            $ = new LbRouteState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backendId The ID of the backend to which the route is associated.
         * 
         * @return builder
         * 
         */
        public Builder backendId(@Nullable Output<String> backendId) {
            $.backendId = backendId;
            return this;
        }

        /**
         * @param backendId The ID of the backend to which the route is associated.
         * 
         * @return builder
         * 
         */
        public Builder backendId(String backendId) {
            return backendId(Output.of(backendId));
        }

        /**
         * @param createdAt The date at which the route was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The date at which the route was created.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param frontendId The ID of the frontend to which the route is associated.
         * 
         * @return builder
         * 
         */
        public Builder frontendId(@Nullable Output<String> frontendId) {
            $.frontendId = frontendId;
            return this;
        }

        /**
         * @param frontendId The ID of the frontend to which the route is associated.
         * 
         * @return builder
         * 
         */
        public Builder frontendId(String frontendId) {
            return frontendId(Output.of(frontendId));
        }

        /**
         * @param matchHostHeader The Host request header specifies the host of the server to which the request is being sent.
         * Only one of `match_sni` and `match_host_header` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder matchHostHeader(@Nullable Output<String> matchHostHeader) {
            $.matchHostHeader = matchHostHeader;
            return this;
        }

        /**
         * @param matchHostHeader The Host request header specifies the host of the server to which the request is being sent.
         * Only one of `match_sni` and `match_host_header` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder matchHostHeader(String matchHostHeader) {
            return matchHostHeader(Output.of(matchHostHeader));
        }

        /**
         * @param matchSni The Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer.
         * Only one of `match_sni` and `match_host_header` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder matchSni(@Nullable Output<String> matchSni) {
            $.matchSni = matchSni;
            return this;
        }

        /**
         * @param matchSni The Server Name Indication TLS extension field from an incoming connection made via an SSL/TLS transport layer.
         * Only one of `match_sni` and `match_host_header` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder matchSni(String matchSni) {
            return matchSni(Output.of(matchSni));
        }

        /**
         * @param updatedAt The date at which the route was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The date at which the route was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public LbRouteState build() {
            return $;
        }
    }

}
