// Code generated by ent, DO NOT EDIT.

package group

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeLocations holds the string denoting the locations edge name in mutations.
	EdgeLocations = "locations"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeLabels holds the string denoting the labels edge name in mutations.
	EdgeLabels = "labels"
	// EdgeDocuments holds the string denoting the documents edge name in mutations.
	EdgeDocuments = "documents"
	// EdgeInvitationTokens holds the string denoting the invitation_tokens edge name in mutations.
	EdgeInvitationTokens = "invitation_tokens"
	// EdgeNotifiers holds the string denoting the notifiers edge name in mutations.
	EdgeNotifiers = "notifiers"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "group_users"
	// LocationsTable is the table that holds the locations relation/edge.
	LocationsTable = "locations"
	// LocationsInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationsInverseTable = "locations"
	// LocationsColumn is the table column denoting the locations relation/edge.
	LocationsColumn = "group_locations"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "group_items"
	// LabelsTable is the table that holds the labels relation/edge.
	LabelsTable = "labels"
	// LabelsInverseTable is the table name for the Label entity.
	// It exists in this package in order to avoid circular dependency with the "label" package.
	LabelsInverseTable = "labels"
	// LabelsColumn is the table column denoting the labels relation/edge.
	LabelsColumn = "group_labels"
	// DocumentsTable is the table that holds the documents relation/edge.
	DocumentsTable = "documents"
	// DocumentsInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentsInverseTable = "documents"
	// DocumentsColumn is the table column denoting the documents relation/edge.
	DocumentsColumn = "group_documents"
	// InvitationTokensTable is the table that holds the invitation_tokens relation/edge.
	InvitationTokensTable = "group_invitation_tokens"
	// InvitationTokensInverseTable is the table name for the GroupInvitationToken entity.
	// It exists in this package in order to avoid circular dependency with the "groupinvitationtoken" package.
	InvitationTokensInverseTable = "group_invitation_tokens"
	// InvitationTokensColumn is the table column denoting the invitation_tokens relation/edge.
	InvitationTokensColumn = "group_invitation_tokens"
	// NotifiersTable is the table that holds the notifiers relation/edge.
	NotifiersTable = "notifiers"
	// NotifiersInverseTable is the table name for the Notifier entity.
	// It exists in this package in order to avoid circular dependency with the "notifier" package.
	NotifiersInverseTable = "notifiers"
	// NotifiersColumn is the table column denoting the notifiers relation/edge.
	NotifiersColumn = "group_id"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldCurrency,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Currency defines the type for the "currency" enum field.
type Currency string

// CurrencyUsd is the default value of the Currency enum.
const DefaultCurrency = CurrencyUsd

// Currency values.
const (
	CurrencyAed Currency = "aed"
	CurrencyAud Currency = "aud"
	CurrencyBgn Currency = "bgn"
	CurrencyBrl Currency = "brl"
	CurrencyCad Currency = "cad"
	CurrencyChf Currency = "chf"
	CurrencyCny Currency = "cny"
	CurrencyCzk Currency = "czk"
	CurrencyDkk Currency = "dkk"
	CurrencyEur Currency = "eur"
	CurrencyGbp Currency = "gbp"
	CurrencyHkd Currency = "hkd"
	CurrencyIdr Currency = "idr"
	CurrencyInr Currency = "inr"
	CurrencyJpy Currency = "jpy"
	CurrencyKrw Currency = "krw"
	CurrencyMxn Currency = "mxn"
	CurrencyNok Currency = "nok"
	CurrencyNzd Currency = "nzd"
	CurrencyPln Currency = "pln"
	CurrencyRmb Currency = "rmb"
	CurrencyRon Currency = "ron"
	CurrencyRub Currency = "rub"
	CurrencySar Currency = "sar"
	CurrencySek Currency = "sek"
	CurrencySgd Currency = "sgd"
	CurrencyThb Currency = "thb"
	CurrencyTry Currency = "try"
	CurrencyUsd Currency = "usd"
	CurrencyXag Currency = "xag"
	CurrencyXau Currency = "xau"
	CurrencyZar Currency = "zar"
)

func (c Currency) String() string {
	return string(c)
}

// CurrencyValidator is a validator for the "currency" field enum values. It is called by the builders before save.
func CurrencyValidator(c Currency) error {
	switch c {
	case CurrencyAed, CurrencyAud, CurrencyBgn, CurrencyBrl, CurrencyCad, CurrencyChf, CurrencyCny, CurrencyCzk, CurrencyDkk, CurrencyEur, CurrencyGbp, CurrencyHkd, CurrencyIdr, CurrencyInr, CurrencyJpy, CurrencyKrw, CurrencyMxn, CurrencyNok, CurrencyNzd, CurrencyPln, CurrencyRmb, CurrencyRon, CurrencyRub, CurrencySar, CurrencySek, CurrencySgd, CurrencyThb, CurrencyTry, CurrencyUsd, CurrencyXag, CurrencyXau, CurrencyZar:
		return nil
	default:
		return fmt.Errorf("group: invalid enum value for currency field: %q", c)
	}
}

// OrderOption defines the ordering options for the Group queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCurrency orders the results by the currency field.
func ByCurrency(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrency, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLocationsCount orders the results by locations count.
func ByLocationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLocationsStep(), opts...)
	}
}

// ByLocations orders the results by locations terms.
func ByLocations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLocationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByItemsCount orders the results by items count.
func ByItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemsStep(), opts...)
	}
}

// ByItems orders the results by items terms.
func ByItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLabelsCount orders the results by labels count.
func ByLabelsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLabelsStep(), opts...)
	}
}

// ByLabels orders the results by labels terms.
func ByLabels(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLabelsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDocumentsCount orders the results by documents count.
func ByDocumentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentsStep(), opts...)
	}
}

// ByDocuments orders the results by documents terms.
func ByDocuments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInvitationTokensCount orders the results by invitation_tokens count.
func ByInvitationTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInvitationTokensStep(), opts...)
	}
}

// ByInvitationTokens orders the results by invitation_tokens terms.
func ByInvitationTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInvitationTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNotifiersCount orders the results by notifiers count.
func ByNotifiersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotifiersStep(), opts...)
	}
}

// ByNotifiers orders the results by notifiers terms.
func ByNotifiers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotifiersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
	)
}
func newLocationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LocationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LocationsTable, LocationsColumn),
	)
}
func newItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
	)
}
func newLabelsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LabelsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LabelsTable, LabelsColumn),
	)
}
func newDocumentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DocumentsTable, DocumentsColumn),
	)
}
func newInvitationTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InvitationTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, InvitationTokensTable, InvitationTokensColumn),
	)
}
func newNotifiersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotifiersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NotifiersTable, NotifiersColumn),
	)
}
