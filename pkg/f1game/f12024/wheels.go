package f12024

const (
	WHEEL_INDEX_REAR_LEFT   = 0
	WHEEL_INDEX_REAR_RIGHT  = 1
	WHEEL_INDEX_FRONT_LEFT  = 2
	WHEEL_INDEX_FRONT_RIGHT = 3
)

type WheelMap[T any] [4]T

func (w *WheelMap[T]) RearLeft() T   { return w[WHEEL_INDEX_REAR_LEFT] }
func (w *WheelMap[T]) RearRight() T  { return w[WHEEL_INDEX_REAR_RIGHT] }
func (w *WheelMap[T]) FrontLeft() T  { return w[WHEEL_INDEX_FRONT_LEFT] }
func (w *WheelMap[T]) FrontRight() T { return w[WHEEL_INDEX_FRONT_RIGHT] }

func NewWheelMap[T any](rear_left, rear_right, front_left, front_right T) (result WheelMap[T]) {
	result[WHEEL_INDEX_REAR_LEFT] = rear_left
	result[WHEEL_INDEX_REAR_RIGHT] = rear_right
	result[WHEEL_INDEX_FRONT_LEFT] = front_left
	result[WHEEL_INDEX_FRONT_RIGHT] = front_right

	return
}
