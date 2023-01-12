// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package actordiff

import (
	"fmt"
	"io"
	"math"
	"sort"

	core "github.com/filecoin-project/lily/pkg/core"
	types "github.com/filecoin-project/lotus/chain/types"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *ActorChange) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.Actor (types.Actor) (struct)
	if len("actor") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"actor\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("actor"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("actor")); err != nil {
		return err
	}

	if err := t.Actor.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Current ([]uint8) (slice)
	if len("current_state") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"current_state\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("current_state"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("current_state")); err != nil {
		return err
	}

	if len(t.Current) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Current was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Current))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Current[:]); err != nil {
		return err
	}

	// t.Previous ([]uint8) (slice)
	if len("previous_state") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"previous_state\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("previous_state"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("previous_state")); err != nil {
		return err
	}

	if len(t.Previous) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Previous was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Previous))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Previous[:]); err != nil {
		return err
	}

	// t.Change (core.ChangeType) (uint8)
	if len("change") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"change\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("change"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("change")); err != nil {
		return err
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Change)); err != nil {
		return err
	}
	return nil
}

func (t *ActorChange) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ActorChange{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("ActorChange: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Actor (types.Actor) (struct)
		case "actor":

			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					t.Actor = new(types.Actor)
					if err := t.Actor.UnmarshalCBOR(cr); err != nil {
						return xerrors.Errorf("unmarshaling t.Actor pointer: %w", err)
					}
				}

			}
			// t.Current ([]uint8) (slice)
		case "current_state":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Current: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Current = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Current[:]); err != nil {
				return err
			}
			// t.Previous ([]uint8) (slice)
		case "previous_state":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Previous: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Previous = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Previous[:]); err != nil {
				return err
			}
			// t.Change (core.ChangeType) (uint8)
		case "change":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}
			if maj != cbg.MajUnsignedInt {
				return fmt.Errorf("wrong type for uint8 field")
			}
			if extra > math.MaxUint8 {
				return fmt.Errorf("integer in input was too large for uint8 field")
			}
			t.Change = core.ChangeType(extra)

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *StateChange) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{161}); err != nil {
		return err
	}

	// t.ActorState (cid.Cid) (struct)
	if len("actors") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"actors\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("actors"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("actors")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.ActorState); err != nil {
		return xerrors.Errorf("failed to write cid field t.ActorState: %w", err)
	}

	return nil
}

func (t *StateChange) UnmarshalCBOR(r io.Reader) (err error) {
	*t = StateChange{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("StateChange: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadString(cr)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.ActorState (cid.Cid) (struct)
		case "actors":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.ActorState: %w", err)
				}

				t.ActorState = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
