// Code generated by ent, DO NOT EDIT.

package ent

import "github.com/google/uuid"

func (a *Attachment) GetID() uuid.UUID {
	return a.ID
}

func (at *AuthTokens) GetID() uuid.UUID {
	return at.ID
}

func (d *Document) GetID() uuid.UUID {
	return d.ID
}

func (dt *DocumentToken) GetID() uuid.UUID {
	return dt.ID
}

func (gr *Group) GetID() uuid.UUID {
	return gr.ID
}

func (i *Item) GetID() uuid.UUID {
	return i.ID
}

func (_if *ItemField) GetID() uuid.UUID {
	return _if.ID
}

func (l *Label) GetID() uuid.UUID {
	return l.ID
}

func (l *Location) GetID() uuid.UUID {
	return l.ID
}

func (u *User) GetID() uuid.UUID {
	return u.ID
}
