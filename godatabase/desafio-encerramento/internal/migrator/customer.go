package migrator

import (
	"app/internal/domain"
)

// NewMigratorCustomerDatabase returns a new MigratorCustomerToDatabase
func NewMigratorCustomerToDatabase(ld domain.LoaderCustomer, rp domain.RepositoryCustomer) (m *MigratorCustomerToDatabase) {
	m = &MigratorCustomerToDatabase{
		ld: ld,
		rp: rp,
	}
	return
}

// MigratorCustomerToDatabase is the implementation of the interface MigratorCustomer
type MigratorCustomerToDatabase struct {
	// ld is the loader to load the data
	ld domain.LoaderCustomer
	// rp is the repository to access the database
	rp domain.RepositoryCustomer
}

// Migrate migrates the data from the a source to a destination
func (m *MigratorCustomerToDatabase) Migrate() (err error) {
	// load the data
	c, err := m.ld.Load()
	if err != nil {
		return
	}

	// save each customer
	for _, v := range c {
		err = m.rp.Save(&v)
		if err != nil {
			return
		}
	}

	return
}
