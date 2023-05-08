// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.CockpitTokenArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.CockpitTokenState;
import io.dirien.scaleway.outputs.CockpitTokenScopes;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Cockpit Tokens.
 * 
 * For more information consult the [documentation](https://www.scaleway.com/en/docs/observability/cockpit/concepts/#tokens).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ScalewayFunctions;
 * import com.pulumi.scaleway.inputs.GetCockpitArgs;
 * import com.pulumi.scaleway.CockpitToken;
 * import com.pulumi.scaleway.CockpitTokenArgs;
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
 *         final var mainCockpit = ScalewayFunctions.getCockpit();
 * 
 *         var mainCockpitToken = new CockpitToken(&#34;mainCockpitToken&#34;, CockpitTokenArgs.builder()        
 *             .projectId(mainCockpit.applyValue(getCockpitResult -&gt; getCockpitResult.projectId()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.ScalewayFunctions;
 * import com.pulumi.scaleway.inputs.GetCockpitArgs;
 * import com.pulumi.scaleway.CockpitToken;
 * import com.pulumi.scaleway.CockpitTokenArgs;
 * import com.pulumi.scaleway.inputs.CockpitTokenScopesArgs;
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
 *         final var mainCockpit = ScalewayFunctions.getCockpit();
 * 
 *         var mainCockpitToken = new CockpitToken(&#34;mainCockpitToken&#34;, CockpitTokenArgs.builder()        
 *             .projectId(mainCockpit.applyValue(getCockpitResult -&gt; getCockpitResult.projectId()))
 *             .scopes(CockpitTokenScopesArgs.builder()
 *                 .queryMetrics(true)
 *                 .writeMetrics(false)
 *                 .queryLogs(true)
 *                 .writeLogs(false)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cockpits can be imported using the token ID, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/cockpitToken:CockpitToken main 11111111-1111-1111-1111-111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/cockpitToken:CockpitToken")
public class CockpitToken extends com.pulumi.resources.CustomResource {
    /**
     * The name of the token
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the token
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * `project_id`) The ID of the project the cockpit is associated with.
     * 
     */
    @Export(name="projectId", refs={String.class}, tree="[0]")
    private Output<String> projectId;

    /**
     * @return `project_id`) The ID of the project the cockpit is associated with.
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }
    /**
     * Allowed scopes
     * 
     */
    @Export(name="scopes", refs={CockpitTokenScopes.class}, tree="[0]")
    private Output<CockpitTokenScopes> scopes;

    /**
     * @return Allowed scopes
     * 
     */
    public Output<CockpitTokenScopes> scopes() {
        return this.scopes;
    }
    /**
     * The secret key of the token
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output<String> secretKey;

    /**
     * @return The secret key of the token
     * 
     */
    public Output<String> secretKey() {
        return this.secretKey;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CockpitToken(String name) {
        this(name, CockpitTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CockpitToken(String name, @Nullable CockpitTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CockpitToken(String name, @Nullable CockpitTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/cockpitToken:CockpitToken", name, args == null ? CockpitTokenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CockpitToken(String name, Output<String> id, @Nullable CockpitTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/cockpitToken:CockpitToken", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretKey"
            ))
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
    public static CockpitToken get(String name, Output<String> id, @Nullable CockpitTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CockpitToken(name, id, state, options);
    }
}
