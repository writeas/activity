//
package streams

import (
	"github.com/writeas/activity/vocab"
	"net/url"
	"time"
)

// Represents an individual person. This is a convenience wrapper of a type with the same name in the vocab package. Accessing it with the Raw function allows direct manipulaton of the object, and does not provide the same integrity guarantees as this package.
type Person struct {
	// The raw type from the vocab package
	raw *vocab.Person
}

// Raw returns the vocab type for manaual manipulation. Note that manipulating the underlying type to be in an inconsistent state may cause this convenience type's methods to later fail.
func (t *Person) Raw() (n *vocab.Person) {
	return t.raw

}

// Serialize turns this object into a map[string]interface{}.
func (t *Person) Serialize() (m map[string]interface{}, err error) {
	return t.raw.Serialize()

}

// GetAltitude attempts to get this 'altitude' property as a float64. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetAltitude() (r Resolution, k float64) {
	r = Unresolved
	handled := false
	if t.raw.IsAltitude() {
		k = t.raw.GetAltitude()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsAltitudeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasAltitude returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasAltitude() (p Presence) {
	p = NoPresence
	if t.raw.IsAltitude() {
		p = ConvenientPresence
	} else if t.raw.IsAltitudeIRI() {
		p = RawPresence
	}
	return

}

// SetAltitude sets the value for property 'altitude'.
func (t *Person) SetAltitude(k float64) {
	t.raw.SetAltitude(k)

}

// LenAttachment returns the number of values this property contains. Each index be used with HasAttachment to determine if ResolveAttachment is safe to call or if raw handling would be needed.
func (t *Person) LenAttachment() (idx int) {
	return t.raw.AttachmentLen()

}

// ResolveAttachment passes the actual concrete type to the resolver for handing property attachment. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveAttachment(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsAttachmentObject(idx) {
		handled, err = r.dispatch(t.raw.GetAttachmentObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsAttachmentLink(idx) {
		handled, err = r.dispatch(t.raw.GetAttachmentLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsAttachmentIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasAttachment returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasAttachment(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsAttachmentObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAttachmentLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAttachmentIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendAttachment appends an 'Object' typed value.
func (t *Person) AppendAttachment(i vocab.ObjectType) {
	t.raw.AppendAttachmentObject(i)

}

// PrependAttachment prepends an 'Object' typed value.
func (t *Person) PrependAttachment(i vocab.ObjectType) {
	t.raw.PrependAttachmentObject(i)

}

// AppendAttachmentLink appends a 'Link' typed value.
func (t *Person) AppendAttachmentLink(i vocab.LinkType) {
	t.raw.AppendAttachmentLink(i)

}

// PrependAttachmentLink prepends a 'Link' typed value.
func (t *Person) PrependAttachmentLink(i vocab.LinkType) {
	t.raw.PrependAttachmentLink(i)

}

// LenAttributedTo returns the number of values this property contains. Each index be used with HasAttributedTo to determine if GetAttributedTo is safe to call or if raw handling would be needed.
func (t *Person) LenAttributedTo() (idx int) {
	return t.raw.AttributedToLen()

}

// GetAttributedTo attempts to get this 'attributedTo' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetAttributedTo(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsAttributedToIRI(idx) {
		k = t.raw.GetAttributedToIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsAttributedToObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsAttributedToLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendAttributedTo appends the value for property 'attributedTo'.
func (t *Person) AppendAttributedTo(k *url.URL) {
	t.raw.AppendAttributedToIRI(k)

}

// PrependAttributedTo prepends the value for property 'attributedTo'.
func (t *Person) PrependAttributedTo(k *url.URL) {
	t.raw.PrependAttributedToIRI(k)

}

// RemoveAttributedTo deletes the value from the specified index for property 'attributedTo'.
func (t *Person) RemoveAttributedTo(idx int) {
	t.raw.RemoveAttributedToIRI(idx)

}

// HasAttributedTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasAttributedTo(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsAttributedToIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAttributedToLink(idx) {
		p = RawPresence
	} else if t.raw.IsAttributedToIRI(idx) {
		p = RawPresence
	}
	return

}

// LenAudience returns the number of values this property contains. Each index be used with HasAudience to determine if GetAudience is safe to call or if raw handling would be needed.
func (t *Person) LenAudience() (idx int) {
	return t.raw.AudienceLen()

}

// GetAudience attempts to get this 'audience' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetAudience(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsAudienceIRI(idx) {
		k = t.raw.GetAudienceIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsAudienceObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsAudienceLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendAudience appends the value for property 'audience'.
func (t *Person) AppendAudience(k *url.URL) {
	t.raw.AppendAudienceIRI(k)

}

// PrependAudience prepends the value for property 'audience'.
func (t *Person) PrependAudience(k *url.URL) {
	t.raw.PrependAudienceIRI(k)

}

// RemoveAudience deletes the value from the specified index for property 'audience'.
func (t *Person) RemoveAudience(idx int) {
	t.raw.RemoveAudienceIRI(idx)

}

// HasAudience returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasAudience(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsAudienceIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsAudienceLink(idx) {
		p = RawPresence
	} else if t.raw.IsAudienceIRI(idx) {
		p = RawPresence
	}
	return

}

// LenContent returns the number of values this property contains. Each index be used with HasContent to determine if GetContent is safe to call or if raw handling would be needed.
func (t *Person) LenContent() (idx int) {
	return t.raw.ContentLen()

}

// GetContent attempts to get this 'content' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetContent(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsContentString(idx) {
		k = t.raw.GetContentString(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsContentLangString(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsContentIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendContent appends the value for property 'content'.
func (t *Person) AppendContent(k string) {
	t.raw.AppendContentString(k)

}

// PrependContent prepends the value for property 'content'.
func (t *Person) PrependContent(k string) {
	t.raw.PrependContentString(k)

}

// RemoveContent deletes the value from the specified index for property 'content'.
func (t *Person) RemoveContent(idx int) {
	t.raw.RemoveContentString(idx)

}

// HasContent returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasContent(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsContentString(idx) {
		p = ConvenientPresence
	} else if t.raw.IsContentLangString(idx) {
		p = RawPresence
	} else if t.raw.IsContentIRI(idx) {
		p = RawPresence
	}
	return

}

// ContentLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Person) ContentLanguages() (l []string) {
	return t.raw.ContentMapLanguages()

}

