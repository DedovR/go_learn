# RabbitMQ

- Publisher публикует сообщение в Exchange
- Exchange пересылает сообщение в Queue или другой Exchange(Binding), в зависимости от адреса получателя. Если не может переслать то удаляет
- Queue хранит сообщения
- На очередь могут быть подписаны один либо несколько Consumer'ов. Если консюмер отправляет очереди подтверждение ack(Acknowledge), сообщение из очереди удаляется

При попадании сообщения в Exchange происходит роутинг(маршрутизация) по routing_key. Типы маршрутизации:
- Direct - отправляет в конкретную очередь, соответствующую routing_key
- Fanout - перешлет во все очереди подключенные к Exchange
- Topic
- Consistent Hashing
- Можно добавить свои через плагины

### Гарантии доставки
- At least once - хотя бы один раз, гарантированно будет доставлено но возможно несколько раз
- At most once - максимум один раз, но при этом может быть не доставлено совсем
- Exactly-once – ровно один раз. Гарантия, которая исключает дубли и потерю данных, но понижает скорость работы взаимодействий.

Гарантии достигаются путем настрой брокеров и очередей

### Нажедность
- Если консюмер не доступен или не может обработать сообщения, есть возможность отправить сообщение в Dead Letter Queue
- Чтобы очередь переживала рестарты, можно использовать флаг durable
- У сообщений в очереди есть возможность указать ttl
- Если хотим чтобы очередь читал гарантированно один косюмер, можно использовать флаг exclusive
- Очередь работает по принципу FIFO
- Если паблишер хочет быть увереным, что сообщение отправилось в какую-то очередь, можно использовать дефолтный ексчейнж пустую строку

Ребит сам отправляет сообщение получателю, то есть всегда открыто соединение. За счет этого:
- Низкий летенси
- Может распределять нагрузку
- Может задедосить консюмер(есть найтройки QoS - число сообщений в канале)
- Широкие возможности для роутинга
- Порядок сообщений не гарантирован, если консюмеров больше одного

# Kafka

Система, реализующая распределенный реплецируемый лог сообщений

- Распределенный лог(топик) не целиком хранится на одной машине, а разбит на несколько частей - партиций, которые хранятся на разных машинах
- Реплецируемый - логи хранятся в нескольких копиях на случай отказов
- Лог - упорядоченная последовательность сообщений

Сообщения хранятся в топиках, могут быть разного формата. У сообщений нет схемы, но можно её добавить через внешние инструменты

Схема работы:
- Продюсер пишет сообщение в конец партиции
- Консюмер читает сообщение по оффсету. Может вычитывать сообщения с начала партиции
- К партиции может быть подключено несколько консюмеров, каждый со своим read offset

В кафке у сообщения может быть ключ(key), на его основе работает алгоритм партиционирования, по факту это деление по модулю от полученного хеша на основе ключа.

Consumer group - логический консюмер, для поддержки многопоточности. Обеспечивает получение сообщения только одним консюмером из группы
- Group coordinator - один из брокеров кафки назначает себя координаторв группы и отвечает за состав и живучесть группы
- Group leader - один из консюмеров группы рандомно назначается Group coordinator'ом, и потом Group leader распределяет консумеров по партициям
- К одной партиции должен быть подключен только один конкурирующий консюмер. Один консюмер может читать из нескольких партиций

Порядок сообщений сохраняется в рамках одной партиции

Offset - метка сообщения, которое консюмер прочитал. Консюмер может сделать commit offset, тогда на стороне кафки будет хранится смещение консюмера. Сommit offset хранится в топике __consumer_offset, который хранится у координатора группы и реплицируется на остальные брокеры

При добавлении нового консюмера в группу происходит перебалансировка, во время которой ни один консюмер не читает сообщения из топика.

У кафки пул-модель, то есть консюмер ходит за сообщениями. Может отдавать сообщения пачками

### Гарантии доставки

- At least once - хотя бы один раз, гарантированно будет доставлено но возможно несколько раз
- At most once - максимум один раз, но при этом может быть не доставлено совсем
- Exactly-once – ровно один раз. Гарантия, которая исключает дубли и потерю данных, но понижает скорость работы взаимодействий.

В кафке гарантии достигаются тем что есть последовательный лог сообщений сохраненный на диске, который хранится определенное время.

### Репликация

Каждая партиция может иметь несколько реплик. Каждый брокер может хранить до тысячи реплик разных топиков. Регулируется фактором репликации: если не нужна то =1, если нужна то 3 и более

Гарантии в кафке при репликации:
- Если продюсер отправил сообщение В после А, то офсет у В будет больше чем у А
- Если продюсер отправил сообщение и все ин-синк реплики подтвердили, что сохранили его. Можно считать, что сообщение закоммиченым
- Закоммиченные сообщения не будут потеряны пока хоть одна реплика жива
- Консюмеры могут читать только закомиченные сообщения

Сообщение можно отправить параметром acks:
- acks = 0 если произошла отправка по сети то оно считается отправленым
- acks = 1 считается успешной если лидер записал сообщение на диск
- acks = all считается успешным если лидер и все ин-синк реплики закомитили сообщения
