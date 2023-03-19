// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.VpcPublicGatewayIpArgs;
import io.dirien.scaleway.inputs.VpcPublicGatewayIpState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway VPC Public Gateway IP.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/vpc-gw/api/v1/#ips-268151).
 * 
 * ## Example
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.VpcPublicGatewayIp;
 * import com.pulumi.scaleway.VpcPublicGatewayIpArgs;
 * import com.pulumi.scaleway.DomainRecord;
 * import com.pulumi.scaleway.DomainRecordArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var main = new VpcPublicGatewayIp(&#34;main&#34;, VpcPublicGatewayIpArgs.builder()        
 *             .reverse(&#34;tf.example.com&#34;)
 *             .build());
 * 
 *         var tfA = new DomainRecord(&#34;tfA&#34;, DomainRecordArgs.builder()        
 *             .data(main.address())
 *             .dnsZone(&#34;example.com&#34;)
 *             .priority(1)
 *             .ttl(3600)
 *             .type(&#34;A&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Public gateway can be imported using the `{zone}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp main fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp")
public class VpcPublicGatewayIp extends com.pulumi.resources.CustomResource {
    /**
     * The IP address itself.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return The IP address itself.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * The date and time of the creation of the public gateway ip.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date and time of the creation of the public gateway ip.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The organization ID the public gateway ip is associated with.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The organization ID the public gateway ip is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * `project_id`) The ID of the project the public gateway ip is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the public gateway ip is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The reverse domain name for the IP address
     * 
     */
    @Export(name="reverse", refs={String.class}, tree="[0]")
    private Output<String> reverse;

    /**
     * @return The reverse domain name for the IP address
     * 
     */
    public Output<String> reverse() {
        return this.reverse;
    }
    /**
     * The tags associated with the public gateway IP.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The tags associated with the public gateway IP.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The date and time of the last update of the public gateway ip.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The date and time of the last update of the public gateway ip.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * `zone`) The zone in which the public gateway ip should be created.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the public gateway ip should be created.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcPublicGatewayIp(String name) {
        this(name, VpcPublicGatewayIpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcPublicGatewayIp(String name, @Nullable VpcPublicGatewayIpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcPublicGatewayIp(String name, @Nullable VpcPublicGatewayIpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, args == null ? VpcPublicGatewayIpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcPublicGatewayIp(String name, Output<String> id, @Nullable VpcPublicGatewayIpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/vpcPublicGatewayIp:VpcPublicGatewayIp", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static VpcPublicGatewayIp get(String name, Output<String> id, @Nullable VpcPublicGatewayIpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcPublicGatewayIp(name, id, state, options);
    }
}
