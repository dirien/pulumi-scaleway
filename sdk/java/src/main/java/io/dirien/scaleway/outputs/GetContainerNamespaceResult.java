// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetContainerNamespaceResult {
    /**
     * @return The description of the namespace.
     * 
     */
    private String description;
    private Boolean destroyRegistry;
    /**
     * @return The environment variables of the namespace.
     * 
     */
    private Map<String,String> environmentVariables;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private @Nullable String name;
    private @Nullable String namespaceId;
    /**
     * @return The organization ID the namespace is associated with.
     * 
     */
    private String organizationId;
    private String projectId;
    private @Nullable String region;
    /**
     * @return The registry endpoint of the namespace.
     * 
     */
    private String registryEndpoint;
    /**
     * @return The registry namespace ID of the namespace.
     * 
     */
    private String registryNamespaceId;
    private Map<String,String> secretEnvironmentVariables;

    private GetContainerNamespaceResult() {}
    /**
     * @return The description of the namespace.
     * 
     */
    public String description() {
        return this.description;
    }
    public Boolean destroyRegistry() {
        return this.destroyRegistry;
    }
    /**
     * @return The environment variables of the namespace.
     * 
     */
    public Map<String,String> environmentVariables() {
        return this.environmentVariables;
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
    public Optional<String> namespaceId() {
        return Optional.ofNullable(this.namespaceId);
    }
    /**
     * @return The organization ID the namespace is associated with.
     * 
     */
    public String organizationId() {
        return this.organizationId;
    }
    public String projectId() {
        return this.projectId;
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    /**
     * @return The registry endpoint of the namespace.
     * 
     */
    public String registryEndpoint() {
        return this.registryEndpoint;
    }
    /**
     * @return The registry namespace ID of the namespace.
     * 
     */
    public String registryNamespaceId() {
        return this.registryNamespaceId;
    }
    public Map<String,String> secretEnvironmentVariables() {
        return this.secretEnvironmentVariables;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerNamespaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Boolean destroyRegistry;
        private Map<String,String> environmentVariables;
        private String id;
        private @Nullable String name;
        private @Nullable String namespaceId;
        private String organizationId;
        private String projectId;
        private @Nullable String region;
        private String registryEndpoint;
        private String registryNamespaceId;
        private Map<String,String> secretEnvironmentVariables;
        public Builder() {}
        public Builder(GetContainerNamespaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.destroyRegistry = defaults.destroyRegistry;
    	      this.environmentVariables = defaults.environmentVariables;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.namespaceId = defaults.namespaceId;
    	      this.organizationId = defaults.organizationId;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.registryEndpoint = defaults.registryEndpoint;
    	      this.registryNamespaceId = defaults.registryNamespaceId;
    	      this.secretEnvironmentVariables = defaults.secretEnvironmentVariables;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder destroyRegistry(Boolean destroyRegistry) {
            this.destroyRegistry = Objects.requireNonNull(destroyRegistry);
            return this;
        }
        @CustomType.Setter
        public Builder environmentVariables(Map<String,String> environmentVariables) {
            this.environmentVariables = Objects.requireNonNull(environmentVariables);
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
        public Builder namespaceId(@Nullable String namespaceId) {
            this.namespaceId = namespaceId;
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder registryEndpoint(String registryEndpoint) {
            this.registryEndpoint = Objects.requireNonNull(registryEndpoint);
            return this;
        }
        @CustomType.Setter
        public Builder registryNamespaceId(String registryNamespaceId) {
            this.registryNamespaceId = Objects.requireNonNull(registryNamespaceId);
            return this;
        }
        @CustomType.Setter
        public Builder secretEnvironmentVariables(Map<String,String> secretEnvironmentVariables) {
            this.secretEnvironmentVariables = Objects.requireNonNull(secretEnvironmentVariables);
            return this;
        }
        public GetContainerNamespaceResult build() {
            final var o = new GetContainerNamespaceResult();
            o.description = description;
            o.destroyRegistry = destroyRegistry;
            o.environmentVariables = environmentVariables;
            o.id = id;
            o.name = name;
            o.namespaceId = namespaceId;
            o.organizationId = organizationId;
            o.projectId = projectId;
            o.region = region;
            o.registryEndpoint = registryEndpoint;
            o.registryNamespaceId = registryNamespaceId;
            o.secretEnvironmentVariables = secretEnvironmentVariables;
            return o;
        }
    }
}
