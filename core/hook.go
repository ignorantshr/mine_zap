package core

type HookFunc func(Entry) error

type hook struct {
	Core
	fns []HookFunc
}

func RegisterHooks(core Core, fns ...HookFunc) Core {
	return &hook{
		Core: core,
		fns:  append([]HookFunc{}, fns...),
	}
}

func (h *hook) Check(ent Entry, ce *CheckedEntry) *CheckedEntry {
	downstream := h.Core.Check(ent, ce) // add core to ce
	if downstream != nil {
		return downstream.AddCore(ent, h) // add self to ce
	}
	return ce
}

func (h *hook) Write(ent Entry, _ []Field) error {
	for _, fn := range h.fns {
		if err := fn(ent); err != nil {
			return err
		}
	}
	return nil
}
