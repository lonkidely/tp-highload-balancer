// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package internal

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB8923099DecodeBalancerInternal(in *jlexer.Lexer, out *ErrResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "error":
			out.ErrMassage = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8923099EncodeBalancerInternal(out *jwriter.Writer, in ErrResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix[1:])
		out.String(string(in.ErrMassage))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8923099EncodeBalancerInternal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8923099EncodeBalancerInternal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8923099DecodeBalancerInternal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8923099DecodeBalancerInternal(l, v)
}
func easyjsonB8923099DecodeBalancerInternal1(in *jlexer.Lexer, out *EchoResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "body":
			out.Body = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB8923099EncodeBalancerInternal1(out *jwriter.Writer, in EchoResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix[1:])
		out.String(string(in.Body))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v EchoResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB8923099EncodeBalancerInternal1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EchoResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB8923099EncodeBalancerInternal1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EchoResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB8923099DecodeBalancerInternal1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EchoResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB8923099DecodeBalancerInternal1(l, v)
}
