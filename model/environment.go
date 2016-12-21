package model

import "github.com/sirupsen/logrus"

type Env struct {
	Log  logrus.FieldLogger
	DataSource DataSource
}