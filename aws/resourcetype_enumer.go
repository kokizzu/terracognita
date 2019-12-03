// Code generated by "enumer -type ResourceType -addprefix aws_ -transform snake -linecomment"; DO NOT EDIT.

package aws

import (
	"fmt"
)

const _ResourceTypeName = "aws_instanceaws_vpcaws_security_groupaws_subnetaws_ebs_volumeaws_elasticache_clusteraws_elbaws_albaws_db_instanceaws_s3_bucketaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_zoneaws_route53_zone_associationaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_ses_active_receipt_rule_setaws_ses_domain_identityaws_ses_domain_identity_verificationaws_ses_domain_dkimaws_ses_domain_mail_fromaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_configuration_setaws_ses_identity_notification_topicaws_ses_templateaws_launch_configurationaws_launch_templateaws_autoscaling_group"

var _ResourceTypeIndex = [...]uint16{0, 12, 19, 37, 47, 61, 84, 91, 98, 113, 126, 153, 190, 215, 236, 267, 280, 304, 324, 355, 379, 410, 424, 436, 455, 485, 506, 532, 544, 573, 592, 622, 648, 672, 693, 711, 727, 755, 784, 821, 852, 875, 911, 930, 954, 976, 996, 1020, 1045, 1080, 1096, 1120, 1139, 1160}

const _ResourceTypeLowerName = "aws_instanceaws_vpcaws_security_groupaws_subnetaws_ebs_volumeaws_elasticache_clusteraws_elbaws_albaws_db_instanceaws_s3_bucketaws_cloudfront_distributionaws_cloudfront_origin_access_identityaws_cloudfront_public_keyaws_iam_account_aliasaws_iam_account_password_policyaws_iam_groupaws_iam_group_membershipaws_iam_group_policyaws_iam_group_policy_attachmentaws_iam_instance_profileaws_iam_openid_connect_provideraws_iam_policyaws_iam_roleaws_iam_role_policyaws_iam_role_policy_attachmentaws_iam_saml_provideraws_iam_server_certificateaws_iam_useraws_iam_user_group_membershipaws_iam_user_policyaws_iam_user_policy_attachmentaws_route53_delegation_setaws_route53_health_checkaws_route53_query_logaws_route53_recordaws_route53_zoneaws_route53_zone_associationaws_route53_resolver_endpointaws_route53_resolver_rule_associationaws_ses_active_receipt_rule_setaws_ses_domain_identityaws_ses_domain_identity_verificationaws_ses_domain_dkimaws_ses_domain_mail_fromaws_ses_receipt_filteraws_ses_receipt_ruleaws_ses_receipt_rule_setaws_ses_configuration_setaws_ses_identity_notification_topicaws_ses_templateaws_launch_configurationaws_launch_templateaws_autoscaling_group"

func (i ResourceType) String() string {
	i -= 1
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i+1)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

