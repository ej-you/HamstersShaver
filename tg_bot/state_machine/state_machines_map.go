package state_machine


// словарь со всеми машинами состояний юзеров
type userStateMachinesMap map[string]UserStateMachine

// получение машины состояний для юзера по его телеграмм ID (если такой ещё нет, то создание)
func (this *userStateMachinesMap) Get(userTelegramID string) UserStateMachine {
	var userStateMachine UserStateMachine

	userStateMachine, ok := (*this)[userTelegramID]

	// если машины состояний для этого юзера ещё нет, то создаём её
	if !ok {
		userStateMachine.setUserTelegramID(userTelegramID)
		(*this)[userTelegramID] = userStateMachine
	}

	return userStateMachine
}


// экземпляр словаря с машинами состояний
var UserStateMachines = make(userStateMachinesMap)
