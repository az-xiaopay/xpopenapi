package xpnotify

import (
	"encoding/json"
	"runtime/debug"
)

type NotifyHandler interface {
	ConsumeSuccess(n *NotifyConsumeRes) error
	ReductionFail(n *NotifyConsumeRes) error
	Reacquire(n *NotifyReacquire) error
	DataChange(n *NotifyDataChange, notify *Notify) error
}

func (notify *Notify) Do(handler NotifyHandler) Resp {
	switch notify.NotifyType {
	case NotifyType_ConsumeSuccess:
		if err := notify.consumeSuccess(handler); err != nil {
			return Resp{"FAIL", err.Error()}
		}
	case NotifyType_ReductionFail:
		if err := notify.reductionFail(handler); err != nil {
			return Resp{"FAIL", err.Error()}
		}
	case NotifyType_Reacquire:
		if err := notify.reacquire(handler); err != nil {
			return Resp{"FAIL", err.Error()}
		}
	case NotifyType_Datachange:
		if err := notify.dataChange(handler); err != nil {
			return Resp{"FAIL", err.Error()}
		}
	}

	return Resp{"SUCCESS", ""}
}

func (notify *Notify) consumeSuccess(handler NotifyHandler) error {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	marshal, err := json.Marshal(notify.NotifyContent)
	if err != nil {
		return err
	}
	var info NotifyConsumeRes
	err = json.Unmarshal(marshal, &info)
	if err != nil {
		return err
	}

	return handler.ConsumeSuccess(&info)
}

func (notify *Notify) reductionFail(handler NotifyHandler) error {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	marshal, err := json.Marshal(notify.NotifyContent)
	if err != nil {
		return err
	}
	var info NotifyConsumeRes
	err = json.Unmarshal(marshal, &info)
	if err != nil {
		return err
	}

	return handler.ReductionFail(&info)
}

func (notify *Notify) reacquire(handler NotifyHandler) (e error) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	marshal, err := json.Marshal(notify.NotifyContent)
	if err != nil {
		return err
	}
	var info NotifyReacquire
	err = json.Unmarshal(marshal, &info)
	if err != nil {
		return err
	}
	return handler.Reacquire(&info)
}

func (notify *Notify) dataChange(handler NotifyHandler) (e error) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()
	marshal, err := json.Marshal(notify.NotifyContent)
	if err != nil {
		return err
	}
	var info NotifyDataChange
	err = json.Unmarshal(marshal, &info)
	if err != nil {
		return err
	}
	return handler.DataChange(&info, notify)
}
