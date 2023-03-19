// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.InstanceServerArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.InstanceServerState;
import io.dirien.scaleway.outputs.InstanceServerPrivateNetwork;
import io.dirien.scaleway.outputs.InstanceServerRootVolume;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Compute Instance servers. For more information, see [the documentation](https://developers.scaleway.com/en/products/instance/api/#servers-8bf7d7).
 * 
 * Please check our [FAQ - Instances](https://www.scaleway.com/en/docs/faq/instances).
 * 
 * ## Examples
 * 
 * ### Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceIp;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
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
 *         var publicIp = new InstanceIp(&#34;publicIp&#34;);
 * 
 *         var web = new InstanceServer(&#34;web&#34;, InstanceServerArgs.builder()        
 *             .type(&#34;DEV1-S&#34;)
 *             .image(&#34;ubuntu_jammy&#34;)
 *             .ipId(publicIp.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### With additional volumes and tags
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceVolume;
 * import com.pulumi.scaleway.InstanceVolumeArgs;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
 * import com.pulumi.scaleway.inputs.InstanceServerRootVolumeArgs;
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
 *         var data = new InstanceVolume(&#34;data&#34;, InstanceVolumeArgs.builder()        
 *             .sizeInGb(100)
 *             .type(&#34;b_ssd&#34;)
 *             .build());
 * 
 *         var web = new InstanceServer(&#34;web&#34;, InstanceServerArgs.builder()        
 *             .type(&#34;DEV1-S&#34;)
 *             .image(&#34;ubuntu_jammy&#34;)
 *             .tags(            
 *                 &#34;hello&#34;,
 *                 &#34;public&#34;)
 *             .rootVolume(InstanceServerRootVolumeArgs.builder()
 *                 .deleteOnTermination(false)
 *                 .build())
 *             .additionalVolumeIds(data.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### With a reserved IP
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceIp;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
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
 *         var ip = new InstanceIp(&#34;ip&#34;);
 * 
 *         var web = new InstanceServer(&#34;web&#34;, InstanceServerArgs.builder()        
 *             .type(&#34;DEV1-S&#34;)
 *             .image(&#34;f974feac-abae-4365-b988-8ec7d1cec10d&#34;)
 *             .tags(            
 *                 &#34;hello&#34;,
 *                 &#34;public&#34;)
 *             .ipId(ip.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### With security group
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceSecurityGroup;
 * import com.pulumi.scaleway.InstanceSecurityGroupArgs;
 * import com.pulumi.scaleway.inputs.InstanceSecurityGroupInboundRuleArgs;
 * import com.pulumi.scaleway.inputs.InstanceSecurityGroupOutboundRuleArgs;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
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
 *         var www = new InstanceSecurityGroup(&#34;www&#34;, InstanceSecurityGroupArgs.builder()        
 *             .inboundDefaultPolicy(&#34;drop&#34;)
 *             .outboundDefaultPolicy(&#34;accept&#34;)
 *             .inboundRules(            
 *                 InstanceSecurityGroupInboundRuleArgs.builder()
 *                     .action(&#34;accept&#34;)
 *                     .port(&#34;22&#34;)
 *                     .ip(&#34;212.47.225.64&#34;)
 *                     .build(),
 *                 InstanceSecurityGroupInboundRuleArgs.builder()
 *                     .action(&#34;accept&#34;)
 *                     .port(&#34;80&#34;)
 *                     .build(),
 *                 InstanceSecurityGroupInboundRuleArgs.builder()
 *                     .action(&#34;accept&#34;)
 *                     .port(&#34;443&#34;)
 *                     .build())
 *             .outboundRules(InstanceSecurityGroupOutboundRuleArgs.builder()
 *                 .action(&#34;drop&#34;)
 *                 .ipRange(&#34;10.20.0.0/24&#34;)
 *                 .build())
 *             .build());
 * 
 *         var web = new InstanceServer(&#34;web&#34;, InstanceServerArgs.builder()        
 *             .type(&#34;DEV1-S&#34;)
 *             .image(&#34;ubuntu_jammy&#34;)
 *             .securityGroupId(www.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### With user data and cloud-init
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
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
 *         var web = new InstanceServer(&#34;web&#34;, InstanceServerArgs.builder()        
 *             .type(&#34;DEV1-S&#34;)
 *             .image(&#34;ubuntu_jammy&#34;)
 *             .userData(Map.ofEntries(
 *                 Map.entry(&#34;foo&#34;, &#34;bar&#34;),
 *                 Map.entry(&#34;cloud-init&#34;, Files.readString(Paths.get(String.format(&#34;%s/cloud-init.yml&#34;, path.module()))))
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### With private network
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.VpcPrivateNetwork;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
 * import com.pulumi.scaleway.inputs.InstanceServerPrivateNetworkArgs;
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
 *         var pn01 = new VpcPrivateNetwork(&#34;pn01&#34;);
 * 
 *         var base = new InstanceServer(&#34;base&#34;, InstanceServerArgs.builder()        
 *             .image(&#34;ubuntu_jammy&#34;)
 *             .type(&#34;DEV1-S&#34;)
 *             .privateNetworks(InstanceServerPrivateNetworkArgs.builder()
 *                 .pnId(pn01.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ### Root volume configuration
 * 
 * #### Resized block volume with installed image
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
 * import com.pulumi.scaleway.inputs.InstanceServerRootVolumeArgs;
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
 *         var image = new InstanceServer(&#34;image&#34;, InstanceServerArgs.builder()        
 *             .image(&#34;ubuntu_jammy&#34;)
 *             .rootVolume(InstanceServerRootVolumeArgs.builder()
 *                 .sizeInGb(100)
 *                 .volumeType(&#34;b_ssd&#34;)
 *                 .build())
 *             .type(&#34;PRO2-XXS&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * #### From snapshot
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ScalewayFunctions;
 * import com.pulumi.scaleway.inputs.GetInstanceSnapshotArgs;
 * import com.pulumi.scaleway.InstanceVolume;
 * import com.pulumi.scaleway.InstanceVolumeArgs;
 * import com.pulumi.scaleway.InstanceServer;
 * import com.pulumi.scaleway.InstanceServerArgs;
 * import com.pulumi.scaleway.inputs.InstanceServerRootVolumeArgs;
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
 *         final var snapshot = ScalewayFunctions.getInstanceSnapshot(GetInstanceSnapshotArgs.builder()
 *             .name(&#34;my_snapshot&#34;)
 *             .build());
 * 
 *         var fromSnapshotInstanceVolume = new InstanceVolume(&#34;fromSnapshotInstanceVolume&#34;, InstanceVolumeArgs.builder()        
 *             .fromSnapshotId(snapshot.applyValue(getInstanceSnapshotResult -&gt; getInstanceSnapshotResult.id()))
 *             .type(&#34;b_ssd&#34;)
 *             .build());
 * 
 *         var fromSnapshotInstanceServer = new InstanceServer(&#34;fromSnapshotInstanceServer&#34;, InstanceServerArgs.builder()        
 *             .type(&#34;PRO2-XXS&#34;)
 *             .rootVolume(InstanceServerRootVolumeArgs.builder()
 *                 .volumeId(fromSnapshotInstanceVolume.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Private Network
 * 
 * &gt; **Important:** Updates to `private_network` will recreate a new private network interface.
 * 
 * - `pn_id` - (Required) The private network ID where to connect.
 * - `mac_address` The private NIC MAC address.
 * - `status` The private NIC state.
 * - `zone` - (Defaults to provider `zone`) The zone in which the server must be created.
 * 
 * &gt; **Important:**
 * 
 * - You can only attach an instance in the same zone as a private network.
 * - Instance supports maximum 8 different private networks.
 * 
 * ## Import
 * 
 * Instance servers can be imported using the `{zone}/{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/instanceServer:InstanceServer web fr-par-1/11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/instanceServer:InstanceServer")
public class InstanceServer extends com.pulumi.resources.CustomResource {
    /**
     * The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
     * attached to the server. Updates to this field will trigger a stop/start of the server.
     * 
     */
    @Export(name="additionalVolumeIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> additionalVolumeIds;

    /**
     * @return The [additional volumes](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39)
     * attached to the server. Updates to this field will trigger a stop/start of the server.
     * 
     */
    public Output<Optional<List<String>>> additionalVolumeIds() {
        return Codegen.optional(this.additionalVolumeIds);
    }
    /**
     * The boot Type of the server. Possible values are: `local`, `bootscript` or `rescue`.
     * 
     */
    @Export(name="bootType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bootType;

    /**
     * @return The boot Type of the server. Possible values are: `local`, `bootscript` or `rescue`.
     * 
     */
    public Output<Optional<String>> bootType() {
        return Codegen.optional(this.bootType);
    }
    /**
     * The ID of the bootscript to use  (set boot_type to `bootscript`).
     * 
     */
    @Export(name="bootscriptId", refs={String.class}, tree="[0]")
    private Output<String> bootscriptId;

    /**
     * @return The ID of the bootscript to use  (set boot_type to `bootscript`).
     * 
     */
    public Output<String> bootscriptId() {
        return this.bootscriptId;
    }
    /**
     * The cloud init script associated with this server
     * 
     */
    @Export(name="cloudInit", refs={String.class}, tree="[0]")
    private Output<String> cloudInit;

    /**
     * @return The cloud init script associated with this server
     * 
     */
    public Output<String> cloudInit() {
        return this.cloudInit;
    }
    /**
     * If true a dynamic IP will be attached to the server.
     * 
     */
    @Export(name="enableDynamicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableDynamicIp;

    /**
     * @return If true a dynamic IP will be attached to the server.
     * 
     */
    public Output<Optional<Boolean>> enableDynamicIp() {
        return Codegen.optional(this.enableDynamicIp);
    }
    /**
     * Determines if IPv6 is enabled for the server.
     * 
     */
    @Export(name="enableIpv6", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableIpv6;

    /**
     * @return Determines if IPv6 is enabled for the server.
     * 
     */
    public Output<Optional<Boolean>> enableIpv6() {
        return Codegen.optional(this.enableIpv6);
    }
    /**
     * The UUID or the label of the base image used by the server. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
     * to find either the right `label` or the right local image `ID` for a given `type`. Optional when creating an instance with an existing root volume.
     * 
     */
    @Export(name="image", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> image;

    /**
     * @return The UUID or the label of the base image used by the server. You can use [this endpoint](https://api-marketplace.scaleway.com/images?page=1&amp;per_page=100)
     * to find either the right `label` or the right local image `ID` for a given `type`. Optional when creating an instance with an existing root volume.
     * 
     */
    public Output<Optional<String>> image() {
        return Codegen.optional(this.image);
    }
    /**
     * = (Optional) The ID of the reserved IP that is attached to the server.
     * 
     */
    @Export(name="ipId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipId;

    /**
     * @return = (Optional) The ID of the reserved IP that is attached to the server.
     * 
     */
    public Output<Optional<String>> ipId() {
        return Codegen.optional(this.ipId);
    }
    /**
     * The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
     * 
     */
    @Export(name="ipv6Address", refs={String.class}, tree="[0]")
    private Output<String> ipv6Address;

    /**
     * @return The default ipv6 address routed to the server. ( Only set when enable_ipv6 is set to true )
     * 
     */
    public Output<String> ipv6Address() {
        return this.ipv6Address;
    }
    /**
     * The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
     * 
     */
    @Export(name="ipv6Gateway", refs={String.class}, tree="[0]")
    private Output<String> ipv6Gateway;

    /**
     * @return The ipv6 gateway address. ( Only set when enable_ipv6 is set to true )
     * 
     */
    public Output<String> ipv6Gateway() {
        return this.ipv6Gateway;
    }
    /**
     * The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
     * 
     */
    @Export(name="ipv6PrefixLength", refs={Integer.class}, tree="[0]")
    private Output<Integer> ipv6PrefixLength;

    /**
     * @return The prefix length of the ipv6 subnet routed to the server. ( Only set when enable_ipv6 is set to true )
     * 
     */
    public Output<Integer> ipv6PrefixLength() {
        return this.ipv6PrefixLength;
    }
    /**
     * The name of the server.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the server.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization ID the server is associated with.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The organization ID the server is associated with.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
     * 
     */
    @Export(name="placementGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> placementGroupId;

    /**
     * @return The [placement group](https://developers.scaleway.com/en/products/instance/api/#placement-groups-d8f653) the server is attached to.
     * 
     */
    public Output<Optional<String>> placementGroupId() {
        return Codegen.optional(this.placementGroupId);
    }
    /**
     * True when the placement group policy is respected.
     * 
     */
    @Export(name="placementGroupPolicyRespected", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> placementGroupPolicyRespected;

    /**
     * @return True when the placement group policy is respected.
     * 
     */
    public Output<Boolean> placementGroupPolicyRespected() {
        return this.placementGroupPolicyRespected;
    }
    /**
     * The Scaleway internal IP address of the server.
     * 
     */
    @Export(name="privateIp", refs={String.class}, tree="[0]")
    private Output<String> privateIp;

    /**
     * @return The Scaleway internal IP address of the server.
     * 
     */
    public Output<String> privateIp() {
        return this.privateIp;
    }
    /**
     * The private network associated with the server.
     * Use the `pn_id` key to attach a [private_network](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea) on your instance.
     * 
     */
    @Export(name="privateNetworks", refs={List.class,InstanceServerPrivateNetwork.class}, tree="[0,1]")
    private Output</* @Nullable */ List<InstanceServerPrivateNetwork>> privateNetworks;

    /**
     * @return The private network associated with the server.
     * Use the `pn_id` key to attach a [private_network](https://developers.scaleway.com/en/products/instance/api/#private-nics-a42eea) on your instance.
     * 
     */
    public Output<Optional<List<InstanceServerPrivateNetwork>>> privateNetworks() {
        return Codegen.optional(this.privateNetworks);
    }
    /**
     * `project_id`) The ID of the project the server is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the server is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * The public IPv4 address of the server.
     * 
     */
    @Export(name="publicIp", refs={String.class}, tree="[0]")
    private Output<String> publicIp;

    /**
     * @return The public IPv4 address of the server.
     * 
     */
    public Output<String> publicIp() {
        return this.publicIp;
    }
    /**
     * Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
     * 
     */
    @Export(name="rootVolume", refs={InstanceServerRootVolume.class}, tree="[0]")
    private Output<InstanceServerRootVolume> rootVolume;

    /**
     * @return Root [volume](https://developers.scaleway.com/en/products/instance/api/#volumes-7e8a39) attached to the server on creation.
     * 
     */
    public Output<InstanceServerRootVolume> rootVolume() {
        return this.rootVolume;
    }
    /**
     * The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return The [security group](https://developers.scaleway.com/en/products/instance/api/#security-groups-8d7f89) the server is attached to.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * The state of the server. Possible values are: `started`, `stopped` or `standby`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state of the server. Possible values are: `started`, `stopped` or `standby`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * The tags associated with the server.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return The tags associated with the server.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The commercial type of the server.
     * You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
     * Updates to this field will recreate a new resource.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The commercial type of the server.
     * You find all the available types on the [pricing page](https://www.scaleway.com/en/pricing/).
     * Updates to this field will recreate a new resource.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The user data associated with the server
     * 
     */
    @Export(name="userData", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> userData;

    /**
     * @return The user data associated with the server
     * 
     */
    public Output<Map<String,String>> userData() {
        return this.userData;
    }
    /**
     * `zone`) The zone in which the server should be created.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return `zone`) The zone in which the server should be created.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceServer(String name) {
        this(name, InstanceServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceServer(String name, InstanceServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceServer(String name, InstanceServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceServer:InstanceServer", name, args == null ? InstanceServerArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceServer(String name, Output<String> id, @Nullable InstanceServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/instanceServer:InstanceServer", name, state, makeResourceOptions(options, id));
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
    public static InstanceServer get(String name, Output<String> id, @Nullable InstanceServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceServer(name, id, state, options);
    }
}
