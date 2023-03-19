// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.outputs;

import com.pulumi.core.annotations.CustomType;
import io.dirien.scaleway.outputs.MnqCredentialSqsSnsCredentialsPermissions;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class MnqCredentialSqsSnsCredentials {
    /**
     * @return The ID of the key.
     * 
     */
    private @Nullable String accessKey;
    /**
     * @return List of permissions associated to this Credential. Only one of permissions may be set.
     * 
     */
    private @Nullable MnqCredentialSqsSnsCredentialsPermissions permissions;
    /**
     * @return The Secret value of the key.
     * 
     */
    private @Nullable String secretKey;

    private MnqCredentialSqsSnsCredentials() {}
    /**
     * @return The ID of the key.
     * 
     */
    public Optional<String> accessKey() {
        return Optional.ofNullable(this.accessKey);
    }
    /**
     * @return List of permissions associated to this Credential. Only one of permissions may be set.
     * 
     */
    public Optional<MnqCredentialSqsSnsCredentialsPermissions> permissions() {
        return Optional.ofNullable(this.permissions);
    }
    /**
     * @return The Secret value of the key.
     * 
     */
    public Optional<String> secretKey() {
        return Optional.ofNullable(this.secretKey);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MnqCredentialSqsSnsCredentials defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String accessKey;
        private @Nullable MnqCredentialSqsSnsCredentialsPermissions permissions;
        private @Nullable String secretKey;
        public Builder() {}
        public Builder(MnqCredentialSqsSnsCredentials defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessKey = defaults.accessKey;
    	      this.permissions = defaults.permissions;
    	      this.secretKey = defaults.secretKey;
        }

        @CustomType.Setter
        public Builder accessKey(@Nullable String accessKey) {
            this.accessKey = accessKey;
            return this;
        }
        @CustomType.Setter
        public Builder permissions(@Nullable MnqCredentialSqsSnsCredentialsPermissions permissions) {
            this.permissions = permissions;
            return this;
        }
        @CustomType.Setter
        public Builder secretKey(@Nullable String secretKey) {
            this.secretKey = secretKey;
            return this;
        }
        public MnqCredentialSqsSnsCredentials build() {
            final var o = new MnqCredentialSqsSnsCredentials();
            o.accessKey = accessKey;
            o.permissions = permissions;
            o.secretKey = secretKey;
            return o;
        }
    }
}
