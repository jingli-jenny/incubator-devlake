/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package archived

import (
	"github.com/apache/incubator-devlake/models/migrationscripts/archived"
)

type GiteeConnection struct {
	RestConnection `mapstructure:",squash"`
	AccessToken    `mapstructure:",squash"`
}

type RestConnection struct {
	BaseConnection `mapstructure:",squash"`
	Endpoint       string `mapstructure:"endpoint" validate:"required" json:"endpoint"`
	Proxy          string `mapstructure:"proxy" json:"proxy"`
	RateLimit      int    `comment:"api request rate limt per hour" json:"rateLimit"`
}

type BaseConnection struct {
	Name string `gorm:"type:varchar(100);uniqueIndex" json:"name" validate:"required"`
	archived.Model
}

type AccessToken struct {
	Token string `mapstructure:"token" validate:"required" json:"token" encrypt:"yes"`
}

// GitlabResponse is this object conforms to what the frontend currently expects.
type GitlabResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	GiteeConnection
}

// Using User because it requires authentication.
type ApiUserResponse struct {
	Id   int
	Name string `json:"name"`
}

type Config struct {
	PrType               string `mapstructure:"prType" env:"GITEE_PR_TYPE" json:"prType"`
	PrComponent          string `mapstructure:"prComponent" env:"GITEE_PR_COMPONENT" json:"prComponent"`
	PrBodyClosePattern   string `mapstructure:"prBodyClosePattern" json:"prBodyClosePattern"`
	IssueSeverity        string `mapstructure:"issueSeverity" env:"GITEE_ISSUE_SEVERITY" json:"issueSeverity"`
	IssuePriority        string `mapstructure:"issuePriority" env:"GITEE_ISSUE_PRIORITY" json:"issuePriority"`
	IssueComponent       string `mapstructure:"issueComponent" env:"GITEE_ISSUE_COMPONENT" json:"issueComponent"`
	IssueTypeBug         string `mapstructure:"issueTypeBug" env:"GITEE_ISSUE_TYPE_BUG" json:"issueTypeBug"`
	IssueTypeIncident    string `mapstructure:"issueTypeIncident" env:"GITEE_ISSUE_TYPE_INCIDENT" json:"issueTypeIncident"`
	IssueTypeRequirement string `mapstructure:"issueTypeRequirement" env:"GITEE_ISSUE_TYPE_REQUIREMENT" json:"issueTypeRequirement"`
}

func (GiteeConnection) TableName() string {
	return "_tool_gitee_connections"
}
