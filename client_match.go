package dota2

import (
	"github.com/Philipp15b/go-steam/v3/protocol/gamecoordinator"

	"github.com/paralin/go-dota2/events"
)

// handleMatchSignedOut handles an incoming steam datagram ticket.
func (d *Dota2) handleMatchSignedOut(packet *gamecoordinator.GCPacket) error {
	ev := &events.MatchSignedOut{}
	if err := d.unmarshalBody(packet, &ev.CMsgGCToClientMatchSignedOut); err != nil {
		return err
	}

	d.emit(ev)
	return nil
}
