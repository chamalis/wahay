// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package gui

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"es": &dictionary{index: esIndex, data: esData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%s\nMeeting ID: %s%%0D%%0A": 11,
	"Accept":                     30,
	"Allow the host to autojoin a meeting when creates new one": 31,
	"An error occurred\n\n%s":                                   16,
	"Are you sure you want to do this action?":                  32,
	"Are you sure you want to end this meeting?":                33,
	"Are you sure you want to leave this meeting?":              34,
	"Autojoin to this meeting":                                  35,
	"Automatically join a meeting":                              36,
	"Automatically join to this meeting when starting it":       37,
	"Be very careful. This information is sensitive and could potentially contain very private information. Only turns this settings if you absolutely need it for debugging.": 38,
	"Browse": 39,
	"By clicking Yes, this meeting will end.":       40,
	"By clicking Yes, you will leave this meeting.": 41,
	"Cancel": 26,
	"Check this option to automatically join every meeting created in the host section": 51,
	"Choose your email service to send invitation":                                      52,
	"Client binary location":                                  42,
	"Configuration settings will be lost in the next session": 43,
	"Configure master password":                               44,
	"Configured port mumble is invalid: %s":                   3,
	"Confirmation":                                            45,
	"Connecting, please wait...":                              46,
	"Continue":                                                47,
	"Copy Invitation":                                         48,
	"Copy Meeting ID":                                         49,
	"Copy URL":                                                50,
	"Debugging":                                               53,
	"Default Email":                                           54,
	"Encrypt the configuration file":                          55,
	"Error":                                                   0,
	"Finish":                                                  56,
	"Finish this meeting":                                     57,
	"Finish this meeting for all":                             58,
	"General":                                                 59,
	"Gmail":                                                   60,
	"Host a new meeting":                                      61,
	"Host meeting":                                            63,
	"Hosting":                                                 62,
	"If you backup the configuration file, we will reset the settings and continue normally. If the configuration file is ecnrypted, then we will ask you for a password to encrypt the new settings file.": 64,
	"If you disable this option, anyone could read your configuration settings":          24,
	"If you set this option to a file name, low level information will be logged there.": 65,
	"Invalid configuration file":                66,
	"Invalid password. Please, try again.":      67,
	"Invite others":                             68,
	"Join":                                      69,
	"Join Wahay Meeting":                        9,
	"Join a meeting":                            70,
	"Join meeting":                              71,
	"Join to a meeting":                         72,
	"Join to the meeting":                       73,
	"Join to this meeting":                      74,
	"Keep configuration file when Wahay closes": 75,
	"Leave":              76,
	"Leave this meeting": 77,
	"Log debug info":     78,
	"Log debug output to the seleted log file. If no file is selected then the log output will written to the default log file.": 79,
	"Master password":                80,
	"Meeting ID":                     81,
	"Meeting ID:":                    82,
	"Meeting password":               83,
	"Mumble":                         84,
	"No, cancel":                     85,
	"Now you are hosting a meeting.": 86,
	"Open":                           27,
	"Open file":                      25,
	"Outlook":                        87,
	"Password":                       88,
	"Please enter the master password for the configuration file.":          89,
	"Please join Wahay meeting with the following details:%%0D%%0A%%0D%%0A": 10,
	"Please, enter the master password for the configuration file":          90,
	"Port":                               91,
	"Port out of range":                  92,
	"Raw log file":                       93,
	"Repeat the password":                94,
	"Save changes":                       95,
	"Security":                           96,
	"Settings":                           97,
	"Show":                               98,
	"Something went wrong: %s":           1,
	"Specify a password for the meeting": 99,
	"Start Meeting":                      13,
	"Start Meeting & Join":               12,
	"Start meeting":                      100,
	"The Meeting ID cannot be blank":     15,
	"The Mumble process is down":         14,
	"The error message":                  101,
	"The invitation email has been copied to Clipboard": 8,
	"The meeting ID has been copied to Clipboard":       7,
	"The meeting ID has been copied to the clipboard":   102,
	"The meeting ID is invalid":                         18,
	"The meeting can't be closed: %s":                   4,
	"The onion service can't be deleted: %s":            6,
	"The range of valid ports are between 1 and 65535":  103,
	"This action cannot be undone":                      104,
	"Toggle password visibility":                        105,
	"Type the Meeting ID (normally an .onion address)":  106,
	"Type the password":                                 107,
	"Type the password to join the meeting":             108,
	"Type your preferred screen name":                   109,
	"Type your screen name":                             110,
	"Username":                                          111,
	"Wahay is ready to use":                             112,
	"We have detected that the configuration file is invalid or corrupted. Do you want to make a copy (backup) of it and continue?": 113,
	"We've found errors": 29,
	"Welcome":            114,
	"When this option is checked, the configuration settings will be stored in the device.": 115,
	"Yahoo Mail":                    116,
	"Yes, backup it &amp; continue": 117,
	"Yes, confirm":                  118,
	"You will not be asked for this password again until you restart Wahay.": 119,
	"enter a password of 6 chars of minimun length":                          23,
	"enter the password confirmation":                                        21,
	"passwords do not match":                                                 22,
	"please enter a valid password":                                          20,
	"the Mumble client can not be used because: %s":                          19,
	"the provided meeting ID is invalid: \n\n%s":                             17,
	"tor can't be used":                                                      28,
	"tor is not defined":                                                     5,
	"we couldn't start the meeting":                                          2,
}

