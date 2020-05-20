package nri

import (
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-winservices/src/scraper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProccessInventory(t *testing.T) {
	entityRules := loadRules()
	i, _ := integration.New("integrationName", "integrationVersion")
	mfbn := scraper.MetricFamiliesByName{
		"wmi_service_info":       metricFamlilyServiceInfo,
		"wmi_service_start_mode": metricFamlilyService,
		"wmi_cs_hostname":        metricFamlilyServiceHostname,
	}

	validator := NewValidator(serviceName, "", "")
	err := ProcessMetrics(i, mfbn, validator)
	require.NoError(t, err)
	require.Greater(t, len(i.Entities), 0)

	err = ProcessInventory(i)
	require.NoError(t, err)
	require.Greater(t, len(i.Entities), 0)

	item, ok := i.Entities[0].Inventory.Item(entityTypeInventory)
	require.True(t, ok)
	require.Equal(t, hostname, item[entityRules.EntityName.HostnameNrdbLabelName])
	require.Equal(t, hostname+":"+serviceName, item["name"])
	require.Equal(t, serviceName, item["windowsService.name"])
	require.Equal(t, serviceDisplayName, item["windowsService.displayName"])
	require.Equal(t, servicePid, item["windowsService.processId"])
	require.Equal(t, serviceStartMode, item["windowsService.startMode"])

}
