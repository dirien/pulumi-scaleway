// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package io.dirien.scaleway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import io.dirien.scaleway.IamApiKeyArgs;
import io.dirien.scaleway.Utilities;
import io.dirien.scaleway.inputs.IamApiKeyState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates and manages Scaleway IAM API Keys. For more information, please
 * check [the documentation](https://developers.scaleway.com/en/products/iam/api/v1alpha1/#api-keys-3665ae)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.scaleway.IamApplication;
 * import com.pulumi.scaleway.IamApiKey;
 * import com.pulumi.scaleway.IamApiKeyArgs;
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
 *         var ciCd = new IamApplication(&#34;ciCd&#34;);
 * 
 *         var main = new IamApiKey(&#34;main&#34;, IamApiKeyArgs.builder()        
 *             .applicationId(scaleway_iam_application.main().id())
 *             .description(&#34;a description&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Api keys can be imported using the `{id}`, e.g. bash
 * 
 * ```sh
 *  $ pulumi import scaleway:index/iamApiKey:IamApiKey main 11111111111111111111
 * ```
 * 
 */
@ResourceType(type="scaleway:index/iamApiKey:IamApiKey")
public class IamApiKey extends com.pulumi.resources.CustomResource {
    /**
     * The access key of the iam api key.
     * 
     */
    @Export(name="accessKey", refs={String.class}, tree="[0]")
    private Output<String> accessKey;

    /**
     * @return The access key of the iam api key.
     * 
     */
    public Output<String> accessKey() {
        return this.accessKey;
    }
    /**
     * ID of the application attached to the api key.
     * Only one of the `application_id` and `user_id` should be specified.
     * 
     */
    @Export(name="applicationId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> applicationId;

    /**
     * @return ID of the application attached to the api key.
     * Only one of the `application_id` and `user_id` should be specified.
     * 
     */
    public Output<Optional<String>> applicationId() {
        return Codegen.optional(this.applicationId);
    }
    /**
     * The date and time of the creation of the iam api key.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date and time of the creation of the iam api key.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The IP Address of the device which created the API key.
     * 
     */
    @Export(name="creationIp", refs={String.class}, tree="[0]")
    private Output<String> creationIp;

    /**
     * @return The IP Address of the device which created the API key.
     * 
     */
    public Output<String> creationIp() {
        return this.creationIp;
    }
    /**
     * The default project ID to use with object storage.
     * 
     */
    @Export(name="defaultProjectId", refs={String.class}, tree="[0]")
    private Output<String> defaultProjectId;

    /**
     * @return The default project ID to use with object storage.
     * 
     */
    public Output<String> defaultProjectId() {
        return this.defaultProjectId;
    }
    /**
     * The description of the iam api key.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the iam api key.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether the iam api key is editable.
     * 
     */
    @Export(name="editable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> editable;

    /**
     * @return Whether the iam api key is editable.
     * 
     */
    public Output<Boolean> editable() {
        return this.editable;
    }
    /**
     * The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     * 
     */
    @Export(name="expiresAt", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> expiresAt;

    /**
     * @return The date and time of the expiration of the iam api key. Please note that in case of change,
     * the resource will be recreated.
     * 
     */
    public Output<Optional<String>> expiresAt() {
        return Codegen.optional(this.expiresAt);
    }
    /**
     * The secret Key of the iam api key.
     * 
     */
    @Export(name="secretKey", refs={String.class}, tree="[0]")
    private Output<String> secretKey;

    /**
     * @return The secret Key of the iam api key.
     * 
     */
    public Output<String> secretKey() {
        return this.secretKey;
    }
    /**
     * The date and time of the last update of the iam api key.
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The date and time of the last update of the iam api key.
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * ID of the user attached to the api key.
     * Only one of the `application_id` and `user_id` should be specified.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userId;

    /**
     * @return ID of the user attached to the api key.
     * Only one of the `application_id` and `user_id` should be specified.
     * 
     */
    public Output<Optional<String>> userId() {
        return Codegen.optional(this.userId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamApiKey(String name) {
        this(name, IamApiKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamApiKey(String name, @Nullable IamApiKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamApiKey(String name, @Nullable IamApiKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/iamApiKey:IamApiKey", name, args == null ? IamApiKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamApiKey(String name, Output<String> id, @Nullable IamApiKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("scaleway:index/iamApiKey:IamApiKey", name, state, makeResourceOptions(options, id));
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
    public static IamApiKey get(String name, Output<String> id, @Nullable IamApiKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamApiKey(name, id, state, options);
    }
}
