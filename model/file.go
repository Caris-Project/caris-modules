package model

// FileConfig ...
type FileConfig struct {
	Host string
}

var c FileConfig

// SetupFileConfig ...
func SetupFileConfig(host string) {
	c.Host = host
}

// GetAvatarURL ...
func GetAvatarURL(avatar string) string {
	if avatar == "" {
		return c.Host + "default-avatar.png"
	}
	return c.Host + avatar
}

// GetNotificationPhotoURL ...
func GetNotificationPhotoURL(name string) string {
	if name == "" {
		return ""
	}
	return c.Host + "sm_" + name
}
