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
// Enable cursor with SetMouseEnabled(true) and read GetCursorPos.
package input
