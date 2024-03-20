package input

import "github.com/hajimehoshi/ebiten/v2"

const (
	Hold = iota
	PRESSED
	RELEASED
)

const (
	KEY_BUTTON = iota
	MOUSE_BUTTON
	GAMEPAD_BUTTON
)

type Action struct {
	inputType     int
	buttonType    int
	keyButton     ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton ebiten.GamepadButton
	gamepadID     ebiten.GamepadID
}

func NewActionKey(keyButton ebiten.Key, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: KEY_BUTTON, keyButton: keyButton}
}

func NewActionMouse(mouseButton ebiten.MouseButton, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: MOUSE_BUTTON, mouseButton: mouseButton}
}

func NewActionGamepad(gamepadButton ebiten.GamepadButton, gamepadID ebiten.GamepadID, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: GAMEPAD_BUTTON, gamepadButton: gamepadButton, gamepadID: gamepadID}
}

func (a *Action) Test() bool {
	res := false

	if a.buttonType == KEY_BUTTON {
		if a.inputType == PRESSED {
			res = ebiten.IsKeyPressed(a.keyButton)
		}
	} else if a.buttonType == MOUSE_BUTTON {
		if a.inputType == PRESSED {
			res = ebiten.IsMouseButtonPressed(a.mouseButton)
		}
	}
	return res
}



// class THOR_API Action
// {
// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Public types
// 	public:

// 		/// @brief Type for actions
// 		///
// 		enum ActionType
// 		{
// 			Hold,			///< Repeated input (e.g. a key that is held down).
// 			PressOnce,		///< Press events that occur only once (e.g. key pressed).
// 			ReleaseOnce,	///< Release events that occur only once (e.g. key released).
// 		};


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Public member functions
// 	public:
// 		/// @brief Construct key action
// 		/// @details Creates an action that is in effect when @a key is manipulated. The second parameter specifies whether
// 		///  KeyPressed events, KeyReleased events or sf::Keyboard::isKeyPressed() act as action source.
// 		explicit					Action(sf::Keyboard::Key key, ActionType action = Hold);

// 		/// @brief Construct mouse button action
// 		/// @details Creates an action that is in effect when @a mouseButton is manipulated. The second parameter specifies whether
// 		///  MouseButtonPressed events, MouseButtonReleased events or sf::Mouse::isButtonPressed() act as action source.
// 		explicit					Action(sf::Mouse::Button mouseButton, ActionType action = Hold);

// 		/// @brief Construct joystick button action
// 		/// @details Creates an action that is in effect when the joystick button stored in @a joystickState is manipulated.
// 		///  The second parameter specifies whether JoyButtonPressed events, JoyButtonReleased events or
// 		///  sf::Joystick::isButtonPressed() act as action source.
// 		explicit					Action(JoystickButton joystickState, ActionType action = Hold);

// 		/// @brief Construct joystick axis action
// 		/// @details Creates an action that is in effect when the absolute value of the joystick axis position exceeds a threshold
// 		///  (both axis and threshold are stored in @a joystickAxis). The source of the action is sf::Joystick::getAxisPosition()
// 		///  and not JoystickMoved events. This implies that the action will also be active if the axis remains unchanged in a
// 		///  position above the threshold.
// 		explicit					Action(JoystickAxis joystickAxis);

// 		/// @brief Construct SFML event action
// 		/// @details Creates an action that is in effect when a SFML event of the type @a eventType is fired.
// 		explicit					Action(sf::Event::EventType eventType);


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Implementation details
// 	public:
// 		// Default constructor: Required for ActionMap::operator[] - constructs an uninitialized action
// 									Action();

// 		// Tests if the event/real time input constellation in the argument is true for this action
// 		bool						isActive(const detail::EventBuffer& buffer) const;

