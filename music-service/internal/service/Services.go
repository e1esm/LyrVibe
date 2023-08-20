package service

type Services struct {
	MusicService MusicServiceProvider
}

type ServicesBuilder struct {
	Services Services
}

func NewServicesBuilder() *ServicesBuilder {
	return &ServicesBuilder{Services: Services{}}
}

func (sb *ServicesBuilder) WithMusicService(musicService MusicServiceProvider) *ServicesBuilder {
	sb.Services.MusicService = musicService
	return sb
}

func (sb *ServicesBuilder) Build() Services {
	return sb.Services
}
