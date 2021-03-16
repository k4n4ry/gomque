package que

import "github.com/streadway/amqp"

type Enq struct {
	address string
	queue   *amqp.Queue
	conn    *amqp.Connection
	ch      *amqp.Channel
}

// NewEnq is constructor
func NewEnq(address, qName string) (*Enq, error) {
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		qName,
		true,
		false,
		false,
		false,
		nil,
	)
	return &Enq{
		address: address,
		queue:   &q,
		conn:    conn,
		ch:      ch,
	}, nil
}

// Enqueue is Publish message
func (e *Enq) Enqueue(msg []byte) error {
	return e.ch.Publish(
		"",           // exchange
		e.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
}

// Close is close channel and queue
func (e *Enq) Close() {
	e.ch.Close()
	e.conn.Close()
}

// type Deq struct {
// 	address string
// 	queue   *amqp.Queue
// 	conn    *amqp.Connection
// 	ch      *amqp.Channel
// }

// // NewEnq is constructor
// func NewDeq(address, qName string) (*Deq, error) {
// 	conn, err := amqp.Dial(address)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return nil, err
// 	}
// 	q, err := ch.QueueDeclare(
// 		qName,
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	return &Deq{
// 		address: address,
// 		queue:   &q,
// 		conn:    conn,
// 		ch:      ch,
// 	}, nil
// }

// func (d *Deq) Dequeue() error {
// 	msgs, err := d.ch.Consume(
// 		d.queue.Name, // queue
// 		"",           // consumer
// 		false,        // auto-ack
// 		false,        // exclusive
// 		false,        // no-local
// 		false,        // no-wait
// 		nil,          // args
// 	)
// }

// func (d *Deq) Close() {
// 	d.ch.Close()
// 	d.conn.Close()
// }
