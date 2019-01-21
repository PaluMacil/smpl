package smpl

import (
	"net/url"
	"testing"
)

func TestAddSubSite(t *testing.T) {
	proxyURL, _ := url.Parse("127.0.0.1:3033")
	site := NewSite("Example Site", "example.com", proxyURL)
	subSite := NewSite("Example SubSite", "example.com", proxyURL)

	err := site.AddSubSite(subSite)
	if err == nil {
		t.Errorf("was able to add %v to %v despite no root sites defined", subSite.Name, site.Name)
	}
	site.Root = site
	err = site.AddSubSite(subSite)
	if err == nil {
		t.Errorf("was able to add %v to %v despite not being a multisite", subSite.Name, site.Name)
	}
	site.MultiSite = true
	err = site.AddSubSite(subSite)
	if err != nil {
		t.Errorf("could not add %v to %v: %v", subSite.Name, site.Name, err)
	}
	err = site.AddSubSite(subSite)
	if err == nil {
		t.Errorf("was able to add %v to %v despite it already being added", subSite.Name, site.Name)
	}
}
