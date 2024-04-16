package iteration

import "testing"

const count = 5

func BenchmarkRepeat(b *testing.B) {

    for i := 0; i < b.N; i++ {
        Repeat("a");
    }
}

func TestRepeat(t *testing.T) {
    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }
}

func Repeat(character string) string {
    var repeated string

    for i := 0; i < count; i++ {
        repeated += character
    }

    return repeated
}