var enIndex = []uint32{ // 121 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x00000022, 0x00000040,
	0x00000069, 0x0000008c, 0x0000009f, 0x000000c9,
	0x000000f5, 0x00000127, 0x0000013a, 0x0000017c,
	0x0000019a, 0x000001af, 0x000001bd, 0x000001d8,
	0x000001f7, 0x00000210, 0x0000023c, 0x00000256,
	0x00000287, 0x000002a5, 0x000002c5, 0x000002dc,
	0x0000030a, 0x00000354, 0x0000035e, 0x00000365,
	0x0000036a, 0x0000037c, 0x0000038f, 0x00000396,
	// Entry 20 - 3F
	0x000003d0, 0x000003f9, 0x00000424, 0x00000451,
	0x0000046a, 0x00000487, 0x000004bb, 0x00000564,
	0x0000056b, 0x00000593, 0x000005c1, 0x000005d8,
	0x00000610, 0x0000062a, 0x00000637, 0x00000652,
	0x0000065b, 0x0000066b, 0x0000067b, 0x00000684,
	0x000006d6, 0x00000703, 0x0000070d, 0x0000071b,
	0x0000073a, 0x00000741, 0x00000755, 0x00000771,
	0x00000779, 0x0000077f, 0x00000792, 0x0000079a,
	// Entry 40 - 5F
	0x000007a7, 0x0000086d, 0x000008c0, 0x000008db,
	0x00000900, 0x0000090e, 0x00000913, 0x00000922,
	0x0000092f, 0x00000941, 0x00000955, 0x0000096a,
	0x00000994, 0x0000099a, 0x000009ad, 0x000009bc,
	0x00000a37, 0x00000a47, 0x00000a52, 0x00000a5e,
	0x00000a6f, 0x00000a76, 0x00000a81, 0x00000aa0,
	0x00000aa8, 0x00000ab1, 0x00000aee, 0x00000b2b,
	0x00000b30, 0x00000b42, 0x00000b4f, 0x00000b63,
	// Entry 60 - 7F
	0x00000b70, 0x00000b79, 0x00000b82, 0x00000b87,
	0x00000baa, 0x00000bb8, 0x00000bca, 0x00000bfa,
	0x00000c2b, 0x00000c48, 0x00000c63, 0x00000c94,
	0x00000ca6, 0x00000ccc, 0x00000cec, 0x00000d02,
	0x00000d0b, 0x00000d21, 0x00000d9f, 0x00000da7,
	0x00000dfd, 0x00000e08, 0x00000e26, 0x00000e33,
	0x00000e7a,
} // Size: 508 bytes