var _ResourceTypeValues = []ResourceType{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:12]:           1,
	_ResourceTypeLowerName[0:12]:      1,
	_ResourceTypeName[12:19]:          2,
	_ResourceTypeLowerName[12:19]:     2,
	_ResourceTypeName[19:37]:          3,
	_ResourceTypeLowerName[19:37]:     3,
	_ResourceTypeName[37:47]:          4,
	_ResourceTypeLowerName[37:47]:     4,
	_ResourceTypeName[47:61]:          5,
	_ResourceTypeLowerName[47:61]:     5,
	_ResourceTypeName[61:84]:          6,
	_ResourceTypeLowerName[61:84]:     6,
	_ResourceTypeName[84:91]:          7,
	_ResourceTypeLowerName[84:91]:     7,
	_ResourceTypeName[91:98]:          8,
	_ResourceTypeLowerName[91:98]:     8,
	_ResourceTypeName[98:113]:         9,
	_ResourceTypeLowerName[98:113]:    9,
	_ResourceTypeName[113:126]:        10,
	_ResourceTypeLowerName[113:126]:   10,
	_ResourceTypeName[126:153]:        11,
	_ResourceTypeLowerName[126:153]:   11,
	_ResourceTypeName[153:190]:        12,
	_ResourceTypeLowerName[153:190]:   12,
	_ResourceTypeName[190:215]:        13,
	_ResourceTypeLowerName[190:215]:   13,
	_ResourceTypeName[215:236]:        14,
	_ResourceTypeLowerName[215:236]:   14,
	_ResourceTypeName[236:267]:        15,
	_ResourceTypeLowerName[236:267]:   15,
	_ResourceTypeName[267:280]:        16,
	_ResourceTypeLowerName[267:280]:   16,
	_ResourceTypeName[280:304]:        17,
	_ResourceTypeLowerName[280:304]:   17,
	_ResourceTypeName[304:324]:        18,
	_ResourceTypeLowerName[304:324]:   18,
	_ResourceTypeName[324:355]:        19,
	_ResourceTypeLowerName[324:355]:   19,
	_ResourceTypeName[355:379]:        20,
	_ResourceTypeLowerName[355:379]:   20,
	_ResourceTypeName[379:410]:        21,
	_ResourceTypeLowerName[379:410]:   21,
	_ResourceTypeName[410:424]:        22,
	_ResourceTypeLowerName[410:424]:   22,
	_ResourceTypeName[424:436]:        23,
	_ResourceTypeLowerName[424:436]:   23,
	_ResourceTypeName[436:455]:        24,
	_ResourceTypeLowerName[436:455]:   24,
	_ResourceTypeName[455:485]:        25,
	_ResourceTypeLowerName[455:485]:   25,
	_ResourceTypeName[485:506]:        26,
	_ResourceTypeLowerName[485:506]:   26,
	_ResourceTypeName[506:532]:        27,
	_ResourceTypeLowerName[506:532]:   27,
	_ResourceTypeName[532:544]:        28,
	_ResourceTypeLowerName[532:544]:   28,
	_ResourceTypeName[544:573]:        29,
	_ResourceTypeLowerName[544:573]:   29,
	_ResourceTypeName[573:592]:        30,
	_ResourceTypeLowerName[573:592]:   30,
	_ResourceTypeName[592:622]:        31,
	_ResourceTypeLowerName[592:622]:   31,
	_ResourceTypeName[622:648]:        32,
	_ResourceTypeLowerName[622:648]:   32,
	_ResourceTypeName[648:672]:        33,
	_ResourceTypeLowerName[648:672]:   33,
	_ResourceTypeName[672:693]:        34,
	_ResourceTypeLowerName[672:693]:   34,
	_ResourceTypeName[693:711]:        35,
	_ResourceTypeLowerName[693:711]:   35,
	_ResourceTypeName[711:727]:        36,
	_ResourceTypeLowerName[711:727]:   36,
	_ResourceTypeName[727:755]:        37,
	_ResourceTypeLowerName[727:755]:   37,
	_ResourceTypeName[755:784]:        38,
	_ResourceTypeLowerName[755:784]:   38,
	_ResourceTypeName[784:821]:        39,
	_ResourceTypeLowerName[784:821]:   39,
	_ResourceTypeName[821:852]:        40,
	_ResourceTypeLowerName[821:852]:   40,
	_ResourceTypeName[852:875]:        41,
	_ResourceTypeLowerName[852:875]:   41,
	_ResourceTypeName[875:911]:        42,
	_ResourceTypeLowerName[875:911]:   42,
	_ResourceTypeName[911:930]:        43,
	_ResourceTypeLowerName[911:930]:   43,
	_ResourceTypeName[930:954]:        44,
	_ResourceTypeLowerName[930:954]:   44,
	_ResourceTypeName[954:976]:        45,
	_ResourceTypeLowerName[954:976]:   45,
	_ResourceTypeName[976:996]:        46,
	_ResourceTypeLowerName[976:996]:   46,
	_ResourceTypeName[996:1020]:       47,
	_ResourceTypeLowerName[996:1020]:  47,
	_ResourceTypeName[1020:1045]:      48,
	_ResourceTypeLowerName[1020:1045]: 48,
	_ResourceTypeName[1045:1080]:      49,
	_ResourceTypeLowerName[1045:1080]: 49,
	_ResourceTypeName[1080:1096]:      50,
	_ResourceTypeLowerName[1080:1096]: 50,
	_ResourceTypeName[1096:1120]:      51,
	_ResourceTypeLowerName[1096:1120]: 51,
	_ResourceTypeName[1120:1139]:      52,
	_ResourceTypeLowerName[1120:1139]: 52,
	_ResourceTypeName[1139:1160]:      53,
	_ResourceTypeLowerName[1139:1160]: 53,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:12],
	_ResourceTypeName[12:19],
	_ResourceTypeName[19:37],
	_ResourceTypeName[37:47],
	_ResourceTypeName[47:61],
	_ResourceTypeName[61:84],
	_ResourceTypeName[84:91],
	_ResourceTypeName[91:98],
	_ResourceTypeName[98:113],
	_ResourceTypeName[113:126],
	_ResourceTypeName[126:153],
	_ResourceTypeName[153:190],
	_ResourceTypeName[190:215],
	_ResourceTypeName[215:236],
	_ResourceTypeName[236:267],
	_ResourceTypeName[267:280],
	_ResourceTypeName[280:304],
	_ResourceTypeName[304:324],
	_ResourceTypeName[324:355],
	_ResourceTypeName[355:379],
	_ResourceTypeName[379:410],
	_ResourceTypeName[410:424],
	_ResourceTypeName[424:436],
	_ResourceTypeName[436:455],
	_ResourceTypeName[455:485],
	_ResourceTypeName[485:506],
	_ResourceTypeName[506:532],
	_ResourceTypeName[532:544],
	_ResourceTypeName[544:573],
	_ResourceTypeName[573:592],
	_ResourceTypeName[592:622],
	_ResourceTypeName[622:648],
	_ResourceTypeName[648:672],
	_ResourceTypeName[672:693],
	_ResourceTypeName[693:711],
	_ResourceTypeName[711:727],
	_ResourceTypeName[727:755],
	_ResourceTypeName[755:784],
	_ResourceTypeName[784:821],
	_ResourceTypeName[821:852],
	_ResourceTypeName[852:875],
	_ResourceTypeName[875:911],
	_ResourceTypeName[911:930],
	_ResourceTypeName[930:954],
	_ResourceTypeName[954:976],
	_ResourceTypeName[976:996],
	_ResourceTypeName[996:1020],
	_ResourceTypeName[1020:1045],
	_ResourceTypeName[1045:1080],
	_ResourceTypeName[1080:1096],
	_ResourceTypeName[1096:1120],
	_ResourceTypeName[1120:1139],
	_ResourceTypeName[1139:1160],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
