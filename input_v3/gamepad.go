package inputv3

import "github.com/hajimehoshi/ebiten/v2"

type GamePadButton struct {
	gamepadID ebiten.GamepadID
	button     ebiten.GamepadButton
}


func NewGamePadButton(gamepadID ebiten.GamepadID, button ebiten.GamepadButton) GamePadButton {
	return GamePadButton{
		gamepadID: gamepadID,
		button:    button,
	}
}
// struct THOR_API JoystickButton
// {
// 	/// @brief Constructor
// 	/// @details Note that you can also construct a joystick id and button property
// 	///  with the following more expressive syntax:
// 	/// @code
// 	/// thor::JoystickButton j = thor::joystick(id).button(b);
// 	/// @endcode
// 								JoystickButton(unsigned int joystickId, unsigned int button);

// 	unsigned int				joystickId;	///< The joystick number
// 	unsigned int				button;		///< The joystick button number
// };


type GamePadAxis struct {
	gamepadID ebiten.GamepadID
	axis      ebiten.GamepadAxisType
	threshold float64
	above     bool // True if position must be above threshold, false if below
}




func NewGamePadAxis(gamepadID ebiten.GamepadID, axis ebiten.GamepadAxisType, threshold float64, above bool) GamePadAxis {
	return GamePadAxis{
		gamepadID: gamepadID,
		axis:      axis,
		threshold: threshold,
		above:     above,
	}
}

// /// @}

// // ---------------------------------------------------------------------------------------------------------------------------

// namespace detail
// {

// 	// Proxy class that allows the joy(id) named parameter syntax
// 	struct THOR_API JoystickBuilder
// 	{
// 		struct THOR_API Axis
// 		{
// 			JoystickAxis					above(float threshold);
// 			JoystickAxis					below(float threshold);

// 			sf::Joystick::Axis				axis;
// 			unsigned int					joystickId;
// 		};

// 		explicit						JoystickBuilder(unsigned int joystickId);
// 		JoystickButton					button(unsigned int button);
// 		Axis							axis(sf::Joystick::Axis axis);

// 		unsigned int					joystickId;
// 	};

// } // namespace detail

// detail::JoystickBuilder THOR_API		joystick(unsigned int joystickId);

// } // namespace thor

// #endif // THOR_JOYSTICK_HPP