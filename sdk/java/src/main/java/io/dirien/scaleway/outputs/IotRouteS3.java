// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IotRouteS3 {
    private String bucketName;
    private String bucketRegion;
    private @Nullable String objectPrefix;
    private String strategy;

    private IotRouteS3() {}
    public String bucketName() {
        return this.bucketName;
    }
    public String bucketRegion() {
        return this.bucketRegion;
    }
    public Optional<String> objectPrefix() {
        return Optional.ofNullable(this.objectPrefix);
    }
    public String strategy() {
        return this.strategy;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IotRouteS3 defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bucketName;
        private String bucketRegion;
        private @Nullable String objectPrefix;
        private String strategy;
        public Builder() {}
        public Builder(IotRouteS3 defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucketName = defaults.bucketName;
    	      this.bucketRegion = defaults.bucketRegion;
    	      this.objectPrefix = defaults.objectPrefix;
    	      this.strategy = defaults.strategy;
        }

        @CustomType.Setter
        public Builder bucketName(String bucketName) {
            this.bucketName = Objects.requireNonNull(bucketName);
            return this;
        }
        @CustomType.Setter
        public Builder bucketRegion(String bucketRegion) {
            this.bucketRegion = Objects.requireNonNull(bucketRegion);
            return this;
        }
        @CustomType.Setter
        public Builder objectPrefix(@Nullable String objectPrefix) {
            this.objectPrefix = objectPrefix;
            return this;
        }
        @CustomType.Setter
        public Builder strategy(String strategy) {
            this.strategy = Objects.requireNonNull(strategy);
            return this;
        }
        public IotRouteS3 build() {
            final var o = new IotRouteS3();
            o.bucketName = bucketName;
            o.bucketRegion = bucketRegion;
            o.objectPrefix = objectPrefix;
            o.strategy = strategy;
            return o;
        }
    }
}
