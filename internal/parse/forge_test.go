package parse

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"mimatache/github.com/forge/internal/forge"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

type ReadCloser struct {
	*strings.Reader
}

func (rc ReadCloser) Close() error {
	return nil
}

var (
	goodForge = &manifest.Forge{
		Forgeries: manifest.Forgeries{
			{
				Name:        "testAction",
				Description: "Test action",
				Pre:         []manifest.ForgeryName{},
				Cmd:         "echo 0",
			},
		},
	}
	nameWithSpaces = &manifest.Forge{
		Forgeries: manifest.Forgeries{
			{
				Name:        "test action",
				Description: "test description",
				Pre:         []manifest.ForgeryName{},
				Cmd:         "echo 0",
			},
		},
	}
)

type ForgeMock struct {
	forges map[string]*manifest.Forge
}

func (m ForgeMock) MockForgeReader(value string) (io.ReadCloser, error) {
	frg, err := yaml.Marshal(m.forges[value])
	fmt.Println(string(frg))
	if err != nil {
		return nil, err
	}
	return &ReadCloser{
		Reader: strings.NewReader(string(frg)),
	}, nil
}

func TestParse(t *testing.T) {
	g := NewWithT(t)

	m := ForgeMock{
		forges: map[string]*manifest.Forge{
			"Forge": goodForge,
		},
	}

	frg, err := GetForge("Forge", m.MockForgeReader)
	g.Expect(err).ShouldNot(HaveOccurred())
	goodForge.Include = []string{"Forge"}
	g.Expect(frg).To(Equal(goodForge))

}
