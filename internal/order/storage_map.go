package order

// StorageOrderMap is a struct that represents an order storage
type StorageOrderMap struct {
	// db is the map where the orders are stored
	// - key: order id
	// - value: order
	db map[int]OrderDB

	// lastId is the next id to be assigned to an order
	lastId int
}

// NewStorageOrderMap returns a new instance of StorageOrderMap
func NewStorageOrderMap(db map[int]OrderDB, lastId int) (s *StorageOrderMap) {
	s = &StorageOrderMap{
		db:     db,
		lastId: lastId,
	}
	return
}

// FindById returns the order with the given Id
func (s *StorageOrderMap) FindById(id int) (o OrderDB, err error) {
	o, ok := s.db[id]
	if !ok {
		err = ErrStorageOrderNotFound
		return
	}

	return
}

// Add adds an order to the storage
func (s *StorageOrderMap) Add(o *OrderDB) (err error) {
	s.lastId++
	o.Id = s.lastId
	s.db[o.Id] = *o
	return
}

// Update updates the order with the given Id
func (s *StorageOrderMap) Update(o *OrderDB) (err error) {
	var exists bool
	var id int
	for ix := range s.db {
		if ix == o.Id {
			exists = true
			id = ix
			break
		}
	}

	if !exists {
		err = ErrStorageOrderNotFound
		return
	}

	s.db[id] = *o
	return
}