package warehouse

type StorageWarehouseMap struct {
	// db is the database of warehouses (key: id, value: warehouse)
	db map[int]WarehouseDB

	// lastId is the next id to be assigned to a warehouse
	lastId int
}

// NewStorageWarehouseMap returns a new instance of StorageWarehouseMap
func NewStorageWarehouseMap(db map[int]WarehouseDB, lastId int) (s *StorageWarehouseMap) {
	s = &StorageWarehouseMap{
		db:     db,
		lastId: lastId,
	}
	return
}

// FindById returns the warehouse with the given Id
func (s *StorageWarehouseMap) FindById(id int) (w WarehouseDB, err error) {
	w, ok := s.db[id]
	if !ok {
		err = ErrStorageWarehouseNotFound
		return
	}

	return
}

// FindByName returns the warehouse with the given name
func (s *StorageWarehouseMap) FindByName(name string) (w WarehouseDB, err error) {
	var exists bool
	var id int
	for ix, wh := range s.db {
		if wh.Name == name {
			exists = true
			id = ix
		}
	}

	if !exists {
		err = ErrStorageWarehouseNotFound
		return
	}

	w = s.db[id]
	return
}

// Add adds a warehouse to the storage
func (s *StorageWarehouseMap) Add(w *WarehouseDB) (err error) {
	s.lastId++
	w.Id = s.lastId
	s.db[s.lastId] = *w
	return
}

// Update updates the warehouse with the given Id
func (s *StorageWarehouseMap) Update(w *WarehouseDB) (err error) {
	var exists bool
	var id int
	for ix := range s.db {
		if ix == w.Id {
			exists = true
			id = ix
		}
	}

	if !exists {
		err = ErrStorageWarehouseNotFound
		return
	}

	s.db[id] = *w
	return
}