package helper

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/google/go-cmp/cmp"
)

func TestSnowflakeQuery(t *testing.T) {
	t.Run("Generate Snowflake id before 3 hour.", func(t *testing.T) {

		const given = 3
		const milli = 3600000

		var got = GenerateSnowflakeIdBefore(TimeUnit{hour: given}).Time()

		node, err := snowflake.NewNode(0)
		if err != nil {
			fmt.Println(err)
			return
		}

		id := node.Generate()

		var want = id.Time() - int64(given*milli)

		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v, given %v", got, want, given)
		}
	})

	t.Run("Generate Snowflake id before 24 hour.", func(t *testing.T) {

		const given = 24
		const milli = 3600000
		SF_ID := GenerateSnowflakeIdBefore(TimeUnit{hour: given})
		var got = SF_ID.Time()

		node, err := snowflake.NewNode(0)
		if err != nil {
			fmt.Println(err)
			return
		}

		id := node.Generate()

		var want = id.Time() - int64(given*milli)

		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v, given %v", got, want, given)
		} else {
			t.Log("\n\n\n Example Snowflake ID : ", SF_ID, " (You able to query greater than this one. it don't care node, step) \n\n")
		}
	})
}
