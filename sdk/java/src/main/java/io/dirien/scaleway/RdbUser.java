// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.RdbUserArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.RdbUserState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway Database Users.
 * For more information, see [the documentation](https://developers.scaleway.com/en/products/rdb/api).
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
 * import com.pulumi.random.RandomPassword;
 * import com.pulumi.random.RandomPasswordArgs;
 * import com.pulumi.scaleway.RdbUser;
 * import com.pulumi.scaleway.RdbUserArgs;
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
 *         var dbPassword = new RandomPassword(&#34;dbPassword&#34;, RandomPasswordArgs.builder()        
 *             .length(16)
 *             .special(true)
 *             .build());
 * 
 *         var dbAdmin = new RdbUser(&#34;dbAdmin&#34;, RdbUserArgs.builder()        
 *             .instanceId(scaleway_rdb_instance.main().id())
 *             .password(dbPassword.result())
 *             .isAdmin(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Database User can be imported using `{region}/{instance_id}/{user_name}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/rdbUser:RdbUser admin fr-par/11111111-1111-1111-1111-111111111111/admin
 * ```
 * 
 */
@ResourceType(type="scaleway:index/rdbUser:RdbUser")
public class RdbUser extends com.pulumi.resources.CustomResource {
    /**
     * UUID of the rdb instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return UUID of the rdb instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Grant admin permissions to the Database User.
     * 
     */
    @Export(name="isAdmin", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isAdmin;

    /**
     * @return Grant admin permissions to the Database User.
     * 
     */
    public Output<Optional<Boolean>> isAdmin() {
        return Codegen.optional(this.isAdmin);
    }
    /**
     * Database User name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Database User name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Database User password.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return Database User password.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The Scaleway region this resource resides in.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The Scaleway region this resource resides in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RdbUser(String name) {
        this(name, RdbUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RdbUser(String name, RdbUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RdbUser(String name, RdbUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/rdbUser:RdbUser", name, args == null ? RdbUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RdbUser(String name, Output<String> id, @Nullable RdbUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/rdbUser:RdbUser", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static RdbUser get(String name, Output<String> id, @Nullable RdbUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RdbUser(name, id, state, options);
    }
}
