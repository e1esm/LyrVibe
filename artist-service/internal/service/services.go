package service

type Services struct {
	ArtistService Service
}

type ServicesBuilder struct {
	Services Services
}

func NewServiceBuilder() *ServicesBuilder {
	return &ServicesBuilder{Services: Services{}}
}

func (s *ServicesBuilder) WithArtistService(service Service) *ServicesBuilder {
	s.Services.ArtistService = service
	return s
}

func (s *ServicesBuilder) Build() Services {
	return s.Services
}
