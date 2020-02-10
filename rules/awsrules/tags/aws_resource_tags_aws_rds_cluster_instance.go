package awsrules

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsRdsClusterInstanceTagsRule checks whether the resource is tagged correctly
type AwsRdsClusterInstanceTagsRule struct {
	resourceType  string
	attributeName string
}


// NewAwsRdsClusterInstanceTagsRule returns new tags rule with default attributes
func NewAwsRdsClusterInstanceTagsRule() *AwsRdsClusterInstanceTagsRule {
	return &AwsRdsClusterInstanceTagsRule{
		resourceType:  "aws_rds_cluster_instance",
		attributeName: "tags",
	}
}

// Name returns the rule name
func (r *AwsRdsClusterInstanceTagsRule) Name() string {
	return "aws_resource_tags_aws_rds_cluster_instance"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRdsClusterInstanceTagsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRdsClusterInstanceTagsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRdsClusterInstanceTagsRule) Link() string {
	return ""
}

// Check checks for matching tags
func (r *AwsRdsClusterInstanceTagsRule) Check(runner *tflint.Runner) error {
	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var tags map[string]string
		err := runner.EvaluateExpr(attribute.Expr, &tags)

		return runner.EnsureNoError(err, func() error {
			configTags := runner.GetConfigTags()
			tagKeys := []string{}
			hash := make(map[string]bool)
			for k, _ := range tags {
				tagKeys = append(tagKeys, k)
				hash[k] = true
			}
			var found []string
			for _, tag := range configTags {
				if _, ok := hash[tag]; ok {
					found = append(found, tag)
				}
			}
			if len(found) != len(configTags) {
				runner.EmitIssue(r, fmt.Sprintf("Wanted tags: %v, found: %v\n", configTags, tags), attribute.Expr.Range() )
			}
			return nil
		})
	})
}
