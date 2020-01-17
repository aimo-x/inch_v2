package log

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestStart(T *testing.T) {
	l := New()
	p := []int{0, 1, 2, 3, 4, 5, 6}
	for i := range p {
		l.WithFields(logrus.Fields{
			"test": "test",
			"i":    i,
		}).Info("info test")
	}
}
