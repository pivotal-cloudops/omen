package errands_test

import (
	"errors"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cloudops/omen/internal/errands"
	"github.com/pivotal-cloudops/omen/internal/errands/errandsfakes"
)

var _ = Describe("Toggle Errands", func() {
	var (
		es *errandsfakes.FakeErrandService
		et errands.ErrandToggler
		rp *errandsfakes.FakeReporter
	)

	BeforeEach(func() {
		es = &errandsfakes.FakeErrandService{}
		rp = &errandsfakes.FakeReporter{}
		et = errands.NewErrandToggler(es, rp)
	})

	It("can be constructed", func() {
		Expect(et).ToNot(BeNil())
	})

	Describe("specific products", func() {
		It("retrieves errand state for only specified products", func() {
			et.Execute([]string{"PEANUTS-and-butter", "almond-butter"})

			Expect(es.ListCallCount()).To(Equal(2))
			product1Id := es.ListArgsForCall(0)
			Expect(product1Id).To(Equal("PEANUTS-and-butter"))

			product2Id := es.ListArgsForCall(1)
			Expect(product2Id).To(Equal("almond-butter"))
		})

		It("fails with error if an invalid product is specified", func() {
			es.ListReturns(api.ErrandsListOutput{}, errors.New("product not found"))
			err := et.Execute([]string{"PEANUTS-and-butter"})

			Expect(err).To(HaveOccurred())
		})

		Context("valid product output", func() {
			errandServiceResponse := api.ErrandsListOutput{
				Errands: []api.Errand{
					{
						Name:       "errand1",
						PostDeploy: true,
						PreDelete:  true,
					},
					{
						Name:       "errand2",
						PostDeploy: false,
					},
					{
						Name:       "errand3",
						PostDeploy: "when-changed",
					},
					{
						Name:       "errand4",
						PostDeploy: "default",
					},
					{
						Name:      "errand5",
						PreDelete: true,
					},
					{
						Name:      "errand6",
						PreDelete: false,
					},
				},
			}
			BeforeEach(func() {
				es.ListReturns(errandServiceResponse, nil)
			})

			Describe("disable", func() {
				It("outputs current and desired state for post-deploy errands of affected products", func() {
					et.Disable().Execute([]string{"PEANUTS-and-butter"})
					output := ""
					for i := 0; i < rp.PrintReportCallCount(); i++ {
						text, err := rp.PrintReportArgsForCall(i)
						Expect(err).NotTo(HaveOccurred())
						output += text
					}

					Expect(output).To(ContainSubstring("Errands for PEANUTS-and-butter"))
					Expect(output).To(MatchRegexp("errand1\\s+enabled => disabled"))
					Expect(output).To(MatchRegexp("errand2\\s+disabled\\n"))
					Expect(output).To(MatchRegexp("errand3\\s+when-changed => disabled"))
					Expect(output).To(MatchRegexp("errand4\\s+default => disabled"))
					Expect(output).NotTo(ContainSubstring("errand5"))
				})

				It("only enables post-deploy errands not at desired state", func() {
					et.Disable().Execute([]string{"PEANUTS-and-butter"})

					Expect(es.SetStateCallCount()).To(Equal(3))

					for i := range []int{0, 1, 2} {
						isFirstRun := i == 0
						productName, errandName, postDeployState, preDeleteState := es.SetStateArgsForCall(i)
						Expect(productName).To(Equal("PEANUTS-and-butter"))
						Expect(postDeployState).To(BeFalse())

						if isFirstRun {
							Expect(errandName).To(Equal(fmt.Sprintf("errand%d", i+1)))
							Expect(preDeleteState).To(BeTrue())
						} else {
							Expect(errandName).To(Equal(fmt.Sprintf("errand%d", i+2)))
							Expect(preDeleteState).To(BeNil())
						}
					}
				})
			})

			Describe("enable", func() {
				It("outputs current and desired state for post-deploy errands of affected products", func() {
					et.Enable().Execute([]string{"PEANUTS-and-butter"})
					output := ""
					for i := 0; i < rp.PrintReportCallCount(); i++ {
						text, err := rp.PrintReportArgsForCall(i)
						Expect(err).NotTo(HaveOccurred())
						output += text
					}

					Expect(output).To(ContainSubstring("Errands for PEANUTS-and-butter"))
					Expect(output).To(MatchRegexp("errand1\\s+enabled\\n"))
					Expect(output).To(MatchRegexp("errand2\\s+disabled => enabled"))
					Expect(output).To(MatchRegexp("errand3\\s+when-changed => enabled"))
					Expect(output).To(MatchRegexp("errand4\\s+default => enabled"))
					Expect(output).NotTo(ContainSubstring("errand5"))
				})

				It("only enables post-deploy errands not at desired state", func() {
					et.Enable().Execute([]string{"PEANUTS-and-butter"})

					Expect(es.SetStateCallCount()).To(Equal(3))

					for i := range []int{0, 1, 2} {
						productName, errandName, postDeployState, preDeleteState := es.SetStateArgsForCall(i)
						Expect(productName).To(Equal("PEANUTS-and-butter"))
						Expect(postDeployState).To(BeTrue())
						Expect(errandName).To(Equal(fmt.Sprintf("errand%d", i+2)))
						Expect(preDeleteState).To(BeNil())
					}
				})
			})

			Describe("default", func() {
				It("outputs current and desired state for post-deploy errands of affected products", func() {
					et.Default().Execute([]string{"PEANUTS-and-butter"})
					output := ""
					for i := 0; i < rp.PrintReportCallCount(); i++ {
						text, err := rp.PrintReportArgsForCall(i)
						Expect(err).NotTo(HaveOccurred())
						output += text
					}

					Expect(output).To(ContainSubstring("Errands for PEANUTS-and-butter"))
					Expect(output).To(MatchRegexp("errand1\\s+enabled => default"))
					Expect(output).To(MatchRegexp("errand2\\s+disabled => default"))
					Expect(output).To(MatchRegexp("errand3\\s+when-changed => default"))
					Expect(output).To(MatchRegexp("errand4\\s+default"))
					Expect(output).NotTo(ContainSubstring("errand5"))
				})

				It("only enables post-deploy errands not at desired state", func() {
					et.Default().Execute([]string{"PEANUTS-and-butter"})

					Expect(es.SetStateCallCount()).To(Equal(3))

					for i := range []int{0, 1, 2} {
						productName, errandName, postDeployState, preDeleteState := es.SetStateArgsForCall(i)
						Expect(productName).To(Equal("PEANUTS-and-butter"))
						Expect(postDeployState).To(Equal("default"))
						Expect(errandName).To(Equal(fmt.Sprintf("errand%d", i+1)))

						if i == 0 {
							Expect(preDeleteState).To(BeTrue())
						} else {
							Expect(preDeleteState).To(BeNil())
						}
					}
				})
			})
		})
	})
})