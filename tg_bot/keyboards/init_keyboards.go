package keyboards


func InitKeyboards() {
	// клавиатура для команды /start
	InlineKeyboardToHome.Inline(
		// кнопка для перехода в главное меню (аналог - команда /home)
		InlineKeyboardToHome.Row(BtnToHome),
	)

	// клавиатура для команды /help
	InlineKeyboardHelp.Inline(
		// кнопка для сокрытия справки (удаление сообщения справки)
		InlineKeyboardHelp.Row(BtnHideHelp),
	)

	// клавиатура для команды /home
	InlineKeyboardMainMenu.Inline(
		// кнопка для перехода в "трейдинг" диалог (аналог - команда /trade)
		InlineKeyboardMainMenu.Row(BtnToTrade),
		// кнопка для перехода в "авто" диалог (аналог - команда /auto)
		InlineKeyboardMainMenu.Row(BtnToAuto),
		// кнопка для перехода в управление сохранёнными токенами (аналог - команда /tokens)
		InlineKeyboardMainMenu.Row(BtnToTokens),
	)
}
