package service

type Services struct {
	ArtistService Service
	RoleService   RolesProvider
	MusicService  MusicServiceProvider
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

func (s *ServicesBuilder) WithMusicService(service MusicServiceProvider) *ServicesBuilder {
	s.Services.MusicService = service
	return s
}

func (s *ServicesBuilder) WithRoleService(service RolesProvider) *ServicesBuilder {
	s.Services.RoleService = service
	return s
}

func (s *ServicesBuilder) Build() Services {
	return s.Services
}
