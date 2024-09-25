package repo_test

import (
	"testing"

	"github.com/carp-cobain/gin-todos/database"
	"github.com/carp-cobain/gin-todos/database/repo"
	"gorm.io/gorm"
)

func createTestDB(t *testing.T) *gorm.DB {
	db, err := database.Connect("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("unable to connect to database: %+v", err)
	}
	if err := database.RunMigrations(db); err != nil {
		t.Fatalf("unable to auto migrate: %+v", err)
	}
	return db
}

func TestStoryRepo(t *testing.T) {
	// Connect to test sqlite db
	db := createTestDB(t)
	storyRepo := repo.NewStoryRepo(db)
	// Create
	story, err := storyRepo.CreateStory("Story 1")
	if err != nil {
		t.Fatalf("unable to create story: %+v", err)
	}
	// Read
	if _, err := storyRepo.GetStory(story.ID); err != nil {
		t.Fatalf("unable to get story by ID: %+v", err)
	}
	_, stories := storyRepo.GetStories(0, 10)
	if len(stories) != 1 || stories[0].ID != story.ID {
		t.Fail()
	}
	// Update
	if _, err := storyRepo.UpdateStory(story.ID, "Story 1 (updated)"); err != nil {
		t.Fatalf("unable to update story: %+v", err)
	}
	// Delete
	if err := storyRepo.DeleteStory(story.ID); err != nil {
		t.Fatalf("unable to delete story: %+v", err)
	}
	// Ensure deleted
	if story, err := storyRepo.GetStory(story.ID); err == nil {
		t.Fatalf("story was not deleted: %+v", story)
	}
}

func TestTaskRepo(t *testing.T) {
	// Connect to test sqlite db
	db := createTestDB(t)
	taskRepo := repo.NewTaskRepo(db)
	storyRepo := repo.NewStoryRepo(db)
	story, err := storyRepo.CreateStory("Test")
	if err != nil {
		t.Fatalf("unable to set up parent story: %+v", err)
	}
	// Create
	task, err := taskRepo.CreateTask(story.ID, "Task")
	if err != nil {
		t.Fatalf("unable to create task: %+v", err)
	}
	// Read
	if _, err := taskRepo.GetTask(task.ID); err != nil {
		t.Fatalf("unable to get task by ID: %+v", err)
	}
	_, tasks := taskRepo.GetTasks(story.ID, 0, 10)
	if len(tasks) != 1 || tasks[0].ID != task.ID {
		t.Fail()
	}
	// Update
	if _, err := taskRepo.UpdateTask(task.ID, task.Title, "complete"); err != nil {
		t.Fatalf("unable to update task: %+v", err)
	}
	// Delete
	if err := taskRepo.DeleteTask(task.ID); err != nil {
		t.Fatalf("unable to delete task: %+v", err)
	}
	// Ensure deleted
	if task, err := taskRepo.GetTask(task.ID); err == nil {
		t.Fatalf("task was not deleted: %+v", task)
	}
	// Cleanup
	if err := storyRepo.DeleteStory(story.ID); err != nil {
		t.Fatalf("cleanup error: %+v", err)
	}
}
