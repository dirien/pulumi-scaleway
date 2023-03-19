// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import io.dirien.scaleway.inputs.RdbAclAclRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RdbAclState extends com.pulumi.resources.ResourceArgs {

    public static final RdbAclState Empty = new RdbAclState();

    /**
     * A list of ACLs (structure is described below)
     * 
     */
    @Import(name="aclRules")
    private @Nullable Output<List<RdbAclAclRuleArgs>> aclRules;

    /**
     * @return A list of ACLs (structure is described below)
     * 
     */
    public Optional<Output<List<RdbAclAclRuleArgs>>> aclRules() {
        return Optional.ofNullable(this.aclRules);
    }

    /**
     * UUID of the rdb instance.
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return UUID of the rdb instance.
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * `region`) The region in which the Database Instance should be created.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return `region`) The region in which the Database Instance should be created.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    private RdbAclState() {}

    private RdbAclState(RdbAclState $) {
        this.aclRules = $.aclRules;
        this.instanceId = $.instanceId;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RdbAclState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RdbAclState $;

        public Builder() {
            $ = new RdbAclState();
        }

        public Builder(RdbAclState defaults) {
            $ = new RdbAclState(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclRules A list of ACLs (structure is described below)
         * 
         * @return builder
         * 
         */
        public Builder aclRules(@Nullable Output<List<RdbAclAclRuleArgs>> aclRules) {
            $.aclRules = aclRules;
            return this;
        }

        /**
         * @param aclRules A list of ACLs (structure is described below)
         * 
         * @return builder
         * 
         */
        public Builder aclRules(List<RdbAclAclRuleArgs> aclRules) {
            return aclRules(Output.of(aclRules));
        }

        /**
         * @param aclRules A list of ACLs (structure is described below)
         * 
         * @return builder
         * 
         */
        public Builder aclRules(RdbAclAclRuleArgs... aclRules) {
            return aclRules(List.of(aclRules));
        }

        /**
         * @param instanceId UUID of the rdb instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId UUID of the rdb instance.
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param region `region`) The region in which the Database Instance should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region `region`) The region in which the Database Instance should be created.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public RdbAclState build() {
            return $;
        }
    }

}
