// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIamGroupResult {
    private List<String> applicationIds;
    private String createdAt;
    private String description;
    private @Nullable String groupId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String name;
    private @Nullable String organizationId;
    private String updatedAt;
    private List<String> userIds;

    private GetIamGroupResult() {}
    public List<String> applicationIds() {
        return this.applicationIds;
    }
    public String createdAt() {
        return this.createdAt;
    }
    public String description() {
        return this.description;
    }
    public Optional<String> groupId() {
        return Optional.ofNullable(this.groupId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Optional<String> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }
    public String updatedAt() {
        return this.updatedAt;
    }
    public List<String> userIds() {
        return this.userIds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIamGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> applicationIds;
        private String createdAt;
        private String description;
        private @Nullable String groupId;
        private String id;
        private @Nullable String name;
        private @Nullable String organizationId;
        private String updatedAt;
        private List<String> userIds;
        public Builder() {}
        public Builder(GetIamGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.applicationIds = defaults.applicationIds;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.groupId = defaults.groupId;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.organizationId = defaults.organizationId;
    	      this.updatedAt = defaults.updatedAt;
    	      this.userIds = defaults.userIds;
        }

        @CustomType.Setter
        public Builder applicationIds(List<String> applicationIds) {
            this.applicationIds = Objects.requireNonNull(applicationIds);
            return this;
        }
        public Builder applicationIds(String... applicationIds) {
            return applicationIds(List.of(applicationIds));
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder groupId(@Nullable String groupId) {
            this.groupId = groupId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(@Nullable String organizationId) {
            this.organizationId = organizationId;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder userIds(List<String> userIds) {
            this.userIds = Objects.requireNonNull(userIds);
            return this;
        }
        public Builder userIds(String... userIds) {
            return userIds(List.of(userIds));
        }
        public GetIamGroupResult build() {
            final var o = new GetIamGroupResult();
            o.applicationIds = applicationIds;
            o.createdAt = createdAt;
            o.description = description;
            o.groupId = groupId;
            o.id = id;
            o.name = name;
            o.organizationId = organizationId;
            o.updatedAt = updatedAt;
            o.userIds = userIds;
            return o;
        }
    }
}
