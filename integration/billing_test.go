package integration

import (
	"io"
	"testing"

	"github.com/databricks/databricks-sdk-go/billing/v2"
	"github.com/databricks/databricks-sdk-go/provisioning/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccUsageDownload(t *testing.T) {
	ctx, cfg := accountTest(t)
	if !cfg.IsAws() {
		t.SkipNow()
	}
	BillingUsageAPI, err := billing.NewBillableUsageClient(cfg)
	require.NoError(t, err)
	resp, err := BillingUsageAPI.Download(ctx, billing.DownloadRequest{
		StartMonth: "2024-08",
		EndMonth:   "2024-09",
	})
	require.NoError(t, err)
	out, err := io.ReadAll(resp.Contents)
	require.NoError(t, err)
	assert.NotEmpty(t, out)
}

func TestMwsAccLogDelivery(t *testing.T) {
	ctx, cfg := accountTest(t)
	if !cfg.IsAws() {
		t.SkipNow()
	}
	CredenialsAPI, err := provisioning.NewCredentialsClient(cfg)
	require.NoError(t, err)
	creds, err := CredenialsAPI.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_LOGDELIVERY_ARN"),
			},
		},
	})
	require.NoError(t, err)
	defer CredenialsAPI.DeleteByCredentialsId(ctx, creds.CredentialsId)

	StorageAPI, err := provisioning.NewStorageClient(cfg)
	bucket, err := StorageAPI.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)
	defer StorageAPI.DeleteByStorageConfigurationId(ctx, bucket.StorageConfigurationId)

	// TODO: OpenAPI: x-databricks-sdk-inline on schema
	LogDeliveryAPI, err := billing.NewLogDeliveryClient(cfg)
	created, err := LogDeliveryAPI.Create(ctx, billing.WrappedCreateLogDeliveryConfiguration{
		LogDeliveryConfiguration: &billing.CreateLogDeliveryConfigurationParams{
			ConfigName:             RandomName("sdk-go-"),
			CredentialsId:          creds.CredentialsId,
			StorageConfigurationId: bucket.StorageConfigurationId,
			LogType:                billing.LogTypeAuditLogs,
			OutputFormat:           billing.OutputFormatJson,
		},
	})
	require.NoError(t, err)
	defer LogDeliveryAPI.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
		LogDeliveryConfigurationId: created.LogDeliveryConfiguration.ConfigId,
		Status:                     billing.LogDeliveryConfigStatusDisabled,
	})

	byId, err := LogDeliveryAPI.GetByLogDeliveryConfigurationId(ctx, created.LogDeliveryConfiguration.ConfigId)
	require.NoError(t, err)
	assert.Equal(t, billing.LogDeliveryConfigStatusDisabled, byId.LogDeliveryConfiguration.Status)

	all, err := LogDeliveryAPI.ListAll(ctx, billing.ListLogDeliveryRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)

	byName, err := LogDeliveryAPI.GetByConfigName(ctx, byId.LogDeliveryConfiguration.ConfigName)
	require.NoError(t, err)

	assert.Equal(t, byId.LogDeliveryConfiguration.ConfigId, byName.ConfigId)
}

func TestMwsAccBudgets(t *testing.T) {
	ctx, cfg := accountTest(t)
	if !cfg.IsAws() {
		t.SkipNow()
	}

	// TODO: OpenAPI: x-databricks-sdk-inline on schema
	BudgetsAPI, err := billing.NewBudgetsClient(cfg)
	created, err := BudgetsAPI.Create(ctx, billing.CreateBudgetConfigurationRequest{
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
	defer BudgetsAPI.DeleteByBudgetId(ctx, created.Budget.BudgetConfigurationId)

	_, err = BudgetsAPI.Update(ctx, billing.UpdateBudgetConfigurationRequest{

		BudgetId: created.Budget.BudgetConfigurationId,
		Budget: billing.UpdateBudgetConfigurationBudget{
			BudgetConfigurationId: created.Budget.BudgetConfigurationId,
			DisplayName:           RandomName("go-sdk-updated-"),
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
					AlertConfigurationId: created.Budget.AlertConfigurations[0].AlertConfigurationId,
					TimePeriod:           billing.AlertConfigurationTimePeriodMonth,
					QuantityType:         billing.AlertConfigurationQuantityTypeListPriceDollarsUsd,
					TriggerType:          billing.AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
					QuantityThreshold:    "50",
					ActionConfigurations: created.Budget.AlertConfigurations[0].ActionConfigurations,
				},
			},
		},
	})
	require.NoError(t, err)

	byId, err := BudgetsAPI.GetByBudgetId(ctx, created.Budget.BudgetConfigurationId)
	require.NoError(t, err)
	assert.NotEqual(t, created.Budget.DisplayName, byId.Budget.DisplayName)

	all, err := BudgetsAPI.ListAll(ctx, billing.ListBudgetConfigurationsRequest{})
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}