const enData string = "" + // Size: 3706 bytes
	"\x02Error\x02Something went wrong: %[1]s\x02we couldn't start the meetin" +
	"g\x02Configured port mumble is invalid: %[1]s\x02The meeting can't be cl" +
	"osed: %[1]s\x02tor is not defined\x02The onion service can't be deleted:" +
	" %[1]s\x02The meeting ID has been copied to Clipboard\x02The invitation " +
	"email has been copied to Clipboard\x02Join Wahay Meeting\x02Please join " +
	"Wahay meeting with the following details:%0D%0A%0D%0A\x02%[1]s\x0aMeetin" +
	"g ID: %[2]s%0D%0A\x02Start Meeting & Join\x02Start Meeting\x02The Mumble" +
	" process is down\x02The Meeting ID cannot be blank\x02An error occurred" +
	"\x0a\x0a%[1]s\x02the provided meeting ID is invalid: \x0a\x0a%[1]s\x02Th" +
	"e meeting ID is invalid\x02the Mumble client can not be used because: %[" +
	"1]s\x02please enter a valid password\x02enter the password confirmation" +
	"\x02passwords do not match\x02enter a password of 6 chars of minimun len" +
	"gth\x02If you disable this option, anyone could read your configuration " +
	"settings\x02Open file\x02Cancel\x02Open\x02tor can't be used\x02We've fo" +
	"und errors\x02Accept\x02Allow the host to autojoin a meeting when create" +
	"s new one\x02Are you sure you want to do this action?\x02Are you sure yo" +
	"u want to end this meeting?\x02Are you sure you want to leave this meeti" +
	"ng?\x02Autojoin to this meeting\x02Automatically join a meeting\x02Autom" +
	"atically join to this meeting when starting it\x02Be very careful. This " +
	"information is sensitive and could potentially contain very private info" +
	"rmation. Only turns this settings if you absolutely need it for debuggin" +
	"g.\x02Browse\x02By clicking Yes, this meeting will end.\x02By clicking Y" +
	"es, you will leave this meeting.\x02Client binary location\x02Configurat" +
	"ion settings will be lost in the next session\x02Configure master passwo" +
	"rd\x02Confirmation\x02Connecting, please wait...\x02Continue\x02Copy Inv" +
	"itation\x02Copy Meeting ID\x02Copy URL\x02Check this option to automatic" +
	"ally join every meeting created in the host section\x02Choose your email" +
	" service to send invitation\x02Debugging\x02Default Email\x02Encrypt the" +
	" configuration file\x02Finish\x02Finish this meeting\x02Finish this meet" +
	"ing for all\x02General\x02Gmail\x02Host a new meeting\x02Hosting\x02Host" +
	" meeting\x02If you backup the configuration file, we will reset the sett" +
	"ings and continue normally. If the configuration file is ecnrypted, then" +
	" we will ask you for a password to encrypt the new settings file.\x02If " +
	"you set this option to a file name, low level information will be logged" +
	" there.\x02Invalid configuration file\x02Invalid password. Please, try a" +
	"gain.\x02Invite others\x02Join\x02Join a meeting\x02Join meeting\x02Join" +
	" to a meeting\x02Join to the meeting\x02Join to this meeting\x02Keep con" +
	"figuration file when Wahay closes\x02Leave\x02Leave this meeting\x02Log " +
	"debug info\x02Log debug output to the seleted log file. If no file is se" +
	"lected then the log output will written to the default log file.\x02Mast" +
	"er password\x02Meeting ID\x02Meeting ID:\x02Meeting password\x02Mumble" +
	"\x02No, cancel\x02Now you are hosting a meeting.\x02Outlook\x02Password" +
	"\x02Please enter the master password for the configuration file.\x02Plea" +
	"se, enter the master password for the configuration file\x02Port\x02Port" +
	" out of range\x02Raw log file\x02Repeat the password\x02Save changes\x02" +
	"Security\x02Settings\x02Show\x02Specify a password for the meeting\x02St" +
	"art meeting\x02The error message\x02The meeting ID has been copied to th" +
	"e clipboard\x02The range of valid ports are between 1 and 65535\x02This " +
	"action cannot be undone\x02Toggle password visibility\x02Type the Meetin" +
	"g ID (normally an .onion address)\x02Type the password\x02Type the passw" +
	"ord to join the meeting\x02Type your preferred screen name\x02Type your " +
	"screen name\x02Username\x02Wahay is ready to use\x02We have detected tha" +
	"t the configuration file is invalid or corrupted. Do you want to make a " +
	"copy (backup) of it and continue?\x02Welcome\x02When this option is chec" +
	"ked, the configuration settings will be stored in the device.\x02Yahoo M" +
	"ail\x02Yes, backup it &amp; continue\x02Yes, confirm\x02You will not be " +
	"asked for this password again until you restart Wahay."

