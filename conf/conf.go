package conf

import "code.google.com/p/goconf/conf"

type ConfigFile struct {
	conf.ConfigFile
}

// GetString gets the string value for the given option in the section.
func (c *ConfigFile) GetStringDef(section string, option string, def string) string {
	value, err := c.GetString(section, option)
	if err != nil {
		return def
	}

	return value
}

// GetInt has the same behaviour as GetString but converts the response to int.
func (c *ConfigFile) GetIntDef(section string, option string, def int) int {
	value, err := c.GetInt(section, option)
	if err != nil {
		return def
	}

	return value
}

// GetFloat has the same behaviour as GetString but converts the response to float.
func (c *ConfigFile) GetFloat64Def(section string, option string, def float64) float64 {
	value, err := c.GetFloat64(section, option)
	if err != nil {
		return def
	}

	return value
}

// GetBool has the same behaviour as GetString but converts the response to bool.
// See constant BoolStrings for string values converted to bool.
func (c *ConfigFile) GetBoolDef(section string, option string, def bool) bool {
	value, err := c.GetBool(section, option)
	if err != nil {
		return def
	}

	return value
}

func ReadConfigFile(fname string) (c *ConfigFile, err error) {
	corg, err := conf.ReadConfigFile(fname)

	if err != nil {
		return nil, err
	}
	return &ConfigFile{*corg}, nil
}
