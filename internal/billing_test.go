package internal

import (
	"io"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccUsageDownload(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	resp, err := a.BillableUsage.Download(ctx, billing.DownloadRequest{
		StartMonth: "2023-01",
		EndMonth:   "2023-02",
	})
	require.NoError(t, err)
	out, err := io.ReadAll(resp.Contents)
	require.NoError(t, err)
	assert.NotEmpty(t, out)
}

func TestMwsAccLogDelivery(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	creds, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_LOGDELIVERY_ARN"),
			},
		},
	})
	require.NoError(t, err)
	defer a.Credentials.DeleteByCredentialsId(ctx, creds.CredentialsId)

	bucket, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)
	defer a.Storage.DeleteByStorageConfigurationId(ctx, bucket.StorageConfigurationId)

	// TODO: OpenAPI: x-databricks-sdk-inline on schema
	created, err := a.LogDelivery.Create(ctx, billing.WrappedCreateLogDeliveryConfiguration{
		LogDeliveryConfiguration: &billing.CreateLogDeliveryConfigurationParams{
			ConfigName:             RandomName("sdk-go-"),
			CredentialsId:          creds.CredentialsId,
			StorageConfigurationId: bucket.StorageConfigurationId,
			LogType:                billing.LogTypeAuditLogs,
			OutputFormat:           billing.OutputFormatJson,
		},
	})
	require.NoError(t, err)
	defer a.LogDelivery.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
		LogDeliveryConfigurationId: created.LogDeliveryConfiguration.ConfigId,
		Status:                     billing.LogDeliveryConfigStatusDisabled,
	})

	byId, err := a.LogDelivery.GetByLogDeliveryConfigurationId(ctx, created.LogDeliveryConfiguration.ConfigId)
	require.NoError(t, err)
	assert.Equal(t, billing.LogDeliveryConfigStatusDisabled, byId.LogDeliveryConfiguration.Status)

	all, err := a.LogDelivery.ListAll(ctx, billing.ListLogDeliveryRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	byName, err := a.LogDelivery.GetByConfigName(ctx, byId.LogDeliveryConfiguration.ConfigName)
	require.NoError(t, err)

	assert.Equal(t, byId.LogDeliveryConfiguration.ConfigId, byName.ConfigId)
}

func TestMwsAccBudgets(t *testing.T) {
	// Currently, the API returns errors
	t.SkipNow()
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	// TODO: OpenAPI: x-databricks-sdk-inline on schema
	created, err := a.Budgets.Create(ctx, billing.CreateBudgetConfigurationRequest{
		Budget: billing.CreateBudgetConfigurationBudget{
			DisplayName: RandomName("go-sdk-"),
			Filter: &billing.BudgetConfigurationFilter{
				Tags: []billing.BudgetConfigurationFilterTagClause{
					{
						Key: "tagName",
						Value: &billing.BudgetConfigurationFilterClause{
							Operator: billing.BudgetConfigurationFilterOperatorIn,
							Values:   []string{"all"},
						},
					},
				}},
			AlertConfigurations: []billing.CreateBudgetConfigurationBudgetAlertConfigurations{
				{
					TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
					QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
					TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
					QuantityThreshold: "100",
					ActionConfigurations: []billing.CreateBudgetConfigurationBudgetActionConfigurations{
						{
							ActionType: billing.ActionConfigurationTypeEmailNotification,
							Target:     "admin@example.com",
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	defer a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetConfigurationId)

	_, err = a.Budgets.Update(ctx, billing.UpdateBudgetConfigurationRequest{

		BudgetId: created.Budget.BudgetConfigurationId,
		Budget: billing.UpdateBudgetConfigurationBudget{
			DisplayName: RandomName("go-sdk-updated-"),
			Filter: &billing.BudgetConfigurationFilter{
				Tags: []billing.BudgetConfigurationFilterTagClause{
					{
						Key: "tagName",
						Value: &billing.BudgetConfigurationFilterClause{
							Operator: billing.BudgetConfigurationFilterOperatorIn,
							Values:   []string{"all"},
						},
					},
				}},
			AlertConfigurations: []billing.AlertConfiguration{
				{
					TimePeriod:        billing.AlertConfigurationTimePeriodMonth,
					QuantityType:      billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
					TriggerType:       billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
					QuantityThreshold: "50",
					ActionConfigurations: []billing.ActionConfiguration{
						{
							ActionType: billing.ActionConfigurationTypeEmailNotification,
							Target:     "admin@example.com",
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)

	byId, err := a.Budgets.GetByBudgetId(ctx, created.Budget.BudgetConfigurationId)
	require.NoError(t, err)
	assert.NotEqual(t, created.Budget.DisplayName, byId.Budget.DisplayName)

	all, err := a.Budgets.ListAll(ctx, billing.ListBudgetConfigurationsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}