var esIndex = []uint32{ // 121 elements
	// Entry 0 - 1F
	0x00000000, 0x00000006, 0x0000001d, 0x0000003d,
	0x00000073, 0x00000098, 0x000000ae, 0x000000db,
	0x0000010f, 0x00000148, 0x00000171, 0x000001cc,
	0x000001f1, 0x0000020c, 0x0000021e, 0x0000023d,
	0x00000266, 0x0000027f, 0x000002b3, 0x000002d5,
	0x00000306, 0x00000334, 0x00000363, 0x00000381,
	0x000003b5, 0x00000413, 0x00000421, 0x0000042a,
	0x00000430, 0x00000445, 0x00000461, 0x00000469,
	// Entry 20 - 3F
	0x000004c1, 0x000004f3, 0x00000529, 0x0000055b,
	0x00000583, 0x000005b8, 0x000005ed, 0x000006a1,
	0x000006aa, 0x000006db, 0x0000070b, 0x00000729,
	0x0000076b, 0x0000078a, 0x00000798, 0x000007b8,
	0x000007c2, 0x000007d5, 0x000007eb, 0x000007f6,
	0x0000085c, 0x000008a1, 0x000008ad, 0x000008c2,
	0x000008e6, 0x000008ef, 0x00000905, 0x00000926,
	0x0000092e, 0x00000934, 0x00000951, 0x00000959,
	// Entry 40 - 5F
	0x0000096d, 0x00000a69, 0x00000ad0, 0x00000af4,
	0x00000b1f, 0x00000b2f, 0x00000b36, 0x00000b4c,
	0x00000b61, 0x00000b77, 0x00000b8c, 0x00000ba3,
	0x00000bd9, 0x00000bdf, 0x00000bf6, 0x00000c1c,
	0x00000cda, 0x00000cee, 0x00000d00, 0x00000d13,
	0x00000d2e, 0x00000d35, 0x00000d42, 0x00000d69,
	0x00000d71, 0x00000d7d, 0x00000dbf, 0x00000e0b,
	0x00000e12, 0x00000e28, 0x00000e3d, 0x00000e53,
	// Entry 60 - 7F
	0x00000e63, 0x00000e6d, 0x00000e7d, 0x00000e85,
	0x00000eb2, 0x00000ec6, 0x00000eda, 0x00000f11,
	0x00000f44, 0x00000f66, 0x00000f8a, 0x00000fcb,
	0x00000fe2, 0x00001013, 0x0000103a, 0x00001057,
	0x00001069, 0x00001087, 0x00001109, 0x00001114,
	0x00001169, 0x00001179, 0x00001196, 0x000011a4,
	0x000011ed,
} // Size: 508 bytes

