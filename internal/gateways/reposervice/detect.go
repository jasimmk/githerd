package reposervice

import "strings"

// DetectRemoteType detects the type of the remote.

func DetectRemoteType(remote string) string {
	if isBitbucket(remote) {
		return "bitbucket"
	}

	if isGitLab(remote) {
		return "gitlab"
	}

	if isGitHub(remote) {
		return "github"
	}

	return ""
}

func isBitbucket(remote string) bool {
	return strings.Contains(remote, "bitbucket.") || strings.Contains(remote, "bitbucketdc.")
}
func isGitLab(remote string) bool {
	return strings.Contains(remote, "gitlab.")
}
func isGitHub(remote string) bool {
	return strings.Contains(remote, "github.")

}
