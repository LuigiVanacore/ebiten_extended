package ui

import (
	"bytes"
	"image/color"
	"testing"

	exampleresources "github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func TestButtonNode_SetSize_GetSize(t *testing.T) {
	im := input.NewInputManager()
	btn := NewButtonNode("btn", 100, 50, im)
	w, h := btn.GetSize()
	if w != 100 || h != 50 {
		t.Errorf("GetSize: got (%v,%v) want (100,50)", w, h)
	}
	btn.SetSize(200, 80)
	w, h = btn.GetSize()
	if w != 200 || h != 80 {
		t.Errorf("SetSize: got (%v,%v) want (200,80)", w, h)
	}
}

func TestButtonNode_SetText(t *testing.T) {
	tt, err := text.NewGoTextFaceSource(bytes.NewReader(exampleresources.DefaultFont))
	if err != nil {
		t.Skip("DefaultFont not available:", err)
	}
	face := &text.GoTextFace{Source: tt, Size: 16}
	im := input.NewInputManager()
	btn := NewButtonNode("btn", 100, 50, im)
	btn.SetText("Click", face, color.White)
	if btn.label == nil {
		t.Error("SetText: label should be created")
	}
	if btn.label.GetMessage() != "Click" {
		t.Errorf("SetText: got %q", btn.label.GetMessage())
	}
}

func TestButtonNode_NilInputManager(t *testing.T) {
	btn := NewButtonNode("btn", 100, 50, nil)
	// Should not panic
	btn.Update()
}

func TestCheckboxNode_InitialState(t *testing.T) {
	im := input.NewInputManager()
	chk := NewCheckboxNode("chk", 30, im)
	if chk.Checked {
		t.Error("new CheckboxNode should be unchecked")
	}
}

func TestCheckboxNode_OnToggle(t *testing.T) {
	im := input.NewInputManager()
	chk := NewCheckboxNode("chk", 30, im)
	var toggled bool
	var checkedVal bool
	chk.OnToggle = func(checked bool) {
		toggled = true
		checkedVal = checked
	}
	chk.ButtonNode.OnClick()
	if !toggled || !checkedVal {
		t.Errorf("OnToggle: toggled=%v checked=%v", toggled, checkedVal)
	}
}

func TestProgressBarNode_SetProgress(t *testing.T) {
	pb := NewProgressBarNode("pb", 200, 20)
	pb.SetProgress(0.5)
	if pb.GetProgress() != 0.5 {
		t.Errorf("GetProgress: got %v", pb.GetProgress())
	}
	pb.SetProgress(2) // clamp to 1
	if pb.GetProgress() != 1 {
		t.Errorf("clamp: got %v", pb.GetProgress())
	}
}

func TestDropdownNode_SetItems_GetSelected(t *testing.T) {
	im := input.NewInputManager()
	dd := NewDropdownNode("dd", 150, 30, nil, im)
	dd.SetItems([]string{"A", "B", "C"})
	if len(dd.GetItems()) != 3 {
		t.Errorf("GetItems: got %d", len(dd.GetItems()))
	}
	if dd.GetSelectedIndex() != -1 {
		t.Errorf("initial selected: got %d", dd.GetSelectedIndex())
	}
	if dd.GetSelectedText() != "" {
		t.Errorf("initial text: got %q", dd.GetSelectedText())
	}
	dd.SetSelectedIndex(1)
	if dd.GetSelectedIndex() != 1 {
		t.Errorf("SetSelectedIndex: got %d", dd.GetSelectedIndex())
	}
	if dd.GetSelectedText() != "B" {
		t.Errorf("GetSelectedText: got %q", dd.GetSelectedText())
	}
	dd.SetItems([]string{"X"})
	if dd.GetSelectedIndex() != -1 {
		t.Errorf("invalid index after SetItems: got %d", dd.GetSelectedIndex())
	}
}

func TestDropdownNode_NilInputManager(t *testing.T) {
	dd := NewDropdownNode("dd", 100, 30, nil, nil)
	dd.SetItems([]string{"A"})
	dd.Update()
}

func TestTooltipNode_SetText(t *testing.T) {
	im := input.NewInputManager()
	panel := NewPanelNode("p", 50, 50)
	tooltip := NewTooltipNode("tip", panel, "help", nil, im)
	if tooltip.text != "help" {
		t.Errorf("tooltip text: got %q", tooltip.text)
	}
	tooltip.SetText("new help")
	if tooltip.text != "new help" {
		t.Errorf("SetText: got %q", tooltip.text)
	}
}

func TestTextInputNode_Clipboard(t *testing.T) {
	im := input.NewInputManager()
	txt := NewTextInputNode("txt", 200, 30, nil, im)
	txt.SetText("hello")
	txt.cursorIndex = 2
	txt.selStart, txt.selEnd = 1, 3
	txt.copySelection()
	if txt.clipboard != "el" {
		t.Errorf("copySelection: got %q", txt.clipboard)
	}
	txt.paste()
	if txt.GetText() != "hello" {
		t.Errorf("paste (replace el with el): got %q", txt.GetText())
	}
	txt.SetText("hxxo")
	txt.cursorIndex = 3
	txt.selStart, txt.selEnd = 1, 3
	txt.paste()
	if txt.GetText() != "helo" {
		t.Errorf("paste (replace xx with el): got %q", txt.GetText())
	}
}
