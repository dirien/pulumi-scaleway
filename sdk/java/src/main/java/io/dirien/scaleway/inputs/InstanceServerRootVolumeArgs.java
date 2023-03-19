// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceServerRootVolumeArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceServerRootVolumeArgs Empty = new InstanceServerRootVolumeArgs();

    @Import(name="boot")
    private @Nullable Output<Boolean> boot;

    public Optional<Output<Boolean>> boot() {
        return Optional.ofNullable(this.boot);
    }

    /**
     * Forces deletion of the root volume on instance termination.
     * 
     */
    @Import(name="deleteOnTermination")
    private @Nullable Output<Boolean> deleteOnTermination;

    /**
     * @return Forces deletion of the root volume on instance termination.
     * 
     */
    public Optional<Output<Boolean>> deleteOnTermination() {
        return Optional.ofNullable(this.deleteOnTermination);
    }

    /**
     * The name of the server.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the server.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Size of the root volume in gigabytes.
     * To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
     * check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
     * Updates to this field will recreate a new resource.
     * 
     */
    @Import(name="sizeInGb")
    private @Nullable Output<Integer> sizeInGb;

    /**
     * @return Size of the root volume in gigabytes.
     * To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
     * check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
     * Updates to this field will recreate a new resource.
     * 
     */
    public Optional<Output<Integer>> sizeInGb() {
        return Optional.ofNullable(this.sizeInGb);
    }

    /**
     * The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
     * 
     */
    @Import(name="volumeId")
    private @Nullable Output<String> volumeId;

    /**
     * @return The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
     * 
     */
    public Optional<Output<String>> volumeId() {
        return Optional.ofNullable(this.volumeId);
    }

    /**
     * Volume type of root volume, can be `b_ssd` or `l_ssd`, default value depends on server type
     * 
     */
    @Import(name="volumeType")
    private @Nullable Output<String> volumeType;

    /**
     * @return Volume type of root volume, can be `b_ssd` or `l_ssd`, default value depends on server type
     * 
     */
    public Optional<Output<String>> volumeType() {
        return Optional.ofNullable(this.volumeType);
    }

    private InstanceServerRootVolumeArgs() {}

    private InstanceServerRootVolumeArgs(InstanceServerRootVolumeArgs $) {
        this.boot = $.boot;
        this.deleteOnTermination = $.deleteOnTermination;
        this.name = $.name;
        this.sizeInGb = $.sizeInGb;
        this.volumeId = $.volumeId;
        this.volumeType = $.volumeType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceServerRootVolumeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceServerRootVolumeArgs $;

        public Builder() {
            $ = new InstanceServerRootVolumeArgs();
        }

        public Builder(InstanceServerRootVolumeArgs defaults) {
            $ = new InstanceServerRootVolumeArgs(Objects.requireNonNull(defaults));
        }

        public Builder boot(@Nullable Output<Boolean> boot) {
            $.boot = boot;
            return this;
        }

        public Builder boot(Boolean boot) {
            return boot(Output.of(boot));
        }

        /**
         * @param deleteOnTermination Forces deletion of the root volume on instance termination.
         * 
         * @return builder
         * 
         */
        public Builder deleteOnTermination(@Nullable Output<Boolean> deleteOnTermination) {
            $.deleteOnTermination = deleteOnTermination;
            return this;
        }

        /**
         * @param deleteOnTermination Forces deletion of the root volume on instance termination.
         * 
         * @return builder
         * 
         */
        public Builder deleteOnTermination(Boolean deleteOnTermination) {
            return deleteOnTermination(Output.of(deleteOnTermination));
        }

        /**
         * @param name The name of the server.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the server.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sizeInGb Size of the root volume in gigabytes.
         * To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
         * check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
         * Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder sizeInGb(@Nullable Output<Integer> sizeInGb) {
            $.sizeInGb = sizeInGb;
            return this;
        }

        /**
         * @param sizeInGb Size of the root volume in gigabytes.
         * To find the right size use [this endpoint](https://api.scaleway.com/instance/v1/zones/fr-par-1/products/servers) and
         * check the `volumes_constraint.{min|max}_size` (in bytes) for your `commercial_type`.
         * Updates to this field will recreate a new resource.
         * 
         * @return builder
         * 
         */
        public Builder sizeInGb(Integer sizeInGb) {
            return sizeInGb(Output.of(sizeInGb));
        }

        /**
         * @param volumeId The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(@Nullable Output<String> volumeId) {
            $.volumeId = volumeId;
            return this;
        }

        /**
         * @param volumeId The volume ID of the root volume of the server, allows you to create server with an existing volume. If empty, will be computed to a created volume ID.
         * 
         * @return builder
         * 
         */
        public Builder volumeId(String volumeId) {
            return volumeId(Output.of(volumeId));
        }

        /**
         * @param volumeType Volume type of root volume, can be `b_ssd` or `l_ssd`, default value depends on server type
         * 
         * @return builder
         * 
         */
        public Builder volumeType(@Nullable Output<String> volumeType) {
            $.volumeType = volumeType;
            return this;
        }

        /**
         * @param volumeType Volume type of root volume, can be `b_ssd` or `l_ssd`, default value depends on server type
         * 
         * @return builder
         * 
         */
        public Builder volumeType(String volumeType) {
            return volumeType(Output.of(volumeType));
        }

        public InstanceServerRootVolumeArgs build() {
            return $;
        }
    }

}
