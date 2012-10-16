package gocached

// ERROR\r\n" - means the client sent a nonexistent command name.
type CommandError struct {
}

// "CLIENT_ERROR <error>\r\n"
// means some sort of client error in the input line, i.e. the input
// doesn't conform to the protocol in some way. <error> is a
// human-readable error string.
type ClientError struct {
	message string
}

// - "SERVER_ERROR <error>\r\n"
// means some sort of server error prevents the server from carrying
// out the command. <error> is a human-readable error string. In cases
// of severe server errors, which make it impossible to continue
// serving the client (this shouldn't normally happen), the server will
// close the connection after sending the error line. This is the only
// case in which the server closes a connection to a client.
type ServerError struct {
	message string
}

type Item struct {
	value []byte
}

type Cache struct {
	store map[string]*Item
}

// Public: store the key-value pair
func (s *Cache) Set(key string, value []byte) (err error) {
	s.store[key] = &Item{value: value}
	return nil
}

// Public: store the key-value pair only if the key does not exist in the store
func (s *Cache) Add(key string, value []byte) (err error) {
	if !s.hasKey(key) {
		s.store[key] = &Item{value: value}
	}
	return nil
}

// Public: replace the value of a key with another value
func (s *Cache) Replace(key string, value []byte) (status bool, err error) {
	if !s.hasKey(key) {
		return false, nil
	}
	s.store[key] = &Item{value: value}
	return true, nil
}

// Public: append the value to an existing key value pair
func (s *Cache) Append(key string, value []byte) (status bool, err error) {
	if !s.hasKey(key) {
		return false, nil
	}
	existingItem, _ := s.Get(key)
	existingItem.value = append(existingItem.value, value...)
	s.store[key] = existingItem
	return true, nil
}

// Public: prepend the value to an existing key value pair
func (s *Cache) Prepend(key string, value []byte) (status bool, err error) {
	if !s.hasKey(key) {
		return false, nil
	}
	existingItem, _ := s.Get(key)
	existingItem.value = append(value, existingItem.value...)
	s.store[key] = existingItem
	return true, nil
}

// Public: get the current data given a key
func (s *Cache) Get(key string) (item *Item, err error) {
	v := s.store[key]
	return v, nil
}

// Private: check existence of the key
func (s *Cache) hasKey(key string) bool {
	_, ok := s.store[key]
	return ok
}

func (s *Cache) Delete(key string) (err error) {
	delete(s.store, key)
	return nil
}

func (s *Cache) Version() string {
	return "0.0.1"
}

func NewCache() *Cache {
	return &Cache{store: make(map[string]*Item)}
}