const esData string = "" + // Size: 4589 bytes
	"\x02Error\x02Algo salió mal: %[1]s\x02no se pudo comenzar la reunión\x02" +
	"El puerto configurado para Mumble es inválido: %[1]s\x02La reunión no se" +
	" pudo cerrar: %[1]s\x02tor no está definido\x02El servicio Onion no se p" +
	"udo eliminar: %[1]s\x02El ID de la reunion ha sido copiado al portapapel" +
	"es\x02El correo de invitación ha sido copiado al portapapeles\x02Unirse " +
	"a una reunión a travéz de Wahay\x02Por favor únete a la reunión a travéz" +
	" de Wahay con los siguientes detalles:%0D%0A%0D%0A\x02%[1]s\x0aID de la " +
	"reunión: %[2]s%0D%0A\x02Comenzar reunión y Unirse\x02Comenzar reunión" +
	"\x02El proceso Mumble está caído\x02El ID de la reunión no puede ser vac" +
	"ío\x02Ocurrió un error\x0a\x0a%[1]s\x02el ID de la reunión provisto es " +
	"inválido: \x0a\x0a%[1]s\x02El ID de la reunión es inválido\x02El cliente" +
	" Mumble no se puede usar porque: %[1]s\x02por favor especifique una cont" +
	"raseña válida\x02especifique la confirmación de la contraseña\x02las con" +
	"traseñas no coinciden\x02especifique una contraseña de mínimo 6 caracter" +
	"es\x02Si deshabilita esta opción, cualquier persona podría leer los pará" +
	"metros de configuración\x02Abrir archivo\x02Cancelar\x02Abrir\x02tor no " +
	"se puede usar\x02Encontramos algunos errores\x02Aceptar\x02Permita que e" +
	"l organizador se una automáticamente a una reunión cuando cree una nueva" +
	"\x02¿Está seguro de que quieres hacer esta acción?\x02¿Está seguro de qu" +
	"e quieres terminar esta reunión?\x02¿Está seguro de que quiere dejar est" +
	"a reunión?\x02Unirse automáticamente a esta reunión\x02Únase automáticam" +
	"ente a esta reunión al iniciarla\x02Únase automáticamente a esta reunión" +
	" al iniciarla\x02Ten mucho cuidado. Esta información es confidencial y p" +
	"odría contener información muy privada. Solo cambia esta configuración s" +
	"i la necesita absolutamente para la depuración.\x02Examinar\x02Al hacer " +
	"clic en Sí, esta reunión finalizará.\x02Al hacer clic en Sí, saldrá de e" +
	"sta reunión.\x02Ubicación del binario Mumble\x02Los ajustes de configura" +
	"ción se perderán en la próxima sesión\x02Configurar contraseña maestra" +
	"\x02Confirmación\x02Conectando, espere por favor...\x02Continuar\x02Copi" +
	"ar invitación\x02Copiar ID de reunión\x02Copiar URL\x02Marque esta opció" +
	"n para unirse automáticamente a cada reunión creada en la sección del an" +
	"fitrión\x02Elija su servicio de correo electrónico para enviar la invita" +
	"ción.\x02Depuración\x02Email predeterminado\x02Cifrar el archivo de conf" +
	"iguración\x02Terminar\x02Termina esta reunión\x02Termina esta reunión pa" +
	"ra todos\x02General\x02Gmail\x02Organizar una nueva reunión\x02Hosting" +
	"\x02Reunión de acogida\x02Si realiza una copia de seguridad del archivo " +
	"de configuración, restableceremos la configuración y continuaremos norma" +
	"lmente. Si el archivo de configuración está cifrado, le pediremos una co" +
	"ntraseña para cifrar el nuevo archivo de configuración.\x02Si establece " +
	"esta opción en un nombre de archivo, la información de bajo nivel se reg" +
	"istrará allí.\x02Archivo de configuración inválido\x02Contraseña invalid" +
	"a. Inténtalo de nuevo.\x02Invitar a otros\x02Unirse\x02Unirse a una reun" +
	"ión\x02Unirse a la reunión\x02Únete a una reunión\x02Únete a la reunión" +
	"\x02Únete a esta reunión\x02Mantener el archivo de configuración al cerr" +
	"ar Wahay\x02Salir\x02Salir de esta reunión\x02Registrar información de d" +
	"epuración\x02Registre la salida de depuración en el archivo de registro " +
	"seleccionado. Si no se selecciona ningún archivo, la salida del registro" +
	" se escribirá en el archivo de registro predeterminado.\x02Contraseña ma" +
	"estra\x02ID de la reunión\x02ID de la reunión:\x02Contraseña de la reuni" +
	"ón\x02Mumble\x02No, cancelar\x02Ahora estás organizando una reunión." +
	"\x02Outlook\x02Contraseña\x02Ingrese la contraseña maestra para el archi" +
	"vo de configuración.\x02Por favor, ingrese la contraseña maestra para el" +
	" archivo de configuración\x02Puerto\x02Puerto fuera de rango\x02Archivo " +
	"de registros\x02Repita la contraseña\x02Guardar cambios\x02Seguridad\x02" +
	"Configuraciones\x02Mostrar\x02Especifique una contraseña para la reunión" +
	"\x02Comience a reunirse\x02El mensaje de error\x02La ID de la reunión se" +
	" ha copiado en el portapapeles.\x02El rango de puertos válidos está entr" +
	"e 1 y 65535\x02Esta acción no se puede deshacer\x02Alternar visibilidad " +
	"de contraseña\x02Escriba la ID de la reunión (normalmente una dirección " +
	".onion)\x02Escribe la contraseña\x02Escriba la contraseña para unirse a " +
	"la reunión\x02Escriba su nombre de usuario preferido\x02Escriba su nombr" +
	"e de usuario\x02Nombre de usuario\x02Wahay está listo para usarse\x02Hem" +
	"os detectado que el archivo de configuración no es válido o está dañado." +
	" ¿Desea hacer una copia de seguridad y continuar?\x02Bienvenido\x02Cuand" +
	"o esta opción está marcada, la configuración se guardará en el dispositi" +
	"vo.\x02Correo de Yahoo\x02Sí, respaldarlo y continuar\x02Si, confirmar" +
	"\x02No se le volverá a solicitar esta contraseña hasta que reinicie Waha" +
	"y."

	// Total table size 9311 bytes (9KiB); checksum: 34E17CB1
