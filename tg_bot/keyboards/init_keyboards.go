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

	// клавиатура для команды /trade
	InlineKeyboardTrade.Inline(
		// кнопка для выбора варианта покупки монет (аналог - команда /buy)
		InlineKeyboardTrade.Row(BtnToBuy),
		// кнопка для выбора варианта продажи монет (аналог - команда /cell)
		InlineKeyboardTrade.Row(BtnToCell),
	)

	// клавиатура для выбора DEX-биржи
	InlineKeyboardChooseDEX.Inline(
		// кнопка для выбора stonfi
		InlineKeyboardChooseDEX.Row(BtnStonfi),
		// кнопка для выбора dedust
		InlineKeyboardChooseDEX.Row(BtnDedust),
	)

	// клавиатура для выбора DEX-биржи
	InlineKeyboardConfirmNewTransaction.Inline(
		// кнопка для подтверждения
		InlineKeyboardConfirmNewTransaction.Row(BtnConfirm),
		// кнопка для отмены
		InlineKeyboardConfirmNewTransaction.Row(BtnCancel),
	)

	// клавиатура с выбором процента проскальзывания
	InlineKeyboardSlippageChoices.Inline(
		// кнопки с разными процентами
		InlineKeyboardSlippageChoices.Row(BtnSlippageChoice15, BtnSlippageChoice30, BtnSlippageChoice60, BtnSlippageChoice100),
	)

	// клавиатура с выбором процента от общего числа монет на балансе аккаунта для продажи
	InlineKeyboardJettonsAmountChoices.Inline(
		// кнопки с разными процентами
		InlineKeyboardJettonsAmountChoices.Row(
			BtnJettonsAmountChoice25,
			BtnJettonsAmountChoice50,
			BtnJettonsAmountChoice75,
			BtnJettonsAmountChoice100,
		),
	)
}
