package subscribers

import (
	"github.com/omid9h/echo-boilerplate/pkg/hook"
)

type UserSubscriber struct {
}

func (s UserSubscriber) Created(payload hook.EventPayload) {

}

func (s UserSubscriber) Deleted(payload hook.EventPayload) {

}

func (s UserSubscriber) Updated(payload hook.EventPayload) {

}
