package input




// namespace InputMapping
// {

// 	// Forward declarations
// 	class InputContext;

type MappedInput struct {
	actions []Action
	states []State
	ranges map[Range]float64
}

func (m *MappedInput) EatAction(action Action) {
	for i, a := range m.actions {
		if a == action {
			m.actions = append(m.actions[:i], m.actions[i+1:]...)
			return
		}
	}
}

func (m *MappedInput) EatState(state State) {
	for i, s := range m.states {
		if s == state {
			m.states = append(m.states[:i], m.states[i+1:]...)
			return
		}
	}
}

func (m *MappedInput) EatRange(range_ Range) {
	delete(m.ranges, range_)
}

// 	// Helper structure
// 	struct MappedInput
// 	{
// 		std::set<Action> Actions;
// 		std::set<State> States;
// 		std::map<Range, double> Ranges;

// 		// Consumption helpers
// 		void EatAction(Action action)		{ Actions.erase(action); }
// 		void EatState(State state)			{ States.erase(state); }
// 		void EatRange(Range range)
// 		{
// 			std::map<Range, double>::iterator iter = Ranges.find(range);
// 			if(iter != Ranges.end())
// 				Ranges.erase(iter);
// 		}
// 	};


type InputCallback func(MappedInput)


type InputMapper struct {
	inputContexts map[string]*InputContext
	activeContexts []*InputContext
	callbackTable map[int]InputCallback
	currentMappedInput MappedInput
}

func NewInputMapper() *InputMapper {
	return &InputMapper{
		inputContexts: make(map[string]*InputContext),
		callbackTable: make(map[int]InputCallback),
	}
}

func (i *InputMapper) Clear() {
	i.currentMappedInput.actions = []Action{}
	i.currentMappedInput.states = []State{}
	i.currentMappedInput.ranges = make(map[Range]float64)
}

// func (i *InputMapper) SetRawButtonState(button RawInputButton, pressed, previouslypressed bool) {
// 	var action Action
// 	var state State

// 	if pressed && !previouslypressed {
// 		if i.mapButtonToAction(button, &action) {
// 			i.currentMappedInput.actions = append(i.currentMappedInput.actions, action)
// 			return
// 		}
// 	}

// 	if pressed {
// 		if i.mapButtonToState(button, &state) {
// 			i.currentMappedInput.states = append(i.currentMappedInput.states, state)
// 			return
// 		}
// 	}

// 	i.mapAndEatButton(button)
// }

// func (i *InputMapper) SetRawAxisValue(axis RawInputAxis, value float64) {
// 	for _, context := range i.activeContexts {
// 		var _range Range
// 		if context.mapAxisToRange(axis, &_range) {
// 			i.currentMappedInput.ranges[_range] = value * context.GetSensitivity(_range)
// 			break
// 		}
// 	}
// }

func (i *InputMapper) Dispatch() {
	input := i.currentMappedInput
	for _, callback := range i.callbackTable {
		callback(input)
	}
}

func (i *InputMapper) AddCallback(callback InputCallback, priority int) {
	i.callbackTable[priority] = callback
}

func (i *InputMapper) PushContext(name string) {
	context, ok := i.inputContexts[name]
	if !ok {
		panic("Invalid input context pushed")
	}

	i.activeContexts = append([]*InputContext{context}, i.activeContexts...)
}

func (i *InputMapper) PopContext() {
	if len(i.activeContexts) == 0 {
		panic("Cannot pop input context, no contexts active!")
	}

	i.activeContexts = i.activeContexts[1:]
}

// func (i *InputMapper) mapButtonToAction(button RawInputButton, action *Action) bool {
// 	for _, context := range i.activeContexts {
// 		if context.mapButtonToAction(button, action) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (i *InputMapper) mapButtonToState(button RawInputButton, state *State) bool {
// 	for _, context := range i.activeContexts {
// 		if context.mapButtonToState(button, state) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (i *InputMapper) mapAndEatButton(button RawInputButton) {
// 	var action Action
// 	var state State

// 	if i.mapButtonToAction(button, &action) {
// 		i.currentMappedInput.EatAction(action)
// 	}

// 	if i.mapButtonToState(button, &state) {
// 		i.currentMappedInput.EatState(state)
// 	}
// }



// 	// Handy type shortcuts
// 	typedef void (*InputCallback)(MappedInput& inputs);


// 	class InputMapper
// 	{
// 	// Construction and destruction
// 	public:
// 		InputMapper();
// 		~InputMapper();

// 	// Raw input interface
// 	public:
// 		void Clear();
// 		void SetRawButtonState(RawInputButton button, bool pressed, bool previouslypressed);
// 		void SetRawAxisValue(RawInputAxis axis, double value);

// 	// Input dispatching interface
// 	public:
// 		void Dispatch() const;

// 	// Input callback registration interface
// 	public:
// 		void AddCallback(InputCallback callback, int priority);

// 	// Context management interface
// 	public:
// 		void PushContext(const std::wstring& name);
// 		void PopContext();

// 	// Internal helpers
// 	private:
// 		bool MapButtonToAction(RawInputButton button, Action& action) const;
// 		bool MapButtonToState(RawInputButton button, State& state) const;
// 		void MapAndEatButton(RawInputButton button);

