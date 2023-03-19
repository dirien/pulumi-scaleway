// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.scaleway.outputs.GetObjectBucketCorsRule;
import io.dirien.scaleway.outputs.GetObjectBucketLifecycleRule;
import io.dirien.scaleway.outputs.GetObjectBucketVersioning;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetObjectBucketResult {
    private String acl;
    private List<GetObjectBucketCorsRule> corsRules;
    /**
     * @return The endpoint URL of the bucket
     * 
     */
    private String endpoint;
    private Boolean forceDestroy;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private List<GetObjectBucketLifecycleRule> lifecycleRules;
    private @Nullable String name;
    private Boolean objectLockEnabled;
    private @Nullable String projectId;
    private @Nullable String region;
    private Map<String,String> tags;
    private List<GetObjectBucketVersioning> versionings;

    private GetObjectBucketResult() {}
    public String acl() {
        return this.acl;
    }
    public List<GetObjectBucketCorsRule> corsRules() {
        return this.corsRules;
    }
    /**
     * @return The endpoint URL of the bucket
     * 
     */
    public String endpoint() {
        return this.endpoint;
    }
    public Boolean forceDestroy() {
        return this.forceDestroy;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public List<GetObjectBucketLifecycleRule> lifecycleRules() {
        return this.lifecycleRules;
    }
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    public Boolean objectLockEnabled() {
        return this.objectLockEnabled;
    }
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public Map<String,String> tags() {
        return this.tags;
    }
    public List<GetObjectBucketVersioning> versionings() {
        return this.versionings;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetObjectBucketResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String acl;
        private List<GetObjectBucketCorsRule> corsRules;
        private String endpoint;
        private Boolean forceDestroy;
        private String id;
        private List<GetObjectBucketLifecycleRule> lifecycleRules;
        private @Nullable String name;
        private Boolean objectLockEnabled;
        private @Nullable String projectId;
        private @Nullable String region;
        private Map<String,String> tags;
        private List<GetObjectBucketVersioning> versionings;
        public Builder() {}
        public Builder(GetObjectBucketResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.acl = defaults.acl;
    	      this.corsRules = defaults.corsRules;
    	      this.endpoint = defaults.endpoint;
    	      this.forceDestroy = defaults.forceDestroy;
    	      this.id = defaults.id;
    	      this.lifecycleRules = defaults.lifecycleRules;
    	      this.name = defaults.name;
    	      this.objectLockEnabled = defaults.objectLockEnabled;
    	      this.projectId = defaults.projectId;
    	      this.region = defaults.region;
    	      this.tags = defaults.tags;
    	      this.versionings = defaults.versionings;
        }

        @CustomType.Setter
        public Builder acl(String acl) {
            this.acl = Objects.requireNonNull(acl);
            return this;
        }
        @CustomType.Setter
        public Builder corsRules(List<GetObjectBucketCorsRule> corsRules) {
            this.corsRules = Objects.requireNonNull(corsRules);
            return this;
        }
        public Builder corsRules(GetObjectBucketCorsRule... corsRules) {
            return corsRules(List.of(corsRules));
        }
        @CustomType.Setter
        public Builder endpoint(String endpoint) {
            this.endpoint = Objects.requireNonNull(endpoint);
            return this;
        }
        @CustomType.Setter
        public Builder forceDestroy(Boolean forceDestroy) {
            this.forceDestroy = Objects.requireNonNull(forceDestroy);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lifecycleRules(List<GetObjectBucketLifecycleRule> lifecycleRules) {
            this.lifecycleRules = Objects.requireNonNull(lifecycleRules);
            return this;
        }
        public Builder lifecycleRules(GetObjectBucketLifecycleRule... lifecycleRules) {
            return lifecycleRules(List.of(lifecycleRules));
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder objectLockEnabled(Boolean objectLockEnabled) {
            this.objectLockEnabled = Objects.requireNonNull(objectLockEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        @CustomType.Setter
        public Builder versionings(List<GetObjectBucketVersioning> versionings) {
            this.versionings = Objects.requireNonNull(versionings);
            return this;
        }
        public Builder versionings(GetObjectBucketVersioning... versionings) {
            return versionings(List.of(versionings));
        }
        public GetObjectBucketResult build() {
            final var o = new GetObjectBucketResult();
            o.acl = acl;
            o.corsRules = corsRules;
            o.endpoint = endpoint;
            o.forceDestroy = forceDestroy;
            o.id = id;
            o.lifecycleRules = lifecycleRules;
            o.name = name;
            o.objectLockEnabled = objectLockEnabled;
            o.projectId = projectId;
            o.region = region;
            o.tags = tags;
            o.versionings = versionings;
            return o;
        }
    }
}
