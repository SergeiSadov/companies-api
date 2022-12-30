package commands

const (
	typeCreate   = "companies-create"
	typeUpdate   = "companies-update"
	typDelete    = "companies-delete"
	createOutbox = "INSERT INTO outbox (id, aggregateid, type, payload) VALUES (?,?,?,?)"
)
