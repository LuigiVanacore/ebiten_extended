// Package input provides a unified input manager for keyboard, mouse, and gamepad.
//
// Two APIs are available:
//
// 1) String-based actions: AddAction(name, RawInputButton) maps names to physical
// inputs. Use IsActionPressed, IsActionJustPressed, IsActionJustReleased, IsActionReleased,
// or IsActionHeld. RawInputButton supports NewKeyRawInputButton, NewMouseRawInputButton,
// NewGamePadRawInputButton.
//
// 2) ActionID-based: RegisterAction(ActionID, Action) with NewKeyAction, NewMouseButtonAction,
// NewGamePadButtonAction, or NewJoystickAxisAction. Use IsActionActive(ActionID).
// Action modes: ActionHold, ActionPressOnce, ActionReleaseOnce.
//
// Gamepad: GetLeftStick, GetRightStick (with StickDeadzone), GamepadIDs,
// IsGamepadButtonPressed, IsGamepadButtonJustPressed, IsStandardGamepadButtonPressed.
//
// Get InputManager from Engine.Input() or create with NewInputManager. Call Update each frame.
// Enable cursor with SetMouseEnabled(true).
//
// Mouse coordinates:
//   - GetCursorPos() returns screen/window coordinates (device pixels, origin top-left).
//     Use for HUD, UI overlays, or when rendering at 1:1 with the window.
//   - Camera.GetCursorCoords(inputMgr) returns world coordinates (scaled by camera zoom/rotation).
//     Use when you need the cursor position in your game world (e.g. click-to-move, OverlapPoint).
package input
