package hotels_test

import (
	"practices/hotels/hotels"
	"reflect"
	"testing"
)

func TestSortHotels(t *testing.T) {
	var testCases = []struct {
		keywords []string
		reviews  []string
		hotelIds []int
		expect   []int
	}{
		{
			[]string{"breakfast", "beach", "citycenter", "location", "metro", "view", "staff", "price"},
			[]string{"This hotel has a nice view of the citycenter. The location is perfect.", "The breakfast is ok. Regarding location, it is quite far from citycenter but price is cheap so it is worth.", "Location is excellent, 5 minutes from citycenter. There is also a metro station very close to the hotel.", "They said I couldn't take my dog and there were other guests with dogs! That is not fair.", "Very friendly staff and good cost-benefit ratio. Its location is a bit far from citycenter."},
			[]int{1, 2, 1, 1, 2},
			[]int{2, 1},
		},
	}
	for _, testCase := range testCases {
		result := hotels.SortHotels(testCase.keywords, testCase.reviews, testCase.hotelIds)
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
