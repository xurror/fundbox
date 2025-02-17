package repositories

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAll() []string {
	return []string{"User A", "User B", "User C"}
}
