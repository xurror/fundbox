package repositories

type FundRepository struct{}

func NewFundRepository() *FundRepository {
	return &FundRepository{}
}

func (r *FundRepository) GetAll() []string {
	return []string{"Fund A", "Fund B", "Fund C"}
}
