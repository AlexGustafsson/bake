package state

import "sync"

type State struct {
	sync.Mutex
	Documents map[string]*Document
}

func CreateState() *State {
	return &State{
		Documents: make(map[string]*Document),
	}
}

func (state *State) CreateDocument(uri string, content string, version int) {
	document := CreateDocument(uri, content, version)
	state.Lock()
	defer state.Unlock()
	// TODO: handle collisions?
	state.Documents[uri] = document
}

func (state *State) RemoveDocument(uri string) {
	state.Lock()
	defer state.Unlock()
	delete(state.Documents, uri)
}
