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

type Server struct {
	store map[string]*Item
}

// Public: store the key-value pair
func (s *Server) Set(key, value string) (err error) {
	s.store[key] = &Item{value: []byte(value)}
	return nil
}

// Public: store the key-value pair only if the key does not exist in the store
func (s *Server) Add(key, value string) (err error) {
	if !s.hasKey(key) {
		s.store[key] = &Item{value: []byte(value)}
	}
	return nil
}

// Public: replace the value of a key with another value
func (s *Server) Replace(key, value string) (status bool, err error) {
	if !s.hasKey(key) {
		return false, nil
	}
	s.store[key] = &Item{value: []byte(value)}
	return true, nil
}

// Public: append the value to an existing key value pair
func (s *Server) Append(key, value string) (status bool, err error) {
	if !s.hasKey(key) {
		return false, nil
	}
	existingItem, _ := s.Get(key)
	existingItem.value = append(existingItem.value, []byte(value)...)
	s.store[key] = existingItem
	return true, nil
}

// Public: prepend the value to an existing key value pair
func (s *Server) Prepend(key, value string) (status bool, err error) {
	if !s.hasKey(key) {
		return false, nil
	}
	existingItem, _ := s.Get(key)
	existingItem.value = append([]byte(value), existingItem.value...)
	s.store[key] = existingItem
	return true, nil
}

// Public: get the current data given a key
func (s *Server) Get(key string) (item *Item, err error) {
	v := s.store[key]
	return v, nil
}

// Private: check existence of the key
func (s *Server) hasKey(key string) bool {
	_, ok := s.store[key]
	return ok
}

func (s *Server) Delete(key string) (err error) {
	delete(s.store, key)
	return nil
}

func NewServer() *Server {
	return &Server{store: make(map[string]*Item)}
}
