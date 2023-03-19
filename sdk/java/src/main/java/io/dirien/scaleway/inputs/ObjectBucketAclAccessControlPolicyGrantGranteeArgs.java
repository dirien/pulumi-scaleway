// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ObjectBucketAclAccessControlPolicyGrantGranteeArgs extends com.pulumi.resources.ResourceArgs {

    public static final ObjectBucketAclAccessControlPolicyGrantGranteeArgs Empty = new ObjectBucketAclAccessControlPolicyGrantGranteeArgs();

    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * The `region`,`bucket` and `acl` separated by (`/`).
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return The `region`,`bucket` and `acl` separated by (`/`).
     * 
     */
    public Output<String> id() {
        return this.id;
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    private ObjectBucketAclAccessControlPolicyGrantGranteeArgs() {}

    private ObjectBucketAclAccessControlPolicyGrantGranteeArgs(ObjectBucketAclAccessControlPolicyGrantGranteeArgs $) {
        this.displayName = $.displayName;
        this.id = $.id;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ObjectBucketAclAccessControlPolicyGrantGranteeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ObjectBucketAclAccessControlPolicyGrantGranteeArgs $;

        public Builder() {
            $ = new ObjectBucketAclAccessControlPolicyGrantGranteeArgs();
        }

        public Builder(ObjectBucketAclAccessControlPolicyGrantGranteeArgs defaults) {
            $ = new ObjectBucketAclAccessControlPolicyGrantGranteeArgs(Objects.requireNonNull(defaults));
        }

        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param id The `region`,`bucket` and `acl` separated by (`/`).
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The `region`,`bucket` and `acl` separated by (`/`).
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ObjectBucketAclAccessControlPolicyGrantGranteeArgs build() {
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
