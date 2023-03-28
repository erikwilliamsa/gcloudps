package cmd_test

import (
	. "github.com/erikwilliamsa/gcloudps/cmd"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subscriber #CheckRequiredFlags", func() {
	BeforeEach(func() {
		TopicName = ""
		ProjectName = ""
	})

	Context("When ProjectName is not set", func() {
		It("Should return an error", func() {
			TopicName = "test"
			ProjectName = ""
			err := CheckRequiredFlags(make(map[string]string))

			Expect(err).ToNot(BeNil())
		})
	})

	Context("When ProjectName is not set", func() {

		It("Should return an error", func() {
			ProjectName = "test"
			TopicName = ""
			err := CheckRequiredFlags(make(map[string]string))

			Expect(err).ToNot(BeNil())
		})
	})

	Context("When project and topic names are provided ", func() {

		It("Should not return an error", func() {
			ProjectName = "test"
			TopicName = "test"
			err := CheckRequiredFlags(make(map[string]string))

			Expect(err).To(BeNil())
		})
	})

	Context("When a map of flags with value is provided", func() {

		It("Should not return an error", func() {
			ProjectName = "test"
			TopicName = "test"
			flagval := "test"
			flags := make(map[string]string)
			flags["flag"] = flagval

			err := CheckRequiredFlags(flags)

			Expect(err).To(BeNil())
		})
	})

	Context("When a map of flags with an empty value is provided", func() {

		It("Should return an error", func() {
			ProjectName = "test"
			TopicName = "test"
			flagval := ""
			flags := make(map[string]string)
			flags["flag"] = flagval

			err := CheckRequiredFlags(flags)

			Expect(err).ToNot(BeNil())
		})
	})

	Context("When a map of flags with multipul values and any are empty", func() {

		It("Should return an error", func() {
			ProjectName = "test"
			TopicName = "test"
			flagval := "aa"
			flagval1 := ""
			flagval2 := "c"

			flags := make(map[string]string)
			flags["flag"] = flagval
			flags["flag1"] = flagval1
			flags["flag2"] = flagval2

			err := CheckRequiredFlags(flags)

			Expect(err).ToNot(BeNil())
		})
	})

})
