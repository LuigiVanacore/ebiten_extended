package ebiten_extended

import (
	"testing"
)

func TestResourceManagerRemoveImage(t *testing.T) {
	rm := NewResourceManager()
	rm.images["hero"] = nil

	if removed := rm.RemoveImage("hero"); !removed {
		t.Fatal("expected RemoveImage to return true for existing image")
	}
	if got := rm.GetImage("hero"); got != nil {
		t.Fatal("expected image to be removed")
	}
	if removed := rm.RemoveImage("hero"); removed {
		t.Fatal("expected RemoveImage to return false for missing image")
	}
}

func TestResourceManagerClearImages(t *testing.T) {
	rm := NewResourceManager()
	rm.images["hero"] = nil
	rm.images["enemy"] = nil

	removed := rm.ClearImages()
	if removed != 2 {
		t.Fatalf("expected 2 removed images, got %d", removed)
	}
	if len(rm.GetImages()) != 0 {
		t.Fatalf("expected no images after clear, got %d", len(rm.GetImages()))
	}
}

func TestResourceManagerClearFonts(t *testing.T) {
	rm := NewResourceManager()
	rm.fonts["f1"] = nil
	rm.fonts["f2"] = nil

	removed := rm.ClearFonts()
	if removed != 2 {
		t.Fatalf("expected 2 removed fonts, got %d", removed)
	}
	if got := rm.GetFont("f1"); got != nil {
		t.Fatal("expected no font at key f1 after clear")
	}
}

func TestResourceManagerLoadImageFromFileInvalidPath(t *testing.T) {
	rm := NewResourceManager()
	err := rm.LoadImageFromFile("test", "nonexistent_image_file_12345.png")
	if err == nil {
		t.Error("expected error when loading from nonexistent path")
	}
	if rm.GetImage("test") != nil {
		t.Error("image should not be cached on error")
	}
}

func TestResourceManagerLoadFontFromFileInvalidPath(t *testing.T) {
	rm := NewResourceManager()
	err := rm.LoadFontFromFile("test", "nonexistent_font_file_12345.otf", 16)
	if err == nil {
		t.Error("expected error when loading from nonexistent path")
	}
	if rm.GetFont("test") != nil {
		t.Error("font should not be cached on error")
	}
}

func TestResourceManagerClear(t *testing.T) {
	rm := NewResourceManager()
	rm.images["hero"] = nil
	rm.fonts["hero_font"] = nil

	rm.Clear()

	if len(rm.GetImages()) != 0 {
		t.Fatalf("expected no images after clear, got %d", len(rm.GetImages()))
	}
	if got := rm.GetFont("hero_font"); got != nil {
		t.Fatal("expected no fonts after clear")
	}
}

func TestResourceManagerPreloadBatch(t *testing.T) {
	rm := NewResourceManager()
	assets := []PreloadAsset{
		{ID: "img1", Path: "nonexistent_1.png", Type: AssetTypeImage},
		{ID: "img2", Path: "nonexistent_2.png", Type: AssetTypeImage},
	}

	completeCh := make(chan []string)
	var progressCount int

	rm.PreloadBatch(assets, func(loaded, total int, lastID string) {
		progressCount++
	}, func(failed []string) {
		completeCh <- failed
	})

	failed := <-completeCh
	if len(failed) != 2 {
		t.Errorf("expected 2 failed assets, got %d", len(failed))
	}
	if progressCount != 2 {
		t.Errorf("expected progress callback to be called 2 times, got %d", progressCount)
	}
}
