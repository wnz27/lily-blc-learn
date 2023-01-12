// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package v0

import (
	"fmt"
	"io"
	"math"
	"sort"

	core "github.com/filecoin-project/lily/pkg/core"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *AddressChange) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.Address ([]uint8) (slice)
	if len("address") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"address\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("address"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("address")); err != nil {
		return err
	}

	if len(t.Address) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Address was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Address))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Address[:]); err != nil {
		return err
	}

	// t.Current (typegen.Deferred) (struct)
	if len("current_actorID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"current_actorID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("current_actorID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("current_actorID")); err != nil {
		return err
	}

	if err := t.Current.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Previous (typegen.Deferred) (struct)
	if len("previous_actorID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"previous_actorID\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("previous_actorID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("previous_actorID")); err != nil {
		return err
	}

	if err := t.Previous.MarshalCBOR(cw); err != nil {
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

func (t *AddressChange) UnmarshalCBOR(r io.Reader) (err error) {
	*t = AddressChange{}

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
		return fmt.Errorf("AddressChange: map struct too large (%d)", extra)
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
		// t.Address ([]uint8) (slice)
		case "address":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Address: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Address = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Address[:]); err != nil {
				return err
			}
			// t.Current (typegen.Deferred) (struct)
		case "current_actorID":

			{

				t.Current = new(cbg.Deferred)

				if err := t.Current.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}
			// t.Previous (typegen.Deferred) (struct)
		case "previous_actorID":

			{

				t.Previous = new(cbg.Deferred)

				if err := t.Previous.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
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

	// t.Addresses (cid.Cid) (struct)
	if len("addresses") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"addresses\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("addresses"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("addresses")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.Addresses); err != nil {
		return xerrors.Errorf("failed to write cid field t.Addresses: %w", err)
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
		// t.Addresses (cid.Cid) (struct)
		case "addresses":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Addresses: %w", err)
				}

				t.Addresses = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
