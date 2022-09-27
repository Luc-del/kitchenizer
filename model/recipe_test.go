package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestRecipe_loadPicture(t *testing.T) {
	r := Recipe{}
	require.NoError(t, r.loadPicture("../.ignore/image.jpg"))

	assert.NoError(t, ioutil.WriteFile("../.ignore/output.jpg", r.Picture, 0777))
}

func TestRecipe_loadWebPicture(t *testing.T) {
	r := Recipe{}
	require.NoError(t, r.loadWebPicture("https://www.ionos.fr/digitalguide/fileadmin/DigitalGuide/Teaser/datensicherung-von-verschiedene-devices-c.jpg"))

	assert.NoError(t, ioutil.WriteFile("../.ignore/output_web.jpg", r.Picture, 0777))
}