// GetContentMap retrieves the value of 'content' for the specified language, or an empty string if it does not exist
func (t *Person) GetContentForLanguage(l string) (v string) {
	return t.raw.GetContentMap(l)

}

// SetContentForLanguage sets the value of 'content' for the specified language
func (t *Person) SetContentForLanguage(l string, v string) {
	t.raw.SetContentMap(l, v)

}

// LenContext returns the number of values this property contains. Each index be used with HasContext to determine if ResolveContext is safe to call or if raw handling would be needed.
func (t *Person) LenContext() (idx int) {
	return t.raw.ContextLen()

}

// ResolveContext passes the actual concrete type to the resolver for handing property context. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveContext(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsContextObject(idx) {
		handled, err = r.dispatch(t.raw.GetContextObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsContextLink(idx) {
		handled, err = r.dispatch(t.raw.GetContextLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsContextIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasContext returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasContext(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsContextObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsContextLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsContextIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendContext appends an 'Object' typed value.
func (t *Person) AppendContext(i vocab.ObjectType) {
	t.raw.AppendContextObject(i)

}

// PrependContext prepends an 'Object' typed value.
func (t *Person) PrependContext(i vocab.ObjectType) {
	t.raw.PrependContextObject(i)

}

// AppendContextLink appends a 'Link' typed value.
func (t *Person) AppendContextLink(i vocab.LinkType) {
	t.raw.AppendContextLink(i)

}

// PrependContextLink prepends a 'Link' typed value.
func (t *Person) PrependContextLink(i vocab.LinkType) {
	t.raw.PrependContextLink(i)

}

// LenName returns the number of values this property contains. Each index be used with HasName to determine if GetName is safe to call or if raw handling would be needed.
func (t *Person) LenName() (idx int) {
	return t.raw.NameLen()

}

// GetName attempts to get this 'name' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetName(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsNameString(idx) {
		k = t.raw.GetNameString(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsNameLangString(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsNameIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendName appends the value for property 'name'.
func (t *Person) AppendName(k string) {
	t.raw.AppendNameString(k)

}

// PrependName prepends the value for property 'name'.
func (t *Person) PrependName(k string) {
	t.raw.PrependNameString(k)

}

// RemoveName deletes the value from the specified index for property 'name'.
func (t *Person) RemoveName(idx int) {
	t.raw.RemoveNameString(idx)

}

// HasName returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasName(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsNameString(idx) {
		p = ConvenientPresence
	} else if t.raw.IsNameLangString(idx) {
		p = RawPresence
	} else if t.raw.IsNameIRI(idx) {
		p = RawPresence
	}
	return

}

// NameLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Person) NameLanguages() (l []string) {
	return t.raw.NameMapLanguages()

}

// GetNameMap retrieves the value of 'name' for the specified language, or an empty string if it does not exist
func (t *Person) GetNameForLanguage(l string) (v string) {
	return t.raw.GetNameMap(l)

}

// SetNameForLanguage sets the value of 'name' for the specified language
func (t *Person) SetNameForLanguage(l string, v string) {
	t.raw.SetNameMap(l, v)

}

// GetEndTime attempts to get this 'endTime' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetEndTime() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsEndTime() {
		k = t.raw.GetEndTime()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsEndTimeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasEndTime returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasEndTime() (p Presence) {
	p = NoPresence
	if t.raw.IsEndTime() {
		p = ConvenientPresence
	} else if t.raw.IsEndTimeIRI() {
		p = RawPresence
	}
	return

}

// SetEndTime sets the value for property 'endTime'.
func (t *Person) SetEndTime(k time.Time) {
	t.raw.SetEndTime(k)

}

// LenGenerator returns the number of values this property contains. Each index be used with HasGenerator to determine if ResolveGenerator is safe to call or if raw handling would be needed.
func (t *Person) LenGenerator() (idx int) {
	return t.raw.GeneratorLen()

}

// ResolveGenerator passes the actual concrete type to the resolver for handing property generator. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveGenerator(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsGeneratorObject(idx) {
		handled, err = r.dispatch(t.raw.GetGeneratorObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsGeneratorLink(idx) {
		handled, err = r.dispatch(t.raw.GetGeneratorLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsGeneratorIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasGenerator returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasGenerator(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsGeneratorObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsGeneratorLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsGeneratorIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendGenerator appends an 'Object' typed value.
func (t *Person) AppendGenerator(i vocab.ObjectType) {
	t.raw.AppendGeneratorObject(i)

}

// PrependGenerator prepends an 'Object' typed value.
func (t *Person) PrependGenerator(i vocab.ObjectType) {
	t.raw.PrependGeneratorObject(i)

}

// AppendGeneratorLink appends a 'Link' typed value.
func (t *Person) AppendGeneratorLink(i vocab.LinkType) {
	t.raw.AppendGeneratorLink(i)

}

// PrependGeneratorLink prepends a 'Link' typed value.
func (t *Person) PrependGeneratorLink(i vocab.LinkType) {
	t.raw.PrependGeneratorLink(i)

}

// LenIcon returns the number of values this property contains. Each index be used with HasIcon to determine if ResolveIcon is safe to call or if raw handling would be needed.
func (t *Person) LenIcon() (idx int) {
	return t.raw.IconLen()

}

// ResolveIcon passes the actual concrete type to the resolver for handing property icon. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveIcon(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsIconImage(idx) {
		handled, err = r.dispatch(t.raw.GetIconImage(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsIconLink(idx) {
		s = RawResolutionNeeded
	} else if t.raw.IsIconIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// AppendIcon appends the value for property 'icon'.
func (t *Person) AppendIcon(i vocab.ImageType) {
	t.raw.AppendIconImage(i)

}

// PrependIcon prepends the value for property 'icon'.
func (t *Person) PrependIcon(i vocab.ImageType) {
	t.raw.PrependIconImage(i)

}

// RemoveIcon deletes the value from the specified index for property 'icon'.
func (t *Person) RemoveIcon(idx int) {
	t.raw.RemoveIconImage(idx)

}

// HasIcon returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasIcon(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsIconImage(idx) {
		p = ConvenientPresence
	} else if t.raw.IsIconLink(idx) {
		p = RawPresence
	} else if t.raw.IsIconIRI(idx) {
		p = RawPresence
	}
	return

}

// GetId attempts to get this 'id' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetId() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasId() {
		k = t.raw.GetId()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasId returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasId() (p Presence) {
	p = NoPresence
	if t.raw.HasId() {
		p = ConvenientPresence
	}
	return

}

// SetId sets the value for property 'id'.
func (t *Person) SetId(k *url.URL) {
	t.raw.SetId(k)

}

// LenImage returns the number of values this property contains. Each index be used with HasImage to determine if ResolveImage is safe to call or if raw handling would be needed.
func (t *Person) LenImage() (idx int) {
	return t.raw.ImageLen()

}

// ResolveImage passes the actual concrete type to the resolver for handing property image. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveImage(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsImageImage(idx) {
		handled, err = r.dispatch(t.raw.GetImageImage(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsImageLink(idx) {
		s = RawResolutionNeeded
	} else if t.raw.IsImageIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// AppendImage appends the value for property 'image'.
func (t *Person) AppendImage(i vocab.ImageType) {
	t.raw.AppendImageImage(i)

}

// PrependImage prepends the value for property 'image'.
func (t *Person) PrependImage(i vocab.ImageType) {
	t.raw.PrependImageImage(i)

}

// RemoveImage deletes the value from the specified index for property 'image'.
func (t *Person) RemoveImage(idx int) {
	t.raw.RemoveImageImage(idx)

}

// HasImage returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasImage(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsImageImage(idx) {
		p = ConvenientPresence
	} else if t.raw.IsImageLink(idx) {
		p = RawPresence
	} else if t.raw.IsImageIRI(idx) {
		p = RawPresence
	}
	return

}

// LenInReplyTo returns the number of values this property contains. Each index be used with HasInReplyTo to determine if GetInReplyTo is safe to call or if raw handling would be needed.
func (t *Person) LenInReplyTo() (idx int) {
	return t.raw.InReplyToLen()

}

// GetInReplyTo attempts to get this 'inReplyTo' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetInReplyTo(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsInReplyToIRI(idx) {
		k = t.raw.GetInReplyToIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsInReplyToObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsInReplyToLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendInReplyTo appends the value for property 'inReplyTo'.
func (t *Person) AppendInReplyTo(k *url.URL) {
	t.raw.AppendInReplyToIRI(k)

}

// PrependInReplyTo prepends the value for property 'inReplyTo'.
func (t *Person) PrependInReplyTo(k *url.URL) {
	t.raw.PrependInReplyToIRI(k)

}

// RemoveInReplyTo deletes the value from the specified index for property 'inReplyTo'.
func (t *Person) RemoveInReplyTo(idx int) {
	t.raw.RemoveInReplyToIRI(idx)

}

// HasInReplyTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasInReplyTo(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsInReplyToIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsInReplyToLink(idx) {
		p = RawPresence
	} else if t.raw.IsInReplyToIRI(idx) {
		p = RawPresence
	}
	return

}

// LenLocation returns the number of values this property contains. Each index be used with HasLocation to determine if ResolveLocation is safe to call or if raw handling would be needed.
func (t *Person) LenLocation() (idx int) {
	return t.raw.LocationLen()

}

// ResolveLocation passes the actual concrete type to the resolver for handing property location. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveLocation(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsLocationObject(idx) {
		handled, err = r.dispatch(t.raw.GetLocationObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLocationLink(idx) {
		handled, err = r.dispatch(t.raw.GetLocationLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLocationIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasLocation returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasLocation(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsLocationObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsLocationLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsLocationIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendLocation appends an 'Object' typed value.
func (t *Person) AppendLocation(i vocab.ObjectType) {
	t.raw.AppendLocationObject(i)

}

// PrependLocation prepends an 'Object' typed value.
func (t *Person) PrependLocation(i vocab.ObjectType) {
	t.raw.PrependLocationObject(i)

}

// AppendLocationLink appends a 'Link' typed value.
func (t *Person) AppendLocationLink(i vocab.LinkType) {
	t.raw.AppendLocationLink(i)

}

// PrependLocationLink prepends a 'Link' typed value.
func (t *Person) PrependLocationLink(i vocab.LinkType) {
	t.raw.PrependLocationLink(i)

}

// LenPreview returns the number of values this property contains. Each index be used with HasPreview to determine if ResolvePreview is safe to call or if raw handling would be needed.
func (t *Person) LenPreview() (idx int) {
	return t.raw.PreviewLen()

}

// ResolvePreview passes the actual concrete type to the resolver for handing property preview. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolvePreview(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsPreviewObject(idx) {
		handled, err = r.dispatch(t.raw.GetPreviewObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsPreviewLink(idx) {
		handled, err = r.dispatch(t.raw.GetPreviewLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsPreviewIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasPreview returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasPreview(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsPreviewObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsPreviewLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsPreviewIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendPreview appends an 'Object' typed value.
func (t *Person) AppendPreview(i vocab.ObjectType) {
	t.raw.AppendPreviewObject(i)

}

// PrependPreview prepends an 'Object' typed value.
func (t *Person) PrependPreview(i vocab.ObjectType) {
	t.raw.PrependPreviewObject(i)

}

// AppendPreviewLink appends a 'Link' typed value.
func (t *Person) AppendPreviewLink(i vocab.LinkType) {
	t.raw.AppendPreviewLink(i)

}

// PrependPreviewLink prepends a 'Link' typed value.
func (t *Person) PrependPreviewLink(i vocab.LinkType) {
	t.raw.PrependPreviewLink(i)

}

// GetPublished attempts to get this 'published' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetPublished() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsPublished() {
		k = t.raw.GetPublished()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsPublishedIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasPublished returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasPublished() (p Presence) {
	p = NoPresence
	if t.raw.IsPublished() {
		p = ConvenientPresence
	} else if t.raw.IsPublishedIRI() {
		p = RawPresence
	}
	return

}

// SetPublished sets the value for property 'published'.
func (t *Person) SetPublished(k time.Time) {
	t.raw.SetPublished(k)

}

// ResolveReplies passes the actual concrete type to the resolver for handing property replies. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveReplies(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsReplies() {
		handled, err = r.dispatch(t.raw.GetReplies())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsRepliesIRI() {
		s = RawResolutionNeeded
	}
	return

}

// HasReplies returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasReplies() (p Presence) {
	p = NoPresence
	if t.raw.IsReplies() {
		p = ConvenientPresence
	} else if t.raw.IsRepliesIRI() {
		p = RawPresence
	}
	return

}

// SetReplies sets this value to be a 'Collection' type.
func (t *Person) SetReplies(i vocab.CollectionType) {
	t.raw.SetReplies(i)

}

// GetStartTime attempts to get this 'startTime' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetStartTime() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsStartTime() {
		k = t.raw.GetStartTime()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsStartTimeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasStartTime returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasStartTime() (p Presence) {
	p = NoPresence
	if t.raw.IsStartTime() {
		p = ConvenientPresence
	} else if t.raw.IsStartTimeIRI() {
		p = RawPresence
	}
	return

}

// SetStartTime sets the value for property 'startTime'.
func (t *Person) SetStartTime(k time.Time) {
	t.raw.SetStartTime(k)

}

// LenSummary returns the number of values this property contains. Each index be used with HasSummary to determine if GetSummary is safe to call or if raw handling would be needed.
func (t *Person) LenSummary() (idx int) {
	return t.raw.SummaryLen()

}

// GetSummary attempts to get this 'summary' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetSummary(idx int) (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsSummaryString(idx) {
		k = t.raw.GetSummaryString(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsSummaryLangString(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsSummaryIRI(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendSummary appends the value for property 'summary'.
func (t *Person) AppendSummary(k string) {
	t.raw.AppendSummaryString(k)

}

// PrependSummary prepends the value for property 'summary'.
func (t *Person) PrependSummary(k string) {
	t.raw.PrependSummaryString(k)

}

// RemoveSummary deletes the value from the specified index for property 'summary'.
func (t *Person) RemoveSummary(idx int) {
	t.raw.RemoveSummaryString(idx)

}

// HasSummary returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasSummary(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsSummaryString(idx) {
		p = ConvenientPresence
	} else if t.raw.IsSummaryLangString(idx) {
		p = RawPresence
	} else if t.raw.IsSummaryIRI(idx) {
		p = RawPresence
	}
	return

}

// SummaryLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Person) SummaryLanguages() (l []string) {
	return t.raw.SummaryMapLanguages()

}

// GetSummaryMap retrieves the value of 'summary' for the specified language, or an empty string if it does not exist
func (t *Person) GetSummaryForLanguage(l string) (v string) {
	return t.raw.GetSummaryMap(l)

}

// SetSummaryForLanguage sets the value of 'summary' for the specified language
func (t *Person) SetSummaryForLanguage(l string, v string) {
	t.raw.SetSummaryMap(l, v)

}

// LenTag returns the number of values this property contains. Each index be used with HasTag to determine if ResolveTag is safe to call or if raw handling would be needed.
func (t *Person) LenTag() (idx int) {
	return t.raw.TagLen()

}

// ResolveTag passes the actual concrete type to the resolver for handing property tag. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveTag(r *Resolver, idx int) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsTagObject(idx) {
		handled, err = r.dispatch(t.raw.GetTagObject(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsTagLink(idx) {
		handled, err = r.dispatch(t.raw.GetTagLink(idx))
		if handled {
			s = Resolved
		}
	} else if t.raw.IsTagIRI(idx) {
		s = RawResolutionNeeded
	}
	return

}

// HasTag returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasTag(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsTagObject(idx) {
		p = ConvenientPresence
	} else if t.raw.IsTagLink(idx) {
		p = ConvenientPresence
	} else if t.raw.IsTagIRI(idx) {
		p = RawPresence
	}
	return

}

// AppendTag appends an 'Object' typed value.
func (t *Person) AppendTag(i vocab.ObjectType) {
	t.raw.AppendTagObject(i)

}

// PrependTag prepends an 'Object' typed value.
func (t *Person) PrependTag(i vocab.ObjectType) {
	t.raw.PrependTagObject(i)

}

// AppendTagLink appends a 'Link' typed value.
func (t *Person) AppendTagLink(i vocab.LinkType) {
	t.raw.AppendTagLink(i)

}

// PrependTagLink prepends a 'Link' typed value.
func (t *Person) PrependTagLink(i vocab.LinkType) {
	t.raw.PrependTagLink(i)

}

// LenType returns the number of values this property contains. Each index be used with HasType to determine if GetType is safe to call or if raw handling would be needed.
func (t *Person) LenType() (idx int) {
	return t.raw.TypeLen()

}

// GetType attempts to get this 'type' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetType(idx int) (r Resolution, s string) {
	r = Unresolved
	if tmp := t.raw.GetType(idx); tmp != nil {
		ok := false
		if s, ok = tmp.(string); ok {
			r = Resolved
		} else {
			r = RawResolutionNeeded
		}
	}
	return

}

// AppendType appends the value for property 'type'.
func (t *Person) AppendType(i interface{}) {
	t.raw.AppendType(i)

}

// PrependType prepends the value for property 'type'.
func (t *Person) PrependType(i interface{}) {
	t.raw.PrependType(i)

}

// RemoveType deletes the value from the specified index for property 'type'.
func (t *Person) RemoveType(idx int) {
	t.raw.RemoveType(idx)

}

// GetUpdated attempts to get this 'updated' property as a time.Time. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetUpdated() (r Resolution, k time.Time) {
	r = Unresolved
	handled := false
	if t.raw.IsUpdated() {
		k = t.raw.GetUpdated()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsUpdatedIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasUpdated returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasUpdated() (p Presence) {
	p = NoPresence
	if t.raw.IsUpdated() {
		p = ConvenientPresence
	} else if t.raw.IsUpdatedIRI() {
		p = RawPresence
	}
	return

}

// SetUpdated sets the value for property 'updated'.
func (t *Person) SetUpdated(k time.Time) {
	t.raw.SetUpdated(k)

}

// LenUrl returns the number of values this property contains. Each index be used with HasUrl to determine if GetUrl is safe to call or if raw handling would be needed.
func (t *Person) LenUrl() (idx int) {
	return t.raw.UrlLen()

}

// GetUrl attempts to get this 'url' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetUrl(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsUrlAnyURI(idx) {
		k = t.raw.GetUrlAnyURI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsUrlLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendUrl appends the value for property 'url'.
func (t *Person) AppendUrl(k *url.URL) {
	t.raw.AppendUrlAnyURI(k)

}

// PrependUrl prepends the value for property 'url'.
func (t *Person) PrependUrl(k *url.URL) {
	t.raw.PrependUrlAnyURI(k)

}

// RemoveUrl deletes the value from the specified index for property 'url'.
func (t *Person) RemoveUrl(idx int) {
	t.raw.RemoveUrlAnyURI(idx)

}

// HasUrl returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasUrl(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsUrlAnyURI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsUrlLink(idx) {
		p = RawPresence
	}
	return

}

// LenTo returns the number of values this property contains. Each index be used with HasTo to determine if GetTo is safe to call or if raw handling would be needed.
func (t *Person) LenTo() (idx int) {
	return t.raw.ToLen()

}

// GetTo attempts to get this 'to' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetTo(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsToIRI(idx) {
		k = t.raw.GetToIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsToObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsToLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendTo appends the value for property 'to'.
func (t *Person) AppendTo(k *url.URL) {
	t.raw.AppendToIRI(k)

}

// PrependTo prepends the value for property 'to'.
func (t *Person) PrependTo(k *url.URL) {
	t.raw.PrependToIRI(k)

}

// RemoveTo deletes the value from the specified index for property 'to'.
func (t *Person) RemoveTo(idx int) {
	t.raw.RemoveToIRI(idx)

}

// HasTo returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasTo(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsToIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsToLink(idx) {
		p = RawPresence
	} else if t.raw.IsToIRI(idx) {
		p = RawPresence
	}
	return

}

// LenBto returns the number of values this property contains. Each index be used with HasBto to determine if GetBto is safe to call or if raw handling would be needed.
func (t *Person) LenBto() (idx int) {
	return t.raw.BtoLen()

}

// GetBto attempts to get this 'bto' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetBto(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsBtoIRI(idx) {
		k = t.raw.GetBtoIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsBtoObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsBtoLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendBto appends the value for property 'bto'.
func (t *Person) AppendBto(k *url.URL) {
	t.raw.AppendBtoIRI(k)

}

// PrependBto prepends the value for property 'bto'.
func (t *Person) PrependBto(k *url.URL) {
	t.raw.PrependBtoIRI(k)

}

// RemoveBto deletes the value from the specified index for property 'bto'.
func (t *Person) RemoveBto(idx int) {
	t.raw.RemoveBtoIRI(idx)

}

// HasBto returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasBto(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsBtoIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsBtoLink(idx) {
		p = RawPresence
	} else if t.raw.IsBtoIRI(idx) {
		p = RawPresence
	}
	return

}

// LenCc returns the number of values this property contains. Each index be used with HasCc to determine if GetCc is safe to call or if raw handling would be needed.
func (t *Person) LenCc() (idx int) {
	return t.raw.CcLen()

}

// GetCc attempts to get this 'cc' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetCc(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsCcIRI(idx) {
		k = t.raw.GetCcIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsCcObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsCcLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendCc appends the value for property 'cc'.
func (t *Person) AppendCc(k *url.URL) {
	t.raw.AppendCcIRI(k)

}

// PrependCc prepends the value for property 'cc'.
func (t *Person) PrependCc(k *url.URL) {
	t.raw.PrependCcIRI(k)

}

// RemoveCc deletes the value from the specified index for property 'cc'.
func (t *Person) RemoveCc(idx int) {
	t.raw.RemoveCcIRI(idx)

}

// HasCc returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasCc(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsCcIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsCcLink(idx) {
		p = RawPresence
	} else if t.raw.IsCcIRI(idx) {
		p = RawPresence
	}
	return

}

// LenBcc returns the number of values this property contains. Each index be used with HasBcc to determine if GetBcc is safe to call or if raw handling would be needed.
func (t *Person) LenBcc() (idx int) {
	return t.raw.BccLen()

}

// GetBcc attempts to get this 'bcc' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetBcc(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.IsBccIRI(idx) {
		k = t.raw.GetBccIRI(idx)
		if handled {
			r = Resolved
		}
	} else if t.raw.IsBccObject(idx) {
		r = RawResolutionNeeded
	} else if t.raw.IsBccLink(idx) {
		r = RawResolutionNeeded
	}
	return

}

