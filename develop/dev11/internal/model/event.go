package model

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var Cache = newEventsCache()

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

type EventCache struct {
	events sync.Map
	nextID int
}

func newEventsCache() *EventCache {
	return &EventCache{
		nextID: 1,
	}
}

func (c *EventCache) CreateEvent(title, description string, start, end time.Time) (Event, error) {
	if start.After(end) {
		return Event{}, errors.New("start time have to be before end time")
	}

	event := Event{
		ID:          c.nextID,
		Title:       title,
		Description: description,
		Start:       start,
		End:         end,
	}

	c.events.Store(event.ID, event)
	c.nextID++

	return event, nil
}

func (c *EventCache) UpdateEvent(eventID int, updatedEvent Event) error {
	if _, exists := c.events.Load(eventID); !exists {
		return fmt.Errorf("event with ID %d does not exist", eventID)
	}
	c.events.Store(eventID, updatedEvent)
	return nil
}

func (c *EventCache) DeleteEvent(eventID int, updatedEvent Event) error {
	if _, exists := c.events.Load(eventID); !exists {
		return fmt.Errorf("event with ID %d does not exist", eventID)
	}
	c.events.Delete(eventID)
	return nil
}

func (c *EventCache) FindEventsByDate(currentTime time.Time, layout string) ([]Event, error) {
	var result []Event
	switch layout {
	case "day":
		Cache.events.Range(func(key, value interface{}) bool {
			event := value.(Event)
			if (event.Start.After(currentTime) || event.Start.Equal(currentTime)) && event.End.After(currentTime.Add(time.Hour*24)) {
				result = append(result, event)
			}
			return true
		})
	case "week":
		Cache.events.Range(func(key, value interface{}) bool {
			event := value.(Event)
			if event.Start.After(currentTime) && event.Start.Before(currentTime.Add(time.Hour*24*7)) {
				result = append(result, event)
			}
			return true
		})
	case "month":
		Cache.events.Range(func(key, value interface{}) bool {
			event := value.(Event)
			if event.Start.After(currentTime) && event.Start.Before(currentTime.Add(time.Hour*24*30)) {
				result = append(result, event)
			}
			return true
		})
	}
	return result, nil
}
