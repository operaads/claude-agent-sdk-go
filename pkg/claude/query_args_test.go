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

func TestBuildArgs_AddsSettingSourcesFlag(t *testing.T) {
	q := &queryImpl{
		opts: &Options{
			SettingSources: []ConfigScope{
				ConfigScopeUser,
				ConfigScopeProject,
			},
		},
	}

	args := q.buildArgs()

	if !hasFlagValue(args, "--setting-sources", "user,project") {
		t.Fatalf("expected args to contain --setting-sources user,project, got %v", args)
	}
}

func TestBuildArgs_AddsEmptySettingSourcesFlagWhenUnset(t *testing.T) {
	q := &queryImpl{
		opts: &Options{},
	}

	args := q.buildArgs()

	if !hasFlagValue(args, "--setting-sources", "") {
		t.Fatalf("expected args to contain --setting-sources with empty value, got %v", args)
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
