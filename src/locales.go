package main

// array[0] - english, array[1] - russian

var ( // keyboards
	cipher_locale             = []string{"Ciphers", "Шифры"}
	chords_locale             = []string{"Chords", "Аккорды"}
	chess_locale              = []string{"Chess", "Шахматы"}
	help_locale               = []string{"Help", "Помощь"}
	author_locale             = []string{"Author", "Автор"}
	version_locale            = []string{"Version", "Версия"}
	time_locale               = []string{"Time", "Время"}
	date_locale               = []string{"Date", "Дата"}
	back_locale               = []string{"Back", "Назад"}
	caesar_locale             = []string{"Caesar", "Цезарь"}
	vigenere_locale           = []string{"Vigenere", "Виженер"}
	encode_locale             = []string{"Encode", "Зашифровать"}
	decode_locale             = []string{"Decode", "Расшифровать"}
	artists_locale            = []string{"Artists", "Исполнители"}
	other_locale              = []string{"Other", "Другое"}
	chess_puzzles_locale      = []string{"Chess puzzles", "Шахматные задачи"}
	mate_locale               = []string{"Mate in", "Мат в"}
	move_locale               = []string{"move", "ход"}
	moves_locale              = []string{"moves", "хода"}
	strings_order_locale      = []string{"Guitar strings order", "Порядок струн"}
	choose_locale             = []string{"Choose", "Выбери"}
	result_locale             = []string{"Result", "Результат"}
	today_locale              = []string{"Today", "Сегодня"}
	in_nsk_locale             = []string{"in Novosibirsk", "в Новосибирске"}
	working_locale            = []string{"I'm working on", "Я работаю на"}
	shift_message_locale      = []string{"Enter the letters shift value (e.g. to transform A to B - 1, A to Z - 25)", "Введите символьное смещение (например, для преобразования А в Б — 1, А в Я — 32)"}
	shift_value_mustbe_locale = []string{"Shift value must be", "Значение смещения должно быть"}
	numeric_locale            = []string{"a number", "числовым"}
	natural_locale            = []string{"more than zero", "больше нуля"}
	enter_cipher_key_locale   = []string{"Enter the cipher key", "Введите ключ шифрования"}
	choose_puzzle_type_locale = []string{"Choose the puzzle type", "Выбери тип шахматной задачи"}
	no_commands_locale        = []string{"I don't have such command. Use the buttons!", "У меня нет такой команды. Используй кнопки!"}
	language_locale           = []string{"Language", "Язык"}
	first_move_locale         = []string{"Enter the first move", "Введите первый ход"}
	coding_locale             = []string{"coding", "шифрования"}
	en_locale                 = []string{"en", ""}
	de_locale                 = []string{"de", "де"}
	enter_phrase_locale       = []string{"Enter a phrase to ", "Введите фразу для "}
	chess_phrases             = [][]string{
		{"❌ Wrong move\nTry another one", "❌ Неверный ход\nПопробуйте другой"},
		{"✅ A nice start. Enter the second move", "✅ Отличное начало. Введите второй ход"},
		{"✅ Well done! Enter the third move", "✅ Замечательно! Введите третий ход"},
		{"🎂 Congrats!\n✅ You've completed this chess puzzle ♟", "🎂 Поздравляю!\n✅ Вы выполнили шахматную задачу ♟"},
		{"✅ Third move is right! Enter the fourth move", "✅ Третий ход верный! Введите четвёртый ход"},
	}
	chess_task_locale = []string{"Task is to checkmate the enemy in ", "Задача — поставить мат за "}
	plural_locale     = []string{"s", "а"}
	chess_move_locale = []string{"To enter the move write: previous_place-new_place\nExample: e2-e4", "Формат хода (в сообщении): прошлое_место_фигуры-новое_место_фигуры\nПример: e2-e4"}
)
