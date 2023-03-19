// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRedisClusterArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRedisClusterArgs Empty = new GetRedisClusterArgs();

    /**
     * The Redis cluster ID.
     * Only one of the `name` and `cluster_id` should be specified.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return The Redis cluster ID.
     * Only one of the `name` and `cluster_id` should be specified.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * The name of the Redis cluster.
     * Only one of the `name` and `cluster_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the Redis cluster.
     * Only one of the `name` and `cluster_id` should be specified.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The zone in which the server exists.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `region`) The zone in which the server exists.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetRedisClusterArgs() {}

    private GetRedisClusterArgs(GetRedisClusterArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRedisClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRedisClusterArgs $;

        public Builder() {
            $ = new GetRedisClusterArgs();
        }

        public Builder(GetRedisClusterArgs defaults) {
            $ = new GetRedisClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The Redis cluster ID.
         * Only one of the `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId The Redis cluster ID.
         * Only one of the `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param name The name of the Redis cluster.
         * Only one of the `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Redis cluster.
         * Only one of the `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param zone `region`) The zone in which the server exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `region`) The zone in which the server exists.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetRedisClusterArgs build() {
            return $;
        }
    }

}
