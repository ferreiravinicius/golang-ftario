package mock

import (
	"github.com/florestario/core/entity"
)

// Reader Mock
type GenusReaderMock struct {
	MockError error
	MockGetGenusByName *entity.Genus
	MockFilterGenusByName []entity.Genus
}

func (reader *GenusReaderMock) GetGenusByName(name string) (*entity.Genus, error) {
	err := reader.MockError
	if err != nil {
		return nil, err
	}
	return reader.MockGetGenusByName, nil
}


func (reader *GenusReaderMock) FilterGenusByName(name string) ([]entity.Genus, error) {
	err := reader.MockError
	if err != nil {
		return nil, err
	}
	return reader.MockFilterGenusByName, nil
}

// Writer Mock
type GenusWriterMock struct {
	MockError error
	MockSaveGenus *entity.Genus
}

func (writer *GenusWriterMock) SaveGenus(genus *entity.Genus) (*entity.Genus, error) {
	err := writer.MockError
	if err != nil {
		return nil, err
	}
	return writer.MockSaveGenus, nil
}

// Genus Persistence Composite Mock
type GenusPersistenceMock struct {
	GenusReaderMock
	GenusWriterMock
}
