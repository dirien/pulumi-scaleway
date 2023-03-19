// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVpcPublicGatewayPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVpcPublicGatewayPlainArgs Empty = new GetVpcPublicGatewayPlainArgs();

    /**
     * Exact name of the public gateway.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Exact name of the public gateway.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="publicGatewayId")
    private @Nullable String publicGatewayId;

    public Optional<String> publicGatewayId() {
        return Optional.ofNullable(this.publicGatewayId);
    }

    /**
     * `zone`) The zone in which
     * the public gateway should be created.
     * 
     */
    @Import(name="zone")
    private @Nullable String zone;

    /**
     * @return `zone`) The zone in which
     * the public gateway should be created.
     * 
     */
    public Optional<String> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetVpcPublicGatewayPlainArgs() {}

    private GetVpcPublicGatewayPlainArgs(GetVpcPublicGatewayPlainArgs $) {
        this.name = $.name;
        this.publicGatewayId = $.publicGatewayId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVpcPublicGatewayPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVpcPublicGatewayPlainArgs $;

        public Builder() {
            $ = new GetVpcPublicGatewayPlainArgs();
        }

        public Builder(GetVpcPublicGatewayPlainArgs defaults) {
            $ = new GetVpcPublicGatewayPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Exact name of the public gateway.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public Builder publicGatewayId(@Nullable String publicGatewayId) {
            $.publicGatewayId = publicGatewayId;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which
         * the public gateway should be created.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable String zone) {
            $.zone = zone;
            return this;
        }

        public GetVpcPublicGatewayPlainArgs build() {
            return $;
        }
    }

}
