// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.scaleway.outputs.ObjectBucketAclAccessControlPolicyGrantGrantee;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ObjectBucketAclAccessControlPolicyGrant {
    private @Nullable ObjectBucketAclAccessControlPolicyGrantGrantee grantee;
    private String permission;

    private ObjectBucketAclAccessControlPolicyGrant() {}
    public Optional<ObjectBucketAclAccessControlPolicyGrantGrantee> grantee() {
        return Optional.ofNullable(this.grantee);
    }
    public String permission() {
        return this.permission;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ObjectBucketAclAccessControlPolicyGrant defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable ObjectBucketAclAccessControlPolicyGrantGrantee grantee;
        private String permission;
        public Builder() {}
        public Builder(ObjectBucketAclAccessControlPolicyGrant defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.grantee = defaults.grantee;
    	      this.permission = defaults.permission;
        }

        @CustomType.Setter
        public Builder grantee(@Nullable ObjectBucketAclAccessControlPolicyGrantGrantee grantee) {
            this.grantee = grantee;
            return this;
        }
        @CustomType.Setter
        public Builder permission(String permission) {
            this.permission = Objects.requireNonNull(permission);
            return this;
        }
        public ObjectBucketAclAccessControlPolicyGrant build() {
            final var o = new ObjectBucketAclAccessControlPolicyGrant();
            o.grantee = grantee;
            o.permission = permission;
            return o;
        }
    }
}
