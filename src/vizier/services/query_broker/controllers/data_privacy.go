/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package controllers

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	pixie "px.dev/pixie/src/operator/apis/px.dev/v1alpha1"
)

func init() {
	pflag.String("data_access", "Full", "The data access level for queries. Options are 'Full' or 'Restricted'")
}

type vizierCachedDataPrivacy struct {
	dataAccess pixie.DataAccessLevel
}

// ShouldRedactSensitiveColumns returns whether we should redact sensitive columns or not.
func (dp *vizierCachedDataPrivacy) ShouldRedactSensitiveColumns(ctx context.Context) (bool, error) {
	return dp.dataAccess == pixie.DataAccessRestricted, nil
}

// CreateDataPrivacyManager creates a privacy manager for the namespace.
func CreateDataPrivacyManager(ns string) (DataPrivacy, error) {
	dataAccessStr := viper.GetString("data_access")
	dataAccess := pixie.DataAccessLevel(dataAccessStr)
	switch dataAccess {
	case pixie.DataAccessFull, pixie.DataAccessRestricted, pixie.DataAccessPIIRestricted:
		return &vizierCachedDataPrivacy{dataAccess}, nil
	default:
		return nil, fmt.Errorf("Invalid DataAccess: '%s'", dataAccessStr)
	}
}
