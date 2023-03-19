// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamGroupArgs Empty = new IamGroupArgs();

    /**
     * The list of IDs of the applications attached to the group.
     * 
     */
    @Import(name="applicationIds")
    private @Nullable Output<List<String>> applicationIds;

    /**
     * @return The list of IDs of the applications attached to the group.
     * 
     */
    public Optional<Output<List<String>>> applicationIds() {
        return Optional.ofNullable(this.applicationIds);
    }

    /**
     * The description of the IAM group.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the IAM group.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the IAM group.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the IAM group.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `organization_id`) The ID of the organization the group is associated with.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return `organization_id`) The ID of the organization the group is associated with.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * The list of IDs of the users attached to the group.
     * 
     */
    @Import(name="userIds")
    private @Nullable Output<List<String>> userIds;

    /**
     * @return The list of IDs of the users attached to the group.
     * 
     */
    public Optional<Output<List<String>>> userIds() {
        return Optional.ofNullable(this.userIds);
    }

    private IamGroupArgs() {}

    private IamGroupArgs(IamGroupArgs $) {
        this.applicationIds = $.applicationIds;
        this.description = $.description;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.userIds = $.userIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamGroupArgs $;

        public Builder() {
            $ = new IamGroupArgs();
        }

        public Builder(IamGroupArgs defaults) {
            $ = new IamGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationIds The list of IDs of the applications attached to the group.
         * 
         * @return builder
         * 
         */
        public Builder applicationIds(@Nullable Output<List<String>> applicationIds) {
            $.applicationIds = applicationIds;
            return this;
        }

        /**
         * @param applicationIds The list of IDs of the applications attached to the group.
         * 
         * @return builder
         * 
         */
        public Builder applicationIds(List<String> applicationIds) {
            return applicationIds(Output.of(applicationIds));
        }

        /**
         * @param applicationIds The list of IDs of the applications attached to the group.
         * 
         * @return builder
         * 
         */
        public Builder applicationIds(String... applicationIds) {
            return applicationIds(List.of(applicationIds));
        }

        /**
         * @param description The description of the IAM group.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the IAM group.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the IAM group.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the IAM group.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId `organization_id`) The ID of the organization the group is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId `organization_id`) The ID of the organization the group is associated with.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param userIds The list of IDs of the users attached to the group.
         * 
         * @return builder
         * 
         */
        public Builder userIds(@Nullable Output<List<String>> userIds) {
            $.userIds = userIds;
            return this;
        }

        /**
         * @param userIds The list of IDs of the users attached to the group.
         * 
         * @return builder
         * 
         */
        public Builder userIds(List<String> userIds) {
            return userIds(Output.of(userIds));
        }

        /**
         * @param userIds The list of IDs of the users attached to the group.
         * 
         * @return builder
         * 
         */
        public Builder userIds(String... userIds) {
            return userIds(List.of(userIds));
        }

        public IamGroupArgs build() {
            return $;
        }
    }

}
