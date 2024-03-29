// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRdbInstanceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRdbInstanceArgs Empty = new GetRdbInstanceArgs();

    /**
     * The RDB instance ID.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return The RDB instance ID.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * The name of the RDB instance.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the RDB instance.
     * Only one of `name` and `instance_id` should be specified.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `region`) The region in which the RDB instance exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the RDB instance exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetRdbInstanceArgs() {}

    private GetRdbInstanceArgs(GetRdbInstanceArgs $) {
        this.instanceId = $.instanceId;
        this.name = $.name;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRdbInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRdbInstanceArgs $;

        public Builder() {
            $ = new GetRdbInstanceArgs();
        }

        public Builder(GetRdbInstanceArgs defaults) {
            $ = new GetRdbInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param instanceId The RDB instance ID.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId The RDB instance ID.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param name The name of the RDB instance.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the RDB instance.
         * Only one of `name` and `instance_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region `region`) The region in which the RDB instance exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the RDB instance exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetRdbInstanceArgs build() {
            return $;
        }
    }

}
