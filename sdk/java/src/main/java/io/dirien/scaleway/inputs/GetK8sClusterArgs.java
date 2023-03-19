// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetK8sClusterArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetK8sClusterArgs Empty = new GetK8sClusterArgs();

    /**
     * The cluster ID. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return The cluster ID. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * The cluster name. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The cluster name. Only one of `name` and `cluster_id` should be specified.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The region in which the cluster exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the cluster exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetK8sClusterArgs() {}

    private GetK8sClusterArgs(GetK8sClusterArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetK8sClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetK8sClusterArgs $;

        public Builder() {
            $ = new GetK8sClusterArgs();
        }

        public Builder(GetK8sClusterArgs defaults) {
            $ = new GetK8sClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId The cluster ID. Only one of `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId The cluster ID. Only one of `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param name The cluster name. Only one of `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The cluster name. Only one of `name` and `cluster_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region `region`) The region in which the cluster exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the cluster exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetK8sClusterArgs build() {
            return $;
        }
    }

}
