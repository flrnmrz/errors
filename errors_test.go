package errors

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestPackageNew(t *testing.T) {
	Convey("Given the error message 'oops'", t, func() {
		msg := "oops"
		Convey("When creating an error from it", func() {
			err := New(msg)
			Convey("Then the error's message should be 'oops", func() {
				So(err.Error(), ShouldEqual, msg)
			})
		})
	})
}

func TestPackageFmt(t *testing.T) {
	Convey("Given the format message 'v: %v' and the value 3", t, func() {
		msg := "v: %v"
		value := 3
		Convey("When creating an error from it", func() {
			err := Fmt(msg, value)
			Convey("Then the error's message should be 'v: 3'", func() {
				So(err.Error(), ShouldEqual, "v: 3")
			})
		})
	})
}

func TestPackageWithFieldNew(t *testing.T) {
	// not doing anything, yet
	Convey("Given the key 'k' and the value 3", t, func() {
		key := "k"
		value := 3
		Convey("When creating an error from it", func() {
			err := WithField(key, value).New("Failed")
			Convey("Then the error's message should be 'oops", func() {
				So(err.Error(), ShouldEqual, "Failed (k: 3)")
			})
		})
	})
}

func TestPackageWithFieldTwiceNew(t *testing.T) {
	// not doing anything, yet
	Convey("Given an error builder with an existing key value pair and the key 'k' and the value 3", t, func() {
		eb := WithField("a", "b")
		key := "k"
		value := 3
		Convey("When creating an error from it", func() {
			err := eb.WithField(key, value).New("Failed")
			Convey("Then the error's message should be 'oops", func() {
				So(err.Error(), ShouldEqual, "Failed (a: b, k: 3)")
			})
		})
	})
}

// add tests with wrapping stuff
