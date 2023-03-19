// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetK8sClusterPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetK8sClusterPlainArgs Empty = new GetK8sClusterPlainArgs();

    /**
     * The cluster ID. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    @Import(name="clusterId")
    private @Nullable String clusterId;

    /**
     * @return The cluster ID. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    public Optional<String> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * The cluster name. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The cluster name. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The region in which the cluster exists.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return `region`) The region in which the cluster exists.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetK8sClusterPlainArgs() {}

    private GetK8sClusterPlainArgs(GetK8sClusterPlainArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetK8sClusterPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetK8sClusterPlainArgs $;

        public Builder() {
            $ = new GetK8sClusterPlainArgs();
        }

        public Builder(GetK8sClusterPlainArgs defaults) {
            $ = new GetK8sClusterPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The cluster ID. Only one of `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param name The cluster name. Only one of `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region `region`) The region in which the cluster exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetK8sClusterPlainArgs build() {
            return $;
        }
    }

}
