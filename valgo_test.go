package valgo_test

import (
	"encoding/json"
	"testing"

	"github.com/cohesivestack/valgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Is(valgo.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20)).
		Is(valgo.Number(17, "age").GreaterThan(18))

	require.False(t, v.Valid())

	out, err := json.Marshal(v.Error())

	require.NoError(t, err)
	assert.Equal(t, `{"age":["Age must be greater than \"18\""],"full_name":["Full name must have a length between \"4\" and \"20\""]}`, string(out))
}

func TestIs(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Is(valgo.String("Bob", "full_name").Not().Blank().OfLengthBetween(4, 20)).
		Is(valgo.Number(17, "age").GreaterThan(18)).
		Is(valgo.String("singl", "status").InSlice([]string{"married", "single"}))

	require.False(t, v.Valid())

	out, err := json.Marshal(v.Error())

	require.NoError(t, err)
	assert.Equal(t, `{"age":["Age must be greater than \"18\""],"full_name":["Full name must have a length between \"4\" and \"20\""],"status":["Status is not valid"]}`, string(out))
}

func TestNew(t *testing.T) {
	t.Parallel()

	val, err := valgo.New()
	require.NoError(t, err)

	assert.False(t, val.Is(valgo.Number(11, "month_day").LessOrEqualTo(10)).Valid())
}

func TestIn(t *testing.T) {
	t.Parallel()

	type Address struct {
		Name   string
		Street string
	}

	type Person struct {
		Name    string
		Address Address
	}

	p := Person{"Bob", Address{"", "1600 Amphitheatre Pkwy"}}

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.String(p.Name, "name").OfLengthBetween(4, 20)).
		In("address",
			valgo.Is(valgo.String(p.Address.Name, "name").Not().Blank()).
				Is(valgo.String(p.Address.Street, "street").Not().Blank()),
		)

	require.False(t, v.Valid())

	out, err := json.Marshal(v.Error())

	require.NoError(t, err)
	assert.Equal(t, `{"address.name":["Name can't be blank"],"name":["Name must have a length between \"4\" and \"20\""]}`, string(out))
}

func TestInRow(t *testing.T) {
	t.Parallel()

	type Address struct {
		Name   string
		Street string
	}

	type Person struct {
		Name      string
		Addresses []Address
	}

	p := Person{
		"Bob",
		[]Address{
			{"", "1600 Amphitheatre Pkwy"},
			{"Home", ""},
		},
	}

	v, err := valgo.New()
	require.NoError(t, err)

	v.Validate().Is(valgo.String(p.Name, "name").OfLengthBetween(4, 20))

	for i, a := range p.Addresses {
		v.InRow("addresses",
			i,
			valgo.Is(valgo.String(a.Name, "name").Not().Blank()).
				Is(valgo.String(a.Street, "street").Not().Blank()))
	}

	require.False(t, v.Valid())

	out, err := json.Marshal(v.Error())

	require.NoError(t, err)
	assert.Equal(t, `{"addresses[0].name":["Name can't be blank"],"addresses[1].street":["Street can't be blank"],"name":["Name must have a length between \"4\" and \"20\""]}`, string(out))
}

func TestCheck(t *testing.T) {
	t.Parallel()

	v, err := valgo.New()
	require.NoError(t, err)

	v.Check(valgo.String("", "full_name").Not().Blank().OfLengthBetween(4, 20))

	require.False(t, v.Valid())

	out, err := json.Marshal(v.Error())
	require.NoError(t, err)

	assert.Equal(t, `{"full_name":["Full name can't be blank","Full name must have a length between \"4\" and \"20\""]}`, string(out))
}
