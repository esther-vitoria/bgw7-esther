package migrator

import "app/internal/domain"

// NewMigratorSaleDatabase returns a new MigratorSaleToDatabase
func NewMigratorSaleToDatabase(ld domain.LoaderSale, rp domain.RepositorySale) (m *MigratorSaleToDatabase) {
	m = &MigratorSaleToDatabase{
		ld: ld,
		rp: rp,
	}
	return
}

// MigratorSaleToDatabase is the implementation of the interface MigratorSale
type MigratorSaleToDatabase struct {
	// ld is the loader to load the data
	ld domain.LoaderSale
	// rp is the repository to access the database
	rp domain.RepositorySale
}

// Migrate migrates the data from the a source to a destination
func (m *MigratorSaleToDatabase) Migrate() (err error) {
	// load the data
	s, err := m.ld.Load()
	if err != nil {
		return
	}

	// save each customer
	for _, v := range s {
		err = m.rp.Save(&v)
		if err != nil {
			return
		}
	}

	return
}