// AppendBcc appends the value for property 'bcc'.
func (t *Person) AppendBcc(k *url.URL) {
	t.raw.AppendBccIRI(k)

}

// PrependBcc prepends the value for property 'bcc'.
func (t *Person) PrependBcc(k *url.URL) {
	t.raw.PrependBccIRI(k)

}

// RemoveBcc deletes the value from the specified index for property 'bcc'.
func (t *Person) RemoveBcc(idx int) {
	t.raw.RemoveBccIRI(idx)

}

// HasBcc returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasBcc(idx int) (p Presence) {
	p = NoPresence
	if t.raw.IsBccIRI(idx) {
		p = ConvenientPresence
	} else if t.raw.IsBccLink(idx) {
		p = RawPresence
	} else if t.raw.IsBccIRI(idx) {
		p = RawPresence
	}
	return

}

// GetMediaType attempts to get this 'mediaType' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetMediaType() (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsMediaType() {
		k = t.raw.GetMediaType()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsMediaTypeIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasMediaType returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasMediaType() (p Presence) {
	p = NoPresence
	if t.raw.IsMediaType() {
		p = ConvenientPresence
	} else if t.raw.IsMediaTypeIRI() {
		p = RawPresence
	}
	return

}

// SetMediaType sets the value for property 'mediaType'.
func (t *Person) SetMediaType(k string) {
	t.raw.SetMediaType(k)

}