// 		// Test if active and store relevant events
// 		bool						isActive(const detail::EventBuffer& buffer, detail::ActionResult& out) const;


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Private member functions
// 	private:
// 		// Construct from movable pointer to the action's operation
// 		explicit					Action(detail::ActionNode::CopiedPtr operation);


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Private variables
// 	private:
// 		detail::ActionNode::CopiedPtr	mOperation;


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Friends
// 	/// @cond FriendsAreAnImplementationDetail
// 	friend Action THOR_API operator|| (const Action& lhs, const Action& rhs);
// 	friend Action THOR_API operator&& (const Action& lhs, const Action& rhs);
// 	friend Action THOR_API operator! (const Action& action);
// 	friend Action THOR_API eventAction(std::function<bool(const sf::Event&)> filter);
// 	friend Action THOR_API realtimeAction(std::function<bool()> filter);
// 	/// @endcond
// };


// /// @relates Action
// /// @brief OR operator of two actions: The resulting action is in effect if at least one of @a lhs and @a rhs is active.
// /// @details This operator can be used when multiple input types are assigned to an action. For example, an action that can be triggered
// ///  using either the keyboard or joystick. Another typical example is to combine both modifier keys, such as left and right shift:
// /// @code
// /// thor::Action shift = thor::Action(sf::Keyboard::LShift) || thor::Action(sf::Keyboard::RShift);
// /// thor::Action x(sf::Keyboard::X, thor::Action::PressOnce);
// /// thor::Action shiftX = shift && x;
// /// @endcode
// Action THOR_API				operator|| (const Action& lhs, const Action& rhs);

// /// @relates Action
// /// @brief AND operator of two actions: The resulting action is in effect if both @a lhs and @a rhs are active.
// /// @details This operator is typically used to implement key combinations such as Shift+X. It does not make sense if both operands
// ///  are event actions, because each of them is only active at one time point, hardly together. Instead, implement modifiers as realtime
// ///  actions and the actual keys as event actions:
// /// @code
// /// thor::Action shift = thor::Action(sf::Keyboard::LShift) || thor::Action(sf::Keyboard::RShift);
// /// thor::Action x(sf::Keyboard::X, thor::Action::PressOnce);
// /// thor::Action shiftX = shift && x;
// /// @endcode
// Action THOR_API				operator&& (const Action& lhs, const Action& rhs);

// /// @relates Action
// /// @brief NOT operator of an action: The resulting action is in effect if @a action is not active.
// /// @details This operator is rarely needed. It can be used to discriminate two actions, for example if X and Shift+X have different
// ///  assignments and you don't want Shift+X to trigger also the action assigned to X.
// /// @code
// /// thor::Action shift = thor::Action(sf::Keyboard::LShift) || thor::Action(sf::Keyboard::RShift);
// /// thor::Action x(sf::Keyboard::X, thor::Action::PressOnce);
// /// thor::Action shiftX = shift && x;
// /// thor::Action onlyX = !shift && x;
// /// @endcode
// Action THOR_API				operator! (const Action& action);

// /// @relates Action
// /// @brief Creates a custom action that operates on events
// /// @param filter Functor that is called for every event (which is passed as a parameter). It shall return true when the
// ///  passed event makes the action active.
// /// @details Code example: An action that is active when the X key is pressed. This is just an example, in this specific case you
// ///  should prefer the equivalent expression <i>thor::Action(sf::Keyboard::X, thor::Action::PressOnce)</i>.
// /// @code
// /// bool isXPressed(const sf::Event& event)
// /// {
// ///     return event.type == sf::Event::KeyPressed && event.key.code == sf::Keyboard::X;
// /// }
// ///
// /// thor::Action xPressed = thor::eventAction(&isXPressed);
// /// @endcode
// Action THOR_API				eventAction(std::function<bool(const sf::Event&)> filter);

// /// @relates Action
// /// @brief Creates a custom action that operates on realtime input
// /// @param filter Functor that is called exactly once per frame, independent of any events. It shall return true when a certain realtime
// ///  input state should make the action active.
// /// @details Code example: An action that is active as long as the X key is held down. This is just an example, in this specific case you
// ///  should prefer the equivalent expression <i>thor::Action(sf::Keyboard::X, thor::Action::Hold)</i>.
// /// @code
// /// bool isXHeldDown()
// /// {
// ///     return sf::Keyboard::isKeyPressed(sf::Keyboard::X);
// /// }
// ///
// /// thor::Action xHeldDown = thor::realtimeAction(&isXHeldDown);
// /// @endcode
// Action THOR_API				realtimeAction(std::function<bool()> filter);


// /// @}

// } // namespace thor
