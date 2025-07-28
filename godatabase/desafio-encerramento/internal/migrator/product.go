package migrator

import "app/internal/domain"

// NewMigratorProductDatabase returns a new MigratorProductToDatabase
func NewMigratorProductToDatabase(ld domain.LoaderProduct, rp domain.RepositoryProduct) (m *MigratorProductToDatabase) {
	m = &MigratorProductToDatabase{
		ld: ld,
		rp: rp,
	}
	return
}

// MigratorProductToDatabase is the implementation of the interface MigratorProduct
type MigratorProductToDatabase struct {
	// ld is the loader to load the data
	ld domain.LoaderProduct
	// rp is the repository to access the database
	rp domain.RepositoryProduct
}

// Migrate migrates the data from the a source to a destination
func (m *MigratorProductToDatabase) Migrate() (err error) {
	// load the data
	p, err := m.ld.Load()
	if err != nil {
		return
	}

	// save each customer
	for _, v := range p {
		err = m.rp.Save(&v)
		if err != nil {
			return
		}
	}

	return
}
