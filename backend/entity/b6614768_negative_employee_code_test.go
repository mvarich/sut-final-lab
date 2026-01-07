package entity

import(
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
	"testing"
	"unittest/entity"
)

func TestEntitythree(t *testing.T){
	t.Run("Check invalid",func(t *testing.T){
		g := NewGomegaWithT(t)
		employee := entity.Employees{
			Name:"Mint",
			Salary:16000,
			EmployeeCode:"H9999",
		}
		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
	})
}