// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetContainerNamespaceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerNamespaceArgs Empty = new GetContainerNamespaceArgs();

    /**
     * The namespace name.
     * Only one of `name` and `namespace_id` should be specified.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The namespace name.
     * Only one of `name` and `namespace_id` should be specified.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace id.
     * Only one of `name` and `namespace_id` should be specified.
     * 
     */
    @Import(name="namespaceId")
    private @Nullable Output<String> namespaceId;

    /**
     * @return The namespace id.
     * Only one of `name` and `namespace_id` should be specified.
     * 
     */
    public Optional<Output<String>> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }

    /**
     * `region`) The region in which the namespace exists.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the namespace exists.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private GetContainerNamespaceArgs() {}

    private GetContainerNamespaceArgs(GetContainerNamespaceArgs $) {
        this.name = $.name;
        this.namespaceId = $.namespaceId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerNamespaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerNamespaceArgs $;

        public Builder() {
            $ = new GetContainerNamespaceArgs();
        }

        public Builder(GetContainerNamespaceArgs defaults) {
            $ = new GetContainerNamespaceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The namespace name.
         * Only one of `name` and `namespace_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The namespace name.
         * Only one of `name` and `namespace_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespaceId The namespace id.
         * Only one of `name` and `namespace_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(@Nullable Output<String> namespaceId) {
            $.namespaceId = namespaceId;
            return this;
        }

        /**
         * @param namespaceId The namespace id.
         * Only one of `name` and `namespace_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder namespaceId(String namespaceId) {
            return namespaceId(Output.of(namespaceId));
        }

        /**
         * @param region `region`) The region in which the namespace exists.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the namespace exists.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public GetContainerNamespaceArgs build() {
            return $;
        }
    }

}
