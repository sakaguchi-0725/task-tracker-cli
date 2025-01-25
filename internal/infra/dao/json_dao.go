package dao

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sakaguchi-0725/task-tracker/internal/util"
)

type JsonDAO[T any] interface {
	Create(val *T) error
	Update(val *T) error
	Find(result *[]T) error
	First(result *T) error
	Delete() error
	Where(field string, val any) JsonDAO[T]
}

type jsonDAO[T any] struct {
	filePath   string
	conditions []condition
}

func (j *jsonDAO[T]) Create(val *T) error {
	records, err := util.ReadJSON[T](j.filePath)
	if err != nil {
		return err
	}

	records = append(records, *val)
	return util.WriteJSON[T](j.filePath, records)
}

func (j *jsonDAO[T]) Delete() error {
	defer func() {
		j.conditions = []condition{}
	}()

	if len(j.conditions) == 0 {
		return errors.New("no conditions specified")
	}

	records, err := util.ReadJSON[T](j.filePath)
	if err != nil {
		return err
	}

	var newRecords []T
	deleted := false

	for _, record := range records {
		matched := true

		for _, c := range j.conditions {
			fieldName := c.field
			expectedVal := c.value

			val := reflect.ValueOf(record)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			fieldVal := val.FieldByName(fieldName)
			if !fieldVal.IsValid() {
				return fmt.Errorf("field %s does not exist in type %T", fieldName, record)
			}

			if !reflect.DeepEqual(fieldVal.Interface(), expectedVal) {
				matched = false
				break
			}
		}

		if matched {
			deleted = true
			continue
		}

		newRecords = append(newRecords, record)
	}

	if !deleted {
		return errors.New("record not found")
	}

	if err := util.WriteJSON[T](j.filePath, newRecords); err != nil {
		return fmt.Errorf("failed to write updated records: %w", err)
	}

	return nil
}

func (j *jsonDAO[T]) Find(result *[]T) error {
	defer func() {
		j.conditions = []condition{}
	}()

	records, err := util.ReadJSON[T](j.filePath)
	if err != nil {
		return err
	}

	if len(j.conditions) == 0 {
		*result = records
		return nil
	}

	var filtered []T
	for _, record := range records {
		match := true
		for _, c := range j.conditions {
			fieldName := c.field
			expectedValue := c.value

			val := reflect.ValueOf(record)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			fieldVal := val.FieldByName(fieldName)
			if !fieldVal.IsValid() {
				return fmt.Errorf("field %s does not exist in type %T", fieldName, record)
			}

			if !reflect.DeepEqual(fieldVal.Interface(), expectedValue) {
				match = false
				break
			}
		}

		if match {
			filtered = append(filtered, record)
		}
	}

	*result = filtered
	return nil
}

func (j *jsonDAO[T]) First(result *T) error {
	defer func() {
		j.conditions = []condition{}
	}()

	if len(j.conditions) == 0 {
		return errors.New("no conditions specified")
	}

	records, err := util.ReadJSON[T](j.filePath)
	if err != nil {
		return err
	}

	for _, record := range records {
		matches := true

		for _, c := range j.conditions {
			fieldName := c.field
			expectedValue := c.value

			val := reflect.ValueOf(record)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}

			fieldVal := val.FieldByName(fieldName)
			if !fieldVal.IsValid() {
				return fmt.Errorf("field %s does not exist in type %T", fieldName, record)
			}

			if !reflect.DeepEqual(fieldVal.Interface(), expectedValue) {
				matches = false
				break
			}
		}

		if matches {
			*result = record
			return nil
		}
	}

	return errors.New("record not found")
}

func (j *jsonDAO[T]) Update(val *T) error {
	records, err := util.ReadJSON[T](j.filePath)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	idField := value.FieldByName("ID")
	if !idField.IsValid() {
		return errors.New("field 'ID' does not exist in the struct")
	}
	id := idField.Interface()

	updated := false
	for i := range records {
		recordVal := reflect.ValueOf(&records[i]).Elem()
		recordIDField := recordVal.FieldByName("ID")
		if !recordIDField.IsValid() {
			return errors.New("field 'ID' does not exist in JSON records")
		}

		if reflect.DeepEqual(recordIDField.Interface(), id) {
			records[i] = *val
			updated = true
			break
		}
	}

	if !updated {
		return errors.New("record with the specified ID not found")
	}

	if err := util.WriteJSON[T](j.filePath, records); err != nil {
		return err
	}

	return nil
}

func (j *jsonDAO[T]) Where(field string, val any) JsonDAO[T] {
	j.conditions = append(j.conditions, condition{
		field: field,
		value: val,
	})

	return j
}

type condition struct {
	field string
	value any
}

func NewJsonDAO[T any](filePath string) JsonDAO[T] {
	return &jsonDAO[T]{filePath, []condition{}}
}
