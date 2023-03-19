// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamApplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamApplicationArgs Empty = new IamApplicationArgs();

    /**
     * The description of the iam application.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the iam application.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * .The name of the iam application.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return .The name of the iam application.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `organization_id`) The ID of the organization the application is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return `organization_id`) The ID of the organization the application is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    private IamApplicationArgs() {}

    private IamApplicationArgs(IamApplicationArgs $) {
        this.description = $.description;
        this.name = $.name;
        this.organizationId = $.organizationId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamApplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamApplicationArgs $;

        public Builder() {
            $ = new IamApplicationArgs();
        }

        public Builder(IamApplicationArgs defaults) {
            $ = new IamApplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description The description of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name .The name of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name .The name of the iam application.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId `organization_id`) The ID of the organization the application is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId `organization_id`) The ID of the organization the application is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        public IamApplicationArgs build() {
            return $;
        }
    }

}
