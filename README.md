Ball Clocks
-----------

Now there are a couple ways to look at ball clocks.
They can be treated as polynomials, much in the same
way 123 can, except they have a variable base, but a little
linear algebra plus some finite fields can come in handy.

A ballclock is defined by both it's state and it's queue size
It returns a result t where t is an element of the finite group
of Z(720) under the operation of addition.

The time t can be represented as a linear combination of the
individual clock elements. Let x1 be the minutes (in Z(5), x2 be the minutes  / 5 (in Z(30)) 
and x3 be the hours (in Z(12)). Then the time t in minutes is given as:

    t := x1 + 5*x2 + 60*x3

or, in simplified notation, 

    let x⃗ = <x1, x2, x3> and c⃗ = <1, 5, 60>,

then t is simply
    
    t = x⃗ · c⃗

Now let's bring into play the balls β The minimum number
of balls required to represent time is of course
12 + 55/5 + 4, or 26. This gives us a time of 12:59 when the static hour ball is included.
At 12:59, the system balls are all once again at the bottom in the queue.
Let's examine out loud the order the balls go through the system.

Every minute, the balls go from some configuration β to β′. There are a few observations we can make

1. For each ball in the queue, it's position is advanced +1 every minute (1,2,3,4,5)
2. For each ball in the minutes queue, it's position is advanced +1 unless
it is the 5th minute. in that case, 4 balls run down the track in reverse order back to the *beginning* of the
queue. Ball number 5 advances to the second track. This is called a carry
3. The 12th ball carried over from the minutes track, returns 11 balls to the queue in reverse order,
leaving only ball 12 to fall to the hours track
3. The 12th ball carried over to the hours track causes 11 balls to return to the queue in reverse order, and then
the 12th ball then falls falls into the to the queue.


Let's simulate:
2,3
(1,2,3,4,5,6) () ()
(1,2,3,4,5) (6) ()
(6,1,2,3,4) () (5)
(6,1,2,3) (4) (5)
(4,6,1,2) () (5,3)
(4,6,1) (2) (5,3)

(1,3,5,2,4,6) () ()
>>>>>>> Initial commit
