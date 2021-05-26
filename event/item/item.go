package item

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/tradersclub/PocArquitetura/app"
	nmodel "github.com/tradersclub/TCNatsModel/TCStartKit"
	"github.com/tradersclub/TCUtils/logger"
	"github.com/tradersclub/TCUtils/tcerr"
)

// Register group health check
func Register(apps *app.Container, conn *nats.Conn) {
	e := &event{
		apps: apps,
		nc:   conn,
	}

	e.nc.Subscribe(nmodel.TCSTARTKIT_GET_ITEM_BY_ID, e.getItemById)
}

type event struct {
	apps *app.Container
	nc   *nats.Conn
}

func (e *event) getItemById(msg *nats.Msg) {
	ctx := context.Background()

	//recupera o objeto através do metódo criado no TCNatsModel
	obj, err := nmodel.GetItemByIdFromBytes(msg.Data)

	// valida erro de parser
	if err != nil {
		logger.ErrorContext(ctx, "event.session.get_session_by_id", err.Error())
		obj := new(nmodel.GetItemById)
		obj.Err = tcerr.NewError(500, "event.session.get_session_by_id", err.Error())
		msg.Respond(obj.ToBytes())
		return
	}

	// pega o item
	session, err := e.apps.Item.GetItemById(ctx, obj.Id)
	if err != nil {
		logger.ErrorContext(ctx, "event.session.get_session_by_id", err.Error())
		obj.Err = tcerr.NewError(500, "event.session.get_session_by_id", err.Error())
	} else {
		obj.Data = session
	}

	msg.Respond(obj.ToBytes())
}
