// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLbArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLbArgs Empty = new GetLbArgs();

    @Import(name="lbId")
    private @Nullable Output<String> lbId;

    public Optional<Output<String>> lbId() {
        return Optional.ofNullable(this.lbId);
    }

    /**
     * The load balancer name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The load balancer name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="releaseIp")
    private @Nullable Output<Boolean> releaseIp;

    public Optional<Output<Boolean>> releaseIp() {
        return Optional.ofNullable(this.releaseIp);
    }

    /**
     * (Defaults to provider `zone`) The zone in which the LB exists.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return (Defaults to provider `zone`) The zone in which the LB exists.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetLbArgs() {}

    private GetLbArgs(GetLbArgs $) {
        this.lbId = $.lbId;
        this.name = $.name;
        this.releaseIp = $.releaseIp;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLbArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLbArgs $;

        public Builder() {
            $ = new GetLbArgs();
        }

        public Builder(GetLbArgs defaults) {
            $ = new GetLbArgs(Objects.requireNonNull(defaults));
        }

        public Builder lbId(@Nullable Output<String> lbId) {
            $.lbId = lbId;
            return this;
        }

        public Builder lbId(String lbId) {
            return lbId(Output.of(lbId));
        }

        /**
         * @param name The load balancer name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The load balancer name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder releaseIp(@Nullable Output<Boolean> releaseIp) {
            $.releaseIp = releaseIp;
            return this;
        }

        public Builder releaseIp(Boolean releaseIp) {
            return releaseIp(Output.of(releaseIp));
        }

        /**
         * @param zone (Defaults to provider `zone`) The zone in which the LB exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone (Defaults to provider `zone`) The zone in which the LB exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetLbArgs build() {
            return $;
        }
    }

}
