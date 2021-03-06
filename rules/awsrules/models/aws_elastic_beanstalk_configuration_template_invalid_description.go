// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule checks the pattern is valid
type AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule returns new rule with default attributes
func NewAwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule() *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule {
	return &AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule{
		resourceType:  "aws_elastic_beanstalk_configuration_template",
		attributeName: "description",
		max:           200,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Name() string {
	return "aws_elastic_beanstalk_configuration_template_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkConfigurationTemplateInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 200 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
