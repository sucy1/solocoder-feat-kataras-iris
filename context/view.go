package context

import (
	"fmt"
	"io"
)

// ErrViewNotExist it's an error.
// It reports whether a template was not found in the parsed templates tree.
type ErrViewNotExist struct {
	Name     string
	IsLayout bool
	Data     any
}

// Error completes the `error` interface.
func (e ErrViewNotExist) Error() string {
	title := "template"
	if e.IsLayout {
		title = "layout"
	}
	return fmt.Sprintf("%s '%s' does not exist", title, e.Name)
}

// ViewEngine is the interface which all view engines should be implemented in order to be registered inside iris.
type ViewEngine interface {
	// Name returns the name of the engine.
	Name() string
	// Load should load the templates from the given FileSystem.
	Load() error
	// ExecuteWriter should execute a template by its filename with an optional layout and bindingData.
	ExecuteWriter(w io.Writer, filename string, layout string, bindingData any) error
	// Ext should return the final file extension (including the dot)
	// which this view engine is responsible to render.
	// If the filename extension on ExecuteWriter is empty then this is appended.
	Ext() string
	// Reload returns whether the engine reloads templates on each render.
	// When false, compiled templates are cached in memory with LRU strategy.
	// When true, templates are recompiled on each ExecuteWriter call.
	GetReload() bool
	// SetReload sets the reload mode.
	SetReload(reload bool)
	// GetMaxCache returns the maximum number of templates to keep in the LRU cache.
	// If MaxCache <= 0, LRU caching is disabled even when Reload is false.
	GetMaxCache() int
	// SetMaxCache sets the maximum number of templates to keep in the LRU cache.
	// If max <= 0, LRU caching is disabled and any existing cache is cleared.
	SetMaxCache(max int)
	// DisableCache disables LRU caching entirely. Equivalent to SetMaxCache(0).
	// This is useful when Reload is false but template result caching is still undesired.
	DisableCache()
}

// ViewEngineFuncer is an addition of a view engine,
// if a view engine implements that interface
// then iris can add some closed-relative iris functions
// like {{ url }}, {{ urlpath }} and {{ tr }}.
type ViewEngineFuncer interface {
	// AddFunc should adds a function to the template's function map.
	AddFunc(funcName string, funcBody any)
}
