package formatter

import "testing"

func TestFormatCreateTable(t *testing.T) {
	input := `CREATE TABLE IF NOT EXISTS notes (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);`

	want := `CREATE TABLE IF NOT EXISTS notes (
    id                  BIGSERIAL           PRIMARY KEY,
    title               TEXT                NOT NULL,
    body                TEXT                NOT NULL DEFAULT '',
    created_at          TIMESTAMPTZ         NOT NULL DEFAULT now()
)
`

	got := Format(input)

	if got != want {
		t.Errorf("unexpected output:\n--- want ---\n%s\n--- got ---\n%s", want, got)
	}
}