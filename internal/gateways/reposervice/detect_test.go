package reposervice

import "testing"

func TestDetectRemoteType(t *testing.T) {
	tests := []struct {
		name   string
		remote string
		want   string
	}{
		{
			name:   "Bitbucket",
			remote: "https://bitbucket.org/user/repo.git",
			want:   "bitbucket",
		},
		{
			name:   "Bitbucket DC",
			remote: "https://bitbucketdc.example.com/user/repo.git",
			want:   "bitbucket",
		},
		{
			name:   "GitLab",
			remote: "https://gitlab.com/user/repo.git",
			want:   "gitlab",
		},
		{
			name:   "GitHub",
			remote: "https://github.com/user/repo.git",
			want:   "github",
		},
		{
			name:   "Unknown",
			remote: "https://example.com/user/repo.git",
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectRemoteType(tt.remote)
			if got != tt.want {
				t.Errorf("DetectRemoteType(%s) = %s, want %s", tt.remote, got, tt.want)
			}
		})
	}
}

func TestIsBitbucket(t *testing.T) {
	tests := []struct {
		name   string
		remote string
		want   bool
	}{
		{
			name:   "Bitbucket",
			remote: "https://bitbucket.org/user/repo.git",
			want:   true,
		},
		{
			name:   "Bitbucket DC",
			remote: "https://bitbucketdc.example.com/user/repo.git",
			want:   true,
		},
		{
			name:   "GitLab",
			remote: "https://gitlab.com/user/repo.git",
			want:   false,
		},
		{
			name:   "GitHub",
			remote: "https://github.com/user/repo.git",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBitbucket(tt.remote)
			if got != tt.want {
				t.Errorf("isBitbucket(%s) = %v, want %v", tt.remote, got, tt.want)
			}
		})
	}
}

func TestIsGitLab(t *testing.T) {
	tests := []struct {
		name   string
		remote string
		want   bool
	}{
		{
			name:   "Bitbucket",
			remote: "https://bitbucket.org/user/repo.git",
			want:   false,
		},
		{
			name:   "Bitbucket DC",
			remote: "https://bitbucketdc.example.com/user/repo.git",
			want:   false,
		},
		{
			name:   "GitLab",
			remote: "https://gitlab.com/user/repo.git",
			want:   true,
		},
		{
			name:   "GitHub",
			remote: "https://github.com/user/repo.git",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isGitLab(tt.remote)
			if got != tt.want {
				t.Errorf("isGitLab(%s) = %v, want %v", tt.remote, got, tt.want)
			}
		})
	}
}
