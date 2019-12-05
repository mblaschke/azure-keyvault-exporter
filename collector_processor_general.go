package main

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
)

type CollectorProcessorGeneralInterface interface {
	Setup(collector *CollectorGeneral)
	Reset()
	Collect(ctx context.Context, callback chan<- func(), subscription subscriptions.Subscription)
}

type CollectorProcessorGeneral struct {
	CollectorProcessorGeneralInterface
	CollectorReference *CollectorGeneral
}

func NewCollectorGeneral(name string, processor CollectorProcessorGeneralInterface) *CollectorGeneral {
	collector := CollectorGeneral{
		CollectorBase: CollectorBase{
			Name:               name,
			AzureSubscriptions: AzureSubscriptions,
		},

		Processor: processor,
	}

	return &collector
}
