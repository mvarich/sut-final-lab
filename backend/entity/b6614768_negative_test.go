package entity

import(
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"testing"
	"unittest/entity"
)

func TestEntity(t *testing.T){
	t.Run("Check valid",func(t *testing.T){
		g := NewGomegaWithT(t)
		employee := entity.Employees{
			Name:"Mint",
			Salary:100,
			EmployeeCode:"HR1024",
		}
		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}