// GetDuration attempts to get this 'duration' property as a time.Duration. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetDuration() (r Resolution, k time.Duration) {
	r = Unresolved
	handled := false
	if t.raw.IsDuration() {
		k = t.raw.GetDuration()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsDurationIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasDuration returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasDuration() (p Presence) {
	p = NoPresence
	if t.raw.IsDuration() {
		p = ConvenientPresence
	} else if t.raw.IsDurationIRI() {
		p = RawPresence
	}
	return

}

// SetDuration sets the value for property 'duration'.
func (t *Person) SetDuration(k time.Duration) {
	t.raw.SetDuration(k)

}

// ResolveSource passes the actual concrete type to the resolver for handing property source. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveSource(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsSource() {
		handled, err = r.dispatch(t.raw.GetSource())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsSourceIRI() {
		s = RawResolutionNeeded
	}
	return

}

// HasSource returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasSource() (p Presence) {
	p = NoPresence
	if t.raw.IsSource() {
		p = ConvenientPresence
	} else if t.raw.IsSourceIRI() {
		p = RawPresence
	}
	return

}

// SetSource sets this value to be a 'Object' type.
func (t *Person) SetSource(i vocab.ObjectType) {
	t.raw.SetSource(i)

}

// ResolveInbox passes the actual concrete type to the resolver for handing property inbox. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveInbox(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsInboxOrderedCollection() {
		handled, err = r.dispatch(t.raw.GetInboxOrderedCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsInboxAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasInbox returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasInbox() (p Presence) {
	p = NoPresence
	if t.raw.IsInboxOrderedCollection() {
		p = ConvenientPresence
	} else if t.raw.IsInboxAnyURI() {
		p = RawPresence
	}
	return

}

// SetInbox sets this value to be a 'OrderedCollection' type.
func (t *Person) SetInbox(i vocab.OrderedCollectionType) {
	t.raw.SetInboxOrderedCollection(i)

}

// ResolveOutbox passes the actual concrete type to the resolver for handing property outbox. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveOutbox(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsOutboxOrderedCollection() {
		handled, err = r.dispatch(t.raw.GetOutboxOrderedCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsOutboxAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasOutbox returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasOutbox() (p Presence) {
	p = NoPresence
	if t.raw.IsOutboxOrderedCollection() {
		p = ConvenientPresence
	} else if t.raw.IsOutboxAnyURI() {
		p = RawPresence
	}
	return

}

// SetOutbox sets this value to be a 'OrderedCollection' type.
func (t *Person) SetOutbox(i vocab.OrderedCollectionType) {
	t.raw.SetOutboxOrderedCollection(i)

}

// ResolveFollowing passes the actual concrete type to the resolver for handing property following. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveFollowing(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsFollowingCollection() {
		handled, err = r.dispatch(t.raw.GetFollowingCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsFollowingOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsFollowingAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasFollowing returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasFollowing() (p Presence) {
	p = NoPresence
	if t.raw.IsFollowingCollection() {
		p = ConvenientPresence
	} else if t.raw.IsFollowingOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsFollowingAnyURI() {
		p = RawPresence
	}
	return

}

// SetFollowing sets this value to be a 'Collection' type.
func (t *Person) SetFollowing(i vocab.CollectionType) {
	t.raw.SetFollowingCollection(i)

}

// ResolveFollowers passes the actual concrete type to the resolver for handing property followers. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveFollowers(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsFollowersCollection() {
		handled, err = r.dispatch(t.raw.GetFollowersCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsFollowersOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsFollowersAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasFollowers returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasFollowers() (p Presence) {
	p = NoPresence
	if t.raw.IsFollowersCollection() {
		p = ConvenientPresence
	} else if t.raw.IsFollowersOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsFollowersAnyURI() {
		p = RawPresence
	}
	return

}

// SetFollowers sets this value to be a 'Collection' type.
func (t *Person) SetFollowers(i vocab.CollectionType) {
	t.raw.SetFollowersCollection(i)

}

// ResolveLiked passes the actual concrete type to the resolver for handing property liked. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveLiked(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsLikedCollection() {
		handled, err = r.dispatch(t.raw.GetLikedCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLikedOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsLikedAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasLiked returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasLiked() (p Presence) {
	p = NoPresence
	if t.raw.IsLikedCollection() {
		p = ConvenientPresence
	} else if t.raw.IsLikedOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsLikedAnyURI() {
		p = RawPresence
	}
	return

}

// SetLiked sets this value to be a 'Collection' type.
func (t *Person) SetLiked(i vocab.CollectionType) {
	t.raw.SetLikedCollection(i)

}

// ResolveLikes passes the actual concrete type to the resolver for handing property likes. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveLikes(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsLikesCollection() {
		handled, err = r.dispatch(t.raw.GetLikesCollection())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsLikesOrderedCollection() {
		s = RawResolutionNeeded
	} else if t.raw.IsLikesAnyURI() {
		s = RawResolutionNeeded
	}
	return

}

// HasLikes returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasLikes() (p Presence) {
	p = NoPresence
	if t.raw.IsLikesCollection() {
		p = ConvenientPresence
	} else if t.raw.IsLikesOrderedCollection() {
		p = RawPresence
	} else if t.raw.IsLikesAnyURI() {
		p = RawPresence
	}
	return

}

// SetLikes sets this value to be a 'Collection' type.
func (t *Person) SetLikes(i vocab.CollectionType) {
	t.raw.SetLikesCollection(i)

}

// LenStreams returns the number of values this property contains. Each index be used with HasStreams to determine if GetStreams is safe to call or if raw handling would be needed.
func (t *Person) LenStreams() (idx int) {
	return t.raw.StreamsLen()

}

// GetStreams attempts to get this 'streams' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetStreams(idx int) (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if /*t.raw.HasStreams(idx)*/ true {
		k = t.raw.GetStreams(idx)
		if handled {
			r = Resolved
		}
	}
	return

}

// AppendStreams appends the value for property 'streams'.
func (t *Person) AppendStreams(k *url.URL) {
	t.raw.AppendStreams(k)

}

// PrependStreams prepends the value for property 'streams'.
func (t *Person) PrependStreams(k *url.URL) {
	t.raw.PrependStreams(k)

}

// RemoveStreams deletes the value from the specified index for property 'streams'.
func (t *Person) RemoveStreams(idx int) {
	t.raw.RemoveStreams(idx)

}

// GetPreferredUsername attempts to get this 'preferredUsername' property as a string. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetPreferredUsername() (r Resolution, k string) {
	r = Unresolved
	handled := false
	if t.raw.IsPreferredUsername() {
		k = t.raw.GetPreferredUsername()
		if handled {
			r = Resolved
		}
	} else if t.raw.IsPreferredUsernameIRI() {
		r = RawResolutionNeeded
	}
	return

}

