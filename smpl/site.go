package smpl

import (
	"fmt"
	"net/url"
	"time"
)

// Site represents either a single site or a MultiSite (site with subdomains)
type Site struct {
	Name      string
	Domain    string
	ProxyURL  *url.URL
	Root      *Site
	Children  []*Site
	MultiSite bool
	Created   time.Time
}

// NewSite creates and returns a pointer to a site that is initialized but
// not yet associated to a server or root site
func NewSite(name string, domain string, proxyURL *url.URL) *Site {
	return &Site{
		Name:     name,
		Domain:   domain,
		ProxyURL: proxyURL,
		Created:  time.Now(),
	}
}

// AddSubSite attempts to add a child to the site and returns an error if it fails
func (s *Site) AddSubSite(newSite *Site) error {
	if s.Root == nil {
		return fmt.Errorf("%v has no root site defined", s.Name)
	}
	if !s.IsRoot() {
		return fmt.Errorf("%v site is not a root", s.Name)
	}
	if !s.MultiSite {
		return fmt.Errorf("%v is not a MultiSite", s.Name)
	}
	if s.HasChildNamed(newSite.Name) {
		return fmt.Errorf("%v is already a subsite", newSite.Name)
	}
	s.Children = append(s.Children, newSite)

	return nil
}

// IsRoot returns whether a site is a root site.
func (s Site) IsRoot() bool {
	return s.Root != nil && s.Root.Name == s.Name && s.Root.Created == s.Created
}

// HasChildNamed returns whether the site has a child site with the given name
func (s Site) HasChildNamed(name string) bool {
	for _, c := range s.Children {
		if c.Name == name {
			return true
		}
	}
	return false
}
