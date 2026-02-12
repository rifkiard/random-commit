package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Username    string
	AccessToken string
	Repos       []string
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	// Load config
	config := loadConfig()
	if len(config.Repos) == 0 {
		fmt.Println("Error: No repositories configured in .env")
		os.Exit(1)
	}

	fmt.Printf("📦 Loaded %d repositories\n", len(config.Repos))

	// Create repos directory if not exists
	reposDir := "repos"
	if err := os.MkdirAll(reposDir, 0755); err != nil {
		fmt.Println("Error creating repos directory:", err)
		os.Exit(1)
	}

	// Clone repositories that don't exist yet
	fmt.Println("\n🔄 Checking repositories...")
	for _, repo := range config.Repos {
		repoPath := filepath.Join(reposDir, getRepoName(repo))
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			fmt.Printf("📥 Cloning %s...\n", repo)
			if err := cloneRepo(config, repo, repoPath); err != nil {
				fmt.Printf("Error cloning %s: %v\n", repo, err)
				continue
			}
			fmt.Printf("✓ Cloned %s\n", repo)
		} else {
			fmt.Printf("✓ %s already exists\n", getRepoName(repo))
		}
	}

	// Pick random repo
	rand.Seed(time.Now().UnixNano())
	selectedRepo := config.Repos[rand.Intn(len(config.Repos))]
	repoPath := filepath.Join(reposDir, getRepoName(selectedRepo))

	fmt.Printf("\n🎲 Selected random repo: %s\n", selectedRepo)

	// Make random change to README.md
	readmePath := filepath.Join(repoPath, "README.md")
	if err := makeRandomChange(readmePath); err != nil {
		fmt.Println("Error making random change:", err)
		os.Exit(1)
	}

	// Git add, commit, push
	fmt.Println("\n📝 Committing changes...")
	if err := gitAddCommitPush(repoPath, config); err != nil {
		fmt.Println("Error during git operations:", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Successfully committed and pushed changes!")
}

// loadConfig loads configuration from environment variables
func loadConfig() Config {
	reposStr := os.Getenv("REPOS")
	repos := strings.Split(reposStr, ",")

	// Trim whitespace from each repo
	for i, repo := range repos {
		repos[i] = strings.TrimSpace(repo)
	}

	return Config{
		Username:    os.Getenv("GIT_USERNAME"),
		AccessToken: os.Getenv("GIT_ACCESS_TOKEN"),
		Repos:       repos,
	}
}

// getRepoName extracts repo name from full repo path (e.g., "user/repo" -> "repo")
func getRepoName(fullRepo string) string {
	parts := strings.Split(fullRepo, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return fullRepo
}

// cloneRepo clones a repository using access token
func cloneRepo(config Config, repo, targetPath string) error {
	// Format: https://username:token@github.com/user/repo.git
	repoURL := fmt.Sprintf("https://%s:%s@github.com/%s.git",
		config.Username, config.AccessToken, repo)

	cmd := exec.Command("git", "clone", repoURL, targetPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// makeRandomChange makes a random change to README.md
func makeRandomChange(readmePath string) error {
	// Check if README.md exists, if not create it
	var content string
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		content = "# Auto-commit Repository\n\nThis repository is managed by auto-commit tool.\n\n"
	} else {
		data, err := os.ReadFile(readmePath)
		if err != nil {
			return err
		}
		content = string(data)
	}

	// Generate random sentence
	randomSentence := gofakeit.Sentence(rand.Intn(10) + 5) // 5-14 words
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	newLine := fmt.Sprintf("\n## Update %s\n\n%s\n", timestamp, randomSentence)
	content += newLine

	return os.WriteFile(readmePath, []byte(content), 0644)
}

// gitAddCommitPush performs git add, commit, and push
func gitAddCommitPush(repoPath string, config Config) error {
	// Change to repo directory
	originalDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(repoPath); err != nil {
		return err
	}

	// Git add
	fmt.Println("  📌 Staging changes...")
	cmd := exec.Command("git", "add", "README.md")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Git commit with random message
	commitMsg := gofakeit.HackerPhrase() // Generate random tech-related phrase
	fmt.Printf("  💾 Creating commit: %s\n", commitMsg)

	cmd = exec.Command("git", "commit", "-m", commitMsg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Git push
	fmt.Println("  🚀 Pushing to remote...")
	cmd = exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