// HasPreferredUsername returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasPreferredUsername() (p Presence) {
	p = NoPresence
	if t.raw.IsPreferredUsername() {
		p = ConvenientPresence
	} else if t.raw.IsPreferredUsernameIRI() {
		p = RawPresence
	}
	return

}

// SetPreferredUsername sets the value for property 'preferredUsername'.
func (t *Person) SetPreferredUsername(k string) {
	t.raw.SetPreferredUsername(k)

}

// PreferredUsernameLanguages returns all languages for this property's language mapping, or nil if there are none.
func (t *Person) PreferredUsernameLanguages() (l []string) {
	return t.raw.PreferredUsernameMapLanguages()

}

// GetPreferredUsernameMap retrieves the value of 'preferredUsername' for the specified language, or an empty string if it does not exist
func (t *Person) GetPreferredUsernameForLanguage(l string) (v string) {
	return t.raw.GetPreferredUsernameMap(l)

}

// SetPreferredUsernameForLanguage sets the value of 'preferredUsername' for the specified language
func (t *Person) SetPreferredUsernameForLanguage(l string, v string) {
	t.raw.SetPreferredUsernameMap(l, v)

}

// ResolveEndpoints passes the actual concrete type to the resolver for handing property endpoints. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) ResolveEndpoints(r *Resolver) (s Resolution, err error) {
	s = Unresolved
	handled := false
	if t.raw.IsEndpoints() {
		handled, err = r.dispatch(t.raw.GetEndpoints())
		if handled {
			s = Resolved
		}
	} else if t.raw.IsEndpointsIRI() {
		s = RawResolutionNeeded
	}
	return

}

