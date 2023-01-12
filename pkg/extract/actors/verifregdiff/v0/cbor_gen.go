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

func (t *StateChange) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{162}); err != nil {
		return err
	}

	// t.Verifiers (cid.Cid) (struct)
	if len("verifiers") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"verifiers\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("verifiers"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("verifiers")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.Verifiers); err != nil {
		return xerrors.Errorf("failed to write cid field t.Verifiers: %w", err)
	}

	// t.Clients (cid.Cid) (struct)
	if len("clients") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"clients\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("clients"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("clients")); err != nil {
		return err
	}

	if err := cbg.WriteCid(cw, t.Clients); err != nil {
		return xerrors.Errorf("failed to write cid field t.Clients: %w", err)
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
		// t.Verifiers (cid.Cid) (struct)
		case "verifiers":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Verifiers: %w", err)
				}

				t.Verifiers = c

			}
			// t.Clients (cid.Cid) (struct)
		case "clients":

			{

				c, err := cbg.ReadCid(cr)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Clients: %w", err)
				}

				t.Clients = c

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *ClientsChange) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.Client ([]uint8) (slice)
	if len("client") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"client\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("client"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("client")); err != nil {
		return err
	}

	if len(t.Client) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Client was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Client))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Client[:]); err != nil {
		return err
	}

	// t.Current (typegen.Deferred) (struct)
	if len("current") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"current\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("current"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("current")); err != nil {
		return err
	}

	if err := t.Current.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Previous (typegen.Deferred) (struct)
	if len("previous") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"previous\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("previous"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("previous")); err != nil {
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

func (t *ClientsChange) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ClientsChange{}

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
		return fmt.Errorf("ClientsChange: map struct too large (%d)", extra)
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
		// t.Client ([]uint8) (slice)
		case "client":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Client: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Client = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Client[:]); err != nil {
				return err
			}
			// t.Current (typegen.Deferred) (struct)
		case "current":

			{

				t.Current = new(cbg.Deferred)

				if err := t.Current.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}
			// t.Previous (typegen.Deferred) (struct)
		case "previous":

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
func (t *VerifiersChange) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write([]byte{164}); err != nil {
		return err
	}

	// t.Verifier ([]uint8) (slice)
	if len("verifier") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"verifier\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("verifier"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("verifier")); err != nil {
		return err
	}

	if len(t.Verifier) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Verifier was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Verifier))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Verifier[:]); err != nil {
		return err
	}

	// t.Current (typegen.Deferred) (struct)
	if len("current") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"current\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("current"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("current")); err != nil {
		return err
	}

	if err := t.Current.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Previous (typegen.Deferred) (struct)
	if len("previous") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"previous\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("previous"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("previous")); err != nil {
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

func (t *VerifiersChange) UnmarshalCBOR(r io.Reader) (err error) {
	*t = VerifiersChange{}

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
		return fmt.Errorf("VerifiersChange: map struct too large (%d)", extra)
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
		// t.Verifier ([]uint8) (slice)
		case "verifier":

			maj, extra, err = cr.ReadHeader()
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Verifier: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}

			if extra > 0 {
				t.Verifier = make([]uint8, extra)
			}

			if _, err := io.ReadFull(cr, t.Verifier[:]); err != nil {
				return err
			}
			// t.Current (typegen.Deferred) (struct)
		case "current":

			{

				t.Current = new(cbg.Deferred)

				if err := t.Current.UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}
			// t.Previous (typegen.Deferred) (struct)
		case "previous":

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
