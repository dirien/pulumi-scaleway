// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.scaleway.inputs.ObjectBucketAclAccessControlPolicyGrantGranteeArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectBucketAclAccessControlPolicyGrantArgs extends com.pulumi.resources.ResourceArgs {

    public static final ObjectBucketAclAccessControlPolicyGrantArgs Empty = new ObjectBucketAclAccessControlPolicyGrantArgs();

    @Import(name="grantee")
    private @Nullable Output<ObjectBucketAclAccessControlPolicyGrantGranteeArgs> grantee;

    public Optional<Output<ObjectBucketAclAccessControlPolicyGrantGranteeArgs>> grantee() {
        return Optional.ofNullable(this.grantee);
    }

    @Import(name="permission", required=true)
    private Output<String> permission;

    public Output<String> permission() {
        return this.permission;
    }

    private ObjectBucketAclAccessControlPolicyGrantArgs() {}

    private ObjectBucketAclAccessControlPolicyGrantArgs(ObjectBucketAclAccessControlPolicyGrantArgs $) {
        this.grantee = $.grantee;
        this.permission = $.permission;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectBucketAclAccessControlPolicyGrantArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectBucketAclAccessControlPolicyGrantArgs $;

        public Builder() {
            $ = new ObjectBucketAclAccessControlPolicyGrantArgs();
        }

        public Builder(ObjectBucketAclAccessControlPolicyGrantArgs defaults) {
            $ = new ObjectBucketAclAccessControlPolicyGrantArgs(Objects.requireNonNull(defaults));
        }

        public Builder grantee(@Nullable Output<ObjectBucketAclAccessControlPolicyGrantGranteeArgs> grantee) {
            $.grantee = grantee;
            return this;
        }

        public Builder grantee(ObjectBucketAclAccessControlPolicyGrantGranteeArgs grantee) {
            return grantee(Output.of(grantee));
        }

        public Builder permission(Output<String> permission) {
            $.permission = permission;
            return this;
        }

        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        public ObjectBucketAclAccessControlPolicyGrantArgs build() {
            $.permission = Objects.requireNonNull($.permission, "expected parameter 'permission' to be non-null");
            return $;
        }
    }

}
