package duration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNil(t *testing.T) {
	assert.Nil(t, nil, "Nil is Nil")
}

func TestWith2EachNoConflict(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{250, 290})
	dl2 = append(dl2, Duration{450, 490})

	dlm := MergeDurations(dl1, dl2)
	assert.Equal(t, 4, len(dlm))

	assert.EqualValues(t, Duration{100, 200}, dlm[0])
	assert.EqualValues(t, Duration{250, 290}, dlm[1])
	assert.EqualValues(t, Duration{300, 400}, dlm[2])
	assert.EqualValues(t, Duration{450, 490}, dlm[3])

	dlm = MergeDurations(dl2, dl1)
	assert.Equal(t, 4, len(dlm))

	assert.EqualValues(t, Duration{100, 200}, dlm[0])
	assert.EqualValues(t, Duration{250, 290}, dlm[1])
	assert.EqualValues(t, Duration{300, 400}, dlm[2])
	assert.EqualValues(t, Duration{450, 490}, dlm[3])


}

func TestWith2_D1eatsD2(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{250, 290})
	dl2 = append(dl2, Duration{350, 390})

	dlm := MergeDurations(dl1, dl2)
	assert.Equal(t, 3, len(dlm))

	assert.EqualValues(t, Duration{100, 200}, dlm[0])
	assert.EqualValues(t, Duration{250, 290}, dlm[1])
	assert.EqualValues(t, Duration{300, 400}, dlm[2])

	dlm = MergeDurations(dl2, dl1)
	assert.Equal(t, 3, len(dlm))

	assert.EqualValues(t, Duration{100, 200}, dlm[0])
	assert.EqualValues(t, Duration{250, 290}, dlm[1])
	assert.EqualValues(t, Duration{300, 400}, dlm[2])

}

func TestWith2ExpandD1(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{150, 290})
	dl2 = append(dl2, Duration{350, 390})

	dlm := MergeDurations(dl1, dl2)
	assert.Equal(t, 2, len(dlm))

	assert.EqualValues(t, Duration{100, 290}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])

	dlm = MergeDurations(dl2, dl1)
	assert.Equal(t, 2, len(dlm))

	assert.EqualValues(t, Duration{100, 290}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])
}

func TestWith2DiffLen1(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{150, 290})
	dl2 = append(dl2, Duration{350, 390})
	dl2 = append(dl2, Duration{392, 396})

	dlm := MergeDurations(dl1, dl2)
	assert.Equal(t, 2, len(dlm))

	assert.EqualValues(t, Duration{100, 290}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])

	dlm = MergeDurations(dl2, dl1)
	assert.Equal(t, 2, len(dlm))

	assert.EqualValues(t, Duration{100, 290}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])
}

func TestWith2DiffLen2(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{150, 290})
	dl2 = append(dl2, Duration{350, 390})
	dl2 = append(dl2, Duration{492, 496})

	dlm := MergeDurations(dl1, dl2)
	assert.Equal(t, 3, len(dlm))

	assert.EqualValues(t, Duration{100, 290}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])
	assert.EqualValues(t, Duration{492, 496}, dlm[2])

	dlm = MergeDurations(dl2, dl1)
	assert.Equal(t, 3, len(dlm))

	assert.EqualValues(t, Duration{100, 290}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])
	assert.EqualValues(t, Duration{492, 496}, dlm[2])
}

func TestNoConflict(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{0, 75})
	dl2 = append(dl2, Duration{201, 299})
	dl2 = append(dl2, Duration{401, 496})

	dlm := HasConflict(dl1, dl2)
	assert.Equal(t, false, dlm)

	dlm = HasConflict(dl2, dl1)
	assert.Equal(t, false, dlm)


}

func TestWith2Equals(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{100, 200})
	dl2 = append(dl2, Duration{300, 400})

	dlm := MergeDurations(dl1, dl2)
	assert.Equal(t, 2, len(dlm))

	assert.EqualValues(t, Duration{100, 200}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])


	dlm = MergeDurations(dl2, dl1)
	assert.Equal(t, 2, len(dlm))

	assert.EqualValues(t, Duration{100, 200}, dlm[0])
	assert.EqualValues(t, Duration{300, 400}, dlm[1])
}

