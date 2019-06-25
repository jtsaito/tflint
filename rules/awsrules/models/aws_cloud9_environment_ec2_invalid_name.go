// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloud9EnvironmentEc2InvalidNameRule checks the pattern is valid
type AwsCloud9EnvironmentEc2InvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloud9EnvironmentEc2InvalidNameRule returns new rule with default attributes
func NewAwsCloud9EnvironmentEc2InvalidNameRule() *AwsCloud9EnvironmentEc2InvalidNameRule {
	return &AwsCloud9EnvironmentEc2InvalidNameRule{
		resourceType:  "aws_cloud9_environment_ec2",
		attributeName: "name",
		max:           60,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Name() string {
	return "aws_cloud9_environment_ec2_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloud9EnvironmentEc2InvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 60 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}