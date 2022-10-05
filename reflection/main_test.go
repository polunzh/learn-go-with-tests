package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Zhang"},
			[]string{"Zhang"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Zhang", "Zheng Zhou"},
			[]string{"Zhang", "Zheng Zhou"},
		},
		{
			"struct with none string fields",
			struct {
				Name string
				Age  int
			}{"Zhang", 33},
			[]string{"Zhang"},
		},
		{
			"nested fields",
			Person{
				"Zhang",
				Profile{33, "Zheng Zhou"},
			},
			[]string{"Zhang", "Zheng Zhou"},
		},
		{
			"pointers to things",
			&Person{
				"Zhang",
				Profile{33, "Zheng Zhou"},
			},
			[]string{"Zhang", "Zheng Zhou"},
		},
		{
			"slices",
			[]Profile{
				{33, "Zheng Zhou"},
				{34, "Zheng Zhou"},
			},
			[]string{"Zheng Zhou", "Zheng Zhou"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Zheng Zhou"},
				{34, "Zheng Zhou"},
			},
			[]string{"Zheng Zhou", "Zheng Zhou"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Zheng Zhou"}
			aChannel <- Profile{34, "He Nan"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Zheng Zhou", "He Nan"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Zheng Zhou"}, Profile{34, "He Nan"}
		}

		var got []string
		want := []string{"Zheng Zhou", "He Nan"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, v := range haystack {
		if v == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain  %q but it didn't", haystack, needle)
	}
}