func TestConflictWith2Equals(t *testing.T) {
	var dl1 []Duration
	var dl2 []Duration
	dl1 = append(dl1, Duration{100, 200})
	dl1 = append(dl1, Duration{300, 400})

	dl2 = append(dl2, Duration{100, 200})
	dl2 = append(dl2, Duration{300, 400})

	dlm := HasConflict(dl1, dl2)
	assert.Equal(t, true, dlm)

	dlm = HasConflict(dl2, dl1)
	assert.Equal(t, true, dlm)


}

func benchmarkHasConflict(b *testing.B, count int) {
	dl1 := make([]Duration, count)
	dl2 := make([]Duration, count)

	for j := 0; j < b.N; j++ {
		for i := 0; i < count; i++ {

			dl1[i] = Duration{uint64(100 * i), uint64(100*i + 2)}
			dl2[i] = Duration{uint64(100*i) + 4, uint64(100*i+2) + 6}
		}
		HasConflict(dl1, dl2)
	}
}

func benchmarkMergeDurations(b *testing.B, count int) {
	// first we add all the durations
	dl1 := make([]Duration, count)
	dl2 := make([]Duration, count)

	for j := 0; j < b.N; j++ {
		for i := 0; i < count; i++ {

			dl1[i] = Duration{uint64(100 * i), uint64(100*i + 2)}
			dl2[i] = Duration{uint64(100*i) + 4, uint64(100*i+2) + 6}
		}
		MergeDurations(dl1, dl2)
	}

}
func BenchmarkMergeDurations10(b *testing.B) {

	benchmarkMergeDurations(b, 10)
}
func BenchmarkMergeDurations100(b *testing.B) {
	benchmarkMergeDurations(b, 100)
}

func BenchmarkMergeDurations1000(b *testing.B) {
	benchmarkMergeDurations(b, 1000)
}

func BenchmarkMergeDurations100000(b *testing.B) {
	benchmarkMergeDurations(b, 100000)
}

func BenchmarkMergeDurations1000000(b *testing.B) {
	benchmarkMergeDurations(b, 1000000)
}

func BenchmarkHasConflict10(b *testing.B) {
	benchmarkHasConflict(b, 10)
}

func BenchmarkHasConflict100(b *testing.B) {
	benchmarkHasConflict(b, 100)
}

func BenchmarkHasConflict1000(b *testing.B) {
	benchmarkHasConflict(b, 1000)
}

func BenchmarkHasConflict10000(b *testing.B) {
	benchmarkHasConflict(b, 10000)
}

func BenchmarkHasConflict100000(b *testing.B) {
	benchmarkHasConflict(b, 100000)
}

func BenchmarkHasConflict10000000(b *testing.B) {
	benchmarkHasConflict(b, 10000000)
}

func TestSort(t *testing.T) {
	var dl []Duration
	dl = append(dl, Duration{201, 299})
	dl = append(dl, Duration{0, 75})
	dl = append(dl, Duration{401, 496})
	Sort(dl)
	assert.EqualValues(t, Duration{0, 75}, dl[0])
	assert.EqualValues(t, Duration{201, 299}, dl[1])
	assert.EqualValues(t, Duration{401, 496}, dl[2])
}

func TestHasOverLap(t *testing.T) {
	var dl []Duration
	dl = append(dl, Duration{100, 200})
	dl = append(dl, Duration{300, 400})
	dl = append(dl, Duration{250, 290})
	dl = append(dl, Duration{450, 490})
	Sort(dl)
	b := HasOverLap(dl)
	assert.EqualValues(t, false, b)

	dl = append(dl, Duration{480, 495})
	b = HasOverLap(dl)
	assert.EqualValues(t, true, b)

	dl = append(dl, Duration{500, 495})
	b = HasOverLap(dl)
	assert.EqualValues(t, true, b)

}
