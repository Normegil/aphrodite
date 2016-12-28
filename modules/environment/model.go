package environment

import (
	"github.com/normegil/aphrodite/modules/datasource"
	"github.com/sirupsen/logrus"
)

type Environment struct {
	Log        logrus.FieldLogger
	DataSource datasource.DataSource
}
