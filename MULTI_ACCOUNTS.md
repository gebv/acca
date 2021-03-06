# MULTI ACCOUNTS

Вводится понятие пользователь (как владельце, для группировки счетов). В рамках пользователя есть 1 и более счетов. Все счета могут использоаться для рассчета.
Информация о пользователе кодируется в `key` счета (см. таблицу `acca.accounts`).
Например `ma.<ownerID>.<type>.some.fields`, где
- `ma` - обязательный префикс для всех случаев с мульти счетами
- `<ownerID>` - идентификатор
- `<type>` - тип
- `some.fields` - прочие значения для другой логики

## Примеры органзиации кредитного и основного счета

- кредитный (C credit) - сколько пользователь должен системе - система предоставила кредитную линию
- основной (M main) - средства доступные для расхода пользователю

Начилсение кредита сопровождается отражением суммы кредита на счете C и увеличения основного счета равного на сумму предоставленного кредита
Например
Пополнение счета на 100
M=100
Предоставили кредит в размере 50
C=50
Увеличили на сумму кредита основной счет M
M=150
В итоге
Собственных средств у пользователя = M - C = 100
Доступно для расхода у пользователя = M

## Пример организации бонусного и основного счета

- кредитный
- основной
- бонусный (B bonus) - аффилиатные средства

Например
Пополнение счета на 100
M=100
Предоставление бонуса в размере 50
B=50
Увеличиваем на сумму бонуса основного счета M
M=150
В итоге
Собственных средств у пользователя = M - B = 100
Доступных для расхода у пользователя = M

Причем средства на бонусном счете заключается в том что это не обязательства перед системой.

---

Ваша логика и реализация может быть иной.

