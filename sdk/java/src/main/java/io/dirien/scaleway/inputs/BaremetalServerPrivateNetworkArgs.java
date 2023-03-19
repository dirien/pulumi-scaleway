// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BaremetalServerPrivateNetworkArgs extends com.pulumi.resources.ResourceArgs {

    public static final BaremetalServerPrivateNetworkArgs Empty = new BaremetalServerPrivateNetworkArgs();

    /**
     * The date and time of the creation of the private network.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return The date and time of the creation of the private network.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * The id of the private network to attach.
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return The id of the private network to attach.
     * 
     */
    public Output<String> id() {
        return this.id;
    }

    /**
     * The private network status.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return The private network status.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The date and time of the last update of the private network.
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return The date and time of the last update of the private network.
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    /**
     * The VLAN ID associated to the private network.
     * 
     */
    @Import(name="vlan")
    private @Nullable Output<Integer> vlan;

    /**
     * @return The VLAN ID associated to the private network.
     * 
     */
    public Optional<Output<Integer>> vlan() {
        return Optional.ofNullable(this.vlan);
    }

    private BaremetalServerPrivateNetworkArgs() {}

    private BaremetalServerPrivateNetworkArgs(BaremetalServerPrivateNetworkArgs $) {
        this.createdAt = $.createdAt;
        this.id = $.id;
        this.status = $.status;
        this.updatedAt = $.updatedAt;
        this.vlan = $.vlan;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BaremetalServerPrivateNetworkArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BaremetalServerPrivateNetworkArgs $;

        public Builder() {
            $ = new BaremetalServerPrivateNetworkArgs();
        }

        public Builder(BaremetalServerPrivateNetworkArgs defaults) {
            $ = new BaremetalServerPrivateNetworkArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt The date and time of the creation of the private network.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt The date and time of the creation of the private network.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param id The id of the private network to attach.
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The id of the private network to attach.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param status The private network status.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status The private network status.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param updatedAt The date and time of the last update of the private network.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt The date and time of the last update of the private network.
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        /**
         * @param vlan The VLAN ID associated to the private network.
         * 
         * @return builder
         * 
         */
        public Builder vlan(@Nullable Output<Integer> vlan) {
            $.vlan = vlan;
            return this;
        }

        /**
         * @param vlan The VLAN ID associated to the private network.
         * 
         * @return builder
         * 
         */
        public Builder vlan(Integer vlan) {
            return vlan(Output.of(vlan));
        }

        public BaremetalServerPrivateNetworkArgs build() {
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            return $;
        }
    }

}
