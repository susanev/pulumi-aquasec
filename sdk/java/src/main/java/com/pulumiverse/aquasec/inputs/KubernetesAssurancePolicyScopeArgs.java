// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumiverse.aquasec.inputs.KubernetesAssurancePolicyScopeVariableArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubernetesAssurancePolicyScopeArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubernetesAssurancePolicyScopeArgs Empty = new KubernetesAssurancePolicyScopeArgs();

    @Import(name="expression")
    private @Nullable Output<String> expression;

    public Optional<Output<String>> expression() {
        return Optional.ofNullable(this.expression);
    }

    @Import(name="variables")
    private @Nullable Output<List<KubernetesAssurancePolicyScopeVariableArgs>> variables;

    public Optional<Output<List<KubernetesAssurancePolicyScopeVariableArgs>>> variables() {
        return Optional.ofNullable(this.variables);
    }

    private KubernetesAssurancePolicyScopeArgs() {}

    private KubernetesAssurancePolicyScopeArgs(KubernetesAssurancePolicyScopeArgs $) {
        this.expression = $.expression;
        this.variables = $.variables;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubernetesAssurancePolicyScopeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubernetesAssurancePolicyScopeArgs $;

        public Builder() {
            $ = new KubernetesAssurancePolicyScopeArgs();
        }

        public Builder(KubernetesAssurancePolicyScopeArgs defaults) {
            $ = new KubernetesAssurancePolicyScopeArgs(Objects.requireNonNull(defaults));
        }

        public Builder expression(@Nullable Output<String> expression) {
            $.expression = expression;
            return this;
        }

        public Builder expression(String expression) {
            return expression(Output.of(expression));
        }

        public Builder variables(@Nullable Output<List<KubernetesAssurancePolicyScopeVariableArgs>> variables) {
            $.variables = variables;
            return this;
        }

        public Builder variables(List<KubernetesAssurancePolicyScopeVariableArgs> variables) {
            return variables(Output.of(variables));
        }

        public Builder variables(KubernetesAssurancePolicyScopeVariableArgs... variables) {
            return variables(List.of(variables));
        }

        public KubernetesAssurancePolicyScopeArgs build() {
            return $;
        }
    }

}
