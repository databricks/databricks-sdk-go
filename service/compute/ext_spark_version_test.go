package compute

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSparkVersionsResponse_Select_WithDuplicateSparkVersionsDifferentScala(t *testing.T) {
	resp := GetSparkVersionsResponse{
		Versions: []SparkVersion{
			{
				Key:  "16.4.x-scala2.12",
				Name: "16.4 LTS (includes Apache Spark 3.5.2, Scala 2.12)",
			},
			{
				Key:  "16.4.x-scala2.13",
				Name: "16.4 LTS (includes Apache Spark 3.5.2, Scala 2.13)",
			},
			{
				Key:  "17.1.x-scala2.13",
				Name: "17.1 (includes Apache Spark 4.0.0, Scala 2.13)",
			},
		},
	}

	// Should match the first found for 16.4.x ignoring Scala version
	selected, err := resp.Select(SparkVersionRequest{LongTermSupport: true, Latest: false, Scala: "2.1"})
	require.NoError(t, err)
	require.Equal(t, "16.4.x-scala2.12", selected)

	selected, err = resp.Select(SparkVersionRequest{Scala: "2.13", Latest: true})
	require.NoError(t, err)
	require.Equal(t, "17.1.x-scala2.13", selected)

	selected, err = resp.Select(SparkVersionRequest{Scala: "2.13", SparkVersion: "3.5.2"})
	require.NoError(t, err)
	require.Equal(t, "16.4.x-scala2.13", selected)
}