// HasEndpoints returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasEndpoints() (p Presence) {
	p = NoPresence
	if t.raw.IsEndpoints() {
		p = ConvenientPresence
	} else if t.raw.IsEndpointsIRI() {
		p = RawPresence
	}
	return

}

// SetEndpoints sets this value to be a 'Object' type.
func (t *Person) SetEndpoints(i vocab.ObjectType) {
	t.raw.SetEndpoints(i)

}

// GetProxyUrl attempts to get this 'proxyUrl' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetProxyUrl() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasProxyUrl() {
		k = t.raw.GetProxyUrl()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasProxyUrl returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasProxyUrl() (p Presence) {
	p = NoPresence
	if t.raw.HasProxyUrl() {
		p = ConvenientPresence
	}
	return

}

// SetProxyUrl sets the value for property 'proxyUrl'.
func (t *Person) SetProxyUrl(k *url.URL) {
	t.raw.SetProxyUrl(k)

}

// GetOauthAuthorizationEndpoint attempts to get this 'oauthAuthorizationEndpoint' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetOauthAuthorizationEndpoint() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasOauthAuthorizationEndpoint() {
		k = t.raw.GetOauthAuthorizationEndpoint()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasOauthAuthorizationEndpoint returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasOauthAuthorizationEndpoint() (p Presence) {
	p = NoPresence
	if t.raw.HasOauthAuthorizationEndpoint() {
		p = ConvenientPresence
	}
	return

}

