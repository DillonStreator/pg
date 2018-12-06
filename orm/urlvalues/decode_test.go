package urlvalues_test

import (
	"time"

	"github.com/go-pg/pg/orm/urlvalues"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Filter struct {
	Field    string
	FieldNEQ string
	FieldLT  int8
	FieldLTE int16
	FieldGT  int32
	FieldGTE int64

	Multi    []string
	MultiNEQ []int

	Time time.Time

	Omit []byte `pg:"-"`
}

var _ = Describe("Decode", func() {
	It("Decodes struct from Values", func() {
		f := &Filter{}
		err := urlvalues.Decode(f, urlvalues.Values{
			"field":      {"one"},
			"field__neq": {"two"},
			"field__lt":  {"1"},
			"field__lte": {"2"},
			"field__gt":  {"3"},
			"field__gte": {"4"},

			"multi":      {"one", "two"},
			"multi__neq": {"3", "4"},

			"time": {"1970-01-01 00:00:00+00:00:00"},
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(f.Field).To(Equal("one"))
		Expect(f.FieldNEQ).To(Equal("two"))
		Expect(f.FieldLT).To(Equal(int8(1)))
		Expect(f.FieldLTE).To(Equal(int16(2)))
		Expect(f.FieldGT).To(Equal(int32(3)))
		Expect(f.FieldGTE).To(Equal(int64(4)))

		Expect(f.Multi).To(Equal([]string{"one", "two"}))
		Expect(f.MultiNEQ).To(Equal([]int{3, 4}))

		Expect(f.Time).To(BeTemporally("==", time.Unix(0, 0)))
	})
})
