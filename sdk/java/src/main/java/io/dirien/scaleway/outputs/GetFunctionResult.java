// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFunctionResult {
    private Integer cpuLimit;
    private Boolean deploy;
    private String description;
    private String domainName;
    private Map<String,String> environmentVariables;
    private @Nullable String functionId;
    private String handler;
    private String httpOption;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Integer maxScale;
    private Integer memoryLimit;
    private Integer minScale;
    private @Nullable String name;
    private String namespaceId;
    private String organizationId;
    private String privacy;
    private String projectId;
    private String region;
    private String runtime;
    private Map<String,String> secretEnvironmentVariables;
    private Integer timeout;
    private String zipFile;
    private String zipHash;

    private GetFunctionResult() {}
    public Integer cpuLimit() {
        return this.cpuLimit;
    }
    public Boolean deploy() {
        return this.deploy;
    }
    public String description() {
        return this.description;
    }
    public String domainName() {
        return this.domainName;
    }
    public Map<String,String> environmentVariables() {
        return this.environmentVariables;
    }
    public Optional<String> functionId() {
        return Optional.ofNullable(this.functionId);
    }
    public String handler() {
        return this.handler;
    }
    public String httpOption() {
        return this.httpOption;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Integer maxScale() {
        return this.maxScale;
    }
    public Integer memoryLimit() {
        return this.memoryLimit;
    }
    public Integer minScale() {
        return this.minScale;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public String namespaceId() {
        return this.namespaceId;
    }
    public String organizationId() {
        return this.organizationId;
    }
    public String privacy() {
        return this.privacy;
    }
    public String projectId() {
        return this.projectId;
    }
    public String region() {
        return this.region;
    }
    public String runtime() {
        return this.runtime;
    }
    public Map<String,String> secretEnvironmentVariables() {
        return this.secretEnvironmentVariables;
    }
    public Integer timeout() {
        return this.timeout;
    }
    public String zipFile() {
        return this.zipFile;
    }
    public String zipHash() {
        return this.zipHash;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFunctionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer cpuLimit;
        private Boolean deploy;
        private String description;
        private String domainName;
        private Map<String,String> environmentVariables;
        private @Nullable String functionId;
        private String handler;
        private String httpOption;
        private String id;
        private Integer maxScale;
        private Integer memoryLimit;
        private Integer minScale;
        private @Nullable String name;
        private String namespaceId;
        private String organizationId;
        private String privacy;
        private String projectId;
        private String region;
        private String runtime;
        private Map<String,String> secretEnvironmentVariables;
        private Integer timeout;
        private String zipFile;
        private String zipHash;
        public Builder() {}
        public Builder(GetFunctionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cpuLimit = defaults.cpuLimit;
    	      this.deploy = defaults.deploy;
    	      this.description = defaults.description;
    	      this.domainName = defaults.domainName;
    	      this.environmentVariables = defaults.environmentVariables;
    	      this.functionId = defaults.functionId;
    	      this.handler = defaults.handler;
    	      this.httpOption = defaults.httpOption;
    	      this.id = defaults.id;
    	      this.maxScale = defaults.maxScale;
    	      this.memoryLimit = defaults.memoryLimit;
    	      this.minScale = defaults.minScale;
    	      this.name = defaults.name;
    	      this.namespaceId = defaults.namespaceId;
    	      this.organizationId = defaults.organizationId;
    	      this.privacy = defaults.privacy;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.runtime = defaults.runtime;
    	      this.secretEnvironmentVariables = defaults.secretEnvironmentVariables;
    	      this.timeout = defaults.timeout;
    	      this.zipFile = defaults.zipFile;
    	      this.zipHash = defaults.zipHash;
        }

        @CustomType.Setter
        public Builder cpuLimit(Integer cpuLimit) {
            this.cpuLimit = Objects.requireNonNull(cpuLimit);
            return this;
        }
        @CustomType.Setter
        public Builder deploy(Boolean deploy) {
            this.deploy = Objects.requireNonNull(deploy);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder domainName(String domainName) {
            this.domainName = Objects.requireNonNull(domainName);
            return this;
        }
        @CustomType.Setter
        public Builder environmentVariables(Map<String,String> environmentVariables) {
            this.environmentVariables = Objects.requireNonNull(environmentVariables);
            return this;
        }
        @CustomType.Setter
        public Builder functionId(@Nullable String functionId) {
            this.functionId = functionId;
            return this;
        }
        @CustomType.Setter
        public Builder handler(String handler) {
            this.handler = Objects.requireNonNull(handler);
            return this;
        }
        @CustomType.Setter
        public Builder httpOption(String httpOption) {
            this.httpOption = Objects.requireNonNull(httpOption);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder maxScale(Integer maxScale) {
            this.maxScale = Objects.requireNonNull(maxScale);
            return this;
        }
        @CustomType.Setter
        public Builder memoryLimit(Integer memoryLimit) {
            this.memoryLimit = Objects.requireNonNull(memoryLimit);
            return this;
        }
        @CustomType.Setter
        public Builder minScale(Integer minScale) {
            this.minScale = Objects.requireNonNull(minScale);
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder namespaceId(String namespaceId) {
            this.namespaceId = Objects.requireNonNull(namespaceId);
            return this;
        }
        @CustomType.Setter
        public Builder organizationId(String organizationId) {
            this.organizationId = Objects.requireNonNull(organizationId);
            return this;
        }
        @CustomType.Setter
        public Builder privacy(String privacy) {
            this.privacy = Objects.requireNonNull(privacy);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        @CustomType.Setter
        public Builder runtime(String runtime) {
            this.runtime = Objects.requireNonNull(runtime);
            return this;
        }
        @CustomType.Setter
        public Builder secretEnvironmentVariables(Map<String,String> secretEnvironmentVariables) {
            this.secretEnvironmentVariables = Objects.requireNonNull(secretEnvironmentVariables);
            return this;
        }
        @CustomType.Setter
        public Builder timeout(Integer timeout) {
            this.timeout = Objects.requireNonNull(timeout);
            return this;
        }
        @CustomType.Setter
        public Builder zipFile(String zipFile) {
            this.zipFile = Objects.requireNonNull(zipFile);
            return this;
        }
        @CustomType.Setter
        public Builder zipHash(String zipHash) {
            this.zipHash = Objects.requireNonNull(zipHash);
            return this;
        }
        public GetFunctionResult build() {
            final var o = new GetFunctionResult();
            o.cpuLimit = cpuLimit;
            o.deploy = deploy;
            o.description = description;
            o.domainName = domainName;
            o.environmentVariables = environmentVariables;
            o.functionId = functionId;
            o.handler = handler;
            o.httpOption = httpOption;
            o.id = id;
            o.maxScale = maxScale;
            o.memoryLimit = memoryLimit;
            o.minScale = minScale;
            o.name = name;
            o.namespaceId = namespaceId;
            o.organizationId = organizationId;
            o.privacy = privacy;
            o.projectId = projectId;
            o.region = region;
            o.runtime = runtime;
            o.secretEnvironmentVariables = secretEnvironmentVariables;
            o.timeout = timeout;
            o.zipFile = zipFile;
            o.zipHash = zipHash;
            return o;
        }
    }
}
