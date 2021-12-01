package day1

import "testing"

func Test_Populate(t *testing.T) {
	t.Run("basic case, window populates correctly", func(t *testing.T) {
		w := window{}
		w.populate(1)
		w.populate(2)
		w.populate(3)

		if w.w[0] != 1 {
			t.Fail()
		}

		if w.w[1] != 2 {
			t.Fail()
		}

		if w.w[2] != 3 {
			t.Fail()
		}
	})

	t.Run("it moves the window a long", func(t *testing.T) {
		w := window{}
		w.populate(1)
		w.populate(2)
		w.populate(3)
		w.populate(5)

		if w.w[0] != 2 {
			t.Fail()
		}

		if w.w[1] != 3 {
			t.Fail()
		}

		if w.w[2] != 5 {
			t.Fail()
		}

		w.populate(6)
		w.populate(33)

		if w.w[0] != 5 {
			t.Fail()
		}

		if w.w[1] != 6 {
			t.Fail()
		}

		if w.w[2] != 33 {
			t.Fail()
		}

	})
}
