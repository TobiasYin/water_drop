package event

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

type Event struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		Id               int         `json:"id"`
		NodeId           string      `json:"node_id"`
		Name             string      `json:"name"`
		FullName         string      `json:"full_name"`
		Private          bool        `json:"private"`
		HtmlUrl          string      `json:"html_url"`
		Description      interface{} `json:"description"`
		Fork             bool        `json:"fork"`
		Url              string      `json:"url"`
		ForksUrl         string      `json:"forks_url"`
		KeysUrl          string      `json:"keys_url"`
		CollaboratorsUrl string      `json:"collaborators_url"`
		TeamsUrl         string      `json:"teams_url"`
		HooksUrl         string      `json:"hooks_url"`
		IssueEventsUrl   string      `json:"issue_events_url"`
		EventsUrl        string      `json:"events_url"`
		AssigneesUrl     string      `json:"assignees_url"`
		BranchesUrl      string      `json:"branches_url"`
		TagsUrl          string      `json:"tags_url"`
		BlobsUrl         string      `json:"blobs_url"`
		GitTagsUrl       string      `json:"git_tags_url"`
		GitRefsUrl       string      `json:"git_refs_url"`
		TreesUrl         string      `json:"trees_url"`
		StatusesUrl      string      `json:"statuses_url"`
		LanguagesUrl     string      `json:"languages_url"`
		StargazersUrl    string      `json:"stargazers_url"`
		ContributorsUrl  string      `json:"contributors_url"`
		SubscribersUrl   string      `json:"subscribers_url"`
		SubscriptionUrl  string      `json:"subscription_url"`
		CommitsUrl       string      `json:"commits_url"`
		GitCommitsUrl    string      `json:"git_commits_url"`
		CommentsUrl      string      `json:"comments_url"`
		IssueCommentUrl  string      `json:"issue_comment_url"`
		ContentsUrl      string      `json:"contents_url"`
		CompareUrl       string      `json:"compare_url"`
		MergesUrl        string      `json:"merges_url"`
		ArchiveUrl       string      `json:"archive_url"`
		DownloadsUrl     string      `json:"downloads_url"`
		IssuesUrl        string      `json:"issues_url"`
		PullsUrl         string      `json:"pulls_url"`
		MilestonesUrl    string      `json:"milestones_url"`
		NotificationsUrl string      `json:"notifications_url"`
		LabelsUrl        string      `json:"labels_url"`
		ReleasesUrl      string      `json:"releases_url"`
		DeploymentsUrl   string      `json:"deployments_url"`
		CreatedAt        int         `json:"created_at"`
		UpdatedAt        string      `json:"updated_at"`
		PushedAt         int         `json:"pushed_at"`
		GitUrl           string      `json:"git_url"`
		SshUrl           string      `json:"ssh_url"`
		CloneUrl         string      `json:"clone_url"`
		SvnUrl           string      `json:"svn_url"`
		Homepage         interface{} `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         interface{} `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		ForksCount       int         `json:"forks_count"`
		MirrorUrl        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		Disabled         bool        `json:"disabled"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		License          interface{} `json:"license"`
		Forks            int         `json:"forks"`
		OpenIssues       int         `json:"open_issues"`
		Watchers         int         `json:"watchers"`
		DefaultBranch    string      `json:"default_branch"`
		Stargazers       int         `json:"stargazers"`
		MasterBranch     string      `json:"master_branch"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Created bool        `json:"created"`
	Deleted bool        `json:"deleted"`
	Forced  bool        `json:"forced"`
	BaseRef interface{} `json:"base_ref"`
	Compare string      `json:"compare"`
}

var (
	lastUpdate time.Time
	run        *exec.Cmd
	lock       sync.Mutex
)

func init() {
	lastUpdate = time.Now()
}

func IsTarget(e Event) bool {
	if e.Ref != "refs/heads/master" {
		return false
	}
	pushAt := time.Unix(int64(e.Repository.PushedAt), 0)

	if lastUpdate.Sub(pushAt) > 0 {
		return false
	}

	return true
}

func StopRun() {
	log.Println("Try Stop old.")
	if run == nil {
		return
	}
	go func() {
		_ = run.Process.Kill()
	}()
	_ = run.Wait()
	run = nil
}

func Clean() {
	log.Println("Try Clean old files.")
	cmd := exec.Command("/bin/bash", "-c", "cd ../;./clean.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		_ = cmd.Wait()
	}
}

func Build() {
	log.Println("Try Build New Target.")
	cmd := exec.Command("/bin/bash", "-c", "cd ../;./build.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		_ = cmd.Wait()
	}
}

func UpdateRepo() {
	log.Println("Try Update Repos.")
	cmd := exec.Command("/bin/bash", "-c", "cd ../;git pull origin master;")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		_ = cmd.Wait()
	}
}

func Run() {
	log.Println("Try Run new.")
	cmd := exec.Command("/bin/bash", "-c", "cd ../;./run.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	run = cmd
}

func Workflow() {
	lock.Lock()
	log.Println("Start to Run workflow!")
	StopRun()
	Clean()
	UpdateRepo()
	Build()
	Run()
	lastUpdate = time.Now()
	lock.Unlock()
}
