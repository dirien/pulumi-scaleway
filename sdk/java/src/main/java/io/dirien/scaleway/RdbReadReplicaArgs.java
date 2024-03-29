// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.scaleway.inputs.RdbReadReplicaDirectAccessArgs;
import io.dirien.scaleway.inputs.RdbReadReplicaPrivateNetworkArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RdbReadReplicaArgs extends com.pulumi.resources.ResourceArgs {

    public static final RdbReadReplicaArgs Empty = new RdbReadReplicaArgs();

    /**
     * Creates a direct access endpoint to rdb replica.
     * 
     */
    @Import(name="directAccess")
    private @Nullable Output<RdbReadReplicaDirectAccessArgs> directAccess;

    /**
     * @return Creates a direct access endpoint to rdb replica.
     * 
     */
    public Optional<Output<RdbReadReplicaDirectAccessArgs>> directAccess() {
        return Optional.ofNullable(this.directAccess);
    }

    /**
     * UUID of the rdb instance.
     * 
     * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
     * 
     */
    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    /**
     * @return UUID of the rdb instance.
     * 
     * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * Create an endpoint in a private network.
     * 
     */
    @Import(name="privateNetwork")
    private @Nullable Output<RdbReadReplicaPrivateNetworkArgs> privateNetwork;

    /**
     * @return Create an endpoint in a private network.
     * 
     */
    public Optional<Output<RdbReadReplicaPrivateNetworkArgs>> privateNetwork() {
        return Optional.ofNullable(this.privateNetwork);
    }

    /**
     * `region`) The region
     * in which the Database read replica should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region
     * in which the Database read replica should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private RdbReadReplicaArgs() {}

    private RdbReadReplicaArgs(RdbReadReplicaArgs $) {
        this.directAccess = $.directAccess;
        this.instanceId = $.instanceId;
        this.privateNetwork = $.privateNetwork;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RdbReadReplicaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RdbReadReplicaArgs $;

        public Builder() {
            $ = new RdbReadReplicaArgs();
        }

        public Builder(RdbReadReplicaArgs defaults) {
            $ = new RdbReadReplicaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param directAccess Creates a direct access endpoint to rdb replica.
         * 
         * @return builder
         * 
         */
        public Builder directAccess(@Nullable Output<RdbReadReplicaDirectAccessArgs> directAccess) {
            $.directAccess = directAccess;
            return this;
        }

        /**
         * @param directAccess Creates a direct access endpoint to rdb replica.
         * 
         * @return builder
         * 
         */
        public Builder directAccess(RdbReadReplicaDirectAccessArgs directAccess) {
            return directAccess(Output.of(directAccess));
        }

        /**
         * @param instanceId UUID of the rdb instance.
         * 
         * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId UUID of the rdb instance.
         * 
         * &gt; **Important:** The replica musts contains at least one of `direct_access` or `private_network`. It can contain both.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param privateNetwork Create an endpoint in a private network.
         * 
         * @return builder
         * 
         */
        public Builder privateNetwork(@Nullable Output<RdbReadReplicaPrivateNetworkArgs> privateNetwork) {
            $.privateNetwork = privateNetwork;
            return this;
        }

        /**
         * @param privateNetwork Create an endpoint in a private network.
         * 
         * @return builder
         * 
         */
        public Builder privateNetwork(RdbReadReplicaPrivateNetworkArgs privateNetwork) {
            return privateNetwork(Output.of(privateNetwork));
        }

        /**
         * @param region `region`) The region
         * in which the Database read replica should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region
         * in which the Database read replica should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public RdbReadReplicaArgs build() {
            $.instanceId = Objects.requireNonNull($.instanceId, "expected parameter 'instanceId' to be non-null");
            return $;
        }
    }

}
