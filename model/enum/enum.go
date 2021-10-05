package enum

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	constantError "github.com/vaniairnanda/send-later/model/constant/error"
	"strings"
)

type Status int

const (
	NEEDS_APPROVAL         Status = 1
	APPROVED               Status = 2
	PROCESSING             Status = 3
	EXPIRED                Status = 4
	SUCCESSFULLY_DISBURSED Status = 5
)


var statuses = []string{"", "NEEDS_APPROVAL", "APPROVED", "PROCESSING", "EXPIRED", "SUCCESSFULLY_DISBURSED"}

func (s Status) String() string {
	return statuses[s]
}

func (s Status) MarshalJSON() ([]byte, error)  {
	return json.Marshal(s.String())
}

func (s *Status) UnmarshalJSON(b []byte) error {
	var strAction = ""
	err := json.Unmarshal(b, &strAction)
	if err != nil {
		return err
	}
	*s, err = StatusFromString(strAction)
	return err
}

// Scan implements the Scanner interface.
func (s *Status) Scan(value interface{}) error {
	sqlData := &sql.NullString{}
	err := sqlData.Scan(value)
	if err != nil || !sqlData.Valid {
		return err
	}

	*s, err = StatusFromString(sqlData.String)
	return err
}

// Value implements the driver Valuer interface.
func (s Status) Value() (driver.Value, error) {
	return s.String(), nil
}


func StatusFromString(str string) (Status, error) {
	lowerStr := strings.ToLower(str)
	for i, j := 0, len(statuses)-1; i <= j; i, j = i+1, j-1 {
		if strings.ToLower(statuses[i]) == lowerStr {
			return Status(i), nil
		}
		if strings.ToLower(statuses[j]) == lowerStr {
			return Status(j), nil
		}
	}
	return -1, constantError.ErrorInvalidEnum
}
