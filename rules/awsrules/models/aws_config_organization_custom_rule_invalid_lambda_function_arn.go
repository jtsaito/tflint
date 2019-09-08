// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule checks the pattern is valid
type AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule returns new rule with default attributes
func NewAwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule() *AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule {
	return &AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule{
		resourceType:  "aws_config_organization_custom_rule",
		attributeName: "lambda_function_arn",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule) Name() string {
	return "aws_config_organization_custom_rule_invalid_lambda_function_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsConfigOrganizationCustomRuleInvalidLambdaFunctionArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"lambda_function_arn must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"lambda_function_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}