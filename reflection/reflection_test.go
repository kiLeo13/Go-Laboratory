package main

import (
    "reflect"
    "testing"
)

func TestWalk(t *testing.T) {

    t.Run("with maps", func(t *testing.T) {
        aMap := map[string]string{
            "Cow":     "Moo",
            "Sheep":   "Baa",
        }

        var got[]string
        walk(aMap, func(input string) {
            got = append(got, input)
        })

        assertContains(t, got, "Moo")
        assertContains(t, got, "Baa")
    })

    t.Run("with channels", func(t *testing.T) {
        aChannel := make(chan Profile)

        go func() {
            aChannel <- Profile{33, "Berlin"}
            aChannel <- Profile{34, "Katowice"}
            close(aChannel)
        }()

        var got []string
        want := []string{"Berlin", "Katowice"}

        walk(aChannel, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v, want %v", got, want)
        }
    })

    t.Run("with function", func(t *testing.T) {
        aFunction := func() (Profile, Profile) {
            return Profile{33, "Berlin"}, Profile{34, "Katowice"}
        }

        var got []string
        want := []string{"Berlin", "Katowice"}

        walk(aFunction, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v, want %v", got, want)
        }
    })
}

type Person struct {
    Name    string
    Profile Profile
}

type Profile struct {
    Age  int
    City string
}

func walk(x interface{}, fn func(input string)) {
    val := getValue(x)

    walkValue := func(value reflect.Value) {
        walk(value.Interface(), fn)
    }

    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        for i := 0; i < val.NumField(); i++ {
            walkValue(val.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i := 0; i < val.Len(); i++ {
            walkValue(val.Index(i))
        }
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walkValue(val.MapIndex(key))
        }
    case reflect.Chan:
        for {
            if v, ok := val.Recv(); ok {
                walkValue(v)
            } else {
                break
            }
        }
    case reflect.Func:
        valFnResult := val.Call(nil)
        for _, res := range valFnResult {
            walkValue(res)
        }
    }
}

func getValue(x interface{}) reflect.Value {
    val := reflect.ValueOf(x)

    if val.Kind() == reflect.Pointer {
        val = val.Elem()
    }

    return val
}

func assertContains(t testing.TB, haystack []string, needle string) {
    t.Helper()
    contains := false

    for _, x := range haystack {
        if x == needle {
            contains = true
        }
    }
    
    if !contains {
        t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
    }
}