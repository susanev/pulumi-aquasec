// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.aquasec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumiverse.aquasec.outputs.GetImageAssurancePolicyScopeVariable;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetImageAssurancePolicyScope {
    private String expression;
    private @Nullable List<GetImageAssurancePolicyScopeVariable> variables;

    private GetImageAssurancePolicyScope() {}
    public String expression() {
        return this.expression;
    }
    public List<GetImageAssurancePolicyScopeVariable> variables() {
        return this.variables == null ? List.of() : this.variables;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetImageAssurancePolicyScope defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String expression;
        private @Nullable List<GetImageAssurancePolicyScopeVariable> variables;
        public Builder() {}
        public Builder(GetImageAssurancePolicyScope defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.expression = defaults.expression;
    	      this.variables = defaults.variables;
        }

        @CustomType.Setter
        public Builder expression(String expression) {
            this.expression = Objects.requireNonNull(expression);
            return this;
        }
        @CustomType.Setter
        public Builder variables(@Nullable List<GetImageAssurancePolicyScopeVariable> variables) {
            this.variables = variables;
            return this;
        }
        public Builder variables(GetImageAssurancePolicyScopeVariable... variables) {
            return variables(List.of(variables));
        }
        public GetImageAssurancePolicyScope build() {
            final var o = new GetImageAssurancePolicyScope();
            o.expression = expression;
            o.variables = variables;
            return o;
        }
    }
}
