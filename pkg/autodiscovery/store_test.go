// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package autodiscovery

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DataDog/datadog-agent/pkg/autodiscovery/integration"
)

func TestServiceToConfig(t *testing.T) {
	s := newStore()
	service := dummyService{
		ID:            "a5901276aed1",
		ADIdentifiers: []string{"redis"},
		Hosts:         map[string]string{"bridge": "127.0.0.1"},
	}
	s.addConfigForService(service.GetEntity(), integration.Config{Name: "foo"})
	s.addConfigForService(service.GetEntity(), integration.Config{Name: "bar"})
	assert.Equal(t, len(s.getConfigsForService(service.GetEntity())), 2)
	s.removeConfigsForService(service.GetEntity())
	s.addConfigForService(service.GetEntity(), integration.Config{Name: "foo"})
	assert.Equal(t, len(s.getConfigsForService(service.GetEntity())), 1)
}

func TestTemplateToConfig(t *testing.T) {
	s := newStore()
	s.addConfigForTemplate("digest1", integration.Config{Name: "foo"})
	s.addConfigForTemplate("digest1", integration.Config{Name: "bar"})
	s.addConfigForTemplate("digest2", integration.Config{Name: "foo"})

	assert.Len(t, s.getConfigsForTemplate("digest1"), 2)
	assert.Len(t, s.getConfigsForTemplate("digest2"), 1)

	s.removeConfigsForTemplate("digest1")
	assert.Len(t, s.getConfigsForTemplate("digest1"), 0)
	assert.Len(t, s.getConfigsForTemplate("digest2"), 1)

	s.addConfigForTemplate("digest1", integration.Config{Name: "foo"})
	assert.Len(t, s.getConfigsForTemplate("digest1"), 1)
	assert.Len(t, s.getConfigsForTemplate("digest2"), 1)
}
