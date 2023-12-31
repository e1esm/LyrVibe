package service

type ServicesBuilder struct {
	Services Services
}

func (sb *ServicesBuilder) WithAuthProvider(provider AuthenticationProvider) *ServicesBuilder {
	sb.Services.AuthService = provider
	return sb
}

func (sb *ServicesBuilder) WithArtistsProvider(provider ArtistServiceProvider) *ServicesBuilder {
	sb.Services.ArtistService = provider
	return sb
}

func (sb *ServicesBuilder) Build() Services {
	return sb.Services
}

func NewServiceBuilder() *ServicesBuilder {
	return &ServicesBuilder{Services: Services{}}
}
