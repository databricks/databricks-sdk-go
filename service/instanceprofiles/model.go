// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instanceprofiles

// all definitions in this file are in alphabetical order

type AddInstanceProfileRequest struct {
    // The AWS ARN of the role in instance profile to register with Databricks. 
    // It should look like: `arn:aws:iam::&lt;account-id&gt;:role/&lt;name&gt;`. 
    IamRoleArn string `json:"iam_role_arn,omitempty"`
    // The AWS ARN of the instance profile to register with Databricks. It 
    // should look like: `arn:aws:iam::&lt;account-id&gt;:instance-profile/&lt;name&gt;`. 
    InstanceProfileArn string `json:"instance_profile_arn"`
    // Boolean flag indicating whether the instance profile should only be used 
    // in credential passthrough scenarios. If true, it means the instance 
    // profile contains an meta IAM role which could assume a wide range of 
    // roles. Therefore it should always be used with authorization. 
    IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`
    // By default, Databricks validates that it has sufficient permissions to 
    // launch instances with the instance profile. This validation uses AWS 
    // dry-run mode for the `RunInstances` API. If validation fails with an 
    // error message that does not indicate an IAM related permission issue, 
    // (e.g. &#34;Your requested instance type is not supported in your requested 
    // Availability Zone&#34;), you may pass this flag to skip the validation and 
    // forcibly add the instance profile. 
    SkipValidation bool `json:"skip_validation,omitempty"`
}


type InstanceProfile struct {
    // The AWS ARN of the role in instance profile to register with Databricks. 
    // It should look like: `arn:aws:iam::&lt;account-id&gt;:role/&lt;name&gt;`. 
    IamRoleArn string `json:"iam_role_arn,omitempty"`
    // The AWS ARN of the instance profile to register with Databricks. It 
    // should look like: `arn:aws:iam::&lt;account-id&gt;:instance-profile/&lt;name&gt;`. 
    InstanceProfileArn string `json:"instance_profile_arn"`
    // Boolean flag indicating whether the instance profile should only be used 
    // in credential passthrough scenarios. If true, it means the instance 
    // profile contains an meta IAM role which could assume a wide range of 
    // roles. Therefore it should always be used with authorization. 
    IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`
}


type ListInstanceProfilesResponse struct {
    // A list of instance profiles that the user can access. 
    InstanceProfiles []InstanceProfile `json:"instance_profiles,omitempty"`
}


type RemoveInstanceProfileRequest struct {
    // The arn of the instance profile to remove. 
    InstanceProfileArn string `json:"instance_profile_arn"`
}

