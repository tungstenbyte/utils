package natstan

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/tungstenbyte/utils/json"
)

type CbHandler func(msg *MsgHandler)

type MsgHandler struct {
	msg *nats.Msg
}

func (m *MsgHandler) Decode(res interface{}) error {
	return json.FromJSON(m.msg.Data, res)
}

func (m *MsgHandler) Respond(req interface{}) error {
	data := json.ToJSON(req)

	return m.msg.Respond(data)
}

type QueueMessager interface {
	Publish(ctx context.Context, subj string, req interface{}) error
	Subscribe(subj string, cb CbHandler) error
	QueueSubscribe(subj string, queue string, cb CbHandler) error
	Close()
}

var ErrResIsNotAPointer = errors.New("field res is not a pointer")

func NewQueueMessagerWithNatsConn(conn *nats.Conn) QueueMessager {
	return &queueMessagerImpl{conn}
}

type queueMessagerImpl struct {
	conn *nats.Conn
}

func (qm *queueMessagerImpl) Publish(ctx context.Context, subj string, req interface{}) error {
	data := json.ToJSON(req)

	return qm.conn.Publish(subj, data)
}

func (qm *queueMessagerImpl) Subscribe(subj string, cb CbHandler) error {
	_, err := qm.conn.Subscribe(subj, func(msg *nats.Msg) {
		cb(&MsgHandler{msg})
	})

	return err
}

func (qm *queueMessagerImpl) QueueSubscribe(subj string, queue string, cb CbHandler) error {
	_, err := qm.conn.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
		cb(&MsgHandler{msg})
	})

	return err
}

func (qm *queueMessagerImpl) Close() {
	qm.conn.Close()
}
