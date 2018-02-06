package gocoin

import "strings"

const (
	// ScopeAccountRead Access to read accounts data.
	ScopeAccountRead = "account_read"

	// ScopeUserRead Access to read users data.
	ScopeUserRead = "user_read"

	// ScopeUserReadWrite Access to read and write user data. Cannot modify password.
	ScopeUserReadWrite = "user_read_write"

	// ScopeUserPasswordWrite Can update user password. Must be combined with user_read_write to access other attributes.
	ScopeUserPasswordWrite = "user_password_write"

	// ScopeMerchantRead Can read all merchant info associated with a user.
	ScopeMerchantRead = "merchant_read"

	// ScopeMerchantReadWrite Can read and write to all merchants associated with a user.
	ScopeMerchantReadWrite = "merchant_read_write"

	// ScopeMerchantSettingsRead Can read all merchant settings associated with a user.
	ScopeMerchantSettingsRead = "merchant_settings_read"

	// ScopeMerchantSettingsReadWrite Can read and write to all merchant settings associated with a user.
	ScopeMerchantSettingsReadWrite = "merchant_settings_read_write"

	// ScopeInvoiceRead Can read all invoice info.
	ScopeInvoiceRead = "invoice_read"

	// ScopeInvoiceWrite Can create an invoice (but cannot read an existing one).
	ScopeInvoiceWrite = "invoice_write"

	// ScopeInvoiceReadWrite Can read and write to invoices.
	ScopeInvoiceReadWrite = "invoice_read_write"
)

// Scopes takes a list of scopes to use and returns the correct Gocoin API scopes string format
func Scopes(scopes ...string) string {
	return strings.Join(scopes, " ")
}
