// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLbAclsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLbAclsArgs Empty = new GetLbAclsArgs();

    /**
     * The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
     * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
     * 
     */
    @Import(name="frontendId", required=true)
    private Output<String> frontendId;

    /**
     * @return The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
     * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
     * 
     */
    public Output<String> frontendId() {
        return this.frontendId;
    }

    /**
     * The ACL name used as filter. ACLs with a name like it are listed.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The ACL name used as filter. ACLs with a name like it are listed.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * `zone`) The zone in which ACLs exist.
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return `zone`) The zone in which ACLs exist.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private GetLbAclsArgs() {}

    private GetLbAclsArgs(GetLbAclsArgs $) {
        this.frontendId = $.frontendId;
        this.name = $.name;
        this.projectId = $.projectId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLbAclsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLbAclsArgs $;

        public Builder() {
            $ = new GetLbAclsArgs();
        }

        public Builder(GetLbAclsArgs defaults) {
            $ = new GetLbAclsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param frontendId The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
         * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
         * 
         * @return builder
         * 
         */
        public Builder frontendId(Output<String> frontendId) {
            $.frontendId = frontendId;
            return this;
        }

        /**
         * @param frontendId The frontend ID this ACL is attached to. ACLs with a frontend ID like it are listed.
         * &gt; **Important:** LB Frontends&#39; IDs are zoned, which means they are of the form `{zone}/{id}`, e.g. `fr-par-1/11111111-1111-1111-1111-111111111111`
         * 
         * @return builder
         * 
         */
        public Builder frontendId(String frontendId) {
            return frontendId(Output.of(frontendId));
        }

        /**
         * @param name The ACL name used as filter. ACLs with a name like it are listed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The ACL name used as filter. ACLs with a name like it are listed.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param zone `zone`) The zone in which ACLs exist.
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone `zone`) The zone in which ACLs exist.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public GetLbAclsArgs build() {
            $.frontendId = Objects.requireNonNull($.frontendId, "expected parameter 'frontendId' to be non-null");
            return $;
        }
    }

}
