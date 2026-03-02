// Package input provides an input manager that maps abstract action names to
// physical buttons (keyboard, mouse, gamepad) and tracks cursor position.
//
// Create an InputManager with NewInputManager. Use AddAction to bind action names
// (e.g. "move_up") to RawInputButton values. Then query IsActionPressed,
// IsActionJustPressed, IsActionJustReleased, IsActionReleased, or IsActionHeld.
// Enable cursor tracking with SetMouseEnabled(true) and read GetCursorPos (screen
// coordinates; use the world's Camera to convert to world space if needed).
// Call Update each frame so cursor and button state are current.
package input