// SetOauthAuthorizationEndpoint sets the value for property 'oauthAuthorizationEndpoint'.
func (t *Person) SetOauthAuthorizationEndpoint(k *url.URL) {
	t.raw.SetOauthAuthorizationEndpoint(k)

}

// GetOauthTokenEndpoint attempts to get this 'oauthTokenEndpoint' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetOauthTokenEndpoint() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasOauthTokenEndpoint() {
		k = t.raw.GetOauthTokenEndpoint()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasOauthTokenEndpoint returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasOauthTokenEndpoint() (p Presence) {
	p = NoPresence
	if t.raw.HasOauthTokenEndpoint() {
		p = ConvenientPresence
	}
	return

}

// SetOauthTokenEndpoint sets the value for property 'oauthTokenEndpoint'.
func (t *Person) SetOauthTokenEndpoint(k *url.URL) {
	t.raw.SetOauthTokenEndpoint(k)

}

// GetProvideClientKey attempts to get this 'provideClientKey' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetProvideClientKey() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasProvideClientKey() {
		k = t.raw.GetProvideClientKey()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasProvideClientKey returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasProvideClientKey() (p Presence) {
	p = NoPresence
	if t.raw.HasProvideClientKey() {
		p = ConvenientPresence
	}
	return

}

// SetProvideClientKey sets the value for property 'provideClientKey'.
func (t *Person) SetProvideClientKey(k *url.URL) {
	t.raw.SetProvideClientKey(k)

}