// 	// Internal tracking
// 	private:
// 		std::map<std::wstring, InputContext*> InputContexts;
// 		std::list<InputContext*> ActiveContexts;

// 		std::multimap<int, InputCallback> CallbackTable;

// 		MappedInput CurrentMappedInput;
// 	};

// }



// InputMapper::InputMapper()
// {
// 	unsigned count;
// 	std::wifstream infile(L"ContextList.txt");
// 	if(!(infile >> count))
// 		throw std::exception("Failed to read ContextList.txt");
// 	for(unsigned i = 0; i < count; ++i)
// 	{
// 		std::wstring name = AttemptRead<std::wstring>(infile);
// 		std::wstring file = AttemptRead<std::wstring>(infile);
// 		InputContexts[name] = new InputContext(file);
// 	}
// }

// //
// // Destruct and clean up an input mapper
// //
// InputMapper::~InputMapper()
// {
// 	for(std::map<std::wstring, InputContext*>::iterator iter = InputContexts.begin(); iter != InputContexts.end(); ++iter)
// 		delete iter->second;
// }


// //
// // Clear all mapped input
// //
// void InputMapper::Clear()
// {
// 	CurrentMappedInput.Actions.clear();
// 	CurrentMappedInput.Ranges.clear();
// 	// Note: we do NOT clear states, because they need to remain set
// 	// across frames so that they don't accidentally show "off" for
// 	// a tick or two while the raw input is still pending.
// }

// //
// // Set the state of a raw button
// //
// void InputMapper::SetRawButtonState(RawInputButton button, bool pressed, bool previouslypressed)
// {
// 	Action action;
// 	State state;

// 	if(pressed && !previouslypressed)
// 	{
// 		if(MapButtonToAction(button, action))
// 		{
// 			CurrentMappedInput.Actions.insert(action);
// 			return;
// 		}
// 	}

// 	if(pressed)
// 	{
// 		if(MapButtonToState(button, state))
// 		{
// 			CurrentMappedInput.States.insert(state);
// 			return;
// 		}
// 	}

// 	MapAndEatButton(button);
// }

// //
// // Set the raw axis value of a given axis
// //
// void InputMapper::SetRawAxisValue(RawInputAxis axis, double value)
// {
// 	for(std::list<InputContext*>::const_iterator iter = ActiveContexts.begin(); iter != ActiveContexts.end(); ++iter)
// 	{
// 		const InputContext* context = *iter;

// 		Range range;
// 		if(context->MapAxisToRange(axis, range))
// 		{
// 			CurrentMappedInput.Ranges[range] = context->GetConversions().Convert(range, value * context->GetSensitivity(range));
// 			break;
// 		}
// 	}
// }


// //
// // Dispatch input to all registered callbacks
// //
// void InputMapper::Dispatch() const
// {
// 	MappedInput input = CurrentMappedInput;
// 	for(std::multimap<int, InputCallback>::const_iterator iter = CallbackTable.begin(); iter != CallbackTable.end(); ++iter)
// 		(*iter->second)(input);
// }

// //
// // Add a callback to the dispatch table
// //
// void InputMapper::AddCallback(InputCallback callback, int priority)
// {
// 	CallbackTable.insert(std::make_pair(priority, callback));
// }


// //
// // Push an active input context onto the stack
// //
// void InputMapper::PushContext(const std::wstring& name)
// {
// 	std::map<std::wstring, InputContext*>::iterator iter = InputContexts.find(name);
// 	if(iter == InputContexts.end())
// 		throw std::exception("Invalid input context pushed");

// 	ActiveContexts.push_front(iter->second);
// }

// //
// // Pop the current input context off the stack
// //
// void InputMapper::PopContext()
// {
// 	if(ActiveContexts.empty())
// 		throw std::exception("Cannot pop input context, no contexts active!");

// 	ActiveContexts.pop_front();
// }


// //
// // Helper: map a button to an action in the active context(s)
// //
// bool InputMapper::MapButtonToAction(RawInputButton button, Action& action) const
// {
// 	for(std::list<InputContext*>::const_iterator iter = ActiveContexts.begin(); iter != ActiveContexts.end(); ++iter)
// 	{
// 		const InputContext* context = *iter;

// 		if(context->MapButtonToAction(button, action))
// 			return true;
// 	}

// 	return false;
// }

// //
// // Helper: map a button to a state in the active context(s)
// //
// bool InputMapper::MapButtonToState(RawInputButton button, State& state) const
// {
// 	for(std::list<InputContext*>::const_iterator iter = ActiveContexts.begin(); iter != ActiveContexts.end(); ++iter)
// 	{
// 		const InputContext* context = *iter;

// 		if(context->MapButtonToState(button, state))
// 			return true;
// 	}

// 	return false;
// }

// //
// // Helper: eat all input mapped to a given button
// //
// void InputMapper::MapAndEatButton(RawInputButton button)
// {
// 	Action action;
// 	State state;

// 	if(MapButtonToAction(button, action))
// 		CurrentMappedInput.EatAction(action);

// 	if(MapButtonToState(button, state))
// 		CurrentMappedInput.EatState(state);
// }
