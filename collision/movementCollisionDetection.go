package collision

import "github.com/LuigiVanacore/ebiten_extended/math2D"

 
func MovingCircleCircleCollide( a, b math2D.Circle, moveA math2D.Vector2D) bool {
	bAbsorbedA := math2D.NewCircle( math2D.NewVector2D( b.GetCenterPosition().X(), b.GetCenterPosition().Y()), b.GetRadius() + a.GetRadius())
	travelA := math2D.NewSegment(a.GetCenterPosition(), math2D.AddVectors(a.GetCenterPosition(), moveA))
	return CircleSegmentCollide(bAbsorbedA, travelA)
} 
// Bool moving_rectangle_rectangle_collide(const Rectangle* a, const Vector2D* moveA, const Rectangle* b)
// {
// 	Rectangle envelope = *a;
// 	envelope.origin = add_vector(&envelope.origin, moveA);
// 	envelope = enlarge_rectangle_rectangle(&envelope, a);

// 	if(rectangles_collide(&envelope, b))
// 	{
// 		const float epsilon = 1.0f / 32.0f;/* improves case a->size is (close to) zero */
// 		const float minimumMoveDistance = maximum(minimum(a->size.x, a->size.y) / 4, epsilon);
// 		const Vector2D halfMoveA = divide_vector(moveA, 2);
// 		if(vector_length(moveA) < minimumMoveDistance)
// 			return yes;

// 		envelope.origin = add_vector(&a->origin, &halfMoveA);
// 		envelope.size = a->size;
// 		return
// 			moving_rectangle_rectangle_collide(a, &halfMoveA, b) ||
// 			moving_rectangle_rectangle_collide(&envelope, &halfMoveA, b);
// 	}
// 	else
// 		return no;
// }
func MovingRectangleRectangleCollide( a math2D.Rectangle, moveA math2D.Vector2D, b math2D.Rectangle) bool {
	
}
// Bool moving_circle_rectangle_collide(const Circle* a, const Vector2D* moveA, const Rectangle* b)
// {
// 	Circle envelope = *a;
// 	const Vector2D halfMoveA = divide_vector(moveA, 2);
// 	const float moveDistance = vector_length(moveA);
// 	envelope.center = add_vector(&a->center, &halfMoveA);
// 	envelope.radius = a->radius + moveDistance / 2;

// 	if(circle_rectangle_collide(&envelope, b))
// 	{
// 		const float epsilon = 1.0f / 32.0f;/* improves case a->radius is (close to) zero */
// 		const float minimumMoveDistance = maximum(a->radius / 4, epsilon);
// 		if(moveDistance < minimumMoveDistance)
// 			return yes;

// 		envelope.radius = a->radius;
// 		return
// 			moving_circle_rectangle_collide(a, &halfMoveA, b) ||
// 			moving_circle_rectangle_collide(&envelope, &halfMoveA, b);
// 	}
// 	else
// 		return no;
// }

// Bool moving_rectangle_circle_collide(const Rectangle* a, const Vector2D* moveA, const Circle* b)
// {
// 	const Vector2D moveB = negate_vector(moveA);
// 	return moving_circle_rectangle_collide(b, &moveB, a);
// }


