package engine

import "fmt"

type Scope struct {
	Variables   map[string]*ScopeVariable
	ParentScope *Scope
}

type ScopeVariable struct {
	Value any
}

func NewScopeVariable(v any) *ScopeVariable {
	return &ScopeVariable{
		Value: v,
	}
}

func NewScope(parentScope *Scope) *Scope {
	return &Scope{
		Variables:   make(map[string]*ScopeVariable),
		ParentScope: parentScope,
	}
}

func (s *Scope) GetVariable(name string) (any, error) {
	value, ok := s.Variables[name]
	if ok {
		return value.Value, nil
	}

	if s.ParentScope != nil {
		return s.ParentScope.GetVariable(name)
	}

	return nil, fmt.Errorf("(Get) Variable '%s' not found", name)
}

func (s *Scope) SetVariable(name string, value any) error {
	_, ok := s.Variables[name]
	if ok {
		s.Variables[name].Value = value

		return nil
	}

	if s.ParentScope != nil {
		return s.ParentScope.SetVariable(name, value)
	}

	return fmt.Errorf("(Set) Variable '%s' not found", name)
}

func (s *Scope) DefineVariable(name string, value any) error {
	if _, ok := s.Variables[name]; ok {
		return fmt.Errorf("(Define) Variable '%s' already defined", name)
	}

	s.Variables[name] = NewScopeVariable(value)
	return nil
}

func (s *Scope) UpdateVariable(name string, value any) error {
	if _, ok := s.Variables[name]; !ok {
		if s.ParentScope != nil {
			return s.ParentScope.UpdateVariable(name, value)
		} else {
			return fmt.Errorf("(Update) Variable '%s' not found", name)
		}
	}

	s.Variables[name].Value = value
	return nil
}

func (s *Scope) DeleteVariable(name string) error {
	if _, ok := s.Variables[name]; !ok {
		if s.ParentScope != nil {
			return s.ParentScope.DeleteVariable(name)
		} else {
			return fmt.Errorf("(Delete) Variable '%s' not found", name)
		}
	}

	delete(s.Variables, name)
	return nil
}

func (s *Scope) GetRawVariable(name string) (*ScopeVariable, error) {
	value, ok := s.Variables[name]
	if ok {
		return value, nil
	}

	if s.ParentScope != nil {
		return s.ParentScope.GetRawVariable(name)
	}

	return nil, fmt.Errorf("(GetRaw) Variable '%s' not found", name)
}

func (s *Scope) GlobalVariable(name string) error {
	if s.ParentScope == nil {
		return nil
	}

	raw, e := s.GetRawVariable(name)

	if e != nil {
		return e
	}

	s.Variables[name] = raw
	s.ParentScope.Variables[name] = raw

	return nil
}
