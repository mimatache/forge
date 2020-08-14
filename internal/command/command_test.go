package command

import (
	"mimatache/github.com/forge/internal/manifest"
	"testing"

	. "github.com/onsi/gomega"
)

func TestBuildCommandList_SimpleList(t *testing.T) {
	g := NewWithT(t)

	forgeries := map[manifest.ForgeryName]manifest.Forgery{
		"cmd1": {
			Pre: []manifest.ForgeryName{"cmd2", "cmd3", "cmd4"},
			Cmd: "cmd1",
		},
		"cmd2": {
			Cmd: "cmd2",
		},
		"cmd3": {
			Cmd: "cmd3",
		},
		"cmd4": {
			Cmd: "cmd4",
		},
	}

	executables, err := NewCommandList("cmd1", forgeries)
	g.Expect(err).ToNot(HaveOccurred(), "error when parsing simple list")
	expectedOrder := []manifest.ForgeryName{"cmd2", "cmd3", "cmd4", "cmd1"}
	for i, v := range executables {
		g.Expect(v).To(Equal(forgeries[expectedOrder[i]]))
	}

}

func TestBuildCommandList_NoInfiniteLoop(t *testing.T) {

	g := NewWithT(t)
	forgeries := map[manifest.ForgeryName]manifest.Forgery{
		"cmd1": {
			Pre: []manifest.ForgeryName{"cmd2"},
			Cmd: "cmd1",
		},
		"cmd2": {
			Pre: []manifest.ForgeryName{"cmd1"},
			Cmd: "cmd2",
		},
	}

	_, err := NewCommandList("cmd1", forgeries)
	g.Expect(err).To(HaveOccurred(), "error when parsing simple list")

}
