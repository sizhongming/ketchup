package store

import (
	"fmt"

	"github.com/ketchuphq/ketchup/proto/ketchup/models"
	"github.com/ketchuphq/ketchup/proto/ketchup/packages"
)

// ErrParsingConfig is returned when theme.json cannot be parsed
var ErrParsingConfig = fmt.Errorf("error parsing theme.json")

// ThemeStore is a interface to support multiple theme backends.
type ThemeStore interface {
	// List all themes in the store
	List() ([]*models.Theme, error)

	// Add a theme from a theme file
	Add(*models.Theme) error

	AddPackage(*packages.Package) error

	// Get a theme by name from the store and the current ref
	Get(string) (Theme, error)

	// Get an asset from the store
	GetAsset(asset string) (*models.ThemeAsset, error)
}

type Theme interface {
	// Ref returns the current ref of the theme, if valid
	Ref() (string, bool)

	// Proto returns the underlying theme proto
	Proto() *models.Theme

	// GetTemplate returns the given template. Return nil, nil if template does not exist.
	GetTemplate(templateName string) (*models.ThemeTemplate, error)

	// GetAsset returns the given asset. Return nil, nil if asset does not exist.
	GetAsset(assetName string) (*models.ThemeAsset, error)
}
