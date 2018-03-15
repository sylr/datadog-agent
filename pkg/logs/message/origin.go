// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

package message

import (
	"strings"

	"github.com/DataDog/datadog-agent/pkg/logs/config"
)

// Origin represents the Origin of a message
type Origin struct {
	Identifier string
	LogsConfig *config.LogsConfig
	Offset     int64
	Timestamp  string
	tags       []string
}

// NewOrigin returns a new Origin
func NewOrigin() *Origin {
	return &Origin{}
}

// Tags returns the tags of the origin.
func (o *Origin) Tags() []string {
	tags := o.tags
	if o.LogsConfig.Source != "" {
		tags = append(tags, "source:"+o.LogsConfig.Source)
	}
	if o.LogsConfig.SourceCategory != "" {
		tags = append(tags, "sourcecategory:"+o.LogsConfig.SourceCategory)
	}

	tags = append(tags, o.LogsConfig.Tags...)
	return tags
}

// TagsPayload returns the raw tag payload of the origin.
func (o *Origin) TagsPayload() []byte {
	var tagsPayload []byte
	if o.LogsConfig.Source != "" {
		tagsPayload = append(tagsPayload, []byte("[dd ddsource=\""+o.LogsConfig.Source+"\"]")...)
	}
	if o.LogsConfig.SourceCategory != "" {
		tagsPayload = append(tagsPayload, []byte("[dd ddsourcecategory=\""+o.LogsConfig.SourceCategory+"\"]")...)
	}

	var tags []string
	tags = append(tags, o.LogsConfig.Tags...)
	tags = append(tags, o.tags...)

	if len(tags) > 0 {
		tagsPayload = append(tagsPayload, []byte("[dd ddtags=\""+strings.Join(tags, ",")+"\"]")...)
	}
	if len(tagsPayload) == 0 {
		tagsPayload = []byte{}
	}
	return tagsPayload
}

// SetTags sets the tags of the origin.
func (o *Origin) SetTags(tags []string) {
	o.tags = tags
}
