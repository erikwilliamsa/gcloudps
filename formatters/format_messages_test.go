package formatters_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/erikwilliamsa/gcloudps/formatters"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#FormatJSON", func() {

	Context("Given a JSON String", func() {
		s := `{"test": "data","isJson": true, "count":1, "subData":{ "sub": "sammicch"}}`
		It("Should return a formatted version of that string", func() {

			testdata := `{
							"test":"data",
							"isJson":true,
							"count":1,
							"subData":{
								"sub":"sammicch"
							}
						}`
			expected := new(bytes.Buffer)
			json.Indent(expected, []byte(testdata), "", "    ")

			r := FormatJSON(s)
			fmt.Println(r)
			Expect(r).To(Equal(expected.String()))
		})
	})

	Context("Given a string that is not JSON", func() {
		s := "Some string"
		It("Should return the string as is", func() {
			r := FormatJSON(s)
			Expect(r).To(Equal(s))
		})
	})

})
