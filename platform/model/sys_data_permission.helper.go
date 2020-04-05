package model

import "html/template"

// ParseRule defined parse rule
func (m *SysDataPermission) ParseRule(roleRules []SysDataPermissionDetail) interface{} {
	roleRule := ""
	for i, rule := range roleRules {
		if i == 0 {
			roleRule = rule.Rule.String
		} else if i > 0 {
			if roleRule != "" {
				roleRule = roleRule + " or " + rule.Rule.String
			} else {
				roleRule = rule.Rule.String
			}
		}
	}
	if len(roleRules) > 1 {
		roleRule = "(" + roleRule + ")"
	}
	return template.HTML(roleRule)
}
