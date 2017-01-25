package environment

import (
	"github.com/Sirupsen/logrus"
	"github.com/normegil/aphrodite/modules/datasource"
)

type Environment struct {
	Log        logrus.FieldLogger
	DataSource datasource.DataSource
}
