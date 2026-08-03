package piefed

// Piefed's real response includes a full my_user object, but mlmym's own
// save-settings flow (checked directly in its source) never reads it —
// it only checks for an error, which Piefed communicates via HTTP status
// rather than a body field. An empty struct here safely discards
// whatever Piefed actually returns.
type SaveUserSettingsResponse struct {
}
