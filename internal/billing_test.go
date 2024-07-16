package internal

import (
	"io"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	out, err := io.ReadAll(resp.Contents)
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	defer a.Credentials.DeleteByCredentialsId(ctx, creds.CredentialsId)

	bucket, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	defer a.LogDelivery.PatchStatus(ctx, billing.UpdateLogDeliveryConfigurationStatusRequest{
		LogDeliveryConfigurationId: created.LogDeliveryConfiguration.ConfigId,
		Status:                     billing.LogDeliveryConfigStatusDisabled,
	})

	byId, err := a.LogDelivery.GetByLogDeliveryConfigurationId(ctx, created.LogDeliveryConfiguration.ConfigId)
	assert.NoError(t, err)
	assert.Equal(t, billing.LogDeliveryConfigStatusDisabled, byId.LogDeliveryConfiguration.Status)

	all, err := a.LogDelivery.ListAll(ctx, billing.ListLogDeliveryRequest{})
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)

	byName, err := a.LogDelivery.GetByConfigName(ctx, byId.LogDeliveryConfiguration.ConfigName)
	assert.NoError(t, err)

	assert.Equal(t, byId.LogDeliveryConfiguration.ConfigId, byName.ConfigId)
}

func TestMwsAccBudgets(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	// TODO: OpenAPI: x-databricks-sdk-inline on schema
	created, err := a.Budgets.Create(ctx, billing.WrappedBudget{
		Budget: billing.Budget{
			Name:         RandomName("go-sdk-"),
			Filter:       "tag.tagName = 'all'",
			Period:       "1 month",
			StartDate:    "2022-01-01",
			TargetAmount: "100",
			Alerts: []billing.BudgetAlert{
				{
					EmailNotifications: []string{"admin@example.com"},
					MinPercentage:      50,
				},
			},
		},
	})
	assert.NoError(t, err)
	defer a.Budgets.DeleteByBudgetId(ctx, created.Budget.BudgetId)

	err = a.Budgets.Update(ctx, billing.WrappedBudget{
		BudgetId: created.Budget.BudgetId,
		Budget: billing.Budget{
			Name:         RandomName("go-sdk-updated-"),
			Filter:       "tag.tagName = 'all'",
			Period:       "1 month",
			StartDate:    "2022-01-01",
			TargetAmount: "100",
			Alerts: []billing.BudgetAlert{
				{
					EmailNotifications: []string{"admin@example.com"},
					MinPercentage:      70,
				},
			},
		},
	})
	assert.NoError(t, err)

	byId, err := a.Budgets.GetByBudgetId(ctx, created.Budget.BudgetId)
	assert.NoError(t, err)
	assert.NotEqual(t, created.Budget.Name, byId.Budget.Name)

	byName, err := a.Budgets.GetByName(ctx, byId.Budget.Name)
	assert.NoError(t, err)
	assert.Equal(t, byId.Budget.BudgetId, byName.BudgetId)

	all, err := a.Budgets.ListAll(ctx)
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)

	names, err := a.Budgets.BudgetWithStatusNameToBudgetIdMap(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(all), len(names))
	assert.Equal(t, created.Budget.BudgetId, names[byId.Budget.Name])
}
