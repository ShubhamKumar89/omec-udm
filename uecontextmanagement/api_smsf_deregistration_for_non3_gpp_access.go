// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uecontextmanagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeregistrationSmsfNon3gppAccess - delete SMSF registration for non 3GPP access
func HTTPDeregistrationSmsfNon3gppAccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
