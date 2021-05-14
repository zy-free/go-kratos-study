package hashid

import (
	"encoding/json"
	"testing"
)

func TestID_MarshalJSON(t *testing.T) {
	// rbna67lx3b7e15lx => 1380
	user := struct {
		ID ID  `json:"id"`
		Name string `json:"name"`
	}{
		ID:1380,
		Name:"zhouyu",
	}

	b,_ := json.Marshal(user)
	if string(b) != `{"id":"rbna67lx3b7e15lx","name":"zhouyu"}` {
		t.Fatalf("json Marshal not equal")
	}

	user2 := struct {
		ID ID  `json:"id"`
		Name string `json:"name"`
	}{}
	s := `{"id":"rbna67lx3b7e15lx","name":"zhouyu"}`
	json.Unmarshal([]byte(s),&user2)
	if user2.ID != 1380{
		t.Fatalf("json Unmarshal not equal")
	}

}

