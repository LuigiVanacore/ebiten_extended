package inputv3




// template <typename ActionId>
// class ActionMap : private aurora::NonCopyable
// {
// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Public types
// 	public:
// 		/// @brief EventSystem to connect ActionMap with event listeners.
// 		///
// 		typedef EventSystem< ActionContext<ActionId>, ActionId > CallbackSystem;


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Public member functions
// 	public:
// 		/// @brief Default constructor
// 		///
// 									ActionMap();

// 		/// @brief Move constructor
// 		///
// 									ActionMap(ActionMap&& source);

// 		/// @brief Move assignment operator
// 		///
// 		ActionMap&					operator= (ActionMap&& source);

// 		/// @brief Clears old events and polls the window for new ones.
// 		/// @param window The SFML window from which events are polled.
// 		/// @details When you invoke this method, you should not call sf::Window::pollEvent() in the same frame, since update()
// 		///  already does that. The following code
// 		/// @code
// 		/// map.update(window);
// 		/// @endcode
// 		/// is equivalent to
// 		/// @code
// 		/// map.clearEvents();
// 		///
// 		/// sf::Event event;
// 		/// while (window.pollEvent(event))
// 		///     map.pushEvent(event);
// 		/// @endcode
// 		void						update(sf::Window& window);

// 		/// @brief Feeds the action map with a SFML event, that can be used during the current frame.
// 		/// @details When you use the update() method, you needn't invoke pushEvent(). This method exists for more flexibility:
// 		///  You can push user-defined events, and you can do something else with the polled events before calling pushEvent().
// 		void						pushEvent(const sf::Event& event);

// 		/// @brief Removes the events that have been temporarily stored.
// 		/// @details You only need this function in combination with pushEvent(), if you want to feed the action map manually with events.
// 		///  Otherwise, you can just call update().
// 		void						clearEvents();

// 		/// @brief Returns the action associated with a given action identifier.
// 		/// @details Use this function to create new associations. If the action id hasn't been stored yet, it is inserted and
// 		///  an empty action is returned. Assign an object of type thor::Action to it. Example:
// 		/// @code
// 		/// thor::ActionMap<std::string> map(...);
// 		/// map["run"] = thor::Action(sf::Keyboard::R);
// 		/// @endcode
// 		Action&						operator[] (const ActionId& id);

// 		/// @brief Remove the action associated with the specified @a id.
// 		///
// 		void						removeAction(const ActionId& id);

// 		/// @brief Removes all actions associated with this map.
// 		///
// 		void						clearActions();

// 		/// @brief Returns whether the action associated with the specified identifier is currently in effect.
// 		/// @details To be in effect, the boolean operation of the assigned action must yield true. Note that in contrast
// 		///  to registered callbacks, isActive() doesn't take into account the situation where multiple events of the same
// 		///  type occur in a single frame.
// 		/// @code
// 		/// // If LCtrl+A or B is pressed, the action is active.
// 		/// thor::Action(sf::Keyboard::A) && thor::Action(sf::Keyboard::LCtrl) || thor::Action(sf::Keyboard::B);
// 		/// @endcode
// 		bool						isActive(const ActionId& id) const;

// 		/// @brief Forwards active actions to a callback system.
// 		/// @param system Callback system of type EventSystem< ActionContext<ActionId>, ActionId >
// 		/// @param window Window pointer which is stored in the ActionContext passed to the callbacks. Can be nullptr.
// 		/// @details For every action that is currently active, the action ID is passed to @a system, where all listener
// 		///  functions associated with the ID are invoked. For a given action registered in the action map, the callback system
// 		///  is invoked at most once, even if the action has been combined with logical operators and multiple sub-actions are
// 		///  active. That is, an action Shift&&X won't lead to an invocation for both Shift and X. This is the expected behavior.
// 		/// @code
// 		/// // Listener function for "run" actions
// 		/// void callback(thor::ActionContext<std::string> context);
// 		///
// 		/// // Map to register "run" action
// 		/// thor::ActionMap<std::string> map;
// 		/// map["run"] = thor::Action(sf::Keyboard::R);
// 		///
// 		/// // Create EventSystem object, connect "run" action to callback
// 		/// thor::ActionMap<std::string>::CallbackSystem system;
// 		/// system.connect("run", &callback);
// 		///
// 		/// // In the main loop: Forward actions to the callback system
// 		/// map.invokeCallbacks(system, &window);
// 		/// @endcode
// 		/// @warning While this function is being called, any modifications to @a system result in undefined behavior. That is,
// 		///  do not remove or insert callbacks during the invocation of another callback. If you need to modify the system, delay
// 		///  the modifications until invokeCallbacks() returns.
// 		/// @see thor::ActionContext
// 		void						invokeCallbacks(CallbackSystem& system, sf::Window* window) const;


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Private types
// 	private:
// 		typedef std::map<ActionId, Action> Map;


// 	// ---------------------------------------------------------------------------------------------------------------------------
// 	// Private variables
// 	private:
// 		Map							mActionMap;
// 		detail::EventBuffer			mEventBuffer;
// };

// /// @}

// }

type ActionID int

type ActionMap struct {
	actionMap  map[ActionID]Action
}



func NewActionMap() *ActionMap {
	return &ActionMap{
		actionMap: make(map[ActionID]Action),
	}
}


func (am *ActionMap) AddAction(id ActionID, action Action) {
	am.actionMap[id] = action
}

func (am *ActionMap) RemoveAction(id ActionID) {
	delete(am.actionMap, id)
}	

func (am *ActionMap) GetAction(id ActionID) (Action, bool) {
	action, exists := am.actionMap[id]
	return action, exists
}

func (am *ActionMap) ClearActions() {
	am.actionMap = make(map[ActionID]Action)
}