// GetSignClientKey attempts to get this 'signClientKey' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetSignClientKey() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasSignClientKey() {
		k = t.raw.GetSignClientKey()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasSignClientKey returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasSignClientKey() (p Presence) {
	p = NoPresence
	if t.raw.HasSignClientKey() {
		p = ConvenientPresence
	}
	return

}

// SetSignClientKey sets the value for property 'signClientKey'.
func (t *Person) SetSignClientKey(k *url.URL) {
	t.raw.SetSignClientKey(k)

}

// GetSharedInbox attempts to get this 'sharedInbox' property as a *url.URL. It returns a Resolution appropriate for clients to determine whether it would be necessary to do raw handling.
func (t *Person) GetSharedInbox() (r Resolution, k *url.URL) {
	r = Unresolved
	handled := false
	if t.raw.HasSharedInbox() {
		k = t.raw.GetSharedInbox()
		if handled {
			r = Resolved
		}
	}
	return

}

// HasSharedInbox returns a Presence appropriate for clients to determine whether it would be necessary to do raw handling, if desired.
func (t *Person) HasSharedInbox() (p Presence) {
	p = NoPresence
	if t.raw.HasSharedInbox() {
		p = ConvenientPresence
	}
	return

}

// SetSharedInbox sets the value for property 'sharedInbox'.
func (t *Person) SetSharedInbox(k *url.URL) {
	t.raw.SetSharedInbox(k)

}
