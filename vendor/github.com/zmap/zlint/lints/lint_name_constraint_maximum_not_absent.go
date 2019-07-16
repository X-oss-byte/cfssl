package lints

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

/************************************************************************
RFC 5280: 4.2.1.10
Within this profile, the minimum and maximum fields are not used with
any name forms, thus, the minimum MUST be zero, and maximum MUST be
absent.  However, if an application encounters a critical name
constraints extension that specifies other values for minimum or
maximum for a name form that appears in a subsequent certificate, the
application MUST either process these fields or reject the
certificate.
************************************************************************/

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type nameConstraintMax struct{}

func (l *nameConstraintMax) Initialize() error {
	return nil
}

func (l *nameConstraintMax) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.NameConstOID)
}

func (l *nameConstraintMax) Execute(c *x509.Certificate) *LintResult {
	for _, i := range c.PermittedDNSNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedDNSNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.PermittedDNSNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedEmailAddresses {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.PermittedIPAddresses {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedIPAddresses {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.PermittedDirectoryNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedDirectoryNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.PermittedEdiPartyNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedEdiPartyNames {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.PermittedRegisteredIDs {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedRegisteredIDs {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.PermittedX400Addresses {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	for _, i := range c.ExcludedX400Addresses {
		if i.Max != 0 {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_name_constraint_maximum_not_absent",
		Description:   "Within the name constraints name form, the maximum field is not used and therefore MUST be absent",
		Citation:      "RFC 5280: 4.2.1.10",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &nameConstraintMax{},
	})
}