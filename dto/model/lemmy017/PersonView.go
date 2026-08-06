package lemmy017

// PersonView matches Lemmy 0.17.2's real PersonViewSafe: person, counts.
// No is_admin field, unlike canonical.
type PersonView struct {
	Person Person           `json:"person" validate:"required"`
	Counts PersonAggregates `json:"counts" validate:"required"`
}
