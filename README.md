Сценарии

I Главный экран

/api/tournaments - получили список турниров (Муж/Жен)
/api/seasons - получиле список сезонов («10/11», «20/21», «21/22»)
есть признак isCurrent - берем текущий сезон

/api/seasons/21-22/stages - получили список этапов соревнований:
квалификация
первый круг
стыковые игры
второй круг
выбираем этап с признаком isCurrent
пока что будет 

/api/leagues/male/21-22/21-22_second_group_round
/api/leagues/female/21-22/21-22_second_group_round
 - получаем список лиг

выводим туринры/сезон/этап/лиги


II Клик на лигу
при клике на лигу запрашиваем список групп:

/api/groups/male/21-22/21-22_second_group_round/super_league
/api/groups/male/21-22/21-22_second_group_round/high_league
/api/groups/male/21-22/21-22_second_group_round/first_league
……

/api/groups/female/21-22/21-22_second_group_round/super_league
/api/groups/female/21-22/21-22_second_group_round/high_league
/api/groups/female/21-22/21-22_second_group_round/first_league
……

получили список групп 


III Переход в группу

Запрашиваем список команд

/api/groups/21-22_second_group_round_first_league_A/teams

и список игр
/api/groups/21-22_second_group_round_first_league_A/matches

на основании это информации можно и составить турнирные таблицы и составить таблицу результатов и показать расписание


IV Клик на команду

Экран команды
/api/teams/1234 - инфа по команде

/api/teams/1234/groups - список групп, в которых участвована команда

V Состав команды

/api/teams/1234/groups/21-22_second_group_round_first_league_A/players - список игроков участвовавших в конкретном этапе соревнований


VI Экран результата игры

/api/matches/4321/result - получение результатов игры
