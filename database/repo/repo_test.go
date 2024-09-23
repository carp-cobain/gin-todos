package repo_test

import (
	"testing"

	"github.com/carp-cobain/gin-todos/database"
	"github.com/carp-cobain/gin-todos/database/model"
	"github.com/carp-cobain/gin-todos/database/repo"
	"gorm.io/gorm"
)

func createTestDB(t *testing.T) *gorm.DB {
	db, err := database.Connect("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Story{}, &model.Task{}); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestStoryRepo(t *testing.T) {
	// Connect to test sqlite db
	db := createTestDB(t)
	repo := repo.NewStoryRepo(db)
	// Create
	story, err := repo.CreateStory("Story 1")
	if err != nil {
		t.Fatalf("unable to create story: %+v", err)
	}
	// Read
	if _, err := repo.GetStory(story.ID); err != nil {
		t.Fatalf("unable to get story by ID: %+v", err)
	}
	stories := repo.GetStories(10, 0)
	if len(stories) != 1 || stories[0].ID != story.ID {
		t.Fail()
	}
	// Update
	if _, err := repo.UpdateStory(story.ID, "Story 1 (updated)"); err != nil {
		t.Fatalf("unable to update story: %+v", err)
	}
	// Delete
	if err := repo.DeleteStory(story.ID); err != nil {
		t.Fatalf("unable to delete story: %+v", err)
	}
	// Ensure deleted
	if story, err := repo.GetStory(story.ID); err == nil {
		t.Fatalf("story should have been deleted but was found: %+v", story)
	}
}

func TestTaskRepo(t *testing.T) {
	// Connect to test sqlite db
	db := createTestDB(t)
	storyRepo := repo.NewStoryRepo(db)
	story, err := storyRepo.CreateStory("Test")
	if err != nil {
		t.Fatalf("unable to set up parent story: %+v", err)
	}
	repo := repo.NewTaskRepo(db)
	// Create
	task, err := repo.CreateTask(story.ID, "Task")
	if err != nil {
		t.Fatalf("unable to create task: %+v", err)
	}
	// Read
	if _, err := repo.GetTask(task.ID); err != nil {
		t.Fatalf("unable to get task by ID: %+v", err)
	}
	tasks := repo.GetTasks(story.ID, 10, 0)
	if len(tasks) != 1 || tasks[0].ID != task.ID {
		t.Fail()
	}
	// Update
	if _, err := repo.UpdateTask(task.ID, task.Name, "complete"); err != nil {
		t.Fatalf("unable to update task: %+v", err)
	}
	// Delete
	if err := repo.DeleteTask(task.ID); err != nil {
		t.Fatalf("unable to delete task: %+v", err)
	}
	// Ensure deleted
	if deleted, err := repo.GetTask(task.ID); err == nil {
		t.Fatalf("story should have been deleted but was found: %+v", deleted)
	}
	// Cleanup
	if err := storyRepo.DeleteStory(story.ID); err != nil {
		t.Fatalf("cleanup error: %+v", err)
	}
}
