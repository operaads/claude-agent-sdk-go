package claude

import "testing"

func TestBuildArgs_AddsToolFlags(t *testing.T) {
	q := &queryImpl{
		opts: &Options{
			Tools: []string{"Read", "Write", "Bash"},
		},
	}

	args := q.buildArgs()

	wantPairs := []string{
		"--tools", "Read",
		"--tools", "Write",
		"--tools", "Bash",
	}

	for i := 0; i < len(wantPairs); i += 2 {
		flag := wantPairs[i]
		value := wantPairs[i+1]
		if !hasFlagValue(args, flag, value) {
			t.Fatalf("expected args to contain %s %s, got %v", flag, value, args)
		}
	}
}

func hasFlagValue(args []string, flag, value string) bool {
	for i := 0; i < len(args)-1; i++ {
		if args[i] == flag && args[i+1] == value {
			return true
		}
	}
	return false
}
